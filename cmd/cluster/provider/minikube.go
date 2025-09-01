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

// MinikubeProvider implements the ClusterProvider interface for Minikube
type MinikubeProvider struct{}

// NewMinikubeProvider creates a new Minikube provider instance
func NewMinikubeProvider() ClusterProvider {
	return &MinikubeProvider{}
}

func (p *MinikubeProvider) GetProviderType() ProviderType {
	return Minikube
}

func (p *MinikubeProvider) Validate() error {
	_, err := exec.LookPath("minikube")
	if err != nil {
		return fmt.Errorf("❌ Minikube is not installed. Please install Minikube to use this command")
	}
	return nil
}

func (p *MinikubeProvider) Create(options *CreateOptions) error {
	if err := p.Validate(); err != nil {
		return err
	}

	if options.ClusterName == "" {
		return fmt.Errorf("❌ The Cluster Name is required")
	}

	// Extract minikube-specific options from ProviderOptions
	driver := config.DefaultDriver
	cni := config.DefaultCni
	if options.ProviderOptions != nil {
		if d, ok := options.ProviderOptions["driver"].(string); ok && d != "" {
			driver = d
		}
		if c, ok := options.ProviderOptions["cni"].(string); ok && c != "" {
			cni = c
		}
	}

	if driver == "" {
		return fmt.Errorf("❌ The Driver is required")
	}

	createCmd := exec.Command(
		"minikube",
		"start",
		"--profile="+options.ClusterName,
		"--driver="+driver,
		"--kubernetes-version="+options.K8sVersion,
		"--extra-config=kubelet.max-pods=100",
		"--cni="+cni,
	)

	createCmd.Stdout = os.Stdout
	createCmd.Stderr = os.Stderr

	fmt.Printf("Debug: Running command: %s\n", createCmd.String())
	fmt.Printf("Debug: Command Args: %v\n", createCmd.Args)
	fmt.Printf("🔄 Running...\n")

	if err := createCmd.Run(); err != nil {
		return fmt.Errorf("❌ Error creating minikube cluster: %v", err)
	}

	fmt.Printf("✅ Minikube cluster '%s' created successfully with %s and %s\n", options.ClusterName, options.K8sVersion, driver)
	fmt.Printf("🔌 CNI: %s\n", cni)
	return nil
}

func (p *MinikubeProvider) Delete(options *DeleteOptions) error {
	if err := p.Validate(); err != nil {
		return err
	}

	if options.ClusterName == "" {
		return fmt.Errorf("❌ The ClusterName is required")
	}

	deleteCmd := exec.Command(
		"minikube",
		"delete",
		"--profile="+options.ClusterName,
	)

	deleteCmd.Stdout = os.Stdout
	deleteCmd.Stderr = os.Stderr

	fmt.Printf("Debug: Delete command: %s\n", deleteCmd.String())
	fmt.Printf("Debug: Command Args: %v\n", deleteCmd.Args)
	fmt.Printf("🔄 Deleting...\n")

	if err := deleteCmd.Run(); err != nil {
		return fmt.Errorf("❌ Error deleting minikube cluster: %v", err)
	}

	fmt.Printf("✅ Minikube cluster '%s' deleted successfully\n", options.ClusterName)
	return nil
}

func (p *MinikubeProvider) List(options *ListOptions) error {
	if err := p.Validate(); err != nil {
		return err
	}

	getCmd := exec.Command("minikube", "profile", "list")
	getCmd.Stdout = os.Stdout
	getCmd.Stderr = os.Stderr

	if err := getCmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "❗No Minikube clusters\n")
		return nil // Don't return error for empty list
	}

	return nil
}

func (p *MinikubeProvider) Upgrade(options *UpgradeOptions) error {
	if err := p.Validate(); err != nil {
		return err
	}

	if options.ClusterName == "" {
		return fmt.Errorf("❌ The Cluster Name is required")
	}

	upgradeCmd := exec.Command(
		"minikube",
		"start",
		"--profile="+options.ClusterName,
		"--kubernetes-version="+options.K8sVersion,
	)

	upgradeCmd.Stdout = os.Stdout
	upgradeCmd.Stderr = os.Stderr

	fmt.Printf("🔄 Upgrading minikube cluster...\n")

	if err := upgradeCmd.Run(); err != nil {
		return fmt.Errorf("❌ Error upgrading minikube cluster: %v", err)
	}

	fmt.Printf("✅ Minikube cluster '%s' upgraded successfully to %s\n", options.ClusterName, options.K8sVersion)
	return nil
}

