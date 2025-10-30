package s3cmd

import (
	"context"
	"graphdbcli/cmd/backupcmd/create/commons"
	channels "graphdbcli/internal/channels/commons"
	"graphdbcli/internal/data_objects/backup_conf"
	"graphdbcli/internal/tool_configurations/logging"
	s "graphdbcli/internal/tui/backuptui/spinner"
	"graphdbcli/internal/tui/common_components"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"go.uber.org/zap"
)

var tui *tea.Program
var log = logging.LOGGER

func CreateS3Backup(
	configurations backup_conf.BackupConfigurations,
	s3ServiceName,
	s3BucketName,
	s3Region,
	accessKeyId,
	accessKeyToken string,
	ctx context.Context,
	ctxCancel context.CancelFunc,
) {
	logging.LOGGER.Debug("Preparing for creation of S3 backup...", zap.Strings("repositories", configurations.Repositories),
		zap.Bool("backupSystemData", configurations.BackupSystemData), zap.String("bucketName", s3BucketName), zap.String("region", s3Region))

	tui = tea.NewProgram(common_components.InitialModel(ctx, ctxCancel, s.BackupPreparationStatuses, &channels.Success, &channels.Failure))
	go func() {
		_, err := tui.Run()
		if err != nil {
			logging.LOGGER.Fatal("TUI was not started", zap.Error(err))
			return
		}
	}()

	bucketUri := constructBucketUri(s3ServiceName, s3BucketName, s3Region, accessKeyId, accessKeyToken, configurations.BackupName)

	commons.SendBackupRequest(configurations, bucketUri, "", tui, ctx, ctxCancel)

	return
}

// constructBucketUri constructs the request that will be used for sending the URL
func constructBucketUri(s3ServiceName, s3BucketName, s3Region, accessKeyId, accessKeyToken, backupName string) string {
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
