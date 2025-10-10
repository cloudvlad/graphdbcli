package helm

var examples = `graphdbcli version install 11.0.0
# Shortcut alternative for installing a version
graphdbcli v i 11.0.0
graphdbcli version install 11.0.0 --integrity-check
# Shortcut version for doing an integrity check
graphdbcli version install 11.0.0 -c
graphdbcli version install
`

var shortDescription = `Prepares the release for the next version of the GraphDB Helm chart.`

var longDescription = `Prepares the release for the next version of the GraphDB Helm chart.`
