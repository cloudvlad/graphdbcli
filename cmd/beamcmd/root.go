package beamcmd

import (
	"context"

	"github.com/spf13/cobra"
)

var graphdbAddr string
var beamPort string

func Beam(ctx context.Context, ctxCancel context.CancelFunc) *cobra.Command {
	command := &cobra.Command{
		Use:   "beam",
		Short: shortDescription,
		Long:  longDescription,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := StartBeamProxyServer(graphdbAddr, beamPort)
			if err != nil {
				return err
			}
			return nil
		},
	}

	command.Flags().StringVarP(&graphdbAddr, "graphdb", "g", "http://localhost:7200", "graphdb address")
	command.Flags().StringVarP(&beamPort, "beamport", "p", "7199", "beam port")

	return command
}
