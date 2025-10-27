package create

import (
	"fmt"
	"graphdbcli/internal/tool_configurations/initialization"
	"graphdbcli/internal/tui/common_components"
	"os"
	"path"

	"github.com/enescakir/emoji"
)

func CleanUp(instanceName string) error {
	instancePath := path.Join(initialization.GetWorkbenchesDirectory(), instanceName)
	fmt.Printf("%s Cleaning up instance after failure\n", common_components.PadStatusIndicator(emoji.Broom.String(), 0))

	return os.RemoveAll(instancePath)
}
