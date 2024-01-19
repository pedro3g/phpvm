package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"

	"github.com/pedro3g/phpvm/handlers"
)

var (
	version = "0.0.40"
)

func main() {
	if runtime.GOOS != "windows" {
		fmt.Println("This tool only works on Windows")
		os.Exit(1)
	}

	flagVersion := flag.Bool("v", false, "Show phpvm version")
	listAll := flag.Bool("list-all", false, "List all PHP versions available")

	flag.Parse()

	if *flagVersion {
		fmt.Println("Version:", version)
		return
	} else if *listAll {
		handlers.ListAllVersions(true)
		return
	} else if flag.Arg(0) == "install" {
		handlers.InstallVersion(flag.Arg(1))
		return
	} else if flag.Arg(0) == "use" {
		handlers.UseVersion(flag.Arg(1))
		return
	} else {
		flag.PrintDefaults()
		return
	}
}
