package main

import "fmt"

import "github.com/Unleash/unleash-server-api-go/swagger"
import (
	"context"
	"flag"
	"net/http"
	"net/http/httputil"
)

func main() {
	enableDebug := flag.Bool("debug-http", false, "Enable HTTP debug logging")

	flag.Parse()

	cfg := swagger.NewConfiguration()

	// Initialize the HTTP client transport request and response logging
	cfg.HTTPClient = &http.Client{
		Transport: &debugTransport{
			Transport:   http.DefaultTransport,
			EnableDebug: *enableDebug,
		},
	}

	cfg.AddDefaultHeader("Authorization", "user:a3c8e3e76e7361c0bc79070accf70409d03fbaef4c2b6b90bff466e8")
	cfg.BasePath = "http://localhost:4242"
	apiClient := swagger.NewAPIClient(cfg)

	ctx := context.Background()
	projects, _, err := apiClient.ProjectsApi.GetProjects(ctx)
	fmt.Println(err)
	users, _, err := apiClient.UsersApi.GetUsers(ctx)
	fmt.Println(err)
	fmt.Println(projects)
	fmt.Println(users)
}

// Custom transport for request and response logging
type debugTransport struct {
	Transport   http.RoundTripper
	EnableDebug bool
}

func (t *debugTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if t.EnableDebug {
		// Log the request details
		requestDump, _ := httputil.DumpRequestOut(req, true)
		fmt.Printf("Request:\n%s\n\n", requestDump)
	}

	// Make the actual request
	resp, err := t.Transport.RoundTrip(req)

	if t.EnableDebug {
		// Log the response details
		responseDump, _ := httputil.DumpResponse(resp, true)
		fmt.Printf("Response:\n%s\n\n", responseDump)
	}

	return resp, err
}
