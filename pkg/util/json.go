package util

import (
	"encoding/json"
	"fmt"
	"github.com/vebrasmusic/curl-echo/pkg"
	"io/ioutil"
	"os"
)

func LoadApiJson() ([]pkg.ApiRoute, *os.File) {
	const filePath = "curl-echo/apis.json"

	// Open the file
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("Failed to open file: %v\n", err)
		return nil, nil
	}

	// Read the current content of the file
	fileData, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("Failed to read file: %v\n", err)
		return nil, nil
	}

	// Initialize the list of API routes
	var apiRoutes []pkg.ApiRoute
	if len(fileData) > 0 {
		// Parse existing data
		err = json.Unmarshal(fileData, &apiRoutes)
		if err != nil {
			fmt.Printf("Failed to parse JSON: %v\n", err)
			return nil, nil
		}
	}

	return apiRoutes, file
}
