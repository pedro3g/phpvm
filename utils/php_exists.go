package utils

import (
	"os/exec"
	"path/filepath"
	"strings"
)

func PhpExists() (exists bool, path string, version string) {
	cmd := exec.Command("where", "php")
	output, err := cmd.CombinedOutput()

	path = ""

	exists = err == nil

	if exists {
		path = filepath.Dir(strings.TrimSpace(string(output)))

		cmd = exec.Command("php", "-r", "echo phpversion();")
		output, err = cmd.CombinedOutput()

		if err != nil {
			panic(err)
		}

		version = strings.TrimSpace(string(output))
	}

	return exists, path, version
}
