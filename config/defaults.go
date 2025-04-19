package config

// Global defaults for the CLI
// CNI for Minikube: auto, bridge, calico, cilium, flannel, kindnet, or path to a CNI manifest (default: auto)
const (
	DefaultK8sVersion       = "1.32.0"
	DefaultContainerRuntime = "podman"
	DefaultClusterName      = "minikube"
	DefaultCni              = "cilium"
)
