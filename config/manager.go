/*
Copyright © 2025 Oneide Luiz Schneider
*/
package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

const (
	// ConfigFileName is the name of the configuration file
	ConfigFileName = "config"
	// ConfigFileType is the type of the configuration file
	ConfigFileType = "yaml"
	// ConfigDirName is the name of the configuration directory
	ConfigDirName = ".blitzctl"
	// EnvPrefix is the prefix for environment variables
	EnvPrefix = "BLITZCTL"
)

// Manager handles configuration management using Viper
type Manager struct {
	viper  *viper.Viper
	config *Config
}

// NewManager creates a new configuration manager
func NewManager() *Manager {
	v := viper.New()
	v.SetConfigName(ConfigFileName)
	v.SetConfigType(ConfigFileType)
	v.SetEnvPrefix(EnvPrefix)
	v.AutomaticEnv()

	return &Manager{
		viper:  v,
		config: GetDefaultConfig(),
	}
}

// Initialize sets up the configuration manager with search paths and loads config
func (m *Manager) Initialize(configPath string) error {
	// Set custom config path if provided
	if configPath != "" {
		m.viper.SetConfigFile(configPath)
	} else {
		// Add configuration search paths in priority order
		m.addConfigPaths()
	}

	// Load configuration
	if err := m.loadConfig(); err != nil {
		return err
	}

	return nil
}

// addConfigPaths adds configuration search paths
func (m *Manager) addConfigPaths() {
	// Current directory (./.blitzctl)
	if pwd, err := os.Getwd(); err == nil {
		m.viper.AddConfigPath(filepath.Join(pwd, ConfigDirName))
	}

	// User home directory (~/.blitzctl)
	if home, err := os.UserHomeDir(); err == nil {
		m.viper.AddConfigPath(filepath.Join(home, ConfigDirName))
	}

	// System-wide configuration (optional)
	m.viper.AddConfigPath("/etc/blitzctl/")
}

// loadConfig loads the configuration from file or creates default
func (m *Manager) loadConfig() error {
	// Try to read existing config
	if err := m.viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found, use defaults and create directory structure
			if err := m.ensureConfigDir(); err != nil {
				return fmt.Errorf("failed to create config directory: %w", err)
			}
			return nil
		}
		return fmt.Errorf("failed to read config file: %w", err)
	}

	// Unmarshal config into struct
	if err := m.viper.Unmarshal(m.config); err != nil {
		return fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return nil
}

// ensureConfigDir ensures the configuration directory exists
func (m *Manager) ensureConfigDir() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	configDir := filepath.Join(home, ConfigDirName)
	return os.MkdirAll(configDir, 0755)
}

// GetConfig returns the current configuration
func (m *Manager) GetConfig() *Config {
	return m.config
}

// GetDefaults returns the default values
func (m *Manager) GetDefaults() Defaults {
	return m.config.Defaults
}

// SaveConfig saves the current configuration to file
func (m *Manager) SaveConfig() error {
	// Update viper with current config values
	m.viper.Set("defaults", m.config.Defaults)
	m.viper.Set("clusters", m.config.Clusters)
	if m.config.CurrentContext != nil {
		m.viper.Set("current_context", m.config.CurrentContext)
	}

	// Determine config file path
	configFile := m.viper.ConfigFileUsed()
	if configFile == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			return fmt.Errorf("failed to get home directory: %w", err)
		}
		configDir := filepath.Join(home, ConfigDirName)
		if err := os.MkdirAll(configDir, 0755); err != nil {
			return fmt.Errorf("failed to create config directory: %w", err)
		}
		configFile = filepath.Join(configDir, ConfigFileName+"."+ConfigFileType)
	}

	return m.viper.WriteConfigAs(configFile)
}

// SetDefault sets a default configuration value
func (m *Manager) SetDefault(key string, value interface{}) error {
	switch key {
	case "k8s_version", "k8s-version":
		m.config.Defaults.K8sVersion = value.(string)
	case "driver":
		m.config.Defaults.Driver = value.(string)
	case "cluster_name", "cluster-name":
		m.config.Defaults.ClusterName = value.(string)
	case "cni":
		m.config.Defaults.CNI = value.(string)
	case "helm_version", "helm-version":
		m.config.Defaults.HelmVersion = value.(string)
	default:
		return fmt.Errorf("unknown configuration key: %s", key)
	}

	return m.SaveConfig()
}

// GetDefault gets a default configuration value
func (m *Manager) GetDefault(key string) (interface{}, error) {
	switch key {
	case "k8s_version", "k8s-version":
		return m.config.Defaults.K8sVersion, nil
	case "driver":
		return m.config.Defaults.Driver, nil
	case "cluster_name", "cluster-name":
		return m.config.Defaults.ClusterName, nil
	case "cni":
		return m.config.Defaults.CNI, nil
	case "helm_version", "helm-version":
		return m.config.Defaults.HelmVersion, nil
	default:
		return nil, fmt.Errorf("unknown configuration key: %s", key)
	}
}

// AddCluster adds a cluster to the configuration
func (m *Manager) AddCluster(cluster ClusterInfo) error {
	// Check if cluster already exists
	for i, existing := range m.config.Clusters {
		if existing.Name == cluster.Name && existing.Provider == cluster.Provider {
			// Update existing cluster
			m.config.Clusters[i] = cluster
			return m.SaveConfig()
		}
	}

	// Add new cluster
	m.config.Clusters = append(m.config.Clusters, cluster)
	return m.SaveConfig()
}

// RemoveCluster removes a cluster from the configuration
func (m *Manager) RemoveCluster(name, provider string) error {
	for i, cluster := range m.config.Clusters {
		if cluster.Name == name && cluster.Provider == provider {
			// Remove cluster
			m.config.Clusters = append(m.config.Clusters[:i], m.config.Clusters[i+1:]...)

			// Clear current context if it was pointing to this cluster
			if m.config.CurrentContext != nil &&
				m.config.CurrentContext.Cluster == name &&
				m.config.CurrentContext.Provider == provider {
				m.config.CurrentContext = nil
			}

			return m.SaveConfig()
		}
	}
	return fmt.Errorf("cluster %s (%s) not found", name, provider)
}

// GetCluster gets a cluster by name and provider
func (m *Manager) GetCluster(name, provider string) (*ClusterInfo, error) {
	for _, cluster := range m.config.Clusters {
		if cluster.Name == name && cluster.Provider == provider {
			return &cluster, nil
		}
	}
	return nil, fmt.Errorf("cluster %s (%s) not found", name, provider)
}

// ListClusters returns all configured clusters
func (m *Manager) ListClusters() []ClusterInfo {
	return m.config.Clusters
}

// SetCurrentContext sets the current active cluster context
func (m *Manager) SetCurrentContext(clusterName, provider string) error {
	// Verify cluster exists
	if _, err := m.GetCluster(clusterName, provider); err != nil {
		return err
	}

	m.config.CurrentContext = &CurrentContext{
		Cluster:  clusterName,
		Provider: provider,
	}

	return m.SaveConfig()
}

// GetCurrentContext returns the current active cluster context
func (m *Manager) GetCurrentContext() *CurrentContext {
	return m.config.CurrentContext
}

// GetConfigFilePath returns the path to the configuration file being used
func (m *Manager) GetConfigFilePath() string {
	return m.viper.ConfigFileUsed()
}
