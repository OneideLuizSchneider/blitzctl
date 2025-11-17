/*
Copyright © 2025 Oneide Luiz Schneider
*/
package provider

import "github.com/spf13/cobra"

// ProviderType represents the type of cluster provider
type ProviderType string

const (
	Kind     ProviderType = "kind"
	Minikube ProviderType = "minikube"
	K3d      ProviderType = "k3d"
)

type Default struct {
	ClusterName string
}

type ClusterOptions struct {
	ClusterName string
	K8sVersion  string
}

type CreateOptions struct {
	ClusterOptions
	// Provider-specific options will be handled via composition or type assertions
	ProviderOptions map[string]interface{}
}

type ListOptions struct {
	// Currently no specific options needed, but keeping for future extensibility
}

type UpgradeOptions struct {
	ClusterOptions
	ProviderOptions map[string]interface{}
}

type InstallOptions struct {
	ClusterOptions
	ProviderOptions map[string]interface{}
}

// ClusterProvider defines the interface that all cluster providers must implement
type ClusterProvider interface {
	GetProviderType() ProviderType
	Create(options *CreateOptions) error
	Delete(options *Default) error
	List(options *ListOptions) error
	Upgrade(options *UpgradeOptions) error
	Install(options *InstallOptions) error
	Start(options *Default) error
	Stop(options *Default) error
	Validate() error

	GetCreateCommand() *cobra.Command
	GetDeleteCommand() *cobra.Command
	GetListCommand() *cobra.Command
	GetUpgradeCommand() *cobra.Command
	GetInstallCommand() *cobra.Command
	GetStartCommand() *cobra.Command
	GetStopCommand() *cobra.Command
}
