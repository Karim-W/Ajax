package clientsvc

import "fmt"

func ClientGenerator(args map[string]string) {
	if args["language"] == "go" && args["path"] != "" {
		fmt.Println("Listing Code Generators")
	}
}
