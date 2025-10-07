// Package start provides the command for managing GraphDB versions.
package start

import (
	"context"
	"graphdbcli/internal/tui/common_components"

	"github.com/spf13/cobra"
)

func Command(ctx context.Context, ctxCancel context.CancelFunc) *cobra.Command {
	command := &cobra.Command{
		Use:     "start <workbench-name>",
		Short:   shortDescription,
		Long:    longDescription,
		Example: common_components.PadExamples(examples),
		Aliases: []string{"s"},
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return cmd.Help()
			}

			startCustomWorkbench(args[0], ctx, ctxCancel)

			return nil
		},
	}

	return command
}
