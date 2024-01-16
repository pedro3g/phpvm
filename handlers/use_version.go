package handlers

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/pedro3g/phpvm/utils"
)

func UseVersion(version string) {
	baseDir := utils.GetBaseDir()
	releasesFolder := filepath.Join(baseDir, "releases")
	filePath := filepath.Join(releasesFolder, "php-"+version)

	if !utils.AlreadyDownloaded(version) {
		fmt.Println("Version", version, "not installed. Please install it first")
		return
	}

	phpExists, phpPath, actualVersion := utils.PhpExists()

	if phpExists && actualVersion == version {
		fmt.Println("PHP version", actualVersion, "already in use")
		return
	}

	if phpExists && actualVersion != version {
		oldVars := os.Getenv("PATH")
		newVars := strings.Replace(oldVars, phpPath, filePath, -1)

		err := exec.Command("setx", "PATH", newVars).Run()

		if err != nil {
			panic(err)
		}

		fmt.Println("PHP version changed from", actualVersion, "to", version)
	}

	oldVars := os.Getenv("PATH")
	newVars := filePath + ";" + oldVars

	err := exec.Command("setx", "PATH", newVars).Run()

	if err != nil {
		panic(err)
	}

	fmt.Println("PHP version changed to", version)
}
