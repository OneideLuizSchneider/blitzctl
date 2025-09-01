# Cluster Provider Factory Pattern

This document describes the Factory Pattern implementation for cluster providers in `blitzctl`.

## Overview

The Factory Pattern refactoring provides a clean, extensible architecture for supporting multiple cluster providers (kind, minikube, k3d, etc.) with a consistent interface.

## Architecture

### Core Components

1. **Interfaces** (`provider/interfaces.go`)
   - `ClusterProvider`: Main interface that all providers must implement
   - `ProviderType`: Enum for supported provider types
   - Options structs: `CreateOptions`, `DeleteOptions`, `ListOptions`, etc.

2. **Factory** (`provider/factory.go`)
   - `ClusterProviderFactory`: Creates provider instances
   - `DefaultFactory`: Global factory instance with built-in providers registered

3. **Provider Implementations**
   - `KindProvider` (`provider/kind.go`)
   - `MinikubeProvider` (`provider/minikube.go`) 
   - `K3dProvider` (`provider/k3d.go`)

### Benefits

1. **Extensibility**: Adding new providers requires only:
   - Creating a new provider struct that implements `ClusterProvider`
   - Adding the provider type constant
   - Registering it in the factory

2. **Consistency**: All providers implement the same interface, ensuring uniform behavior

3. **Maintainability**: Common logic is centralized, reducing code duplication

4. **Testability**: Easy to mock providers for testing

## Adding a New Provider

To add a new cluster provider (e.g., "docker-desktop"):

### Step 1: Add Provider Type
```go
// In provider/interfaces.go
const (
    Kind           ProviderType = "kind"
    Minikube       ProviderType = "minikube"
    K3d            ProviderType = "k3d"
    DockerDesktop  ProviderType = "docker-desktop"  // Add this
)
```

### Step 2: Create Provider Implementation
```go
// Create provider/docker_desktop.go
type DockerDesktopProvider struct{}

func NewDockerDesktopProvider() ClusterProvider {
    return &DockerDesktopProvider{}
}

func (p *DockerDesktopProvider) GetProviderType() ProviderType {
    return DockerDesktop
}

// Implement all interface methods...
```

### Step 3: Register the Provider
```go
// In provider/factory.go registerBuiltInProviders()
f.RegisterProvider(DockerDesktop, NewDockerDesktopProvider)
```

That's it! The provider will automatically appear in all cluster commands.

## Command Structure

The factory pattern enables dynamic command registration:

```
blitzctl cluster create
├── kind (from KindProvider.GetCreateCommand())
├── minikube (from MinikubeProvider.GetCreateCommand())  
└── k3d (from K3dProvider.GetCreateCommand())
```

## Provider Interface

Each provider must implement:

- `GetProviderType() ProviderType`
- `Create(options *CreateOptions) error`
- `Delete(options *DeleteOptions) error` 
- `List(options *ListOptions) error`
- `Upgrade(options *UpgradeOptions) error`
- `Install(options *InstallOptions) error`
- `Validate() error`
- Command builders: `GetCreateCommand()`, `GetDeleteCommand()`, etc.

## Migration From Old Structure

The old structure:
```
cmd/cluster/
├── kind/
│   ├── create.go
│   ├── delete.go
│   └── list.go
└── minikube/
    ├── create.go
    ├── delete.go
    └── list.go
```

Is now replaced with:
```
cmd/cluster/
├── provider/
│   ├── interfaces.go
│   ├── factory.go
│   ├── kind.go
│   ├── minikube.go
│   └── k3d.go
├── create.go (uses factory)
├── delete.go (uses factory)
└── list.go (uses factory)
```

## Example Usage

```bash
# All these work automatically with any registered provider
blitzctl cluster create kind --cluster-name=my-kind-cluster
blitzctl cluster create minikube --cluster-name=my-mini-cluster --driver=docker
blitzctl cluster create k3d --cluster-name=my-k3d-cluster

blitzctl cluster list kind
blitzctl cluster list minikube  
blitzctl cluster list k3d

blitzctl cluster delete kind --cluster-name=my-kind-cluster
```
