/*
Copyright © 2025 Oneide Luiz Schneider
*/
package delete

import (
	"fmt"

	"github.com/OneideLuizSchneider/blitzctl/cmd/cluster/provider"
	"github.com/OneideLuizSchneider/blitzctl/config"
	"github.com/spf13/cobra"
	"k8s.io/kubectl/pkg/util/i18n"
	"k8s.io/kubectl/pkg/util/templates"
)

var (
	clusterExamples = templates.Examples(i18n.T(`
		# Delete a minikube cluster (default provider)
		blitzctl delete cluster --cluster-name=mycluster

		# Delete a kind cluster
		blitzctl delete cluster --provider kind --cluster-name=mycluster
	`))

	clusterCmd = &cobra.Command{
		Use:     "cluster",
		Short:   "Delete a k8s cluster",
		Long:    `Delete a local k8s cluster using the specified provider.`,
		Example: clusterExamples,
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			if !cmd.Flags().Changed("cluster-name") {
				clusterName = config.GetManager().GetDefaults().ClusterName
			}
			providerType, err := provider.ParseProvider(clusterProvider)
			if err != nil {
				return err
			}

			clusterProviderInstance, ok := provider.GetProviderByType(providerType)
			if !ok {
				return fmt.Errorf("❌ Unsupported provider: %s (supported: minikube, kind)", clusterProvider)
			}

			return clusterProviderInstance.Delete(&provider.Default{
				ClusterName: clusterName,
			})
		},
	}

	clusterProvider string
	clusterName     string
)

func init() {
	clusterCmd.Flags().StringVarP(&clusterProvider, "provider", "p", string(provider.Minikube), i18n.T("Cluster provider (minikube or kind)."))
	clusterCmd.Flags().StringVar(&clusterName, "cluster-name", "", i18n.T("Cluster Name."))
}
