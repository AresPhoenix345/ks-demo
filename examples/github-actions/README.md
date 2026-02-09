# GitHub Actions Integration Examples

## Files

- `workflow-example.yaml` — Example workflow_dispatch for Console trigger

## Console Card Config

```json
{
  "card_type": "github_workflows",
  "config": {
    "owner": "kubestellar",
    "repo": "console"
  }
}
```

## Webhook Secret

Generate for GitHub webhook validation:

```bash
openssl rand -hex 32
```
