package cmd

import (
	"fmt"
	"os/exec"

	"github.com/karim-w/Ajax/Utils/ExecuteCommandsUtility"
	"github.com/spf13/cobra"
)

var dockerPullCmd = &cobra.Command{
	Use:     "get",
	Aliases: []string{"-g", "pull"},
	Short:   "Pull Docker Containers",
	Long:    `Pull Docker Containers`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("No Docker Image Specified to Pull")
			cmd.Help()
			return
		}
		cName := args[0]
		if cName == "" {
			fmt.Println("Error: Missing required flag(s)")
			cmd.Help()
			return
		}
		fmt.Println("Pulling Docker Image: " + cName + " ...")
		cm := exec.Command("docker", "pull", cName)
		c := make(chan struct{})

		go ExecuteCommandsUtility.StreamCommandOutput(cm, c)

		c <- struct{}{}
		cm.Start()

		<-c
		if err := cm.Wait(); err != nil {
			fmt.Println(err)
		}
		fmt.Println("Docker Image: " + cName + " Pulled")
	},
}

func init() {
	DockerCmd.AddCommand(dockerPullCmd)
}
