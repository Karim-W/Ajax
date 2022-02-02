package DockerServices

import (
	"fmt"
	"os/exec"
	"strings"

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
	f := exec.Command("docker", "ps", "--format", "table {{.ID}}//\\//\\{{.Labels}}//\\//\\")
	stdout, _ := f.CombinedOutput()
	stdout = []byte(strings.ReplaceAll(string(stdout), "\n", ""))
	data := strings.Split(string(stdout), "//\\//\\")
	data = data[2:]
	fmt.Println("#\tContainerID\t\t Label")
	fmt.Println("--------------------------------------------------------------------------------------------")
	// max := len(data)
	for i, v := range data {
		if i%2 != 0 {
			index := i/2 + 1
			f := fmt.Sprintf("%d\t%s\t\t%s", index, data[i-1], v)
			fmt.Println(f)
		}
	}
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
