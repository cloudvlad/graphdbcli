package upload

var examples = `# Specifying the baseUri with the corner brackets is required
graphdbcli rdf4j sparql upload ./eg-data/f100e1000t100.zip --repository my-data-repo --baseUri '<http://test.org/eg1>' 
graphdbcli rdf4j sparql upload ./eg-data/f100e1000t100.zip --repository my-data-repo
graphdbcli rdf4j sparql upload ~/data.ttl --repository my-data-repo
graphdbcli rdf4j sparql upload /home/vlad/data.nt --repository my-data-repo --rdfFormat 'application/n-triples'
graphdbcli rdf4j sparql upload ~/data.ttl --repository my-data-repo --graphdb-address http://localhost:7200
`

var shortDescription = `Uploads data to a repository`

var longDescription = `Uploads data to a repository, 
by uploading a RDF document, that is added alongside the existing data.

API Documentation: https://rdf4j.org/documentation/reference/rest-api/#tag/SPARQL/paths/~1repositories~1%7BrepositoryID%7D~1statements/post

Implementation details:
- The command does not support updates. This will be added as part of another command.

Compatible with: 
- GraphDB 11.1.0 (RDF4J 5.1.4)
`
