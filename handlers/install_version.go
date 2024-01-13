package handlers

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/pedro3g/phpvm/utils"
)

func InstallVersion(version string) {

	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("An error occurred while getting the current directory")
		panic(err)
	}

	baseDir := utils.Ternary(os.Getenv("MODE") == "dev", wd, filepath.Dir(os.Args[0])).(string)
	releasesFolder := filepath.Join(baseDir, "releases")

	phpExists, phpPath, actualVersion := utils.PhpExists()

	if phpExists && actualVersion == version {
		fmt.Println("PHP version", version, "already installed at", phpPath, ". Nothing to do.")
		os.Exit(0)
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

	filePath := releasesFolder + "/php-" + version

	err = utils.DownloadFile(downloadReleaseUrl, filePath+".zip")

	if err != nil {
		fmt.Println("An error occurred while downloading PHP")
		panic(err)
	}

	fmt.Println("Downloaded PHP version", version)
	fmt.Println("Extracting PHP version", version)

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		os.Mkdir(filePath, os.ModePerm)
	}

	cmd := exec.Command("tar", "-xzf", filePath+".zip", "-C", filePath)
	err = cmd.Run()

	if err != nil {
		fmt.Println("An error occurred while extracting PHP")
		panic(err)
	}

	err = os.Remove(filePath + ".zip")

	if err != nil {
		fmt.Println("An error occurred while removing the downloaded zip file")
		panic(err)
	}

	if phpExists {
		fmt.Println("Moving the old PHP version to", phpPath+".old")

		err := os.Rename(phpPath, phpPath+".old")

		if err != nil {
			fmt.Println("An error occurred while renaming the old PHP version")
			panic(err)
		}
	}

	if _, err := os.Stat(phpPath); os.IsNotExist(err) {
		os.Mkdir(phpPath, os.ModePerm)
	}

	if phpExists {
		// perguntar se deseja manter o php.ini atual
		var keepIni string
		fmt.Print("Do you want to keep the current php.ini file? (y/n) ")
		fmt.Scan(&keepIni)

		err = os.Rename(filePath, phpPath)

		if keepIni == "y" {
			err = os.Rename(phpPath+".old/php.ini", phpPath+"/php.ini")
		}
	} else {
		_, err := os.Stat("C:/php")

		if err == nil {
			err = os.RemoveAll("C:/php")

			if err != nil {
				fmt.Println("An error occurred while removing the old PHP version")
				panic(err)
			}
		}

		err = os.Rename(filePath, "C:/php")

		if err != nil {
			fmt.Println("An error occurred while renaming the new PHP version")
			panic(err)
		}

		oldVars := os.Getenv("PATH")

		err = exec.Command("setx", "PATH", "C:\\php;"+oldVars).Run()

		if err != nil {
			fmt.Println("An error occurred while adding PHP to the PATH")
			panic(err)
		}
	}

	if err != nil {
		fmt.Println("An error occurred while renaming the new PHP version")
		panic(err)
	}

	fmt.Println("Installation complete! Please restart your terminal to apply the changes.")
}
