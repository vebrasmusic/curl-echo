package util

import (
	"fmt"
	"os"
	"path/filepath"
)

func CreateFolders(folderPath string) {
	// Get the current working directory
	cwd, err1 := os.Getwd()
	if err1 != nil {
		fmt.Printf("Failed to get current working directory: %v\n", err1)
		return
	}

	path := filepath.Join(cwd, "curl-echo", folderPath)
	err2 := os.MkdirAll(path, os.ModePerm)
	if err2 != nil {
		fmt.Printf("Failed to create directory: %v\n", err2)
		return
	}
}
