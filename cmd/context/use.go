/*
Copyright © 2025 Oneide Luiz Schneider
*/
package context

import (
	"fmt"

	"github.com/OneideLuizSchneider/blitzctl/config"
	"github.com/spf13/cobra"
)

var useCmd = &cobra.Command{
	Use:   "use <cluster-name> <provider>",
	Short: "Set the active cluster context",
	Long: `Set the active cluster context to the specified cluster and provider.
The cluster must be already managed by blitzctl.`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		clusterName := args[0]
		provider := args[1]

		manager := config.GetManager()

		if err := manager.SetCurrentContext(clusterName, provider); err != nil {
			fmt.Printf("❌ Error setting context: %v\n", err)
			return
		}

		fmt.Printf("✅ Switched to context: %s (%s)\n", clusterName, provider)
	},
}
