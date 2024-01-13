package handlers

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/pedro3g/phpvm/utils"
)

var (
	baseDir        = utils.GetBaseDir()
	releasesFolder = filepath.Join(baseDir, "releases")
)

func InstallVersion(version string) {
	phpExists, phpPath, actualVersion := utils.PhpExists()

	if phpExists && filepath.Join(releasesFolder, "php-"+actualVersion) != phpPath {
		fmt.Println("PHP version", actualVersion, "found at", phpPath)
		fmt.Println("If you can not use phpvm, please uninstall it first")
		os.Exit(1)
	}

	if utils.AlreadyDownloaded(version) {
		fmt.Println("PHP version", version, "already downloaded. If you want to use this version, use phpvm --use", version)
		os.Exit(1)
	}

	versionAvailable, source := CheckVersionAvailability(version)

	if !versionAvailable {
		fmt.Println("Version", version, "not available")
		os.Exit(1)
	}

	fmt.Println("Downloading PHP version", version)

	if _, err := os.Stat(releasesFolder); os.IsNotExist(err) {
		os.Mkdir(releasesFolder, os.ModePerm)
	}

	downloadReleaseUrl := "https://raw.githubusercontent.com/pedro3g/win-php-bin/master/releases/" + source

	filePath := filepath.Join(releasesFolder, "php-"+version)
	splitDotLen := strings.Split(source, ".")
	extension := splitDotLen[len(splitDotLen)-1]

	err := utils.DownloadFile(downloadReleaseUrl, filePath+"."+extension)

	if err != nil {
		fmt.Println("An error occurred while downloading PHP")
		panic(err)
	}

	fmt.Println("Extracting PHP version", version)

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		os.Mkdir(filePath, os.ModePerm)
	}

	cmd := exec.Command("tar", "-xzf", filePath+"."+extension, "-C", filePath)
	err = cmd.Run()

	if err != nil {
		fmt.Println("An error occurred while extracting PHP")
		panic(err)
	}

	err = os.Remove(filePath + "." + extension)

	if err != nil {
		fmt.Println("An error occurred while removing the downloaded zip file")
		panic(err)
	}

	if _, err := os.Stat(phpPath); os.IsNotExist(err) {
		os.Mkdir(phpPath, os.ModePerm)
	}

	oldVars := os.Getenv("PATH")

	if phpPath != "" {
		oldVars = strings.Replace(oldVars, phpPath+";", "", -1)
	}

	err = exec.Command("setx", "PATH", filePath+";"+oldVars).Run()

	if err != nil {
		fmt.Println("An error occurred while adding PHP to the PATH")
		panic(err)
	}

	fmt.Println("Installation complete! Please restart your terminal to apply the changes.")
}
