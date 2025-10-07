// Package purge provides the command for managing GraphDB versions.
package purge

import (
	"context"
	"graphdbcli/internal/tui/common_components"

	"github.com/spf13/cobra"
)

func Command(ctx context.Context) *cobra.Command {
	command := &cobra.Command{
		Use:     "purge <workbench-name>",
		Short:   shortDescription,
		Long:    longDescription,
		Example: common_components.PadExamples(examples),
		Aliases: []string{"remove", "rm", "p"},
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return cmd.Help()
			}

			purgeWorkbench(args[0])

			return nil
		},
	}

	return command
}
