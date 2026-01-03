# NLG Blockchain Deployment Guide

Complete guide for deploying the NLG blockchain in various environments.

## Table of Contents

1. [Local Development](#local-development)
2. [Single Node Deployment](#single-node-deployment)
3. [Multi-Node Testnet](#multi-node-testnet)
4. [Production Deployment](#production-deployment)
5. [Cloud Deployment](#cloud-deployment)
6. [Monitoring and Maintenance](#monitoring-and-maintenance)

## Local Development

### Quick Start

```bash
# Start development chain with hot reload
ignite chain serve

# Or start manually
ignite chain build
nlgd start
```

### Development Configuration

The dev chain uses `config.yml` for configuration:

```yaml
version: 1
validation: sovereign
default_denom: nlg
accounts: 
- name: alice
  coins: [40000000nlg]
- name: bob
  coins: [30000000nlg]
- name: charlie
  coins: [30000000nlg]
validators:
- name: alice
  bonded: 20000000nlg
```

### Development Endpoints

- RPC: http://localhost:26657
- API: http://localhost:1317
- gRPC: localhost:9090
- P2P: localhost:26656

## Single Node Deployment

### Prerequisites

- Ubuntu 20.04+ or similar Linux distribution
- 4GB+ RAM
- 100GB+ disk space
- Open ports: 26656 (P2P), 26657 (RPC), 1317 (API)

### Step 1: Install Binary

```bash
# Build from source
git clone <repository-url>
cd nlg
make install

# Verify installation
nlgd version
```

### Step 2: Initialize Node

```bash
# Set chain ID
CHAIN_ID="nlg-1"
MONIKER="my-node"

# Initialize node
nlgd init $MONIKER --chain-id $CHAIN_ID

# This creates ~/.nlg/ directory with:
# - config/config.toml (Tendermint config)
# - config/app.toml (Application config)
# - config/genesis.json (Genesis state)
# - data/ (Blockchain data)
```

### Step 3: Configure Genesis

```bash
# Create genesis accounts
nlgd genesis add-genesis-account alice 40000000nlg
nlgd genesis add-genesis-account bob 30000000nlg
nlgd genesis add-genesis-account charlie 30000000nlg

# Configure genesis parameters
cat <<EOF > genesis_patch.json
{
  "app_state": {
    "staking": {
      "params": {
        "bond_denom": "nlg",
        "unbonding_time": "1814400s"
      }
    },
    "crisis": {
      "constant_fee": {
        "denom": "nlg",
        "amount": "1000"
      }
    },
    "gov": {
      "deposit_params": {
        "min_deposit": [
          {
            "denom": "nlg",
            "amount": "10000000"
          }
        ]
      }
    }
  }
}
EOF

# Apply patch (use jq)
jq -s '.[0] * .[1]' ~/.nlg/config/genesis.json genesis_patch.json > genesis_new.json
mv genesis_new.json ~/.nlg/config/genesis.json
```

### Step 4: Create Validator

```bash
# Create validator key
nlgd keys add validator

# Create genesis transaction
nlgd genesis gentx validator 20000000nlg \
  --chain-id $CHAIN_ID \
  --moniker $MONIKER \
  --commission-rate 0.1 \
  --commission-max-rate 0.2 \
  --commission-max-change-rate 0.01 \
  --min-self-delegation 1

# Collect genesis transactions
nlgd genesis collect-gentxs

# Validate genesis
nlgd genesis validate-genesis
```

### Step 5: Configure Node

Edit `~/.nlg/config/config.toml`:

```toml
# RPC Server Configuration
[rpc]
laddr = "tcp://0.0.0.0:26657"
cors_allowed_origins = ["*"]

# P2P Configuration
[p2p]
laddr = "tcp://0.0.0.0:26656"
persistent_peers = ""
max_num_inbound_peers = 40
max_num_outbound_peers = 10

# Consensus Configuration
[consensus]
timeout_commit = "5s"
```

Edit `~/.nlg/config/app.toml`:

```toml
# API Configuration
[api]
enable = true
swagger = true
address = "tcp://0.0.0.0:1317"

# gRPC Configuration
[grpc]
enable = true
address = "0.0.0.0:9090"

# State Sync
[state-sync]
snapshot-interval = 1000
snapshot-keep-recent = 2
```

### Step 6: Create Systemd Service

```bash
sudo tee /etc/systemd/system/nlgd.service > /dev/null <<EOF
[Unit]
Description=NLG Blockchain Daemon
After=network-online.target

[Service]
User=$USER
ExecStart=$(which nlgd) start
Restart=on-failure
RestartSec=3
LimitNOFILE=65535

[Install]
WantedBy=multi-user.target
EOF

# Enable and start service
sudo systemctl daemon-reload
sudo systemctl enable nlgd
sudo systemctl start nlgd

# Check status
sudo systemctl status nlgd

# View logs
sudo journalctl -u nlgd -f
```

## Multi-Node Testnet

### Using Ignite CLI

```bash
# Generate 4-node testnet
ignite chain testnet \
  --validator-count 4 \
  --chain-id nlg-testnet-1 \
  --output-dir ./testnet

# This creates:
# testnet/
#   node0/
#   node1/
#   node2/
#   node3/
```

### Manual Multi-Node Setup

#### On Node 1 (Genesis Node)

```bash
# Initialize
nlgd init node1 --chain-id nlg-testnet-1

# Create validator key
nlgd keys add validator1

# Add genesis accounts (repeat for each validator)
nlgd genesis add-genesis-account $(nlgd keys show validator1 -a) 50000000nlg

# Create gentx
nlgd genesis gentx validator1 25000000nlg --chain-id nlg-testnet-1

# Get node ID
NODE1_ID=$(nlgd tendermint show-node-id)
NODE1_IP="<public-ip>"
```

#### On Additional Nodes

```bash
# Initialize
nlgd init node2 --chain-id nlg-testnet-1

# Copy genesis.json from Node 1
scp user@node1:~/.nlg/config/genesis.json ~/.nlg/config/

# Set persistent peers
PEERS="${NODE1_ID}@${NODE1_IP}:26656"
sed -i "s/persistent_peers = \"\"/persistent_peers = \"${PEERS}\"/" ~/.nlg/config/config.toml

# Start node
nlgd start
```

#### Collect Genesis Transactions

On Node 1:

```bash
# Collect gentx files from all validators
scp user@node2:~/.nlg/config/gentx/* ~/.nlg/config/gentx/
scp user@node3:~/.nlg/config/gentx/* ~/.nlg/config/gentx/

# Collect all gentxs
nlgd genesis collect-gentxs

# Validate
nlgd genesis validate-genesis

# Share final genesis with all nodes
for node in node2 node3 node4; do
  scp ~/.nlg/config/genesis.json user@$node:~/.nlg/config/
done
```

## Production Deployment

### Architecture

```
                    ┌─────────────┐
                    │   Load      │
                    │   Balancer  │
                    └──────┬──────┘
                           │
        ┌──────────────────┼──────────────────┐
        │                  │                  │
   ┌────▼────┐       ┌────▼────┐       ┌────▼────┐
   │ Sentry  │       │ Sentry  │       │ Sentry  │
   │ Node 1  │       │ Node 2  │       │ Node 3  │
   └────┬────┘       └────┬────┘       └────┬────┘
        │                  │                  │
        └──────────────────┼──────────────────┘
                           │
                    ┌──────▼──────┐
                    │  Validator  │
                    │    Node     │
                    └─────────────┘
```

### Sentry Node Configuration

```toml
# config.toml
[p2p]
# Disable RPC
pex = true
private_peer_ids = ""
unconditional_peer_ids = "<validator-node-id>"

# Only connect to validator
persistent_peers = "<validator-node-id>@<validator-private-ip>:26656"

[rpc]
# Enable public RPC
laddr = "tcp://0.0.0.0:26657"
```

### Validator Node Configuration

```toml
# config.toml
[p2p]
# Only connect to sentry nodes
pex = false
persistent_peers = "<sentry1-id>@<sentry1-ip>:26656,<sentry2-id>@<sentry2-ip>:26656"
private_peer_ids = ""

# Disable public RPC
[rpc]
laddr = "tcp://127.0.0.1:26657"
```

### Security Hardening

```bash
# 1. Firewall rules (validator)
sudo ufw default deny incoming
sudo ufw default allow outgoing
sudo ufw allow from <sentry1-ip> to any port 26656
sudo ufw allow from <sentry2-ip> to any port 26656
sudo ufw allow 22/tcp  # SSH
sudo ufw enable

# 2. Firewall rules (sentry)
sudo ufw default deny incoming
sudo ufw default allow outgoing
sudo ufw allow 26656/tcp  # P2P
sudo ufw allow 26657/tcp  # RPC
sudo ufw allow 1317/tcp   # API
sudo ufw allow 22/tcp     # SSH
sudo ufw enable

# 3. Fail2ban for SSH
sudo apt install fail2ban
sudo systemctl enable fail2ban

# 4. Disable root login
sudo sed -i 's/PermitRootLogin yes/PermitRootLogin no/' /etc/ssh/sshd_config
sudo systemctl restart sshd

# 5. Key management (use HSM for production)
# Store validator keys in encrypted location
# Consider using Tendermint KMS
```

### Backup Strategy

```bash
# 1. Backup validator key
cp ~/.nlg/config/priv_validator_key.json ~/validator_key_backup.json
# Store securely offline

# 2. Backup node key
cp ~/.nlg/config/node_key.json ~/node_key_backup.json

# 3. Automated state backup
#!/bin/bash
BACKUP_DIR="/backup/nlg"
DATE=$(date +%Y%m%d_%H%M%S)

# Stop node
sudo systemctl stop nlgd

# Backup data directory
tar -czf $BACKUP_DIR/nlg_data_$DATE.tar.gz ~/.nlg/data

# Restart node
sudo systemctl start nlgd

# Keep last 7 days
find $BACKUP_DIR -name "nlg_data_*.tar.gz" -mtime +7 -delete
```

## Cloud Deployment

### AWS Deployment

#### Using EC2

```bash
# 1. Launch EC2 instance
# Instance type: t3.large (2 vCPU, 8GB RAM)
# Storage: 100GB gp3 SSD
# AMI: Ubuntu Server 20.04 LTS

# 2. Configure security group
# Inbound rules:
# - SSH (22): Your IP
# - P2P (26656): 0.0.0.0/0
# - RPC (26657): Your IP or load balancer
# - API (1317): Your IP or load balancer

# 3. Install and configure node (see single node deployment)

# 4. Setup CloudWatch logging
aws logs create-log-group --log-group-name /nlg/node
aws logs create-log-stream --log-group-name /nlg/node --log-stream-name node1

# Install CloudWatch agent
wget https://s3.amazonaws.com/amazoncloudwatch-agent/ubuntu/amd64/latest/amazon-cloudwatch-agent.deb
sudo dpkg -i amazon-cloudwatch-agent.deb

# Configure agent to send nlgd logs
```

#### Using ECS/EKS

See `kubernetes/` directory for Kubernetes manifests.

### GCP Deployment

```bash
# Create compute instance
gcloud compute instances create nlg-validator \
  --machine-type=e2-standard-2 \
  --zone=us-central1-a \
  --image-family=ubuntu-2004-lts \
  --image-project=ubuntu-os-cloud \
  --boot-disk-size=100GB \
  --boot-disk-type=pd-ssd

# Setup firewall rules
gcloud compute firewall-rules create nlg-p2p \
  --allow tcp:26656 \
  --source-ranges 0.0.0.0/0

# SSH into instance
gcloud compute ssh nlg-validator --zone=us-central1-a

# Follow single node deployment steps
```

### DigitalOcean Deployment

```bash
# Create droplet via API
curl -X POST "https://api.digitalocean.com/v2/droplets" \
  -H "Authorization: Bearer $DO_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "nlg-node",
    "region": "nyc3",
    "size": "s-2vcpu-4gb",
    "image": "ubuntu-20-04-x64",
    "ssh_keys": ["<ssh-key-id>"],
    "backups": true,
    "monitoring": true
  }'

# Follow single node deployment steps
```

## Monitoring and Maintenance

### Prometheus Metrics

```yaml
# prometheus.yml
global:
  scrape_interval: 15s

scrape_configs:
  - job_name: 'nlg-node'
    static_configs:
      - targets: ['localhost:26660']
```

### Grafana Dashboard

```json
{
  "dashboard": {
    "title": "NLG Blockchain",
    "panels": [
      {
        "title": "Block Height",
        "targets": [{"expr": "tendermint_consensus_height"}]
      },
      {
        "title": "Transactions",
        "targets": [{"expr": "rate(tendermint_consensus_total_txs[5m])"}]
      },
      {
        "title": "Peers",
        "targets": [{"expr": "tendermint_p2p_peers"}]
      }
    ]
  }
}
```

### Health Checks

```bash
#!/bin/bash
# health_check.sh

# Check if node is running
if ! systemctl is-active --quiet nlgd; then
    echo "ERROR: nlgd service is not running"
    exit 1
fi

# Check if node is synced
CATCHING_UP=$(curl -s http://localhost:26657/status | jq -r '.result.sync_info.catching_up')
if [ "$CATCHING_UP" = "true" ]; then
    echo "WARNING: Node is still syncing"
fi

# Check block height
HEIGHT=$(curl -s http://localhost:26657/status | jq -r '.result.sync_info.latest_block_height')
echo "Current block height: $HEIGHT"

# Check validator status
VOTING_POWER=$(nlgd query staking validator $(nlgd keys show validator -a --bech val) -o json | jq -r '.tokens')
echo "Validator voting power: $VOTING_POWER"
```

### Automated Updates

```bash
#!/bin/bash
# update_node.sh

# Backup current version
cp $(which nlgd) ~/nlgd_backup_$(date +%Y%m%d)

# Stop node
sudo systemctl stop nlgd

# Update binary
cd ~/nlg
git pull
make install

# Restart node
sudo systemctl start nlgd

# Monitor logs
journalctl -u nlgd -f -n 100
```

### Log Rotation

```bash
# /etc/logrotate.d/nlgd
/var/log/nlgd/*.log {
    daily
    rotate 14
    compress
    delaycompress
    notifempty
    create 0640 nlg nlg
    sharedscripts
    postrotate
        systemctl reload nlgd > /dev/null 2>&1 || true
    endscript
}
```

## Troubleshooting

### Node won't sync

```bash
# Reset data and resync from genesis
nlgd tendermint unsafe-reset-all

# Or use state sync
# Edit config.toml:
[statesync]
enable = true
rpc_servers = "https://rpc1.nlg.network:26657,https://rpc2.nlg.network:26657"
trust_height = <trusted-height>
trust_hash = "<trusted-hash>"
```

### High memory usage

```bash
# Increase Go GC frequency
export GOGC=50
systemctl restart nlgd
```

### Database corruption

```bash
# Restore from backup
sudo systemctl stop nlgd
rm -rf ~/.nlg/data
tar -xzf /backup/nlg_data_YYYYMMDD.tar.gz -C ~/
sudo systemctl start nlgd
```

## Support

For deployment assistance:
- Discord: #deployment-support
- Documentation: https://docs.nlg.network
- Email: support@nlg.network

---

Last updated: 2024
