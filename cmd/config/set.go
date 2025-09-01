/*
Copyright © 2025 Oneide Luiz Schneider
*/
package config

import (
	"fmt"

	"github.com/OneideLuizSchneider/blitzctl/config"
	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:   "set <key> <value>",
	Short: "Set a configuration value",
	Long: `Set a configuration value and save it to the config file.
	
Valid keys:
  - driver (e.g., docker, podman, virtualbox)
  - k8s-version (e.g., 1.32.0, 1.33.4)
  - cluster-name (e.g., my-cluster)
  - cni (e.g., cilium, flannel, calico)
  - helm-version (e.g., v3.18.6)`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		key := args[0]
		value := args[1]

		manager := config.GetManager()

		if err := manager.SetDefault(key, value); err != nil {
			fmt.Printf("❌ Error setting configuration: %v\n", err)
			return
		}

		fmt.Printf("✅ Set %s = %s\n", key, value)

		// Show config file location
		if configPath := manager.GetConfigFilePath(); configPath != "" {
			fmt.Printf("Configuration saved to: %s\n", configPath)
		}
	},
}
