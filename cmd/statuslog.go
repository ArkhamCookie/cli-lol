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

func viewCmd(address, status string) {
	if address == "" {
		println("Usage: lol status v[iew] <address> <status-id>")
		return
	}
	if status == "" {
		println("Usage: lol status v[iew] <address> <status-id>")
		return
	}

	result, err := statuslog.Retrieve(address, status)
	if err != nil {
		println(err)
		return
	}

	println(result.Response.Message, "\n")
	println(result.Response.Status.Emoji, result.Response.Status.RelativeTime)
	println(result.Response.Status.Content)
	// println(result.Response.Status.ExternalURL) // Planned verbose response
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
		case "v":
			fallthrough
		case "view":
			if len(args) >= 3 {
				viewCmd(args[1], args[2])
			} else {
				viewCmd("", "")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(statuslogCmd)
}
