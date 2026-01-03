# EVM Integration Guide for NLG Blockchain

This guide explains how to add Ethereum Virtual Machine (EVM) support to the NLG blockchain using Ethermint.

## Overview

Ethermint enables running Ethereum smart contracts on Cosmos SDK chains, providing:
- Full Ethereum compatibility
- Web3 JSON-RPC support
- Solidity smart contract deployment
- MetaMask integration
- Cross-chain compatibility with IBC

## Prerequisites

- NLG blockchain installed and running
- Go 1.24.1+
- Node.js and npm (for testing with Hardhat/Truffle)

## Step 1: Add Ethermint Dependencies

Update `go.mod` to include Ethermint:

```go
require (
    github.com/evmos/ethermint v0.23.0
    github.com/ethereum/go-ethereum v1.13.0
)
```

Run:
```bash
go mod tidy
```

## Step 2: Import EVM Module

Edit `app/app.go` to import Ethermint modules:

```go
import (
    // ... existing imports
    
    evmkeeper "github.com/evmos/ethermint/x/evm/keeper"
    evmtypes "github.com/evmos/ethermint/x/evm/types"
    evmrest "github.com/evmos/ethermint/x/evm/client/rest"
    
    feemarket "github.com/evmos/ethermint/x/feemarket"
    feemarketkeeper "github.com/evmos/ethermint/x/feemarket/keeper"
    feemarkettypes "github.com/evmos/ethermint/x/feemarket/types"
)
```

## Step 3: Add Module Keepers

In the `App` struct, add EVM and fee market keepers:

```go
type App struct {
    // ... existing keepers
    
    EvmKeeper       *evmkeeper.Keeper
    FeeMarketKeeper feemarketkeeper.Keeper
}
```

## Step 4: Initialize Keepers

In `NewApp()` function:

```go
// Initialize fee market keeper
app.FeeMarketKeeper = feemarketkeeper.NewKeeper(
    appCodec,
    keys[feemarkettypes.StoreKey],
    app.GetSubspace(feemarkettypes.ModuleName),
    nil,
)

// Initialize EVM keeper
app.EvmKeeper = evmkeeper.NewKeeper(
    appCodec,
    keys[evmtypes.StoreKey],
    tkeys[evmtypes.TransientKey],
    app.GetSubspace(evmtypes.ModuleName),
    app.AccountKeeper,
    app.BankKeeper,
    app.StakingKeeper,
    &app.FeeMarketKeeper,
    nil,
    geth.NewEVM,
    "",
)
```

## Step 5: Register Modules

Add modules to the module manager:

```go
app.mm = module.NewManager(
    // ... existing modules
    evm.NewAppModule(app.EvmKeeper, app.AccountKeeper, app.GetSubspace(evmtypes.ModuleName)),
    feemarket.NewAppModule(app.FeeMarketKeeper, app.GetSubspace(feemarkettypes.ModuleName)),
)
```

## Step 6: Configure Genesis

Update `app/genesis.go` to include EVM genesis state:

```go
func (app *App) DefaultGenesis() map[string]json.RawMessage {
    genesis := app.BasicModuleManager.DefaultGenesis(app.appCodec)
    
    // EVM genesis
    evmGenesis := evmtypes.DefaultGenesisState()
    evmGenesis.Params.EvmDenom = "nlg"
    evmGenesis.Params.EnableCreate = true
    evmGenesis.Params.EnableCall = true
    genesis[evmtypes.ModuleName] = app.appCodec.MustMarshalJSON(evmGenesis)
    
    // Fee market genesis
    feeMarketGenesis := feemarkettypes.DefaultGenesisState()
    genesis[feemarkettypes.ModuleName] = app.appCodec.MustMarshalJSON(feeMarketGenesis)
    
    return genesis
}
```

## Step 7: Enable Web3 API

Create `cmd/nlgd/cmd/web3.go`:

```go
package cmd

import (
    "github.com/cosmos/cosmos-sdk/server"
    "github.com/evmos/ethermint/server/config"
    "github.com/spf13/cobra"
)

func addWeb3Flags(cmd *cobra.Command) {
    cmd.Flags().String("json-rpc.address", "0.0.0.0:8545", "the JSON-RPC server address to listen on")
    cmd.Flags().String("json-rpc.ws-address", "0.0.0.0:8546", "the JSON-RPC WS server address to listen on")
    cmd.Flags().Bool("json-rpc.enable", true, "enable the JSON-RPC server")
    cmd.Flags().String("json-rpc.api", "eth,net,web3,debug,txpool", "API's offered over the JSON-RPC interface")
}
```

