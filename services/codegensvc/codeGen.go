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
	} else if args["service"] != "" {
		handleGenerateService(args["service"])
	} else if args["router"] != "" {
		handleGenerateRouter(args["router"])
	} else if args["index"] != "" {
		handleGenerateIndexPage(args["index"])
	}
}

func handleGenerateApiController(args string) {
	fmt.Println("Generating Api Controller: " + args)
	cap := cases.Title(language.Und, cases.NoLower).String(args)
	smol := cases.Lower(language.Und).String(args)
	apiFile := strings.ReplaceAll(templates.APIControllerGin, "{{.Name}}", cap)
	apiFile = strings.ReplaceAll(apiFile, "{{.name}}", smol)
	filename := smol + "Controller.go"
	writeTofile(filename, apiFile)
}

func handleGenerateService(args string) {
	fmt.Println("Generating Service: " + args)
	cap := cases.Title(language.Und, cases.NoLower).String(args)
	smol := cases.Lower(language.Und).String(args)
	apiFile := strings.ReplaceAll(templates.SvcTemplate, "{{.Name}}", cap)
	apiFile = strings.ReplaceAll(apiFile, "{{.name}}", smol)
	filename := smol + "Service.go"
	writeTofile(filename, apiFile)
}

func handleGenerateRouter(args string) {
	lArgs := cases.Lower(language.Und).String(args)
	switch lArgs {
	case "gin":
		fmt.Println("Generating Gin Index page")
		filename := "ginRouter.go"
		writeTofile(filename, templates.GinRouterTemplate)
	default:
		fmt.Println("Unknown Router Type")

	}
}

func handleGenerateIndexPage(args string) {
	fmt.Println("Generating Index Page")
	lArgs := cases.Lower(language.Und).String(args)
	filename := "index.html"
	switch lArgs {
	case "gin":
		fmt.Println("Generating Gin Router")
		writeTofile(filename, templates.GinIndexHtml)
	default:
		fmt.Println("Unknown Router Type")

	}
}

func writeTofile(fileName string, content string) {
	var _, err = os.Stat(fileName)
	if os.IsNotExist(err) {
		if _, err := os.Create(fileName); err != nil {
			fmt.Println(err)
			return
		} else {
			if file, err := os.OpenFile(fileName, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660); err != nil {
				fmt.Println(err)
				return
			} else {
				defer file.Close()
				fmt.Fprintf(file, "%s\n", content)
			}
		}
	} else {
		fmt.Println("File already exists!", fileName)
		return
	}
}
