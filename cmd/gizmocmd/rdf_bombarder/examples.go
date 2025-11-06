package rdf_bombarder

var examples = `# Sending 10 statements to each of 10 named graphs (100 statements total), 5 at a time, using 10 threads
graphdbcli gizmo rdf_bombarder -r my_repo \
--numberOfNamedGraphs 10 \
--numberOfStatementsPerNamedGraph 10 \
--numberOfThreads 10 \
--numberOfStatementsPerRequest 5
# Using the default value for numberOfNamedGraphs (0) will do nothing 
# (reserved for future development)
graphdbcli gizmo rdf_bombarder -r my_repo --numberOfNamedGraphs 0
`
