/*
Copyright © 2025 Oneide Luiz Schneider
*/
package context

import (
	"fmt"

	"github.com/OneideLuizSchneider/blitzctl/config"
	"github.com/spf13/cobra"
)

var currentCmd = &cobra.Command{
	Use:   "current",
	Short: "Show the current active cluster context",
	Long:  `Display the current active cluster context.`,
	Run: func(cmd *cobra.Command, args []string) {
		manager := config.GetManager()
		currentContext := manager.GetCurrentContext()

		if currentContext == nil {
			fmt.Println("No active cluster context set")
			fmt.Println("Use 'blitzctl context use <cluster> <provider>' to set one")
			return
		}

		fmt.Printf("Current Context: %s (%s)\n", currentContext.Cluster, currentContext.Provider)
	},
}
