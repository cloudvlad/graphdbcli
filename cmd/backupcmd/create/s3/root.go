// Package s3cmd provides the command for creating S3 backups.
package s3cmd

import (
	"context"
	"graphdbcli/internal/data_objects/authentication"
	"graphdbcli/internal/data_objects/backup_conf"
	"graphdbcli/internal/tui/common_components"

	"github.com/spf13/cobra"
)

var (
	s3ServiceName  = ""
	s3BucketName   = ""
	s3Region       = ""
	accessKeyId    = ""
	accessKeyToken = ""
)

func Command(ctx context.Context, ctxCancel context.CancelFunc) *cobra.Command {
	command := &cobra.Command{
		Use:     "s3 --bucket <bucket name>",
		Short:   shortDescription,
		Long:    longDescription,
		Example: common_components.PadExamples(examples),
		RunE: func(cmd *cobra.Command, args []string) error {
			authentication.SetupS3Authentication()

			CreateS3Backup(
				backup_conf.Configurations,
				s3ServiceName,
				s3BucketName,
				s3Region,
				accessKeyId,
				accessKeyToken,
				ctx,
				ctxCancel,
			)

			return nil
		},
	}

	command.Flags().StringVarP(&s3ServiceName, "service", "", "", "S3-compatible service")
	command.Flags().StringVarP(&s3BucketName, "bucket", "b", "", "Bucket name (and prefix)")
	command.Flags().StringVarP(&s3Region, "region", "", "", "S3 Region")
	command.Flags().StringVarP(&accessKeyId, "access-key-id", "", "", "Access Key ID")
	command.Flags().StringVarP(&accessKeyToken, "access-key-token", "", "", "Access Key Token")

	command.MarkFlagsRequiredTogether("access-key-id", "access-key-token")
	command.MarkFlagRequired("bucket")

	return command
}
