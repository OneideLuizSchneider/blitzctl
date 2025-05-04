/*
Copyright © 2025 Oneide Luiz Schneider
*/
package cluster

import (
	"github.com/OneideLuizSchneider/blitzctl/cmd/cluster/kind"
	"github.com/OneideLuizSchneider/blitzctl/cmd/cluster/minikube"
	"github.com/spf13/cobra"
	"k8s.io/kubectl/pkg/util/i18n"
	"k8s.io/kubectl/pkg/util/templates"
)

var (
	upgradeExamples = templates.Examples(i18n.T(`
		# Minikube Doc: https://minikube.sigs.k8s.io/docs/start/

		# minikube
		blitzctl clusters upgrade minikube
		# kind
		blitzctl clusters upgrade kind
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
	upgradeCmd.AddCommand(minikube.NewUpgradeMinikubeCmd())
	upgradeCmd.AddCommand(kind.NewUpgradeKindCmd())
}
