/*
Copyright © 2025 Oneide Luiz Schneider
*/
package cluster

import (
	"fmt"
	"os"
	"os/exec"

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

func listMinikubeClusters(cmd *cobra.Command, args []string) {
	// Check if Minikube is installed
	_, err := exec.LookPath("minikube")
	if err != nil {
		fmt.Println("❌ Minikube is not installed. Please install Minikube to use this command.")
		os.Exit(1)
	}
	// List Minikube clusters
	getCmd := exec.Command("minikube", "profile", "list")
	output, err := getCmd.Output()
	if err != nil {
		fmt.Fprintf(os.Stderr, "❌ Error getting Minikube clusters: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Available Minikube clusters:")
	fmt.Println(string(output))
}

func listKindClusters(cmd *cobra.Command, args []string) {
	// Check if Kind is installed
	_, err := exec.LookPath("kind")
	if err != nil {
		fmt.Println("❌ Kind is not installed. Please install Kind to use this command.")
		os.Exit(1)
	}
	// List Kind clusters
	getCmd := exec.Command("kind", "get", "clusters")
	output, err := getCmd.Output()
	if err != nil {
		fmt.Fprintf(os.Stderr, "❌ Error getting Kind clusters: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Available Kind clusters:")
	fmt.Println(string(output))
}

func init() {
	listCmd.AddCommand(listMinikubeCmd)
	listCmd.AddCommand(listKindCmd)
}
