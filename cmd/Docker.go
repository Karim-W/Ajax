package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

var lsCmd = &cobra.Command{
	Use:     "d",
	Aliases: []string{"docker"},
	Short:   "Manage Docker Containers",
	Long:    `Manage Docker Containers`,
	Run: func(cmd *cobra.Command, args []string) {
		f := exec.Command("ls", "-l")
		stdout, _ := f.Output()
		fmt.Println(string(stdout))

	},
}

func init() {
	rootCmd.AddCommand(lsCmd)
}
