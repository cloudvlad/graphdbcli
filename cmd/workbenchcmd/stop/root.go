// Package stop provides the command for managing GraphDB versions.
package stop

import (
	"context"
	"graphdbcli/internal/tui/common_components"

	"github.com/spf13/cobra"
)

func Command(ctx context.Context, ctxCancel context.CancelFunc) *cobra.Command {
	command := &cobra.Command{
		Use:     "stop <workbench-name>",
		Short:   shortDescription,
		Long:    longDescription,
		Example: common_components.PadExamples(examples),
		Aliases: []string{"s"},
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return cmd.Help()
			}

			stopInstance(args[0])

			return nil
		},
	}

	return command
}
