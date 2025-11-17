/*
Copyright © 2025 Oneide Luiz Schneider
*/
package provider

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/OneideLuizSchneider/blitzctl/config"
	"github.com/spf13/cobra"
	"k8s.io/kubectl/pkg/util/i18n"
)

// KindProvider implements the ClusterProvider interface for Kind
type KindProvider struct{}

// NewKindProvider creates a new Kind provider instance
func NewKindProvider() ClusterProvider {
	return &KindProvider{}
}

func (p *KindProvider) GetProviderType() ProviderType {
	return Kind
}

func (p *KindProvider) Validate() error {
	_, err := exec.LookPath("kind")
	if err != nil {
		return fmt.Errorf("❌ Kind is not installed. Please install Kind to use this command")
	}

	_, err = exec.LookPath("docker")
	if err != nil {
		return fmt.Errorf("❌ Docker is not installed. Please install Docker to use this command")
	}

	return nil
}

func (p *KindProvider) Create(options *CreateOptions) error {
	if err := p.Validate(); err != nil {
		return err
	}

	if options.ClusterName == "" {
		return fmt.Errorf("❌ The Cluster Name is required")
	}

	createCmd := exec.Command(
		"kind",
		"create",
		"cluster",
		"--image=kindest/node:v"+options.K8sVersion,
		"--name="+options.ClusterName,
	)

	// Set up real-time output streaming
	createCmd.Stdout = os.Stdout
	createCmd.Stderr = os.Stderr

	fmt.Printf("Debug: Running command: %s\n", createCmd.String())
	fmt.Printf("Debug: Command Args: %v\n", createCmd.Args)
	fmt.Printf("🔄 Running...\n")

	if err := createCmd.Run(); err != nil {
		return fmt.Errorf("❌ Error creating Kind cluster: %v", err)
	}

	fmt.Printf("✅ Kind cluster '%s' created successfully\n", options.ClusterName)

	// Save cluster information to config
	configManager := config.GetManager()
	clusterInfo := config.ClusterInfo{
		Name:       options.ClusterName,
		Provider:   string(Kind),
		K8sVersion: options.K8sVersion,
		Status:     "running",
		CreatedAt:  time.Now(),
		Options:    make(map[string]string),
	}

	// Add provider-specific options to the cluster info
	if options.ProviderOptions != nil {
		for k, v := range options.ProviderOptions {
			if str, ok := v.(string); ok {
				clusterInfo.Options[k] = str
			}
		}
	}

	if err := configManager.AddCluster(clusterInfo); err != nil {
		fmt.Printf("⚠️ Warning: Failed to save cluster information: %v\n", err)
	}

	return nil
}

func (p *KindProvider) Delete(options *Default) error {
	if err := p.Validate(); err != nil {
		return err
	}

	if options.ClusterName == "" {
		return fmt.Errorf("❌ The ClusterName is required")
	}

	deleteCmd := exec.Command(
		"kind",
		"delete",
		"cluster",
		"--name="+options.ClusterName,
	)

	deleteCmd.Stdout = os.Stdout
	deleteCmd.Stderr = os.Stderr

	fmt.Printf("Debug: Delete command: %s\n", deleteCmd.String())
	fmt.Printf("Debug: Command Args: %v\n", deleteCmd.Args)
	fmt.Printf("🔄 Deleting...\n")

	if err := deleteCmd.Run(); err != nil {
		return fmt.Errorf("❌ Error deleting Kind cluster: %v", err)
	}

	fmt.Printf("✅ Kind cluster '%s' deleted successfully\n", options.ClusterName)

	// Remove cluster information from config
	configManager := config.GetManager()
	if err := configManager.RemoveCluster(options.ClusterName, string(Kind)); err != nil {
		fmt.Printf("⚠️ Warning: Failed to remove cluster from configuration: %v\n", err)
	}

	return nil
}

func (p *KindProvider) Start(options *Default) error {
	return fmt.Errorf("❌ kind doesn't support cluster start. Please delete and recreate the cluster")
}

func (p *KindProvider) Stop(options *Default) error {
	return fmt.Errorf("❌ kind doesn't support cluster stop. Please delete and recreate the cluster")
}

func (p *KindProvider) GetStartCommand() *cobra.Command {
	return &cobra.Command{
		Use:     "kind",
		Short:   "Start a kind cluster",
		Long:    `Kind doesn't support direct cluster upgrades. You need to delete and recreate the cluster.`,
		Example: `blitzctl cluster start kind`,
		Aliases: []string{"kind", "k"},
		RunE: func(cmd *cobra.Command, args []string) error {
			return p.Start(&Default{})
		},
	}
}

