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

func NewUpgradeMinikubeCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "minikube",
		Short:   "Upgrade minikube",
		Long:    `Upgrade minikube to latest vesion`,
		Example: `blitzctl upgrade cluster minikube`,
		Aliases: []string{"minikube", "mini", "m"},
		Args:    cobra.NoArgs,
		Run:     upgradeMinikube,
	}
}

func upgradeMinikube(cmd *cobra.Command, args []string) {
	_, err := exec.LookPath("minikube")
	if err != nil {
		fmt.Println("❌ Minikube is not installed. Please install Minikube to use this command.")
		os.Exit(1)
	}
	checkCmd := exec.Command("minikube", "update-check")
	// Set up real-time output streaming
	checkCmd.Stdout = os.Stdout
	checkCmd.Stderr = os.Stderr
	// Start and wait for the command to complete
	if err := checkCmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "❗Error Checking Minikube\n")
		os.Exit(1)
	}

	switch runtime.GOOS {
	case "darwin":
		upgradeCmd := exec.Command("brew", "upgrade", "minikube")
		// Set up real-time output streaming
		upgradeCmd.Stdout = os.Stdout
		upgradeCmd.Stderr = os.Stderr
		if err := upgradeCmd.Run(); err != nil {
			fmt.Fprintf(os.Stderr, "❗Error Upgrading Minikube\n")
			os.Exit(1)
		}
	case "linux":
		// Step 1: Download the latest Minikube binary
		downloadCmd := exec.Command("curl", "-LO", "https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64")
		downloadCmd.Stdout = os.Stdout
		downloadCmd.Stderr = os.Stderr
		if err := downloadCmd.Run(); err != nil {
			fmt.Fprintf(os.Stderr, "❗Error Downloading Minikube\n")
			os.Exit(1)
		}

		// Step 2: Install the downloaded binary
		installCmd := exec.Command("sudo", "install", "minikube-linux-amd64", "/usr/local/bin/minikube")
		installCmd.Stdout = os.Stdout
		installCmd.Stderr = os.Stderr
		if err := installCmd.Run(); err != nil {
			fmt.Fprintf(os.Stderr, "❗Error Installing Minikube\n")
			os.Exit(1)
		}
	case "windows":
		fmt.Println("❌ Running on an unsupported OS")
		os.Exit(1)
	default:
		fmt.Println("❌ Running on an unsupported OS")
		os.Exit(1)
	}
	fmt.Printf("✅ Minikube Upgraded successfully!")
}
