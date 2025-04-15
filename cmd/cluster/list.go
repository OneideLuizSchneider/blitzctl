/*
Copyright © 2025 Oneide Luiz Schneider <...@...>
Licensed under the Apache License, Version 2.0 (the "License");
*/
package cluster

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

func listMinikubeClusters(cmd *cobra.Command, args []string) {
	// Check if Minikube is installed
	_, err := exec.LookPath("minikube")
	if err != nil {
		fmt.Println("Minikube is not installed. Please install Minikube to use this command.")
		os.Exit(1)
	}
	// List Minikube clusters
	getCmd := exec.Command("minikube", "profile", "list")
	output, err := getCmd.Output()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting Minikube clusters: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Available Minikube clusters:")
	fmt.Println(string(output))
}

func listKindClusters(cmd *cobra.Command, args []string) {
	// Check if Kind is installed
	_, err := exec.LookPath("kind")
	if err != nil {
		fmt.Println("Kind is not installed. Please install Kind to use this command.")
		os.Exit(1)
	}
	// List Kind clusters
	getCmd := exec.Command("kind", "get", "clusters")
	output, err := getCmd.Output()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting Kind clusters: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Available Kind clusters:")
	fmt.Println(string(output))
}
