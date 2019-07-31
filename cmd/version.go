package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Ebot",
	Long:  `All software has versions. This is Ebot's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Ebot 1.0")
	},
}
