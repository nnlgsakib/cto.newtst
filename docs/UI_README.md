# NLG Chain Inbuilt UI

The NLG blockchain comes with a comprehensive inbuilt web interface that provides real-time monitoring and interaction capabilities.

## Features

### Dashboard
- **Real-time Block Explorer**: View latest blocks, transactions, and network status
- **Network Statistics**: Monitor block height, active validators, token supply, and transaction volume
- **Auto-refresh**: Data updates automatically every 10 seconds

### Social Feed
- **Decentralized Social**: View posts from the NLG social media module
- **Real-time Updates**: Live feed of social content on the blockchain
- **Post Details**: View creator, timestamp, likes, comments, and reputation

### Block Explorer
- **Block Details**: View comprehensive information about each block
- **Transaction List**: See all transactions included in each block
- **Proposer Information**: Track block proposers and their activity

### Validator Monitor
- **Validator List**: View all active validators on the network
- **Staking Information**: See token delegations and commission rates
- **Ranking**: Validators ranked by stake

### API Documentation
- **Interactive Swagger UI**: Explore all available API endpoints
- **Try It Out**: Test API calls directly from the browser
- **Auto-generated**: Documentation automatically generated from OpenAPI spec

## Accessing the UI

The UI is built into the NLG blockchain binary and runs automatically when you start the node.

### Start the Node
```bash
nlgd start
```

### Access the Dashboard
Open your browser and navigate to:
```
http://localhost:1317/dashboard
```

### Access API Documentation
```
http://localhost:1317/
```

## Configuration

The UI uses the default API server configuration:

- **Default Port**: 1317
- **Dashboard Path**: /dashboard
- **API Docs Path**: /
- **API Endpoints**: /cosmos/* and /nlg/*

### Custom Port
To use a custom port, set the API address in your configuration:
```bash
nlgd start --api.address tcp://0.0.0.0:26657
```

## UI Pages

1. **Dashboard** (`/dashboard#dashboard`)
   - Overview of network statistics
   - Recent blocks and transactions
   - Social feed preview

2. **Blocks** (`/dashboard#blocks`)
   - List of recent blocks
   - Click on a block for detailed view

3. **Validators** (`/dashboard#validators`)
   - All active validators
   - Staking and commission information

4. **Social Feed** (`/dashboard#social`)
   - Decentralized social media posts
   - Real-time content updates

5. **API Documentation** (`/dashboard#api` or `/`)
   - Full API reference
   - Interactive testing interface

## Development

### UI Structure
```
docs/
├── static/
│   ├── css/
│   │   └── styles.css          # Main stylesheet
│   ├── js/
│   │   └── app.js             # Dashboard JavaScript
│   └── openapi.json           # API specification
└── template/
    ├── index.tpl              # API docs template
    └── dashboard.tpl          # Main dashboard template
```

### Customization

To customize the UI, modify the following files:

1. **Styling**: `docs/static/css/styles.css`
   - Change colors, fonts, and layout
   - Responsive design settings

2. **Functionality**: `docs/static/js/app.js`
   - API calls and data fetching
   - UI updates and interactions
   - Navigation logic

3. **Layout**: `docs/template/dashboard.tpl`
   - HTML structure
   - Page layouts
   - Navigation elements

### Adding New Features

1. Add new API endpoints in the appropriate module
2. Update `docs/static/js/app.js` to fetch the new data
3. Modify `docs/template/dashboard.tpl` to display it

## API Integration

The UI communicates with the blockchain through REST API endpoints:

### Core Cosmos SDK Endpoints
- `/cosmos/base/tendermint/v1beta1/*` - Block and node information
- `/cosmos/bank/v1beta1/*` - Token balances
- `/cosmos/staking/v1beta1/*` - Validator information
- `/cosmos/auth/v1beta1/*` - Account information

### NLG Module Endpoints
- `/nlg/socialmedia/v1/*` - Social media module
- `/nlg/socialmedia/v1/posts` - List all posts
- `/nlg/socialmedia/v1/profile/{address}` - Get user profile

## Browser Compatibility

The UI works with all modern browsers:
- Chrome/Edge (recommended)
- Firefox
- Safari
- Opera

Minimum versions:
- Chrome 90+
- Firefox 88+
- Safari 14+
- Edge 90+

## Performance

- **Auto-refresh**: 10-second intervals
- **API caching**: Minimal client-side caching for better performance
- **Responsive**: Works on desktop, tablet, and mobile devices
- **Lightweight**: Pure vanilla JavaScript, no external dependencies (except Swagger UI)

## Troubleshooting

### UI Not Loading
1. Ensure the node is running: `nlgd start`
2. Check if API server is enabled (enabled by default)
3. Verify correct port: http://localhost:1317/dashboard

### Data Not Updating
1. Check if the node is syncing properly
2. Verify API server is responding
3. Check browser console for errors

### API Errors
1. Check the API logs: `nlgd logs`
2. Verify CORS settings if accessing from different domain
3. Check firewall settings

## Security Notes

- The UI is read-only for blockchain data
- No private keys are stored or transmitted
- All API calls are made from the browser to the local node
- Consider adding authentication for production deployments

## Future Enhancements

Planned features for future versions:
- Wallet integration
- Transaction signing from the UI
- Governance proposal voting interface
- Advanced charts and analytics
- Real-time WebSocket updates
- Dark mode theme
- Multi-language support

## Support

For issues or questions:
1. Check the main [README.md](../README.md)
2. Review [DEVELOPMENT.md](../DEVELOPMENT.md)
3. Check [API.md](../API.md) for API documentation
4. Open an issue on GitHub

---

Built with Cosmos SDK and ❤️ for the decentralized future.
