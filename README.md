GraphDB CLI
===

This project is not affiliated with or endorsed by Ontotext/Graphwise. It is an independent side project using only publicly available resources.

## Functionalities
- Pull platform independent distribution files
- Manage licenses locally

## Install using Homebrew

```shell
brew install graphdbcli
```

## Manual build from Source code

```shell
go build .
./graphdbcli -v
# Copy in the appropriate directory
```

## Install locally using Homebrew Formulae

```shell
git archive --format=tar.gz HEAD -o /tmp/graphdb-cli.tar.gz
sed -i 's|url ".*"|url "file:///tmp/graphdb-cli.tar.gz"|' graphdb-cli.rb
brew install --build-from-source graphdb-cli.rb
```

## Contacts
[vladislav.nikolov@graphwise.ai](mailto://vladislav.nikolov@graphwise.ai)