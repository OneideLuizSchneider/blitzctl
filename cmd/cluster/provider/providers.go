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
