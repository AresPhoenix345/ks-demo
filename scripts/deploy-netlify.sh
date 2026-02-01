#!/usr/bin/env bash
# deploy-netlify.sh — Build docs-site and optionally deploy via Netlify CLI
set -euo pipefail

REPO_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$REPO_ROOT/docs-site"

echo "Building docs-site..."
npm ci --legacy-peer-deps || npm install --legacy-peer-deps
npm run build

echo "Build complete. To deploy with Netlify CLI:"
echo "  npx netlify deploy --dir=.next  # or netlify deploy --prod"
echo "Or connect the repo to Netlify and push to trigger automatic deploys."
