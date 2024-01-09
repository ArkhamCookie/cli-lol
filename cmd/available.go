package cmd

import (
	"fmt"
	"omglol/address"

	"github.com/spf13/cobra"
)

var availableCmd = &cobra.Command{
	Use:   "available <address>",
	Short: "Get the availability of an address",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		response, err := address.Available(args[0])
		if err != nil {
			fmt.Println(err)
		}

		// TODO: Format the response before printing it
		fmt.Println(response)
	},
}

func init() {
	rootCmd.AddCommand(availableCmd)
}
