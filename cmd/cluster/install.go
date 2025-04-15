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

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install a k8s cluster",
	Long: `Install a k8s cluster
This command will install a k8s cluster
using the specified driver and configuration.
You can use this command to quickly set up a k8s cluster
for development and testing purposes.`,
	Example: `blitzctl cluster install minikube`,
	Aliases: []string{"i"},
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var installMinikubeCmd = &cobra.Command{
	Use:     "minikube",
	Short:   "Install a minikube cluster",
	Long:    `Install a minikube cluster using the specified driver and configuration.`,
	Example: `blitzctl cluster install minikube`,
	Aliases: []string{"mini", "m"},
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// Check if Minikube is installed
		_, err := exec.LookPath("brew")
		if err != nil {
			fmt.Println("Brew is not installed. Please install Brew to use this command.")
			os.Exit(1)
		}
		getCmd := exec.Command("brew", "install", "minikube")
		output, err := getCmd.Output()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error installing minikube: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Installing minikube...")
		fmt.Println(string(output))
	},
}

var installKindCmd = &cobra.Command{
	Use:     "kind",
	Short:   "Install a kind cluster",
	Long:    `Install a kind cluster using the specified driver and configuration.`,
	Example: `blitzctl cluster install kind`,
	Aliases: []string{"k"},
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// Check if Minikube is installed
		_, err := exec.LookPath("brew")
		if err != nil {
			fmt.Println("Brew is not installed. Please install Brew to use this command.")
			os.Exit(1)
		}
		getCmd := exec.Command("brew", "install", "kind")
		output, err := getCmd.Output()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error installing kind: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Installing kind...")
		fmt.Println(string(output))
	},
}

func init() {
	installCmd.AddCommand(installMinikubeCmd)
	installCmd.AddCommand(installKindCmd)
}
