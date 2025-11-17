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

	// Get configuration manager for defaults
	configManager := config.GetManager()
	defaults := configManager.GetDefaults()

	// Extract minikube-specific options from ProviderOptions with config defaults
	driver := defaults.Driver
	cni := defaults.CNI
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

	// Save cluster information to config
	clusterInfo := config.ClusterInfo{
		Name:       options.ClusterName,
		Provider:   string(Minikube),
		K8sVersion: options.K8sVersion,
		Status:     "running",
		CreatedAt:  time.Now(),
		Driver:     driver,
		CNI:        cni,
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

func (p *MinikubeProvider) Delete(options *Default) error {
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

	// Remove cluster information from config
	configManager := config.GetManager()
	if err := configManager.RemoveCluster(options.ClusterName, string(Minikube)); err != nil {
		fmt.Printf("⚠️ Warning: Failed to remove cluster from configuration: %v\n", err)
	}

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
			linkCmd := exec.Command("brew", "link", "--overwrite", "minikube")
			linkCmd.Stdout = os.Stdout
			linkCmd.Stderr = os.Stderr
			if err := linkCmd.Run(); err != nil {
				fmt.Fprintf(os.Stderr, "❗Error Linking Minikube\n")
				os.Exit(1)
			}
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

	return nil
}

func (p *MinikubeProvider) Install(options *InstallOptions) error {
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

	return nil
}

func (p *MinikubeProvider) Stop(options *Default) error {
	if err := p.Validate(); err != nil {
		return err
	}

	if options.ClusterName == "" {
		return fmt.Errorf("❌ The Cluster Name is required")
	}

	createCmd := exec.Command(
		"minikube",
		"stop",
		"--profile="+options.ClusterName,
	)

	createCmd.Stdout = os.Stdout
	createCmd.Stderr = os.Stderr

	fmt.Printf("Debug: Running command: %s\n", createCmd.String())
	fmt.Printf("Debug: Command Args: %v\n", createCmd.Args)
	fmt.Printf("🔄 Running...\n")

	if err := createCmd.Run(); err != nil {
		return fmt.Errorf("❌ Error creating minikube cluster: %v", err)
	}

	fmt.Printf("✅ Minikube cluster '%s' created successfully\n", options.ClusterName)

	return nil
}

func (p *MinikubeProvider) Start(options *Default) error {
	if err := p.Validate(); err != nil {
		return err
	}

	if options.ClusterName == "" {
		return fmt.Errorf("❌ The Cluster Name is required")
	}

	createCmd := exec.Command(
		"minikube",
		"start",
		"--profile="+options.ClusterName,
	)

	createCmd.Stdout = os.Stdout
	createCmd.Stderr = os.Stderr

	fmt.Printf("Debug: Running command: %s\n", createCmd.String())
	fmt.Printf("Debug: Command Args: %v\n", createCmd.Args)
	fmt.Printf("🔄 Running...\n")

	if err := createCmd.Run(); err != nil {
		return fmt.Errorf("❌ Error creating minikube cluster: %v", err)
	}

	fmt.Printf("✅ Minikube cluster '%s' created successfully\n", options.ClusterName)

	return nil
}

// Command builders
func (p *MinikubeProvider) GetCreateCommand() *cobra.Command {
	var clusterName, k8sVersion, driver, cni string

	cmd := &cobra.Command{
		Use:     "minikube",
		Short:   "Create a minikube cluster",
		Long:    `Create a minikube cluster using the specified driver and configuration.`,
		Example: `blitzctl cluster create minikube --cluster-name=mycluster --k8s-version=1.32.0 --driver=docker --cni=cilium`,
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
		Example: `blitzctl cluster upgrade minikube`,
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

func (p *MinikubeProvider) GetStartCommand() *cobra.Command {
	var clusterName string

	cmd := &cobra.Command{
		Use:     "minikube",
		Short:   "Start a minikube cluster",
		Long:    `Start a minikube cluster.`,
		Example: `blitzctl cluster start minikube --cluster-name <cluster-name>`,
		Aliases: []string{"mini", "m"},
		RunE: func(cmd *cobra.Command, args []string) error {
			options := &Default{
				ClusterName: clusterName,
			}
			return p.Start(options)
		},
	}
	cmd.Flags().StringVar(&clusterName, "cluster-name", config.DefaultClusterName, i18n.T("Cluster Name."))
	if err := cmd.MarkFlagRequired("cluster-name"); err != nil {
		panic(fmt.Sprintf("❌ Failed to mark 'cluster-name' flag as required: %v", err))
	}

	return cmd
}

func (p *MinikubeProvider) GetStopCommand() *cobra.Command {
	var clusterName string

	cmd := &cobra.Command{
		Use:     "minikube",
		Short:   "Stop a minikube cluster",
		Long:    `Stop a minikube cluster.`,
		Example: `blitzctl cluster stop minikube --cluster-name <cluster-name>`,
		Aliases: []string{"mini", "m"},
		RunE: func(cmd *cobra.Command, args []string) error {
			options := &Default{
				ClusterName: clusterName,
			}
			return p.Stop(options)
		},
	}
	cmd.Flags().StringVar(&clusterName, "cluster-name", config.DefaultClusterName, i18n.T("Cluster Name."))
	if err := cmd.MarkFlagRequired("cluster-name"); err != nil {
		panic(fmt.Sprintf("❌ Failed to mark 'cluster-name' flag as required: %v", err))
	}

	return cmd
}
