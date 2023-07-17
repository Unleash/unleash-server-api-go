# {{classname}}

All URIs are relative to *https://us.app.unleash-hosted.com/ushosted*

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

# **ChangePassword**
> ChangePassword(ctx, body)
Changes a user password

Allows users with a valid reset token to reset their password without remembering their old password

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ChangePasswordSchema**](ChangePasswordSchema.md)| changePasswordSchema | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGoogleSettings**
> GoogleSettingsSchema GetGoogleSettings(ctx, )
Get Google auth settings

Returns the current settings for Google Authentication (deprecated, please use OpenID instead)

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GoogleSettingsSchema**](googleSettingsSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOidcSettings**
> OidcSettingsSchema GetOidcSettings(ctx, )
Get OIDC auth settings

Returns the current settings for OIDC Authentication

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**OidcSettingsSchema**](oidcSettingsSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPermissions**
> AdminPermissionsSchema GetPermissions(ctx, )
Gets available permissions

Returns a list of available permissions

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**AdminPermissionsSchema**](adminPermissionsSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSamlSettings**
> SamlSettingsSchema GetSamlSettings(ctx, )
Get SAML auth settings

Returns the current settings for SAML authentication

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**SamlSettingsSchema**](samlSettingsSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSimpleSettings**
> PasswordAuthSchema GetSimpleSettings(ctx, )
Get Simple auth settings

Is simple authentication (username/password) enabled for this server

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**PasswordAuthSchema**](passwordAuthSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Login**
> UserSchema Login(ctx, body)
Log in

Logs in the user and creates an active session

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**LoginSchema**](LoginSchema.md)| loginSchema | 

### Return type

[**UserSchema**](userSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SendResetPasswordEmail**
> SendResetPasswordEmail(ctx, body)
Reset password

Requests a password reset email for the user. This email can be used to reset the password for a user that has forgotten their password

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EmailSchema**](EmailSchema.md)| emailSchema | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetGoogleSettings**
> GoogleSettingsSchema SetGoogleSettings(ctx, body)
Set Google auth options

Updates the settings for Google Authentication (deprecated, please use OpenID instead)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GoogleSettingsSchema**](GoogleSettingsSchema.md)| googleSettingsSchema | 

### Return type

[**GoogleSettingsSchema**](googleSettingsSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetOidcSettings**
> OidcSettingsSchema SetOidcSettings(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**OidcSettingsSchema**](OidcSettingsSchema.md)| oidcSettingsSchema | 

### Return type

[**OidcSettingsSchema**](oidcSettingsSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetSamlSettings**
> SamlSettingsSchema SetSamlSettings(ctx, body)
Update SAML auth settings

Updates the settings for SAML Authentication

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SamlSettingsSchema**](SamlSettingsSchema.md)| samlSettingsSchema | 

### Return type

[**SamlSettingsSchema**](samlSettingsSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetSimpleSettings**
> PasswordAuthSchema SetSimpleSettings(ctx, body)
Update Simple auth settings

Enable or disable simple authentication (username/password)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PasswordAuthSchema**](PasswordAuthSchema.md)| passwordAuthSchema | 

### Return type

[**PasswordAuthSchema**](passwordAuthSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ValidatePassword**
> ValidatePassword(ctx, body)
Validates password

Verifies that the password adheres to the [Unleash password guidelines](https://docs.getunleash.io/reference/deploy/securing-unleash#password-requirements)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ValidatePasswordSchema**](ValidatePasswordSchema.md)| validatePasswordSchema | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ValidateToken**
> TokenUserSchema ValidateToken(ctx, )
Validates a token

If the token is valid returns the user that owns the token

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**TokenUserSchema**](tokenUserSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

