package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var DockerCmd = &cobra.Command{
	Use:     "d",
	Aliases: []string{"docker"},
	Short:   "Manage Docker Containers",
	Long:    `Manage Docker Containers`,
	Run: func(cmd *cobra.Command, args []string) {
		// ls := cmd.Flag("-o").Value.String()
		fmt.Println("Docker Managment Module")
	},
}

func init() {
	RootCmd.AddCommand(DockerCmd)
}
