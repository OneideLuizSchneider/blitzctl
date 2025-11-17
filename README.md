# blitzctl

`blitzctl` is a CLI tool for managing local Kubernetes environments. It simplifies the creation, deletion, upgrading, and management of Kubernetes clusters using tools like `Minikube` and `Kind`.

Currently supports `macOS` and `Linux`.

## Key Features

- 🚀 **Multi-Provider Support**: Kind and Minikube
- ⚙️ **Smart Configuration**: Powered by Viper with file, environment, and flag support
- 🔄 **Context Switching**: Easy switching between different cluster environments
- 📊 **Cluster State Tracking**: Automatic tracking of created clusters and their metadata
- 🎯 **Flexible Defaults**: Customizable defaults for all cluster operations
- 📁 **Project-Specific Config**: Support for both global and project-specific configurations

## Install

- Quick install (latest):

```sh
curl -fsSL https://raw.githubusercontent.com/OneideLuizSchneider/blitzctl/main/scripts/blitzctl.sh | sh -
```

- Pin a specific version:

```sh
BLITZCTL_VERSION=v0.0.3 \
  curl -fsSL https://raw.githubusercontent.com/OneideLuizSchneider/blitzctl/main/scripts/blitzctl.sh | sh -
```

The script detects your OS/arch (macOS/Linux, amd64/arm64), downloads the matching `blitzctl` binary from GitHub Releases, and installs it into `/usr/local/bin` (or `~/.local/bin` if not writable).

## Quick Start

```sh
# 1. Set your preferences (optional - uses defaults otherwise)
blitzctl config set driver docker
blitzctl config set k8s-version 1.33.4

# 2. Create a cluster
blitzctl cluster create kind --cluster-name my-dev-cluster

# 3. View your managed clusters
blitzctl config list

# 4. Switch context to your cluster
blitzctl context use my-dev-cluster kind

# 5. Create another cluster for testing
blitzctl cluster create minikube --cluster-name my-test-cluster

# 6. Switch between clusters easily
blitzctl context list               # See all available contexts
blitzctl context use my-test-cluster minikube
blitzctl context current          # Check current context
```

## Usage

#### General Command Structure

```sh
blitzctl <command> <subcommand> [flags]
```

#### Commands

##### Cluster Commands

- `create`: Create a Kubernetes cluster.
- `delete`: Delete a Kubernetes cluster.
- `list`: List all available clusters.
- `install`: Install tools like Minikube or Kind.
- `upgrade`: Upgrade tools like Minikube or Kind to their latest versions.
- `start` `start`: Only available for `minikube`
  - It'll `start` or `stop` a cluster

##### Configuration Commands

- `config get`: View current configuration values.
- `config set <key> <value>`: Set configuration defaults.
- `config list`: List all configuration and managed clusters.
- `config view`: View raw configuration file contents.

##### Context Commands

- `context current`: Show current active cluster context.
- `context list`: List all available cluster contexts.
- `context use <cluster> <provider>`: Switch to a specific cluster context.

##### Tools Commands

- `tools install`: Install additional tools such as Helm.

#### Flags

- `--cluster-name`: Specify the name of the cluster.
- `--k8s-version`: Specify the Kubernetes version.
- `--driver`: Specify the driver (e.g., Docker, Podman, virtualbox, parallels, hyperkit, vmware, qemu2, vfkit).
  - For `kind`, only Docker.

## Configuration

