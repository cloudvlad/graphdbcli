// Package add provides the command for managing GraphDB versions.
package add

import (
	"context"
	"graphdbcli/internal/tui/common_components"

	"github.com/spf13/cobra"
)

func Command(ctx context.Context) *cobra.Command {
	command := &cobra.Command{
		Use:     "add <workbench-name> <path-to-source>",
		Short:   shortDescription,
		Long:    longDescription,
		Example: common_components.PadExamples(examples),
		Aliases: []string{"a"},
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 2 {
				return cmd.Help()
			}

			AddCustomWorkbench(args[0], args[1])

			return nil
		},
	}

	return command
}
