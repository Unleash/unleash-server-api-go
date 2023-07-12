# {{classname}}

All URIs are relative to *https://us.app.unleash-hosted.com/ushosted*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddAccessToProject**](ProjectsApi.md#AddAccessToProject) | **Post** /api/admin/projects/{projectId}/role/{roleId}/access | 
[**AddDefaultStrategyToProjectEnvironment**](ProjectsApi.md#AddDefaultStrategyToProjectEnvironment) | **Post** /api/admin/projects/{projectId}/environments/{environment}/default-strategy | 
[**AddEnvironmentToProject**](ProjectsApi.md#AddEnvironmentToProject) | **Post** /api/admin/projects/{projectId}/environments | Add an environment to a project.
[**AddRoleToUser**](ProjectsApi.md#AddRoleToUser) | **Post** /api/admin/projects/{projectId}/users/{userId}/roles/{roleId} | 
[**ChangeRoleForGroup**](ProjectsApi.md#ChangeRoleForGroup) | **Put** /api/admin/projects/{projectId}/groups/{groupId}/roles/{roleId} | 
[**ChangeRoleForUser**](ProjectsApi.md#ChangeRoleForUser) | **Put** /api/admin/projects/{projectId}/users/{userId}/roles/{roleId} | 
[**CreateProject**](ProjectsApi.md#CreateProject) | **Post** /api/admin/projects | 
[**CreateProjectApiToken**](ProjectsApi.md#CreateProjectApiToken) | **Post** /api/admin/projects/{projectId}/api-tokens | Create a project API token.
[**DeleteProject**](ProjectsApi.md#DeleteProject) | **Delete** /api/admin/projects/{projectId} | 
[**DeleteProjectApiToken**](ProjectsApi.md#DeleteProjectApiToken) | **Delete** /api/admin/projects/{projectId}/api-tokens/{token} | Delete a project API token.
[**GetProjectAccess**](ProjectsApi.md#GetProjectAccess) | **Get** /api/admin/projects/{projectId}/access | 
[**GetProjectApiTokens**](ProjectsApi.md#GetProjectApiTokens) | **Get** /api/admin/projects/{projectId}/api-tokens | Get api tokens for project.
[**GetProjectHealthReport**](ProjectsApi.md#GetProjectHealthReport) | **Get** /api/admin/projects/{projectId}/health-report | Get a health report for a project.
[**GetProjectOverview**](ProjectsApi.md#GetProjectOverview) | **Get** /api/admin/projects/{projectId} | Get an overview of a project.
[**GetProjectSettings**](ProjectsApi.md#GetProjectSettings) | **Get** /api/admin/projects/{projectId}/settings | 
[**GetProjectUsers**](ProjectsApi.md#GetProjectUsers) | **Get** /api/admin/projects/{projectId}/users | 
[**GetProjects**](ProjectsApi.md#GetProjects) | **Get** /api/admin/projects | Get a list of all projects.
[**RemoveEnvironmentFromProject**](ProjectsApi.md#RemoveEnvironmentFromProject) | **Delete** /api/admin/projects/{projectId}/environments/{environment} | Remove an environment from a project.
[**RemoveRoleForUser**](ProjectsApi.md#RemoveRoleForUser) | **Delete** /api/admin/projects/{projectId}/users/{userId}/roles/{roleId} | 
[**RemoveRoleFromGroup**](ProjectsApi.md#RemoveRoleFromGroup) | **Delete** /api/admin/projects/{projectId}/groups/{groupId}/roles/{roleId} | 
[**UpdateProject**](ProjectsApi.md#UpdateProject) | **Put** /api/admin/projects/{projectId} | 
[**ValidateProject**](ProjectsApi.md#ValidateProject) | **Post** /api/admin/projects/validate | 

# **AddAccessToProject**
> AddAccessToProject(ctx, projectId, roleId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**|  | 
  **roleId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddDefaultStrategyToProjectEnvironment**
> CreateFeatureStrategySchema AddDefaultStrategyToProjectEnvironment(ctx, body, projectId, environment)


Adds a default strategy for this environment. Unleash will use this strategy by default when enabling a toggle. Use the wild card \"*\" for `:environment` to add to all environments. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateFeatureStrategySchema**](CreateFeatureStrategySchema.md)| createFeatureStrategySchema | 
  **projectId** | **string**|  | 
  **environment** | **string**|  | 

### Return type

[**CreateFeatureStrategySchema**](createFeatureStrategySchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddEnvironmentToProject**
> AddEnvironmentToProject(ctx, body, projectId)
Add an environment to a project.

This endpoint adds the provided environment to the specified project, with optional support for enabling and disabling change requests for the environment and project.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ProjectEnvironmentSchema**](ProjectEnvironmentSchema.md)| projectEnvironmentSchema | 
  **projectId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddRoleToUser**
> AddRoleToUser(ctx, projectId, userId, roleId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**|  | 
  **userId** | **string**|  | 
  **roleId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChangeRoleForGroup**
> ChangeRoleForGroup(ctx, projectId, groupId, roleId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**|  | 
  **groupId** | **string**|  | 
  **roleId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChangeRoleForUser**
> ChangeRoleForUser(ctx, projectId, userId, roleId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**|  | 
  **userId** | **string**|  | 
  **roleId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateProject**
> ProjectCreatedSchema CreateProject(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateProjectSchema**](CreateProjectSchema.md)| createProjectSchema | 

### Return type

[**ProjectCreatedSchema**](projectCreatedSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateProjectApiToken**
> ApiTokenSchema CreateProjectApiToken(ctx, body, projectId)
Create a project API token.

Endpoint that allows creation of [project API tokens](https://docs.getunleash.io/reference/api-tokens-and-client-keys#api-token-visibility) for the specified project.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateApiTokenSchema**](CreateApiTokenSchema.md)| createApiTokenSchema | 
  **projectId** | **string**|  | 

### Return type

[**ApiTokenSchema**](apiTokenSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteProject**
> DeleteProject(ctx, projectId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteProjectApiToken**
> DeleteProjectApiToken(ctx, projectId, token)
Delete a project API token.

This operation deletes the API token specified in the request URL. If the token doesn't exist, returns an OK response (status code 200).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**|  | 
  **token** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProjectAccess**
> ProjectAccessSchema GetProjectAccess(ctx, projectId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**|  | 

### Return type

[**ProjectAccessSchema**](projectAccessSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProjectApiTokens**
> ApiTokensSchema GetProjectApiTokens(ctx, projectId)
Get api tokens for project.

Returns the [project API tokens](https://docs.getunleash.io/how-to/how-to-create-project-api-tokens) that have been created for this project.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**|  | 

### Return type

[**ApiTokensSchema**](apiTokensSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProjectHealthReport**
> HealthReportSchema GetProjectHealthReport(ctx, projectId)
Get a health report for a project.

This endpoint returns a health report for the specified project. This data is used for [the technical debt dashboard](https://docs.getunleash.io/reference/technical-debt#the-technical-debt-dashboard)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**|  | 

### Return type

[**HealthReportSchema**](healthReportSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProjectOverview**
> ProjectOverviewSchema GetProjectOverview(ctx, projectId)
Get an overview of a project.

This endpoint returns an overview of the specified projects stats, project health, number of members, which environments are configured, and the features in the project.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**|  | 

### Return type

[**ProjectOverviewSchema**](projectOverviewSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProjectSettings**
> ProjectSettingsSchema GetProjectSettings(ctx, projectId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**|  | 

### Return type

[**ProjectSettingsSchema**](projectSettingsSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProjectUsers**
> ProjectUsersSchema GetProjectUsers(ctx, projectId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**|  | 

### Return type

[**ProjectUsersSchema**](projectUsersSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProjects**
> ProjectsSchema GetProjects(ctx, )
Get a list of all projects.

This endpoint returns an list of all the projects in the Unleash instance.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ProjectsSchema**](projectsSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveEnvironmentFromProject**
> RemoveEnvironmentFromProject(ctx, projectId, environment)
Remove an environment from a project.

This endpoint removes the specified environment from the project.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**|  | 
  **environment** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveRoleForUser**
> RemoveRoleForUser(ctx, projectId, userId, roleId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**|  | 
  **userId** | **string**|  | 
  **roleId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveRoleFromGroup**
> RemoveRoleFromGroup(ctx, projectId, groupId, roleId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**|  | 
  **groupId** | **string**|  | 
  **roleId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateProject**
> UpdateProject(ctx, body, projectId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateProjectSchema**](UpdateProjectSchema.md)| updateProjectSchema | 
  **projectId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ValidateProject**
> ValidateProject(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ValidateProjectSchema**](ValidateProjectSchema.md)| validateProjectSchema | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

