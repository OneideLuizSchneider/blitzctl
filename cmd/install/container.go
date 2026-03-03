/*
Copyright © 2025 Oneide Luiz Schneider
*/
package install

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/OneideLuizSchneider/blitzctl/cmd/cluster/provider"
	"github.com/spf13/cobra"
	"k8s.io/kubectl/pkg/util/i18n"
	"k8s.io/kubectl/pkg/util/templates"
)

var (
	listExample = templates.Examples(i18n.T(`
		# this will open the webbrowser for you to install the docker desktop
		blitzctl install container --driver=docker
		# this will open the webbrowser for you to install the podman desktop
		blitzctl install container --driver=podman
	`))

	containerCmd = &cobra.Command{
		Use:     "container",
		Example: listExample,
		Aliases: []string{"c"},
		Short:   "Install a container driver",
		Long:    `It opens the webbrowser to install docker desktop manually for example.`,
		Run: func(cmd *cobra.Command, args []string) {
			if containerDriver == "" {
				fmt.Fprintf(os.Stderr, "❗Error: --driver flag is required (docker or podman)\n")
				os.Exit(1)
			}
			switch containerDriver {
			case string(provider.Docker):
				fmt.Fprintf(os.Stderr, "❗Opening docker website...\n")
				dockerCmd.Run(cmd, args)
			case string(provider.Podman):
				fmt.Fprintf(os.Stderr, "❗Opening podman website...\n")
				podmanCmd.Run(cmd, args)
			default:
				fmt.Fprintf(os.Stderr, "❗Error: invalid driver '%s'. Valid options are 'docker' or 'podman'\n", containerDriver)
				os.Exit(1)
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

	containerDriver string
)

func GetContainerCmd() *cobra.Command {
	return containerCmd
}

func init() {
	containerCmd.Flags().StringVarP(&containerDriver, "driver", "d", string(provider.Docker), i18n.T("Container driver (docker or podman)."))
}
