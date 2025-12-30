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
	createExamples = templates.Examples(i18n.T(`
		# Minikube Doc: https://minikube.sigs.k8s.io/docs/start/
		# Kind Doc: https://kind.sigs.k8s.io/docs/user/quick-start/

		# Create a minikube cluster
		blitzctl cluster create minikube --cluster-name=mycluster --k8s-version=1.32.0 --driver=docker
		
		# Create a kind cluster
		blitzctl cluster create kind --cluster-name=mycluster
	`))

	createCmd = &cobra.Command{
		Use:   "create",
		Short: "Create a k8s cluster",
		Long: `Create a k8s cluster
This command will create a local k8s cluster
using the specified provider and configuration.
You can use this command to quickly set up a k8s cluster
for development and testing purposes.`,
		Example: createExamples,
		Aliases: []string{"create", "c", "new"},
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
		createCmd.AddCommand(clusterProvider.GetCreateCommand())
	}
}
