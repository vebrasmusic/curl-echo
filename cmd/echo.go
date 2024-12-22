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
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/spf13/cobra"
	"github.com/vebrasmusic/curl-echo/pkg"
	"github.com/vebrasmusic/curl-echo/pkg/util"
	"log"
	"os"
	"path/filepath"
)

var (
	routeToEcho    string
	nicknameToEcho string
	groupToEcho    string
)

var client = resty.New()

// echoCmd represents the echo command
var echoCmd = &cobra.Command{
	Use:   "echo [-r route] [-n nickname] [-g groupToEcho]",
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
		apiRoutes, _, err := util.LoadApiJson()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		switch {
		case routeToEcho != "":
			fmt.Printf("Echoing route: %s...\n", routeToEcho)
			filteredRoutes, _ := util.FilterApiRoutes(apiRoutes, func(r pkg.ApiRoute) bool {
				return r.Route == routeToEcho
			})
			echo(filteredRoutes)
		case nicknameToEcho != "":
			fmt.Printf("Echoing nickname: %s...\n", nicknameToEcho)
			filteredRoutes, _ := util.FilterApiRoutes(apiRoutes, func(r pkg.ApiRoute) bool {
				return r.Nickname == nicknameToEcho
			})
			echo(filteredRoutes)
		case groupToEcho != "":
			fmt.Printf("Echoing group: %s...\n", groupToEcho)
			filteredRoutes, _ := util.FilterApiRoutes(apiRoutes, func(r pkg.ApiRoute) bool {
				return r.Group == groupToEcho
			})
			echo(filteredRoutes)
		default:
			fmt.Println("Echoing all available routes...")
			echo(apiRoutes)
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
		cwd, _ := os.Getwd()
		path := filepath.Join(cwd, "curl-echo", "echoes")
		if apiRoute.Group != "" {
			path += "/" + apiRoute.Group
			folderPath := "echoes/" + apiRoute.Group
			util.CreateFolders(folderPath)
		}
		filePath := apiRoute.Nickname + ".json"

		// TODO: implement for other HTTP methods
		response, err := runGetRequest(apiRoute)
		if err != nil {
			fmt.Println("Error while executing request:", err)
			// add to log here
			continue
		}
		responseContent, err := parseHttp(response)
		if err != nil {
			fmt.Println("Error while parsing response:", err)
			continue
		}
		util.CreateJson(path, responseContent, filePath)
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

func constructCompleteRoute(apiRoute pkg.ApiRoute) string {
	// read config.json.rootApiPath
	config, err := util.LoadConfigJson()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return config.RootApiPath + apiRoute.Route
}

func runGetRequest(apiRoute pkg.ApiRoute) (*resty.Response, error) {
	// Create a Resty Client
	// TODO: add reading the config to find the root path
	resp, err := client.R().
		EnableTrace().
		Get(constructCompleteRoute(apiRoute))

	if err != nil {
		fmt.Println("Error:", err)
		return nil, fmt.Errorf("failed to make GET request: %w", err)
	}
	return resp, nil
}

func parseHttp(resp *resty.Response) (pkg.ResponseContent, error) {
	// Populate the ResponseContent struct
	content := pkg.ResponseContent{
		StatusCode: resp.StatusCode(),
		Body:       make(map[string]interface{}),
		Headers:    make(map[string]string),
	}

	// Parse the response body into a map
	err := json.Unmarshal(resp.Body(), &content.Body)
	if err != nil {
		log.Printf("Error unmarshalling response body: %v", err)
		return pkg.ResponseContent{}, err
	}

	// Populate headers
	for k, v := range resp.Header() {
		if len(v) > 0 {
			content.Headers[k] = v[0] // Use the first value if multiple
		}
	}

	return content, nil
}

func init() {
	rootCmd.AddCommand(echoCmd)

	// Here you will define your flags and configuration settings.
	echoCmd.Flags().StringVarP(&routeToEcho, "route", "r", "", "Route to echo")
	echoCmd.Flags().StringVarP(&groupToEcho, "group", "g", "", "Group to echo")
	echoCmd.Flags().StringVarP(&nicknameToEcho, "nickname", "n", "", "Nickname to echo")

}
