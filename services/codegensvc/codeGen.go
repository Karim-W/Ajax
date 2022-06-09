package codegensvc

import (
	"fmt"
	"os"
	"strings"

	"github.com/karim-w/Ajax/templates"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func CodeGenerator(args map[string]string) {
	if args["list"] == "true" {
		fmt.Println("Listing Code Generators")
	} else if args["controller"] != "" {
		handleGenerateApiController(args["controller"])
	} else if args["kill"] != "" {
	} else if args["get"] != "" {
	} else if args["pregen"] != "" {
		fmt.Println("Pregen")
	}
}

func handleGenerateApiController(args string) {
	fmt.Println("Generating Api Controller: " + args)
	cap := cases.Title(language.Und, cases.NoLower).String(args)
	smol := cases.Lower(language.Und).String(args)
	apiFile := strings.ReplaceAll(templates.APIControllerGin, "{{.Name}}", cap)
	apiFile = strings.ReplaceAll(apiFile, "{{.name}}", smol)
	filename := smol + "Controller.go"
	var _, err = os.Stat(filename)
	if os.IsNotExist(err) {
		if _, err := os.Create(filename); err != nil {
			fmt.Println(err)
			return
		} else {
			if file, err := os.OpenFile(filename, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660); err != nil {
				fmt.Println(err)
				return
			} else {
				defer file.Close()
				fmt.Fprintf(file, "%s\n", apiFile)
			}
		}
	} else {
		fmt.Println("File already exists!", filename)
		return
	}
}
