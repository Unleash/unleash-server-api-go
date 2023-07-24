package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"net/http/httputil"
	"os"

	openapiclient "github.com/Unleash/unleash-server-api-go/client"
)

func main() {
	unleashApiUrl := flag.String("unleash-api", "", "Specifies the Unleash API URL (e.g. http://localhost:4242 if you're running Unleash locally)")
	authorization := flag.String("authorization", "", "Token to use on Authorization header")
	enableDebug := flag.Bool("debug-http", false, "Enable HTTP debug logging")

	flag.Parse()

	name := "A name"
	email := "tester1@getunleash.io"
	username := "username1"
	sendEmail := false
	rootRoleId := float32(1)

	createUserSchema := *openapiclient.NewCreateUserSchemaWithDefaults()
	createUserSchema.Name = &name
	createUserSchema.Email = &email
	createUserSchema.Username = &username
	createUserSchema.RootRole = rootRoleId
	createUserSchema.SendEmail = &sendEmail

	configuration := openapiclient.NewConfiguration()
	configuration.HTTPClient = &http.Client{
		Transport: &debugTransport{
			Transport:   http.DefaultTransport,
			EnableDebug: *enableDebug,
		},
	}
	configuration.Servers = openapiclient.ServerConfigurations{
		{
			URL: *unleashApiUrl,
		},
	}
	configuration.AddDefaultHeader("Authorization", *authorization)
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersApi.CreateUser(context.Background()).CreateUserSchema(createUserSchema).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.CreateUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateUser`: CreateUserResponseSchema
	fmt.Fprintf(os.Stdout, "Response from `UsersApi.CreateUser`: %v\n", resp)
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

	fmt.Printf("Err:\n%s\n\n", err)
	if t.EnableDebug && resp != nil {
		// Log the response details
		responseDump, _ := httputil.DumpResponse(resp, true)
		fmt.Printf("Response:\n%s\n\n", responseDump)
	}

	return resp, err
}
