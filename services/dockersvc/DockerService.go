package dockersvc

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"

	"github.com/karim-w/Ajax/Utils/ExecuteCommandsUtility"
)

func DockerService(args map[string]string) {
	if args["list"] == "true" {
		handleListContainers()
	} else if args["get"] != "" {
		handlePullContainer(args["get"])
	} else if args["kill"] != "" {
		handleKillContainer(args["kill"])
	} else if args["pregen"] != "" {
		handlePregenContainer(args["pregen"])
	}
}

func handlePregenContainer(cName string) {
	fmt.Println("Generating Docker Container: " + cName + " ...")
	fmt.Println(DBExamples[cName])
	arr := strings.Split(DBExamples[cName], " ")
	cm := exec.Command(arr[0], arr[1:]...)
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

var DBExamples = map[string]string{
	"redis":      "docker run --name my-redis -p 6379:6379 -d redis",
	"mssql":      `docker run -e "ACCEPT_EULA=Y" -e "SA_PASSWORD=secret" -p 1433:1433 --name sql1 --hostname sql1 -d mcr.microsoft.com/mssql/server:2019-latest`,
	"postgres":   "docker container run -d --name=pg -p 5432:5432 -e POSTGRES_PASSWORD=secret postgres:13",
	"helloWorld": "docker run -p 8080:8080 -d hello-world",
}

func handleListContainers() {
	f := exec.Command("docker", "ps", "--format", "table {{.ID}}//\\//\\{{.Names}}//\\//\\")
	stdout, _ := f.CombinedOutput()
	stdout = []byte(strings.ReplaceAll(string(stdout), "\n", ""))
	data := strings.Split(string(stdout), "//\\//\\")
	total := len(data)
	fmt.Println(total)
	data = data[2:]
	fmt.Println("#\tContainerID\t\t Label")
	fmt.Println("--------------------------------------------------------------------------------------------")
	// max := len(data)
	for i, v := range data {
		if i%2 != 0 {
			index := i/2 + 1
			f := fmt.Sprintf("%d\t%s\t\t%s", index-1, data[i-1], v)
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

func handleKillContainer(index string) {
	f := exec.Command("docker", "ps", "--format", "table {{.ID}}")
	stdout, _ := f.CombinedOutput()
	data := strings.Split(string(stdout), "\n")
	deletedIndex, err := strconv.ParseInt(index, 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	total := len(data)
	if total > 2 {
		total = total - 2
	}
	idel := int(deletedIndex)

	if idel >= total {
		fmt.Println("only found " + strconv.Itoa(total) + " containers")
		return
	} else {
		fmt.Println("Killing Docker Container: " + data[deletedIndex+1] + " ...")
		f := exec.Command("docker", "kill", data[deletedIndex+1])
		stdout, _ := f.CombinedOutput()
		fmt.Println(string(stdout))
	}
}
