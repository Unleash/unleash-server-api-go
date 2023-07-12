# {{classname}}

All URIs are relative to *https://us.app.unleash-hosted.com/ushosted*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloneEnvironment**](EnvironmentsApi.md#CloneEnvironment) | **Post** /api/admin/environments/{name}/clone | Clones an environment
[**CreateEnvironment**](EnvironmentsApi.md#CreateEnvironment) | **Post** /api/admin/environments | Creates a new environment
[**GetAllEnvironments**](EnvironmentsApi.md#GetAllEnvironments) | **Get** /api/admin/environments | Get all environments
[**GetEnvironment**](EnvironmentsApi.md#GetEnvironment) | **Get** /api/admin/environments/{name} | Get the environment with &#x60;name&#x60;
[**GetProjectEnvironments**](EnvironmentsApi.md#GetProjectEnvironments) | **Get** /api/admin/environments/project/{projectId} | Get the environments available to a project
[**RemoveEnvironment**](EnvironmentsApi.md#RemoveEnvironment) | **Delete** /api/admin/environments/{name} | Deletes an environment by name
[**ToggleEnvironmentOff**](EnvironmentsApi.md#ToggleEnvironmentOff) | **Post** /api/admin/environments/{name}/off | Toggle the environment with &#x60;name&#x60; off
[**ToggleEnvironmentOn**](EnvironmentsApi.md#ToggleEnvironmentOn) | **Post** /api/admin/environments/{name}/on | Toggle the environment with &#x60;name&#x60; on
[**UpdateEnvironment**](EnvironmentsApi.md#UpdateEnvironment) | **Put** /api/admin/environments/update/{name} | Updates an environment by name
[**UpdateSortOrder**](EnvironmentsApi.md#UpdateSortOrder) | **Put** /api/admin/environments/sort-order | Update environment sort orders
[**ValidateEnvironmentName**](EnvironmentsApi.md#ValidateEnvironmentName) | **Post** /api/admin/environments/validate | Validates if an environment name exists

# **CloneEnvironment**
> CloneEnvironment(ctx, name)
Clones an environment

Given an existing environment name and a set of options, this will create a copy of that environment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateEnvironment**
> EnvironmentSchema CreateEnvironment(ctx, body)
Creates a new environment

Uses the details provided in the payload to create a new environment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateEnvironmentSchema**](CreateEnvironmentSchema.md)| createEnvironmentSchema | 

### Return type

[**EnvironmentSchema**](environmentSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllEnvironments**
> EnvironmentsSchema GetAllEnvironments(ctx, )
Get all environments

Retrieves all environments that exist in this Unleash instance.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**EnvironmentsSchema**](environmentsSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEnvironment**
> EnvironmentSchema GetEnvironment(ctx, name)
Get the environment with `name`

Retrieves the environment with `name` if it exists in this Unleash instance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**|  | 

### Return type

[**EnvironmentSchema**](environmentSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProjectEnvironments**
> EnvironmentsProjectSchema GetProjectEnvironments(ctx, projectId)
Get the environments available to a project

Gets the environments that are available for this project. An environment is available for a project if enabled in the [project configuration](https://docs.getunleash.io/reference/environments#step-1-enable-new-environments-for-your-project)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**|  | 

### Return type

[**EnvironmentsProjectSchema**](environmentsProjectSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveEnvironment**
> RemoveEnvironment(ctx, name)
Deletes an environment by name

Given an existing environment by name, this endpoint will attempt to delete it

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ToggleEnvironmentOff**
> ToggleEnvironmentOff(ctx, name)
Toggle the environment with `name` off

Removes this environment from the list of available environments for projects to use

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ToggleEnvironmentOn**
> ToggleEnvironmentOn(ctx, name)
Toggle the environment with `name` on

Makes it possible to enable this environment for a project. An environment must first be globally enabled using this endpoint before it can be enabled for a project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateEnvironment**
> EnvironmentSchema UpdateEnvironment(ctx, body, name)
Updates an environment by name

Given an environment by name updates the environment with the given payload. Note that `name`, `enabled` and `protected` cannot be changed by this API

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateEnvironmentSchema**](UpdateEnvironmentSchema.md)| updateEnvironmentSchema | 
  **name** | **string**|  | 

### Return type

[**EnvironmentSchema**](environmentSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSortOrder**
> UpdateSortOrder(ctx, body)
Update environment sort orders

Updates sort orders for the named environments. Environments not specified are unaffected.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**map[string]float64**](map.md)| sortOrderSchema | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ValidateEnvironmentName**
> ValidateEnvironmentName(ctx, body)
Validates if an environment name exists

Uses the name provided in the body of the request to validate if the given name exists or not

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**NameSchema**](NameSchema.md)| nameSchema | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

