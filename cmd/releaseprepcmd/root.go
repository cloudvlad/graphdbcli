// Package versioncmd provides the command for managing GraphDB versions.
package versioncmd

import (
	"context"
	"graphdbcli/cmd/versioncmd/install"
	"graphdbcli/cmd/versioncmd/list"
	"graphdbcli/internal/tui/common_components"

	"github.com/spf13/cobra"
)

func Command(ctx context.Context) *cobra.Command {
	command := &cobra.Command{
		Use:     "releaseprep",
		Short:   shortDescription,
		Long:    longDescription,
		Example: common_components.PadExamples(examples),
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return cmd.Help()
			}
			return nil
		},
		Hidden: true,
	}

	command.AddCommand(list.Command(ctx))
	command.AddCommand(install.Command(ctx))

	return command
}
