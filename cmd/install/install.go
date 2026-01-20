/*
Copyright © 2025 Oneide Luiz Schneider
*/
package install

import (
	"github.com/spf13/cobra"
	"k8s.io/kubectl/pkg/util/i18n"
	"k8s.io/kubectl/pkg/util/templates"
)

var (
	installExamples = templates.Examples(i18n.T(`
		# Install minikube (default provider)
		blitzctl install cluster
		
		# Install kind
		blitzctl install cluster --provider kind
	`))

	installCmd = &cobra.Command{
		Use:     "install",
		Short:   "Install resources",
		Long:    `Install resources such as local Kubernetes tools.`,
		Example: installExamples,
		Run: func(cmd *cobra.Command, args []string) {
			if err := cmd.Help(); err != nil {
				cmd.PrintErrln("❌ Error displaying help:", err)
			}
		},
	}
)

func GetInstallCmd() *cobra.Command {
	return installCmd
}

func init() {
	installCmd.AddCommand(clusterCmd)
}
