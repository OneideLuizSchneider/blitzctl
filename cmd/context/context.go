/*
Copyright © 2025 Oneide Luiz Schneider
*/
package context

import (
	"github.com/spf13/cobra"
	"k8s.io/kubectl/pkg/util/i18n"
	"k8s.io/kubectl/pkg/util/templates"
)

var (
	contextExample = templates.Examples(i18n.T(`
		# Get current active cluster context
		blitzctl context current

		# List all available cluster contexts
		blitzctl context list

		# Switch to a specific cluster context
		blitzctl context use test-cluster kind
	`))

	contextCmd = &cobra.Command{
		Use:     "context",
		Example: contextExample,
		Aliases: []string{"ctx"},
		Short:   "Manage cluster contexts",
		Long: `Manage cluster contexts for switching between different clusters.
Context management allows you to easily switch between different clusters
without having to specify cluster names and providers repeatedly.`,
		Run: func(cmd *cobra.Command, args []string) {
			if err := cmd.Help(); err != nil {
				cmd.PrintErrln("❌ Error displaying help:", err)
			}
		},
	}
)

// GetContextCmd returns the context command
func GetContextCmd() *cobra.Command {
	return contextCmd
}

func init() {
	contextCmd.AddCommand(currentCmd)
	contextCmd.AddCommand(listContextCmd)
	contextCmd.AddCommand(useCmd)
}
