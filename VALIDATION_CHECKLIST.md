# Validation Checklist

Use this checklist to validate the ks-demo project before release or submission.

## Code Quality

- [ ] All Go code passes `go vet`
- [ ] All tests pass (`make test` in each integration)
- [ ] Code coverage >80% (where applicable)
- [ ] No critical security vulnerabilities
- [ ] Consistent code style (e.g. golangci-lint)

## Documentation Site

- [ ] Site builds without errors (`npm run build` in docs-site)
- [ ] All pages render correctly
- [ ] Navigation works
- [ ] Code examples are accurate
- [ ] Images/diagrams display correctly
- [ ] Mobile responsive
- [ ] Acceptable load time (<3s target)

## Integration Functionality

### ArgoCD Integration

- [ ] CRDs install successfully
- [ ] Controller starts without errors
- [ ] Reconciliation logic works (ArgoCDBinding → Applications)
- [ ] Status updates correctly
- [ ] Error handling and cleanup work

### Terraform Integration

- [ ] Provider builds and loads
- [ ] Resource CRUD works (create/read/update/delete BindingPolicy)
- [ ] Data source works (if implemented)
- [ ] Examples run successfully

## Examples

- [ ] All examples run successfully
- [ ] Example documentation is clear
- [ ] Examples demonstrate key features

## Deployment

- [ ] Netlify build succeeds
- [ ] Production site is accessible
- [ ] All routes work in production
- [ ] Environment variables configured as needed

## User Experience

- [ ] Installation process is clear and works
- [ ] Documentation is helpful and accurate
- [ ] Examples are easy to follow
- [ ] Troubleshooting guides are useful

## LFX Deliverables (Project-Specific)

- [ ] 2 production-ready integrations
- [ ] Integration documentation with setup guides
- [ ] 2 demo videos (overview + setup)
- [ ] Sample implementations and templates
- [ ] Integration maintenance guide
- [ ] User feedback mechanism (e.g. GitHub issues template)
