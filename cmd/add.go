/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add new api routes to be echoed.",
	Long: `The "add" command allows you to add new API routes to be echoed by the curl-echo tool.

This command accepts an API route as an argument and then prompts you to configure additional details for the route, such as:
- HTTP Method (e.g., GET, POST, PUT, DELETE)
- Response content type (e.g., JSON, plain text)
- Response body

Example Usage:

  curl-echo add

This will start an interactive survey where you'll be asked about the HTTP method and other options for the new route. 
Once configured, the route will be added to the curl-echo registry, making it available for use.`,
	Run: func(cmd *cobra.Command, args []string) {

		// Ensure the directory exists
		dir := "curl-echo" // Example directory
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			fmt.Printf("Directory '%s' does not exist. Please run `curl-echo init` first.\n", dir)
			os.Exit(1) // Exit the program
		}

		fmt.Printf("Running add new route wizard...")
		apiRoute := runAddSurvey()
		addToApiJson(&apiRoute)
		fmt.Println("\nAPI route added successfully!")
	},
}

type ApiRoute struct {
	Nickname   string `json:"nickname"`
	Group      string `json:"group"`
	Route      string `json:"route"`
	HTTPMethod string `json:"http_method"`
}

func runAddSurvey() ApiRoute {
	questions := []*survey.Question{
		{
			Name:     "nickname",
			Prompt:   &survey.Input{Message: "Enter a nickname for this api: "},
			Validate: survey.Required,
		},
		{
			Name:     "group",
			Prompt:   &survey.Input{Message: "Enter the group this api should be long to: "},
			Validate: survey.Required,
		},
		{
			Name:     "route",
			Prompt:   &survey.Input{Message: "Enter the api route, including query params if needed: "},
			Validate: survey.Required,
		},
		{
			Name:     "httpMethod",
			Prompt:   &survey.Select{Message: "HTTP Method", Options: []string{"GET", "POST", "PUT", "DELETE"}},
			Validate: survey.Required,
		},
	}

	answers := struct {
		Nickname   string `survey:"nickname"`
		Group      string `survey:"group"`
		Route      string `survey:"route"`
		HTTPMethod string `survey:"httpMethod"`
	}{}

	err := survey.Ask(questions, &answers)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	apiRoute := ApiRoute{
		Nickname:   answers.Nickname,
		Group:      answers.Group,
		Route:      answers.Route,
		HTTPMethod: answers.HTTPMethod,
	}

	return apiRoute
}

func addToApiJson(apiRoute *ApiRoute) {
	const filePath = "curl-echo/apis.json"

	// Open the file
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("Failed to open file: %v\n", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Printf("Failed to close file: %v\n", err)
		}
	}(file)

	// Read the current content of the file
	fileData, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("Failed to read file: %v\n", err)
		return
	}

	// Initialize the list of API routes
	var apiRoutes []ApiRoute
	if len(fileData) > 0 {
		// Parse existing data
		err = json.Unmarshal(fileData, &apiRoutes)
		if err != nil {
			fmt.Printf("Failed to parse JSON: %v\n", err)
			return
		}
	}

	// Append the new API route
	apiRoutes = append(apiRoutes, *apiRoute)

	// Serialize the updated data back to JSON
	newData, err := json.MarshalIndent(apiRoutes, "", "  ")
	if err != nil {
		fmt.Printf("Failed to marshal JSON: %v\n", err)
		return
	}

	// Truncate the file before writing
	err = file.Truncate(0)
	if err != nil {
		fmt.Printf("Failed to truncate file: %v\n", err)
		return
	}

	// Reset the file cursor to the beginning
	_, err = file.Seek(0, 0)
	if err != nil {
		fmt.Printf("Failed to seek file: %v\n", err)
		return
	}

	// Write the updated JSON back to the file
	_, err = file.Write(newData)
	if err != nil {
		fmt.Printf("Failed to write to file: %v\n", err)
		return
	}
}

func init() {
	rootCmd.AddCommand(addCmd)
}
