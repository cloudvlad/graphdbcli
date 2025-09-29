// Package statements_stats provides the command for retrieving statistics about RDF statements across all repositories.
// Accept the path to the directory containing the properties file for the repositories state.
package statements_stats

import (
	"graphdbcli/internal/tui/common_components"

	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	var command = &cobra.Command{
		Use:     "statements_stats",
		Short:   "Statistics for the RDF statements across all repositories",
		Example: common_components.PadExamples(examples),
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return cmd.Help()
			}

			PrintRDFStatementsStats(args[0])

			return nil
		},
	}
	return command
}
