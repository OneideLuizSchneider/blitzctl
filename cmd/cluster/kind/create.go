/*
Copyright © 2025 Oneide Luiz Schneider
*/
package kind

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/OneideLuizSchneider/blitzctl/config"
	"github.com/spf13/cobra"
	"k8s.io/kubectl/pkg/util/i18n"
)

type CreateOptions struct {
	ClusterName string
	K8sVersion  string
}

// NewCreateOptions initializes CreateOptions with default values
func NewCreateOptions() *CreateOptions {
	return &CreateOptions{
		ClusterName: config.DefaultClusterName,
		K8sVersion:  config.DefaultK8sVersion,
	}
}

// Complete processes command-line arguments and updates CreateOptions
func (o *CreateOptions) Complete(cmd *cobra.Command, args []string) error {
	return nil
}

// Validate ensures that all required fields in CreateOptions are set
func (o *CreateOptions) Validate() error {
	_, err := exec.LookPath("docker")
	if err != nil {
		fmt.Println("❌ Docker is not installed. Please install Docker to use this command.")
		os.Exit(1)
	}

	if o.ClusterName == "" {
		return fmt.Errorf("❌ The Cluster Name is required")
	}
	return nil
}

// Run executes the create logic
func (o *CreateOptions) Run() error {
	createCmd := exec.Command(
		"kind",
		"create",
		"cluster",
		"--image=kindest/node:v"+o.K8sVersion,
		"--name="+o.ClusterName,
	)
	// Set up real-time output streaming
	createCmd.Stdout = os.Stdout
	createCmd.Stderr = os.Stderr

	fmt.Printf("Debug: Running command: %s\n", createCmd.String())
	fmt.Printf("Debug: Command Args: %v\n", createCmd.Args)
	fmt.Printf("🔄 Running...\n")

	// Start and wait for the command to complete
	if err := createCmd.Run(); err != nil {
		return fmt.Errorf("❌ Error creating Kind cluster: %v", err)
	}

	fmt.Printf("✅ Kind cluster '%s' created successfully\n", o.ClusterName)
	return nil
}

func NewCreatkindCmd() *cobra.Command {
	o := NewCreateOptions()

	cmd := &cobra.Command{
		Use:     "kind",
		Short:   "Create a kind cluster",
		Long:    `Create a kind cluster using the specified driver and configuration (kind only works with docker driver).`,
		Example: `blitzctl cluster create kind --cluster-name=mycluster`,
		Aliases: []string{"kind", "k"},
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
	// Check the error returned by MarkFlagRequired
	if err := cmd.MarkFlagRequired("cluster-name"); err != nil {
		// Handle the error (e.g., log it or panic)
		panic(fmt.Sprintf("❌ Failed to mark 'cluster-name' flag as required: %v", err))
	}

	return cmd
}
