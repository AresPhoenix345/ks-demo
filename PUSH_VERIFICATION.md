# Push verification checklist (kss-demo)

Use this checklist to verify that all required content is present and correctly pushed to the remote for the **kss-demo** repository (renamed from ks-demo).

## Before pushing

- [ ] **Repo name**: Remote repository is named `kss-demo` (or update all `kubestellar/kss-demo` references in docs to your actual org/repo).
- [ ] **Go modules**: Run `go mod tidy` in `integrations/argocd` and `integrations/terraform` so `go.sum` is generated and committed (needed for reproducible builds).
- [ ] **Docs site**: From `docs-site`, run `npm install` and `npm run build` to ensure the site builds.
- [ ] **No secrets**: No `.env`, API keys, or kubeconfigs are committed.

## Required structure (all present and pushed)

### Root

- [ ] `README.md` — Project overview, quick start (clone `kss-demo`, `cd kss-demo`).
- [ ] `SETUP.md` — Dev setup (clone `kss-demo`).
- [ ] `DEPLOYMENT.md` — Netlify deployment.
- [ ] `VALIDATION_CHECKLIST.md` — Release validation.
- [ ] `ANALYSIS.md` — Repository analysis.
- [ ] `IMPLEMENTATION_PLAN.md` — Implementation plan.
- [ ] `PUSH_VERIFICATION.md` — This file.
- [ ] `netlify.toml` — Build base `docs-site`, command, publish `.next`.
- [ ] `.gitignore` — node_modules, .next, bin, .env, etc.
- [ ] `scripts/setup-dev-env.sh`, `install-integrations.sh`, `deploy-netlify.sh`.

### Docs site (`docs-site/`)

- [ ] `package.json`, `package-lock.json`, `next.config.js`, `theme.config.tsx`, `tailwind.config.js`, `postcss.config.js`, `tsconfig.json`.
- [ ] `netlify.toml` (optional; root `netlify.toml` can be used).
- [ ] `pages/index.mdx`, `getting-started/index.mdx`, `integrations/argocd/*.mdx`, `integrations/terraform/*.mdx`, `api-reference/index.mdx`, `community/index.mdx`.
- [ ] All clone/URL references use **kss-demo** (not ks-demo).
- [ ] Theme `project.link` and `docsRepositoryBase` point to `kubestellar/kss-demo` (or your org/repo).

### Integrations

- [ ] **ArgoCD** (`integrations/argocd/`): `api/v1alpha1/*.go`, `cmd/controller/main.go`, `pkg/reconciler/*.go`, `config/samples/*.yaml`, `config/crd/bases/README.md`, `config/rbac/README.md`, `deploy/kustomize/README.md`, `go.mod` (and `go.sum` after `go mod tidy`), `Makefile`.
- [ ] **Terraform** (`integrations/terraform/`): `main.go`, `internal/provider/provider.go`, `go.mod` (and `go.sum` after `go mod tidy`).
- [ ] Go module paths use `github.com/kubestellar/kss-demo/...` (or your actual module path).

### Examples and CI

- [ ] `examples/argocd-basic/` — README, `bindingpolicy.yaml`, `argocd-binding.yaml` (repo URL: kss-demo).
- [ ] `examples/terraform-simple/` — README, `main.tf`.
- [ ] `.github/workflows/docs-build.yaml` — Builds `docs-site` on push/PR.
- [ ] `.github/workflows/integration-ci.yaml` — Builds/tests ArgoCD integration.

## After pushing

- [ ] **Remote**: `git remote -v` shows correct remote URL (e.g. `.../kss-demo.git`).
- [ ] **Branch**: Default branch (e.g. `main`) is pushed: `git push -u origin main`.
- [ ] **GitHub**: Repo on GitHub/GitLab shows all folders and files; no large or sensitive files.
- [ ] **Netlify**: If connected, build runs from base `docs-site` and succeeds.
- [ ] **Clone test**: From a clean directory, `git clone <your-kss-demo-url> && cd kss-demo && cd docs-site && npm install && npm run build` succeeds.

## If your GitHub repo name is different

If the remote repository is not named `kss-demo` (e.g. it is still `ks-demo` or something else):

1. Update `docs-site/theme.config.tsx`: `project.link` and `docsRepositoryBase` to your repo URL.
2. Update all `git clone` and repo links in README, SETUP, DEPLOYMENT, and `docs-site/pages/*.mdx` to your actual repo URL.
3. Go module paths in `integrations/argocd/go.mod` and `integrations/terraform/go.mod` (and imports) should match the repo path (e.g. `github.com/yourorg/your-repo-name`).
