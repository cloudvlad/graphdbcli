package create

var examples = `
# Creates a repository called 'my-repo' with the default repository properties
graphdbcli repository create my-repo
# Creates a repository called 'another-repo' using locally placed file
graphdbcli repository create another-repo --file my-repo.ttl
# Creates a repository called 'my-repo' specifying entity id size
graphdbcli repository create my-repo entity-id-size 40
`

var shortDescription = `Creates a GraphDb repository`

var longDescription = `Creates a GraphDB repository

API Documentation: https://rdf4j.org/documentation/reference/rest-api/#tag/Graph-Store/paths/~1repositories~1%7BrepositoryID%7D~1rdf-graphs~1%7Bgraph%7D/post

Compatible with: 
- GraphDB 11.1.0 (RDF4J 5.1.4)
`
