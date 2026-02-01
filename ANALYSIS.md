# KubeStellar Repository Analysis

**LFX Mentorship Project**: CNCF - KubeStellar: Integration and ecosystem development specialist (2026 Term 1)

This document summarizes the analysis of the two source repositories used to build the ks-demo prototype.

---

## kubestellar/kubestellar Analysis

### Tech Stack

| Component | Version |
|-----------|---------|
| **Go** | 1.24.1 |
| **Kubernetes (client-go, apimachinery, etc.)** | 0.30.14 |
| **Controller Runtime** | 0.18.7 |
| **KubeStellar/KubeFlex** | v0.9.1 |
| **Open Cluster Management (OCM)** | v0.15.0 |
| **Prometheus client** | v1.18.0 |
| **CEL (Common Expression Language)** | v0.17.8 |

### Key Dependencies

- `sigs.k8s.io/controller-runtime` – controller framework
- `k8s.io/client-go`, `k8s.io/apimachinery` – Kubernetes APIs
- `open-cluster-management.io/api` – cluster registration and inventory
- `github.com/kubestellar/kubeflex` – KubeFlex (control plane hosting)
- `github.com/google/cel-go` – status/expression evaluation
- `github.com/prometheus/client_golang` – metrics

### Architecture Overview

- **Controller Manager** (`cmd/controller-manager/main.go`): Runs binding and status controllers; connects to ITS (Inventory and Transport Space) and WDS (Workload Definition Space).
- **Binding subsystem** (`pkg/binding/`): Resolves BindingPolicies into Bindings; watches workload objects and cluster inventory; uses informers, workqueues, and reconciliation.
- **Status subsystem** (`pkg/status/`): Aggregates status from WECs (Workload Execution Clusters) into CombinedStatus; uses StatusCollectors and CEL.
- **Transport** (`pkg/transport/`): OCM-based transport controller; syncs workload objects to downstream clusters.
- **API** (`api/control/v1alpha1/`): BindingPolicy, Binding, StatusCollector, CombinedStatus, CustomTransform.

### API Structure

| Custom Resource | API Group | Scope | Purpose |
|-----------------|-----------|--------|---------|
| **BindingPolicy** | control.kubestellar.io/v1alpha1 | Cluster | Binds “what” (workloads) to “where” (clusters) via clusterSelectors and downsync clauses |
| **Binding** | control.kubestellar.io/v1alpha1 | Cluster | Resolved binding: explicit workload refs + destinations |
| **StatusCollector** | control.kubestellar.io/v1alpha1 | Cluster | Defines how to aggregate status from WECs (CEL expressions) |
| **CombinedStatus** | control.kubestellar.io/v1alpha1 | Namespaced | Aggregated status per (workload, BindingPolicy) |
| **CustomTransform** | control.kubestellar.io/v1alpha1 | Cluster | JSONPath-based transforms on objects during propagation |

**BindingPolicy** highlights:

- `clusterSelectors`: label selectors for clusters (OCM ManagedClusters).
- `downsync`: list of `DownsyncPolicyClause` (object selection + modulation: createOnly, statusCollectors, WantSingletonReportedState, WantMultiWECReportedState).
- Object selection: APIGroup, resources, namespaces, namespaceSelectors, objectSelectors, objectNames.

### Integration Points

1. **BindingPolicy as integration anchor** – External tools (e.g. ArgoCD, Terraform) can drive or consume BindingPolicies to define multi-cluster placement.
2. **Binding status** – Controllers can watch Bindings/CombinedStatus for sync/health to feed into GitOps or dashboards.
3. **WDS as source of truth** – Workloads in WDS are what get propagated; integrations can create/update WDS objects and BindingPolicies.
4. **OCM cluster inventory** – Cluster identity and labels come from OCM; integrations can assume ManagedCluster CRs and labels.
5. **No dedicated “integrations” or “plugins” directory** – Integration code would live in a separate repo (e.g. ks-demo) and depend on this repo’s API and client types.

### Code Patterns

- **Controller setup**: Informers for BindingPolicy, Binding, cluster listers, dynamic informers for workload types; workqueue for reconciliation; rate limiting.
- **Reconciliation**: Resolve policy → compute Binding spec (workload refs + destinations) → update Binding; status controller updates conditions/errors.
- **Client usage**: Generated clientset `ksclient`, typed control client, dynamic client for arbitrary GVKs, OCM cluster client.
- **Logging**: `klog/v2`; context-based logger.
- **Metrics**: Prometheus metrics via controller-runtime and custom metrics in `pkg/metrics`.

### Build and Deploy

- **Makefile**: `make build`, `make test`, `make deploy` (with KIND_HOSTING_CLUSTER, DEFAULT_WDS_NAME); image targets for controller-manager and transport.
- **Config**: `config/` – RBAC, CRDs, sample manifests; `core-chart/` – Helm for core KubeStellar.
- **Scripts**: `scripts/create-kubestellar-demo-env.sh` – demo env (e.g. Kind); `hack/` – codegen and tooling.

---

## kubestellar/docs (ks-docs) Analysis

### Tech Stack

