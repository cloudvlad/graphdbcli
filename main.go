package main

import (
	"graphdbcli/cmd"
	tc "graphdbcli/internal/tool_configurations/initialization"
	"graphdbcli/internal/tool_configurations/logging"

	"go.uber.org/zap"
)

func main() {
	tc.InitializeCLIHomeDirectory()
	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {
			logger.Fatal("Failed to sync logger", zap.Error(err))
		}
	}(logging.GetLogger())
	err := cmd.Execute()
	if err != nil {
		println(err.Error())
		return
	}
}
