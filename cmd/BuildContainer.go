package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

var dockerBuildCmd = &cobra.Command{
	Use:     "b",
	Aliases: []string{"build", "-b"},
	Short:   "Build Docker Containers",
	Long:    `Build Docker Containers`,
	Run: func(cmd *cobra.Command, args []string) {
		f := exec.Command("docker", "build", ".")
		stdout, _ := f.CombinedOutput()
		fmt.Println(string(stdout))
	},
}

func init() {
	DockerCmd.AddCommand(dockerBuildCmd)
}