## Step 8: Update App Configuration

Edit `app/config.go`:

```go
func initAppConfig() (string, interface{}) {
    type CustomAppConfig struct {
        serverconfig.Config
        EVM     config.EVMConfig     `mapstructure:"evm"`
        JSONRPC config.JSONRPCConfig `mapstructure:"json-rpc"`
    }

    customAppConfig := CustomAppConfig{
        Config: *serverconfig.DefaultConfig(),
        EVM: config.DefaultEVMConfig(),
        JSONRPC: config.JSONRPCConfig{
            Enable:   true,
            Address:  "0.0.0.0:8545",
            WSAddress: "0.0.0.0:8546",
            API:      []string{"eth", "net", "web3", "txpool", "debug"},
        },
    }

    customAppTemplate := serverconfig.DefaultConfigTemplate + config.DefaultEVMConfigTemplate

    return customAppTemplate, customAppConfig
}
```

## Step 9: Build and Test

```bash
# Rebuild the chain
ignite chain build

# Start with EVM enabled
nlgd start --json-rpc.enable --json-rpc.api eth,net,web3
```

## Testing EVM Integration

### Using curl

```bash
# Get chain ID
curl -X POST --data '{"jsonrpc":"2.0","method":"eth_chainId","params":[],"id":1}' \
  -H "Content-Type: application/json" http://localhost:8545

# Get latest block
curl -X POST --data '{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":1}' \
  -H "Content-Type: application/json" http://localhost:8545
```

### Using MetaMask

1. Open MetaMask
2. Add custom network:
   - Network Name: NLG Chain
   - RPC URL: http://localhost:8545
   - Chain ID: (from eth_chainId)
   - Currency Symbol: NLG

### Deploy Smart Contract with Hardhat

Create `hardhat.config.js`:

```javascript
require("@nomiclabs/hardhat-waffle");

module.exports = {
  solidity: "0.8.19",
  networks: {
    nlg: {
      url: "http://localhost:8545",
      accounts: ["<private-key>"],
      chainId: 9000
    }
  }
};
```

Example contract `contracts/SocialToken.sol`:

```solidity
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

contract SocialToken {
    string public name = "Social Media Token";
    string public symbol = "SMT";
    uint256 public totalSupply = 1000000 * 10**18;
    
    mapping(address => uint256) public balanceOf;
    mapping(address => mapping(address => uint256)) public allowance;
    
    event Transfer(address indexed from, address indexed to, uint256 value);
    event Approval(address indexed owner, address indexed spender, uint256 value);
    
    constructor() {
        balanceOf[msg.sender] = totalSupply;
    }
    
    function transfer(address to, uint256 value) public returns (bool) {
        require(balanceOf[msg.sender] >= value, "Insufficient balance");
        balanceOf[msg.sender] -= value;
        balanceOf[to] += value;
        emit Transfer(msg.sender, to, value);
        return true;
    }
    
    function approve(address spender, uint256 value) public returns (bool) {
        allowance[msg.sender][spender] = value;
        emit Approval(msg.sender, spender, value);
        return true;
    }
    
    function transferFrom(address from, address to, uint256 value) public returns (bool) {
        require(balanceOf[from] >= value, "Insufficient balance");
        require(allowance[from][msg.sender] >= value, "Insufficient allowance");
        balanceOf[from] -= value;
        balanceOf[to] += value;
        allowance[from][msg.sender] -= value;
        emit Transfer(from, to, value);
        return true;
    }
}
```

Deploy:
```bash
npx hardhat run scripts/deploy.js --network nlg
```

## EVM + Cosmos SDK Integration Patterns

### Pattern 1: EVM Contract calls Cosmos Module

Use precompiled contracts to expose Cosmos modules to EVM:

```go
// Create precompiled contract for social media module
type SocialMediaPrecompile struct {
    keeper socialMediaKeeper.Keeper
}

func (s *SocialMediaPrecompile) Run(input []byte) ([]byte, error) {
    // Decode input and call Cosmos module
    // Return result
}
```