| Component | Version |
|-----------|---------|
| **Next.js** | ^15.5.9 |
| **Nextra** | ^4.6.1 |
| **React** | ^19.2.0 |
| **Tailwind CSS** | ^4 |
| **TypeScript** | ^5 |
| **next-intl** | ^4.3.12 |
| **Node (Netlify)** | 20.11.1 |

Additional: `framer-motion`, `lucide-react`, `mermaid`, `@theguild/remark-mermaid`, `three`, `@react-three/fiber`, `@react-three/drei`.

### Documentation Structure

- **App Router**: `src/app/` – `[locale]/` (marketing, quick-installation, programs, etc.), `docs/` (catch-all `[...slug]`), `api/` (docs-image proxy, search).
- **Docs content**: Driven by `docs/content/` and `src/app/docs/page-map.ts`; multi-project (a2a, kubeflex, multi-plugin, kubestellar-mcp, console) with `getContentPath` / `getBasePath`.
- **i18n**: `next-intl`, `messages/*.json`, `src/i18n/`.
- **Components**: `src/components/` – docs (DocsLayout, DocsSidebar, TableOfContents, ThemeToggle, VersionSelector), master-page sections, animations (globe, GridLines, StarField), etc.

### Build Configuration

- **Build command**: `npm run build`
- **Output**: `.next` (Netlify uses `@netlify/plugin-nextjs`)
- **Publish**: `.next` (Netlify)
- **Config**: `next.config.ts` – Nextra (latex, search), next-intl plugin; rewrites for `/docs-images/` → `/api/docs-image/`; redirects for /quickstart, /slack, etc.

### Netlify Configuration (`netlify.toml`)

- `[build]`: command `npm run build`, publish `.next`, NODE_VERSION 20.11.1.
- `[[plugins]]`: `@netlify/plugin-nextjs`.
- Context-specific commands for branch-deploy and version branches (`docs/0.28.0`, etc.).
- Many `[[redirects]]` for short URLs (e.g. /slack, /quickstart) and legacy domains (a2a, kubeflex, multi, mcp, etc.) → kubestellar.io/docs/...
- CORS headers for `/config/*`.
- Forms spam protection.

### Reusable Components and Conventions

- **Layout**: DocsLayout, DocsSidebar, DocsNavbar, DocsFooter, MobileSidebarToggle, MobileTOC.
- **Content**: EditPageLink, EditViewSourceButtons, TableOfContents, RelatedProjects, ThemeToggle, VersionSelector.
- **Styling**: Tailwind with custom theme (e.g. `space-dark`), dark mode (class), animations (e.g. fade-in-up, status-glow).
- **Docs**: MDX under `docs/content/`; page-map built from filesystem; version selector and project-specific base paths.

---

## Integration Opportunities

### From kubestellar/kubestellar

1. **ArgoCD**: Consume BindingPolicy + Binding (and optionally CombinedStatus); create/update Argo CD Application(s) per cluster or per BindingPolicy; sync policy from KubeStellar into GitOps flows.
2. **Terraform**: Provider or module that manages BindingPolicy (and optionally Binding) resources; data sources for clusters/destinations from OCM or Binding status.
3. **CI/CD (e.g. GitHub Actions)**: Scripts or actions that apply BindingPolicies and verify Binding/CombinedStatus after deployments.
4. **Observability**: Export Binding/BindingPolicy/CombinedStatus metrics or events to Prometheus/Grafana (core already has some Prometheus client usage).

### From Documentation Needs

1. **Integration-specific docs**: Installation, architecture, user guide, troubleshooting, examples (aligned with ks-docs style and Nextra).
2. **Code samples**: YAML and CLI snippets for BindingPolicy + integration tool (ArgoCD Application, Terraform resource).
3. **Demo flow**: Quick start that uses both KubeStellar and the integration (e.g. create BindingPolicy → see Argo CD Applications).

---

## Recommendations for Prototype (ks-demo)

1. **Two integrations**
   - **Primary: ArgoCD** – High demand, clear mapping (BindingPolicy → Applications per cluster/destination).
   - **Secondary: Terraform** – IaC story; Terraform resource/data source for BindingPolicy (and optionally cluster list).

2. **Standalone repo (ks-demo)**
   - Own Go modules under `integrations/argocd/` and `integrations/terraform/`; depend on `github.com/kubestellar/kubestellar` for API types and clients (or copy minimal types if needed for portability).
   - Docs site in `docs-site/` – Nextra + Next.js, simplified compared to full ks-docs (single project, no multi-version), deployable to Netlify.

3. **Reuse patterns from kubestellar**
   - Controller that watches a custom resource (e.g. ArgoCDBinding) and optionally BindingPolicy/Binding; uses controller-runtime, informers, and similar client patterns.
   - CRDs and RBAC in `config/crd`, `config/rbac`; samples in `config/samples`.

4. **Reuse patterns from ks-docs**
   - Nextra theme, Tailwind, TypeScript; minimal `theme.config` and layout; docs under `docs-site/pages/` or `docs-site/src/app/docs/` depending on chosen structure; `netlify.toml` for build and publish.

5. **Deliverables**
   - Two production-ready integrations (controller/operator + CRD + docs + examples).
   - One deployable docs site (Netlify) that documents both integrations and points to examples and repos.
