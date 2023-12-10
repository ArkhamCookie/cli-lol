package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func logo() {
	fmt.Println(`
      ____     ____
  ,-""    "-.-"    ""-,
 /  __  .  . .  .  __  \
|  (  ) '--' '--' (  )  |
 \  ""   ,     ,   ""  /
  ",      "---"      ,"
    ",             ,"
      "-,       ,-"
   sjw   "-,_,-"
   `)
}

var Verbose bool

var rootCmd = &cobra.Command{
	Use:   "lol",
	Short: "cli-lol is a commandline interface for omg.lol",
	// Long: "cli-lol is an open-source commandline interface for omg.lol built with Golang.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("cli-lol")
		if Verbose {
			logo()
		}
	},
}

func Execute() {
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
