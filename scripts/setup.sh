#!/bin/bash
# NLG Blockchain Setup Script

set -e

echo "======================================"
echo "NLG Blockchain Setup"
echo "======================================"
echo ""

# Colors
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo -e "${RED}Go is not installed!${NC}"
    echo "Installing Go 1.24.1..."
    
    cd /tmp
    curl -OL https://go.dev/dl/go1.24.1.linux-amd64.tar.gz
    sudo rm -rf /usr/local/go
    sudo tar -C /usr/local -xzf go1.24.1.linux-amd64.tar.gz
    
    # Add to PATH
    echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
    echo 'export PATH=$PATH:$HOME/go/bin' >> ~/.bashrc
    export PATH=$PATH:/usr/local/go/bin
    export PATH=$PATH:$HOME/go/bin
    
    echo -e "${GREEN}Go installed successfully!${NC}"
else
    GO_VERSION=$(go version | awk '{print $3}')
    echo -e "${GREEN}Go is already installed: $GO_VERSION${NC}"
fi

# Check if Ignite CLI is installed
if ! command -v ignite &> /dev/null; then
    echo -e "${YELLOW}Ignite CLI is not installed!${NC}"
    echo "Installing Ignite CLI..."
    
    curl https://get.ignite.com/cli! | bash
    
    echo -e "${GREEN}Ignite CLI installed successfully!${NC}"
else
    IGNITE_VERSION=$(ignite version 2>&1 | head -1 || echo "unknown")
    echo -e "${GREEN}Ignite CLI is already installed: $IGNITE_VERSION${NC}"
fi

echo ""
echo "======================================"
echo "Building NLG Blockchain..."
echo "======================================"

# Navigate to project directory
cd "$(dirname "$0")/.."

# Install dependencies
echo "Installing Go dependencies..."
go mod tidy

# Build the chain
echo "Building chain binary..."
ignite chain build

# Add to PATH
export PATH=$PATH:$HOME/go/bin

# Verify installation
if command -v nlgd &> /dev/null; then
    echo ""
    echo -e "${GREEN}======================================"
    echo "Setup Complete!"
    echo "======================================${NC}"
    echo ""
    echo "NLG blockchain binary 'nlgd' is now available!"
    echo ""
    echo "Next steps:"
    echo "  1. Add to your PATH: export PATH=\$PATH:\$HOME/go/bin"
    echo "  2. Start development chain: ignite chain serve"
    echo "  3. Or start manually: nlgd start"
    echo ""
    echo "For more information, see README.md"
else
    echo -e "${RED}Error: nlgd binary not found after build${NC}"
    exit 1
fi
