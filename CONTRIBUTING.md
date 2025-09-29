# Contributing to graphdbcli

Thank you for considering contributing to this project!

## Conventional Commits

All commit messages **must** follow the [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/) specification. This helps keep a clean, readable history and enables better automation.

**Example commit message:**

```
feat(parser): add new parsing logic
fix: correct typo in error message
docs: update README with usage example
```

## Commit Message Hook

To ensure compliance, a Git commit hook is provided. This hook will block any commit that does not follow the Conventional Commits format.

### Setting up the commit hook

1. Copy the provided `commit-msg` hook to your local `.git/hooks` directory:

   ```shell
   cp .githooks/commit-msg .git/hooks/commit-msg
   ```

2. Make sure the hook is executable:

   ```shell
   chmod +x .git/hooks/commit-msg
   ```

3. Now, any commit that does not follow the Conventional Commits format will be rejected.

## Pull Requests

- Ensure your branch is up to date with `main`.
- Follow the Conventional Commits format for all commits.
- Describe your changes clearly in the PR description.

Thank you for helping improve this project!
