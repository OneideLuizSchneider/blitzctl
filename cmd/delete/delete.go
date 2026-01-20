/*
Copyright © 2025 Oneide Luiz Schneider
*/
package delete

import (
	"github.com/spf13/cobra"
	"k8s.io/kubectl/pkg/util/i18n"
	"k8s.io/kubectl/pkg/util/templates"
)

var (
	deleteExamples = templates.Examples(i18n.T(`
		# Delete a minikube cluster (default provider)
		blitzctl delete cluster --cluster-name=mycluster
		
		# Delete a minikube cluster
		blitzctl delete cluster --provider minikube --cluster-name=mycluster
	`))

	deleteCmd = &cobra.Command{
		Use:     "delete",
		Short:   "Delete resources",
		Long:    `Delete resources such as local Kubernetes clusters.`,
		Example: deleteExamples,
		Run: func(cmd *cobra.Command, args []string) {
			if err := cmd.Help(); err != nil {
				cmd.PrintErrln("❌ Error displaying help:", err)
			}
		},
	}
)

func GetDeleteCmd() *cobra.Command {
	return deleteCmd
}

func init() {
	deleteCmd.AddCommand(clusterCmd)
}
