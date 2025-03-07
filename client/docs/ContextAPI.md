# \ContextAPI

All URIs are relative to *https://us.app.unleash-hosted.com/ushosted*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateContextField**](ContextAPI.md#CreateContextField) | **Post** /api/admin/context | Create a context field
[**DeleteContextField**](ContextAPI.md#DeleteContextField) | **Delete** /api/admin/context/{contextField} | Delete an existing context field
[**GetContextField**](ContextAPI.md#GetContextField) | **Get** /api/admin/context/{contextField} | Gets context field
[**GetContextFields**](ContextAPI.md#GetContextFields) | **Get** /api/admin/context | Gets configured context fields
[**UpdateContextField**](ContextAPI.md#UpdateContextField) | **Put** /api/admin/context/{contextField} | Update an existing context field



## CreateContextField

> ContextFieldSchema CreateContextField(ctx).CreateContextFieldSchema(createContextFieldSchema).Execute()

Create a context field



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
	createContextFieldSchema := *openapiclient.NewCreateContextFieldSchema("subscriptionTier") // CreateContextFieldSchema | createContextFieldSchema

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContextAPI.CreateContextField(context.Background()).CreateContextFieldSchema(createContextFieldSchema).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContextAPI.CreateContextField``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateContextField`: ContextFieldSchema
	fmt.Fprintf(os.Stdout, "Response from `ContextAPI.CreateContextField`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateContextFieldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createContextFieldSchema** | [**CreateContextFieldSchema**](CreateContextFieldSchema.md) | createContextFieldSchema | 

### Return type

[**ContextFieldSchema**](ContextFieldSchema.md)

### Authorization

[bearerToken](../README.md#bearerToken), [apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteContextField

> DeleteContextField(ctx, contextField).Execute()

Delete an existing context field



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
	contextField := "contextField_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ContextAPI.DeleteContextField(context.Background(), contextField).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContextAPI.DeleteContextField``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contextField** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteContextFieldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerToken](../README.md#bearerToken), [apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContextField

> ContextFieldSchema GetContextField(ctx, contextField).Execute()

Gets context field



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
	contextField := "contextField_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContextAPI.GetContextField(context.Background(), contextField).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContextAPI.GetContextField``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContextField`: ContextFieldSchema
	fmt.Fprintf(os.Stdout, "Response from `ContextAPI.GetContextField`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contextField** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContextFieldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ContextFieldSchema**](ContextFieldSchema.md)

### Authorization

[bearerToken](../README.md#bearerToken), [apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContextFields

> []ContextFieldSchema GetContextFields(ctx).Execute()

Gets configured context fields



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContextAPI.GetContextFields(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContextAPI.GetContextFields``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContextFields`: []ContextFieldSchema
	fmt.Fprintf(os.Stdout, "Response from `ContextAPI.GetContextFields`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetContextFieldsRequest struct via the builder pattern


### Return type

[**[]ContextFieldSchema**](ContextFieldSchema.md)

### Authorization

[bearerToken](../README.md#bearerToken), [apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateContextField

> UpdateContextField(ctx, contextField).UpdateContextFieldSchema(updateContextFieldSchema).Execute()

Update an existing context field



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
	contextField := "contextField_example" // string | 
	updateContextFieldSchema := *openapiclient.NewUpdateContextFieldSchema() // UpdateContextFieldSchema | updateContextFieldSchema

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ContextAPI.UpdateContextField(context.Background(), contextField).UpdateContextFieldSchema(updateContextFieldSchema).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContextAPI.UpdateContextField``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contextField** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateContextFieldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateContextFieldSchema** | [**UpdateContextFieldSchema**](UpdateContextFieldSchema.md) | updateContextFieldSchema | 

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

