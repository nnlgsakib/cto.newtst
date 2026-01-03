# NLG Blockchain Usage Examples

Real-world examples of using the NLG blockchain's social media features.

## Table of Contents

1. [Basic Post Management](#basic-post-management)
2. [User Profile Setup](#user-profile-setup)
3. [Social Interactions](#social-interactions)
4. [Content Moderation](#content-moderation)
5. [Advanced Scenarios](#advanced-scenarios)

## Prerequisites

```bash
# Ensure chain is running
ignite chain serve

# Or manually
nlgd start

# Check available accounts
nlgd keys list
```

## Basic Post Management

### Example 1: Create Your First Post

```bash
# Create a post about blockchain
nlgd tx socialmedia create-post \
  "Introduction to NLG Blockchain" \
  "Welcome to NLG, a decentralized social media platform built on Cosmos SDK. This blockchain enables censorship-resistant content sharing with community governance." \
  "QmYwAPJzv5CZsnA7JkGG1gpXAaWBeaE6HAoQVMfw5gReD8" \
  $(date +%s) \
  0 \
  --from alice \
  --chain-id nlg-1 \
  --gas auto \
  --gas-adjustment 1.3 \
  --yes

# Query to see your post
nlgd query socialmedia list-post
```

### Example 2: Update a Post

```bash
# Update post with ID "0"
nlgd tx socialmedia update-post 0 \
  "Introduction to NLG Blockchain [UPDATED]" \
  "Welcome to NLG! This updated version includes more information about our governance features and moderation system." \
  "QmNewHash456..." \
  $(date +%s) \
  15 \
  --from alice \
  --yes

# View the updated post
nlgd query socialmedia show-post 0
```

### Example 3: Query Posts via API

```bash
# Using curl
curl http://localhost:1317/nlg/socialmedia/post | jq .

# Filter for specific field
curl http://localhost:1317/nlg/socialmedia/post | jq '.post[] | select(.creator=="nlg1...")'

# Count total posts
curl http://localhost:1317/nlg/socialmedia/post | jq '.post | length'
```

## User Profile Setup

### Example 4: Create a Complete Profile

```bash
# Alice creates her profile
nlgd tx socialmedia create-profile \
  "alice_blockchain" \
  "Core developer of NLG blockchain. Passionate about decentralization and Web3. Building the future of social media. 🚀" \
  "QmProfilePicABCD1234567890" \
  100 \
  0 \
  0 \
  --from alice \
  --yes

# Bob creates his profile
nlgd tx socialmedia create-profile \
  "bob_crypto" \
  "Crypto enthusiast and investor. Love exploring new blockchain projects." \
  "QmBobAvatar9876543210" \
  50 \
  0 \
  0 \
  --from bob \
  --yes
```

### Example 5: Update Profile Reputation

```bash
# Update Alice's profile with increased reputation
nlgd tx socialmedia update-profile 0 \
  "alice_blockchain" \
  "Core developer of NLG blockchain. Speaker at Blockchain Conference 2024." \
  "QmProfilePicABCD1234567890" \
  150 \
  25 \
  10 \
  --from alice \
  --yes
```

### Example 6: Query Profile Information

```bash
# Get Alice's profile
ALICE_ADDR=$(nlgd keys show alice -a)
nlgd query socialmedia list-profile | jq --arg addr "$ALICE_ADDR" '.profile[] | select(.creator==$addr)'

# Or via REST API
curl http://localhost:1317/nlg/socialmedia/profile/0 | jq .
```

## Social Interactions

### Example 7: Follow a User

```bash
# Bob follows Alice
BOB_ADDR=$(nlgd keys show bob -a)
ALICE_ADDR=$(nlgd keys show alice -a)

nlgd tx socialmedia create-social-connection \
  $BOB_ADDR \
  $ALICE_ADDR \
  $(date +%s) \
  --from bob \
  --yes

# Charlie follows Bob
CHARLIE_ADDR=$(nlgd keys show charlie -a)

nlgd tx socialmedia create-social-connection \
  $CHARLIE_ADDR \
  $BOB_ADDR \
  $(date +%s) \
  --from charlie \
  --yes
```

### Example 8: Query Social Graph

```bash
# Get all connections
nlgd query socialmedia list-social-connection

# Find who Bob is following
BOB_ADDR=$(nlgd keys show bob -a)
nlgd query socialmedia list-social-connection | \
  jq --arg addr "$BOB_ADDR" '.socialConnection[] | select(.follower==$addr)'

# Find Bob's followers
nlgd query socialmedia list-social-connection | \
  jq --arg addr "$BOB_ADDR" '.socialConnection[] | select(.following==$addr)'
```

### Example 9: Comment on a Post

```bash
# Bob comments on Alice's post (ID: 0)
BOB_ADDR=$(nlgd keys show bob -a)

nlgd tx socialmedia create-comment \
  "0" \
  "Great post Alice! I'm really excited about the potential of decentralized social media." \
  $(date +%s) \
  $BOB_ADDR \
  --from bob \
  --yes

# Charlie also comments
CHARLIE_ADDR=$(nlgd keys show charlie -a)

nlgd tx socialmedia create-comment \
  "0" \
  "This is exactly what we need. Looking forward to the EVM integration!" \
  $(date +%s) \
  $CHARLIE_ADDR \
  --from charlie \
  --yes
```

### Example 10: Like Posts

```bash
# Bob likes Alice's post
nlgd tx socialmedia like-post "0" --from bob --yes

# Charlie also likes it
nlgd tx socialmedia like-post "0" --from charlie --yes

# Query updated post to see likes count
nlgd query socialmedia show-post 0
```

## Content Moderation

### Example 11: Flag Inappropriate Content

```bash
# User flags a post as inappropriate
nlgd tx socialmedia moderate-content "post-0" "flag" \
  --from charlie \
  --yes

# Multiple users can vote
nlgd tx socialmedia moderate-content "post-0" "flag" \
  --from bob \
  --yes
```

### Example 12: Approve Quality Content

```bash
# Vote to approve good content
nlgd tx socialmedia moderate-content "post-1" "approve" \
  --from alice \
  --yes

nlgd tx socialmedia moderate-content "post-1" "approve" \
  --from bob \
  --yes
```

## Advanced Scenarios

### Example 13: Batch Operations Script

```bash
#!/bin/bash
# batch_posts.sh - Create multiple posts

for i in {1..5}; do
  nlgd tx socialmedia create-post \
    "Post $i" \
    "Content for post number $i" \
    "QmHash$i" \
    $(date +%s) \
    0 \
    --from alice \
    --yes \
    --chain-id nlg-1
  
  echo "Created post $i"
  sleep 6  # Wait for block confirmation
done
```

### Example 14: Monitor New Posts

```bash
#!/bin/bash
# monitor_posts.sh - Watch for new posts

LAST_COUNT=0

while true; do
  CURRENT_COUNT=$(curl -s http://localhost:1317/nlg/socialmedia/post | jq '.post | length')
  
  if [ "$CURRENT_COUNT" -gt "$LAST_COUNT" ]; then
    echo "New post detected! Total posts: $CURRENT_COUNT"
    LAST_COUNT=$CURRENT_COUNT
  fi
  
  sleep 10
done
```

### Example 15: Export User Data

```bash
#!/bin/bash
# export_user_data.sh - Export all data for a user

USER_ADDR=$(nlgd keys show alice -a)
OUTPUT_FILE="alice_data.json"

echo "{" > $OUTPUT_FILE

# Export profile
echo "  \"profile\": " >> $OUTPUT_FILE
curl -s http://localhost:1317/nlg/socialmedia/profile | \
  jq --arg addr "$USER_ADDR" '.profile[] | select(.creator==$addr)' >> $OUTPUT_FILE

echo "," >> $OUTPUT_FILE

# Export posts
echo "  \"posts\": " >> $OUTPUT_FILE
curl -s http://localhost:1317/nlg/socialmedia/post | \
  jq --arg addr "$USER_ADDR" '[.post[] | select(.creator==$addr)]' >> $OUTPUT_FILE

echo "," >> $OUTPUT_FILE

# Export comments
echo "  \"comments\": " >> $OUTPUT_FILE
curl -s http://localhost:1317/nlg/socialmedia/comment | \
  jq --arg addr "$USER_ADDR" '[.comment[] | select(.author==$addr)]' >> $OUTPUT_FILE

echo "}" >> $OUTPUT_FILE

echo "User data exported to $OUTPUT_FILE"
```

### Example 16: Token Transfer Before Post Creation

```bash
# Charlie needs tokens to pay for transactions
nlgd tx bank send bob $(nlgd keys show charlie -a) 1000000nlg \
  --yes \
  --chain-id nlg-1

# Now Charlie can create posts
nlgd tx socialmedia create-post \
  "My First Post" \
  "Thanks Bob for the tokens!" \
  "QmHash..." \
  $(date +%s) \
  0 \
  --from charlie \
  --yes
```

### Example 17: Query with Pagination

```bash
# Get first 5 posts
curl "http://localhost:1317/nlg/socialmedia/post?pagination.limit=5" | jq .

# Get next 5 posts
curl "http://localhost:1317/nlg/socialmedia/post?pagination.limit=5&pagination.offset=5" | jq .

# Get total count without data
curl "http://localhost:1317/nlg/socialmedia/post?pagination.count_total=true&pagination.limit=1" | \
  jq '.pagination.total'
```

### Example 18: Building a Simple Feed

```bash
#!/bin/bash
# feed.sh - Display a simple social media feed

echo "=== NLG Social Media Feed ==="
echo ""

# Get all posts
curl -s http://localhost:1317/nlg/socialmedia/post | \
  jq -r '.post[] | "[\(.id)] \(.title)\nBy: \(.creator)\nLikes: \(.likesCount)\n---"'

echo ""
echo "Total posts: $(curl -s http://localhost:1317/nlg/socialmedia/post | jq '.post | length')"
```

### Example 19: Governance Proposal for Parameter Change

```bash
# Create a proposal to change module parameters
nlgd tx gov submit-proposal param-change proposal.json \
  --from alice \
  --deposit 10000000nlg \
  --yes

# proposal.json content:
cat > proposal.json << EOF
{
  "title": "Update Social Media Module Parameters",
  "description": "Increase max post length to 10000 characters",
  "changes": [
    {
      "subspace": "socialmedia",
      "key": "MaxPostLength",
      "value": "10000"
    }
  ]
}
EOF

# Vote on proposal
nlgd tx gov vote 1 yes --from alice --yes
nlgd tx gov vote 1 yes --from bob --yes
```

### Example 20: WebSocket Event Subscription

```javascript
// subscribe.js - Subscribe to new post events
const WebSocket = require('ws');

const ws = new WebSocket('ws://localhost:26657/websocket');

ws.on('open', () => {
  // Subscribe to all transactions
  const subscribeMsg = {
    jsonrpc: '2.0',
    method: 'subscribe',
    id: 1,
    params: {
      query: "tm.event='Tx'"
    }
  };
  
  ws.send(JSON.stringify(subscribeMsg));
  console.log('Subscribed to new transactions');
});

ws.on('message', (data) => {
  const event = JSON.parse(data);
  
  if (event.result && event.result.data) {
    console.log('New transaction:', event.result.data);
  }
});
```

## Testing Scenarios

### Example 21: Load Testing

```bash
#!/bin/bash
# load_test.sh - Simple load test

NUM_POSTS=100
CONCURRENT=10

echo "Creating $NUM_POSTS posts with $CONCURRENT concurrent users"

for i in $(seq 1 $NUM_POSTS); do
  (
    nlgd tx socialmedia create-post \
      "Load Test Post $i" \
      "This is load test post number $i" \
      "QmLoadTest$i" \
      $(date +%s) \
      0 \
      --from $([ $((i % 3)) -eq 0 ] && echo "alice" || ([ $((i % 3)) -eq 1 ] && echo "bob" || echo "charlie")) \
      --yes \
      --broadcast-mode async \
      &> /dev/null
  ) &
  
  if [ $((i % CONCURRENT)) -eq 0 ]; then
    wait
  fi
done

wait
echo "Load test complete"
```

## Best Practices

1. **Always use `--gas auto`** for automatic gas estimation
2. **Add `--gas-adjustment 1.3`** for buffer in gas calculation
3. **Use `--yes`** flag for non-interactive mode in scripts
4. **Check transaction hash** after submission for confirmation
5. **Wait for block confirmation** before querying new data
6. **Use pagination** for large datasets
7. **Implement error handling** in production scripts
8. **Store private keys securely** - never commit them

## Troubleshooting Examples

### Check Transaction Status

```bash
# Get transaction by hash
nlgd query tx <TX_HASH>

# Check if transaction succeeded
nlgd query tx <TX_HASH> | jq '.code'
# 0 = success, non-zero = error
```

### Verify Account Balance

```bash
# Check if user has enough tokens
nlgd query bank balances $(nlgd keys show alice -a)
```

### Debug Failed Transactions

```bash
# Get detailed error information
nlgd query tx <TX_HASH> | jq '.raw_log'
```

## Additional Resources

- See [API.md](docs/API.md) for complete API reference
- See [README.md](README.md) for detailed documentation
- Check [DEVELOPMENT.md](docs/DEVELOPMENT.md) for development guidelines

---

For more examples and use cases, join our Discord community or check GitHub discussions.
