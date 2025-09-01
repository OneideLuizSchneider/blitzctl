# Viper Configuration Integration

This document describes the Viper integration added to blitzctl for configuration and cluster state management.

## Features Added

### 1. Configuration Management
- **Configuration File**: `~/.blitzctl/config.yaml` (user-global) or `./.blitzctl/config.yaml` (project-specific)
- **Environment Variables**: Support for `BLITZCTL_*` environment variables
- **Command Line Flags**: `--config` flag to specify custom config file location
- **Priority Order**: CLI flags > Environment variables > Config file > Defaults

### 2. Default Values Management
```bash
# Set default values
blitzctl config set driver docker
blitzctl config set k8s-version 1.32.0
blitzctl config set cluster-name my-cluster
blitzctl config set cni cilium
blitzctl config set helm-version v3.18.6

# Get current configuration
blitzctl config get
blitzctl config get driver

# List all configuration
blitzctl config list

# View configuration file
blitzctl config view
```

### 3. Cluster State Tracking
- **Automatic Tracking**: Clusters are automatically tracked when created/deleted
- **Cluster Information**: Name, provider, K8s version, status, creation time, driver, CNI, options
- **Persistent Storage**: Cluster information persists across blitzctl sessions

### 4. Context Management
```bash
# List available cluster contexts
blitzctl context list

# Show current active context
blitzctl context current

# Switch to a specific cluster context
blitzctl context use test-cluster kind
```

## Configuration Structure

```yaml
defaults:
  k8s_version: "1.33.4"
  driver: "podman"
  cluster_name: "blitz-cluster1"
  cni: "cilium"
  helm_version: "v3.18.6"

clusters:
  - name: "test-cluster"
    provider: "kind"
    k8s_version: "1.32.0"
    status: "running"
    created_at: "2025-09-01T10:22:59Z"
    driver: "docker"
    cni: "cilium"
    options:
      custom_option: "value"

current_context:
  cluster: "test-cluster"
  provider: "kind"
```

## Integration Points

### 1. Provider Integration
- **Minikube Provider**: Uses config defaults for driver and CNI
- **Kind Provider**: Uses config defaults and tracks cluster state
- **Automatic State Management**: Create/delete operations update cluster registry

### 2. Global Configuration Manager
- **Singleton Pattern**: Global configuration manager instance
- **Thread-Safe**: Concurrent access protection
- **Lazy Initialization**: Config loaded on first access

### 3. Command Structure
```
blitzctl
├── config
│   ├── get [key]
│   ├── set <key> <value>
│   ├── list
│   └── view
├── context
│   ├── current
│   ├── list
│   └── use <cluster> <provider>
└── cluster (updated to use config)
    ├── create (saves cluster info)
    └── delete (removes cluster info)
```

## Benefits

1. **User Experience**: 
   - Persistent preferences
   - No need to repeat common options
   - Easy context switching

2. **State Management**:
   - Track created clusters
   - Persist cluster metadata
   - Context awareness

3. **Flexibility**:
   - Multiple configuration locations
   - Environment variable support
   - Profile-based configurations (ready for future)

4. **Maintainability**:
   - Centralized configuration logic
   - Clean separation of concerns
   - Backward compatibility maintained

## Future Enhancements

1. **Configuration Profiles**: Support for multiple named configurations
2. **Import/Export**: Configuration backup and sharing
3. **Validation**: Enhanced configuration validation
4. **Auto-completion**: Context-aware command completion
5. **Cluster Health**: Integration with cluster status checking
