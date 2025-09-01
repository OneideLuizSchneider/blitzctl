/*
Copyright © 2025 Oneide Luiz Schneider
*/
package config

import (
	"fmt"
	"time"

	"github.com/OneideLuizSchneider/blitzctl/config"
	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:    "test-add-cluster",
	Short:  "Add a test cluster to configuration (for testing)",
	Hidden: true, // Hide from help
	Run: func(cmd *cobra.Command, args []string) {
		manager := config.GetManager()

		// Add a test cluster
		testCluster := config.ClusterInfo{
			Name:       "test-cluster",
			Provider:   "kind",
			K8sVersion: "1.32.0",
			Status:     "running",
			CreatedAt:  time.Now(),
			Options:    make(map[string]string),
		}

		if err := manager.AddCluster(testCluster); err != nil {
			fmt.Printf("❌ Error adding test cluster: %v\n", err)
			return
		}

		fmt.Println("✅ Test cluster added successfully")
		fmt.Println("Run 'blitzctl config list' to see the tracked cluster")
	},
}

func init() {
	configCmd.AddCommand(testCmd)
}
