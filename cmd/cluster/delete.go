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
	deleteExamples = templates.Examples(i18n.T(`
		# minikube
		blitzctl clusters delete minikube --cluster-name=mycluster
		# kind
		blitzctl clusters delete kind --cluster-name=mycluster
	`))

	deleteCmd = &cobra.Command{
		Use:   "delete",
		Short: "Delete a k8s cluster",
		Long: `Delete a k8s cluster
This command will Delete a local k8s cluster
based on cluster the name.`,
		Example: deleteExamples,
		Aliases: []string{"c"},
		Args:    cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			if err := cmd.Help(); err != nil {
				cmd.PrintErrln("❌ Error displaying help:", err)
			}
		},
	}
)

func init() {
	deleteCmd.AddCommand(minikube.NewDeleteCmd())
	deleteCmd.AddCommand(kind.NewDeleteCmd())
}
