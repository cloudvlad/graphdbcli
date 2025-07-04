// Package create provides the command for creating backups.
package create

import (
	"context"
	s3cmd "graphdbcli/cmd/backupcmd/create/s3"
	"graphdbcli/internal/tui/common_components"

	"github.com/spf13/cobra"
)

func Command(ctx context.Context, ctxCancel context.CancelFunc) *cobra.Command {
	var command = &cobra.Command{
		Use:     "create",
		Short:   shortDescription,
		Long:    longDescription,
		Example: common_components.PadExamples(examples),
		Aliases: []string{"c"},
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}

	command.AddCommand(s3cmd.Command(ctx, ctxCancel))

	return command
}
