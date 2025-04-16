/*
Copyright © 2025 Oneide Luiz Schneider
*/
package cluster

import (
	"fmt"
	"os/exec"

	"github.com/OneideLuizSchneider/blitzctl/config"
	"github.com/spf13/cobra"
	"k8s.io/kubectl/pkg/util/i18n"
	"k8s.io/kubectl/pkg/util/templates"
)

var (
	createExamples = templates.Examples(i18n.T(`
		# Minikube Doc: https://minikube.sigs.k8s.io/docs/start/

		# podman
		blitzctl clusters create minikube --cluster-name=mycluster --k8s-version=1.32.0 --container-runtime=podman
		# docker
		blitzctl clusters create minikube --cluster-name=mycluster --k8s-version=1.32.0 --container-runtime=docker
	`))

	createCmd = &cobra.Command{
		Use:   "create",
		Short: "Create a k8s cluster",
		Long: `Create a k8s cluster
This command will create a local k8s cluster
using the specified driver and configuration.
You can use this command to quickly set up a k8s cluster
for development and testing purposes.`,
		Example: createExamples,
		Aliases: []string{"c"},
		Args:    cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			if err := cmd.Help(); err != nil {
				cmd.PrintErrln("❌ Error displaying help:", err)
			}
		},
	}
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
	fmt.Printf("Debug: ClusterName=%s, K8sVersion=%s, ContainerRuntime=%s\n, Cni=%s\n",
		o.ClusterName, o.K8sVersion, o.ContainerRuntime, o.Cni)

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
		"--driver="+o.ContainerRuntime,
		"--kubernetes-version="+o.K8sVersion,
		"--extra-config=kubelet.max-pods=100",
		"--cni="+o.Cni,
	)
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

func init() {
	createCmd.AddCommand(NewCreateMinikubeCmd())
}
