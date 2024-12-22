/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
	"github.com/vebrasmusic/curl-echo/pkg"
	"github.com/vebrasmusic/curl-echo/pkg/util"
	"os"
	"path/filepath"
	"strconv"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize curl-echo and add to project.",
	Long:  `Initializes curl-echo for your project. Defaults are listed.`,
	Run: func(cmd *cobra.Command, args []string) {
		asciiTitle := `
                 _                 _           
  ___ _   _ _ __| |       ___  ___| |__   ___  
 / __| | | | '__| |_____ / _ \/ __| '_ \ / _ \ 
| (__| |_| | |  | |_____|  __/ (__| | | | (_) |
 \___|\__,_|_|  |_|      \___|\___|_| |_|\___/

`
		fmt.Print(asciiTitle)
		fmt.Println("made with <3 by andrés")

		config := runInitSurvey()
		loading := true
		go util.ShowLoading(&loading)
		createFiles(config)
		fmt.Println("\nInitialization complete!")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func runInitSurvey() pkg.Config {
	// Confirmation question
	confirmation := false
	confirmationQuestion := &survey.Confirm{
		Message: `Ready to init curl-echo in your project? This will create a 'curl-echo' directory.`,
	}

	// Ask the confirmation question
	err := survey.AskOne(confirmationQuestion, &confirmation)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	// If not confirmed, abort
	if !confirmation {
		fmt.Println("Aborted. curl-echo was not initialized.")
		os.Exit(1)
	}

	// Input question
	var rootApiPath string
	inputQuestion := &survey.Input{
		Message: "What's the root path of your API? (ie. http://localhost:8080/api): ",
	}
	err = survey.AskOne(inputQuestion, &rootApiPath, survey.WithValidator(survey.Required))
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	var maxEchoTimeout int
	defaultTimeout := 20

	timeoutQuestion := &survey.Input{
		Message: fmt.Sprintf("Set the max timeout (secs) for an echo response (default: %d): ", defaultTimeout),
	}

	var timeoutInput string
	err1 := survey.AskOne(timeoutQuestion, &timeoutInput)
	if err1 != nil {
		fmt.Printf("Error: %v\n", err1)
		os.Exit(1)
	}

	// Handle user input: if empty, use default value; if not, parse to int
	if timeoutInput == "" {
		maxEchoTimeout = defaultTimeout
	} else {
		maxEchoTimeout, err = strconv.Atoi(timeoutInput)
		if err != nil {
			fmt.Printf("Invalid input. Please enter a number.\n")
			os.Exit(1)
		}
	}

	fmt.Println("Confirmed. Initializing curl-echo with root API path:", rootApiPath)
	return pkg.Config{
		RootApiPath:    rootApiPath,
		MaxEchoTimeout: maxEchoTimeout,
	}
}

func createFiles(config pkg.Config) {

	folders := []string{
		"echoes",
	}

	for index, folder := range folders {
		fmt.Println(index, folder)
		util.CreateFolders(folder)
	}
	fmt.Printf("\nFolder structure created")
	cwd, _ := os.Getwd()

	// Create and write to the configuration file in CWD
	filePath := filepath.Join(cwd, "curl-echo")
	apiRoutes := []pkg.ApiRoute{}
	util.CreateJson(filePath, apiRoutes, "apis.json")

	util.CreateJson(filePath, config, "config.json")
}
