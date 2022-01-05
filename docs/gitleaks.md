## Setup Gitleaks

The [`Brewfile`](../Brewfile) ensures the [`gitleaks`](https://github.com/zricethezav/gitleaks) is installed.

Next, configure Git hooks:
```bash
git config core.hooksPath ./internal/git-hooks
```
