/*
Copyright © 2025 Oneide Luiz Schneider
*/
package tools

import (
	"github.com/spf13/cobra"
	"k8s.io/kubectl/pkg/util/i18n"
	"k8s.io/kubectl/pkg/util/templates"
)

var (
	toolsExample = templates.Examples(i18n.T(`
		# Install Tools like Helm...
		blitzctl tools install`))

	toolsCmd = &cobra.Command{
		Use:     "tools",
		Example: toolsExample,
		Aliases: []string{"t"},
		Short:   "Install tools",
		Long:    `Install All Tools like Helm`,
		Run: func(cmd *cobra.Command, args []string) {
			if err := cmd.Help(); err != nil {
				// Handle the error (e.g., log it or print it)
				cmd.PrintErrln("❌ Error displaying help:", err)
			}
		},
	}
)

func GetToolsmd() *cobra.Command {
	return toolsCmd
}

func init() {
	toolsCmd.AddCommand(installCmd)
}
