# NLG Chain UI Quick Start Guide

Get started with the NLG Chain web dashboard in minutes!

## Prerequisites

Ensure you have the NLG Chain binary installed and ready to run.

## Step 1: Start the Chain

Start your local NLG Chain development node:

```bash
# Using Ignite CLI (recommended for development)
ignite chain serve

# Or using the binary directly
nlgd start
```

Wait for the node to start producing blocks. You should see logs indicating:
- "Starting node"
- "Service started"
- Block production messages

## Step 2: Access the Dashboard

Open your web browser and navigate to:

```
http://localhost:1317/dashboard
```

You should see the NLG Chain Dashboard with:
- Network statistics
- Recent blocks
- Social feed preview
- Navigation menu

## Step 3: Explore the Interface

### Dashboard Page (Default)

The main dashboard shows:
- **Block Height**: Current blockchain height
- **Validator Count**: Number of active validators
- **Total Supply**: Total NLG tokens in circulation
- **Transactions Today**: Daily transaction volume

### Blocks Page

Click "Blocks" in the navigation menu to:
- View list of recent blocks
- Click on any block to see detailed information
- See proposer, transaction count, and timestamp

### Validators Page

Click "Validators" in the navigation menu to:
- View all active validators
- See staking amounts and commission rates
- Monitor validator rankings

### Social Feed Page

Click "Social Feed" in the navigation menu to:
- View decentralized social media posts
- See post creators, content, and engagement metrics
- Watch the feed update in real-time

### API Documentation

Click "API Docs" or visit `http://localhost:1317/` to:
- Explore all available API endpoints
- Test API calls directly from the browser
- View request/response formats

## Step 4: Interact with the Chain

### View Account Balances

1. Open API Documentation (`/dashboard#api`)
2. Find the `/cosmos/bank/v1beta1/balances/{address}` endpoint
3. Enter an address (e.g., one of the genesis accounts)
4. Click "Try it out" to query the balance

### Query Social Posts

1. Navigate to Social Feed page
2. Posts automatically load from the blockchain
3. Click on the API Docs to explore more endpoints

### Create a Social Post

Using the CLI:

```bash
# Create a post (from alice)
nlgd tx socialmedia create-post \
  "My First Post" \
  "Hello NLG Chain! This is my first post." \
  "" \
  $(date +%s) \
  0 \
  --from alice \
  --yes
```

Refresh the Social Feed page to see your new post!

### Create a Profile

```bash
# Create a profile for alice
nlgd tx socialmedia create-profile \
  "alice" \
  "Blockchain enthusiast" \
  "" \
  100 \
  0 \
  0 \
  --from alice \
  --yes
```

## Step 5: Monitor in Real-Time

The dashboard automatically refreshes every 10 seconds:
- Watch the block height increase
- See new transactions appear
- View new social posts as they're created

## Common Use Cases

### For Developers

1. **API Testing**: Use the Swagger UI to test your integration
2. **Transaction Monitoring**: Watch transactions as they propagate
3. **Block Analysis**: Examine block structure and content

### For Validators

1. **Performance Tracking**: Monitor your validator's rank and status
2. **Network Health**: Check block production and transaction volume
3. **Peer Monitoring**: View other validators on the network

### For Users

1. **Social Interaction**: Browse and interact with social content
2. **Balance Checking**: View token balances via API
3. **Transaction History**: Track your transactions

## Tips and Tricks

### Bookmark Pages

Bookmark specific sections for quick access:
- Dashboard: `http://localhost:1317/dashboard#dashboard`
- Blocks: `http://localhost:1317/dashboard#blocks`
- Validators: `http://localhost:1317/dashboard#validators`
- Social: `http://localhost:1317/dashboard#social`

### Multiple Tabs

Open multiple tabs to monitor different views simultaneously:
- Tab 1: Dashboard (overview)
- Tab 2: Blocks (live explorer)
- Tab 3: Social Feed (social activity)
- Tab 4: API Docs (development)

### Developer Console

Press F12 to open the browser console and see:
- API request logs
- Error messages
- Performance data

## Troubleshooting

### Dashboard Won't Load

**Problem**: Browser shows "Connection refused"

**Solution**:
1. Check if the node is running: `nlgd status`
2. Verify API server is enabled (default: enabled)
3. Check the port: Ensure port 1317 is available

**Alternative ports**: If you changed the API port, use:
```
http://localhost:YOUR_PORT/dashboard
```

### Data Not Updating

**Problem**: Block height or stats stay the same

**Solution**:
1. Check if the node is syncing: `nlgd status`
2. Look for errors in the node logs
3. Wait a few seconds (auto-refreshes every 10s)
4. Refresh the browser page (F5)

### Social Feed Empty

**Problem**: No posts showing in the social feed

**Solution**:
1. Create a post using the CLI (see Step 4)
2. Check the API is responding: `curl http://localhost:1317/nlg/socialmedia/v1/posts`
3. Verify the socialmedia module is enabled

### API Errors

**Problem**: "404 Not Found" or "500 Internal Server Error"

**Solution**:
1. Check the exact endpoint path in API Docs
2. Ensure the node is fully synced
3. Check node logs for error details

## Next Steps

- Read the [full UI documentation](UI_README.md) for advanced features
- Check the [API documentation](API.md) for all available endpoints
- Review [DEVELOPMENT.md](../DEVELOPMENT.md) for building custom features
- Explore the social media module to understand content moderation

## Getting Help

If you encounter issues:

1. Check the main [README.md](../README.md)
2. Review troubleshooting sections in documentation
3. Check browser console for errors (F12)
4. Verify the node is running correctly
5. Open an issue on GitHub with:
   - Your operating system and browser
   - Error messages or screenshots
   - Steps to reproduce

## Security Notes

- The dashboard connects to your local node
- No private keys are stored in the browser
- All interactions are read-only (except via CLI)
- Consider adding authentication for production deployments

---

Happy exploring! 🚀
