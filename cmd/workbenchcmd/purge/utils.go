package purge

import (
	"fmt"
	"os"
	"path/filepath"

	"graphdbcli/internal/tool_configurations/initialization"
	"graphdbcli/internal/tool_configurations/logging"

	"go.uber.org/zap"
)

// purgeWorkbench removes the specified workbench directory under the
// configured workbench directory. It performs safety checks and logs the
// outcome.
func purgeWorkbench(workbenchName string) {
	workbenchesDir := initialization.GetWorkbenchDirectory()
	workbenchPath := filepath.Join(workbenchesDir, workbenchName)

	_, err := os.Stat(workbenchPath)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("Workbench '%s' does not exist\n", workbenchName)
			return
		}

		logging.LOGGER.Error("failed to stat workbench path", zap.Error(err))
		fmt.Printf("Error accessing workbench '%s': %v\n", workbenchName, err)

		return
	}

	// Perform removal
	if err := os.RemoveAll(workbenchPath); err != nil {
		logging.LOGGER.Error("failed to remove workbench directory", zap.Error(err))
		fmt.Printf("Failed to remove workbench '%s': %v\n", workbenchName, err)
		return
	}

	fmt.Printf("Workbench '%s' removed\n", workbenchName)
	logging.LOGGER.Info("workbench purged", zap.String("workbench", workbenchName), zap.String("path", workbenchPath))
}
