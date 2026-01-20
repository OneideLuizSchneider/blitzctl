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

The factory pattern enables provider-backed commands:

```
blitzctl create cluster --provider <provider>
blitzctl delete cluster --provider <provider>
blitzctl list cluster --provider <provider>
blitzctl install cluster --provider <provider>
blitzctl upgrade cluster --provider <provider>
blitzctl start cluster --provider <provider>
blitzctl stop cluster --provider <provider>
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
└── provider/
    ├── interfaces.go
    ├── factory.go
    ├── kind.go
    ├── minikube.go
    └── k3d.go
cmd/create/
├── create.go
└── cluster.go
cmd/delete/
├── delete.go
└── cluster.go
cmd/list/
├── list.go
└── cluster.go
```

## Example Usage

```bash
# All these work automatically with any registered provider
blitzctl create cluster --provider kind --cluster-name=my-kind-cluster
blitzctl create cluster --provider minikube --cluster-name=my-mini-cluster --driver=docker
blitzctl create cluster --provider k3d --cluster-name=my-k3d-cluster

blitzctl list cluster --provider kind
blitzctl list cluster --provider minikube  
blitzctl list cluster --provider k3d

blitzctl delete cluster --provider kind --cluster-name=my-kind-cluster
```
