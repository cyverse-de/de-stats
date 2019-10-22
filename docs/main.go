// Package docs DE Statistics API
//
// Documentation of the DE Stats API
//
//     Schemes: http
//     BasePath: /
//     Version: 1.0.0
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
package docs

import "github.com/cyverse-de/de-stats/api"

// swagger:route GET / misc getRoot
// Returns general information about the API.
// responses:
//   200: rootResponse

// General information about the API.
// swagger:response rootResponse
type rootResponseWrapper struct {
	// in:body
	Body api.RootResponse
}

type appsResponseWrapper struct {
	Body api.AppsResponse
}
