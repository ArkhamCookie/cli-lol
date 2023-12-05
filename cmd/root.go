package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "lol",
	Short: "cli-lol is a commandline interface for omg.lol",
	Long: `cli-lol is an open-source commandline interface for https://omg.lol built with Golang.
You can find the source code at https://git.theuntitledproject.dev/cli-lol, and the documentation at https://docs.theuntitledproject.dev/cli-lol.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
