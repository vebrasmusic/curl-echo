/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize curl-echo and add to project.",
	Long:  `Initializes curl-echo for your project. Defaults are listed.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("curl-echo made with <3 by Andrés")
		fmt.Print("Echoed responses folder: ")
		fmt.Scanln() // Waits for enter and any text input
		fmt.Print("Initializing curl-echo")

		// Animate the dots for a loading effect
		for i := 0; i < 3; i++ {
			time.Sleep(500 * time.Millisecond) // Wait half a second
			fmt.Print(".")
		}
		fmt.Println("\nInitialization complete!")
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
