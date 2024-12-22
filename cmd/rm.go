/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"os"

	"github.com/spf13/cobra"
)

// rmCmd represents the rm command
var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Delete the curl-echo folder and all its contents.",
	Long: `The "rm" command permanently deletes the curl-echo folder and all its contents. 

Warning: This action cannot be reversed. Use with caution.

Example:
  curl-echo rm
  # Prompts: "Are you sure? This cannot be reversed."
  # If confirmed, deletes the folder.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			fmt.Println("Too many arguments. Run again with no arguments.")
			os.Exit(1)
		}
		fmt.Println("Running uninstall...")
		runDeleteSurvey()
	},
}

func runDeleteSurvey() {
	question := &survey.Confirm{
		Message: "Are you sure you want to remove curl-echo from your project? This cannot be reversed.",
	}
	var confirmation bool

	// Run the survey
	err := survey.AskOne(question, &confirmation)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	if confirmation {
		fmt.Println("Confirmed. Deleting curl-echo folder...")
		err := os.RemoveAll("./curl-echo") // Update the folder path as needed
		if err != nil {
			fmt.Printf("Failed to delete folder: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("The curl-echo folder and its contents have been deleted.")
		os.Exit(0)
	}

	fmt.Println("Aborted. The curl-echo folder was not deleted.")
	os.Exit(1)
}

func init() {
	rootCmd.AddCommand(rmCmd)
}
