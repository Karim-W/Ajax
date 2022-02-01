package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

var dockerLsCmd = &cobra.Command{
	Use:     "l",
	Aliases: []string{"list"},
	Short:   "Manage Docker Containers",
	Long:    `Manage Docker Containers`,
	Run: func(cmd *cobra.Command, args []string) {
		f := exec.Command("docker", "ps")
		stdout, _ := f.CombinedOutput()
		fmt.Println(string(stdout))
	},
}

func init() {
	DockerCmd.AddCommand(dockerLsCmd)
}
