# blitzctl

`blitzctl` is a CLI tool for managing local Kubernetes environments. It simplifies the creation, deletion, upgrading, and management of Kubernetes clusters using tools like `Minikube` and `Kind`.

Currently supports `macOS` and `Linux`.

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

#### Flags

- `--cluster-name`: Specify the name of the cluster.
- `--k8s-version`: Specify the Kubernetes version.
- `--driver`: Specify the driver (e.g., Docker, Podman, virtualbox, parallels, hyperkit, vmware, qemu2, vfkit).
  - For `kind`, only Docker.

## Configuration

Default configurations are defined in `config/defaults.go`:

- **Kubernetes Version**: `1.32.0`
- **Driver**: `docker`
- **Cluster Name**: `blitz-cluster1`
- **CNI Plugin**: `cilium`

---

## Examples

#### Create a Minikube Cluster

Create a Kubernetes cluster using Minikube with a specific Kubernetes version and container runtime:

```sh
blitzctl cluster create minikube --cluster-name=mycluster --k8s-version=1.32.0 --driver=docker
```

#### Delete a Kind Cluster

Delete a Kubernetes cluster created with Kind:

```sh
blitzctl cluster delete kind --cluster-name=mycluster
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

#### [Kubectl Utilities](https://kubernetes.io/docs/reference/kubectl/)
- The project uses utilities from the Kubernetes `kubectl` package for handling Kubernetes-related operations.

#### [spf13/viper](https://github.com/spf13/viper) (Optional, if used)
- Viper is a configuration management library for Go.
- It simplifies reading configuration files, environment variables, and flags.

These tools and libraries enable `blitzctl` to provide a robust and user-friendly experience for managing Kubernetes clusters.