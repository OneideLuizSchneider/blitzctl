/*
Copyright © 2025 Oneide Luiz Schneider
*/
package context

import (
	"fmt"

	"github.com/OneideLuizSchneider/blitzctl/config"
	"github.com/spf13/cobra"
)

var listContextCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "List all available cluster contexts",
	Long:    `List all managed clusters that can be used as contexts.`,
	Run: func(cmd *cobra.Command, args []string) {
		manager := config.GetManager()
		clusters := manager.ListClusters()
		currentContext := manager.GetCurrentContext()

		if len(clusters) == 0 {
			fmt.Println("No managed clusters available")
			fmt.Println("Create clusters using 'blitzctl create cluster' first")
			return
		}

		fmt.Println("Available Cluster Contexts:")
		fmt.Println("==========================")

		for _, cluster := range clusters {
			marker := "  "
			if currentContext != nil &&
				currentContext.Cluster == cluster.Name &&
				currentContext.Provider == cluster.Provider {
				marker = "* " // Mark current context
			}

			status := "❓"
			switch cluster.Status {
			case "running":
				status = "✅"
			case "stopped":
				status = "⏹️"
			case "deleted":
				status = "❌"
			}

			fmt.Printf("%s%s %s (%s) - %s\n", marker, status, cluster.Name, cluster.Provider, cluster.K8sVersion)
		}

		if currentContext != nil {
			fmt.Printf("\nCurrent: %s (%s)\n", currentContext.Cluster, currentContext.Provider)
		}
	},
}
