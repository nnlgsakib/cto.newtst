# NLG Blockchain

A complete Cosmos SDK blockchain with Ethereum compatibility and distributed social media capabilities.

## Overview

NLG is a fully-functional blockchain built using Ignite CLI and the Cosmos SDK. It features:

- **Native Token**: NLG with 100M total supply
- **Chain ID**: nlg-1
- **Address Prefix**: nlg
- **Consensus**: Tendermint BFT
- **Standard Cosmos Modules**: bank, auth, staking, distribution, slashing, governance, feegrant, authz
- **Custom Social Media Module**: Decentralized social networking capabilities

## Token Economics

- **Total Supply**: 100,000,000 NLG
- **Distribution**:
  - Alice: 40,000,000 NLG (40%)
  - Bob: 30,000,000 NLG (30%)
  - Charlie: 30,000,000 NLG (30%)
- **Validator Stake**: 20,000,000 NLG bonded by Alice

## Social Media Module

The custom `socialmedia` module provides decentralized social networking features:

### Features

1. **Posts**
   - Create, update, and delete posts
   - Store content with IPFS hash references
   - Track likes and engagement
   - Timestamp tracking

2. **User Profiles**
   - Username management
   - Bio and avatar (IPFS)
   - Reputation system
   - Follower/following counts

3. **Comments**
   - Comment on posts
   - Nested discussions
   - Author attribution
   - Timestamp tracking

4. **Social Connections**
   - Follow/unfollow users
   - Track follower relationships
   - Connection timestamps

5. **Content Moderation**
   - Community-driven moderation
   - Voting mechanism
   - Stake-weighted decisions

### Module Dependencies

The social media module integrates with:
- **Bank Module**: For token transfers and rewards
- **Staking Module**: For reputation and moderation weight

## Prerequisites

- Go 1.24.1 or later
- Ignite CLI v29.6.2 or later

## Installation

### Install Go

```bash
curl -OL https://go.dev/dl/go1.24.1.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.24.1.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin
```

### Install Ignite CLI

```bash
curl https://get.ignite.com/cli! | bash
```

### Build the Chain

```bash
# Clone the repository
git clone <repository-url>
cd nlg

# Install dependencies
go mod tidy

# Build the chain
ignite chain build
```

## Running the Chain

### Start Local Development Chain

```bash
ignite chain serve
```

This command will:
- Build the chain binary (`nlgd`)
- Initialize a development chain
- Create validator accounts (alice, bob, charlie)
- Start the chain with live reloading

The chain will be available at:
- RPC: `http://localhost:26657`
- API: `http://localhost:1317`
- gRPC: `localhost:9090`
- **Web Dashboard**: `http://localhost:1317/dashboard`

### Chain Commands

Once the chain is running, you can interact with it using the `nlgd` CLI:

```bash
# Query chain status
nlgd status

# Query balances
nlgd query bank balances $(nlgd keys show alice -a)

# Create a post
nlgd tx socialmedia create-post "My First Post" \
  "This is the content of my post" \
  "QmXXXIPFSHash" \
  $(date +%s) \
  0 \
  --from alice

# Query all posts
nlgd query socialmedia list-post

# Create a profile
nlgd tx socialmedia create-profile "alice" \
  "Alice's bio" \
  "QmYYYAvatarHash" \
  100 \
  0 \
  0 \
  --from alice

# Query profiles
nlgd query socialmedia list-profile

# Follow a user
nlgd tx socialmedia create-social-connection \
  $(nlgd keys show alice -a) \
  $(nlgd keys show bob -a) \
  $(date +%s) \
  --from alice

# Like a post
nlgd tx socialmedia like-post "post-id" --from bob

# Moderate content
nlgd tx socialmedia moderate-content "content-id" "flag" --from charlie
```

## Web Dashboard

NLG Chain includes a comprehensive web dashboard that provides real-time monitoring and interaction capabilities.

### Features

- **Dashboard Overview**: Real-time statistics (block height, validators, token supply)
- **Block Explorer**: Browse blocks and transactions
- **Validator Monitor**: View active validators and their status
- **Social Feed**: Real-time feed of decentralized social media posts
- **API Documentation**: Interactive Swagger UI for all API endpoints

### Accessing the Dashboard

Start the chain and open your browser:
```bash
ignite chain serve
```

Then navigate to:
```
http://localhost:1317/dashboard
```

### Dashboard Pages

1. **Main Dashboard** (`/dashboard#dashboard`)
   - Network statistics and metrics
   - Recent blocks and social posts

2. **Blocks Explorer** (`/dashboard#blocks`)
   - List of latest blocks
   - Detailed block information

3. **Validators** (`/dashboard#validators`)
   - Active validator list
   - Staking and commission info

4. **Social Feed** (`/dashboard#social`)
   - Decentralized social media posts
   - Real-time content updates

5. **API Documentation** (`/dashboard#api`)
   - Full API reference
   - Interactive testing

