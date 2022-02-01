package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "ajax",
	Short: "Ajax is a tool for simlifying dev tools",
	Long:  "Ajax is a tool for simlifying dev tools.",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		fmt.Println(`    _          _      _     __  __
   / \        | |    / \    \ \/ /
  / _ \    _  | |   / _ \    \  / 
 / ___ \  | |_| |  / ___ \   /  \ 
/_/   \_\  \___/  /_/   \_\ /_/\_\
                                  `)
		fmt.Println("Welcome To Ajax")
		fmt.Println("How Can I Help")

	},
}

func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
