# Deployment Guide: Frontend (Netlify) + Backend (Render)

This guide covers deploying the **frontend** (docs-site) on **Netlify** first, then the **backend** on **Render**.

---

# Part 1: Deploy frontend on Netlify

Deploy the **kss-demo docs-site** (Next.js/Nextra) to Netlify.

## Prerequisites

- **kss-demo** repository pushed to **GitHub** (or GitLab/Bitbucket).
- **Netlify account** — [Sign up](https://app.netlify.com/signup).

## Step 1: Open Netlify and add a new site

1. Go to **[https://app.netlify.com](https://app.netlify.com)** and log in.
2. Click **Add new site** → **Import an existing project**.
3. Choose **GitHub** (or your Git provider).
4. Authorize Netlify if asked, then **select the repository** that contains **kss-demo** (e.g. `kss-demo` or the repo where `kss-demo` is the root).

## Step 2: Configure build settings

Netlify may read `netlify.toml` from the repo. If it does, you only need to set the **branch** (e.g. `main`). Otherwise set these manually:

| Setting            | Value                    |
|--------------------|--------------------------|
| **Branch to deploy** | `main` (or your default) |
| **Base directory**   | `docs-site`              |
| **Build command**   | `npm run build`           |
| **Publish directory** | `.next`                 |

Important:

- **Base directory** must be **`docs-site`** so Netlify builds the Next.js app (not the repo root).
- **Publish directory** is **`.next`** (relative to the base directory, i.e. `docs-site/.next`).  
  If you use **@netlify/plugin-nextjs** (see below), the plugin may override this; that’s fine.

## Step 3: Install Netlify Next.js plugin (recommended)

For Next.js 14, Netlify’s Next.js plugin is recommended:

1. In Netlify: **Site settings** → **Build & deploy** → **Build**.
2. Under **Build plugins**, click **Add plugin** → **Netlify** → **Next.js** (`@netlify/plugin-nextjs`).
3. Save.

Your repo’s **root** `netlify.toml` already includes:

```toml
[[plugins]]
  package = "@netlify/plugin-nextjs"
```

So if Netlify uses the repo’s `netlify.toml`, the plugin is already configured.

## Step 4: Set environment variables

1. **Site settings** → **Environment variables** → **Add a variable** (or **Add from .env**).
2. Add:

| Key           | Value              | Scopes   |
|---------------|--------------------|----------|
| `NODE_VERSION` | `18`               | All      |
| `NPM_FLAGS`   | `--legacy-peer-deps` | All (optional) |

Use `NPM_FLAGS` if the build fails with dependency/peer dependency errors.

## Step 5: Deploy

1. Click **Deploy site** (or **Trigger deploy** → **Deploy site**).
2. Wait for the build to finish. Logs are under **Deploys** → click the deploy.
3. When it succeeds, the site URL will be like:  
   **`https://<random-name>.netlify.app`**

## Step 6: Verify

- Open the Netlify URL. You should see the KubeStellar Integrations docs (landing page).
- Click **Getting started**, **Integrations** (ArgoCD, Terraform), etc. All routes should load.

## Optional: Custom domain

1. **Site settings** → **Domain management** → **Add custom domain**.
2. Enter your domain and follow Netlify’s DNS instructions (CNAME or A record).
3. HTTPS is provided automatically.

## Optional: Deploy from CLI

```bash
cd kss-demo
npm install -g netlify-cli
netlify login
netlify init   # choose “Create & configure a new site” or “Link to existing site”
# Set build command: npm run build (from docs-site)
# Set publish directory: docs-site/.next (or .next if base is docs-site)
netlify deploy --build
netlify deploy --build --prod
```

---

# Part 2: Deploy backend on Render

Deploy a **backend** (API or server) on Render **after** the frontend is on Netlify.

## When to use this

- You add an **API** (e.g. Node/Express, Go, Python) that the frontend calls.
- kss-demo today is **docs-site only**; the integrations (ArgoCD, Terraform) are Go code that run in Kubernetes, not a hosted API. When you introduce a backend service (e.g. REST API for integrations), use the steps below.

## Prerequisites

- **Render account** — [Sign up](https://render.com).
- Backend code in your repo (e.g. `backend/` or `api/` with its own `package.json` or `go.mod`).

## Step 1: Create a Web Service on Render

1. Go to **[https://dashboard.render.com](https://dashboard.render.com)**.
2. Click **New +** → **Web Service**.
3. Connect your **GitHub** (or Git) account and select the **same repository** as Netlify (e.g. kss-demo).
4. Configure:
   - **Name**: e.g. `kss-demo-api`.
   - **Region**: choose one.
   - **Branch**: `main` (or your default).
   - **Root directory**: folder that contains the backend (e.g. `backend` or `api`). Leave blank if the backend is at repo root.
   - **Runtime**: **Node** or **Go** (or Python, etc.) depending on your backend.
   - **Build command**: e.g. `npm install` (Node) or `go build -o server` (Go).
   - **Start command**: e.g. `npm start` or `./server` (must start the HTTP server).
5. **Instance type**: Free or paid.
6. Click **Create Web Service**.

## Step 2: Environment variables (backend)

In the Render service → **Environment** tab, add any env vars your backend needs (e.g. `DATABASE_URL`, `API_KEY`, `NODE_ENV=production`).

## Step 3: Use the backend URL in the frontend

1. Render gives you a URL like **`https://kss-demo-api.onrender.com`**.
2. In your **frontend** (docs-site or any React/Next app), call this URL for API requests (e.g. `fetch('https://kss-demo-api.onrender.com/...')`).
3. If you need this URL at build time, add it as a **Netlify env var** (e.g. `NEXT_PUBLIC_API_URL`) so the frontend can use it.

## CORS (if frontend and backend are on different origins)

- Frontend: `https://your-site.netlify.app`
- Backend: `https://kss-demo-api.onrender.com`

If the browser blocks requests, configure **CORS** on the backend to allow your Netlify origin (e.g. `https://your-site.netlify.app` or `https://*.netlify.app`).

---

# Summary

| What        | Where   | Repo path   | Build command   | Publish / Start      |
|------------|---------|-------------|------------------|----------------------|
| **Frontend** | Netlify | `docs-site` | `npm run build`  | `.next` (Next.js)    |
| **Backend**  | Render  | e.g. `backend` | `npm install` / `go build` | `npm start` / `./server` |

1. Deploy **frontend** on Netlify first (Part 1).
2. When you have backend code, deploy **backend** on Render (Part 2) and point the frontend to its URL.
