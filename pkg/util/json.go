package util

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/vebrasmusic/curl-echo/pkg"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
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

func LoadApiJson() ([]pkg.ApiRoute, *os.File, error) {
	const filePath = "curl-echo/apis.json"

	// Open the file
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("Failed to open file: %v\n", err)
		return nil, nil, err
	}

	// Read the current content of the file
	fileData, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("Failed to read file: %v\n", err)
		return nil, nil, err
	}

	// Initialize the list of API routes
	var apiRoutes []pkg.ApiRoute
	if len(fileData) > 0 {
		// Parse existing data
		err = json.Unmarshal(fileData, &apiRoutes)
		if err != nil {
			fmt.Printf("Failed to parse JSON: %v\n", err)
			return nil, nil, err
		}
	}

	return apiRoutes, file, nil
}

func CreateJson(filePath string, content string, fileName string) {
	path := filepath.Join(filePath, fileName)
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

	_, err = file.Write([]byte(content))
	if err != nil {
		fmt.Printf("Failed to write to file %s: %v\n", path, err)
		return
	}
	fmt.Printf("\nFile %s created\n", path)
}

type ResponseContent struct {
	StatusCode int                    `json:"statusCode"`
	Body       map[string]interface{} `json:"body"`
	Headers    map[string]string      `json:"headers"`
}

func ParseHttpToJson(resp *resty.Response) (string, error) {
	// Populate the ResponseContent struct
	content := ResponseContent{
		StatusCode: resp.StatusCode(),
		Body:       make(map[string]interface{}),
		Headers:    make(map[string]string),
	}

	// Parse the response body into a map
	err := json.Unmarshal(resp.Body(), &content.Body)
	if err != nil {
		log.Fatalf("Error unmarshalling response body: %v", err)
		return "", err
	}

	// Populate headers
	for k, v := range resp.Header() {
		content.Headers[k] = v[0] // Use the first value if multiple
	}

	// Convert the struct to JSON
	finalJSON, err := json.MarshalIndent(content, "", "  ")
	if err != nil {
		log.Fatalf("Error marshalling JSON: %v", err)
		return "", err
	}

	// Print the result
	return string(finalJSON), nil
}
