# NLG Blockchain - Quick Start Guide

Get up and running with the NLG blockchain in 5 minutes!

## Prerequisites

- Linux or macOS
- 4GB RAM minimum
- 10GB free disk space

## Option 1: Automated Setup (Recommended)

```bash
# Run the setup script
./scripts/setup.sh

# Add binary to PATH
export PATH=$PATH:$HOME/go/bin

# Start the development chain
ignite chain serve
```

That's it! The chain is now running with:
- RPC: http://localhost:26657
- API: http://localhost:1317
- gRPC: localhost:9090

## Option 2: Manual Setup

### Step 1: Install Go

```bash
curl -OL https://go.dev/dl/go1.24.1.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.24.1.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin:$HOME/go/bin
```

### Step 2: Install Ignite CLI

```bash
curl https://get.ignite.com/cli! | bash
```

### Step 3: Build the Chain

```bash
# Install dependencies
go mod tidy

# Build
ignite chain build

# Verify
nlgd version
```

### Step 4: Start the Chain

```bash
ignite chain serve
```

## Quick Examples

### Create a Post

```bash
nlgd tx socialmedia create-post \
  "My First Post" \
  "Hello, decentralized world!" \
  "QmExampleHash" \
  $(date +%s) \
  0 \
  --from alice \
  --chain-id nlg-1 \
  --yes
```

### Query Posts

```bash
# CLI
nlgd query socialmedia list-post

# REST API
curl http://localhost:1317/nlg/socialmedia/post
```

### Create Profile

```bash
nlgd tx socialmedia create-profile \
  "alice_nlg" \
  "Blockchain developer" \
  "QmAvatarHash" \
  100 \
  0 \
  0 \
  --from alice \
  --yes
```

### Follow a User

```bash
nlgd tx socialmedia create-social-connection \
  $(nlgd keys show alice -a) \
  $(nlgd keys show bob -a) \
  $(date +%s) \
  --from alice \
  --yes
```

### Like a Post

```bash
nlgd tx socialmedia like-post "0" --from bob --yes
```

## Available Accounts

The development chain comes with pre-configured accounts:

- **alice**: 40,000,000 nlg (Validator)
- **bob**: 30,000,000 nlg
- **charlie**: 30,000,000 nlg

## Default Ports

| Service | Port |
|---------|------|
| RPC | 26657 |
| API | 1317 |
| gRPC | 9090 |
| P2P | 26656 |

## Stop the Chain

Press `Ctrl+C` in the terminal running `ignite chain serve`

## Reset Chain Data

```bash
nlgd tendermint unsafe-reset-all
```

Or remove the data directory:

```bash
rm -rf ~/.nlg
```

## Next Steps

1. **Read the Full Documentation**: See [README.md](README.md)
2. **Explore the API**: Check [docs/API.md](docs/API.md)
3. **Deploy to Testnet**: Follow [docs/DEPLOYMENT.md](docs/DEPLOYMENT.md)
4. **Add EVM Support**: See [docs/EVM_INTEGRATION.md](docs/EVM_INTEGRATION.md)
5. **Start Developing**: Read [docs/DEVELOPMENT.md](docs/DEVELOPMENT.md)

## Troubleshooting

### Port Already in Use

```bash
# Find process using port 26657
lsof -i :26657

# Kill the process
kill -9 <PID>
```

### Chain Won't Start

```bash
# Reset everything
rm -rf ~/.nlg
ignite chain serve
```

### Module Not Found Error

```bash
go mod tidy
go clean -modcache
ignite chain build
```

## Get Help

- **Documentation**: Check the `/docs` directory
- **Issues**: Open a GitHub issue
- **Discord**: Join our community server

## What's Next?

The NLG blockchain is now running! Here are some ideas:

- Build a frontend app using the REST API
- Deploy smart contracts (after EVM integration)
- Set up a multi-node testnet
- Integrate with IPFS for content storage
- Create custom governance proposals

Happy building! 🚀
