# {{classname}}

All URIs are relative to *https://us.app.unleash-hosted.com/ushosted*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddChangeRequestComment**](UnstableApi.md#AddChangeRequestComment) | **Post** /api/admin/projects/{projectId}/change-requests/{id}/comments | This endpoint will add a comment to a change request
[**ChangeRequest**](UnstableApi.md#ChangeRequest) | **Post** /api/admin/projects/{projectId}/environments/{environment}/change-requests | Create/Add change to a change request
[**DeleteChange**](UnstableApi.md#DeleteChange) | **Delete** /api/admin/projects/{projectId}/change-requests/{changeRequestId}/changes/{changeId} | Discards a change from a change request by change id
[**DeleteChangeRequest**](UnstableApi.md#DeleteChangeRequest) | **Delete** /api/admin/projects/{projectId}/change-requests/{id} | Deletes a change request by id
[**EditChange**](UnstableApi.md#EditChange) | **Put** /api/admin/projects/{projectId}/change-requests/{changeRequestId}/changes/{changeId} | Edits a single change in a change request
[**GetAdvancedPlayground**](UnstableApi.md#GetAdvancedPlayground) | **Post** /api/admin/playground/advanced | Batch evaluate an Unleash context against a set of environments and projects.
[**GetChangeRequest**](UnstableApi.md#GetChangeRequest) | **Get** /api/admin/projects/{projectId}/change-requests/{id} | Retrieves one change request by id
[**GetChangeRequestsForProject**](UnstableApi.md#GetChangeRequestsForProject) | **Get** /api/admin/projects/{projectId}/change-requests | Retrieves all change requests for a project
[**GetLoginHistory**](UnstableApi.md#GetLoginHistory) | **Get** /api/admin/logins | Get all login events.
[**GetNotifications**](UnstableApi.md#GetNotifications) | **Get** /api/admin/notifications | Retrieves a list of notifications
[**GetOpenChangeRequestsForUser**](UnstableApi.md#GetOpenChangeRequestsForUser) | **Get** /api/admin/projects/{projectId}/change-requests/open | Retrieves pending change requests in configured environments
[**GetPendingChangeRequestsForFeature**](UnstableApi.md#GetPendingChangeRequestsForFeature) | **Get** /api/admin/projects/{projectId}/change-requests/pending/{featureName} | Retrieves all pending change requests referencing a feature in the project
[**GetPendingChangeRequestsForUser**](UnstableApi.md#GetPendingChangeRequestsForUser) | **Get** /api/admin/projects/{projectId}/change-requests/pending | Retrieves pending change requests in configured environments
[**GetProjectChangeRequestConfig**](UnstableApi.md#GetProjectChangeRequestConfig) | **Get** /api/admin/projects/{projectId}/change-requests/config | Retrieves change request configuration for a project
[**MarkNotificationsAsRead**](UnstableApi.md#MarkNotificationsAsRead) | **Post** /api/admin/notifications/read | Mark notifications as read
[**UpdateChangeRequestState**](UnstableApi.md#UpdateChangeRequestState) | **Put** /api/admin/projects/{projectId}/change-requests/{id}/state | This endpoint will update the state of a change request
[**UpdateChangeRequestTitle**](UnstableApi.md#UpdateChangeRequestTitle) | **Put** /api/admin/projects/{projectId}/change-requests/{id}/title | This endpoint will update the custom title of a change request
[**UpdateProjectChangeRequestConfig**](UnstableApi.md#UpdateProjectChangeRequestConfig) | **Put** /api/admin/projects/{projectId}/environments/{environment}/change-requests/config | Updates change request configuration for an environment in the project

# **AddChangeRequestComment**
> AddChangeRequestComment(ctx, body, projectId, id)
This endpoint will add a comment to a change request

This endpoint will add a comment to a change request for the user making the request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ChangeRequestAddCommentSchema**](ChangeRequestAddCommentSchema.md)| changeRequestAddCommentSchema | 
  **projectId** | **string**|  | 
  **id** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChangeRequest**
> ChangeRequestSchema ChangeRequest(ctx, body, projectId, environment)
Create/Add change to a change request

Given a change request exists, this endpoint will attempt to add a change to                          an existing change request for the user. If a change request does not exist.                          It will attempt to create it.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ChangeRequestOneOrManyCreateSchema**](ChangeRequestOneOrManyCreateSchema.md)| changeRequestOneOrManyCreateSchema | 
  **projectId** | **string**|  | 
  **environment** | **string**|  | 

### Return type

[**ChangeRequestSchema**](changeRequestSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteChange**
> DeleteChange(ctx, projectId, changeRequestId, changeId)
Discards a change from a change request by change id

This endpoint will discard one change from a change request if it matches the provided id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**|  | 
  **changeRequestId** | **string**|  | 
  **changeId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteChangeRequest**
> DeleteChangeRequest(ctx, projectId, id)
Deletes a change request by id

This endpoint will delete one change request if it matches the provided id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**|  | 
  **id** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditChange**
> ChangeRequestSchema EditChange(ctx, body, projectId, changeRequestId, changeId)
Edits a single change in a change request

This endpoint will edit one change from a change request if it matches the provided id. The edit removes previous change and inserts a new one. You                     should not rely on the changeId for subsequent edits and always check the most recent changeId.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ChangeRequestCreateSchema**](ChangeRequestCreateSchema.md)| changeRequestCreateSchema | 
  **projectId** | **string**|  | 
  **changeRequestId** | **string**|  | 
  **changeId** | **string**|  | 

### Return type

[**ChangeRequestSchema**](changeRequestSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAdvancedPlayground**
> AdvancedPlaygroundResponseSchema GetAdvancedPlayground(ctx, body)
Batch evaluate an Unleash context against a set of environments and projects.

Use the provided `context`, `environments`, and `projects` to evaluate toggles on this Unleash instance. You can use comma-separated values to provide multiple values to each context field. Returns a combinatorial list of all toggles that match the parameters and what they evaluate to. The response also contains the input parameters that were provided.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AdvancedPlaygroundRequestSchema**](AdvancedPlaygroundRequestSchema.md)| advancedPlaygroundRequestSchema | 

### Return type

[**AdvancedPlaygroundResponseSchema**](advancedPlaygroundResponseSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetChangeRequest**
> ChangeRequestSchema GetChangeRequest(ctx, projectId, id)
Retrieves one change request by id

This endpoint will retrieve one change request if it matches the provided id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**|  | 
  **id** | **string**|  | 

### Return type

[**ChangeRequestSchema**](changeRequestSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetChangeRequestsForProject**
> Array GetChangeRequestsForProject(ctx, projectId)
Retrieves all change requests for a project

This endpoint will retrieve all change requests regardless of status for a given project. There's an upper limit of last 300 change requests ordered by creation date.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**|  | 

### Return type

[**Array**](array.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLoginHistory**
> LoginHistorySchema GetLoginHistory(ctx, )
Get all login events.

Returns **all** login events in the Unleash system. You can optionally get them in CSV format by specifying the `Accept` header as `text/csv`.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**LoginHistorySchema**](loginHistorySchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNotifications**
> Array GetNotifications(ctx, )
Retrieves a list of notifications

A user may get relevant notifications from the projects they are a member of

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**Array**](array.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOpenChangeRequestsForUser**
> Array GetOpenChangeRequestsForUser(ctx, projectId)
Retrieves pending change requests in configured environments

This endpoint will retrieve the pending change requests in the configured environments for the project, for the current user performing the request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**|  | 

### Return type

[**Array**](array.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPendingChangeRequestsForFeature**
> Array GetPendingChangeRequestsForFeature(ctx, projectId, featureName)
Retrieves all pending change requests referencing a feature in the project

This endpoint will retrieve all pending change requests (change requests with a status of Draft | In review | Approved) referencing the given feature toggle name.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**|  | 
  **featureName** | **string**|  | 

### Return type

[**Array**](array.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPendingChangeRequestsForUser**
> Array GetPendingChangeRequestsForUser(ctx, projectId)
Retrieves pending change requests in configured environments

This endpoint will retrieve the pending change requests in the configured environments for the project, for the current user performing the request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**|  | 

### Return type

[**Array**](array.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProjectChangeRequestConfig**
> Array GetProjectChangeRequestConfig(ctx, projectId)
Retrieves change request configuration for a project

Given a projectId, this endpoint will retrieve change request configuration for the project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**|  | 

### Return type

[**Array**](array.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MarkNotificationsAsRead**
> MarkNotificationsAsRead(ctx, body)
Mark notifications as read

Allow to select which notifications were read and saving a read date

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ModelMap**](map.md)| markNotificationsAsReadSchema | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateChangeRequestState**
> ChangeRequestStateSchema UpdateChangeRequestState(ctx, projectId, id)
This endpoint will update the state of a change request

This endpoint will update the state of a change request if the business rules allow it. The state can be one of the following: Draft, In review, Approved, Cancelled, Applied. In order to be approved, the change request must have at least one change and the number of approvals must be greater than or equal to the number of approvals required for the environment.                      Once a change request has been approved, it can be applied. Once a change request has been applied, it cannot be changed. Once a change request has been cancelled, it cannot be changed. Any change to a change request in the state of Approved will result in the state being set to In Review and the number of approvals will be reset.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**|  | 
  **id** | **string**|  | 

### Return type

[**ChangeRequestStateSchema**](changeRequestStateSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateChangeRequestTitle**
> UpdateChangeRequestTitle(ctx, projectId, id)
This endpoint will update the custom title of a change request

Change requests get a default title e.g. Change Request #1. This endpoint allows to make the title                     more meaningful and describe the intent behind the Change Request

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**|  | 
  **id** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateProjectChangeRequestConfig**
> UpdateProjectChangeRequestConfig(ctx, body, projectId, environment)
Updates change request configuration for an environment in the project

This endpoint will change the change request             configuration for a given environment, set it to either on/off and optionally configure the number of approvals needed.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateChangeRequestEnvironmentConfigSchema**](UpdateChangeRequestEnvironmentConfigSchema.md)| updateChangeRequestEnvironmentConfigSchema | 
  **projectId** | **string**|  | 
  **environment** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

