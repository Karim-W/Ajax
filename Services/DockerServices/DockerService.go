package DockerServices

import (
	"fmt"
	"os/exec"

	"github.com/karim-w/Ajax/Utils/ExecuteCommandsUtility"
)

func DockerService(args map[string]string) {
	if args["list"] == "true" {
		handleListContainers()
	} else if args["get"] != "" {
		handlePullContainer(args["get"])
	}

}

func handleListContainers() {
	fmt.Println("Listing Containers...")
	f := exec.Command("docker", "ps")
	stdout, _ := f.CombinedOutput()
	fmt.Println(string(stdout))
}

func handlePullContainer(cName string) {
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
	fmt.Println("Docker Image: " + cName + " pulled")
}
