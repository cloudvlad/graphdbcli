package config

import (
	"fmt"
	"graphdbcli/internal/tool_configurations/initialization"
	"graphdbcli/internal/tool_configurations/logging"
	"os"
	"path/filepath"
	"regexp"

	"go.uber.org/zap"
)

var (
	graphdbHost   string
	graphdbPort   uint
	workbenchPort uint
)

func ConfigureCustomWorkbench(workbenchName string) {
	workbenchesDirectory := initialization.GetWorkbenchDirectory()

	workbenchPath := filepath.Join(workbenchesDirectory, workbenchName)

	overrideConfigurations(workbenchPath)
}

func overrideConfigurations(workbenchPath string) {
	webpackFile := filepath.Join(workbenchPath, "webpack.config.dev.js")

	data, err := os.ReadFile(webpackFile)
	if err != nil {
		logging.LOGGER.Debug("webpack config not found or unreadable", zap.Error(err))
		fmt.Printf("No webpack config found at %s, skipping overrides\n", webpackFile)
		return
	}

	content := string(data)

	reHost := regexp.MustCompile(`(?m)const\s+host\s*=\s*['"].*?['"];?`)
	replacement := fmt.Sprintf("const host = '%s';", graphdbHost)
	content = reHost.ReplaceAllString(content, replacement)

	rePortHere := regexp.MustCompile(`(?m)const\s+portHere\s*=\s*\d+;?`)
	replacement = fmt.Sprintf("const portHere = %d;", workbenchPort)
	content = rePortHere.ReplaceAllString(content, replacement)

	rePortThere := regexp.MustCompile(`(?m)const\s+portThere\s*=\s*\d+;?`)
	replacement = fmt.Sprintf("const portThere = %d;", graphdbPort)
	content = rePortThere.ReplaceAllString(content, replacement)

	if string(data) != content {
		if err := os.WriteFile(webpackFile, []byte(content), 0o644); err != nil {
			logging.LOGGER.Error("failed to write webpack config", zap.Error(err))
			fmt.Printf("Failed to write webpack config %s: %v\n", webpackFile, err)
			return
		}
		fmt.Printf("Updated webpack config at %s\n", webpackFile)
	} else {
		fmt.Printf("No overrides applied to %s\n", webpackFile)
	}
}
