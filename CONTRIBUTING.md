# Contributing to NLG Blockchain

Thank you for your interest in contributing to the NLG blockchain! This document provides guidelines and instructions for contributing.

## Table of Contents

- [Code of Conduct](#code-of-conduct)
- [Getting Started](#getting-started)
- [Development Process](#development-process)
- [Submitting Changes](#submitting-changes)
- [Coding Standards](#coding-standards)
- [Testing Guidelines](#testing-guidelines)
- [Documentation](#documentation)

## Code of Conduct

### Our Pledge

We are committed to providing a welcoming and inspiring community for all. Please be respectful and constructive in your interactions.

### Expected Behavior

- Be respectful and inclusive
- Accept constructive criticism gracefully
- Focus on what is best for the community
- Show empathy towards other community members

### Unacceptable Behavior

- Harassment or discriminatory language
- Trolling or insulting comments
- Public or private harassment
- Publishing others' private information

## Getting Started

### Prerequisites

- Go 1.24.1+
- Ignite CLI v29.6.2+
- Git
- Basic understanding of Cosmos SDK

### Fork and Clone

```bash
# Fork the repository on GitHub
# Then clone your fork
git clone https://github.com/YOUR_USERNAME/nlg.git
cd nlg

# Add upstream remote
git remote add upstream https://github.com/original/nlg.git
```

### Setup Development Environment

```bash
# Install dependencies
go mod tidy

# Build the chain
ignite chain build

# Run tests
go test ./...
```

## Development Process

### 1. Create a Branch

```bash
# Update main branch
git checkout main
git pull upstream main

# Create feature branch
git checkout -b feature/your-feature-name
```

Branch naming conventions:
- `feature/` - New features
- `fix/` - Bug fixes
- `docs/` - Documentation changes
- `refactor/` - Code refactoring
- `test/` - Test additions or changes

### 2. Make Your Changes

- Write clean, readable code
- Follow existing code style
- Add tests for new features
- Update documentation as needed
- Keep commits focused and atomic

### 3. Commit Your Changes

We use [Conventional Commits](https://www.conventionalcommits.org/):

```
<type>(<scope>): <subject>

<body>

<footer>
```

**Types:**
- `feat`: New feature
- `fix`: Bug fix
- `docs`: Documentation only
- `style`: Code style changes (formatting, etc.)
- `refactor`: Code refactoring
- `test`: Adding or updating tests
- `chore`: Maintenance tasks

**Examples:**

```bash
git commit -m "feat(socialmedia): add post sharing functionality"
git commit -m "fix(keeper): resolve state corruption in post deletion"
git commit -m "docs(api): add examples for profile queries"
```

### 4. Push and Create Pull Request

```bash
# Push to your fork
git push origin feature/your-feature-name

# Create pull request on GitHub
```

## Submitting Changes

### Pull Request Process

1. **Update Documentation**
   - Update README.md if needed
   - Add/update API documentation
   - Update CHANGELOG.md

2. **Ensure Tests Pass**
   ```bash
   go test ./...
   golangci-lint run
   ```

3. **Describe Your Changes**
   - Clear title and description
   - Reference related issues
   - Include screenshots/examples if applicable

4. **Wait for Review**
   - Address reviewer feedback
   - Keep PR updated with main branch
   - Be patient and responsive

### Pull Request Template

```markdown
## Description
Brief description of the changes

## Type of Change
- [ ] Bug fix
- [ ] New feature
- [ ] Breaking change
- [ ] Documentation update

## Testing
How was this tested?

## Checklist
- [ ] Tests pass locally
- [ ] Code follows style guidelines
- [ ] Documentation updated
- [ ] Changelog updated
- [ ] No breaking changes (or documented)
```

## Coding Standards

### Go Style Guide

Follow [Effective Go](https://golang.org/doc/effective_go.html) and the [Cosmos SDK Style Guide](https://github.com/cosmos/cosmos-sdk/blob/main/CODING_GUIDELINES.md).

### Code Formatting

```bash
# Format code
gofmt -w .

# Or use goimports
goimports -w .
```

### Linting

```bash
# Run linter
golangci-lint run

# Auto-fix issues
golangci-lint run --fix
```

### Naming Conventions

**Variables:**
```go
// Good
userProfile := types.Profile{}
postID := "123"

// Bad
up := types.Profile{}
id := "123"
```

**Functions:**
```go
// Good - Clear and descriptive
func (k Keeper) GetPostByID(ctx sdk.Context, id string) (types.Post, error)

// Bad - Unclear
func (k Keeper) Get(ctx sdk.Context, id string) (types.Post, error)
```

**Constants:**
```go
const (
    MaxPostLength = 5000
    MinReputationToModerate = 100
)
```

### Error Handling

```go
// Good - Wrap errors with context
if err != nil {
    return sdkerrors.Wrap(err, "failed to create post")
}

// Bad - Generic error
if err != nil {
    return err
}
```

### Logging

```go
// Use structured logging
k.Logger(ctx).Info("post created",
    "post_id", postID,
    "creator", msg.Creator,
)

// Include error details
k.Logger(ctx).Error("failed to delete post",
    "error", err,
    "post_id", postID,
)
```

## Testing Guidelines

### Unit Tests

```go
func TestCreatePost(t *testing.T) {
    // Setup
    keeper, ctx := setupKeeper(t)
    
    // Test cases
    tests := []struct {
        name    string
        post    types.Post
        wantErr bool
    }{
        {
            name: "valid post",
            post: types.Post{
                Title:   "Test Post",
                Content: "Test content",
            },
            wantErr: false,
        },
        {
            name: "empty title",
            post: types.Post{
                Title:   "",
                Content: "Test content",
            },
            wantErr: true,
        },
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            err := keeper.CreatePost(ctx, tt.post)
            if (err != nil) != tt.wantErr {
                t.Errorf("CreatePost() error = %v, wantErr %v", err, tt.wantErr)
            }
        })
    }
}
```

### Test Coverage

Aim for at least 80% code coverage:

```bash
# Generate coverage report
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

### Integration Tests

Place integration tests in separate files with build tags:

```go
//go:build integration
// +build integration

package keeper_test

func TestIntegration(t *testing.T) {
    // Integration test code
}
```

Run with:
```bash
go test -tags=integration ./...
```

## Documentation

### Code Comments

```go
// CreatePost creates a new post in the blockchain state.
// It validates the post data and assigns a unique ID.
//
// Parameters:
//   - ctx: SDK context
//   - msg: Message containing post data
//
// Returns:
//   - Post ID on success
//   - Error if validation fails or state operation fails
func (k Keeper) CreatePost(ctx sdk.Context, msg *types.MsgCreatePost) (string, error) {
    // Implementation
}
```

### Proto Documentation

```protobuf
// Post represents a social media post on the blockchain
message Post {
  // Unique identifier for the post
  string id = 1;
  
  // Post title (max 200 characters)
  string title = 2;
  
  // Post content (max 5000 characters)
  string content = 3;
}
```

### API Documentation

Update `docs/API.md` when adding new endpoints:

```markdown
### Create Post

Creates a new post on the blockchain.

#### Endpoint
POST /nlg/socialmedia/post

#### Request Body
\`\`\`json
{
  "title": "Post Title",
  "content": "Post content"
}
\`\`\`

#### Response
\`\`\`json
{
  "post_id": "123",
  "tx_hash": "ABC..."
}
\`\`\`
```

## Communication

### Channels

- **GitHub Issues**: Bug reports and feature requests
- **GitHub Discussions**: General questions and ideas
- **Discord**: Real-time chat and support
- **Dev Calls**: Weekly development meetings

### Asking for Help

- Search existing issues first
- Provide clear reproduction steps
- Include relevant logs and code
- Be specific about your environment

### Reporting Bugs

Include:
1. Description of the bug
2. Steps to reproduce
3. Expected behavior
4. Actual behavior
5. Environment details (OS, Go version, etc.)
6. Relevant logs

### Suggesting Features

Include:
1. Clear description of the feature
2. Use case and motivation
3. Proposed implementation (if any)
4. Potential impact on existing features

## Review Process

### What We Look For

- **Correctness**: Does it work as intended?
- **Quality**: Is the code clean and maintainable?
- **Tests**: Are there adequate tests?
- **Documentation**: Is it well documented?
- **Performance**: Are there any performance concerns?

### Responding to Feedback

- Be open to suggestions
- Ask questions if unclear
- Make requested changes promptly
- Explain your reasoning when disagreeing

### Getting Approved

- Two approvals from maintainers required
- All tests must pass
- No unresolved conversations
- Up to date with main branch

## Release Process

1. Maintainers prepare release notes
2. Version bump and tag created
3. Binaries built and tested
4. Release published on GitHub
5. Documentation updated

## Recognition

Contributors are recognized in:
- CHANGELOG.md
- Release notes
- Project README
- Annual contributor spotlight

## Questions?

Feel free to reach out:
- Open a discussion on GitHub
- Join our Discord server
- Attend weekly dev calls

Thank you for contributing to NLG blockchain! 🚀
