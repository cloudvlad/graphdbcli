package local

import (
	"context"
	"graphdbcli/internal/data_objects/backup_conf"
	"graphdbcli/internal/tui/common_components"

	"github.com/spf13/cobra"
)

var backupSaveDirPath string

func Command(ctx context.Context, ctxCancel context.CancelFunc) *cobra.Command {
	command := &cobra.Command{
		Use:     "local",
		Short:   shortDescription,
		Long:    longDescription,
		Example: common_components.PadExamples(examples),
		RunE: func(cmd *cobra.Command, args []string) error {

			createLocalBackup(backup_conf.Configurations, backupSaveDirPath, ctx, ctxCancel)

			return nil
		},
	}

	command.PersistentFlags().StringVarP(&backupSaveDirPath, "backupSaveDirPath", "d", "./", "save directory path")

	return command
}
