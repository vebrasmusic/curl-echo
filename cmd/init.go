/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

func showLoading(loading *bool) {
	for *loading {
		for _, r := range `-\|/` {
			fmt.Printf("\rInitializing curl-echo... %c", r)
			time.Sleep(100 * time.Millisecond)
		}
	}
}

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

		// Use bufio for input handling
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Folder path for storing echoed responses? [default: /echoes]: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}

		// Trim input and set default if empty
		echoPath := strings.TrimSpace(input)
		if echoPath == "" {
			echoPath = "/echoes"
		}

		// Display initializing message with loading animation
		loading := true
		fmt.Print("Initializing curl-echo")
		go showLoading(&loading)

		// Get the current working directory
		cwd, err := os.Getwd()
		if err != nil {
			fmt.Printf("Failed to get current working directory: %v\n", err)
			return
		}

		// Create the directory for echoed responses
		echoDirPath := filepath.Join(cwd, echoPath) // Ensure path is relative to CWD
		err = os.MkdirAll(echoDirPath, os.ModePerm)
		if err != nil {
			fmt.Printf("\nFailed to create directory %s: %v\n", echoDirPath, err)
			return
		}
		fmt.Printf("\nDirectory created at %s\n", echoDirPath)

		// Create and write to the configuration file in CWD
		configFilePath := filepath.Join(cwd, "curl-echo.config.json")
		file, err := os.Create(configFilePath)
		if err != nil {
			fmt.Printf("Failed to create file %s: %v\n", configFilePath, err)
			return
		}
		defer func(file *os.File) {
			err := file.Close()
			if err != nil {
				if err != nil {
					fmt.Printf("Failed to create file %s: %v\n", configFilePath, err)
					return
				}
			}
		}(file)

		content := `{
    "echoPath": "` + echoPath + `",
    "configVersion": "1.0.0",
}`
		_, err = file.Write([]byte(content))
		if err != nil {
			fmt.Printf("Failed to write to file %s: %v\n", configFilePath, err)
			return
		}
		fmt.Printf("Config file created at %s\n", configFilePath)
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
