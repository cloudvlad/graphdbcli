// Package join provides the command for joining a node to a cluster.
package join

import "github.com/spf13/cobra"

func Command() *cobra.Command {
	var command = &cobra.Command{
		Use:     "join",
		Short:   "Join a node to form a cluster",
		Example: "destroy <name>\n",
		Aliases: []string{"a"},
		RunE: func(cmd *cobra.Command, args []string) error {

			return nil
		},
	}

	command.Flags().UintVarP(&electionMinTimeout, "electionMinTimeout", "", 8000, "election minimum timeout")
	command.Flags().UintVarP(&electionRangeTimeout, "electionRangeTimeout", "", 6000, "election range timeout")
	command.Flags().UintVarP(&heartbeatInterval, "heartbeatInterval", "", 2000, "heartbeat interval")
	command.Flags().UintVarP(&messageSizeKB, "messageSizeKB", "", 64, "message size KB")
	command.Flags().UintVarP(&transactionLogMaximumSizeGB, "transactionLogMaximumSizeGB", "", 50, "transaction log maximum size")
	command.Flags().UintVarP(&verificationTimeout, "verificationTimeout", "", 1500, "verification timeout")

	return command
}
