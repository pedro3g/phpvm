package utils

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func GetBaseDir() (baseDir string) {
	godotenv.Load()

	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("An error occurred while getting the current directory")
		panic(err)
	}
	baseDir = Ternary(os.Getenv("MODE") == "dev", wd, filepath.Dir(os.Args[0])).(string)
	return
}
