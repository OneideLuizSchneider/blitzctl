# blitzctl

`blitzctl` is a CLI tool for managing local Kubernetes environments. It simplifies the creation, deletion, and management of Kubernetes clusters using tools like `Minikube` and `Kind`.

## Usage

### General Command Structure

```sh
blitzctl <command> <subcommand> [flags]
```

### Commands

#### Cluster Commands

- `create`: Create a Kubernetes cluster.
- `delete`: Delete a Kubernetes cluster.
- `list`: List all available clusters.
- `install`: Install tools like Minikube or Kind.

### Flags

- `--cluster-name`: Specify the name of the cluster.
- `--k8s-version`: Specify the Kubernetes version.
- `--driver`: Specify the driver (e.g., Docker, Podman, virtualbox, parallels, hyperkit, vmware, qemu2, vfkit).
  - For `kind`, only Docker.

## Configuration

Default configurations are defined in `config/defaults.go`:

- **Kubernetes Version**: `1.32.0`
- **Driver**: `podman`
- **Cluster Name**: `minikube`
- **CNI Plugin**: `cilium`

---

## Examples

### Create a Minikube Cluster

Create a Kubernetes cluster using Minikube with a specific Kubernetes version and container runtime:

```sh
blitzctl cluster create minikube --cluster-name=mycluster --k8s-version=1.32.0 --driver=docker
```

### Delete a Kind Cluster

Delete a Kubernetes cluster created with Kind:

```sh
blitzctl cluster delete kind --cluster-name=mycluster
```

### List All Minikube Clusters

List all Kubernetes clusters managed by Minikube:

```sh
blitzctl cluster list minikube
```

### Install Minikube

Install Minikube on your system:

```sh
blitzctl cluster install minikube
```

### Create a Kind Cluster with Default Settings

Create a Kubernetes cluster using Kind with default configurations:

```sh
blitzctl cluster create kind
```

### Delete a Cluster with Custom Name

Delete a cluster with a custom name:

```sh
blitzctl cluster delete kind --cluster-name=custom-cluster
```

### Debugging a Cluster Deletion

Run a cluster deletion command with debug output enabled:

```sh
blitzctl cluster delete kind --cluster-name=mycluster --debug
```