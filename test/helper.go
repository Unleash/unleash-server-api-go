package client

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"os"

	"github.com/Unleash/unleash-server-api-go/client"
)

func enterpriseEnvironmentAvailable() bool {
	return os.Getenv("UNLEASH_ENTERPRISE") == "true"
}

func testClient() *client.APIClient {
	configuration := client.NewConfiguration()
	configuration.Servers = client.ServerConfigurations{
		{
			URL: "http://localhost:4242",
		},
	}
	configuration.HTTPClient = &http.Client{
		Transport: &debugTransport{
			Transport:   http.DefaultTransport,
			EnableDebug: false,
		},
	}
	configuration.AddDefaultHeader("Authorization", "*:*.unleash-insecure-admin-api-token")
	return client.NewAPIClient(configuration)
}

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

	if err != nil {
		fmt.Printf("Err:\n%s\n\n", err)
	}

	if t.EnableDebug && resp != nil {
		// Log the response details
		responseDump, _ := httputil.DumpResponse(resp, true)
		fmt.Printf("Response:\n%s\n\n", responseDump)
	}

	return resp, err
}
