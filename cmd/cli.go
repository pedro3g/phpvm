package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"

	"github.com/pedro3g/phpvm/handlers"
)

var (
	version = "1.0.31"
)

func main() {
	if runtime.GOOS != "windows" {
		fmt.Println("This tool only works on Windows")
		os.Exit(1)
	}

	flagVersion := flag.Bool("v", false, "Show phpvm version")
	listAll := flag.Bool("list-all", false, "List all PHP versions available")
	install := flag.String("install", "", "Install a PHP version")

	flag.Parse()

	if *flagVersion {
		fmt.Println("Version:", version)
		return
	} else if *listAll {
		handlers.ListAllVersions(true)
		return
	} else if *install != "" {
		handlers.InstallVersion(*install)
		return
	}
}
