/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cluster

import (
	"github.com/spf13/cobra"
	"k8s.io/kubectl/pkg/util/i18n"
	"k8s.io/kubectl/pkg/util/templates"
)

var (
	listExample = templates.Examples(i18n.T(`
		# List Local Minikube clusters
		blitzctl clusters list minikube

		# List Local Minikube clusters
		blitzctl clusters list kind`))
)

var clusterCmd = &cobra.Command{
	Use:     "cluster",
	Aliases: []string{"clusters", "cl"},
	Short:   "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list [subcommand]",
	Short: "List all available local environments",
	Long: `List all available local environments
This command will display all the local environments
that are currently available on your system.
You can use this command to quickly see which environments
are available for use and their current status.`,
	Example: listExample,
	Aliases: []string{"ls"},
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var listMinikubeCmd = &cobra.Command{
	Use:     "minikube",
	Short:   "List all minikube clusters",
	Long:    `List all available minikube local clusters`,
	Example: `blitzctl list clusters minikube`,
	Aliases: []string{"mini", "m"},
	Args:    cobra.NoArgs,
	Run:     listMinikubeClusters,
}

var listKindCmd = &cobra.Command{
	Use:     "kind",
	Short:   "List all kind clusters",
	Long:    `List all available kind local clusters`,
	Example: `blitzctl list clusters <kind>`,
	Aliases: []string{"kind", "k"},
	Args:    cobra.NoArgs,
	Run:     listKindClusters,
}

func GetClusterCmd() *cobra.Command {
	return clusterCmd
}

func init() {
	listCmd.AddCommand(listMinikubeCmd)
	listCmd.AddCommand(listKindCmd)

	clusterCmd.AddCommand(listCmd)

}
