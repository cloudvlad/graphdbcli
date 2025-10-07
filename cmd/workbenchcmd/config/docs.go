package config

var examples = `# Configure GraphDB port
graphdbcli workbench config my-custom-wb --graphdb-port 7770
# Configure the GraphDB host address
graphdbcli workbench config my-custom-wb --graphdb-host localhost
# Configure the Workbench port
graphdbcli workbench config my-custom-wb --workbench-port 9009
# Full command for creating non-default workbench instance
graphdbcli workbench config my-custom-wb \
--workbench-port 9009 \
--graphdb-host localhost \
--graphdb-port 9009
# Reset to the default values
graphdbcli workbench config my-custom-wb
`

var shortDescription = `Add GraphDB Workbench instances`

var longDescription = `Add GraphDB Workbench instances

Copies the provided directory inside the dedicated tool directory for further management and usage.
`
