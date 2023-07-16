package main

import (
	"context"
	"fmt"
	openapiclient "github.com/Unleash/unleash-server-api-go/client"
	"net/http"
	"net/http/httputil"
	"os"
)

func main() {
	name := "A name"
	email := "test@getunleash.io"
	username := "myusername"
	sendEmail := false
	rootRoleId := int32(1)
	rootRole := openapiclient.Int32AsCreateUserSchemaRootRole(&rootRoleId)

	createUserSchema := *openapiclient.NewCreateUserSchemaWithDefaults()
	createUserSchema.Name = &name
	createUserSchema.Email = &email
	createUserSchema.Username = &username
	createUserSchema.RootRole = rootRole
	createUserSchema.SendEmail = &sendEmail
	fmt.Println(createUserSchema)

	configuration := openapiclient.NewConfiguration()
	configuration.HTTPClient = &http.Client{
		Transport: &debugTransport{
			Transport:   http.DefaultTransport,
			EnableDebug: true,
		},
	}
	configuration.Servers = openapiclient.ServerConfigurations{
		{
			URL: "http://localhost:3000",
		},
	}
	configuration.AddDefaultHeader("Authorization", "*:*.964a287e1b728cb5f4f3e0120df92cb5")
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
