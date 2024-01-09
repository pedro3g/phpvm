package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Announcement struct {
	English string `json:"English,omitempty"`
}

type Source struct {
	Filename string `json:"filename"`
	Name     string `json:"name"`
	SHA256   string `json:"sha256"`
	Date     string `json:"date"`
}

type VersionInfo struct {
	Announcement interface{} `json:"announcement"`
	Tags         []string    `json:"tags"`
	Date         string      `json:"date"`
	Source       []Source    `json:"source"`
	Museum       bool        `json:"museum,omitempty"`
}

type Response map[string]VersionInfo

func ListAllVersions() {
	availableVersions := [5]int{3, 4, 5, 7, 8}

	versionsMap := make(map[string]VersionInfo)

	for _, i := range availableVersions {
		resp, err := http.Get("https://www.php.net/releases/index.php?json=1&version=" + fmt.Sprint(i) + "&max=1000")

		if err != nil {
			fmt.Println("An error occurred while getting PHP list")
			return
		}

		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)

		if err != nil {
			fmt.Println("An error occurred while reading PHP list")
			return
		}

		var versions Response

		err = json.Unmarshal(body, &versions)

		if err != nil {
			fmt.Println("An error occurred while parsing PHP list")
			return
		}

		for version, info := range versions {
			versionsMap[version] = info
		}
	}

	for version, info := range versionsMap {
		fmt.Println(version, info.Date)
	}
}
