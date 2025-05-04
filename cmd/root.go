/*
Copyright © 2025 Oneide Luiz Schneider
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/OneideLuizSchneider/blitzctl/cmd/cluster"
	"github.com/OneideLuizSchneider/blitzctl/cmd/container"
)

var rootCmd = &cobra.Command{
	Use:   "blitzctl",
	Short: "The k8s local environment manager",
	Long: `A simple CLI tool to manage local Kubernetes environments.
It allows you to create, delete, and manage local Kubernetes clusters
using various drivers and configurations. It is designed to be
lightweight and easy to use, making it ideal for developers
who need a quick and efficient way to set up and manage
local Kubernetes environments.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := cmd.Help(); err != nil {
			// Handle the error (e.g., log it or print it)
			cmd.PrintErrln("❌ Error displaying help:", err)
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Add the cluster command to the root command
	rootCmd.AddCommand(cluster.GetClusterCmd())
	rootCmd.AddCommand(container.GetContainerCmd())
}
