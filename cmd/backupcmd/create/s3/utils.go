package s3cmd

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	cc "graphdbcli/internal/channels/commons"
	channels "graphdbcli/internal/channels/commons"
	"graphdbcli/internal/data_objects/authentication"
	"graphdbcli/internal/tool_configurations/logging"
	s "graphdbcli/internal/tui/backuptui/spinner"
	"graphdbcli/internal/tui/common_components"
	"io"
	"mime/multipart"
	"net/http"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/enescakir/emoji"
	"go.uber.org/zap"
)

var (
	s3ServiceName  = ""
	s3BucketName   = ""
	s3Region       = ""
	accessKeyId    = ""
	accessKeyToken = ""
)

var tui *tea.Program
var log = logging.LOGGER

func CreateS3Backup(
	baseURL string,
	repositories []string,
	backupSystemData bool,
	backupName string,
	ctx context.Context,
	ctxCancel context.CancelFunc,
) {
	logging.LOGGER.Debug("Preparing for creation of S3 backup...")

	tui = tea.NewProgram(common_components.InitialModel(ctx, ctxCancel, s.CreatingBackupStatuses, &channels.Success, &channels.Failure))
	go func() {
		_, err := tui.Run()
		if err != nil {
			logging.LOGGER.Fatal("TUI was not started", zap.Error(err))
			return
		}
	}()

	params := map[string]interface{}{
		"repositories":     repositories,
		"backupSystemData": backupSystemData,
		"bucketUri":        constructRequestURL(backupName),
	}

	log.Debug("Repositories", zap.Strings("repositories", repositories))
	log.Debug("Backup System Data", zap.Bool("backupSystemData", backupSystemData))
	log.Debug("S3 Bucket Name ", zap.String("bucketName", s3BucketName))

	paramsJson, err := json.Marshal(params)
	if err != nil {
		cc.HandleEvent(&channels.Failure, tui)
		log.Fatal("Failed to marshal params: %v", zap.Error(err))
	}

	var body bytes.Buffer
	writer := multipart.NewWriter(&body)

	fw, err := writer.CreateFormField("params")
	if err != nil {
		cc.HandleEvent(&channels.Failure, tui)
		log.Fatal("Failed to create form field", zap.Error(err))
	}

	if _, err := fw.Write(paramsJson); err != nil {
		cc.HandleEvent(&channels.Failure, tui)
		log.Fatal("Failed to write params to form: %v", zap.Error(err))
	}

	if err := writer.Close(); err != nil {
		cc.HandleEvent(&channels.Failure, tui)
		log.Fatal("Failed to close multipart writer", zap.Error(err))
	}
	log.Debug("Multipart form prepared")

	// Prepare request
	url := strings.TrimRight(baseURL, "/") + "/rest/recovery/cloud-backup"
	// Create request bound to the operation context so it is canceled if ctx is canceled
	req, err := http.NewRequestWithContext(ctx, "POST", url, &body)
	if err != nil {
		log.Fatal("Failed to create HTTP request", zap.Error(err))
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Accept", "application/json")

	if authentication.AuthToken.AuthToken != "" {
		req.Header.Set("Authorization", "Bearer "+authentication.AuthToken.AuthToken)
		log.Debug("Using Bearer token authentication.")
	} else {
		if authentication.BasicCredentials.Username != "" && authentication.BasicCredentials.Password != "" {
			req.SetBasicAuth(authentication.BasicCredentials.Username, authentication.BasicCredentials.Password)
			log.Debug("Using basic authentication user", zap.String("username", authentication.BasicCredentials.Username))
		}
	}
	
	cc.HandleEvent(&channels.Success, tui)
	tui = tea.NewProgram(common_components.InitialModel(ctx, ctxCancel, s.BackupCreationStatuses, &channels.Success, &channels.Failure))
	go func() {
		_, err = tui.Run()
		if err != nil {
			logging.LOGGER.Fatal("TUI was not started", zap.Error(err))
			return
		}
	}()

	log.Info("Creating a backup...")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		if ctx.Err() != nil {
			log.Fatal("Request canceled by context", zap.Error(ctx.Err()))
		}
		cc.HandleEvent(&channels.Failure, tui)
		log.Fatal("Request failed", zap.Error(err))
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		b, _ := io.ReadAll(resp.Body)
		cc.HandleEvent(&channels.Failure, tui)
		fmt.Printf("%s An error occured while backup creation was in progres. Status code %s.", emoji.Information, resp.Status)
		log.Fatal("Backup failed: %s: %s", zap.String("status", resp.Status), zap.String("text", string(b)))
	}

	log.Info("Backup request successful!")
	cc.HandleEvent(&channels.Success, tui)

	return
}

// constructRequestURL constructs the request that will be used for sending the URL
func constructRequestURL(backupName string) string {
	var uriBuilder strings.Builder
	uriBuilder.WriteString("s3://")

	if s3ServiceName != "" {
		uriBuilder.WriteString(s3ServiceName)
		log.Debug("Using S3 service:", zap.String("s3ServiceName", s3ServiceName))
	} else {
		log.Debug("No S3 service name provided, using AWS S3.")
	}

	uriBuilder.WriteString("/")
	uriBuilder.WriteString(s3BucketName)
	uriBuilder.WriteString("/")
	uriBuilder.WriteString(backupName)

	log.Debug("Bucket name: %s, Backup name: %s", zap.String("bucketName", s3BucketName), zap.String("backupName", backupName))

	var queryParam []string
	if s3Region != "" {
		queryParam = append(queryParam, "region="+s3Region)
		log.Debug("Region set", zap.String("region", s3Region))
	} else {
		log.Debug("No region specified, omitting from URI.")
	}
	if accessKeyId != "" {
		queryParam = append(queryParam, "AWS_ACCESS_KEY_ID="+accessKeyId)
	}

	if accessKeyToken != "" {
		queryParam = append(queryParam, "AWS_SECRET_ACCESS_KEY="+accessKeyToken)
	}

	if len(queryParam) > 0 {
		uriBuilder.WriteString("?")
		uriBuilder.WriteString(strings.Join(queryParam, "&"))
	}

	return uriBuilder.String()
}
