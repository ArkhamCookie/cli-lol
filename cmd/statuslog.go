package cmd

import (
	"omglol/statuslog"

	"github.com/spf13/cobra"
)

func listCmd(address string) {
	if address != "" {
		statuslog.List(address)
	} else {
		println("Usage:")
	}
}

var statuslogCmd = &cobra.Command{
	Use:   "status <command> [flags]",
	Short: "View or manage statuses",
	Long:  "Core command for viewing and managing statuslogs.",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "l":
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