For more details on the web UI, see [docs/UI_README.md](docs/UI_README.md).

## Architecture

### Module Structure

```
x/socialmedia/
├── keeper/          # State management and business logic
├── types/           # Type definitions and protobuf messages
├── module/          # Module initialization and registration
└── simulation/      # Simulation tests
```

### Data Models

#### Post
```protobuf
message Post {
  string id = 1;
  string title = 2;
  string content = 3;
  string ipfsHash = 4;
  uint64 timestamp = 5;
  uint64 likesCount = 6;
  string creator = 7;
}
```

#### Profile
```protobuf
message Profile {
  string id = 1;
  string username = 2;
  string bio = 3;
  string avatarIpfsHash = 4;
  uint64 reputation = 5;
  uint64 followersCount = 6;
  uint64 followingCount = 7;
  string creator = 8;
}
```

#### Comment
```protobuf
message Comment {
  string id = 1;
  string postId = 2;
  string content = 3;
  uint64 timestamp = 4;
  string author = 5;
  string creator = 6;
}
```

#### SocialConnection
```protobuf
message SocialConnection {
  string id = 1;
  string follower = 2;
  string following = 3;
  uint64 timestamp = 4;
  string creator = 5;
}
```

## API Endpoints

### REST API

The chain exposes a REST API at `http://localhost:1317`:

- `GET /nlg/socialmedia/post` - List all posts
- `GET /nlg/socialmedia/post/{id}` - Get post by ID
- `GET /nlg/socialmedia/profile` - List all profiles
- `GET /nlg/socialmedia/profile/{id}` - Get profile by ID
- `GET /nlg/socialmedia/comment` - List all comments
- `GET /nlg/socialmedia/social-connection` - List all connections

### gRPC

gRPC services are available at `localhost:9090` for all query and transaction methods.

## Development

### Running Tests

```bash
# Run all tests
go test ./...

# Run module-specific tests
go test ./x/socialmedia/...

# Run with coverage
go test -cover ./x/socialmedia/...
```

### Code Generation

```bash
# Generate protobuf files
ignite generate proto-go

# Generate OpenAPI documentation
ignite generate openapi
```

### Linting

```bash
# Run linter
golangci-lint run

# Format code
gofmt -w .
```

## EVM Integration (Future Enhancement)

To add Ethereum compatibility via Ethermint:

1. Add Ethermint as a dependency in `go.mod`
2. Import EVM module in `app/app.go`
3. Configure EVM module in `app/app_config.go`
4. Add EVM-specific genesis parameters
5. Enable Web3 API endpoints

Example:
```go
import (
    evmkeeper "github.com/evmos/ethermint/x/evm/keeper"
    evmtypes "github.com/evmos/ethermint/x/evm/types"
)
```

## Genesis Configuration

The genesis file is automatically generated based on `config.yml`:

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

## Deployment

### Single Node

```bash
# Initialize chain
nlgd init my-node --chain-id nlg-1

# Add genesis accounts
nlgd genesis add-genesis-account alice 40000000nlg
nlgd genesis add-genesis-account bob 30000000nlg
nlgd genesis add-genesis-account charlie 30000000nlg

# Create validator
nlgd genesis gentx alice 20000000nlg --chain-id nlg-1

# Collect genesis transactions
nlgd genesis collect-gentxs

# Start the chain
nlgd start
```

### Multi-Node Testnet

```bash
# Use Ignite testnet command
ignite chain testnet --validator-count 3 --chain-id nlg-testnet
```

## Monitoring

The chain exposes Prometheus metrics at `http://localhost:26660/metrics` for monitoring:

- Block height
- Transaction throughput
- Validator performance
- Module-specific metrics

## Security Considerations

1. **Key Management**: Store validator keys securely
2. **Firewall**: Restrict P2P and RPC ports
3. **DDoS Protection**: Use rate limiting and proper network configuration
4. **Content Moderation**: Implement proper staking requirements for moderation
5. **IPFS Content**: Validate and sanitize content references

## Troubleshooting

### Chain Won't Start

```bash
# Reset chain data
nlgd tendermint unsafe-reset-all

# Remove old data
rm -rf ~/.nlg
```

### Module Errors

```bash
# Regenerate proto files
ignite generate proto-go

# Rebuild
ignite chain build --rebuild
```

### Port Conflicts

Edit `config.toml` and `app.toml` in `~/.nlg/config/` to change ports.

## Resources

- [Cosmos SDK Documentation](https://docs.cosmos.network/)
- [Ignite CLI Documentation](https://docs.ignite.com/)
- [Tendermint Core](https://docs.tendermint.com/)
- [IPFS Documentation](https://docs.ipfs.io/)

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Submit a pull request

## License

This project is licensed under the Apache 2.0 License.

## Support

For issues and questions:
- Open an issue on GitHub
- Join our Discord community
- Check the documentation

---

Built with ❤️ using Cosmos SDK and Ignite CLI
