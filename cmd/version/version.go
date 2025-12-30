/*
Copyright © 2025 Oneide Luiz Schneider
*/
package version

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/OneideLuizSchneider/blitzctl/internal/version"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the blitzctl version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version.String())
	},
}

// GetVersionCmd exposes the version command to the root command.
func GetVersionCmd() *cobra.Command {
	return versionCmd
}
