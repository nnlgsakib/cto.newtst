// NLG Chain Dashboard JavaScript

const API_BASE = window.location.origin;

// State management
let currentState = {
  blockHeight: 0,
  totalTxs: 0,
  validators: [],
  recentBlocks: [],
  recentTxs: [],
  posts: []
};

// API calls
async function fetchWithFallback(url, fallback = null) {
  try {
    const response = await fetch(url);
    if (!response.ok) throw new Error(`HTTP ${response.status}`);
    return await response.json();
  } catch (error) {
    console.warn(`Failed to fetch ${url}:`, error);
    return fallback;
  }
}

async function getStatus() {
  return fetchWithFallback(`${API_BASE}/cosmos/base/tendermint/v1beta1/node_info`);
}

async function getBlock(height = null) {
  const endpoint = height
    ? `${API_BASE}/cosmos/base/tendermint/v1beta1/blocks/${height}`
    : `${API_BASE}/cosmos/base/tendermint/v1beta1/blocks/latest`;
  return fetchWithFallback(endpoint);
}

async function getValidatorSet() {
  return fetchWithFallback(`${API_BASE}/cosmos/staking/v1beta1/validators?status=BOND_STATUS_BONDED`);
}

async function getAccount(address) {
  return fetchWithFallback(`${API_BASE}/cosmos/auth/v1beta1/accounts/${address}`);
}

async function getBalance(address, denom = 'nlg') {
  return fetchWithFallback(`${API_BASE}/cosmos/bank/v1beta1/balances/${address}`);
}

async function getTxsByEvents(events, page = 1, limit = 10) {
  return fetchWithFallback(`${API_BASE}/cosmos/tx/v1beta1/txs?events=${encodeURIComponent(events)}&pagination.limit=${limit}&pagination.offset=${(page - 1) * limit}`);
}

async function getSocialPosts() {
  return fetchWithFallback(`${API_BASE}/nlg/socialmedia/v1/posts`);
}

async function getSocialProfile(address) {
  return fetchWithFallback(`${API_BASE}/nlg/socialmedia/v1/profile/${address}`);
}

// UI Updates
function updateBlockHeight(height) {
  const element = document.getElementById('block-height');
  if (element) {
    element.textContent = height.toLocaleString();
    if (currentState.blockHeight > 0) {
      element.classList.add('updated');
      setTimeout(() => element.classList.remove('updated'), 500);
    }
    currentState.blockHeight = height;
  }
}

function updateValidatorCount(count) {
  const element = document.getElementById('validator-count');
  if (element) {
    element.textContent = count.toLocaleString();
  }
}

function updateTotalSupply(supply) {
  const element = document.getElementById('total-supply');
  if (element) {
    const formatted = supply / 1000000; // Convert from unlg to nlg
    element.textContent = `${formatted.toLocaleString()} nlg`;
  }
}

function renderRecentBlocks(blocks) {
  const container = document.getElementById('recent-blocks');
  if (!container) return;

  if (!blocks || blocks.length === 0) {
    container.innerHTML = '<div class="loading">No blocks available</div>';
    return;
  }

  const html = blocks.map(block => `
    <tr>
      <td><a href="#block/${block.block?.header?.height || 'N/A'}" class="hash">${block.block?.header?.height || 'N/A'}</a></td>
      <td><span class="hash">${block.block?.header?.proposer_address || 'N/A'}</span></td>
      <td>${block.block?.data?.txs?.length || 0}</td>
      <td>${new Date(block.block?.header?.time || Date.now()).toLocaleString()}</td>
    </tr>
  `).join('');

  container.innerHTML = html;
}

function renderValidators(validators) {
  const container = document.getElementById('validators');
  if (!container) return;

  if (!validators || validators.length === 0) {
    container.innerHTML = '<div class="loading">No validators available</div>';
    return;
  }

  const html = validators.slice(0, 10).map(validator => `
    <tr>
      <td>${validator.rank || 'N/A'}</td>
      <td><span class="hash">${validator.operator_address || 'N/A'}</span></td>
      <td>${validator.description?.moniker || 'Unknown'}</td>
      <td>${(parseFloat(validator.tokens || 0) / 1000000).toLocaleString()} nlg</td>
      <td>${validator.commission?.commission_rates?.rate ? (parseFloat(validator.commission.commission_rates.rate) * 100).toFixed(2) : 'N/A'}%</td>
    </tr>
  `).join('');

  container.innerHTML = html;
}

