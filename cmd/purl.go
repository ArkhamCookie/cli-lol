package cmd

import (
	"fmt"
	"omglol/purl"

	"github.com/spf13/cobra"
)

func purlListCmd(address string) {
	if address == "" {
		fmt.Println("Usage: lol purl l[ist] <address> [flags]")
		return
	}

	result, err := purl.List(address)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result.Response.Message)
	fmt.Println(result.Response.Purls)
}

var purlCmd = &cobra.Command{
	Use: "purl <command> [flags]",
	Short: "View or manage purls",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "l":
			fallthrough
		case "list":
			if len(args) > 1 {
				purlListCmd(args[1])
			} else {
				purlListCmd("")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(purlCmd)
}
