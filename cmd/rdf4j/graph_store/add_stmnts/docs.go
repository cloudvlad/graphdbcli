package add_stmnts

var examples = `graphdbcli rdf4j graph-store add-statements --repository my-data-repo --graph 'http://test.org/eg1' ./eg-data/f100e1000t100.zip
graphdbcli rdf4j graph-store add-statements --repository my-data-repo ./eg-data/f100e1000t100.zip
graphdbcli rdf4j graph-store add-statements --repository my-data-repo ~/data.ttl
graphdbcli rdf4j graph-store add-statements /home/vlad/data.nt --repository my-data-repo --rdfFormat 'application/n-triples'
graphdbcli rdf4j graph-store add-statements /home/vlad/data.ttl --repository my-data-repo --graphdb-address http://localhost:7200
`

var shortDescription = `Add statements to a directly referenced named graph.`

var longDescription = `Add statements to a directly referenced named graph.

API Documentation: https://rdf4j.org/documentation/reference/rest-api/#tag/Graph-Store/paths/~1repositories~1%7BrepositoryID%7D~1rdf-graphs~1%7Bgraph%7D/post

Compatible with: 
- GraphDB 11.1.0 (RDF4J 5.1.4)
`
