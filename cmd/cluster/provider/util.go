/*
Copyright © 2026 Oneide Luiz Schneider
*/
package provider

import (
	"fmt"
	"strings"
)

// ProviderAliases maps supported aliases to their provider type.
var ProviderAliases = map[string]ProviderType{
	"minikube": Minikube,
	"mini":     Minikube,
	"m":        Minikube,
	"kind":     Kind,
	"k":        Kind,
}

// ParseProvider converts user input into a ProviderType.
func ParseProvider(input string) (ProviderType, error) {
	normalized := strings.TrimSpace(strings.ToLower(input))
	if providerType, ok := ProviderAliases[normalized]; ok {
		return providerType, nil
	}

	return "", fmt.Errorf("❌ Unsupported provider: %s (supported: minikube, kind)", input)
}
