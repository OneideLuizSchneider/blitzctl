/*
Copyright © 2025 Oneide Luiz Schneider
*/
package config

import (
	"fmt"

	"github.com/OneideLuizSchneider/blitzctl/config"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get [key]",
	Short: "Get a configuration value",
	Long: `Get a configuration value. If no key is specified, returns all current configuration.
	
Valid keys:
  - driver
  - k8s-version
  - cluster-name
  - cni
  - helm-version`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		manager := config.GetManager()

		if len(args) == 0 {
			// Display all configuration
			cfg := manager.GetConfig()
			fmt.Println("Current Configuration:")
			fmt.Println("===================")
			fmt.Printf("Driver: %s\n", cfg.Defaults.Driver)
			fmt.Printf("K8s Version: %s\n", cfg.Defaults.K8sVersion)
			fmt.Printf("Cluster Name: %s\n", cfg.Defaults.ClusterName)
			fmt.Printf("CNI: %s\n", cfg.Defaults.CNI)
			fmt.Printf("Helm Version: %s\n", cfg.Defaults.HelmVersion)

			if cfg.CurrentContext != nil {
				fmt.Println("\nCurrent Context:")
				fmt.Printf("Cluster: %s (%s)\n", cfg.CurrentContext.Cluster, cfg.CurrentContext.Provider)
			}

			if len(cfg.Clusters) > 0 {
				fmt.Printf("\nManaged Clusters: %d\n", len(cfg.Clusters))
			}
			return
		}

		key := args[0]
		value, err := manager.GetDefault(key)
		if err != nil {
			fmt.Printf("❌ Error getting configuration: %v\n", err)
			return
		}

		fmt.Printf("%s: %v\n", key, value)
	},
}
