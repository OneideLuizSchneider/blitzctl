/*
Copyright © 2025 Oneide Luiz Schneider
*/
package install

import (
	"fmt"

	"github.com/OneideLuizSchneider/blitzctl/cmd/cluster/provider"
	"github.com/spf13/cobra"
	"k8s.io/kubectl/pkg/util/i18n"
	"k8s.io/kubectl/pkg/util/templates"
)

var (
	clusterExamples = templates.Examples(i18n.T(`
		# Install minikube (default provider)
		blitzctl install cluster

		# Install kind
		blitzctl install cluster --provider kind
	`))

	clusterCmd = &cobra.Command{
		Use:     "cluster",
		Short:   "Install a k8s tool",
		Long:    `Install a local k8s tool using the specified provider.`,
		Example: clusterExamples,
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			providerType, err := provider.ParseProvider(clusterProvider)
			if err != nil {
				return err
			}

			clusterProviderInstance, ok := provider.GetProviderByType(providerType)
			if !ok {
				return fmt.Errorf("❌ Unsupported provider: %s (supported: minikube, kind)", clusterProvider)
			}

			return clusterProviderInstance.Install(&provider.InstallOptions{})
		},
	}

	clusterProvider string
)

func init() {
	clusterCmd.Flags().StringVarP(&clusterProvider, "provider", "p", string(provider.Minikube), i18n.T("Cluster provider (minikube or kind)."))
}
