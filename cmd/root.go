package cmd

import (
	"context"
	"github.com/spf13/cobra"
	"graphdbcli/cmd/licensecmd"
	"graphdbcli/cmd/versioncmd"
	"graphdbcli/internal/tool_configurations/logging"
	"graphdbcli/internal/tool_configurations/statics"
	"strconv"
)

// Execute is the command line applications entry function
func Execute() error {
	rootCmd := &cobra.Command{
		Version: "v0.1.0",
		Use:     "graphdbcli",
		Short:   "GraphDB Command Line Tool",
		Example: "graphdbcli",
	}

	rootCmd.AddCommand(versioncmd.Version())
	rootCmd.AddCommand(licensecmd.License())
	rootCmd.PersistentFlags().BoolVarP(&statics.IsTuiDisabled, "notui", "", false, "Stops the TUI")

	logging.LOGGER.Info("TUI enabled: " + strconv.FormatBool(statics.IsTuiDisabled))

	return rootCmd.ExecuteContext(context.Background())
}
