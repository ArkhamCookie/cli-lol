package cmd

import (
	"omglol/statuslog"

	"github.com/spf13/cobra"
)

func listCmd(address string) {
	if address == "" {
		println("Usage: lol status l[ist] <address>")
		return
	}

	result, err := statuslog.List(address)
	if err != nil {
		println(err)
		return
	}

	println(result.Response.Message)
	println(result.Response.Statuses)
}

var statuslogCmd = &cobra.Command{
	Use:   "status <command> [flags]",
	Short: "View or manage statuses",
	Long:  "Core command for viewing and managing statuslogs.",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "l":
			fallthrough
		case "list":
			if len(args) > 1 {
				listCmd(args[1])
			} else {
				listCmd("")
			}
		case "list":
			if len(args) > 1 {
				listCmd(args[1])
			} else {
				listCmd("")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(statuslogCmd)
}
