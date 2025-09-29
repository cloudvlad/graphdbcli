// Package logging /*
//
// The package is responsible for setting up the
// logger used access the application.
package logging

import (
	"fmt"
	"graphdbcli/internal/tool_configurations/statics"
	"os"
	"path/filepath"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var LOGGER = GetLogger()

// GetLogger initializes a ZAP logger
func GetLogger() *zap.Logger {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error: Unable to get user home directory:", err)
		return nil
	}

	gdbDir := filepath.Join(homeDir, statics.HomeDirSpaceName)
	logsDir := filepath.Join(gdbDir, statics.LogsDirName)
	filename := "logs"
	logsFilePath := filepath.Join(logsDir, filename)

	writer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   logsFilePath,
		MaxSize:    100,
		MaxBackups: 3,
		MaxAge:     10,
	})

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		writer,
		zap.DebugLevel,
	)

	logger := zap.New(core)

	logger.Debug("Logger with log rotation initialized")

	return logger
}
