// Package versioncmd provides the command for managing GraphDB versions.
package versioncmd

import (
	"context"
	"graphdbcli/cmd/versioncmd/install"
	"graphdbcli/cmd/versioncmd/show"

	"github.com/spf13/cobra"
)

func Version(ctx context.Context) *cobra.Command {
	command := &cobra.Command{
		Use:     "version",
		Short:   shortDescription,
		Long:    longDescription,
		Example: examples,
		Aliases: []string{"versions", "v"},
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return cmd.Help()
			}
			return nil
		},
	}

	command.AddCommand(show.Command(ctx))
	command.AddCommand(install.Command(ctx))

	return command
}
