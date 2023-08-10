# \AuthApi

All URIs are relative to *https://app.unleash-hosted.com/hosted*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChangePassword**](AuthApi.md#ChangePassword) | **Post** /auth/reset/password | Changes a user password
[**GetGoogleSettings**](AuthApi.md#GetGoogleSettings) | **Get** /api/admin/auth/google/settings | Get Google auth settings
[**GetOidcSettings**](AuthApi.md#GetOidcSettings) | **Get** /api/admin/auth/oidc/settings | Get OIDC auth settings
[**GetPermissions**](AuthApi.md#GetPermissions) | **Get** /api/admin/permissions | Gets available permissions
[**GetSamlSettings**](AuthApi.md#GetSamlSettings) | **Get** /api/admin/auth/saml/settings | Get SAML auth settings
[**GetSimpleSettings**](AuthApi.md#GetSimpleSettings) | **Get** /api/admin/auth/simple/settings | Get Simple auth settings
[**Login**](AuthApi.md#Login) | **Post** /auth/simple/login | Log in
[**SendResetPasswordEmail**](AuthApi.md#SendResetPasswordEmail) | **Post** /auth/reset/password-email | Reset password
[**SetGoogleSettings**](AuthApi.md#SetGoogleSettings) | **Post** /api/admin/auth/google/settings | Set Google auth options
[**SetOidcSettings**](AuthApi.md#SetOidcSettings) | **Post** /api/admin/auth/oidc/settings | 
[**SetSamlSettings**](AuthApi.md#SetSamlSettings) | **Post** /api/admin/auth/saml/settings | Update SAML auth settings
[**SetSimpleSettings**](AuthApi.md#SetSimpleSettings) | **Post** /api/admin/auth/simple/settings | Update Simple auth settings
[**ValidatePassword**](AuthApi.md#ValidatePassword) | **Post** /auth/reset/validate-password | Validates password
[**ValidateToken**](AuthApi.md#ValidateToken) | **Get** /auth/reset/validate | Validates a token



## ChangePassword

> ChangePassword(ctx).ChangePasswordSchema(changePasswordSchema).Execute()

Changes a user password



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
    changePasswordSchema := *openapiclient.NewChangePasswordSchema("$2a$15$QzeW/y5/MEppCWVEkoX5euejobYOLSd4We21LQjjKlWH9l2I3wCke", "correct horse battery staple") // ChangePasswordSchema | changePasswordSchema

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AuthApi.ChangePassword(context.Background()).ChangePasswordSchema(changePasswordSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.ChangePassword``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChangePasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **changePasswordSchema** | [**ChangePasswordSchema**](ChangePasswordSchema.md) | changePasswordSchema | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGoogleSettings

> GoogleSettingsSchema GetGoogleSettings(ctx).Execute()

Get Google auth settings



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
    resp, r, err := apiClient.AuthApi.GetGoogleSettings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.GetGoogleSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGoogleSettings`: GoogleSettingsSchema
    fmt.Fprintf(os.Stdout, "Response from `AuthApi.GetGoogleSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetGoogleSettingsRequest struct via the builder pattern


### Return type

[**GoogleSettingsSchema**](GoogleSettingsSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
    resp, r, err := apiClient.AuthApi.GetOidcSettings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.GetOidcSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOidcSettings`: OidcSettingsSchema
    fmt.Fprintf(os.Stdout, "Response from `AuthApi.GetOidcSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetOidcSettingsRequest struct via the builder pattern


### Return type

[**OidcSettingsSchema**](OidcSettingsSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

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
    resp, r, err := apiClient.AuthApi.GetPermissions(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.GetPermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPermissions`: AdminPermissionsSchema
    fmt.Fprintf(os.Stdout, "Response from `AuthApi.GetPermissions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPermissionsRequest struct via the builder pattern


### Return type

[**AdminPermissionsSchema**](AdminPermissionsSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSamlSettings

> SamlSettingsSchema GetSamlSettings(ctx).Execute()

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
    resp, r, err := apiClient.AuthApi.GetSamlSettings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.GetSamlSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSamlSettings`: SamlSettingsSchema
    fmt.Fprintf(os.Stdout, "Response from `AuthApi.GetSamlSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSamlSettingsRequest struct via the builder pattern


### Return type

[**SamlSettingsSchema**](SamlSettingsSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

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
    resp, r, err := apiClient.AuthApi.GetSimpleSettings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.GetSimpleSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSimpleSettings`: PasswordAuthSchema
    fmt.Fprintf(os.Stdout, "Response from `AuthApi.GetSimpleSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSimpleSettingsRequest struct via the builder pattern


### Return type

[**PasswordAuthSchema**](PasswordAuthSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Login

> UserSchema Login(ctx).LoginSchema(loginSchema).Execute()

Log in



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
    loginSchema := *openapiclient.NewLoginSchema("user", "hunter2") // LoginSchema | loginSchema

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthApi.Login(context.Background()).LoginSchema(loginSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.Login``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Login`: UserSchema
    fmt.Fprintf(os.Stdout, "Response from `AuthApi.Login`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLoginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **loginSchema** | [**LoginSchema**](LoginSchema.md) | loginSchema | 

### Return type

[**UserSchema**](UserSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendResetPasswordEmail

> SendResetPasswordEmail(ctx).EmailSchema(emailSchema).Execute()

Reset password



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
    emailSchema := *openapiclient.NewEmailSchema("test@example.com") // EmailSchema | emailSchema

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AuthApi.SendResetPasswordEmail(context.Background()).EmailSchema(emailSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.SendResetPasswordEmail``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendResetPasswordEmailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **emailSchema** | [**EmailSchema**](EmailSchema.md) | emailSchema | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetGoogleSettings

> GoogleSettingsSchema SetGoogleSettings(ctx).GoogleSettingsSchema(googleSettingsSchema).Execute()

Set Google auth options



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
    googleSettingsSchema := *openapiclient.NewGoogleSettingsSchema("your-client-id", "your-client-secret", "localhost") // GoogleSettingsSchema | googleSettingsSchema

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthApi.SetGoogleSettings(context.Background()).GoogleSettingsSchema(googleSettingsSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.SetGoogleSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetGoogleSettings`: GoogleSettingsSchema
    fmt.Fprintf(os.Stdout, "Response from `AuthApi.SetGoogleSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetGoogleSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **googleSettingsSchema** | [**GoogleSettingsSchema**](GoogleSettingsSchema.md) | googleSettingsSchema | 

### Return type

[**GoogleSettingsSchema**](GoogleSettingsSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetOidcSettings

> OidcSettingsSchema SetOidcSettings(ctx).OidcSettingsSchema(oidcSettingsSchema).Execute()



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
    oidcSettingsSchema := *openapiclient.NewOidcSettingsSchema("FB87266D-CDDB-4BCF-BB1F-8392FD0EDC1B", "qjcVfeFjEfoYAF3AEsX2IMUWYuUzAbXO") // OidcSettingsSchema | oidcSettingsSchema

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthApi.SetOidcSettings(context.Background()).OidcSettingsSchema(oidcSettingsSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.SetOidcSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetOidcSettings`: OidcSettingsSchema
    fmt.Fprintf(os.Stdout, "Response from `AuthApi.SetOidcSettings`: %v\n", resp)
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

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetSamlSettings

> SamlSettingsSchema SetSamlSettings(ctx).SamlSettingsSchema(samlSettingsSchema).Execute()

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
    samlSettingsSchema := *openapiclient.NewSamlSettingsSchema("http://localhost:8080/auth/realms/master", "http://localhost:8080/auth/realms/master/protocol/saml", "MIIE/zCCBGgCAg4CMA0GCSqGSIb3DQEBBQUAMIGbMQswCQYDVQQGEwJKUDEOMAwG
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
") // SamlSettingsSchema | samlSettingsSchema

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthApi.SetSamlSettings(context.Background()).SamlSettingsSchema(samlSettingsSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.SetSamlSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetSamlSettings`: SamlSettingsSchema
    fmt.Fprintf(os.Stdout, "Response from `AuthApi.SetSamlSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetSamlSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **samlSettingsSchema** | [**SamlSettingsSchema**](SamlSettingsSchema.md) | samlSettingsSchema | 

### Return type

[**SamlSettingsSchema**](SamlSettingsSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

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
    resp, r, err := apiClient.AuthApi.SetSimpleSettings(context.Background()).PasswordAuthSchema(passwordAuthSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.SetSimpleSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetSimpleSettings`: PasswordAuthSchema
    fmt.Fprintf(os.Stdout, "Response from `AuthApi.SetSimpleSettings`: %v\n", resp)
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

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidatePassword

> ValidatePassword(ctx).ValidatePasswordSchema(validatePasswordSchema).Execute()

Validates password



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
    validatePasswordSchema := *openapiclient.NewValidatePasswordSchema("hunter2") // ValidatePasswordSchema | validatePasswordSchema

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AuthApi.ValidatePassword(context.Background()).ValidatePasswordSchema(validatePasswordSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.ValidatePassword``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidatePasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **validatePasswordSchema** | [**ValidatePasswordSchema**](ValidatePasswordSchema.md) | validatePasswordSchema | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateToken

> TokenUserSchema ValidateToken(ctx).Execute()

Validates a token



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
    resp, r, err := apiClient.AuthApi.ValidateToken(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.ValidateToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ValidateToken`: TokenUserSchema
    fmt.Fprintf(os.Stdout, "Response from `AuthApi.ValidateToken`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiValidateTokenRequest struct via the builder pattern


### Return type

[**TokenUserSchema**](TokenUserSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

