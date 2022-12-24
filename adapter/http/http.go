// Package http (adapter) provides an HTTP client interface for the entire application
package http

import (
	"net/http"
)

type (
	// HttpClient is the http wrapper for the application
	HttpClient interface {
		HttpGetter
	}

	// HttpGetter holds fields and dependencies for executing an http GET request
	HttpGetter interface {
		// Get executes a GET http request
		Get(url string) (*http.Response, error)
	}
)
