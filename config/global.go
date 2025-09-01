/*
Copyright © 2025 Oneide Luiz Schneider
*/
package config

import (
	"sync"
)

var (
	// GlobalManager is the global configuration manager instance
	GlobalManager *Manager
	// once ensures GlobalManager is initialized only once
	once sync.Once
)

// GetManager returns the global configuration manager instance
func GetManager() *Manager {
	once.Do(func() {
		GlobalManager = NewManager()
	})
	return GlobalManager
}

// InitializeGlobalManager initializes the global configuration manager
func InitializeGlobalManager(configPath string) error {
	manager := GetManager()
	return manager.Initialize(configPath)
}
