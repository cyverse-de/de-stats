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

// swagger:route GET /apps GetTopApps
// Returns the top apps in the given time period.
// Responses:
// 200: appsResponse
// 400: errorResponse

//Top apps in time period.
//swagger:response appsResponse
type appsResponseWrapper struct {
	// in:body
	Body api.AppsResponse
}

// swagger:route GET /users GetTopUsers
// Returns the top users in the given time period.
// Responses:
// 200: usersResponse
// 400: errorResponse

//Top users in time period
//swagger:response usersResponse
type usersResponseWrapper struct {
	// in:body
	Body api.UsersResponse
}

// swagger:route GET /jobs/submitted GetSubmittedJobCounts
// Returns the number of submitted jobs in the given time period for each job type.
// Responses:
// 200: usersResponse
// 400: errorResponse

//Top users in time period
//swagger:response usersResponse
type jobsResponseWrapper struct {
	// in:body
	Body api.JobsResponse
}