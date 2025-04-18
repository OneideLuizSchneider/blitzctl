/*
Copyright © 2025 Oneide Luiz Schneider
*/
package cluster

import (
	"github.com/OneideLuizSchneider/blitzctl/cmd/cluster/minikube"
	"github.com/spf13/cobra"
	"k8s.io/kubectl/pkg/util/i18n"
	"k8s.io/kubectl/pkg/util/templates"
)

var (
	createExamples = templates.Examples(i18n.T(`
		# Minikube Doc: https://minikube.sigs.k8s.io/docs/start/

		# podman
		blitzctl clusters create minikube --cluster-name=mycluster --k8s-version=1.32.0 --container-runtime=podman
		# docker
		blitzctl clusters create minikube --cluster-name=mycluster --k8s-version=1.32.0 --container-runtime=docker
	`))

	createCmd = &cobra.Command{
		Use:   "create",
		Short: "Create a k8s cluster",
		Long: `Create a k8s cluster
This command will create a local k8s cluster
using the specified driver and configuration.
You can use this command to quickly set up a k8s cluster
for development and testing purposes.`,
		Example: createExamples,
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
	createCmd.AddCommand(minikube.NewCreateMinikubeCmd())
}
