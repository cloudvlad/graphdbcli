// Package perf_test provides the command for testing instance performance.
package perf_test

import (
	"graphdbcli/internal/tool_configurations/logging"
	"graphdbcli/internal/tui/common_components"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func Command() *cobra.Command {
	var command = &cobra.Command{
		Use:     "perf_test",
		Short:   shortDescription,
		Long:    longDescription,
		Example: common_components.PadExamples(examples),
		RunE: func(cmd *cobra.Command, args []string) error {
			ContextIndexTests()

			return nil
		},
	}

	command.Flags().StringVarP(&instanceAddress, "instanceAddress", "a", "", "instance address")
	command.Flags().StringVarP(&loaderQueriesDir, "loaders", "l", "", "directory with executeLoaders queries")
	command.Flags().StringVarP(&testQueriesDir, "tests", "t", "", "directory with select queries")
	command.Flags().StringVarP(&repoName, "repositoryName", "r", "", "repository name")
	command.Flags().IntVarP(&runs, "runs", "", 1, "number of runs for each test")
	command.Flags().BoolVarP(&markdownResultTable, "md", "", false, "generates a Markdown table with the results")

	err := cobra.MarkFlagRequired(command.Flags(), "instanceAddress")
	if err != nil {
		logging.LOGGER.Error("Error marking flag as required", zap.Error(err), zap.String("flag", "instanceAddress"), zap.String("command", command.Use))
		return nil
	}

	err = cobra.MarkFlagRequired(command.Flags(), "repositoryName")
	if err != nil {
		logging.LOGGER.Error("Error marking flag as required", zap.Error(err), zap.String("flag", "repositoryName"), zap.String("command", command.Use))
		return nil
	}

	return command
}
