# \AuthAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOidcSettings**](AuthAPI.md#GetOidcSettings) | **Get** /api/admin/auth/oidc/settings | Get OIDC auth settings
[**GetPermissions**](AuthAPI.md#GetPermissions) | **Get** /api/admin/permissions | Gets available permissions
[**GetSamlSettings**](AuthAPI.md#GetSamlSettings) | **Get** /api/admin/auth/saml/settings | Get SAML auth settings
[**GetSimpleSettings**](AuthAPI.md#GetSimpleSettings) | **Get** /api/admin/auth/simple/settings | Get Simple auth settings
[**SetOidcSettings**](AuthAPI.md#SetOidcSettings) | **Post** /api/admin/auth/oidc/settings | Set OIDC settings
[**SetSamlSettings**](AuthAPI.md#SetSamlSettings) | **Post** /api/admin/auth/saml/settings | Update SAML auth settings
[**SetSimpleSettings**](AuthAPI.md#SetSimpleSettings) | **Post** /api/admin/auth/simple/settings | Update Simple auth settings



## GetOidcSettings

> OidcSettingsResponseSchema GetOidcSettings(ctx).Execute()

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
	// response from `GetOidcSettings`: OidcSettingsResponseSchema
	fmt.Fprintf(os.Stdout, "Response from `AuthAPI.GetOidcSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetOidcSettingsRequest struct via the builder pattern


### Return type

[**OidcSettingsResponseSchema**](OidcSettingsResponseSchema.md)

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


## GetSamlSettings

> SamlSettingsResponseSchema GetSamlSettings(ctx).Execute()

Get SAML auth settings



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
	resp, r, err := apiClient.AuthAPI.GetSamlSettings(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.GetSamlSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSamlSettings`: SamlSettingsResponseSchema
	fmt.Fprintf(os.Stdout, "Response from `AuthAPI.GetSamlSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSamlSettingsRequest struct via the builder pattern


### Return type

[**SamlSettingsResponseSchema**](SamlSettingsResponseSchema.md)

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

> OidcSettingsResponseSchema SetOidcSettings(ctx).OidcSettingsSchema(oidcSettingsSchema).Execute()

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
	// response from `SetOidcSettings`: OidcSettingsResponseSchema
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

[**OidcSettingsResponseSchema**](OidcSettingsResponseSchema.md)

### Authorization

[bearerToken](../README.md#bearerToken), [apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetSamlSettings

> SamlSettingsResponseSchema SetSamlSettings(ctx).SamlSettingsSchema(samlSettingsSchema).Execute()

Update SAML auth settings



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
	samlSettingsSchema := openapiclient.samlSettingsSchema{SamlSettingsSchemaOneOf: openapiclient.NewSamlSettingsSchemaOneOf("http://localhost:8080/auth/realms/master", "http://localhost:8080/auth/realms/master/protocol/saml", "MIIE/zCCBGgCAg4CMA0GCSqGSIb3DQEBBQUAMIGbMQswCQYDVQQGEwJKUDEOMAwG
A1UECBMFVG9reW8xEDAOBgNVBAcTB0NodW8ta3UxETAPBgNVBAoTCEZyYW5rNERE
MRgwFgYDVQQLEw9XZWJDZXJ0IFN1cHBvcnQxGDAWBgNVBAMTD0ZyYW5rNEREIFdl
YiBDQTEjMCEGCSqGSIb3DQEJARYUc3VwcG9ydEBmcmFuazRkZC5jb20wHhcNMTIw
ODIyMDcyNzIyWhcNMTcwODIxMDcyNzIyWjBKMQswCQYDVQQGEwJKUDEOMAwGA1UE
CAwFVG9reW8xETAPBgNVBAoMCEZyYW5rNEREMRgwFgYDVQQDDA93d3cuZXhhbXBs
ZS5jb20wggNHMIICOQYHKoZIzjgEATCCAiwCggEBAMbmu6uSdZWRxnO5PteARz5I
nrdM7vJadzJcY9Spf0cAhUDWyINCrUDn8h8QjbMiAxd+E7v5N85TbfvrW+/g7lYa
e7DB0uX02Rw29yoK+TE/znNTNq6HdPn/H4ll77uJqpkWgQwXgAQ3qDKRv96QaTfr
fSXYRxG9NvKzUBs9S7Woby7K6Pgh7/EmgeiOHKSX49XD+ihvkFRSFdeL5qV5hzDf
epfk8Ghl0cyK9jpM/yKlRuUUkP1pSMwUrCrptyRpqqXmam0UfFuFhMT2SJyNTyo2
SfnyZb78lbYcVLJQLJA+O3l469eOH3Odv/Pq7bvOstBKg96Q7imz5t0honf63EMC
IQCw7FeY0QQbxWYy+MI8/0m2kiRzIruA8RUPcEs4il1mwwKCAQB4W3QLepQRTWhR
69Xv+fC5JLEDyngw9KEalKorlg6o4Z9hASMbFMwECjlXZCxKd/NFjUMbtPcrMCoa
/KuaqRvHLs2bqe94X5VR4lWCv0SgOunKBj58jnVuN/OkkLu4cSZQ/jia/yPkdcMv
w8ZjF7zjPXGVhh9XC0QU9ipVfrreGaBSN+0zODKY5TyQI84FsZFZNetOTIT0HT2S
fIDRGYaL/0xMfQx070Z07cdTTuibzJHVr38qjKqEDiwAUyjXVdE+GJ15ZD4d56Ef
0qgRpzDmuvUjOtv1t8Hr2O2HTABqRMtAKZsLEVPjwnpKpcStixfg0uIPGVIKbej4
FzHHpO6bA4IBBgACggEBAJNRaaTFe253sOVm/JmUgsO1QB5GI5hOEWLpC8KHxgwn
nf/GQUaJLrN8TT4hXgJM2CdvdAkY6et1HpT6BUoz1cYTgsE3ToIsbH3SzPJvU7jz
cPOvY1jQv+xVBrU8Ydw2D8pydbAcw/L6JZnGpFBqeHa1iFAQc0B8ToXEgxnmGAdP
IOAKAHX0S4m6CrP5fKwYbmzu8WuWO4bRqvX7QJofrs2RaGFESulw0VrMFffJ/guf
HTvhDaMW7TSCKo1tBZK9SdEbWCQN2stnfnRSyZFQ+v02oyQtLg+3vSuCx4PS9DM9
/Uh3r9JDDH3GveUMbqw8Dmy6WH9iV3oOJt8aVF8F4CMwDQYJKoZIhvcNAQEFBQAD
gYEAbxDoJM8vKVfhltpfG3YXmBKnoGb2UpdKpcjmxMt1/yX8lWJaRBwUDeiDqjVC
JGi9gXO2SDAtXl7GI1cXTs/l7QlmoTmnc6kDwqk3pl6jC72rQH3E/Fpg7hBkSWL9
3V1dbLU5id63lVD8sUEULyfWFGk3L+Uka5oiSsxwZhdIb/Q=
")} // SamlSettingsSchema | samlSettingsSchema

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthAPI.SetSamlSettings(context.Background()).SamlSettingsSchema(samlSettingsSchema).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.SetSamlSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetSamlSettings`: SamlSettingsResponseSchema
	fmt.Fprintf(os.Stdout, "Response from `AuthAPI.SetSamlSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetSamlSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **samlSettingsSchema** | [**SamlSettingsSchema**](SamlSettingsSchema.md) | samlSettingsSchema | 

### Return type

[**SamlSettingsResponseSchema**](SamlSettingsResponseSchema.md)

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

