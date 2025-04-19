/*
Copyright © 2025 Oneide Luiz Schneider
*/
package kind

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/spf13/cobra"
)

func installKind(cmd *cobra.Command, args []string) {
	_, err := exec.LookPath("docker")
	if err != nil {
		fmt.Println("❌ Docker is not installed. Please install Docker to use this command.")
		os.Exit(1)
	}

	switch runtime.GOOS {
	case "darwin":
		fmt.Println("Installing kind on macOS...")
		fmt.Println("Please make sure you have Brew installed.")
		fmt.Println("You can install Brew by running the following command:")
		fmt.Println("https://brew.sh/")
		// Check if Minikube is installed
		_, err := exec.LookPath("brew")
		if err != nil {
			fmt.Println("❌ Brew is not installed. Please install Brew to use this command.")
			os.Exit(1)
		}
		getCmd := exec.Command("brew", "install", "kind")
		output, err := getCmd.Output()
		if err != nil {
			fmt.Fprintf(os.Stderr, "❌ Error installing kind: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Installing kind...")
		fmt.Println(string(output))
	case "linux":
		fmt.Println("Installing kind on Linux...")
		fmt.Println("Please make sure you have curl installed.")
		fmt.Println("You can install curl by running the following command:")
		fmt.Println("sudo apt-get install curl")
		getCmd := exec.Command("curl", "-Lo", "/usr/local/bin/kind", "https://kind.sigs.k8s.io/dl/v0.11.0/kind-$(uname)-amd64")
		output, err := getCmd.Output()
		if err != nil {
			fmt.Fprintf(os.Stderr, "❌ Error installing kind: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(string(output))
	case "windows":
		fmt.Println("❌ Running on an unsupported OS")
	default:
		fmt.Println("❌ Running on an unsupported OS")
	}
}

func NewInstallKindCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "kind",
		Short:   "Install a kind cluster",
		Long:    `Install a kind cluster using the specified driver and configuration.`,
		Example: `blitzctl cluster install kind`,
		Aliases: []string{"k"},
		Args:    cobra.NoArgs,
		Run:     installKind,
	}
}
