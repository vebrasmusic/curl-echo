package util

import "github.com/vebrasmusic/curl-echo/pkg"

func FilterApiRoutes(apiRoutes []pkg.ApiRoute, condition func(pkg.ApiRoute) bool) ([]pkg.ApiRoute, error) {
	var filtered []pkg.ApiRoute
	for _, apiRoute := range apiRoutes {
		if condition(apiRoute) {
			filtered = append(filtered, apiRoute)
		}
	}
	return filtered, nil
}
