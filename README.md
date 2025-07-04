# GraphDB CLI

GraphDB CLI is an open-source command-line tool for managing GraphDB instances, licenses, and cloud backups.

## Features
- Manage GraphDB instances
- Manage licenses
- Create S3 cloud backups

## Installation

### Homebrew (recommended)

```shell
brew install graphdbcli
```

### Manual build

```shell
make build
sudo make install
```

### Linux binary installation

1. Install the needed CLI tool version.

2. Rename the tool as needed
    ```shell
    mv ./graphdbcli-linux-amd64 ./graphdbcli
    ```
3. Install the tool for global access
    ```shell
    mv ./graphdbcli /usr/local/bin/graphdbcli
    ```

## Usage

Run `graphdbcli --help` to see available commands and options.

## Changelog

All notable changes to this project are documented in [CHANGELOG.md](./CHANGELOG.md).
Please refer to it for release notes and version history.

## Contributing

Contributions are welcome! Please read the [CONTRIBUTING.md](./CONTRIBUTING.md) guide for:
- How to propose changes or report issues
- Coding standards and commit message conventions
- How to set up your development environment


## Contacts
[Vladislav Nikolov (Graphwise)](mailto:vladislav.nikolov@graphwise.ai)
