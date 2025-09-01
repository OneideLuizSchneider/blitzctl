package tools

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/OneideLuizSchneider/blitzctl/config"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:     "install",
	Aliases: []string{"i"},
	Short:   "Install Tools like Helm",
	Long:    `Install All Tools like Helm`,
	Run:     InstallTools,
}

func InstallTools(cmd *cobra.Command, args []string) {
	var url, archive, bin string

	_, err := exec.LookPath("helm")
	if err == nil {
		fmt.Println("✅ Helm is already installed.")
		os.Exit(0)
	}

	switch runtime.GOOS {
	case "linux":
		url = "https://get.helm.sh/helm-v" + config.DefaultHelmVersion + "-linux-amd64.tar.gz"
		archive = "helm-linux-amd64.tar.gz"
		bin = "linux-amd64/helm"
	case "darwin":
		url = "https://get.helm.sh/helm-v" + config.DefaultHelmVersion + "-darwin-amd64.tar.gz"
		archive = "helm-darwin-amd64.tar.gz"
		bin = "darwin-amd64/helm"
	default:
		fmt.Println("❌ Running on an unsupported OS")
		os.Exit(1)
	}

	fmt.Println("⬇️ Downloading Helm...")
	if err := exec.Command("curl", "-LO", url).Run(); err != nil {
		fmt.Println("❌ failed to download helm: %w", err)
		Cleanup()
		os.Exit(1)
	}

	fmt.Println("📦 Extracting Helm...")
	if err := exec.Command("tar", "xzvf", archive).Run(); err != nil {
		fmt.Println("❌ failed to extract helm: %w", err)
		Cleanup()
		os.Exit(1)
	}

	fmt.Println("📦Moving Helm binary to /usr/local/bin...")
	if err := exec.Command("mv", bin, "/usr/local/bin/helm").Run(); err != nil {
		fmt.Println("❌ failed to move helm binary: %w", err)
		Cleanup()
		os.Exit(1)
	}

	if err := exec.Command("chmod", "+x", "/usr/local/bin/helm").Run(); err != nil {
		fmt.Println("❌ failed to chmod helm binary: %w", err)
		Cleanup()
		os.Exit(1)
	}

	Cleanup()

	fmt.Println("✅ Helm installed successfully.")
}

func Cleanup() {
	if err := os.Remove("helm-linux-amd64.tar.gz"); err != nil && !os.IsNotExist(err) {
		fmt.Printf("⚠️  Failed to remove helm-linux-amd64.tar.gz: %v\n", err)
	}
	if err := os.Remove("helm-darwin-amd64.tar.gz"); err != nil && !os.IsNotExist(err) {
		fmt.Printf("⚠️  Failed to remove helm-darwin-amd64.tar.gz: %v\n", err)
	}
	if err := os.RemoveAll("linux-amd64"); err != nil {
		fmt.Printf("⚠️  Failed to remove linux-amd64 directory: %v\n", err)
	}
	if err := os.RemoveAll("darwin-amd64"); err != nil {
		fmt.Printf("⚠️  Failed to remove darwin-amd64 directory: %v\n", err)
	}
	fmt.Println("🧹 Cleanup completed.")
}
