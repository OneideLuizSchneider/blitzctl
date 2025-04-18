/*
Copyright © 2025 Oneide Luiz Schneider
*/
package cluster

import (
	"github.com/spf13/cobra"
	"k8s.io/kubectl/pkg/util/i18n"
	"k8s.io/kubectl/pkg/util/templates"
)

var (
	listExample = templates.Examples(i18n.T(`
		# Kind Doc: https://kind.sigs.k8s.io/docs/user/quick-start/
		# Minikube Doc: https://minikube.sigs.k8s.io/docs/start/

		# List Local Minikube clusters
		blitzctl clusters list minikube

		# List Local Minikube clusters
		blitzctl clusters list kind
		
		# Install a Minikube cluster
		blitzctl clusters install minikube
		
		# Install a Kind cluster
		blitzctl clusters install kind`))

	clusterCmd = &cobra.Command{
		Use:     "cluster",
		Example: listExample,
		Aliases: []string{"clusters", "cl"},
		Short:   "A brief description of your command",
		Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		Run: func(cmd *cobra.Command, args []string) {
			if err := cmd.Help(); err != nil {
				// Handle the error (e.g., log it or print it)
				cmd.PrintErrln("❌ Error displaying help:", err)
			}
		},
	}
)

func GetClusterCmd() *cobra.Command {
	return clusterCmd
}

func init() {
	clusterCmd.AddCommand(listCmd)
	clusterCmd.AddCommand(installCmd)
	clusterCmd.AddCommand(createCmd)
}
