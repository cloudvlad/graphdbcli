// Package commons provides functions that are used across many instance commands
package commons

import (
	"context"
	"fmt"
	"graphdbcli/internal/channels/commons"
	c "graphdbcli/internal/data_objects/graphdb_cluster"
	ini "graphdbcli/internal/tool_configurations/initialization"
	"graphdbcli/internal/tool_configurations/logging"
	tc "graphdbcli/internal/tool_configurations/statics"
	"graphdbcli/internal/tui/common_components"
	sp "graphdbcli/internal/tui/instancetui/create"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/enescakir/emoji"
	"go.uber.org/zap"
)

var p *tea.Program

// StartInstance start the instance by default, unless the
// necessary flag was set to override this behaviour.
func StartInstance(ctx context.Context, ctxCancel context.CancelFunc, successChannel *chan bool, failureChannel *chan bool) {
	p = tea.NewProgram(common_components.InitialModel(ctx, ctxCancel, sp.StartingGraphDBInstance, successChannel, failureChannel))
	go func() {
		p.Run()
	}()

	if c.Instance.IsActive {
		logging.LOGGER.Debug("Instance" + c.Instance.Name + " is set to active")
	} else {
		logging.LOGGER.Debug("Instance " + c.Instance.Name + " is not active")
		return
	}

	logging.LOGGER.Info("Starting cluster",
		zap.String("cluster name", c.Instance.Name))

	segments := append([]string{ini.GetWorkbenchesDirectory(), c.Instance.Name}, tc.DefaultExecutablePath...)
	executablePath := path.Join(segments...)

	logging.LOGGER.Debug("Starting the GraphDB instance...",
		zap.String("executablePath", executablePath))
	err := startMainProcess(executablePath)
	if err != nil {
		commons.HandleEvent(failureChannel, p)
		logging.LOGGER.Fatal("Error starting instance", zap.Error(err))
	}

	commons.HandleEvent(successChannel, p)
}

func startMainProcess(executablePath string) error {
	cmd := exec.Command(executablePath)

	// Start the process in the background
	if err := cmd.Start(); err != nil {
		fmt.Printf("%s Error uncounted while starting the process", common_components.PadStatusIndicator(emoji.CheckMark.String(), tc.NotTUIStatusIndicatorAdditionalPadding))
		logging.LOGGER.Fatal("Error uncounted when starting the process",
			zap.String("executablePath", executablePath),
			zap.Error(err))
		return err
	}

	parentDir := filepath.Dir(executablePath)
	grandParentDir := filepath.Dir(parentDir)
	pidFile := filepath.Join(grandParentDir, ".instance_pid")
	pid := strconv.Itoa(cmd.Process.Pid)
	if err := os.WriteFile(pidFile, []byte(pid), 0644); err != nil {
		fmt.Println("Error uncounted while writing process id to file")
		logging.LOGGER.Fatal("Error uncounted while writing process id to file",
			zap.String("executablePath", executablePath),
			zap.Error(err))
		return err
	}

	return nil
}
