package start

import (
	"graphdbcli/internal/tool_configurations/logging"
	"os/exec"

	"go.uber.org/zap"
)

func PrerequisitesAreFulfilled() bool {
	cmd := exec.Command("npm", "-v")
	if err := cmd.Run(); err != nil {
		logging.LOGGER.Error("npm prerequisite not fulfilled", zap.Error(err))
		return false
	}
	logging.LOGGER.Info("npm prerequisite fulfilled")
	return true
}
