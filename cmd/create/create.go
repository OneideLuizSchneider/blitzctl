/*
Copyright © 2025 Oneide Luiz Schneider
*/
package create

import (
	"github.com/spf13/cobra"
	"k8s.io/kubectl/pkg/util/i18n"
	"k8s.io/kubectl/pkg/util/templates"
)

var (
	createExamples = templates.Examples(i18n.T(`
		# Create a minikube cluster (default provider)
		blitzctl create cluster --cluster-name=mycluster --k8s-version=1.32.0 --driver=docker
		
		# Create a kind cluster
		blitzctl create cluster --provider kind --cluster-name=mycluster
	`))

	createCmd = &cobra.Command{
		Use:     "create",
		Short:   "Create resources",
		Long:    `Create resources such as local Kubernetes clusters.`,
		Example: createExamples,
		Run: func(cmd *cobra.Command, args []string) {
			if err := cmd.Help(); err != nil {
				cmd.PrintErrln("❌ Error displaying help:", err)
			}
		},
	}
)

func GetCreateCmd() *cobra.Command {
	return createCmd
}

func init() {
	createCmd.AddCommand(clusterCmd)
}
