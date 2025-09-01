/*
Copyright © 2025 Oneide Luiz Schneider
*/
package config

import (
	"fmt"

	"github.com/OneideLuizSchneider/blitzctl/config"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "List all configuration values",
	Long:    `List all configuration values including defaults and managed clusters.`,
	Run: func(cmd *cobra.Command, args []string) {
		manager := config.GetManager()
		cfg := manager.GetConfig()

		fmt.Println("Configuration Values:")
		fmt.Println("====================")
		fmt.Println("\nDefaults:")
		fmt.Printf("  Driver: %s\n", cfg.Defaults.Driver)
		fmt.Printf("  K8s Version: %s\n", cfg.Defaults.K8sVersion)
		fmt.Printf("  Cluster Name: %s\n", cfg.Defaults.ClusterName)
		fmt.Printf("  CNI: %s\n", cfg.Defaults.CNI)
		fmt.Printf("  Helm Version: %s\n", cfg.Defaults.HelmVersion)

		if cfg.CurrentContext != nil {
			fmt.Println("\nCurrent Context:")
			fmt.Printf("  Cluster: %s\n", cfg.CurrentContext.Cluster)
			fmt.Printf("  Provider: %s\n", cfg.CurrentContext.Provider)
		}

		if len(cfg.Clusters) > 0 {
			fmt.Println("\nManaged Clusters:")
			for _, cluster := range cfg.Clusters {
				status := "❓"
				switch cluster.Status {
				case "running":
					status = "✅"
				case "stopped":
					status = "⏹️"
				case "deleted":
					status = "❌"
				}
				fmt.Printf("  %s %s (%s) - %s - Created: %s\n",
					status, cluster.Name, cluster.Provider, cluster.K8sVersion,
					cluster.CreatedAt.Format("2006-01-02 15:04:05"))
			}
		} else {
			fmt.Println("\nManaged Clusters: None")
		}

		// Show config file location
		if configPath := manager.GetConfigFilePath(); configPath != "" {
			fmt.Printf("\nConfig File: %s\n", configPath)
		} else {
			fmt.Println("\nConfig File: Not yet created (using defaults)")
		}
	},
}
