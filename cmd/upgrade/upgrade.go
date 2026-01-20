/*
Copyright © 2025 Oneide Luiz Schneider
*/
package upgrade

import (
	"github.com/spf13/cobra"
	"k8s.io/kubectl/pkg/util/i18n"
	"k8s.io/kubectl/pkg/util/templates"
)

var (
	upgradeExamples = templates.Examples(i18n.T(`
		# Upgrade minikube (default provider)
		blitzctl upgrade cluster
		
		# Upgrade kind
		blitzctl upgrade cluster --provider kind
	`))

	upgradeCmd = &cobra.Command{
		Use:     "upgrade",
		Short:   "Upgrade resources",
		Long:    `Upgrade resources such as local Kubernetes tools.`,
		Example: upgradeExamples,
		Run: func(cmd *cobra.Command, args []string) {
			if err := cmd.Help(); err != nil {
				cmd.PrintErrln("❌ Error displaying help:", err)
			}
		},
	}
)

func GetUpgradeCmd() *cobra.Command {
	return upgradeCmd
}

func init() {
	upgradeCmd.AddCommand(clusterCmd)
}
