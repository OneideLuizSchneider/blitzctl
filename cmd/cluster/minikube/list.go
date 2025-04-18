/*
Copyright © 2025 Oneide Luiz Schneider
*/
package minikube

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

func NewlistMinikubeCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "minikube",
		Short:   "List all minikube clusters",
		Long:    `List all available minikube local clusters`,
		Example: `blitzctl list clusters minikube`,
		Aliases: []string{"minikube", "mini", "m"},
		Args:    cobra.NoArgs,
		Run:     listMinikubeClusters,
	}
}

// listMinikubeClusters lists all available Minikube clusters
func listMinikubeClusters(cmd *cobra.Command, args []string) {
	_, err := exec.LookPath("minikube")
	if err != nil {
		fmt.Println("❌ Minikube is not installed. Please install Minikube to use this command.")
		os.Exit(1)
	}
	getCmd := exec.Command("minikube", "profile", "list")
	// Set up real-time output streaming
	getCmd.Stdout = os.Stdout
	getCmd.Stderr = os.Stderr
	// Start and wait for the command to complete
	if err := getCmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "❗No Minikube clusters\n")
		os.Exit(1)
	}
}