function renderSocialPosts(posts) {
  const container = document.getElementById('social-feed-dashboard');
  if (!container) return;

  if (!posts || posts.length === 0) {
    container.innerHTML = '<div class="loading">No posts available. Connect your wallet to create your first post!</div>';
    return;
  }

  const html = posts.slice(0, 5).map(post => `
    <div class="post-card">
      <div class="post-header">
        <div class="post-avatar">${(post.creator || 'U').substring(0, 2).toUpperCase()}</div>
        <div>
          <div class="post-author"><span class="hash">${post.creator || 'Anonymous'}</span></div>
          <div class="post-time">${new Date(post.created_at || Date.now()).toLocaleString()}</div>
        </div>
      </div>
      <div class="post-content">${post.content || post.title || 'No content'}</div>
      <div class="post-footer">
        <span>💬 ${post.comments_count || 0} comments</span>
        <span>❤️ ${post.likes_count || 0} likes</span>
        <span>🏆 ${post.reputation || 0} rep</span>
      </div>
    </div>
  `).join('');

  container.innerHTML = html;
}

// Main update function
async function updateDashboard() {
  console.log('Updating dashboard...');

  // Fetch block info
  const block = await getBlock();
  if (block && block.block) {
    updateBlockHeight(block.block.header.height);
  }

  // Fetch validators
  const validators = await getValidatorSet();
  if (validators && validators.validators) {
    currentState.validators = validators.validators;
    updateValidatorCount(validators.validators.length);
    renderValidators(validators.validators);
  }

  // Fetch social posts
  const posts = await getSocialPosts();
  if (posts && posts.posts) {
    currentState.posts = posts.posts;
    renderSocialPosts(posts.posts);
  }
}

// Navigation
function navigateTo(page) {
  window.location.hash = page;
}

function handleNavigation() {
  const hash = window.location.hash;

  // Hide all pages
  document.querySelectorAll('.page').forEach(page => {
    page.style.display = 'none';
  });

  // Show active page
  if (hash === '#dashboard' || hash === '') {
    document.getElementById('dashboard-page').style.display = 'block';
    updateDashboard();
  } else if (hash === '#blocks') {
    document.getElementById('blocks-page').style.display = 'block';
    loadBlocksList();
  } else if (hash === '#validators') {
    document.getElementById('validators-page').style.display = 'block';
    loadValidatorsList();
  } else if (hash === '#social') {
    document.getElementById('social-page').style.display = 'block';
    loadSocialFeed();
  } else if (hash.startsWith('#block/')) {
    document.getElementById('block-detail-page').style.display = 'block';
    const height = hash.split('/')[1];
    loadBlockDetail(height);
  } else if (hash === '#api') {
    document.getElementById('api-page').style.display = 'block';
  }
}

async function loadBlockDetail(height) {
  const block = await getBlock(height);
  const container = document.getElementById('block-detail');
  
  if (!block || !block.block) {
    container.innerHTML = '<div class="loading">Block not found</div>';
    return;
  }

  const html = `
    <div class="card">
      <h3>Block #${block.block.header.height}</h3>
      <table>
        <tr><th>Height</th><td>${block.block.header.height}</td></tr>
        <tr><th>Hash</th><td><span class="hash">${block.block_id?.hash || 'N/A'}</span></td></tr>
        <tr><th>Proposer</th><td><span class="hash">${block.block.header.proposer_address}</span></td></tr>
        <tr><th>Timestamp</th><td>${new Date(block.block.header.time).toLocaleString()}</td></tr>
        <tr><th>Transactions</th><td>${block.block.data?.txs?.length || 0}</td></tr>
        <tr><th>Last Block Hash</th><td><span class="hash">${block.block.header.last_block_id?.hash || 'N/A'}</span></td></tr>
      </table>
    </div>
  `;

  container.innerHTML = html;
}

