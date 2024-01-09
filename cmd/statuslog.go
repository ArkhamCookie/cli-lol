package cmd

import (
	"fmt"

	"omglol/statuslog"

	"github.com/charmbracelet/glamour"
	"github.com/spf13/cobra"
)

func listCmd(address string) {
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

func viewCmd(address, status string) {
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

	fmt.Println(result.Response.Message)
	fmt.Println() // To remove error about no new line with `fmt.Println()`
	fmt.Println(result.Response.Status.Emoji, result.Response.Status.RelativeTime)
	fmt.Println(result.Response.Status.Content)
	// fmt.Println(result.Response.Status.ExternalURL) // Planned verbose response
}

func bioCmd(address string) {
	if address == "" {
		fmt.Println("Usage lol status b[io] <address> [flags]")
		return
	}

	result, err := statuslog.BioView(address)
	if err != nil {
		fmt.Println(err)
		return
	}

	out, _ := glamour.Render(result.Response.Bio, "dark")
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
		case "b":
			fallthrough
		case "bio":
			if len(args) > 1 {
				bioCmd(args[1])
			} else {
				bioCmd("")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(statuslogCmd)
}
