// Package versioncmd /*
//
// Copntains the logic for the GraphDB
// version managing capabilities.
package versioncmd

import (
	"graphdbcli/cmd/versioncmd/install"
	"graphdbcli/cmd/versioncmd/list"
	"github.com/spf13/cobra"
)

func Version() *cobra.Command {
	command := &cobra.Command{
		Use:     "version",
		Short:   "Manages GraphDB versions",
		Example: "version",
		Aliases: []string{"v", "versions"},
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return cmd.Help()
			}
			return nil
		},
	}

	command.AddCommand(list.Command())
	command.AddCommand(install.Command())

	return command
}
