# Changelog

All notable changes to the NLG blockchain will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- Initial blockchain scaffold using Ignite CLI v29.6.2
- Custom social media module with comprehensive features
- Post creation, update, and deletion functionality
- User profile management system
- Comment system for posts
- Social connection system (follow/unfollow)
- Content moderation with voting mechanism
- Like functionality for posts
- 100M NLG token total supply configuration
- Chain ID: nlg-1
- Address prefix: nlg
- Standard Cosmos SDK modules: bank, auth, staking, distribution, slashing, governance, feegrant, authz
- IPFS hash storage for decentralized content
- Comprehensive API documentation
- REST API endpoints for all social media features
- gRPC services for module queries and transactions
- CLI commands for all operations
- Development setup with Ignite serve
- Multi-node testnet support
- EVM integration guide for Ethereum compatibility (documentation)
- Deployment guides for various environments
- API documentation with examples in multiple languages
- Development guide for contributors

### Changed
- Default denomination from "stake" to "nlg"
- Genesis configuration to distribute 100M tokens
- Module dependencies to include bank and staking keepers

### Security
- Transaction signing required for all state changes
- Validator-based consensus using Tendermint BFT
- Gas metering for resource management

## [0.1.0] - 2024-01-03

### Initial Release
- NLG blockchain with social media capabilities
- Cosmos SDK v0.53.4
- Go 1.24.1
- Ignite CLI v29.6.2

---

## Version History

### Version Numbering

This project uses semantic versioning:
- MAJOR version for incompatible API changes
- MINOR version for added functionality in a backwards compatible manner
- PATCH version for backwards compatible bug fixes

### Upgrade Path

When upgrading between versions, please follow the migration guides in the `/docs/migrations/` directory.

---

For more information about releases, see the [GitHub Releases page](https://github.com/your-org/nlg/releases).
