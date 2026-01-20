/*
Copyright © 2025 Oneide Luiz Schneider
*/
package create

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
		# Create a minikube cluster (default provider)
		blitzctl create cluster --cluster-name=mycluster --k8s-version=1.32.0 --driver=docker

		# Create a minikube cluster
		blitzctl create cluster --provider minikube --cluster-name=mycluster
	`))

	clusterCmd = &cobra.Command{
		Use:     "cluster",
		Short:   "Create a k8s cluster",
		Long:    `Create a local k8s cluster using the specified provider and configuration.`,
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

			options := &provider.CreateOptions{
				ClusterOptions: provider.ClusterOptions{
					ClusterName: clusterName,
					K8sVersion:  k8sVersion,
				},
			}

			if providerType == provider.Minikube {
				options.ProviderOptions = map[string]interface{}{
					"driver": driver,
					"cni":    cni,
				}
			}

			return clusterProviderInstance.Create(options)
		},
	}

	clusterProvider string
	clusterName     string
	k8sVersion      string
	driver          string
	cni             string
)

func init() {
	clusterCmd.Flags().StringVarP(&clusterProvider, "provider", "p", string(provider.Minikube), i18n.T("Cluster provider (minikube or kind)."))
	clusterCmd.Flags().StringVar(&clusterName, "cluster-name", "", i18n.T("Cluster Name."))
	clusterCmd.Flags().StringVar(&k8sVersion, "k8s-version", config.DefaultK8sVersion, i18n.T("K8s Version."))
	clusterCmd.Flags().StringVar(&driver, "driver", config.DefaultDriver, i18n.T("Driver (minikube only)."))
	clusterCmd.Flags().StringVar(&cni, "cni", config.DefaultCni, i18n.T("CNI (minikube only)."))
}
