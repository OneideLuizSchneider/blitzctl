/*
Copyright © 2025 Oneide Luiz Schneider
*/
package provider

import (
	"fmt"
	"os"
	"os/exec"

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
	return nil
}

func (p *KindProvider) Delete(options *DeleteOptions) error {
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
	return nil
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
	// Kind doesn't support direct cluster upgrades, would need to recreate
	return fmt.Errorf("❌ Kind doesn't support cluster upgrades. Please delete and recreate the cluster")
}

func (p *KindProvider) Install(options *InstallOptions) error {
	// This would handle Kind installation logic
	return fmt.Errorf("❌ Kind installation not implemented yet")
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
