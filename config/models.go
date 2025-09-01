/*
Copyright © 2025 Oneide Luiz Schneider
*/
package config

import (
	"time"
)

// Config represents the complete blitzctl configuration
type Config struct {
	Defaults       Defaults        `yaml:"defaults" mapstructure:"defaults"`
	Clusters       []ClusterInfo   `yaml:"clusters" mapstructure:"clusters"`
	CurrentContext *CurrentContext `yaml:"current_context,omitempty" mapstructure:"current_context"`
}

// Defaults holds the default configuration values
type Defaults struct {
	K8sVersion  string `yaml:"k8s_version" mapstructure:"k8s_version"`
	Driver      string `yaml:"driver" mapstructure:"driver"`
	ClusterName string `yaml:"cluster_name" mapstructure:"cluster_name"`
	CNI         string `yaml:"cni" mapstructure:"cni"`
	HelmVersion string `yaml:"helm_version" mapstructure:"helm_version"`
}

// ClusterInfo represents information about a managed cluster
type ClusterInfo struct {
	Name       string            `yaml:"name" mapstructure:"name"`
	Provider   string            `yaml:"provider" mapstructure:"provider"`
	K8sVersion string            `yaml:"k8s_version" mapstructure:"k8s_version"`
	Status     string            `yaml:"status" mapstructure:"status"`
	CreatedAt  time.Time         `yaml:"created_at" mapstructure:"created_at"`
	Driver     string            `yaml:"driver,omitempty" mapstructure:"driver"`
	CNI        string            `yaml:"cni,omitempty" mapstructure:"cni"`
	Options    map[string]string `yaml:"options,omitempty" mapstructure:"options"`
}

// CurrentContext represents the current active cluster context
type CurrentContext struct {
	Cluster  string `yaml:"cluster" mapstructure:"cluster"`
	Provider string `yaml:"provider" mapstructure:"provider"`
}

// GetDefaultConfig returns a Config struct with default values
func GetDefaultConfig() *Config {
	return &Config{
		Defaults: Defaults{
			K8sVersion:  DefaultK8sVersion,
			Driver:      DefaultDriver,
			ClusterName: DefaultClusterName,
			CNI:         DefaultCni,
			HelmVersion: DefaultHelmVersion,
		},
		Clusters:       []ClusterInfo{},
		CurrentContext: nil,
	}
}
