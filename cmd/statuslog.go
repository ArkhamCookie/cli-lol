package cmd

import (
	"fmt"

	"omglol/statuslog"

	"github.com/charmbracelet/glamour"
	"github.com/spf13/cobra"
)

func statuslogListCmd(address string) {
	if address == "" {
		fmt.Println("Usage: lol status l[ist] <address> [flags]")
		return
	}

	result, err := statuslog.List(address)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result.Response.Message)
	fmt.Println(result.Response.Statuses)
}

func statuslogViewCmd(address, status string) {
	if address == "" {
		fmt.Println("Usage: lol status v[iew] <address> <status-id> [flags]")
		return
	}
	if status == "" {
		fmt.Println("Usage: lol status v[iew] <address> <status-id> [flags]")
		return
	}

	result, err := statuslog.Retrieve(address, status)
	if err != nil {
		fmt.Println(err)
		return
	}

	rendered := fmt.Sprintf("%s %s\n\n%s", result.Response.Status.Emoji, result.Response.Status.RelativeTime, result.Response.Status.Content)
	out, err := glamour.Render(rendered, "dark")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(out)
	// fmt.Println(result.Response.Status.ExternalURL) // Planned verbose response
}

func StatuslogBioCmd(address string) {
	if address == "" {
		fmt.Println("Usage lol status b[io] <address> [flags]")
		return
	}

	result, err := statuslog.BioView(address)
	if err != nil {
		fmt.Println(err)
		return
	}

	out, err := glamour.Render(result.Response.Bio, "dark")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(out)
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
				statuslogListCmd(args[1])
			} else {
				statuslogListCmd("")
			}
		case "v":
			fallthrough
		case "view":
			if len(args) >= 3 {
				statuslogViewCmd(args[1], args[2])
			} else {
				statuslogViewCmd("", "")
			}
		case "b":
			fallthrough
		case "bio":
			if len(args) > 1 {
				StatuslogBioCmd(args[1])
			} else {
				StatuslogBioCmd("")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(statuslogCmd)
}
