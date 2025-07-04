package stop

var examples = `
# Stops an instance (Full syntax)
graphdbcli instance stop my-instance
# Stops an instance (Shortened)
graphdbcli i s my-instance
# Stops an instance with force (ungraceful)
graphdbcli i s my-instance --force
`
var shortDescription = `Shutdowns GraphDB instances`

var longDescription = `Shutdown graphDb instances gracefully,
whilst providin the option for doing it ungracefully, also known as "with force".

The shutdown operation may not finish right away, so subsequent invocations
of the show sub-command of the same super-command may be done to check the status.

The command also supports Regular Expressions, so multiple instances could be shutted down at the
same time.
`
