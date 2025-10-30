package backupcmd

var examples = `Check the examples from the sub-commands below`

var shortDescription = `Manages GraphDB backup related operations`
var longDescription = `Manages GraphDB backup related operations by using the GraphDB Command API itself.

The Authentication for GraphDB could happen through two ways - passing username (--username / -u) and password (--password / -p), 
and by passing authentication token (--authToken / -t).

In the cases where the authentication data should not be left in the terminal, environment variables could be used. The
GraphDB username could be passed with the "GRAPHDB_USERNAME" environment variable, the password - "GRAPHDB_PASSWORD", and
the authentication token ("GRAPHDB_AUTH_TOKEN").

When passing the repositories, separate them with a comma, so a comma-separated list is formed. Example - "life-data,fibo,repo3".
`
