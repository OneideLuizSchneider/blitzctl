/*
Copyright © 2025 Oneide Luiz Schneider
*/
package cluster

import (
	"github.com/OneideLuizSchneider/blitzctl/cmd/cluster/provider"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install a k8s cluster provider",
	Long: `Install a k8s cluster provider
This command will install a k8s cluster provider
(like kind, minikube, etc.) on your system.
You can use this command to quickly set up the necessary
tools for cluster management.`,
	Example: `blitzctl cluster install minikube`,
	Aliases: []string{"i"},
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if err := cmd.Help(); err != nil {
			cmd.PrintErrln("❌ Error displaying help:", err)
		}
	},
}

func init() {
	for _, clusterProvider := range provider.GetProviders() {
		installCmd.AddCommand(clusterProvider.GetInstallCommand())
	}
}
