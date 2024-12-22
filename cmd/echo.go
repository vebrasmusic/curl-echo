/*
Copyright © 2024 Andres Duvvuri <hspncs@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/vebrasmusic/curl-echo/pkg"
	"github.com/vebrasmusic/curl-echo/pkg/util"
	"os"
)

var (
	routeToEcho    string
	nicknameToEcho string
	groupToEcho    string
)

// echoCmd represents the echo command
var echoCmd = &cobra.Command{
	Use:   "curl-echo echo [-r route] [-n nickname] [-g groupToEcho]",
	Short: "Run API routes and save their responses to files.",
	Long: `The "curl-echo echo" command runs API routes defined in the "apis.json" file and echoes their responses to files.

Usage:
- By default (no flags), it will run all routes in "apis.json" and save their responses.
- With the "-r route" flag, it will only run the specific route matching the provided route string.
- With the "-n nickname" flag, it will only run the route identified by the specified nickname.
- With the "-g groupToEcho" flag, it will run all routes associated with the specified group.
- You cannot chain these, ie. only use one of the flags or it will throw an error.

Examples:
1. Run all routes and save their responses:
   curl-echo echo

2. Run a specific route by its path:
   curl-echo echo -r "/api/v1/resource"

3. Run a specific route by its nickname:
   curl-echo echo -n "getResource"

4. Run all routes in a group:
   curl-echo echo -g "admin"`,
	Run: func(cmd *cobra.Command, args []string) {
		loading := true
		go util.ShowLoading(&loading)
		// if more than 1 is defined, throw error
		countFlags()
		// choose path based on what was chosen
		switch {
		case routeToEcho != "":
			fmt.Printf("Echoing route: %s...\n", routeToEcho)
		case nicknameToEcho != "":
			fmt.Printf("Echoing nickname: %s...\n", nicknameToEcho)
		case groupToEcho != "":
			fmt.Printf("Echoing group: %s...\n", group)
		default:
			fmt.Println("Echoing all available routes...")
		}

		loading = false
	},
}

func echo(apiRoutes []pkg.ApiRoute) {
	if len(apiRoutes) == 0 {
		fmt.Println("No routes to echo. Please add some using 'curl-echo add'.")
		os.Exit(0)
	}
	for _, apiRoute := range apiRoutes {
		// in curl-echo/echoes, create folder for group if not already there

		// create file titled {nickname}.json, or if it exists already override

		// run curl cmd w the api route

		// copy full response (headers, body, all that) into the json file

		// close and save, or if error then skip that one and write to echo.log. if it works, just write "route: xxx successfully echoed."

		// stop loading,

		// if timeout reached, exit w/ error timeout and log that too
	}
}

func countFlags() {
	// Count how many flags are set
	count := 0
	if routeToEcho != "" {
		count++
	}
	if nicknameToEcho != "" {
		count++
	}
	if groupToEcho != "" {
		count++
	}

	// If more than one flag is set, throw an error
	if count > 1 {
		fmt.Println("Error: Only one of 'route', 'nickname', or 'group' can be specified at a time.")
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(echoCmd)

	// Here you will define your flags and configuration settings.
	echoCmd.Flags().StringVarP(&routeToEcho, "route", "r", "", "Route to echo")
	echoCmd.Flags().StringVarP(&groupToEcho, "group", "g", "", "Group to echo")
	echoCmd.Flags().StringVarP(&nicknameToEcho, "nickname", "n", "", "Nickname to echo")

}
