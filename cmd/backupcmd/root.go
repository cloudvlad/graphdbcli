// Package backupcmd provides the command for managing backups.
package backupcmd

import (
	"context"
	"graphdbcli/cmd/backupcmd/create"
	"graphdbcli/internal/data_objects/authentication"
	"graphdbcli/internal/data_objects/backup_conf"
	"graphdbcli/internal/tui/common_components"

	"github.com/spf13/cobra"
)

func Command(ctx context.Context, ctxCancel context.CancelFunc) *cobra.Command {
	command := &cobra.Command{
		Use:     "backup",
		Short:   shortDescription,
		Long:    longDescription,
		Example: common_components.PadExamples(examples),
		Aliases: []string{"backups", "b"},
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return cmd.Help()
			}

			// Set up all the related authentication
			authentication.SetupGraphDBAuthentication()

			return nil
		},
	}

	command.PersistentFlags().StringVarP(&authentication.BasicCredentials.Username, "username", "u", "", "GraphDB username")
	command.PersistentFlags().StringVarP(&authentication.BasicCredentials.Password, "password", "p", "", "GraphDB password")
	command.PersistentFlags().StringVarP(&authentication.AuthToken.AuthToken, "authToken", "t", "", "GraphDB authentication token")

	command.PersistentFlags().StringVarP(&backup_conf.Configurations.GraphDBLocation, "location", "l", "http://localhost:7200", "Location of the GraphDB instance")
	command.PersistentFlags().StringSliceVarP(&backup_conf.Configurations.Repositories, "repositories", "r", []string{}, "names of the backed up repositories")
	command.PersistentFlags().BoolVarP(&backup_conf.Configurations.BackupSystemData, "backupSystemData", "s", false, "Includes the system data in the backup")
	command.PersistentFlags().StringVarP(&backup_conf.Configurations.BackupName, "backupName", "n", "", "Command name")

	command.MarkFlagsRequiredTogether("username", "password")

	command.AddCommand(create.Command(ctx, ctxCancel))

	return command
}
