/*
Copyright © 2025 Oneide Luiz Schneider
*/
package install

import (
	"fmt"
	"os"

	"github.com/OneideLuizSchneider/blitzctl/cmd/tools"
	"github.com/spf13/cobra"
	"k8s.io/kubectl/pkg/util/i18n"
	"k8s.io/kubectl/pkg/util/templates"
)

var (
	toolsExample = templates.Examples(i18n.T(`
		# Install Tools like Helm...
		blitzctl install tool --help
		blitzctl install tool --tool=helm
		blitzctl install tool --tool=all
	`))

	toolsCmd = &cobra.Command{
		Use:     "tool",
		Example: toolsExample,
		Aliases: []string{"t"},
		Short:   "Install tools",
		Long:    `Install All Tools like Helm`,
		Run: func(cmd *cobra.Command, args []string) {
			if toolName == "" {
				fmt.Fprintf(os.Stderr, "❗Error: --tool flag is required (helm or all)\n")
				os.Exit(1)
			}
			switch toolName {
			case "helm", "all":
				tools.InstallHelmCmd.Run(cmd, args)
			default:
				fmt.Fprintf(os.Stderr, "❗Error: invalid tool '%s'. Valid options are 'helm' or 'all'\n", toolName)
				os.Exit(1)
			}
			if err := cmd.Help(); err != nil {
				cmd.PrintErrln("❌ Error displaying help:", err)
			}
		},
	}

	toolName string
)

func GetToolsCmd() *cobra.Command {
	return toolsCmd
}

func init() {
	toolsCmd.Flags().StringVarP(&toolName, "tool", "t", "", i18n.T("Tool name (Helm...)."))
}
