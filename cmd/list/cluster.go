/*
Copyright © 2025 Oneide Luiz Schneider
*/
package list

import (
	"fmt"

	"github.com/OneideLuizSchneider/blitzctl/cmd/cluster/provider"
	"github.com/spf13/cobra"
	"k8s.io/kubectl/pkg/util/i18n"
	"k8s.io/kubectl/pkg/util/templates"
)

var (
	clusterExamples = templates.Examples(i18n.T(`
		# List minikube clusters (default provider)
		blitzctl list clusters

		# List minikube clusters
		blitzctl list clusters --provider minikube
	`))

	clusterCmd = &cobra.Command{
		Use:     "clusters",
		Aliases: []string{"cluster", "c"},
		Short:   "List k8s clusters",
		Long:    `List local k8s clusters using the specified provider.`,
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

			return clusterProviderInstance.List(&provider.ListOptions{})
		},
	}

	clusterProvider string
)

func init() {
	clusterCmd.Flags().StringVarP(&clusterProvider, "provider", "p", string(provider.Minikube), i18n.T("Cluster provider (minikube or kind)."))
}
