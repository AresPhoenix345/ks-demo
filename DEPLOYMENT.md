# Netlify Deployment Guide

## Automatic Deployment

This project is configured for automatic deployment of the **docs-site** to Netlify.

### Prerequisites

1. **Netlify account** — [Sign up](https://netlify.com)
2. **GitHub (or Git) access** — Connect repository to Netlify

## Configuration

### 1. Create New Site on Netlify

1. Log in to Netlify.
2. **Add new site** → **Import an existing project**.
3. Connect to GitHub and select the repository.
4. Configure build settings:
   - **Base directory**: `docs-site`
   - **Build command**: `npm run build`
   - **Publish directory**: `docs-site/.next` (or use Netlify Next.js plugin)
   - **Node version**: 18 or 20

### 2. Environment Variables

In Netlify: **Site settings** → **Environment variables**:

- `NODE_VERSION` = `18` (or `20`)
- `NPM_FLAGS` = `--legacy-peer-deps` (if needed for dependency resolution)

### 3. Build Configuration

The `docs-site/netlify.toml` file contains build configuration. Key settings:

- `[build]`: base, command, publish
- `[build.environment]`: NODE_VERSION, NPM_FLAGS
- `[[plugins]]`: `@netlify/plugin-nextjs` for Next.js

## Manual Deployment

### Using Netlify CLI

```bash
npm install -g netlify-cli
netlify login
cd docs-site
netlify link   # link to existing site
netlify deploy        # preview
netlify deploy --prod # production
```

## Deploy Previews

- **Pull requests**: Netlify can create deploy previews per PR.
- **URL pattern**: `https://deploy-preview-<PR_NUMBER>--<site-name>.netlify.app`

## Custom Domain (Optional)

1. **Site settings** → **Domain management** → **Add custom domain**.
2. Configure DNS (CNAME or A record as instructed by Netlify).
3. HTTPS is automatic with Netlify.

## Troubleshooting

### Build Failures

- **Node version**: Set `NODE_VERSION=18` (or 20) in environment variables.
- **Dependencies**: Set `NPM_FLAGS=--legacy-peer-deps` if needed.
- **Out of memory**: Reduce concurrency or upgrade Netlify plan.

### Preview Deploy Issues

- Ensure Netlify app has access to the repository.
- Check `netlify.toml` and base directory.

## Rollback

In Netlify **Deploys** tab, find a previous successful deploy and click **Publish deploy** to roll back.
