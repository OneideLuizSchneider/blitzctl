/*
Copyright © 2025 Oneide Luiz Schneider
*/
package list

import (
	"github.com/spf13/cobra"
	"k8s.io/kubectl/pkg/util/i18n"
	"k8s.io/kubectl/pkg/util/templates"
)

var (
	listExamples = templates.Examples(i18n.T(`
		# List minikube clusters (default provider)
		blitzctl list clusters
		
		# List kind clusters
		blitzctl list clusters --provider kind
	`))

	listCmd = &cobra.Command{
		Use:     "list",
		Short:   "List resources",
		Long:    `List resources such as local Kubernetes clusters.`,
		Example: listExamples,
		Run: func(cmd *cobra.Command, args []string) {
			if err := cmd.Help(); err != nil {
				cmd.PrintErrln("❌ Error displaying help:", err)
			}
		},
	}
)

func GetListCmd() *cobra.Command {
	return listCmd
}

func init() {
	listCmd.AddCommand(clusterCmd)
}
