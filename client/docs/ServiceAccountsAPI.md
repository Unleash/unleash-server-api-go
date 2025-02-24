# \ServiceAccountsAPI

All URIs are relative to *https://us.app.unleash-hosted.com/ushosted*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateServiceAccount**](ServiceAccountsAPI.md#CreateServiceAccount) | **Post** /api/admin/service-account | Create a service account.
[**CreateServiceAccountToken**](ServiceAccountsAPI.md#CreateServiceAccountToken) | **Post** /api/admin/service-account/{id}/token | Create a token for a service account.
[**DeleteServiceAccount**](ServiceAccountsAPI.md#DeleteServiceAccount) | **Delete** /api/admin/service-account/{id} | Delete a service account.
[**DeleteServiceAccountToken**](ServiceAccountsAPI.md#DeleteServiceAccountToken) | **Delete** /api/admin/service-account/{id}/token/{tokenId} | Delete a token for a service account.
[**GetServiceAccountTokens**](ServiceAccountsAPI.md#GetServiceAccountTokens) | **Get** /api/admin/service-account/{id}/token | List all tokens for a service account.
[**GetServiceAccounts**](ServiceAccountsAPI.md#GetServiceAccounts) | **Get** /api/admin/service-account | List service accounts.
[**UpdateServiceAccount**](ServiceAccountsAPI.md#UpdateServiceAccount) | **Put** /api/admin/service-account/{id} | Update a service account.



## CreateServiceAccount

> ServiceAccountSchema CreateServiceAccount(ctx).CreateServiceAccountSchema(createServiceAccountSchema).Execute()

Create a service account.



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
	createServiceAccountSchema := *openapiclient.NewCreateServiceAccountSchema("service-account-1", "Service Account 1", int32(1)) // CreateServiceAccountSchema | #/components/schemas/createServiceAccountSchema

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceAccountsAPI.CreateServiceAccount(context.Background()).CreateServiceAccountSchema(createServiceAccountSchema).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountsAPI.CreateServiceAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateServiceAccount`: ServiceAccountSchema
	fmt.Fprintf(os.Stdout, "Response from `ServiceAccountsAPI.CreateServiceAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateServiceAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createServiceAccountSchema** | [**CreateServiceAccountSchema**](CreateServiceAccountSchema.md) | #/components/schemas/createServiceAccountSchema | 

### Return type

[**ServiceAccountSchema**](ServiceAccountSchema.md)

### Authorization

[bearerToken](../README.md#bearerToken), [apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateServiceAccountToken

> PatSchema CreateServiceAccountToken(ctx, id).CreatePatSchema(createPatSchema).Execute()

Create a token for a service account.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/Unleash/unleash-server-api-go"
)

func main() {
	id := "id_example" // string | 
	createPatSchema := *openapiclient.NewCreatePatSchema("user:xyzrandomstring", time.Now()) // CreatePatSchema | #/components/schemas/createPatSchema

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceAccountsAPI.CreateServiceAccountToken(context.Background(), id).CreatePatSchema(createPatSchema).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountsAPI.CreateServiceAccountToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateServiceAccountToken`: PatSchema
	fmt.Fprintf(os.Stdout, "Response from `ServiceAccountsAPI.CreateServiceAccountToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateServiceAccountTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createPatSchema** | [**CreatePatSchema**](CreatePatSchema.md) | #/components/schemas/createPatSchema | 

### Return type

[**PatSchema**](PatSchema.md)

### Authorization

[bearerToken](../README.md#bearerToken), [apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteServiceAccount

> DeleteServiceAccount(ctx, id).Execute()

Delete a service account.



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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ServiceAccountsAPI.DeleteServiceAccount(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountsAPI.DeleteServiceAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteServiceAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerToken](../README.md#bearerToken), [apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteServiceAccountToken

> DeleteServiceAccountToken(ctx, id, tokenId).Execute()

Delete a token for a service account.



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
	id := "id_example" // string | 
	tokenId := "tokenId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ServiceAccountsAPI.DeleteServiceAccountToken(context.Background(), id, tokenId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountsAPI.DeleteServiceAccountToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**tokenId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteServiceAccountTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[bearerToken](../README.md#bearerToken), [apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceAccountTokens

> PatsSchema GetServiceAccountTokens(ctx, id).Execute()

List all tokens for a service account.



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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceAccountsAPI.GetServiceAccountTokens(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountsAPI.GetServiceAccountTokens``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServiceAccountTokens`: PatsSchema
	fmt.Fprintf(os.Stdout, "Response from `ServiceAccountsAPI.GetServiceAccountTokens`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceAccountTokensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PatsSchema**](PatsSchema.md)

### Authorization

[bearerToken](../README.md#bearerToken), [apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceAccounts

> ServiceAccountsSchema GetServiceAccounts(ctx).Execute()

List service accounts.



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
	resp, r, err := apiClient.ServiceAccountsAPI.GetServiceAccounts(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountsAPI.GetServiceAccounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServiceAccounts`: ServiceAccountsSchema
	fmt.Fprintf(os.Stdout, "Response from `ServiceAccountsAPI.GetServiceAccounts`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceAccountsRequest struct via the builder pattern


### Return type

[**ServiceAccountsSchema**](ServiceAccountsSchema.md)

### Authorization

[bearerToken](../README.md#bearerToken), [apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateServiceAccount

> ServiceAccountSchema UpdateServiceAccount(ctx, id).UpdateServiceAccountSchema(updateServiceAccountSchema).Execute()

Update a service account.



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
	id := "id_example" // string | 
	updateServiceAccountSchema := *openapiclient.NewUpdateServiceAccountSchema() // UpdateServiceAccountSchema | #/components/schemas/updateServiceAccountSchema

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceAccountsAPI.UpdateServiceAccount(context.Background(), id).UpdateServiceAccountSchema(updateServiceAccountSchema).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountsAPI.UpdateServiceAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateServiceAccount`: ServiceAccountSchema
	fmt.Fprintf(os.Stdout, "Response from `ServiceAccountsAPI.UpdateServiceAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateServiceAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateServiceAccountSchema** | [**UpdateServiceAccountSchema**](UpdateServiceAccountSchema.md) | #/components/schemas/updateServiceAccountSchema | 

### Return type

[**ServiceAccountSchema**](ServiceAccountSchema.md)

### Authorization

[bearerToken](../README.md#bearerToken), [apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