func (p *MinikubeProvider) Install(options *InstallOptions) error {
	// This would handle Minikube installation logic
	return fmt.Errorf("❌ Minikube installation not implemented yet")
}

// Command builders
func (p *MinikubeProvider) GetCreateCommand() *cobra.Command {
	var clusterName, k8sVersion, driver, cni string

	cmd := &cobra.Command{
		Use:     "minikube",
		Short:   "Create a minikube cluster",
		Long:    `Create a minikube cluster using the specified driver and configuration.`,
		Example: `blitzctl cluster create minikube --cluster-name=mycluster --k8s-version=1.32.0 --driver=docker`,
		Aliases: []string{"mini", "m"},
		RunE: func(cmd *cobra.Command, args []string) error {
			options := &CreateOptions{
				ClusterOptions: ClusterOptions{
					ClusterName: clusterName,
					K8sVersion:  k8sVersion,
				},
				ProviderOptions: map[string]interface{}{
					"driver": driver,
					"cni":    cni,
				},
			}
			return p.Create(options)
		},
	}

	cmd.Flags().StringVar(&clusterName, "cluster-name", config.DefaultClusterName, i18n.T("Cluster Name."))
	cmd.Flags().StringVar(&k8sVersion, "k8s-version", config.DefaultK8sVersion, i18n.T("K8s Version."))
	cmd.Flags().StringVar(&driver, "driver", config.DefaultDriver, i18n.T("Driver."))
	cmd.Flags().StringVar(&cni, "cni", config.DefaultCni, i18n.T("CNI."))

	if err := cmd.MarkFlagRequired("driver"); err != nil {
		panic(fmt.Sprintf("❌ Failed to mark 'driver' flag as required: %v", err))
	}

	return cmd
}

func (p *MinikubeProvider) GetDeleteCommand() *cobra.Command {
	var clusterName string

	cmd := &cobra.Command{
		Use:     "minikube",
		Short:   "Delete a minikube cluster",
		Long:    `Delete a minikube cluster by profile name.`,
		Example: `blitzctl cluster delete minikube --cluster-name=mycluster`,
		Aliases: []string{"mini", "m"},
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

func (p *MinikubeProvider) GetListCommand() *cobra.Command {
	return &cobra.Command{
		Use:     "minikube",
		Short:   "List all minikube clusters",
		Long:    `List all available minikube local clusters`,
		Example: `blitzctl cluster list minikube`,
		Aliases: []string{"minikube", "mini", "m"},
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return p.List(&ListOptions{})
		},
	}
}

func (p *MinikubeProvider) GetUpgradeCommand() *cobra.Command {
	var clusterName, k8sVersion string

	cmd := &cobra.Command{
		Use:     "minikube",
		Short:   "Upgrade a minikube cluster",
		Long:    `Upgrade a minikube cluster to a new Kubernetes version.`,
		Example: `blitzctl cluster upgrade minikube --cluster-name=mycluster --k8s-version=1.33.1`,
		Aliases: []string{"mini", "m"},
		RunE: func(cmd *cobra.Command, args []string) error {
			options := &UpgradeOptions{
				ClusterOptions: ClusterOptions{
					ClusterName: clusterName,
					K8sVersion:  k8sVersion,
				},
			}
			return p.Upgrade(options)
		},
	}

	cmd.Flags().StringVar(&clusterName, "cluster-name", config.DefaultClusterName, i18n.T("Cluster Name."))
	cmd.Flags().StringVar(&k8sVersion, "k8s-version", config.DefaultK8sVersion, i18n.T("K8s Version."))

	if err := cmd.MarkFlagRequired("cluster-name"); err != nil {
		panic(fmt.Sprintf("❌ Failed to mark 'cluster-name' flag as required: %v", err))
	}

	return cmd
}

func (p *MinikubeProvider) GetInstallCommand() *cobra.Command {
	return &cobra.Command{
		Use:     "minikube",
		Short:   "Install minikube",
		Long:    `Install minikube cluster provider.`,
		Example: `blitzctl cluster install minikube`,
		Aliases: []string{"mini", "m"},
		RunE: func(cmd *cobra.Command, args []string) error {
			return p.Install(&InstallOptions{})
		},
	}
}
