package config

// Global defaults for the CLI
// CNI for Minikube: auto, bridge, calico, cilium, flannel, kindnet, or path to a CNI manifest (default: auto)
// Driver for Kind: docker, containerd, or path to a driver binary (default: podman)
// Driver for Minikube: docker, podman, virtualbox, vmware, kvm2, hyperkit, qemu, ssh, or path to a driver binary (default: docker)
// - Minikube - container-runtime: docker, containerd, cri-o, or auto (default: auto)
//   - Note: Not implemented yet
const (
	DefaultK8sVersion  = "1.33.4"
	DefaultDriver      = "podman"
	DefaultClusterName = "blitz-cluster1"
	DefaultCni         = "cilium"

	DefaultHelmVersion = "v3.18.6"
)
