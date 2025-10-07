package add

import (
	"fmt"
	"graphdbcli/internal/tool_configurations/initialization"
	"graphdbcli/internal/tool_configurations/logging"
	"graphdbcli/internal/tui/common_components"
	"io"
	"os"
	"path/filepath"

	"github.com/enescakir/emoji"
	"go.uber.org/zap"
)

func AddCustomWorkbench(workbenchName string, workbenchPath string) {
	workbenchDir := filepath.Join(initialization.GetWorkbenchDirectory(), workbenchName)

	if _, err := os.Stat(workbenchDir); !os.IsNotExist(err) {
		fmt.Printf("%s Workbench '%s' already exists\n", common_components.PadStatusIndicator(emoji.CrossMark.String(), 0), workbenchName)
		logging.LOGGER.Fatal("workbench already exists", zap.String("workbench", workbenchName))
	}

	logging.LOGGER.Debug("workbench does not exist", zap.String("workbench", workbenchName))

	if err := os.MkdirAll(workbenchDir, 0755); err != nil {
		fmt.Printf("%s Unable to create workbench directory '%s': %v\n", common_components.PadStatusIndicator(emoji.CrossMark.String(), 0), workbenchDir, err)
		logging.LOGGER.Fatal("Unable to create workbench directory", zap.String("workbench", workbenchDir))
	}

	logging.LOGGER.Debug("workbench directory created", zap.String("workbench", workbenchName))

	if err := duplicateContentInternally(workbenchPath, workbenchDir); err != nil {
		fmt.Printf("%s Failed to copy workbench files from '%s' to '%s'\n", common_components.PadStatusIndicator(emoji.CrossMark.String(), 0), workbenchPath, workbenchDir)
		logging.LOGGER.Fatal("failed to copy workbench files", zap.String("workbench source", workbenchDir))
	}

	fmt.Printf("%s Workbench '%s' added\n", common_components.PadStatusIndicator(emoji.RaisingHands.String(), 0), workbenchName)
}

func duplicateContentInternally(src string, dest string) error {
	src = filepath.Clean(src)
	dest = filepath.Clean(dest)

	info, err := os.Stat(src)
	if err != nil {
		return err
	}
	if !info.IsDir() {
		return fmt.Errorf("source is not a directory: %s", src)
	}

	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}

		targetPath := filepath.Join(dest, relPath)

		if info.IsDir() {
			if err := os.MkdirAll(targetPath, info.Mode()); err != nil {
				return err
			}
			return nil
		}

		if err := copyFile(path, targetPath, info.Mode()); err != nil {
			return err
		}

		return nil
	})
}

func copyFile(srcFile string, destFile string, mode os.FileMode) error {
	in, err := os.Open(srcFile)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.OpenFile(destFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, mode)
	if err != nil {
		return err
	}
	defer func() {
		_ = out.Close()
	}()

	if _, err := io.Copy(out, in); err != nil {
		return err
	}

	return nil
}
