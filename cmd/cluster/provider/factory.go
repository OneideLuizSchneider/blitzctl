/*
Copyright © 2025 Oneide Luiz Schneider
*/
package provider

import (
	"fmt"
	"strings"
)

// ClusterProviderFactory is responsible for creating cluster providers
type ClusterProviderFactory struct {
	providers map[ProviderType]func() ClusterProvider
}

// NewClusterProviderFactory creates a new factory instance
func NewClusterProviderFactory() *ClusterProviderFactory {
	factory := &ClusterProviderFactory{
		providers: make(map[ProviderType]func() ClusterProvider),
	}

	// Register built-in providers
	factory.registerBuiltInProviders()

	return factory
}

// registerBuiltInProviders registers the built-in providers
func (f *ClusterProviderFactory) registerBuiltInProviders() {
	f.RegisterProvider(Kind, NewKindProvider)
	f.RegisterProvider(Minikube, NewMinikubeProvider)
}

// RegisterProvider allows registering custom providers
func (f *ClusterProviderFactory) RegisterProvider(providerType ProviderType, constructor func() ClusterProvider) {
	f.providers[providerType] = constructor
}

// CreateProvider creates a new provider instance
func (f *ClusterProviderFactory) CreateProvider(providerType ProviderType) (ClusterProvider, error) {
	constructor, exists := f.providers[providerType]
	if !exists {
		return nil, fmt.Errorf("❌ unsupported provider: %s. Available providers: %s",
			providerType, f.GetAvailableProviders())
	}

	return constructor(), nil
}

// GetAvailableProviders returns a list of available provider types
func (f *ClusterProviderFactory) GetAvailableProviders() string {
	var providers []string
	for providerType := range f.providers {
		providers = append(providers, string(providerType))
	}
	return strings.Join(providers, ", ")
}

// GetSupportedProviders returns all supported provider types
func (f *ClusterProviderFactory) GetSupportedProviders() []ProviderType {
	var providers []ProviderType
	for providerType := range f.providers {
		providers = append(providers, providerType)
	}
	return providers
}

// Global factory instance
var DefaultFactory = NewClusterProviderFactory()
