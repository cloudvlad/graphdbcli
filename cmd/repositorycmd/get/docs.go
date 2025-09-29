package get

var examples = `
# Fetches the configuration file for a repository
graphdbcli repository get repo1 --location http://localhost:7200
# Fetches the configuration file for a repository and stores it locally
graphdbcli repository get repo1 --location http://localhost:7200 --save
# Fetches the configuration file for a repository from local storage
graphdbcli repository get repo1
`
