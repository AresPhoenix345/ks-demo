# Implementation Plan for KubeStellar Integrations

**LFX Mentorship Project**: CNCF - KubeStellar: Integration and ecosystem development specialist (2026 Term 1)

This document outlines the implementation plan for two production-ready integrations and the standalone kss-demo project deployable on Netlify.

---

## Integration 1: ArgoCD

### Rationale

- Most widely used GitOps tool; clear user demand.
- Natural fit: BindingPolicy defines “what” and “where”; Argo CD Applications define “what” and “where” in GitOps terms.
- Medium complexity: controller watches ArgoCDBinding CR; reads BindingPolicy/Binding from KubeStellar; creates/updates Argo CD Application resources.

### Architecture

```
┌─────────────────────────────────────────────────────────────────┐
│  KubeStellar WDS (Workload Definition Space)                     │
│  - BindingPolicy (clusterSelectors + downsync)                  │
│  - Binding (resolved workload refs + destinations)               │
└───────────────────────────────┬─────────────────────────────────┘
                                │
                                ▼
┌─────────────────────────────────────────────────────────────────┐
│  ArgoCD Integration Controller (kss-demo)                          │
│  - Watches ArgoCDBinding CR                                      │
│  - Reads BindingPolicy / Binding from KubeStellar API             │
│  - For each destination: create/update Argo CD Application       │
│  - Optionally: sync status from Argo CD back to ArgoCDBinding    │
└───────────────────────────────┬─────────────────────────────────┘
                                │
                                ▼
┌─────────────────────────────────────────────────────────────────┐
│  Argo CD                                                         │
│  - Application CRs (one per cluster/destination)                  │
│  - Syncs from Git to target clusters                             │
└─────────────────────────────────────────────────────────────────┘
```

### Components to Build

1. **Custom Resource**: `ArgoCDBinding` (integrations.kubestellar.io/v1alpha1)
   - `bindingPolicyRef`: reference to KubeStellar BindingPolicy (name or name+namespace if namespaced).
   - `applicationTemplate`: template for Argo CD Application (project, source, destination, syncPolicy).
   - Status: conditions, observedGeneration, lastSyncStatus.

2. **Controller/Reconciler**
   - Reconcile ArgoCDBinding: fetch BindingPolicy → resolve Binding (or use BindingPolicy + cluster list) → for each destination create/update Application with correct `destination.server`/`destination.name`.
   - Use KubeStellar clientset for BindingPolicy/Binding; Argo CD clientset or dynamic client for Application.

3. **CRD + RBAC**
   - CRD YAML for ArgoCDBinding.
   - ServiceAccount, Role, RoleBinding for controller.

4. **Deployment**
   - Kustomize overlay and/or Helm chart for controller + CRD.

5. **Documentation**
   - Installation, user guide, architecture, troubleshooting, examples (see docs-site structure below).

6. **Examples**
   - Basic: one BindingPolicy → one ArgoCDBinding → N Applications.
   - Multi-cluster: multiple destinations, single Git repo/path.

### File Structure (integrations/argocd/)

```
integrations/argocd/
├── api/v1alpha1/
│   ├── argocd_binding_types.go
│   └── zz_generated.deepcopy.go
├── cmd/controller/main.go
├── pkg/
│   ├── argocd/client.go
│   ├── kubestellar/client.go
│   └── reconciler/argocd_binding_reconciler.go
├── config/crd/bases/
├── config/rbac/
├── config/samples/
├── deploy/kustomize/
├── deploy/helm/
├── go.mod
├── go.sum
└── Makefile
```

### Dependencies

- `github.com/kubestellar/kubestellar` (API types, optional: client if same repo or vendored).
- `sigs.k8s.io/controller-runtime`.
- Argo CD Application API (argoproj.io/v1alpha1) – client or unstructured.

### Timeline (High Level)

- **Week 1**: CRD types, codegen, controller skeleton, CRD + RBAC manifests.
- **Week 2**: Full reconciliation (BindingPolicy → Binding → Applications), unit tests, samples, basic docs.
- **Week 3**: Helm/Kustomize, integration tests, documentation polish, demo video outline.

---

## Integration 2: Terraform

### Rationale

- Popular IaC tool; complements GitOps.
- Enables BindingPolicy (and optionally cluster discovery) to be managed as Terraform resources/data sources.
- Fits “infrastructure as code” and platform-team workflows.

