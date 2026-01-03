# NLG Blockchain API Documentation

Complete API reference for the NLG blockchain social media module.

## Table of Contents

1. [Overview](#overview)
2. [Authentication](#authentication)
3. [Posts API](#posts-api)
4. [Profiles API](#profiles-api)
5. [Comments API](#comments-api)
6. [Social Connections API](#social-connections-api)
7. [Moderation API](#moderation-api)
8. [Query Examples](#query-examples)

## Overview

The NLG blockchain exposes three types of APIs:

- **REST API**: HTTP endpoints at port 1317
- **gRPC API**: gRPC services at port 9090
- **CLI**: Command-line interface via `nlgd` binary

### Base URLs

- REST: `http://localhost:1317`
- gRPC: `localhost:9090`
- RPC: `http://localhost:26657`

## Authentication

All transactions require signing with a private key. Use the `--from` flag with CLI commands or sign transactions programmatically.

### Generate Account

```bash
# Create new account
nlgd keys add myaccount

# Import existing account
nlgd keys add myaccount --recover

# List accounts
nlgd keys list
```

### Transaction Signing

```bash
# Sign and broadcast in one step
nlgd tx socialmedia <command> --from myaccount --chain-id nlg-1 --yes

# Sign transaction offline
nlgd tx socialmedia <command> --from myaccount --generate-only > unsigned.json
nlgd tx sign unsigned.json --from myaccount > signed.json
nlgd tx broadcast signed.json
```

## Posts API

### Create Post

Create a new post on the blockchain.

#### CLI Command

```bash
nlgd tx socialmedia create-post [title] [content] [ipfs-hash] [timestamp] [likes-count] \
  --from [creator] \
  --chain-id nlg-1 \
  --gas auto \
  --gas-adjustment 1.3 \
  --fees 5000nlg
```

#### Example

```bash
nlgd tx socialmedia create-post \
  "Welcome to NLG!" \
  "This is my first post on the decentralized social network" \
  "QmYwAPJzv5CZsnANT5TXLk68TnSqKHAKERw2qH5ZYi8KqR" \
  $(date +%s) \
  0 \
  --from alice \
  --chain-id nlg-1 \
  --yes
```

#### REST API

```http
POST /nlg/socialmedia/post
Content-Type: application/json

{
  "creator": "nlg1...",
  "title": "Welcome to NLG!",
  "content": "This is my first post",
  "ipfsHash": "QmYwAPJzv5CZsnANT5TXLk68TnSqKHAKERw2qH5ZYi8KqR",
  "timestamp": 1704312000,
  "likesCount": 0
}
```

#### Response

```json
{
  "code": 0,
  "txhash": "A7F5B2...",
  "height": "12345",
  "logs": [{
    "msg_index": 0,
    "events": [{
      "type": "create_post",
      "attributes": [{
        "key": "post_id",
        "value": "0"
      }]
    }]
  }]
}
```

### Update Post

Update an existing post.

#### CLI Command

```bash
nlgd tx socialmedia update-post [id] [title] [content] [ipfs-hash] [timestamp] [likes-count] \
  --from [creator]
```

#### Example

```bash
nlgd tx socialmedia update-post 0 \
  "Updated Title" \
  "Updated content" \
  "QmNewHash..." \
  $(date +%s) \
  5 \
  --from alice
```

### Delete Post

Delete a post by ID.

#### CLI Command

```bash
nlgd tx socialmedia delete-post [id] --from [creator]
```

#### Example

```bash
nlgd tx socialmedia delete-post 0 --from alice
```

### Query Posts

#### Get All Posts

```bash
# CLI
nlgd query socialmedia list-post

# REST
curl http://localhost:1317/nlg/socialmedia/post

# Response
{
  "post": [
    {
      "id": "0",
      "title": "Welcome to NLG!",
      "content": "This is my first post",
      "ipfsHash": "QmYwAPJzv5CZsnANT5TXLk68TnSqKHAKERw2qH5ZYi8KqR",
      "timestamp": "1704312000",
      "likesCount": "5",
      "creator": "nlg1..."
    }
  ],
  "pagination": {
    "next_key": null,
    "total": "1"
  }
}
```

#### Get Post by ID

```bash
# CLI
nlgd query socialmedia show-post [id]

# REST
curl http://localhost:1317/nlg/socialmedia/post/0

# Response
{
  "post": {
    "id": "0",
    "title": "Welcome to NLG!",
    "content": "This is my first post",
    "ipfsHash": "QmYwAPJzv5CZsnANT5TXLk68TnSqKHAKERw2qH5ZYi8KqR",
    "timestamp": "1704312000",
    "likesCount": "5",
    "creator": "nlg1..."
  }
}
```

## Profiles API

### Create Profile

Create a user profile.

#### CLI Command

```bash
nlgd tx socialmedia create-profile \
  [username] \
  [bio] \
  [avatar-ipfs-hash] \
  [reputation] \
  [followers-count] \
  [following-count] \
  --from [creator]
```

#### Example

```bash
nlgd tx socialmedia create-profile \
  "alice_nlg" \
  "Blockchain enthusiast and developer" \
  "QmProfileAvatar..." \
  100 \
  50 \
  30 \
  --from alice
```

### Update Profile

```bash
nlgd tx socialmedia update-profile [id] \
  "alice_updated" \
  "Updated bio" \
  "QmNewAvatar..." \
  150 \
  60 \
  35 \
  --from alice
```

### Query Profiles

#### Get All Profiles

```bash
# CLI
nlgd query socialmedia list-profile

# REST
curl http://localhost:1317/nlg/socialmedia/profile

# Response
{
  "profile": [
    {
      "id": "0",
      "username": "alice_nlg",
      "bio": "Blockchain enthusiast and developer",
      "avatarIpfsHash": "QmProfileAvatar...",
      "reputation": "100",
      "followersCount": "50",
      "followingCount": "30",
      "creator": "nlg1..."
    }
  ],
  "pagination": {
    "next_key": null,
    "total": "1"
  }
}
```

#### Get Profile by ID

```bash
# CLI
nlgd query socialmedia show-profile [id]

# REST
curl http://localhost:1317/nlg/socialmedia/profile/0
```

## Comments API

### Create Comment

Add a comment to a post.

#### CLI Command

```bash
nlgd tx socialmedia create-comment \
  [post-id] \
  [content] \
  [timestamp] \
  [author] \
  --from [creator]
```

#### Example

```bash
nlgd tx socialmedia create-comment \
  "0" \
  "Great post! Thanks for sharing." \
  $(date +%s) \
  "nlg1..." \
  --from bob
```

### Update Comment

```bash
nlgd tx socialmedia update-comment [id] \
  "0" \
  "Updated comment content" \
  $(date +%s) \
  "nlg1..." \
  --from bob
```

### Delete Comment

```bash
nlgd tx socialmedia delete-comment [id] --from bob
```

### Query Comments

#### Get All Comments

```bash
# CLI
nlgd query socialmedia list-comment

# REST
curl http://localhost:1317/nlg/socialmedia/comment

# Response
{
  "comment": [
    {
      "id": "0",
      "postId": "0",
      "content": "Great post! Thanks for sharing.",
      "timestamp": "1704312100",
      "author": "nlg1...",
      "creator": "nlg1..."
    }
  ],
  "pagination": {
    "next_key": null,
    "total": "1"
  }
}
```

#### Filter Comments by Post

```bash
# CLI with custom query
nlgd query socialmedia list-comment --output json | jq '.comment[] | select(.postId=="0")'

# Using gRPC reflection (if implemented)
grpcurl -plaintext -d '{"postId":"0"}' localhost:9090 nlg.socialmedia.v1.Query/CommentsByPost
```

## Social Connections API

### Create Connection (Follow)

Follow another user.

#### CLI Command

```bash
nlgd tx socialmedia create-social-connection \
  [follower] \
  [following] \
  [timestamp] \
  --from [creator]
```

#### Example

```bash
nlgd tx socialmedia create-social-connection \
  $(nlgd keys show alice -a) \
  $(nlgd keys show bob -a) \
  $(date +%s) \
  --from alice
```

### Delete Connection (Unfollow)

```bash
nlgd tx socialmedia delete-social-connection [id] --from alice
```

### Query Connections

#### Get All Connections

```bash
# CLI
nlgd query socialmedia list-social-connection

# REST
curl http://localhost:1317/nlg/socialmedia/social-connection

# Response
{
  "socialConnection": [
    {
      "id": "0",
      "follower": "nlg1alice...",
      "following": "nlg1bob...",
      "timestamp": "1704312200",
      "creator": "nlg1alice..."
    }
  ],
  "pagination": {
    "next_key": null,
    "total": "1"
  }
}
```

## Moderation API

### Like Post

Like a specific post.

#### CLI Command

```bash
nlgd tx socialmedia like-post [post-id] --from [creator]
```

#### Example

```bash
nlgd tx socialmedia like-post "0" --from bob
```

### Moderate Content

Vote to moderate content (flag, approve, etc.).

#### CLI Command

```bash
nlgd tx socialmedia moderate-content [content-id] [vote-type] --from [creator]
```

#### Example

```bash
# Flag inappropriate content
nlgd tx socialmedia moderate-content "post-0" "flag" --from charlie

# Approve content
nlgd tx socialmedia moderate-content "post-0" "approve" --from alice
```

#### Vote Types

- `flag`: Flag content as inappropriate
- `approve`: Approve content
- `spam`: Mark as spam
- `remove`: Vote to remove content

## Query Examples

### Using cURL

#### Get All Posts

```bash
curl http://localhost:1317/nlg/socialmedia/post | jq .
```

#### Get Post with Pagination

```bash
curl "http://localhost:1317/nlg/socialmedia/post?pagination.limit=10&pagination.offset=0" | jq .
```

### Using gRPC

#### Install grpcurl

```bash
go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
```

#### List Available Services

```bash
grpcurl -plaintext localhost:9090 list
```

#### Query Posts

```bash
grpcurl -plaintext localhost:9090 nlg.socialmedia.v1.Query/PostAll
```

#### Query Specific Post

```bash
grpcurl -plaintext -d '{"id":"0"}' localhost:9090 nlg.socialmedia.v1.Query/Post
```

### Using JavaScript/TypeScript

#### Install Dependencies

```bash
npm install @cosmjs/stargate @cosmjs/proto-signing
```

#### Query Example

```typescript
import { StargateClient } from "@cosmjs/stargate";

async function queryPosts() {
  const client = await StargateClient.connect("http://localhost:26657");
  
  const response = await client.queryContractSmart(
    "nlg1...",
    { list_posts: {} }
  );
  
  console.log("Posts:", response);
}

queryPosts();
```

#### Transaction Example

```typescript
import { SigningStargateClient } from "@cosmjs/stargate";
import { DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";

async function createPost() {
  const mnemonic = "your mnemonic here";
  const wallet = await DirectSecp256k1HdWallet.fromMnemonic(mnemonic, {
    prefix: "nlg"
  });
  
  const [account] = await wallet.getAccounts();
  
  const client = await SigningStargateClient.connectWithSigner(
    "http://localhost:26657",
    wallet
  );
  
  const msg = {
    typeUrl: "/nlg.socialmedia.v1.MsgCreatePost",
    value: {
      creator: account.address,
      title: "My Post",
      content: "Post content",
      ipfsHash: "QmHash...",
      timestamp: Math.floor(Date.now() / 1000),
      likesCount: 0
    }
  };
  
  const fee = {
    amount: [{ denom: "nlg", amount: "5000" }],
    gas: "200000"
  };
  
  const result = await client.signAndBroadcast(
    account.address,
    [msg],
    fee,
    "Creating post"
  );
  
  console.log("Transaction hash:", result.transactionHash);
}

createPost();
```

### Using Python

#### Install Dependencies

```bash
pip install cosmpy
```

#### Query Example

```python
import requests
import json

def get_posts():
    url = "http://localhost:1317/nlg/socialmedia/post"
    response = requests.get(url)
    data = response.json()
    
    for post in data['post']:
        print(f"Post {post['id']}: {post['title']}")
    
    return data

posts = get_posts()
```

#### Transaction Example

```python
from cosmpy.aerial.client import LedgerClient, NetworkConfig
from cosmpy.aerial.wallet import LocalWallet
from cosmpy.crypto.keypairs import PrivateKey

# Setup
network = NetworkConfig(
    chain_id="nlg-1",
    url="http://localhost:26657",
    fee_minimum_gas_price=1,
    fee_denomination="nlg",
    staking_denomination="nlg"
)

wallet = LocalWallet(PrivateKey("your-private-key-hex"))
client = LedgerClient(network)

# Create transaction
tx = {
    "@type": "/nlg.socialmedia.v1.MsgCreatePost",
    "creator": str(wallet.address()),
    "title": "Python Post",
    "content": "Posted from Python",
    "ipfsHash": "QmHash...",
    "timestamp": int(time.time()),
    "likesCount": 0
}

# Sign and broadcast
result = client.broadcast_tx(tx, wallet)
print(f"Transaction hash: {result.tx_hash}")
```

## Error Codes

| Code | Description |
|------|-------------|
| 0 | Success |
| 1 | Internal error |
| 2 | Transaction parsing error |
| 3 | Invalid sequence |
| 4 | Unauthorized |
| 5 | Insufficient funds |
| 11 | Out of gas |
| 12 | Invalid memo |
| 13 | Insufficient fee |
| 32 | Item not found |

## Rate Limiting

The API implements rate limiting to prevent abuse:

- REST API: 100 requests per minute per IP
- gRPC: 1000 requests per minute per connection
- Transactions: Limited by gas and fees

## Pagination

All list queries support pagination:

```bash
# CLI
nlgd query socialmedia list-post --page 1 --limit 10

# REST
curl "http://localhost:1317/nlg/socialmedia/post?pagination.limit=10&pagination.key=<next_key>"
```

## WebSocket Subscriptions

Subscribe to blockchain events:

```bash
# Using wscat
wscat -c ws://localhost:26657/websocket

# Subscribe to new posts (custom implementation needed)
{"jsonrpc":"2.0","method":"subscribe","params":["tm.event='Tx' AND create_post.creator EXISTS"],"id":1}
```

## Best Practices

1. **Always validate input** before submitting transactions
2. **Use pagination** for large data sets
3. **Implement retry logic** for failed requests
4. **Cache frequently accessed data** to reduce API calls
5. **Use WebSocket subscriptions** for real-time updates
6. **Handle errors gracefully** with proper error messages
7. **Set appropriate gas limits** to avoid out-of-gas errors

## Support

For API support:
- GitHub Issues: Report bugs and request features
- Discord: #api-support channel
- Documentation: https://docs.nlg.network/api

---

API Version: 1.0  
Last Updated: 2024
