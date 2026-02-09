#!/usr/bin/env bash
# install-console-integrations.sh — Merge kss-demo integrations into a Console fork
set -euo pipefail

REPO_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$REPO_ROOT"

# Path to your Console fork (default: sibling directory)
CONSOLE_DIR="${CONSOLE_DIR:-$(dirname "$REPO_ROOT")/console}"

if [[ ! -d "$CONSOLE_DIR" ]]; then
  echo "Console directory not found: $CONSOLE_DIR"
  echo ""
  echo "Clone Console first:"
  echo "  git clone https://github.com/kubestellar/console.git $CONSOLE_DIR"
  echo ""
  echo "Or set CONSOLE_DIR:"
  echo "  CONSOLE_DIR=/path/to/console ./scripts/install-console-integrations.sh"
  exit 1
fi

echo "Installing Console integrations from $REPO_ROOT into $CONSOLE_DIR"

# Copy integration packages
if [[ -d "$REPO_ROOT/console-integrations/pkg/integrations" ]]; then
  echo "Copying pkg/integrations..."
  mkdir -p "$CONSOLE_DIR/pkg/integrations"
  cp -r "$REPO_ROOT/console-integrations/pkg/integrations/"* "$CONSOLE_DIR/pkg/integrations/"
fi

# Copy React cards
if [[ -d "$REPO_ROOT/web-cards" ]]; then
  echo "Copying web cards..."
  mkdir -p "$CONSOLE_DIR/web/src/components/cards"
  for f in "$REPO_ROOT/web-cards"/*.tsx; do
    [[ -f "$f" ]] && cp "$f" "$CONSOLE_DIR/web/src/components/cards/"
  done
fi

echo ""
echo "Done. Next steps:"
echo "  1. Add API handlers in Console for /api/integrations/prometheus/*, /api/integrations/github/*"
echo "  2. Register new card types in pkg/models/card.go and frontend card registry"
echo "  3. Add env vars: PROMETHEUS_URL, GRAFANA_URL, GRAFANA_API_KEY, GITHUB_TOKEN"
echo "  4. Rebuild: cd $CONSOLE_DIR && ./scripts/prod.sh"
