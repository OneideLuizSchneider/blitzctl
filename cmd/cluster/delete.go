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
	deleteExamples = templates.Examples(i18n.T(`
		# minikube
		blitzctl cluster delete minikube --cluster-name=mycluster
		# kind
		blitzctl cluster delete kind --cluster-name=mycluster
	`))

	deleteCmd = &cobra.Command{
		Use:   "delete",
		Short: "Delete a k8s cluster",
		Long: `Delete a k8s cluster
This command will Delete a local k8s cluster
based on cluster the name.`,
		Example: deleteExamples,
		Aliases: []string{"delete", "del", "d", "rm", "remove"},
		Args:    cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			if err := cmd.Help(); err != nil {
				cmd.PrintErrln("❌ Error displaying help:", err)
			}
		},
	}
)

func init() {
	for _, clusterProvider := range provider.GetProviders() {
		deleteCmd.AddCommand(clusterProvider.GetDeleteCommand())
	}
}
