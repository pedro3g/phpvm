package main

import (
	"flag"
	"fmt"

	"github.com/pedro3g/phpvm/cmd/handlers"
)

func main() {
	flagVersion := flag.Bool("v", false, "Show phpvm version")
	listAll := flag.Bool("list-all", false, "List all PHP versions available")

	flag.Parse()

	if *flagVersion {
		fmt.Println("Version:", version)
		return
	} else if *listAll {
		handlers.ListAllVersions()
		return
	}
}
