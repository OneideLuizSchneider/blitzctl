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

// ClusterOptions represents common options for cluster operations
type ClusterOptions struct {
	ClusterName string
	K8sVersion  string
}

// CreateOptions represents options for creating a cluster
type CreateOptions struct {
	ClusterOptions
	// Provider-specific options will be handled via composition or type assertions
	ProviderOptions map[string]interface{}
}

// DeleteOptions represents options for deleting a cluster
type DeleteOptions struct {
	ClusterName string
}

// ListOptions represents options for listing clusters
type ListOptions struct {
	// Currently no specific options needed, but keeping for future extensibility
}

// UpgradeOptions represents options for upgrading a cluster
type UpgradeOptions struct {
	ClusterOptions
	ProviderOptions map[string]interface{}
}

// InstallOptions represents options for installing a cluster
type InstallOptions struct {
	ClusterOptions
	ProviderOptions map[string]interface{}
}

// ClusterProvider defines the interface that all cluster providers must implement
type ClusterProvider interface {
	// GetProviderType returns the type of this provider
	GetProviderType() ProviderType

	// Create creates a new cluster with the given options
	Create(options *CreateOptions) error

	// Delete deletes a cluster with the given options
	Delete(options *DeleteOptions) error

	// List lists all clusters for this provider
	List(options *ListOptions) error

	// Upgrade upgrades a cluster with the given options
	Upgrade(options *UpgradeOptions) error

	// Install installs a cluster with the given options
	Install(options *InstallOptions) error

	// Validate validates that the provider is available and properly configured
	Validate() error

	// GetCreateCommand returns the cobra command for creating clusters with this provider
	GetCreateCommand() *cobra.Command

	// GetDeleteCommand returns the cobra command for deleting clusters with this provider
	GetDeleteCommand() *cobra.Command

	// GetListCommand returns the cobra command for listing clusters with this provider
	GetListCommand() *cobra.Command

	// GetUpgradeCommand returns the cobra command for upgrading clusters with this provider
	GetUpgradeCommand() *cobra.Command

	// GetInstallCommand returns the cobra command for installing clusters with this provider
	GetInstallCommand() *cobra.Command
}
