
# Contributing to graphdbcli

Thank you for considering contributing to this project!

## How to Propose Changes or Report Issues

- For bug reports, feature requests, or questions, please [open an issue](https://github.com/cloudvlad/graphdbcli/issues) on GitHub.
- For code changes, fork the repository, create a feature branch, and submit a pull request (PR) with a clear description of your changes.
- Please ensure your branch is up to date with `main` before submitting a PR.

## Coding Standards and Commit Message Conventions

All code should be clean, well-documented, and follow Go good practices. Use descriptive variable and function names, and add comments where helpful.

All commit messages **must** follow the [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/) specification. This helps keep a clean, readable history and enables better automation.

## Conventional Commits

All commit messages **must** follow the [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/) specification. This helps keep a clean, readable history and enables better automation.

**Example commit message:**

```
feat(backup): add cloud backup capability
fix: correct typo in error message
docs: update README with usage example
ci: fix pipeline condition
```

## Commit Message Hook

To ensure compliance, a Git commit hook is provided. This hook will stop any commit that does not follow the Conventional Commits format.

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

## Setting Up Your Development Environment

1. Clone the repository:
    ```shell
    git clone https://github.com/cloudvlad/graphdbcli.git
    cd graphdbcli
    ```
2. Install Go (version 1.25.0 or later recommended).
3. Install dependencies:
    ```shell
    make build
    ```
4. Run tests:
    ```shell
    make test
    ```
5. Set up pre-commit hooks as described above.

## Pull Requests

- Ensure your branch is up to date with `main`.
- Follow the Conventional Commits format for all commits.
- Describe your changes clearly in the PR description.


### Thank you for helping improve this project! :tada: 


## Useful links
[How to write CLI Help](https://bettercli.org/design/cli-help-page/#how-to-write-cli-help)
