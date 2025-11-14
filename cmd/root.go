// Package cmd provides the root command for the CLI application.
// Every subcommand is defined in its own package.
package cmd

import (
	"context"
	"graphdbcli/cmd/backupcmd"
	"graphdbcli/cmd/beamcmd"
	"graphdbcli/cmd/gizmocmd"
	"graphdbcli/cmd/instancecmd"
	"graphdbcli/cmd/licensecmd"
	"graphdbcli/cmd/rdf4j"
	"graphdbcli/cmd/repositorycmd"
	"graphdbcli/cmd/resourcecmd"
	"graphdbcli/cmd/versioncmd"
	"graphdbcli/cmd/workbenchcmd"

	"github.com/spf13/cobra"
)

var Version = "x.y.z"

// Execute is the command line applications entry function
func Execute() error {
	rootCmd := &cobra.Command{
		Version: Version,
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
	rootCmd.AddCommand(backupcmd.Command(ctx, cancel))
	rootCmd.AddCommand(resourcecmd.Resource(ctx))
	rootCmd.AddCommand(repositorycmd.Repository(ctx, cancel))
	rootCmd.AddCommand(workbenchcmd.Workbench(ctx, cancel))
	rootCmd.AddCommand(rdf4j.Rdf4J(ctx, cancel))
	rootCmd.AddCommand(beamcmd.Beam(ctx, cancel))

	return rootCmd.ExecuteContext(ctx)
}
