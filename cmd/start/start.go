/*
Copyright © 2025 Oneide Luiz Schneider
*/
package start

import (
	"github.com/spf13/cobra"
	"k8s.io/kubectl/pkg/util/i18n"
	"k8s.io/kubectl/pkg/util/templates"
)

var (
	startExamples = templates.Examples(i18n.T(`
		# Start a minikube cluster (default provider)
		blitzctl start cluster --cluster-name <cluster-name>
		
		# Start a kind cluster
		blitzctl start cluster --provider kind --cluster-name <cluster-name>
	`))

	startCmd = &cobra.Command{
		Use:     "start",
		Short:   "Start resources",
		Long:    `Start resources such as local Kubernetes clusters.`,
		Example: startExamples,
		Run: func(cmd *cobra.Command, args []string) {
			if err := cmd.Help(); err != nil {
				cmd.PrintErrln("❌ Error displaying help:", err)
			}
		},
	}
)

func GetStartCmd() *cobra.Command {
	return startCmd
}

func init() {
	startCmd.AddCommand(clusterCmd)
}
