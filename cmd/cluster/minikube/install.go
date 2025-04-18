/*
Copyright © 2025 Oneide Luiz Schneider
*/
package minikube

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/spf13/cobra"
)

func NewInstallMinikubeCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "minikube",
		Short:   "Install a minikube cluster",
		Long:    `Install a minikube cluster using the specified driver and configuration.`,
		Example: `blitzctl cluster install minikube`,
		Aliases: []string{"mini", "m"},
		Args:    cobra.NoArgs,
		Run:     installMinikube,
	}
}

func installMinikube(cmd *cobra.Command, args []string) {
	switch runtime.GOOS {
	case "darwin":
		fmt.Println("Installing minikube on macOS...")
		fmt.Println("Please make sure you have Brew installed.")
		fmt.Println("You can install Brew by running the following command:")
		fmt.Println("https://brew.sh/")
		_, err := exec.LookPath("brew")
		if err != nil {
			fmt.Println("❌ Brew is not installed. Please install Brew to use this command.")
			os.Exit(1)
		}
		getCmd := exec.Command("brew", "install", "minikube")
		output, err := getCmd.Output()
		if err != nil {
			fmt.Fprintf(os.Stderr, "❌ Error installing minikube: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Installing minikube...")
		fmt.Println(string(output))
	case "linux":
		fmt.Println("Installing minikube on Linux...")
		fmt.Println("Please make sure you have curl installed.")
		fmt.Println("You can install curl by running the following command:")
		fmt.Println("sudo apt-get install curl")
		getCmd := exec.Command("curl", "-LO", "https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64")
		output, err := getCmd.Output()
		if err != nil {
			fmt.Fprintf(os.Stderr, "❌ Error installing minikube: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Installing minikube...")
		fmt.Println(string(output))
	case "windows":
		fmt.Println("❌ Running on an unsupported OS")
	default:
		fmt.Println("❌ Running on an unsupported OS")
	}
}
