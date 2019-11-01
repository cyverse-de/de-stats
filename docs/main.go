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
//   400: errorResponse

// General information about the API.
// swagger:response rootResponse
type rootResponseWrapper struct {
	// in:body
	Body api.RootResponse
}

// Basic error response.
// swagger:response errorResponse
type errorResponseWrapper struct {
	// in:body
	Body api.ErrorResponse
}