### Architecture

```
┌─────────────────────────────────────────────────────────────────┐
│  Terraform                                                       │
│  - kubestellar_binding_policy resource                           │
│  - kubestellar_clusters data source (optional)                   │
└───────────────────────────────┬─────────────────────────────────┘
                                │
                                ▼
┌─────────────────────────────────────────────────────────────────┐
│  Terraform Provider (kss-demo)                                     │
│  - CRUD for BindingPolicy via KubeStellar API                     │
│  - Read cluster list from OCM or KubeStellar context              │
└───────────────────────────────┬─────────────────────────────────┘
                                │
                                ▼
┌─────────────────────────────────────────────────────────────────┐
│  KubeStellar API (WDS)                                           │
│  - BindingPolicy CRUD                                             │
│  - (Optional) cluster inventory via OCM                          │
└─────────────────────────────────────────────────────────────────┘
```

### Components to Build

1. **Terraform Provider**
   - Provider schema: kubeconfig path or in-cluster config for KubeStellar WDS.
   - Resource: `kubestellar_binding_policy` – map to BindingPolicy (clusterSelectors, downsync clauses).
   - Data source: `kubestellar_clusters` (optional) – list clusters (e.g. from OCM ManagedCluster).

2. **Implementation**
   - Use Terraform Plugin Framework (recommended) or SDK v2.
   - KubeStellar clientset (or generated client for control.v1alpha1) for CRUD.

3. **Documentation**
   - Provider docs (resources/data sources), examples, release to Terraform Registry (optional).

4. **Examples**
   - `examples/terraform/simple/` – one BindingPolicy.
   - `examples/terraform/multi-cluster/` – multiple policies, optional data source.

### File Structure (integrations/terraform/)

```
integrations/terraform/
├── main.go
├── internal/provider/
│   ├── provider.go
│   ├── resource_binding_policy.go
│   └── datasource_clusters.go
├── examples/
├── go.mod
├── go.sum
└── Makefile
```

### Dependencies

- Terraform Plugin Framework (or SDK v2).
- `github.com/kubestellar/kubestellar` (API types + client for BindingPolicy).

### Timeline (High Level)

- **Week 1**: Provider skeleton, resource_binding_policy (create/read/update/delete).
- **Week 2**: Data source (if applicable), docs, examples, tests.
- **Week 3**: Polish, optional Terraform Registry submission, demo content.

---

## Testing Strategy

- **Unit**: Controller logic with fake clients; provider CRUD with mock API.
- **Integration**: Run controller against real KubeStellar + Argo CD (e.g. Kind); run provider against real cluster.
- **E2E**: Full flow – apply ArgoCDBinding → verify Applications; apply Terraform → verify BindingPolicy.

---

## Documentation Strategy

- **Single docs site** in `kss-demo/docs-site/` (Nextra), deployable on Netlify.
- **Sections**: Getting Started, Integrations (ArgoCD, Terraform), API Reference (high-level), Examples, Community.
- **Per integration**: Installation, User Guide, Architecture, Troubleshooting, Examples (links to `examples/` in repo).

---

## Deployment Strategy (Netlify)

- **Build context**: `docs-site/` (base directory for Netlify).
- **Build command**: `npm run build`.
- **Publish**: `.next` (with `@netlify/plugin-nextjs` if using Next.js App Router) or `out` if static export.
- **Environment**: NODE_VERSION=18 or 20; NPM_FLAGS=--legacy-peer-deps if needed.
- **Redirects**: Optional short URLs (e.g. /argocd → /integrations/argocd).

---

## kss-demo Repository Layout (Summary)

```
kss-demo/
├── README.md
├── SETUP.md
├── DEPLOYMENT.md
├── ANALYSIS.md          (copy or link from parent)
├── IMPLEMENTATION_PLAN.md
├── integrations/
│   ├── argocd/          (as above)
│   └── terraform/       (as above)
├── docs-site/           (Nextra + Next.js, Netlify)
├── examples/            (argocd-basic, terraform-simple, etc.)
├── scripts/             (setup-dev-env.sh, install-integrations.sh, deploy-netlify.sh)
├── .github/workflows/    (CI for integrations, docs build, optional Netlify deploy)
└── package.json         (optional root workspace)
```

This plan aligns with the LFX project outcomes: two production-ready integrations, clear setup guides, sample implementations, and a single deployable documentation site for Netlify.
