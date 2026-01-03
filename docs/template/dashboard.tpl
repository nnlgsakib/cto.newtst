<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="utf-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>NLG Chain - Dashboard</title>
    <link rel="stylesheet" href="/static/css/styles.css" />
    <link rel="icon" type="image/svg+xml" href="data:image/svg+xml,<svg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 24 24' fill='%233b82f6'><path d='M12 2L2 7l10 5 10-5-10-5zM2 17l10 5 10-5M2 12l10 5 10-5'/></svg>">
</head>
<body>
    <!-- Navigation -->
    <nav>
        <div class="container">
            <div class="logo">
                <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor">
                    <path d="M12 2L2 7l10 5 10-5-10-5zM2 17l10 5 10-5M2 12l10 5 10-5"/>
                </svg>
                <span>NLG Chain</span>
            </div>
            <ul>
                <li><a href="#dashboard">
                    <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                        <rect x="3" y="3" width="7" height="7"></rect>
                        <rect x="14" y="3" width="7" height="7"></rect>
                        <rect x="14" y="14" width="7" height="7"></rect>
                        <rect x="3" y="14" width="7" height="7"></rect>
                    </svg>
                    Dashboard
                </a></li>
                <li><a href="#blocks">
                    <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                        <rect x="2" y="7" width="20" height="14" rx="2" ry="2"></rect>
                        <path d="M16 21V5a2 2 0 0 0-2-2h-4a2 2 0 0 0-2 2v16"></path>
                    </svg>
                    Blocks
                </a></li>
                <li><a href="#validators">
                    <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                        <path d="M12 22s8-4 8-10V5l-8-3-8 3v7c0 6 8 10 8 10z"></path>
                    </svg>
                    Validators
                </a></li>
                <li><a href="#social">
                    <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                        <path d="M17 21v-2a4 4 0 0 0-4-4H5a4 4 0 0 0-4 4v2"></path>
                        <circle cx="9" cy="7" r="4"></circle>
                        <path d="M23 21v-2a4 4 0 0 0-3-3.87"></path>
                        <path d="M16 3.13a4 4 0 0 1 0 7.75"></path>
                    </svg>
                    Social Feed
                </a></li>
                <li><a href="#api">
                    <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                        <polyline points="16 18 22 12 16 6"></polyline>
                        <polyline points="8 6 2 12 8 18"></polyline>
                    </svg>
                    API Docs
                </a></li>
            </ul>
        </div>
    </nav>

    <!-- Hero Section -->
    <div class="hero">
        <div class="container">
            <h1>NLG Chain Dashboard</h1>
            <p>Explore the decentralized social blockchain with real-time data</p>
        </div>
    </div>

    <!-- Main Content -->
    <div class="container" style="padding-top: 2rem; padding-bottom: 2rem;">
        
        <!-- Dashboard Page -->
        <div id="dashboard-page" class="page">
            <div class="stats-grid">
                <div class="card">
                    <div class="value" id="block-height">-</div>
                    <div class="label">Current Block Height</div>
                </div>
                <div class="card">
                    <div class="value" id="validator-count">-</div>
                    <div class="label">Active Validators</div>
                </div>
                <div class="card">
                    <div class="value" id="total-supply">-</div>
                    <div class="label">Total Token Supply</div>
                </div>
                <div class="card">
                    <div class="value" id="tx-count">-</div>
                    <div class="label">Transactions Today</div>
                </div>
            </div>

            <h2 style="margin: 2rem 0 1rem;">Recent Social Posts</h2>
            <div id="social-feed-dashboard"></div>

            <h2 style="margin: 2rem 0 1rem;">Recent Blocks</h2>
            <div class="card">
                <table>
                    <thead>
                        <tr>
                            <th>Height</th>
                            <th>Proposer</th>
                            <th>Txs</th>
                            <th>Time</th>
                        </tr>
                    </thead>
                    <tbody id="recent-blocks">
                        <tr><td colspan="4"><div class="loading">Loading...</div></td></tr>
                    </tbody>
                </table>
            </div>
        </div>

        <!-- Blocks Page -->
        <div id="blocks-page" class="page" style="display: none;">
            <h2 style="margin-bottom: 1rem;">Latest Blocks</h2>
            <div class="card">
                <table>
                    <thead>
                        <tr>
                            <th>Height</th>
                            <th>Hash</th>
                            <th>Proposer</th>
                            <th>Txs</th>
                            <th>Time</th>
                        </tr>
                    </thead>
                    <tbody id="blocks-list">
                        <tr><td colspan="5"><div class="loading">Loading...</div></td></tr>
                    </tbody>
                </table>
            </div>
        </div>

        <!-- Validators Page -->
        <div id="validators-page" class="page" style="display: none;">
            <h2 style="margin-bottom: 1rem;">Active Validators</h2>
            <div class="card">
                <table>
                    <thead>
                        <tr>
                            <th>Rank</th>
                            <th>Operator Address</th>
                            <th>Moniker</th>
                            <th>Tokens</th>
                            <th>Commission</th>
                        </tr>
                    </thead>
                    <tbody id="validators-list">
                        <tr><td colspan="5"><div class="loading">Loading...</div></td></tr>
                    </tbody>
                </table>
            </div>
        </div>

        <!-- Social Feed Page -->
        <div id="social-page" class="page" style="display: none;">
            <div style="display: flex; justify-content: space-between; align-items: center; margin-bottom: 1rem;">
                <h2>Social Feed</h2>
                <button class="btn" onclick="alert('Connect your wallet to create a post')">Create Post</button>
            </div>
            <div id="social-feed"></div>
        </div>

        <!-- Block Detail Page -->
        <div id="block-detail-page" class="page" style="display: none;">
            <div style="margin-bottom: 1rem;">
                <a href="#blocks" style="color: var(--primary-color); text-decoration: none;">← Back to Blocks</a>
            </div>
            <div id="block-detail">
                <div class="loading">Loading block details...</div>
            </div>
        </div>

        <!-- API Documentation Page -->
        <div id="api-page" class="page" style="display: none;">
            <h2 style="margin-bottom: 1rem;">API Documentation</h2>
            <div id="swagger-ui"></div>
        </div>
    </div>

    <!-- Footer -->
    <footer>
        <div class="container">
            <p>&copy; 2024 NLG Chain. Built with Cosmos SDK.</p>
            <p style="margin-top: 0.5rem; opacity: 0.7;">
                <a href="#dashboard" style="color: white; margin: 0 1rem;">Dashboard</a>
                <a href="#api" style="color: white; margin: 0 1rem;">API Docs</a>
                <a href="https://github.com/cosmos/cosmos-sdk" style="color: white; margin: 0 1rem;">GitHub</a>
            </p>
        </div>
    </footer>

    <!-- Scripts -->
    <script src="/static/js/app.js"></script>
    <script src="//unpkg.com/swagger-ui-dist@3.40.0/swagger-ui-bundle.js"></script>
    <script>
        // Initialize Swagger UI when API page is shown
        window.addEventListener('hashchange', function() {
            if (window.location.hash === '#api') {
                setTimeout(function() {
                    window.ui = SwaggerUIBundle({
                        url: '/static/openapi.json',
                        dom_id: '#swagger-ui',
                        deepLinking: true,
                        layout: "BaseLayout",
                    });
                }, 100);
            }
        });
    </script>
</body>
</html>
