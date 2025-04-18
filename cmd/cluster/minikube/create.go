/*
Copyright © 2025 Oneide Luiz Schneider
*/
package minikube

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/OneideLuizSchneider/blitzctl/config"
	"github.com/spf13/cobra"
	"k8s.io/kubectl/pkg/util/i18n"
)

type CreateOptions struct {
	ClusterName      string
	K8sVersion       string
	ContainerRuntime string
	Cni              string
}

// NewCreateOptions initializes CreateOptions with default values
func NewCreateOptions() *CreateOptions {
	return &CreateOptions{
		ClusterName:      config.DefaultClusterName,
		K8sVersion:       config.DefaultK8sVersion,
		ContainerRuntime: config.DefaultContainerRuntime,
		Cni:              config.DefaultCni,
	}
}

// Complete processes command-line arguments and updates CreateOptions
func (o *CreateOptions) Complete(cmd *cobra.Command, args []string) error {
	return nil
}

// Validate ensures that all required fields in CreateOptions are set
func (o *CreateOptions) Validate() error {
	if o.ContainerRuntime == "" {
		return fmt.Errorf("❌ The container runtime is required")
	}
	return nil
}

// Run executes the create logic
func (o *CreateOptions) Run() error {
	createCmd := exec.Command(
		"minikube",
		"start",
		"--profile="+o.ClusterName,
		"--driver="+o.ContainerRuntime,
		"--kubernetes-version="+o.K8sVersion,
		"--extra-config=kubelet.max-pods=100",
		"--cni="+o.Cni,
	)
	// Set up real-time output streaming
	createCmd.Stdout = os.Stdout
	createCmd.Stderr = os.Stderr

	fmt.Printf("Debug: Running command: %s\n", createCmd.String())
	fmt.Printf("Debug: Command Args: %v\n", createCmd.Args)
	fmt.Printf("🔄 Running...\n")

	// Start and wait for the command to complete
	if err := createCmd.Run(); err != nil {
		return fmt.Errorf("❌ Error creating minikube cluster: %v", err)
	}

	fmt.Printf("✅ Minikube cluster '%s' created successfully with %s and %s\n", o.ClusterName, o.K8sVersion, o.ContainerRuntime)
	fmt.Printf("✅ CNI: %s\n", o.Cni)
	return nil
}

func NewCreateMinikubeCmd() *cobra.Command {
	o := NewCreateOptions()

	cmd := &cobra.Command{
		Use:     "minikube",
		Short:   "Create a minikube cluster",
		Long:    `Create a minikube cluster using the specified driver and configuration.`,
		Example: `blitzctl cluster create minikube --cluster-name=mycluster --k8s-version=1.25.0 --container-runtime=docker`,
		Aliases: []string{"mini", "m"},
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := o.Complete(cmd, args); err != nil {
				return err
			}
			if err := o.Validate(); err != nil {
				return err
			}
			return o.Run()
		},
	}

	// Bind flags to CreateOptions
	cmd.Flags().StringVar(&o.ClusterName, "cluster-name", o.ClusterName, i18n.T("Cluster Name."))
	cmd.Flags().StringVar(&o.K8sVersion, "k8s-version", o.K8sVersion, i18n.T("K8s Version."))
	cmd.Flags().StringVar(&o.ContainerRuntime, "container-runtime", o.ContainerRuntime, i18n.T("Container Runtime."))
	// Check the error returned by MarkFlagRequired
	if err := cmd.MarkFlagRequired("container-runtime"); err != nil {
		// Handle the error (e.g., log it or panic)
		panic(fmt.Sprintf("❌ Failed to mark 'container-runtime' flag as required: %v", err))
	}

	return cmd
}