### Pattern 2: Cosmos Module triggers EVM

```go
// In keeper logic
func (k Keeper) CreatePostWithReward(ctx sdk.Context, post types.Post) error {
    // Create post in Cosmos module
    k.SetPost(ctx, post)
    
    // Trigger EVM contract to mint reward tokens
    evmMsg := evmtypes.NewMsgEthereumTx(...)
    _, err := k.evmKeeper.EthereumTx(ctx, evmMsg)
    
    return err
}
```

### Pattern 3: Cross-module Communication

Use IBC to connect EVM state with other chains:

```go
// Send token from EVM to another chain
func (k Keeper) TransferToIBC(
    ctx sdk.Context,
    sender common.Address,
    recipient string,
    amount *big.Int,
) error {
    // Convert EVM address to Cosmos address
    cosmosAddr := sdk.AccAddress(sender.Bytes())
    
    // Create IBC transfer
    return k.ibcTransferKeeper.SendTransfer(...)
}
```

## Advanced Features

### ERC-20 to Native Token Bridge

Map Cosmos SDK tokens to ERC-20:

```solidity
interface IERC20Token {
    function transfer(address to, uint256 amount) external returns (bool);
    function balanceOf(address account) external view returns (uint256);
}

// In precompile
func (b *TokenBridge) ConvertToERC20(
    ctx sdk.Context,
    account sdk.AccAddress,
    amount sdk.Coin,
) error {
    // Burn native token
    err := b.bankKeeper.SendCoinsFromAccountToModule(ctx, account, types.ModuleName, sdk.NewCoins(amount))
    
    // Mint ERC-20
    evmAddr := common.BytesToAddress(account.Bytes())
    b.evmKeeper.AddBalance(ctx, evmAddr, amount.Amount.BigInt())
    
    return err
}
```

### Gas Price Strategy

Configure dynamic gas pricing:

```go
// In app/app.go
app.FeeMarketKeeper = feemarketkeeper.NewKeeper(
    appCodec,
    keys[feemarkettypes.StoreKey],
    app.GetSubspace(feemarkettypes.ModuleName),
    &feemarketkeeper.Config{
        BaseFee:       sdk.NewInt(1000000000), // 1 gwei
        MinGasPrice:   sdk.NewDecWithPrec(1, 9),
        MinGasMultiplier: sdk.NewDecWithPrec(5, 1), // 0.5
    },
)
```

## Security Considerations

1. **Gas Limits**: Set appropriate gas limits to prevent DOS
2. **Contract Verification**: Implement contract verification system
3. **Precompile Security**: Audit custom precompiles thoroughly
4. **State Sync**: Ensure EVM and Cosmos state consistency
5. **Upgrade Safety**: Test EVM upgrades on testnet first

## Performance Optimization

```go
// Enable EVM optimizations
evmParams := evmtypes.DefaultParams()
evmParams.ExtraEIPs = []int64{3855, 3860} // PUSH0, INITCODE operations
evmParams.EnableCreate = true
evmParams.EnableCall = true
app.EvmKeeper.SetParams(ctx, evmParams)
```

## Monitoring EVM Activity

```bash
# Query EVM module status
nlgd query evm params

# Get EVM account balance
nlgd query evm balance <address>

# Get contract code
nlgd query evm code <contract-address>

# Query transaction receipt
nlgd query evm receipt <tx-hash>
```

## Troubleshooting

### EVM Module Not Loading
- Check module registration in `app.go`
- Verify genesis state includes EVM params
- Ensure all dependencies are installed

### JSON-RPC Not Responding
- Check firewall settings
- Verify port 8545/8546 availability
- Review logs: `nlgd start --log_level debug`

### Contract Deployment Fails
- Verify gas price and limit
- Check account balance
- Ensure EVM params allow contract creation

## Resources

- [Ethermint Documentation](https://docs.ethermint.zone/)
- [Evmos GitHub](https://github.com/evmos/evmos)
- [EVM Module Source](https://github.com/evmos/ethermint/tree/main/x/evm)
- [Web3.js Documentation](https://web3js.readthedocs.io/)

## Next Steps

1. Test basic EVM transactions
2. Deploy sample contracts
3. Create precompiles for social media module
4. Build Web3 frontend integration
5. Set up monitoring and alerts

---

For questions or issues, please open a GitHub issue or join our developer Discord.
