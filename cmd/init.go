/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
	"github.com/vebrasmusic/curl-echo/pkg/util"
	"os"
	"path/filepath"
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

		runInitSurvey()
		loading := true
		fmt.Print("Initializing curl-echo")
		go util.ShowLoading(&loading)
		createDirectories()
		fmt.Println("Initialization complete!")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func runInitSurvey() {
	questions := []*survey.Question{
		{
			Name:     "affirmation",
			Prompt:   &survey.Select{Message: "Init curl-echo in your project?:", Options: []string{"yes", "no"}},
			Validate: survey.Required,
		},
	}

	answers := struct {
		Affirmation string
	}{}

	err := survey.Ask(questions, &answers)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if answers.Affirmation == "no" {
		fmt.Println("Initialization aborted by user.")
		os.Exit(0)
	}
}

func createDirectories() {

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
	apiFilePath := filepath.Join(cwd, "/curl-echo/apis.json")
	file, err := os.Create(apiFilePath)
	if err != nil {
		fmt.Printf("Failed to create file %s: %v\n", apiFilePath, err)
		return
	}
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Printf("Failed to close file %s: %v\n", apiFilePath, err)
		}
	}()

	content := `[]`
	_, err = file.Write([]byte(content))
	if err != nil {
		fmt.Printf("Failed to write to file %s: %v\n", apiFilePath, err)
		return
	}
	fmt.Printf("\nApi spec file created")
}
