package add

import (
	"fmt"
	"graphdbcli/internal/ll_commons"
	"graphdbcli/internal/tool_configurations/initialization"
	"graphdbcli/internal/tool_configurations/logging"
	"graphdbcli/internal/tui/common_components"
	"os"
	"path/filepath"

	"github.com/enescakir/emoji"
	"go.uber.org/zap"
)

// addCustomWorkbench adds new workbench by duplicating the content of the passed for adding workbench source
// code directory
func addCustomWorkbench(workbenchName string, addedWorkbenchPath string) {
	workbenchDir := filepath.Join(initialization.GetWorkbenchDirectory(), workbenchName)

	if _, err := os.Stat(workbenchDir); !os.IsNotExist(err) {
		fmt.Printf("%s Workbench '%s' already exists\n", common_components.PadStatusIndicator(emoji.CrossMark.String(), 0), workbenchName)
		logging.LOGGER.Fatal("workbench already exists", zap.String("workbench", workbenchName))
	}

	logging.LOGGER.Debug("workbench does not exist", zap.String("workbench", workbenchName))

	if err := os.MkdirAll(workbenchDir, 0755); err != nil {
		fmt.Printf("%s Unable to create interna workbench directory '%s': %v\n", common_components.PadStatusIndicator(emoji.CrossMark.String(), 0), workbenchDir, err)
		logging.LOGGER.Fatal("Unable to create workbench directory", zap.String("workbench", workbenchDir))
	}

	logging.LOGGER.Debug("workbench directory created", zap.String("workbench", workbenchName))

	if err := ll_commons.DuplicateContentInternally(addedWorkbenchPath, workbenchDir); err != nil {
		fmt.Printf("%s Failed to copy workbench files from '%s' to '%s'\n", common_components.PadStatusIndicator(emoji.CrossMark.String(), 0), addedWorkbenchPath, workbenchDir)
		logging.LOGGER.Fatal("failed to copy workbench files", zap.String("workbench source", workbenchDir))
	}

	createWorkbenchMetadata(workbenchDir, workbenchName)

	fmt.Printf("%s Workbench '%s' added\n", common_components.PadStatusIndicator(emoji.RaisingHands.String(), 0), workbenchName)
}
