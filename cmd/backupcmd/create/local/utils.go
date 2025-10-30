package local

import (
	"context"
	"graphdbcli/cmd/backupcmd/create/commons"
	channels "graphdbcli/internal/channels/commons"
	"graphdbcli/internal/data_objects/backup_conf"
	"graphdbcli/internal/tool_configurations/logging"
	s "graphdbcli/internal/tui/backuptui/spinner"
	"graphdbcli/internal/tui/common_components"

	tea "github.com/charmbracelet/bubbletea"
	"go.uber.org/zap"
)

var tui *tea.Program

func createLocalBackup(configurations backup_conf.BackupConfigurations, backupSaveDirPath string, ctx context.Context, ctxCancel context.CancelFunc) {
	logging.LOGGER.Debug("Preparing for creation of local backup...", zap.Strings("repositories", configurations.Repositories),
		zap.Bool("backupSystemData", configurations.BackupSystemData), zap.String("backupSaveDirPath", backupSaveDirPath))

	tui = tea.NewProgram(common_components.InitialModel(ctx, ctxCancel, s.BackupPreparationStatuses, &channels.Success, &channels.Failure))
	go func() {
		_, err := tui.Run()
		if err != nil {
			logging.LOGGER.Fatal("TUI was not started", zap.Error(err))
			return
		}
	}()

	commons.SendBackupRequest(configurations, "", backupSaveDirPath, tui, ctx, ctxCancel)
}
