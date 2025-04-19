/*
Copyright © 2025 Oneide Luiz Schneider
*/
package kind

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

func listKindClusters(cmd *cobra.Command, args []string) {
	_, err := exec.LookPath("kind")
	if err != nil {
		fmt.Println("❌ Kind is not installed. Please install Kind to use this command.")
		os.Exit(1)
	}
	getCmd := exec.Command("kind", "get", "clusters")
	// Set up real-time output streaming
	getCmd.Stdout = os.Stdout
	getCmd.Stderr = os.Stderr
	if err := getCmd.Run(); err != nil {
		fmt.Println("❌ Error to get Kind clusters!")
		os.Exit(1)
	}
	// fmt.Println("✅ Available Kind clusters:")
	// fmt.Println(string(output))
}

func NewlistKindCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "kind",
		Short:   "List all kind clusters",
		Long:    `List all available kind local clusters`,
		Example: `blitzctl list clusters <kind>`,
		Aliases: []string{"kind", "k"},
		Args:    cobra.NoArgs,
		Run:     listKindClusters,
	}
}
