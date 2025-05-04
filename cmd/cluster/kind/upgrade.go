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

func NewUpgradeKindCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "kind",
		Short:   "Upgrade kind",
		Long:    `Upgrade kind to latest vesion`,
		Example: `blitzctl upgrade cluster kind`,
		Aliases: []string{"kind", "mini", "m"},
		Args:    cobra.NoArgs,
		Run:     upgradeKind,
	}
}

func upgradeKind(cmd *cobra.Command, args []string) {
	_, err := exec.LookPath("kind")
	if err != nil {
		fmt.Println("❌ kind is not installed. Please install kind to use this command.")
		os.Exit(1)
	}
	switch runtime.GOOS {
	case "darwin":
		upgradeCmd := exec.Command("brew", "upgrade", "kind")
		// Set up real-time output streaming
		upgradeCmd.Stdout = os.Stdout
		upgradeCmd.Stderr = os.Stderr
		if err := upgradeCmd.Run(); err != nil {
			fmt.Fprintf(os.Stderr, "❗Error Upgrading kind\n")
			os.Exit(1)
		}
	case "linux":
		upgradeCmd := exec.Command("sh", "-c", `
            curl -Lo /usr/local/bin/kind https://kind.sigs.k8s.io/dl/latest/kind-linux-amd64 &&
            chmod +x /usr/local/bin/kind
        `)
		// Set up real-time output streaming
		upgradeCmd.Stdout = os.Stdout
		upgradeCmd.Stderr = os.Stderr
		if err := upgradeCmd.Run(); err != nil {
			fmt.Fprintf(os.Stderr, "❗Error Upgrading kind\n")
			os.Exit(1)
		}
	case "windows":
		fmt.Println("❌ Running on an unsupported OS")
		os.Exit(1)
	default:
		fmt.Println("❌ Running on an unsupported OS")
		os.Exit(1)
	}
	fmt.Printf("✅ kind Upgraded successfully!")
}
