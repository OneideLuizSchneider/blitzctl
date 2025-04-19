/*
Copyright © 2025 Oneide Luiz Schneider
*/
package cluster

import (
	"github.com/OneideLuizSchneider/blitzctl/cmd/cluster/kind"
	"github.com/OneideLuizSchneider/blitzctl/cmd/cluster/minikube"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install a k8s cluster",
	Long: `Install a k8s cluster
This command will install a k8s cluster
using the specified driver and configuration.
You can use this command to quickly set up a k8s cluster
for development and testing purposes.`,
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
	installCmd.AddCommand(minikube.NewInstallMinikubeCmd())
	installCmd.AddCommand(kind.NewInstallKindCmd())
}
