package config

// Global defaults for the CLI
// CNI for Minikube: auto, bridge, calico, cilium, flannel, kindnet, or path to a CNI manifest (default: auto)
// Driver for Kind: docker, containerd, or path to a driver binary (default: docker)
// Driver for Minikube: docker, podman, virtualbox, vmware, kvm2, hyperkit, qemu, ssh, or path to a driver binary (default: docker)
// - Minikube - container-runtime: docker, containerd, cri-o, or auto (default: auto)
//   - Note: Not implemented yet
//
// - k8s release versions can be found at:
//   - https://kubernetes.io/releases/
const (
	DefaultK8sVersion  = "1.35.0"
	DefaultDriver      = "docker"
	DefaultClusterName = "blitz-cluster1"
	DefaultCni         = "cilium"

	DefaultHelmVersion = "3.19.5"
)
