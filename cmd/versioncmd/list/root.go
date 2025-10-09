// Package show provides the command for showing available versions.
package list

import (
	"context"
	"graphdbcli/internal/tui/common_components"

	"github.com/spf13/cobra"
)

func Command(ctx context.Context) *cobra.Command {
	var command = &cobra.Command{
		Use:     "list",
		Short:   shortDescription,
		Long:    longDescription,
		Example: common_components.PadExamples(examples),
		Aliases: []string{"l"},
		RunE: func(cmd *cobra.Command, args []string) error {
			listAvailableVersions()
			return nil
		},
	}

	return command
}
