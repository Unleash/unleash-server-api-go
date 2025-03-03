# \ChangeRequestsAPI

All URIs are relative to *https://us.app.unleash-hosted.com/ushosted*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetProjectChangeRequestConfig**](ChangeRequestsAPI.md#GetProjectChangeRequestConfig) | **Get** /api/admin/projects/{projectId}/change-requests/config | Retrieves change request configuration for a project
[**UpdateProjectChangeRequestConfig**](ChangeRequestsAPI.md#UpdateProjectChangeRequestConfig) | **Put** /api/admin/projects/{projectId}/environments/{environment}/change-requests/config | Updates change request configuration for an environment in the project



## GetProjectChangeRequestConfig

> []ChangeRequestEnvironmentConfigSchema GetProjectChangeRequestConfig(ctx, projectId).Execute()

Retrieves change request configuration for a project



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Unleash/unleash-server-api-go"
)

func main() {
	projectId := "projectId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChangeRequestsAPI.GetProjectChangeRequestConfig(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChangeRequestsAPI.GetProjectChangeRequestConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectChangeRequestConfig`: []ChangeRequestEnvironmentConfigSchema
	fmt.Fprintf(os.Stdout, "Response from `ChangeRequestsAPI.GetProjectChangeRequestConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectChangeRequestConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ChangeRequestEnvironmentConfigSchema**](ChangeRequestEnvironmentConfigSchema.md)

### Authorization

[bearerToken](../README.md#bearerToken), [apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProjectChangeRequestConfig

> UpdateProjectChangeRequestConfig(ctx, projectId, environment).UpdateChangeRequestEnvironmentConfigSchema(updateChangeRequestEnvironmentConfigSchema).Execute()

Updates change request configuration for an environment in the project



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Unleash/unleash-server-api-go"
)

func main() {
	projectId := "projectId_example" // string | 
	environment := "environment_example" // string | 
	updateChangeRequestEnvironmentConfigSchema := *openapiclient.NewUpdateChangeRequestEnvironmentConfigSchema(false) // UpdateChangeRequestEnvironmentConfigSchema | updateChangeRequestEnvironmentConfigSchema

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ChangeRequestsAPI.UpdateProjectChangeRequestConfig(context.Background(), projectId, environment).UpdateChangeRequestEnvironmentConfigSchema(updateChangeRequestEnvironmentConfigSchema).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChangeRequestsAPI.UpdateProjectChangeRequestConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 
**environment** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProjectChangeRequestConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateChangeRequestEnvironmentConfigSchema** | [**UpdateChangeRequestEnvironmentConfigSchema**](UpdateChangeRequestEnvironmentConfigSchema.md) | updateChangeRequestEnvironmentConfigSchema | 

### Return type

 (empty response body)

### Authorization

[bearerToken](../README.md#bearerToken), [apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