func (p *KindProvider) GetStopCommand() *cobra.Command {
	return &cobra.Command{
		Use:     "kind",
		Short:   "Stop a kind cluster",
		Long:    `Kind doesn't support direct cluster upgrades. You need to delete and recreate the cluster.`,
		Example: `blitzctl cluster stop kind`,
		Aliases: []string{"kind", "k"},
		RunE: func(cmd *cobra.Command, args []string) error {
			return p.Stop(&Default{})
		},
	}
}

func (p *KindProvider) List(options *ListOptions) error {
	if err := p.Validate(); err != nil {
		return err
	}

	getCmd := exec.Command("kind", "get", "clusters")
	getCmd.Stdout = os.Stdout
	getCmd.Stderr = os.Stderr

	if err := getCmd.Run(); err != nil {
		return fmt.Errorf("❌ Error listing Kind clusters: %v", err)
	}

	return nil
}

func (p *KindProvider) Upgrade(options *UpgradeOptions) error {
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

	return nil
}

func (p *KindProvider) Install(options *InstallOptions) error {
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

	return nil
}

// Command builders - these create cobra commands that use the provider
func (p *KindProvider) GetCreateCommand() *cobra.Command {
	var clusterName, k8sVersion string

	cmd := &cobra.Command{
		Use:     "kind",
		Short:   "Create a kind cluster",
		Long:    `Create a kind cluster using the specified configuration (kind only works with docker driver).`,
		Example: `blitzctl cluster create kind --cluster-name=mycluster`,
		Aliases: []string{"kind", "k"},
		RunE: func(cmd *cobra.Command, args []string) error {
			options := &CreateOptions{
				ClusterOptions: ClusterOptions{
					ClusterName: clusterName,
					K8sVersion:  k8sVersion,
				},
			}
			return p.Create(options)
		},
	}

	cmd.Flags().StringVar(&clusterName, "cluster-name", config.DefaultClusterName, i18n.T("Cluster Name."))
	cmd.Flags().StringVar(&k8sVersion, "k8s-version", config.DefaultK8sVersion, i18n.T("K8s Version."))

	if err := cmd.MarkFlagRequired("cluster-name"); err != nil {
		panic(fmt.Sprintf("❌ Failed to mark 'cluster-name' flag as required: %v", err))
	}

	return cmd
}

func (p *KindProvider) GetDeleteCommand() *cobra.Command {
	var clusterName string

	cmd := &cobra.Command{
		Use:     "kind",
		Short:   "Delete a kind cluster",
		Long:    `Delete a kind cluster by name.`,
		Example: `blitzctl cluster delete kind --cluster-name=mycluster`,
		Aliases: []string{"kind", "k"},
		RunE: func(cmd *cobra.Command, args []string) error {
			options := &Default{
				ClusterName: clusterName,
			}
			return p.Delete(options)
		},
	}

	cmd.Flags().StringVar(&clusterName, "cluster-name", config.DefaultClusterName, i18n.T("Cluster Name."))

	if err := cmd.MarkFlagRequired("cluster-name"); err != nil {
		panic(fmt.Sprintf("❌ Failed to mark 'cluster-name' flag as required: %v", err))
	}

	return cmd
}

func (p *KindProvider) GetListCommand() *cobra.Command {
	return &cobra.Command{
		Use:     "kind",
		Short:   "List all kind clusters",
		Long:    `List all available kind local clusters`,
		Example: `blitzctl cluster list kind`,
		Aliases: []string{"kind", "k"},
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return p.List(&ListOptions{})
		},
	}
}

func (p *KindProvider) GetUpgradeCommand() *cobra.Command {
	return &cobra.Command{
		Use:     "kind",
		Short:   "Upgrade a kind cluster",
		Long:    `Kind doesn't support direct cluster upgrades. You need to delete and recreate the cluster.`,
		Example: `blitzctl cluster upgrade kind`,
		Aliases: []string{"kind", "k"},
		RunE: func(cmd *cobra.Command, args []string) error {
			return p.Upgrade(&UpgradeOptions{})
		},
	}
}

func (p *KindProvider) GetInstallCommand() *cobra.Command {
	return &cobra.Command{
		Use:     "kind",
		Short:   "Install kind",
		Long:    `Install kind cluster provider.`,
		Example: `blitzctl cluster install kind`,
		Aliases: []string{"kind", "k"},
		RunE: func(cmd *cobra.Command, args []string) error {
			return p.Install(&InstallOptions{})
		},
	}
}
