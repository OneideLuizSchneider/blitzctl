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

type DeleteOptions struct {
	ClusterName string
}

func NewDeleteOptions() *DeleteOptions {
	return &DeleteOptions{
		ClusterName: config.DefaultClusterName,
	}
}

func (o *DeleteOptions) Complete(cmd *cobra.Command, args []string) error {
	return nil
}

func (o *DeleteOptions) Validate() error {
	if o.ClusterName == "" {
		return fmt.Errorf("❌ The ClusterName is required")
	}
	return nil
}

func (o *DeleteOptions) Run() error {
	deleteCmd := exec.Command(
		"minikube",
		"delete",
		"--profile="+o.ClusterName,
	)
	// Set up real-time output streaming
	deleteCmd.Stdout = os.Stdout
	deleteCmd.Stderr = os.Stderr

	fmt.Printf("Debug: Delete command: %s\n", deleteCmd.String())
	fmt.Printf("Debug: Command Args: %v\n", deleteCmd.Args)
	fmt.Printf("🔄 Deleting...\n")

	if err := deleteCmd.Run(); err != nil {
		return fmt.Errorf("❌ Error deleting minikube cluster: %v", err)
	}

	fmt.Printf("✅ Minikube cluster '%s' deleted successfully!", o.ClusterName)
	return nil
}

func NewDeleteCmd() *cobra.Command {
	o := NewDeleteOptions()

	cmd := &cobra.Command{
		Use:     "minikube",
		Short:   "Delete a minikube cluster",
		Long:    `Delete a minikube cluster using the specified driver and configuration.`,
		Example: `blitzctl cluster delete minikube --cluster-name=mycluster`,
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

	// Bind flags to DeleteOptions
	cmd.Flags().StringVar(&o.ClusterName, "cluster-name", o.ClusterName, i18n.T("Cluster Name."))
	// Check the error returned by MarkFlagRequired
	if err := cmd.MarkFlagRequired("cluster-name"); err != nil {
		// Handle the error (e.g., log it or panic)
		panic(fmt.Sprintf("❌ Failed to mark 'cluster-name' flag as required: %v", err))
	}

	return cmd
}
