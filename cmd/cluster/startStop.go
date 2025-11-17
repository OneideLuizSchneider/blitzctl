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
	startStopExamples = templates.Examples(i18n.T(`
		# Minikube Doc: https://minikube.sigs.k8s.io/docs/start/

		# minikube
		blitzctl cluster start minikube --cluster-name <profile-name>
		blitzctl cluster stop minikube --cluster-name <profile-name>
	`))

	startCmd = &cobra.Command{
		Use:   "start",
		Short: "Start a k8s engine",
		Long: `Start a k8s Cluster engine
This command will start a local k8s engine
using the specified driver and configuration.`,
		Example: startStopExamples,
		Aliases: []string{},
		Args:    cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			if err := cmd.Help(); err != nil {
				cmd.PrintErrln("❌ Error displaying help:", err)
			}
		},
	}

	stopCmd = &cobra.Command{
		Use:   "stop",
		Short: "Stop a k8s engine",
		Long: `Stop a k8s Cluster engine
This command will stop a local k8s engine
using the specified driver and configuration.`,
		Example: startStopExamples,
		Aliases: []string{},
		Args:    cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			if err := cmd.Help(); err != nil {
				cmd.PrintErrln("❌ Error displaying help:", err)
			}
		},
	}
)

func init() {
	factory := provider.DefaultFactory
	// Register provider commands dynamically
	for _, providerType := range factory.GetSupportedProviders() {
		if clusterProvider, err := factory.CreateProvider(providerType); err == nil {
			startCmd.AddCommand(clusterProvider.GetStartCommand())
			stopCmd.AddCommand(clusterProvider.GetStopCommand())
		}
	}
}
