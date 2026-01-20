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

	fmt.Println("OS: " + runtime.GOOS)

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
	fmt.Printf("URL: %s\n", url)

	// Get current working directory
	cwd, _ := os.Getwd()
	fmt.Printf("Working directory: %s\n", cwd)

	downloadCmd := exec.Command("curl", "-L", "-f", "-o", archive, url)
	output, err := downloadCmd.CombinedOutput()
	fmt.Printf("Curl output: %s\n", string(output))
	if err != nil {
		fmt.Printf("❌ failed to download helm: %v\n", err)
		Cleanup(archive, bin)
		os.Exit(1)
	}

	// Verify the file was downloaded
	if _, err := os.Stat(archive); os.IsNotExist(err) {
		fmt.Printf("❌ Download failed: %s does not exist\n", archive)
		Cleanup(archive, bin)
		os.Exit(1)
	}

	fileInfo, _ := os.Stat(archive)
	fmt.Printf("✅ Downloaded: %s (%.2f MB)\n", archive, float64(fileInfo.Size())/1024/1024)

	fmt.Println("📦 Extracting Helm...")
	extractCmd := exec.Command("tar", "-xzf", archive)
	if output, err := extractCmd.CombinedOutput(); err != nil {
		fmt.Printf("❌ failed to extract helm: %v\n", err)
		fmt.Printf("Output: %s\n", output)
		Cleanup(archive, bin)
		os.Exit(1)
	}

	fmt.Println("📦 Moving Helm binary to /usr/local/bin...")
	fmt.Println("🔐 This requires administrator privileges. You may be prompted for your password.")
	moveCmd := exec.Command("sudo", "mv", bin, "/usr/local/bin/helm")
	moveCmd.Stdin = os.Stdin
	moveCmd.Stdout = os.Stdout
	moveCmd.Stderr = os.Stderr
	if err := moveCmd.Run(); err != nil {
		fmt.Printf("❌ failed to move helm binary: %v\n", err)
		Cleanup(archive, bin)
		os.Exit(1)
	}

	chmodCmd := exec.Command("sudo", "chmod", "+x", "/usr/local/bin/helm")
	chmodCmd.Stdin = os.Stdin
	chmodCmd.Stdout = os.Stdout
	chmodCmd.Stderr = os.Stderr
	if err := chmodCmd.Run(); err != nil {
		fmt.Printf("❌ failed to chmod helm binary: %v\n", err)
		Cleanup(archive, bin)
		os.Exit(1)
	}

	Cleanup(archive, bin)

	fmt.Println("✅ Helm installed successfully.")
}

func Cleanup(archive, bin string) {
	// Remove the downloaded archive
	if err := os.Remove(archive); err != nil && !os.IsNotExist(err) {
		fmt.Printf("⚠️  Failed to remove %s: %v\n", archive, err)
	}

	// Extract directory name from bin path (e.g., "darwin-amd64/helm" -> "darwin-amd64")
	var extractDir string
	switch runtime.GOOS {
	case "linux":
		extractDir = "linux-amd64"
	case "darwin":
		extractDir = "darwin-amd64"
	}

	if extractDir != "" {
		if err := os.RemoveAll(extractDir); err != nil && !os.IsNotExist(err) {
			fmt.Printf("⚠️  Failed to remove %s directory: %v\n", extractDir, err)
		}
	}

	fmt.Println("🧹 Cleanup completed.")
}
