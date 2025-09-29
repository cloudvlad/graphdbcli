// Package cmd provides the root command for the CLI application.
// Every subcommand is defined in its own package.
package cmd

import (
	"context"
	"graphdbcli/cmd/backupcmd"
	"graphdbcli/cmd/gizmocmd"
	"graphdbcli/cmd/instancecmd"
	"graphdbcli/cmd/licensecmd"
	"graphdbcli/cmd/repositorycmd"
	"graphdbcli/cmd/resourcecmd"
	"graphdbcli/cmd/versioncmd"

	"github.com/spf13/cobra"
)

// Execute is the command line applications entry function
func Execute() error {
	rootCmd := &cobra.Command{
		Version: "v0.1.0",
		Use:     "graphdbcli",
		Short:   shortDescription,
		Long:    longDescription,
		Example: "graphdbcli",
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	rootCmd.AddCommand(versioncmd.Version(ctx))
	rootCmd.AddCommand(licensecmd.License(ctx))
	rootCmd.AddCommand(instancecmd.Cluster(ctx, cancel))
	rootCmd.AddCommand(gizmocmd.Gizmo())
	rootCmd.AddCommand(backupcmd.Backup(ctx, cancel))
	rootCmd.AddCommand(resourcecmd.Resource(ctx))
	rootCmd.AddCommand(repositorycmd.Repository(ctx))

	return rootCmd.ExecuteContext(ctx)
}
