package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	rootCmd.AddCommand(versionCmd)

	generateCmd.Flags().BoolP("designed", "d", false, "whether the module is designed or not(default is false)")
	generateCmd.Flags().StringP("workspace", "w", "", "the parent path of the repo(default is $GOPATH)")
	rootCmd.AddCommand(generateCmd)
}

var rootCmd = &cobra.Command{
	Use:   "ebot",
	Short: "Ebot is a tool for generating code",
	//Run: func(cmd *cobra.Command, args []string) {
	//	// Do Stuff Here
	//},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
