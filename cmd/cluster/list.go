/*
Copyright © 2025 Oneide Luiz Schneider
*/
package cluster

import (
	"github.com/OneideLuizSchneider/blitzctl/cmd/cluster/provider"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list [subcommand]",
	Short: "List all available local environments",
	Long: `List all available local environments
This command will display all the local environments
that are currently available on your system.
You can use this command to quickly see which environments
are available for use and their current status.`,
	Example: listExample,
	Aliases: []string{"ls"},
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if err := cmd.Help(); err != nil {
			// Handle the error (e.g., log it or print it)
			cmd.PrintErrln("Error displaying help:", err)
		}
	},
}

func init() {
	for _, clusterProvider := range provider.GetProviders() {
		listCmd.AddCommand(clusterProvider.GetListCommand())
	}
}
