package show

var shortDescription = `Shows the local or remote repositories`

var longDescription = `Shows the local or remote repositories based on the passed flags.
If the --location flag is passed, the show command displays the repositories that are
located in the GraphDB instance itself. If not - it shows the configuration files that are stored locally.
`

var examples = `
# Shows local repositories configurations
graphdbcli repository show
# Shows the remote repositories
graphdbcli repository show --location
`
