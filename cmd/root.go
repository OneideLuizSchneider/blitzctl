/*
Copyright © 2025 Oneide Luiz Schneider
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	configCmd "github.com/OneideLuizSchneider/blitzctl/cmd/config"
	"github.com/OneideLuizSchneider/blitzctl/cmd/container"
	contextCmd "github.com/OneideLuizSchneider/blitzctl/cmd/context"
	createCmd "github.com/OneideLuizSchneider/blitzctl/cmd/create"
	deleteCmd "github.com/OneideLuizSchneider/blitzctl/cmd/delete"
	installCmd "github.com/OneideLuizSchneider/blitzctl/cmd/install"
	listCmd "github.com/OneideLuizSchneider/blitzctl/cmd/list"
	startCmd "github.com/OneideLuizSchneider/blitzctl/cmd/start"
	stopCmd "github.com/OneideLuizSchneider/blitzctl/cmd/stop"
	"github.com/OneideLuizSchneider/blitzctl/cmd/tools"
	upgradeCmd "github.com/OneideLuizSchneider/blitzctl/cmd/upgrade"
	versionCmdPkg "github.com/OneideLuizSchneider/blitzctl/cmd/version"
	"github.com/OneideLuizSchneider/blitzctl/config"
	"github.com/OneideLuizSchneider/blitzctl/internal/version"
)

var (
	configFile string
	rootCmd    = &cobra.Command{
		Use:     "blitzctl",
		Version: version.String(),
		Short:   "The k8s local environment manager",
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
)

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Initialize configuration before command execution
	cobra.OnInitialize(initConfig)
	rootCmd.SetVersionTemplate("{{printf \"%s\\n\" .Version}}")

	// Add global flags
	rootCmd.PersistentFlags().StringVar(&configFile, "config", "", "config file (default is $HOME/.blitzctl/config.yaml)")

	// Add commands to the root command
	rootCmd.AddCommand(createCmd.GetCreateCmd())
	rootCmd.AddCommand(deleteCmd.GetDeleteCmd())
	rootCmd.AddCommand(listCmd.GetListCmd())
	rootCmd.AddCommand(installCmd.GetInstallCmd())
	rootCmd.AddCommand(upgradeCmd.GetUpgradeCmd())
	rootCmd.AddCommand(startCmd.GetStartCmd())
	rootCmd.AddCommand(stopCmd.GetStopCmd())
	rootCmd.AddCommand(configCmd.GetConfigCmd())
	rootCmd.AddCommand(container.GetContainerCmd())
	rootCmd.AddCommand(contextCmd.GetContextCmd())
	rootCmd.AddCommand(tools.GetToolsmd())
	rootCmd.AddCommand(versionCmdPkg.GetVersionCmd())
}

// initConfig reads in config file and ENV variables
func initConfig() {
	if err := config.InitializeGlobalManager(configFile); err != nil {
		fmt.Printf("❌ Error initializing configuration: %v\n", err)
		os.Exit(1)
	}
}
