# \AuthAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOidcSettings**](AuthAPI.md#GetOidcSettings) | **Get** /api/admin/auth/oidc/settings | Get OIDC auth settings
[**GetPermissions**](AuthAPI.md#GetPermissions) | **Get** /api/admin/permissions | Gets available permissions
[**GetSimpleSettings**](AuthAPI.md#GetSimpleSettings) | **Get** /api/admin/auth/simple/settings | Get Simple auth settings
[**SetOidcSettings**](AuthAPI.md#SetOidcSettings) | **Post** /api/admin/auth/oidc/settings | Set OIDC settings
[**SetSimpleSettings**](AuthAPI.md#SetSimpleSettings) | **Post** /api/admin/auth/simple/settings | Update Simple auth settings



## GetOidcSettings

> OidcSettingsSchema GetOidcSettings(ctx).Execute()

Get OIDC auth settings



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
	resp, r, err := apiClient.AuthAPI.GetOidcSettings(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.GetOidcSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOidcSettings`: OidcSettingsSchema
	fmt.Fprintf(os.Stdout, "Response from `AuthAPI.GetOidcSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetOidcSettingsRequest struct via the builder pattern


### Return type

[**OidcSettingsSchema**](OidcSettingsSchema.md)

### Authorization

[bearerToken](../README.md#bearerToken), [apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPermissions

> AdminPermissionsSchema GetPermissions(ctx).Execute()

Gets available permissions



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
	resp, r, err := apiClient.AuthAPI.GetPermissions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.GetPermissions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPermissions`: AdminPermissionsSchema
	fmt.Fprintf(os.Stdout, "Response from `AuthAPI.GetPermissions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPermissionsRequest struct via the builder pattern


### Return type

[**AdminPermissionsSchema**](AdminPermissionsSchema.md)

### Authorization

[bearerToken](../README.md#bearerToken), [apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSimpleSettings

> PasswordAuthSchema GetSimpleSettings(ctx).Execute()

Get Simple auth settings



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
	resp, r, err := apiClient.AuthAPI.GetSimpleSettings(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.GetSimpleSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSimpleSettings`: PasswordAuthSchema
	fmt.Fprintf(os.Stdout, "Response from `AuthAPI.GetSimpleSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSimpleSettingsRequest struct via the builder pattern


### Return type

[**PasswordAuthSchema**](PasswordAuthSchema.md)

### Authorization

[bearerToken](../README.md#bearerToken), [apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetOidcSettings

> OidcSettingsSchema SetOidcSettings(ctx).OidcSettingsSchema(oidcSettingsSchema).Execute()

Set OIDC settings



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
	oidcSettingsSchema := openapiclient.oidcSettingsSchema{OidcSettingsSchemaOneOf: openapiclient.NewOidcSettingsSchemaOneOf(true, "FB87266D-CDDB-4BCF-BB1F-8392FD0EDC1B", "qjcVfeFjEfoYAF3AEsX2IMUWYuUzAbXO")} // OidcSettingsSchema | oidcSettingsSchema

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthAPI.SetOidcSettings(context.Background()).OidcSettingsSchema(oidcSettingsSchema).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.SetOidcSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetOidcSettings`: OidcSettingsSchema
	fmt.Fprintf(os.Stdout, "Response from `AuthAPI.SetOidcSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetOidcSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oidcSettingsSchema** | [**OidcSettingsSchema**](OidcSettingsSchema.md) | oidcSettingsSchema | 

### Return type

[**OidcSettingsSchema**](OidcSettingsSchema.md)

### Authorization

[bearerToken](../README.md#bearerToken), [apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetSimpleSettings

> PasswordAuthSchema SetSimpleSettings(ctx).PasswordAuthSchema(passwordAuthSchema).Execute()

Update Simple auth settings



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
	passwordAuthSchema := *openapiclient.NewPasswordAuthSchema() // PasswordAuthSchema | passwordAuthSchema

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthAPI.SetSimpleSettings(context.Background()).PasswordAuthSchema(passwordAuthSchema).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.SetSimpleSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetSimpleSettings`: PasswordAuthSchema
	fmt.Fprintf(os.Stdout, "Response from `AuthAPI.SetSimpleSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetSimpleSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **passwordAuthSchema** | [**PasswordAuthSchema**](PasswordAuthSchema.md) | passwordAuthSchema | 

### Return type

[**PasswordAuthSchema**](PasswordAuthSchema.md)

### Authorization

[bearerToken](../README.md#bearerToken), [apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