`blitzctl` uses [Viper](https://github.com/spf13/viper) for configuration management, providing flexible configuration through files, environment variables, and command-line flags.

### Configuration File Locations (Priority Order)

1. `--config` flag (highest priority)
2. `BLITZCTL_CONFIG` environment variable
3. `./.blitzctl/config.yaml` (project-specific)
4. `~/.blitzctl/config.yaml` (user global)
5. Built-in defaults (fallback)

### Default Values

Default configurations are defined in `config/defaults.go`:

- **Kubernetes Version**: `1.33.4`
- **Driver**: `podman`
- **Cluster Name**: `blitz-cluster1`
- **CNI Plugin**: `cilium`
- **Helm Version**: `v3.18.6`

### Configuration Management

```sh
# View current configuration
blitzctl config get

# Set default values
blitzctl config set driver docker
blitzctl config set k8s-version 1.33.4
blitzctl config set cluster-name my-default-cluster
blitzctl config set cni flannel

# Get specific configuration values
blitzctl config get driver
blitzctl config get k8s-version

# List all configuration and managed clusters
blitzctl config list

# View raw configuration file and location
blitzctl config view
```

### Cluster State Management

`blitzctl` automatically tracks created clusters and their metadata:

```sh
# Clusters are automatically tracked when created
blitzctl cluster create minikube --cluster-name prod-cluster --driver docker

# View tracked clusters
blitzctl config list

# Clusters are automatically removed when deleted
blitzctl cluster delete minikube --cluster-name prod-cluster
```

### Context Management

Switch between different clusters easily:

```sh
# List available cluster contexts
blitzctl context list

# Show current active context
blitzctl context current

# Switch to a specific cluster context
blitzctl context use my-cluster minikube
blitzctl context use dev-cluster kind
```

### Environment Variables

All configuration can be overridden with environment variables:

```sh
export BLITZCTL_DRIVER=docker
export BLITZCTL_K8S_VERSION=1.33.4
export BLITZCTL_CNI=cilium
blitzctl cluster create minikube
```

### Custom Configuration File

```sh
# Use a custom configuration file
blitzctl --config ./my-project/.blitzctl/config.yaml cluster create kind

# Or set via environment
export BLITZCTL_CONFIG=./my-project/.blitzctl/config.yaml
blitzctl cluster create kind
```

### Configuration File Format

The configuration file uses YAML format and includes three main sections:

```yaml
# Default values for cluster operations
defaults:
  k8s_version: "1.34.4"
  driver: "docker"
  cluster_name: "my-default-cluster"
  cni: "cilium"
  helm_version: "v3.18.6"

# Tracked clusters with metadata
clusters:
  - name: "prod-cluster"
    provider: "minikube"
    k8s_version: "1.34.4"
    status: "running"
    created_at: "2025-09-01T10:30:00Z"
    driver: "docker"
    cni: "cilium"
  - name: "dev-cluster"
    provider: "kind"
    k8s_version: "1.31.0"
    status: "running"
    created_at: "2025-09-01T11:15:00Z"

# Current active cluster context
current_context:
  cluster: "dev-cluster"
  provider: "kind"
```

---

## Examples

### Configuration Management Examples

#### Setting Up Your Preferences

```sh
# Set your preferred defaults
blitzctl config set driver docker
blitzctl config set k8s-version 1.34.4
blitzctl config set cluster-name my-default-cluster
blitzctl config set cni flannel

# View your configuration
blitzctl config get
# Output:
# Current Configuration:
# ===================
# Driver: docker
# K8s Version: 1.34.4
# Cluster Name: my-default-cluster
# CNI: flannel
# Helm Version: v3.18.6
```

#### Working with Different Configuration Files

```sh
# Use project-specific configuration
mkdir my-project && cd my-project
mkdir .blitzctl
blitzctl config set driver podman
blitzctl config set cluster-name project-cluster

# Use a custom configuration file
blitzctl --config ./custom-config.yaml config set driver virtualbox

# Check which config file is being used
blitzctl config view
# Shows: Configuration file: /path/to/config.yaml
```

#### Environment Variable Override

```sh
# Override configuration with environment variables
export BLITZCTL_DRIVER=containerd
export BLITZCTL_CNI=calico
blitzctl config get driver  # Shows: containerd
```

### Cluster Management with State Tracking

#### Creating and Tracking Clusters

```sh
# Create clusters with your configured defaults
blitzctl cluster create minikube --cluster-name prod-cluster
blitzctl cluster create kind --cluster-name dev-cluster

# View all managed clusters
blitzctl config list
# Output:
# Managed Clusters:
#   ✅ prod-cluster (minikube) - 1.34.4 - Created: 2025-09-01 10:30:00
#   ✅ dev-cluster (kind) - 1.34.4 - Created: 2025-09-01 10:35:00
```

### Context Management Examples

#### Switching Between Clusters

```sh
# List available contexts
blitzctl context list
# Output:
# Available Cluster Contexts:
# ==========================
#   ✅ prod-cluster (minikube) - 1.34.4
#   ✅ dev-cluster (kind) - 1.34.4

# Switch to development cluster
blitzctl context use dev-cluster kind

# Check current context
blitzctl context current
# Output: Current Context: dev-cluster (kind)

# Switch to production cluster
blitzctl context use prod-cluster minikube
```

#### Development Workflow Example

```sh
# 1. Set up your development environment preferences
blitzctl config set driver docker
blitzctl config set k8s-version 1.31.0
blitzctl config set cni cilium

# 2. Create different clusters for different purposes
blitzctl cluster create kind --cluster-name frontend-dev
blitzctl cluster create minikube --cluster-name backend-dev --driver podman

# 3. Switch between environments as needed
blitzctl context use frontend-dev kind     # Work on frontend
blitzctl context use backend-dev minikube  # Work on backend

# 4. Check what you're currently working on
blitzctl context current

# 5. See all your environments
blitzctl context list
```

### Team Configuration Example

```sh
# Share team configuration via version control
# Create .blitzctl/config.yaml in your project root
blitzctl --config ./.blitzctl/config.yaml config set driver docker
blitzctl --config ./.blitzctl/config.yaml config set k8s-version 1.31.0
blitzctl --config ./.blitzctl/config.yaml config set cni cilium

# Team members can use the same configuration
git add .blitzctl/config.yaml
git commit -m "Add team blitzctl configuration"

# Other team members can now use:
blitzctl cluster create kind  # Uses team configuration
```

### Complete Workflow Example

```sh
# 1. Initial setup - configure your preferences
blitzctl config set driver docker
blitzctl config set k8s-version 1.34.4

# 2. Create multiple environments
blitzctl cluster create kind --cluster-name frontend-dev
blitzctl cluster create minikube --cluster-name backend-dev --driver podman  

# 3. View all your environments
blitzctl config list
# Shows all tracked clusters with their status and creation time

# 4. Work with different environments
blitzctl context use frontend-dev kind
# ... do frontend development work ...

blitzctl context use backend-dev minikube  
# ... do backend development work ...

# 5. Check what you're currently working on
blitzctl context current

# 6. Clean up when done
blitzctl cluster delete kind --cluster-name frontend-dev
# Automatically removes from tracking

# 7. View remaining environments
blitzctl context list
```

---

## Basic Cluster Operations

#### Create a Minikube Cluster

Create a Kubernetes cluster using Minikube with a specific Kubernetes version and container runtime:

```sh
# podman
blitzctl cluster create minikube --cluster-name=mycluster --k8s-version=1.33.1 --driver=podman

# docker
blitzctl cluster create minikube --cluster-name=mycluster --k8s-version=1.33.1 --driver=docker
```

#### Delete a Cluster

Delete a Kubernetes cluster:

```sh
# kind
blitzctl cluster delete kind --cluster-name=mycluster
# minikube
blitzctl cluster delete minikube --cluster-name=mycluster
```

#### List All Minikube Clusters

List all Kubernetes clusters managed by Minikube:

```sh
blitzctl cluster list minikube
```

#### Install Minikube

Install Minikube on your system:

```sh
blitzctl cluster install minikube
```

#### Upgrade Minikube

Upgrade Minikube to the latest version:

```sh
blitzctl cluster upgrade minikube
```

#### Upgrade Kind

Upgrade Kind to the latest version:

```sh
blitzctl cluster upgrade kind
```

#### Create a Kind Cluster with Default Settings

Create a Kubernetes cluster using Kind with default configurations:

```sh
blitzctl cluster create kind
```

#### Delete a Cluster with Custom Name

Delete a cluster with a custom name:

```sh
blitzctl cluster delete kind --cluster-name=custom-cluster
```

#### Debugging a Cluster Deletion

Run a cluster deletion command with debug output enabled:

```sh
blitzctl cluster delete kind --cluster-name=mycluster --debug
```

#### Install Helm

Install Helm on your system:

```sh
blitzctl tools install
```

---

## Tools and Libraries

`blitzctl` is built using the following tools and libraries:

#### [Go](https://golang.org/)
- The project is written in Go, a statically typed, compiled programming language designed for simplicity and performance.
- Go provides excellent support for building CLI tools with its standard library and ecosystem.

#### [Cobra](https://github.com/spf13/cobra)
- Cobra is a powerful library for creating modern CLI applications in Go.
- It provides features like command hierarchies, flag parsing, and built-in help generation.
- `blitzctl` uses Cobra to define commands like `create`, `delete`, `list`, `install`, and `upgrade`.

#### [Kind](https://kind.sigs.k8s.io/)
- Kind (Kubernetes IN Docker) is a tool for running local Kubernetes clusters using Docker containers.
- `blitzctl` integrates with Kind to create, manage, and upgrade Kubernetes clusters.

#### [Minikube](https://minikube.sigs.k8s.io/docs/)
- Minikube is a tool that lets you run Kubernetes locally.
- `blitzctl` supports Minikube for creating, managing, and upgrading clusters with various drivers and configurations.

#### [Helm](https://helm.sh/)
- Helm is a package manager for Kubernetes.
- `blitzctl` can install Helm for you using the `blitzctl tools install` command.

#### [Kubectl Utilities](https://kubernetes.io/docs/reference/kubectl/)
- The project uses utilities from the Kubernetes `kubectl` package for handling Kubernetes-related operations.

#### [spf13/viper](https://github.com/spf13/viper)

- Viper is a configuration management library for Go used by `blitzctl`.
- It enables flexible configuration through files, environment variables, and command-line flags.
- `blitzctl` uses Viper to manage defaults, cluster state, and user preferences.

These tools and libraries enable `blitzctl` to provide a robust and user-friendly experience for managing Kubernetes clusters.
