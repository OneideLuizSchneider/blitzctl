/*
Copyright © 2025 Oneide Luiz Schneider
*/
package cluster

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/OneideLuizSchneider/blitzctl/cmd/cluster/minikube"
	"github.com/spf13/cobra"
)

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
		if err := cmd.Help(); err != nil {
			// Handle the error (e.g., log it or print it)
			cmd.PrintErrln("Error displaying help:", err)
		}
	},
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

// listKindClusters lists all available Kind clusters
func listKindClusters(cmd *cobra.Command, args []string) {
	_, err := exec.LookPath("kind")
	if err != nil {
		fmt.Println("❌ Kind is not installed. Please install Kind to use this command.")
		os.Exit(1)
	}
	getCmd := exec.Command("kind", "get", "clusters")
	output, err := getCmd.Output()
	if err != nil {
		fmt.Fprintf(os.Stderr, "❌ Error getting Kind clusters: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("✅ Available Kind clusters:")
	fmt.Println(string(output))
}

func init() {
	listCmd.AddCommand(minikube.NewlistMinikubeCmd())
	listCmd.AddCommand(listKindCmd)
}
