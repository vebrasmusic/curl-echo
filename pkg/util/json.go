package util

import (
	"encoding/json"
	"fmt"
	"github.com/vebrasmusic/curl-echo/pkg"
	"io/ioutil"
	"os"
	"path/filepath"
)

type FilterSpec struct {
	Param     string
	ParamType string
}

// Filter function map with cleaner usage of function types
var filterFunctions = map[string]func(pkg.ApiRoute) string{
	"Nickname": func(r pkg.ApiRoute) string { return r.Nickname },
	"Group":    func(r pkg.ApiRoute) string { return r.Group },
	"Route":    func(r pkg.ApiRoute) string { return r.Route },
}

func FilterRoutes(array []pkg.ApiRoute, spec FilterSpec) ([]pkg.ApiRoute, error) {
	// Lookup filter function based on ParamType
	filterFunction, exists := filterFunctions[spec.ParamType]
	if !exists {
		return nil, fmt.Errorf("unknown filter type: %s", spec.ParamType)
	}

	// Filter the routes
	var filteredApiRoutes []pkg.ApiRoute
	for _, apiRoute := range array {
		if filterFunction(apiRoute) == spec.Param {
			filteredApiRoutes = append(filteredApiRoutes, apiRoute)
		}
	}

	return filteredApiRoutes, nil
}

func LoadJsonFromFile[T any](filePath string) (T, *os.File, error) {
	// Open the file
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("Failed to open file: %v\n", err)
		var zero T
		return zero, nil, err
	}

	// Read the current content of the file
	fileData, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("Failed to read file: %v\n", err)
		var zero T
		return zero, nil, err
	}

	// Initialize the target object
	var result T
	if len(fileData) > 0 {
		// Parse existing data
		err = json.Unmarshal(fileData, &result)
		if err != nil {
			fmt.Printf("Failed to parse JSON: %v\n", err)
			var zero T
			return zero, nil, err
		}
	}

	return result, file, nil
}

func LoadConfigJson() (pkg.Config, error) {
	filePath := "curl-echo/config.json"

	config, _, err := LoadJsonFromFile[pkg.Config](filePath)
	if err != nil {
		fmt.Printf("Failed to load config file: %v\n", err)
		return pkg.Config{}, err
	}
	return config, nil
}

func LoadApiJson() ([]pkg.ApiRoute, *os.File, error) {
	filePath := "curl-echo/apis.json"

	// Use the generic JSON loader
	apiRoutes, file, err := LoadJsonFromFile[[]pkg.ApiRoute](filePath)
	if err != nil {
		return nil, nil, err
	}

	return apiRoutes, file, nil
}

func CreateJson[T any](filePath string, data T, fileName string) {
	// Marshal the data into JSON
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Printf("Failed to marshal data to JSON: %v\n", err)
		return
	}

	// Construct the full file path
	path := filepath.Join(filePath, fileName)

	// Create the file
	file, err := os.Create(path)
	if err != nil {
		fmt.Printf("Failed to create file %s: %v\n", path, err)
		return
	}

	defer func() {
		if err := file.Close(); err != nil {
			fmt.Printf("Failed to close file %s: %v\n", path, err)
		}
	}()

	// Write the JSON data to the file
	_, err = file.Write(jsonData)
	if err != nil {
		fmt.Printf("Failed to write to file %s: %v\n", path, err)
		return
	}
	fmt.Printf("\nFile %s created successfully.\n", path)
}
