package cmd

import (
	"bufio"
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

func run(cmd *exec.Cmd, c chan struct{}) {
	defer func() { c <- struct{}{} }()
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}
	<-c
	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		m := scanner.Text()
		fmt.Println(m)
	}
}

var dockerPullCmd = &cobra.Command{
	Use:     "get",
	Aliases: []string{"-g", "pull"},
	Short:   "Manage Docker Containers",
	Long:    `Manage Docker Containers`,
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

		go run(cm, c)

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
