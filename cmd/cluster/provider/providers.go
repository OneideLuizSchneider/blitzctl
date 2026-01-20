/*
Copyright © 2025 Oneide Luiz Schneider
*/
package provider

// GetProviders returns the built-in cluster providers supported by blitzctl.
func GetProviders() []ClusterProvider {
	return []ClusterProvider{
		NewKindProvider(),
		NewMinikubeProvider(),
	}
}

// GetProviderByType returns the provider matching the given type.
func GetProviderByType(providerType ProviderType) (ClusterProvider, bool) {
	for _, p := range GetProviders() {
		if p.GetProviderType() == providerType {
			return p, true
		}
	}
	return nil, false
}
