package create

var examples = `
# Creates a repository called 'my-repo' with the default repository properties
graphdbcli repository create --repository my-repo
# Creates a repository called 'my-repo', setting the Entity ID size to 40
graphdbcli repository create --repository my-repo --eidSize40
`

var shortDescription = `Creates a GraphDB repository`

var longDescription = `Creates a GraphDB repository

Compatible with: 11.1.2, 11.1.3
`
