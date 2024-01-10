package utils

import (
	"os/exec"
	"path/filepath"
	"strings"
)

func PhpExists() (bool, string) {
	cmd := exec.Command("where", "php")
	output, err := cmd.CombinedOutput()

	path := filepath.Dir(strings.TrimSpace(string(output)))

	return err == nil, path
}
