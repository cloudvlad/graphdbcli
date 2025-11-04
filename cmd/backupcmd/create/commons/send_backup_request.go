package commons

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	cc "graphdbcli/internal/channels/commons"
	"graphdbcli/internal/data_objects/authentication"
	"graphdbcli/internal/data_objects/backup_conf"
	"graphdbcli/internal/tool_configurations/logging"
	s "graphdbcli/internal/tui/backuptui/spinner"
	"graphdbcli/internal/tui/common_components"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/enescakir/emoji"
	"go.uber.org/zap"
)

// SendBackupRequest is a generic function for creating backups.
// if the bucketUri is not empty string, it will use the cloud backup endpoint and specific header
// if it is, will create a local backup, and will use another header
func SendBackupRequest(configurations backup_conf.BackupConfigurations, bucketUri, localBackupSaveDirPath string, tui *tea.Program, ctx context.Context, ctxCancel context.CancelFunc) {

	params := make(map[string]interface{})
	if len(configurations.Repositories) > 0 {
		params["repositories"] = configurations.Repositories
	}

	if configurations.BackupSystemData {
		params["backupSystemData"] = configurations.BackupSystemData
	}

	if bucketUri != "" {
		params["bucketUri"] = bucketUri
	}

	paramsJson, err := json.Marshal(params)
	if err != nil {
		cc.HandleEvent(&cc.Failure, tui)
		logging.LOGGER.Fatal("Failed to marshal params", zap.Error(err))
	}

	var body *bytes.Buffer
	var writer *multipart.Writer
	if bucketUri != "" {
		body = &bytes.Buffer{}
		writer = multipart.NewWriter(body)
		fw, err := writer.CreateFormField("params")
		if err != nil {
			cc.HandleEvent(&cc.Failure, tui)
			logging.LOGGER.Fatal("Failed to create form field", zap.Error(err))
		}
		if _, err := fw.Write(paramsJson); err != nil {
			cc.HandleEvent(&cc.Failure, tui)
			logging.LOGGER.Fatal("Failed to write params to form: %v", zap.Error(err))
		}
		if err := writer.Close(); err != nil {
			cc.HandleEvent(&cc.Failure, tui)
			logging.LOGGER.Fatal("Failed to close multipart writer", zap.Error(err))
		}
		logging.LOGGER.Debug("Multipart form prepared", zap.ByteString("params", paramsJson))
	} else {
		body = bytes.NewBuffer(paramsJson)
		logging.LOGGER.Debug("JSON body prepared", zap.ByteString("params", paramsJson))
	}

	var endpoint string
	if bucketUri != "" {
		logging.LOGGER.Debug("Creating Cloud backups")
		endpoint = "/rest/recovery/cloud-backup"
	} else {
		logging.LOGGER.Debug("Creating local backup")
		endpoint = "/rest/recovery/backup"
	}

	// Prepare request
	url := strings.TrimRight(configurations.GraphDBLocation, "/") + endpoint
	req, err := http.NewRequestWithContext(ctx, "POST", url, body)
	if err != nil {
		log.Fatal("Failed to create HTTP request", zap.Error(err))
	}

	if bucketUri != "" && writer != nil {
		// Set Content-Type to multipart/form-data with boundary
		req.Header.Set("Content-Type", writer.FormDataContentType())
		req.Header.Set("Accept", "text/plain")
	} else {
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Accept", "*/*")
	}

	if authentication.AuthToken.AuthToken != "" {
		req.Header.Set("Authorization", "Bearer "+authentication.AuthToken.AuthToken)
		logging.LOGGER.Debug("Using Bearer token authentication.")
	} else {
		if authentication.BasicCredentials.Username != "" && authentication.BasicCredentials.Password != "" {
			req.SetBasicAuth(authentication.BasicCredentials.Username, authentication.BasicCredentials.Password)
			logging.LOGGER.Debug("Using basic authentication user", zap.String("username", authentication.BasicCredentials.Username))
		}
	}

	cc.HandleEvent(&cc.Success, tui)
	tui = tea.NewProgram(common_components.InitialModel(ctx, ctxCancel, s.BackupCreationStatuses, &cc.Success, &cc.Failure))
	go func() {
		_, err = tui.Run()
		if err != nil {
			logging.LOGGER.Fatal("TUI was not started", zap.Error(err))
			return
		}
	}()

	logging.LOGGER.Info("Creating the backup...")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		if ctx.Err() != nil {
			logging.LOGGER.Fatal("Request canceled by context", zap.Error(ctx.Err()))
		}
		cc.HandleEvent(&cc.Failure, tui)
		logging.LOGGER.Fatal("Request failed", zap.Error(err))
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		b, _ := io.ReadAll(resp.Body)
		cc.HandleEvent(&cc.Failure, tui)
		fmt.Printf("%s An error occured while backup creation was in progress. Status code %s.", emoji.Information, resp.Status)
		logging.LOGGER.Fatal("Backup failed", zap.String("status", resp.Status), zap.String("text", string(b)))
	}

	// If localBackupSaveDirPath is set, save the response body to a file named by configurations.BackupName in that directory
	if localBackupSaveDirPath != "" {
		backupFileName := configurations.BackupName
		if backupFileName == "" {
			backupFileName = "backup.tar"
		}
		fullPath := localBackupSaveDirPath
		if !strings.HasSuffix(localBackupSaveDirPath, "/") {
			fullPath += "/"
		}
		fullPath += backupFileName
		outFile, err := os.Create(fullPath)
		if err != nil {
			cc.HandleEvent(&cc.Failure, tui)
			logging.LOGGER.Fatal("Failed to create backup file", zap.Error(err))
		}
		defer outFile.Close()
		_, err = io.Copy(outFile, resp.Body)
		if err != nil {
			cc.HandleEvent(&cc.Failure, tui)
			logging.LOGGER.Fatal("Failed to write backup to file", zap.Error(err))
		}
		logging.LOGGER.Info("Backup saved to file", zap.String("path", fullPath))
	} else {
		logging.LOGGER.Info("Backup request finished successful!")
	}
	cc.HandleEvent(&cc.Success, tui)
}
