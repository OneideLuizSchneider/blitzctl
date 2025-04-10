/*
Copyright © 2025 Oneide Luiz Schneider <...@...>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
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
		cmd.Help()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(listCmd)
}
