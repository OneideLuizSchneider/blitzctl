/*
Copyright © 2025 Oneide Luiz Schneider
*/
package config

import (
	"fmt"
	"os"

	"github.com/OneideLuizSchneider/blitzctl/config"
	"github.com/spf13/cobra"
)

var viewCmd = &cobra.Command{
	Use:   "view",
	Short: "View configuration file contents and location",
	Long: `View the current configuration file contents and show where it's located.
This command displays the raw YAML configuration file.`,
	Run: func(cmd *cobra.Command, args []string) {
		manager := config.GetManager()
		configPath := manager.GetConfigFilePath()

		if configPath == "" {
			fmt.Println("No configuration file found. Using default values.")
			fmt.Println("Run 'blitzctl config set <key> <value>' to create a configuration file.")
			return
		}

		fmt.Printf("Configuration file: %s\n", configPath)
		fmt.Println("Content:")
		fmt.Println("========")

		// Read and display file contents
		content, err := os.ReadFile(configPath)
		if err != nil {
			fmt.Printf("❌ Error reading configuration file: %v\n", err)
			return
		}

		fmt.Println(string(content))
	},
}
