/*
Copyright © 2025 Oneide Luiz Schneider
*/
package cluster

import (
	"github.com/OneideLuizSchneider/blitzctl/cmd/cluster/provider"
	"github.com/spf13/cobra"
	"k8s.io/kubectl/pkg/util/i18n"
	"k8s.io/kubectl/pkg/util/templates"
)

var (
	upgradeExamples = templates.Examples(i18n.T(`
		# Minikube Doc: https://minikube.sigs.k8s.io/docs/start/

		# minikube
		blitzctl cluster upgrade minikube --cluster-name=mycluster --k8s-version=1.33.1
		# kind (not supported - will show error)
		blitzctl cluster upgrade kind
	`))

	upgradeCmd = &cobra.Command{
		Use:   "upgrade",
		Short: "Upgrade a k8s engine",
		Long: `Upgrade a k8s Cluster engine
This command will upgrade a local k8s engine
using the specified driver and configuration.`,
		Example: upgradeExamples,
		Aliases: []string{"u"},
		Args:    cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			if err := cmd.Help(); err != nil {
				cmd.PrintErrln("❌ Error displaying help:", err)
			}
		},
	}
)

func init() {
	factory := provider.DefaultFactory

	// Register provider commands dynamically
	for _, providerType := range factory.GetSupportedProviders() {
		if clusterProvider, err := factory.CreateProvider(providerType); err == nil {
			upgradeCmd.AddCommand(clusterProvider.GetUpgradeCommand())
		}
	}
}
