/*
Copyright © 2025 Oneide Luiz Schneider
*/
package provider

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/OneideLuizSchneider/blitzctl/config"
	"github.com/spf13/cobra"
	"k8s.io/kubectl/pkg/util/i18n"
)

// K3dProvider implements the ClusterProvider interface for K3d
type K3dProvider struct{}

// NewK3dProvider creates a new K3d provider instance
func NewK3dProvider() ClusterProvider {
	return &K3dProvider{}
}

func (p *K3dProvider) GetProviderType() ProviderType {
	return K3d
}

func (p *K3dProvider) Validate() error {
	_, err := exec.LookPath("k3d")
	if err != nil {
		return fmt.Errorf("❌ K3d is not installed. Please install K3d to use this command")
	}

	_, err = exec.LookPath("docker")
	if err != nil {
		return fmt.Errorf("❌ Docker is not installed. Please install Docker to use this command")
	}

	return nil
}

func (p *K3dProvider) Create(options *CreateOptions) error {
	if err := p.Validate(); err != nil {
		return err
	}

	if options.ClusterName == "" {
		return fmt.Errorf("❌ The Cluster Name is required")
	}

	createCmd := exec.Command(
		"k3d",
		"cluster",
		"create",
		options.ClusterName,
		"--image=rancher/k3s:v"+options.K8sVersion+"-k3s1",
	)

	createCmd.Stdout = os.Stdout
	createCmd.Stderr = os.Stderr

	fmt.Printf("Debug: Running command: %s\n", createCmd.String())
	fmt.Printf("🔄 Creating K3d cluster...\n")

	if err := createCmd.Run(); err != nil {
		return fmt.Errorf("❌ Error creating K3d cluster: %v", err)
	}

	fmt.Printf("✅ K3d cluster '%s' created successfully\n", options.ClusterName)

	// Save cluster information to config
	configManager := config.GetManager()
	clusterInfo := config.ClusterInfo{
		Name:       options.ClusterName,
		Provider:   string(K3d),
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

func (p *K3dProvider) Delete(options *DeleteOptions) error {
	if err := p.Validate(); err != nil {
		return err
	}

	if options.ClusterName == "" {
		return fmt.Errorf("❌ The ClusterName is required")
	}

	deleteCmd := exec.Command(
		"k3d",
		"cluster",
		"delete",
		options.ClusterName,
	)

	deleteCmd.Stdout = os.Stdout
	deleteCmd.Stderr = os.Stderr

	fmt.Printf("🔄 Deleting K3d cluster...\n")

	if err := deleteCmd.Run(); err != nil {
		return fmt.Errorf("❌ Error deleting K3d cluster: %v", err)
	}

	fmt.Printf("✅ K3d cluster '%s' deleted successfully\n", options.ClusterName)

	// Remove cluster information from config
	configManager := config.GetManager()
	if err := configManager.RemoveCluster(options.ClusterName, string(K3d)); err != nil {
		fmt.Printf("⚠️ Warning: Failed to remove cluster from configuration: %v\n", err)
	}

	return nil
}

func (p *K3dProvider) List(options *ListOptions) error {
	if err := p.Validate(); err != nil {
		return err
	}

	getCmd := exec.Command("k3d", "cluster", "list")
	getCmd.Stdout = os.Stdout
	getCmd.Stderr = os.Stderr

	if err := getCmd.Run(); err != nil {
		return fmt.Errorf("❌ Error listing K3d clusters: %v", err)
	}

	return nil
}

func (p *K3dProvider) Upgrade(options *UpgradeOptions) error {
	// K3d doesn't support direct cluster upgrades, would need to recreate
	return fmt.Errorf("❌ K3d doesn't support cluster upgrades. Please delete and recreate the cluster")
}

func (p *K3dProvider) Install(options *InstallOptions) error {
	installCmd := exec.Command("bash", "-c", "curl -s https://raw.githubusercontent.com/k3d-io/k3d/main/install.sh | bash")
	installCmd.Stdout = os.Stdout
	installCmd.Stderr = os.Stderr

	fmt.Printf("🔄 Installing K3d...\n")

	if err := installCmd.Run(); err != nil {
		return fmt.Errorf("❌ Error installing K3d: %v", err)
	}

	fmt.Printf("✅ K3d installed successfully\n")
	return nil
}

// Command builders
func (p *K3dProvider) GetCreateCommand() *cobra.Command {
	var clusterName, k8sVersion string

	cmd := &cobra.Command{
		Use:     "k3d",
		Short:   "Create a k3d cluster",
		Long:    `Create a k3d cluster using the specified configuration (k3d runs on docker).`,
		Example: `blitzctl cluster create k3d --cluster-name=mycluster`,
		Aliases: []string{"k3d", "k3"},
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

func (p *K3dProvider) GetDeleteCommand() *cobra.Command {
	var clusterName string

	cmd := &cobra.Command{
		Use:     "k3d",
		Short:   "Delete a k3d cluster",
		Long:    `Delete a k3d cluster by name.`,
		Example: `blitzctl cluster delete k3d --cluster-name=mycluster`,
		Aliases: []string{"k3d", "k3"},
		RunE: func(cmd *cobra.Command, args []string) error {
			options := &DeleteOptions{
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

func (p *K3dProvider) GetListCommand() *cobra.Command {
	return &cobra.Command{
		Use:     "k3d",
		Short:   "List all k3d clusters",
		Long:    `List all available k3d local clusters`,
		Example: `blitzctl cluster list k3d`,
		Aliases: []string{"k3d", "k3"},
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return p.List(&ListOptions{})
		},
	}
}

func (p *K3dProvider) GetUpgradeCommand() *cobra.Command {
	return &cobra.Command{
		Use:     "k3d",
		Short:   "Upgrade a k3d cluster",
		Long:    `K3d doesn't support direct cluster upgrades. You need to delete and recreate the cluster.`,
		Example: `blitzctl cluster upgrade k3d`,
		Aliases: []string{"k3d", "k3"},
		RunE: func(cmd *cobra.Command, args []string) error {
			return p.Upgrade(&UpgradeOptions{})
		},
	}
}

func (p *K3dProvider) GetInstallCommand() *cobra.Command {
	return &cobra.Command{
		Use:     "k3d",
		Short:   "Install k3d",
		Long:    `Install k3d cluster provider.`,
		Example: `blitzctl cluster install k3d`,
		Aliases: []string{"k3d", "k3"},
		RunE: func(cmd *cobra.Command, args []string) error {
			return p.Install(&InstallOptions{})
		},
	}
}
