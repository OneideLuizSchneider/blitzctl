/*
Copyright © 2025 Oneide Luiz Schneider
*/
package stop

import (
	"github.com/spf13/cobra"
	"k8s.io/kubectl/pkg/util/i18n"
	"k8s.io/kubectl/pkg/util/templates"
)

var (
	stopExamples = templates.Examples(i18n.T(`
		# Stop a minikube cluster (default provider)
		blitzctl stop cluster --cluster-name <cluster-name>
		
		# Stop a kind cluster
		blitzctl stop cluster --provider kind --cluster-name <cluster-name>
	`))

	stopCmd = &cobra.Command{
		Use:     "stop",
		Short:   "Stop resources",
		Long:    `Stop resources such as local Kubernetes clusters.`,
		Example: stopExamples,
		Run: func(cmd *cobra.Command, args []string) {
			if err := cmd.Help(); err != nil {
				cmd.PrintErrln("❌ Error displaying help:", err)
			}
		},
	}
)

func GetStopCmd() *cobra.Command {
	return stopCmd
}

func init() {
	stopCmd.AddCommand(clusterCmd)
}
