package beamcmd

var shortDescription = `Translates shortened GraphDB API payloads to their expected format.`

var longDescription = `Translates shortened GraphDB API payloads to their expected format, by
adding all fields, if not mentioned, and the expected attributes for them. Also cleans up the
structure by flattening it to a two levels.`

var examples = `# Start the daemon
graphdbcli beam
`
