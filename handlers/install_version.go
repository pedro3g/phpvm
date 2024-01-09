package handlers

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/pedro3g/phpvm/utils"
)

var (
	baseDir        = filepath.Dir(os.Args[0])
	releasesFolder = baseDir + "/releases"
)

func InstallVersion(version string) {
	fmt.Println("Downloading PHP version", version)

	if _, err := os.Stat(releasesFolder); os.IsNotExist(err) {
		os.Mkdir(releasesFolder, os.ModePerm)
	}

	downloadReleaseUrl := "https://windows.php.net/downloads/releases/php-" + version + "-nts-Win32-vs16-x64.zip"
	filePath := releasesFolder + "/php-" + version

	err := utils.DownloadFile(downloadReleaseUrl, filePath+".zip")

	if err != nil {
		fmt.Println("An error occurred while downloading PHP")
		return
	}

	fmt.Println("Downloaded PHP version", version)
	fmt.Println("Extracting PHP version", version)

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		os.Mkdir(filePath, os.ModePerm)
	}

	cmd := exec.Command("tar", "-xzf", filePath+".zip", "-C", filePath)
	err = cmd.Run()

	if err != nil {
		fmt.Println("An error occurred while extracting PHP:", err)
		return
	}

	err = os.Remove(filePath + ".zip")

	if err != nil {
		fmt.Println("An error occurred while removing the downloaded zip file:", err)
		return
	}

	fmt.Println("Installation complete!")
}
