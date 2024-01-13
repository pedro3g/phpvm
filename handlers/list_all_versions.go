package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime"
)

type Release struct {
	Tag    string `json:"tag"`
	Source Source `json:"source"`
}

type Source map[string]string

func ListAllVersions() {
	if runtime.GOOS != "windows" {
		fmt.Println("This tool only works on Windows")
		os.Exit(1)
	}

	releasesUrl := "https://raw.githubusercontent.com/pedro3g/win-php-bin/master/releases.json"

	resp, err := http.Get(releasesUrl)
	if err != nil {
		fmt.Println("Error while fetching releases list")
		os.Exit(1)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Error while reading releases list")
		os.Exit(1)
	}

	var releases []Release
	err = json.Unmarshal(body, &releases)

	if err != nil {
		fmt.Println("Error while parsing releases list")
		os.Exit(1)
	}

	for _, release := range releases {
		fmt.Println(release.Tag)
	}
}
