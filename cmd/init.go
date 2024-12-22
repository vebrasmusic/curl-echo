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
		go util.ShowLoading(&loading)
		createDirectories()
		fmt.Println("\nInitialization complete!")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func runInitSurvey() {
	question := &survey.Confirm{
		Message: `Ready to init curl-echo in your project? This will create a 'curl-echo' directory.
If you want to remove 'curl-echo' from your project in the future, you can with the cmd 'curl-echo rm'`,
	}
	var confirmation bool

	// Run the survey
	err := survey.AskOne(question, &confirmation)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	if confirmation {
		fmt.Println("Confirmed. Initializing curl-echo...")
		return
	}

	fmt.Println("Aborted. curl-echo was not initialized.")
	os.Exit(1)
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
	apiFilePath := filepath.Join(cwd, "curl-echo")
	util.CreateJson(apiFilePath, "[]", "apis.json")
}