async function loadBlocksList() {
  const container = document.getElementById('blocks-list');
  if (!container) return;

  const block = await getBlock();
  if (!block || !block.block) {
    container.innerHTML = '<tr><td colspan="5"><div class="loading">Failed to load blocks</div></td></tr>';
    return;
  }

  const currentHeight = parseInt(block.block.header.height);
  const blocks = [];

  // Fetch last 10 blocks
  for (let i = 0; i < 10; i++) {
    const h = currentHeight - i;
    if (h < 1) break;
    const b = await getBlock(h);
    if (b && b.block) {
      blocks.push(b);
    }
  }

  const html = blocks.map(b => `
    <tr>
      <td><a href="#block/${b.block.header.height}" style="color: var(--primary-color); text-decoration: none;">${b.block.header.height}</a></td>
      <td><span class="hash">${b.block_id?.hash || 'N/A'}</span></td>
      <td><span class="hash">${b.block.header.proposer_address || 'N/A'}</span></td>
      <td>${b.block.data?.txs?.length || 0}</td>
      <td>${new Date(b.block.header.time).toLocaleString()}</td>
    </tr>
  `).join('');

  container.innerHTML = html;
}

async function loadValidatorsList() {
  const container = document.getElementById('validators-list');
  if (!container) return;

  const validators = await getValidatorSet();
  if (!validators || !validators.validators) {
    container.innerHTML = '<tr><td colspan="5"><div class="loading">Failed to load validators</div></td></tr>';
    return;
  }

  const html = validators.validators.map((validator, index) => `
    <tr>
      <td>${index + 1}</td>
      <td><span class="hash">${validator.operator_address || 'N/A'}</span></td>
      <td>${validator.description?.moniker || 'Unknown'}</td>
      <td>${(parseFloat(validator.tokens || 0) / 1000000).toLocaleString()} nlg</td>
      <td>${validator.commission?.commission_rates?.rate ? (parseFloat(validator.commission.commission_rates.rate) * 100).toFixed(2) : 'N/A'}%</td>
    </tr>
  `).join('');

  container.innerHTML = html;
}

async function loadSocialFeed() {
  const container = document.getElementById('social-feed');
  if (!container) return;

  const posts = await getSocialPosts();
  if (!posts || !posts.posts) {
    container.innerHTML = '<div class="loading">Failed to load posts</div>';
    return;
  }

  if (posts.posts.length === 0) {
    container.innerHTML = '<div class="loading">No posts available. Connect your wallet to create your first post!</div>';
    return;
  }

  const html = posts.posts.map(post => `
    <div class="post-card">
      <div class="post-header">
        <div class="post-avatar">${(post.creator || 'U').substring(0, 2).toUpperCase()}</div>
        <div>
          <div class="post-author"><span class="hash">${post.creator || 'Anonymous'}</span></div>
          <div class="post-time">${new Date(post.created_at || Date.now()).toLocaleString()}</div>
        </div>
      </div>
      <div class="post-content">${post.content || post.title || 'No content'}</div>
      <div class="post-footer">
        <span>💬 ${post.comments_count || 0} comments</span>
        <span>❤️ ${post.likes_count || 0} likes</span>
        <span>🏆 ${post.reputation || 0} rep</span>
      </div>
    </div>
  `).join('');

  container.innerHTML = html;
}

// Initialize
document.addEventListener('DOMContentLoaded', () => {
  console.log('NLG Chain Dashboard initialized');
  
  // Handle initial navigation
  handleNavigation();
  
  // Listen for hash changes
  window.addEventListener('hashchange', handleNavigation);
  
  // Auto-update every 10 seconds
  setInterval(updateDashboard, 10000);
  
  // Initial update
  updateDashboard();
});

// Utility functions
function formatAddress(address) {
  if (!address) return 'N/A';
  return `${address.substring(0, 10)}...${address.substring(address.length - 6)}`;
}

function formatHash(hash) {
  if (!hash) return 'N/A';
  return `${hash.substring(0, 16)}...${hash.substring(hash.length - 8)}`;
}

function truncate(text, maxLength = 100) {
  if (!text) return 'N/A';
  return text.length > maxLength ? `${text.substring(0, maxLength)}...` : text;
}

// Export for use in HTML
window.NLGDashboard = {
  updateDashboard,
  getStatus,
  getBlock,
  getValidatorSet,
  getAccount,
  getBalance,
  getSocialPosts,
  formatAddress,
  formatHash,
  truncate
};
