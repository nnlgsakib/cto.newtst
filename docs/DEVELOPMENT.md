# Development Guide for NLG Blockchain

Guide for developers working on the NLG blockchain.

## Table of Contents

1. [Development Setup](#development-setup)
2. [Project Structure](#project-structure)
3. [Building and Testing](#building-and-testing)
4. [Adding New Features](#adding-new-features)
5. [Contributing Guidelines](#contributing-guidelines)

## Development Setup

### Prerequisites

- Go 1.24.1+
- Ignite CLI v29.6.2+
- Git
- Make

### Clone and Build

```bash
# Clone repository
git clone <repository-url>
cd nlg

# Install dependencies
go mod tidy

# Build the chain
ignite chain build

# Add to PATH
export PATH=$PATH:$HOME/go/bin

# Verify installation
nlgd version
```

### IDE Setup

#### VS Code

Recommended extensions:
- Go (golang.go)
- Protobuf (zxh404.vscode-proto3)
- YAML (redhat.vscode-yaml)

Settings (`.vscode/settings.json`):

```json
{
  "go.useLanguageServer": true,
  "go.lintTool": "golangci-lint",
  "go.lintOnSave": "workspace",
  "go.formatTool": "goimports",
  "[go]": {
    "editor.formatOnSave": true,
    "editor.codeActionsOnSave": {
      "source.organizeImports": true
    }
  }
}
```

#### GoLand

1. Open project directory
2. Enable Go Modules support
3. Configure code style to match project conventions
4. Install Protobuf plugin

## Project Structure

```
nlg/
├── app/                    # Application setup and configuration
│   ├── app.go             # Main application logic
│   ├── app_config.go      # Dependency injection configuration
│   ├── config.go          # Application configuration
│   └── genesis.go         # Genesis initialization
├── cmd/                    # Command-line interface
│   └── nlgd/
│       ├── main.go        # Entry point
│       └── cmd/
│           ├── commands.go # Command definitions
│           ├── root.go    # Root command
│           └── config.go  # CLI configuration
├── x/                      # Custom modules
│   └── socialmedia/       # Social media module
│       ├── keeper/        # State management
│       ├── types/         # Type definitions
│       ├── module/        # Module implementation
│       └── simulation/    # Simulation tests
├── proto/                  # Protocol buffer definitions
│   └── nlg/
│       └── socialmedia/
│           └── v1/
├── docs/                   # Documentation
├── testutil/              # Testing utilities
├── config.yml             # Ignite configuration
└── go.mod                 # Go dependencies
```

## Building and Testing

### Build Commands

```bash
# Build binary
ignite chain build

# Build with debug symbols
go build -gcflags="all=-N -l" -o build/nlgd ./cmd/nlgd

# Clean build
rm -rf build/ && ignite chain build
```

### Running the Chain

```bash
# Start development chain with hot reload
ignite chain serve

# Start without reload
ignite chain build
nlgd start

# Start with custom home directory
nlgd start --home ~/.nlg-dev

# Start with specific chain ID
nlgd start --chain-id nlg-dev-1
```

### Testing

#### Unit Tests

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run tests with verbose output
go test -v ./...

# Run specific module tests
go test ./x/socialmedia/...

# Run specific test
go test -run TestCreatePost ./x/socialmedia/keeper/...

# Generate coverage report
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out -o coverage.html
```

#### Integration Tests

```bash
# Run integration tests (if implemented)
go test -tags=integration ./...
```

#### Simulation Tests

```bash
# Run simulation
ignite chain simulate

# Or manually
go test -run TestFullAppSimulation ./app/...
```

### Code Quality

#### Linting

```bash
# Install golangci-lint
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Run linter
golangci-lint run

# Auto-fix issues
golangci-lint run --fix
```

#### Formatting

```bash
# Format all Go files
gofmt -w .

# Or use goimports
go install golang.org/x/tools/cmd/goimports@latest
goimports -w .
```

## Adding New Features

### Adding a New Message Type

```bash
# Scaffold a new message
ignite scaffold message [message-name] [field1:type] [field2:type] \
  --module socialmedia \
  --signer creator

# Example: Add a "share post" message
ignite scaffold message sharePost postId:string --module socialmedia --signer creator
```

### Adding a New Query

Edit proto file and implement keeper method:

```protobuf
// proto/nlg/socialmedia/v1/query.proto
service Query {
  rpc PostsByUser(QueryPostsByUserRequest) returns (QueryPostsByUserResponse) {
    option (google.api.http).get = "/nlg/socialmedia/posts/{user}";
  }
}

message QueryPostsByUserRequest {
  string user = 1;
}

message QueryPostsByUserResponse {
  repeated Post posts = 1;
}
```

Generate code:

```bash
ignite generate proto-go
```

Implement in keeper:

```go
// x/socialmedia/keeper/query_posts_by_user.go
package keeper

import (
    "context"
    sdk "github.com/cosmos/cosmos-sdk/types"
    "nlg/x/socialmedia/types"
)

func (k Keeper) PostsByUser(goCtx context.Context, req *types.QueryPostsByUserRequest) (*types.QueryPostsByUserResponse, error) {
    ctx := sdk.UnwrapSDKContext(goCtx)
    
    var posts []types.Post
    store := k.storeService.OpenKVStore(ctx)
    
    // Implement query logic
    // ...
    
    return &types.QueryPostsByUserResponse{
        Posts: posts,
    }, nil
}
```

### Adding Module Parameters

```protobuf
// proto/nlg/socialmedia/v1/params.proto
message Params {
  option (gogoproto.goproto_stringer) = false;
  
  uint64 max_post_length = 1;
  uint64 min_reputation_to_moderate = 2;
  string moderation_fee = 3 [
    (cosmos_proto.scalar) = "cosmos.Int",
    (gogoproto.customtype) = "cosmossdk.io/math.Int",
    (gogoproto.nullable) = false
  ];
}
```

Update default params:

```go
// x/socialmedia/types/params.go
func DefaultParams() Params {
    return Params{
        MaxPostLength: 5000,
        MinReputationToModerate: 100,
        ModerationFee: sdk.NewInt(1000),
    }
}
```

### Adding Events

```go
// Emit event in keeper
ctx.EventManager().EmitEvent(
    sdk.NewEvent(
        "post_created",
        sdk.NewAttribute("post_id", postID),
        sdk.NewAttribute("creator", msg.Creator),
        sdk.NewAttribute("timestamp", fmt.Sprintf("%d", msg.Timestamp)),
    ),
)
```

### Adding Custom Errors

```go
// x/socialmedia/types/errors.go
package types

import (
    "cosmossdk.io/errors"
)

var (
    ErrPostNotFound = errors.Register(ModuleName, 1100, "post not found")
    ErrUnauthorized = errors.Register(ModuleName, 1101, "unauthorized")
    ErrInvalidContent = errors.Register(ModuleName, 1102, "invalid content")
)
```

Use in keeper:

```go
if post == nil {
    return nil, types.ErrPostNotFound
}
```

## Working with Protocol Buffers

### Generate Proto Files

```bash
# Generate Go code from proto files
ignite generate proto-go

# Generate OpenAPI documentation
ignite generate openapi

# Generate TypeScript client
ignite generate ts-client
```

### Proto Best Practices

1. Use semantic versioning for proto packages
2. Never remove or change field numbers
3. Use reserved fields for deprecated fields
4. Document all messages and fields
5. Use appropriate scalar types

Example:

```protobuf
syntax = "proto3";

package nlg.socialmedia.v1;

option go_package = "nlg/x/socialmedia/types";

// Post represents a social media post
message Post {
  option (gogoproto.equal) = false;
  
  string id = 1; // Unique post identifier
  string title = 2; // Post title
  string content = 3; // Post content
  
  // Reserved for future use
  reserved 4;
  reserved "deprecated_field";
}
```

## Database and State Management

### Accessing State

```go
// In keeper methods
func (k Keeper) GetPost(ctx sdk.Context, id string) (types.Post, bool) {
    store := k.storeService.OpenKVStore(ctx)
    
    bz, err := store.Get(types.PostKey(id))
    if err != nil || bz == nil {
        return types.Post{}, false
    }
    
    var post types.Post
    k.cdc.MustUnmarshal(bz, &post)
    return post, true
}

func (k Keeper) SetPost(ctx sdk.Context, post types.Post) {
    store := k.storeService.OpenKVStore(ctx)
    
    bz := k.cdc.MustMarshal(&post)
    store.Set(types.PostKey(post.Id), bz)
}
```

### Iterating State

```go
func (k Keeper) GetAllPosts(ctx sdk.Context) []types.Post {
    store := k.storeService.OpenKVStore(ctx)
    iterator := storetypes.KVStorePrefixIterator(store, types.PostKeyPrefix)
    defer iterator.Close()
    
    var posts []types.Post
    for ; iterator.Valid(); iterator.Next() {
        var post types.Post
        k.cdc.MustUnmarshal(iterator.Value(), &post)
        posts = append(posts, post)
    }
    
    return posts
}
```

## Debugging

### Enable Debug Logging

```bash
# Start with debug logs
nlgd start --log_level debug

# Or set in config.toml
sed -i 's/log_level = "info"/log_level = "debug"/' ~/.nlg/config/config.toml
```

### Using Delve Debugger

```bash
# Install delve
go install github.com/go-delve/delve/cmd/dlv@latest

# Start chain with debugger
dlv exec ./build/nlgd -- start

# Set breakpoint
(dlv) break x/socialmedia/keeper/msg_server_post.go:15
(dlv) continue
```

### Logging Best Practices

```go
// Use structured logging
k.Logger(ctx).Info("creating post",
    "post_id", postID,
    "creator", msg.Creator,
)

k.Logger(ctx).Error("failed to create post",
    "error", err,
    "post_id", postID,
)
```

## Performance Optimization

### Caching

```go
// Use collections for efficient state access
import "cosmossdk.io/collections"

type Keeper struct {
    Posts collections.Map[string, types.Post]
}
```

### Gas Optimization

```go
// Charge gas for expensive operations
ctx.GasMeter().ConsumeGas(1000, "post creation")

// Check remaining gas
if ctx.GasMeter().IsOutOfGas() {
    return nil, sdkerrors.Wrap(sdkerrors.ErrOutOfGas, "insufficient gas")
}
```

## Common Issues and Solutions

### Issue: Proto files not generating

```bash
# Clear cache and regenerate
rm -rf proto/
ignite generate proto-go
```

### Issue: Module not found

```bash
# Clean and rebuild
go clean -modcache
go mod tidy
ignite chain build
```

### Issue: Chain won't start

```bash
# Reset chain state
nlgd tendermint unsafe-reset-all

# Remove all data
rm -rf ~/.nlg
```

## Contributing Guidelines

### Pull Request Process

1. Fork the repository
2. Create a feature branch: `git checkout -b feature/my-feature`
3. Make your changes
4. Write tests
5. Run linter and tests
6. Commit with conventional commits
7. Push and create PR

### Commit Message Format

```
type(scope): subject

body

footer
```

Types:
- feat: New feature
- fix: Bug fix
- docs: Documentation
- test: Tests
- refactor: Code refactoring
- chore: Maintenance

Example:
```
feat(socialmedia): add post sharing functionality

Implement message type and keeper methods for sharing posts.
Includes tests and documentation.

Closes #123
```

### Code Review Checklist

- [ ] Tests pass
- [ ] Code is formatted
- [ ] No linting errors
- [ ] Documentation updated
- [ ] Changelog updated
- [ ] No breaking changes (or properly documented)

## Resources

- [Cosmos SDK Documentation](https://docs.cosmos.network/)
- [Ignite CLI Docs](https://docs.ignite.com/)
- [Go Documentation](https://golang.org/doc/)
- [Protocol Buffers Guide](https://developers.google.com/protocol-buffers)

## Support

- GitHub Discussions
- Discord: #development channel
- Weekly dev calls (see calendar)

---

Happy coding! 🚀
