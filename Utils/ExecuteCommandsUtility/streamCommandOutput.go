package ExecuteCommandsUtility

import (
	"bufio"
	"fmt"
	"os/exec"
)

func StreamCommandOutput(cmd *exec.Cmd, c chan struct{}) {
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
