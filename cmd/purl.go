package cmd

import (
	"fmt"
	"omglol/purl"

	"github.com/spf13/cobra"
)

func purlRetrieveCmd(address, targetPurl string) {
	if address == "" {
		fmt.Println("Usage: lol purl g[et] <address> <purl> [flags]")
		return
	}
	if targetPurl == "" {
		fmt.Println("Usage: lol purl g[et] <address> <purl> [flags]")
		return
	}

	result, err := purl.Retrieve(address, targetPurl)
	if err != nil {
		fmt.Println(err)
		return
	}

	// fmt.Println(result.Response.Purl.Name) // Planned verbose response
	fmt.Println(result.Response.Purl.URL)
	// fmt.Println("Counter:", result.Response.Purl.Counter) // Planned verbose response
}

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

	for _, purl := range result.Response.Purls {
		var x struct {
			Name string
			URL string
			Counter int
		}

		x.Name = purl.Name
		x.URL = purl.URL
		x.Counter = purl.Counter

		fmt.Printf("https://purl.%s.omg.lol/%s = %s\n", address, x.Name, x.URL)
	}
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
		case "g":
			fallthrough
		case "get":
			if len(args) >= 3 {
				purlRetrieveCmd(args[1], args[2])
			} else {
				purlRetrieveCmd("", "")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(purlCmd)
}
