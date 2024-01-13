package utils

import (
	"os"
	"path/filepath"
)

var (
	baseDir = GetBaseDir()
)

func AlreadyDownloaded(version string) bool {
	_, err := os.Stat(filepath.Join(baseDir, "releases", "php-"+version))

	return !os.IsNotExist(err)
}
