/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/vebrasmusic/curl-echo/pkg/util"

	"github.com/spf13/cobra"
)

var group string

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list [-g group]",
	Short: "Display all available routes or filter by a specific group.",
	Long: `The "list" command allows you to view all the routes configured in the "curl-echo" service.

By default, running "curl-echo list" returns all available routes, including their nicknames, endpoints, and HTTP methods. This provides an overview of all exposed APIs in your project.

If you want to narrow down the output to routes within a specific group, you can use the "-g" flag followed by the group name. This is useful for filtering large numbers of routes into manageable categories.

Examples:

1. List all routes:
   $ curl-echo list

   Output:
   nickname: get-all-studies
   route: /api/proxy/get_all_studies
   httpMethod: GET

   nickname: create-study
   route: /api/proxy/create_study
   httpMethod: POST

2. List routes for a specific group:
   $ curl-echo list -g studies

   Output:
   nickname: get-all-studies
   route: /api/proxy/get_all_studies
   httpMethod: GET

   nickname: create-study
   route: /api/proxy/create_study
   httpMethod: POST

Usage:
- "curl-echo list" to view all routes.
- "curl-echo list -g <group>" to filter by group.

Replace "<group>" with the desired group name, e.g., "studies".

Ensure your configuration defines valid groups for accurate filtering.`,
	Run: func(cmd *cobra.Command, args []string) {
		apiRoutes, _ := util.LoadApiJson()
		// Filter routes by group if the --group flag is provided
		if group != "" {
			filterSpec := util.FilterSpec{
				Param:     group,
				ParamType: "Group",
			}
			apiRoutes = util.FilterRoutes(apiRoutes, filterSpec)
		}
		if (apiRoutes == nil) || (len(apiRoutes) == 0) {
			fmt.Println("No routes found. Try 'curl-echo add' to add some new routes.")
			return
		}
		fmt.Println("Available routes:")
		for _, apiRoute := range apiRoutes {
			fmt.Printf("Nickname:   %s\n", apiRoute.Nickname)
			fmt.Printf("Route:      %s\n", apiRoute.Route)
			fmt.Printf("HTTPMethod: %s\n", apiRoute.HTTPMethod)
			fmt.Printf("Group:      %s\n", apiRoute.Group)
			fmt.Println("---")
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().StringVarP(&group, "group", "g", "", "Filter by group name")

}
