# {{classname}}

All URIs are relative to *https://us.app.unleash-hosted.com/ushosted*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddPublicSignupTokenUser**](PublicSignupTokensApi.md#AddPublicSignupTokenUser) | **Post** /invite/{token}/signup | Add a user via a signup token
[**CreatePublicSignupToken**](PublicSignupTokensApi.md#CreatePublicSignupToken) | **Post** /api/admin/invite-link/tokens | Create a public signup token
[**GetAllPublicSignupTokens**](PublicSignupTokensApi.md#GetAllPublicSignupTokens) | **Get** /api/admin/invite-link/tokens | Retrieve all existing public signup tokens
[**GetPublicSignupToken**](PublicSignupTokensApi.md#GetPublicSignupToken) | **Get** /api/admin/invite-link/tokens/{token} | Retrieve a token
[**UpdatePublicSignupToken**](PublicSignupTokensApi.md#UpdatePublicSignupToken) | **Put** /api/admin/invite-link/tokens/{token} | Update a public signup token
[**ValidatePublicSignupToken**](PublicSignupTokensApi.md#ValidatePublicSignupToken) | **Get** /invite/{token}/validate | Check whether a public sign-up token exists, has not expired and is enabled

# **AddPublicSignupTokenUser**
> UserSchema AddPublicSignupTokenUser(ctx, body, token)
Add a user via a signup token

Create a user with the viewer root role and link them to the provided signup token

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateInvitedUserSchema**](CreateInvitedUserSchema.md)| createInvitedUserSchema | 
  **token** | **string**|  | 

### Return type

[**UserSchema**](userSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreatePublicSignupToken**
> PublicSignupTokenSchema CreatePublicSignupToken(ctx, body)
Create a public signup token

Lets administrators create a invite link to share with colleagues.  People that join using the public invite are assigned the `Viewer` role

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PublicSignupTokenCreateSchema**](PublicSignupTokenCreateSchema.md)| publicSignupTokenCreateSchema | 

### Return type

[**PublicSignupTokenSchema**](publicSignupTokenSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllPublicSignupTokens**
> PublicSignupTokensSchema GetAllPublicSignupTokens(ctx, )
Retrieve all existing public signup tokens

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**PublicSignupTokensSchema**](publicSignupTokensSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPublicSignupToken**
> PublicSignupTokenSchema GetPublicSignupToken(ctx, token)
Retrieve a token

Get information about a specific token. The `:token` part of the URL should be the token's secret.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**|  | 

### Return type

[**PublicSignupTokenSchema**](publicSignupTokenSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePublicSignupToken**
> PublicSignupTokenSchema UpdatePublicSignupToken(ctx, body, token)
Update a public signup token

Update information about a specific token. The `:token` part of the URL should be the token's secret.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PublicSignupTokenUpdateSchema**](PublicSignupTokenUpdateSchema.md)| publicSignupTokenUpdateSchema | 
  **token** | **string**|  | 

### Return type

[**PublicSignupTokenSchema**](publicSignupTokenSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ValidatePublicSignupToken**
> ValidatePublicSignupToken(ctx, token)
Check whether a public sign-up token exists, has not expired and is enabled

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

