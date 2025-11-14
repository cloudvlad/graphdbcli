package add

import (
	"fmt"
	"graphdbcli/internal/tool_configurations/initialization"
)

var examples = `graphdbcli workbench add test-workbench ./new-feature
graphdbcli wb add test-workbench ./new-feature
graphdbcli wb add test-workbench /home/mark/dev/new-feature
`

var shortDescription = `Add a GraphDB Workbench codebase instance`

var longDescription = fmt.Sprintf(`Add a GraphDB Workbench codebase instance.

This command registers a GraphDB Workbench codebase in the CLI's internal
workspace so it can be managed using the CLI. Provide the path to a
directory containing a Workbench codebase; the command copies that
directory into the CLI-managed workbenches directory.

  Workbenches directory: %s
`, initialization.GetWorkbenchDirectory())
