package util

import (
	"encoding/json"
	"fmt"
	"github.com/vebrasmusic/curl-echo/pkg"
	"io/ioutil"
	"os"
)

type FilterSpec struct {
	Param     string
	ParamType string
}

var filterFunctions = map[string]func(pkg.ApiRoute) string{
	"Nickname": func(r pkg.ApiRoute) string { return r.Nickname },
	"Group":    func(r pkg.ApiRoute) string { return r.Group },
	"Route":    func(r pkg.ApiRoute) string { return r.Route },
}

func FilterRoutes(array []pkg.ApiRoute, spec FilterSpec) []pkg.ApiRoute {
	// filter based on the cmd line arg
	var filteredApiRoutes []pkg.ApiRoute

	filterFunction, exists := filterFunctions[spec.Param]
	if !exists {
		fmt.Println("Filter function error")
		return nil
	}
	for _, apiRoute := range array {
		if filterFunction(apiRoute) == spec.Param {
			filteredApiRoutes = append(filteredApiRoutes, apiRoute)
		}
	}
	return filteredApiRoutes
}

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
