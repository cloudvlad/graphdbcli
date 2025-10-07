package stop

var examples = `# Stop running Workbench instances
graphdbcli workbench stop my-instance
# Shortcut for stopping a running Workbench instance
graphdbcli wb s my-instance`

var shortDescription = `Stops a running Workbench instance`

var longDescription = `Stops a running Workbench instance

Stops a running GraphDB Workbench instance by proving the instance name.
It uses the process ID that is stored inside the .instance_pid that is placed
inside the dedicated workbench directory name.
`
