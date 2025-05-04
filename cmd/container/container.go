/*
Copyright © 2025 Oneide Luiz Schneider
*/
package container

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
	"k8s.io/kubectl/pkg/util/i18n"
	"k8s.io/kubectl/pkg/util/templates"
)

var (
	listExample = templates.Examples(i18n.T(`
		# this will open the webbrowser for you to install the docker desktop
		blitzctl container install docker`))

	containerCmd = &cobra.Command{
		Use:     "container",
		Example: listExample,
		Aliases: []string{"c"},
		Short:   "Install a container driver",
		Long:    `It opens the webbrowser to install docker desktop manually for example.`,
		Run: func(cmd *cobra.Command, args []string) {
			if err := cmd.Help(); err != nil {
				// Handle the error (e.g., log it or print it)
				cmd.PrintErrln("❌ Error displaying help:", err)
			}
		},
	}

	dockerCmd = &cobra.Command{
		Use:     "docker",
		Example: listExample,
		Aliases: []string{"d"},
		Short:   "Docker website",
		Long:    `Docker website`,
		Run: func(cmd *cobra.Command, args []string) {
			getCmd := exec.Command("open", "https://www.docker.com/get-started/")
			// Set up real-time output streaming
			getCmd.Stdout = os.Stdout
			getCmd.Stderr = os.Stderr
			// Start and wait for the command to complete
			if err := getCmd.Run(); err != nil {
				fmt.Fprintf(os.Stderr, "❗Error opening docker website\n")
				os.Exit(1)
			}
		},
	}

	podmanCmd = &cobra.Command{
		Use:     "podman",
		Example: listExample,
		Aliases: []string{"p"},
		Short:   "Podman website",
		Long:    `Podman website`,
		Run: func(cmd *cobra.Command, args []string) {
			getCmd := exec.Command("open", "https://podman-desktop.io/downloads")
			// Set up real-time output streaming
			getCmd.Stdout = os.Stdout
			getCmd.Stderr = os.Stderr
			// Start and wait for the command to complete
			if err := getCmd.Run(); err != nil {
				fmt.Fprintf(os.Stderr, "❗Error opening podman website\n")
				os.Exit(1)
			}
		},
	}
)

func GetContainerCmd() *cobra.Command {
	return containerCmd
}

func init() {
	containerCmd.AddCommand(dockerCmd)
	containerCmd.AddCommand(podmanCmd)
}
