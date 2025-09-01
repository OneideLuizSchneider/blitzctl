/*
Copyright © 2025 Oneide Luiz Schneider
*/
package config

import (
	"github.com/spf13/cobra"
	"k8s.io/kubectl/pkg/util/i18n"
	"k8s.io/kubectl/pkg/util/templates"
)

var (
	configExample = templates.Examples(i18n.T(`
		# Get current configuration
		blitzctl config get

		# Set a default value
		blitzctl config set driver docker
		blitzctl config set k8s-version 1.32.0
		blitzctl config set cluster-name my-cluster

		# Get a specific configuration value
		blitzctl config get driver
		blitzctl config get k8s-version

		# List all configuration values
		blitzctl config list

		# Show configuration file location
		blitzctl config view
	`))

	configCmd = &cobra.Command{
		Use:     "config",
		Example: configExample,
		Short:   "Manage blitzctl configuration",
		Long: `Manage blitzctl configuration settings including defaults for cluster creation,
driver preferences, and other settings. Configuration is stored in YAML format
and can be customized per user or per project.`,
		Run: func(cmd *cobra.Command, args []string) {
			if err := cmd.Help(); err != nil {
				cmd.PrintErrln("❌ Error displaying help:", err)
			}
		},
	}
)

// GetConfigCmd returns the config command
func GetConfigCmd() *cobra.Command {
	return configCmd
}

func init() {
	configCmd.AddCommand(getCmd)
	configCmd.AddCommand(setCmd)
	configCmd.AddCommand(listCmd)
	configCmd.AddCommand(viewCmd)
}
