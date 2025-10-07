package start

import (
	"context"
	"fmt"
	cc "graphdbcli/internal/channels/commons"
	"graphdbcli/internal/tui/common_components"
	sp "graphdbcli/internal/tui/workbenchtui/spinner"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"

	"graphdbcli/internal/tool_configurations/initialization"

	tea "github.com/charmbracelet/bubbletea"
)

var p *tea.Program

func startCustomWorkbench(workbenchName string, ctx context.Context, ctxCancel context.CancelFunc) {
	workbenchDir := filepath.Join(initialization.GetWorkbenchDirectory(), workbenchName)

	p = tea.NewProgram(common_components.InitialModel(ctx, ctxCancel, sp.InitializingWorkbenchStatuses, &cc.Success, &cc.Failure))
	go func() {
		p.Run()
	}()

	if _, err := os.Stat(workbenchDir); os.IsNotExist(err) {
		fmt.Printf("workbench path does not exist: %s\n", workbenchDir)
		cc.HandleEvent(&cc.Failure, p)
		return
	}

	cmd := exec.Command("npm", "run", "clean-install")
	cmd.Dir = workbenchDir

	if err := cmd.Run(); err != nil {
		fmt.Printf("npm run clean-install failed: %v\n", err)
		cc.HandleEvent(&cc.Failure, p)
		return
	}

	cc.HandleEvent(&cc.Success, p)

	p = tea.NewProgram(common_components.InitialModel(ctx, ctxCancel, sp.StartingWorkbenchStatuses, &cc.Success, &cc.Failure))
	go func() {
		p.Run()
	}()

	cmd = exec.Command("npm", "run", "start")
	cmd.Dir = workbenchDir

	// Redirect stdout/stderr to log files so the subprocess doesn't block on IO
	outPath := filepath.Join(workbenchDir, "instance.out.log")
	errPath := filepath.Join(workbenchDir, "instance.err.log")
	outFile, outErr := os.OpenFile(outPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if outErr != nil {
		fmt.Printf("failed to open stdout log: %v\n", outErr)
	}
	errFile, errErr := os.OpenFile(errPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if errErr != nil {
		fmt.Printf("failed to open stderr log: %v\n", errErr)
	}
	if outFile != nil {
		cmd.Stdout = outFile
	}
	if errFile != nil {
		cmd.Stderr = errFile
	}

	// Start the process (don't wait) so it runs as a subprocess
	if err := cmd.Start(); err != nil {
		fmt.Printf("npm run start failed to start: %v\n", err)
		cc.HandleEvent(&cc.Failure, p)
		return
	}

	// Close the file descriptors in the parent process; child has its own copy
	if outFile != nil {
		_ = outFile.Close()
	}
	if errFile != nil {
		_ = errFile.Close()
	}

	// Persist the PID so other tooling can manage this instance
	pidFilePath := filepath.Join(workbenchDir, ".instance_pid")
	pidStr := strconv.Itoa(cmd.Process.Pid)
	if werr := os.WriteFile(pidFilePath, []byte(pidStr), 0644); werr != nil {
		fmt.Printf("failed to write pid file: %v\n", werr)
		// attempt to kill the started process to avoid leaking
		_ = cmd.Process.Kill()
		cc.HandleEvent(&cc.Failure, p)
		return
	}

	cc.HandleEvent(&cc.Success, p)
}
