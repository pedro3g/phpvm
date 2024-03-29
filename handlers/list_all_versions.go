package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/pedro3g/phpvm/types"
)

func ListAllVersions(echo bool) []types.Release {
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

	var releases []types.Release
	err = json.Unmarshal(body, &releases)

	if err != nil {
		fmt.Println("Error while parsing releases list")
		os.Exit(1)
	}

	if !echo {
		return releases
	}

	for _, release := range releases {
		fmt.Println(release.Tag)
	}

	return releases
}
