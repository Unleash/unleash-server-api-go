# \UnstableApi

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



## AddChangeRequestComment

> AddChangeRequestComment(ctx, projectId, id).ChangeRequestAddCommentSchema(changeRequestAddCommentSchema).Execute()

This endpoint will add a comment to a change request



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Unleash/unleash-server-api-go/client"
)

func main() {
    projectId := "projectId_example" // string | 
    id := "id_example" // string | 
    changeRequestAddCommentSchema := *openapiclient.NewChangeRequestAddCommentSchema("This is a comment") // ChangeRequestAddCommentSchema | changeRequestAddCommentSchema

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UnstableApi.AddChangeRequestComment(context.Background(), projectId, id).ChangeRequestAddCommentSchema(changeRequestAddCommentSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UnstableApi.AddChangeRequestComment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddChangeRequestCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **changeRequestAddCommentSchema** | [**ChangeRequestAddCommentSchema**](ChangeRequestAddCommentSchema.md) | changeRequestAddCommentSchema | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChangeRequest

> ChangeRequestSchema ChangeRequest(ctx, projectId, environment).ChangeRequestCreateSchema(changeRequestCreateSchema).Execute()

Create/Add change to a change request



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Unleash/unleash-server-api-go/client"
)

func main() {
    projectId := "projectId_example" // string | 
    environment := "environment_example" // string | 
    changeRequestCreateSchema := *openapiclient.NewChangeRequestCreateSchema("my-best-feature") // ChangeRequestCreateSchema | changeRequestOneOrManyCreateSchema

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UnstableApi.ChangeRequest(context.Background(), projectId, environment).ChangeRequestCreateSchema(changeRequestCreateSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UnstableApi.ChangeRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ChangeRequest`: ChangeRequestSchema
    fmt.Fprintf(os.Stdout, "Response from `UnstableApi.ChangeRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 
**environment** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChangeRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **changeRequestCreateSchema** | [**ChangeRequestCreateSchema**](ChangeRequestCreateSchema.md) | changeRequestOneOrManyCreateSchema | 

### Return type

[**ChangeRequestSchema**](ChangeRequestSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteChange

> DeleteChange(ctx, projectId, changeRequestId, changeId).Execute()

Discards a change from a change request by change id



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Unleash/unleash-server-api-go/client"
)

func main() {
    projectId := "projectId_example" // string | 
    changeRequestId := "changeRequestId_example" // string | 
    changeId := "changeId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UnstableApi.DeleteChange(context.Background(), projectId, changeRequestId, changeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UnstableApi.DeleteChange``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 
**changeRequestId** | **string** |  | 
**changeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteChangeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteChangeRequest

> DeleteChangeRequest(ctx, projectId, id).Execute()

Deletes a change request by id



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Unleash/unleash-server-api-go/client"
)

func main() {
    projectId := "projectId_example" // string | 
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UnstableApi.DeleteChangeRequest(context.Background(), projectId, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UnstableApi.DeleteChangeRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteChangeRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditChange

> ChangeRequestSchema EditChange(ctx, projectId, changeRequestId, changeId).ChangeRequestCreateSchema(changeRequestCreateSchema).Execute()

Edits a single change in a change request



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Unleash/unleash-server-api-go/client"
)

func main() {
    projectId := "projectId_example" // string | 
    changeRequestId := "changeRequestId_example" // string | 
    changeId := "changeId_example" // string | 
    changeRequestCreateSchema := *openapiclient.NewChangeRequestCreateSchema("my-best-feature") // ChangeRequestCreateSchema | changeRequestCreateSchema

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UnstableApi.EditChange(context.Background(), projectId, changeRequestId, changeId).ChangeRequestCreateSchema(changeRequestCreateSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UnstableApi.EditChange``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditChange`: ChangeRequestSchema
    fmt.Fprintf(os.Stdout, "Response from `UnstableApi.EditChange`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 
**changeRequestId** | **string** |  | 
**changeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditChangeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **changeRequestCreateSchema** | [**ChangeRequestCreateSchema**](ChangeRequestCreateSchema.md) | changeRequestCreateSchema | 

### Return type

[**ChangeRequestSchema**](ChangeRequestSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAdvancedPlayground

> AdvancedPlaygroundResponseSchema GetAdvancedPlayground(ctx).AdvancedPlaygroundRequestSchema(advancedPlaygroundRequestSchema).Execute()

Batch evaluate an Unleash context against a set of environments and projects.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Unleash/unleash-server-api-go/client"
)

func main() {
    advancedPlaygroundRequestSchema := *openapiclient.NewAdvancedPlaygroundRequestSchema([]string{"Environments_example"}, *openapiclient.NewSdkContextSchema("My cool application.")) // AdvancedPlaygroundRequestSchema | advancedPlaygroundRequestSchema

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UnstableApi.GetAdvancedPlayground(context.Background()).AdvancedPlaygroundRequestSchema(advancedPlaygroundRequestSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UnstableApi.GetAdvancedPlayground``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAdvancedPlayground`: AdvancedPlaygroundResponseSchema
    fmt.Fprintf(os.Stdout, "Response from `UnstableApi.GetAdvancedPlayground`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAdvancedPlaygroundRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **advancedPlaygroundRequestSchema** | [**AdvancedPlaygroundRequestSchema**](AdvancedPlaygroundRequestSchema.md) | advancedPlaygroundRequestSchema | 

### Return type

[**AdvancedPlaygroundResponseSchema**](AdvancedPlaygroundResponseSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChangeRequest

> ChangeRequestSchema GetChangeRequest(ctx, projectId, id).Execute()

Retrieves one change request by id



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Unleash/unleash-server-api-go/client"
)

func main() {
    projectId := "projectId_example" // string | 
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UnstableApi.GetChangeRequest(context.Background(), projectId, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UnstableApi.GetChangeRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetChangeRequest`: ChangeRequestSchema
    fmt.Fprintf(os.Stdout, "Response from `UnstableApi.GetChangeRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetChangeRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ChangeRequestSchema**](ChangeRequestSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChangeRequestsForProject

> []ChangeRequestSchema GetChangeRequestsForProject(ctx, projectId).Execute()

Retrieves all change requests for a project



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Unleash/unleash-server-api-go/client"
)

func main() {
    projectId := "projectId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UnstableApi.GetChangeRequestsForProject(context.Background(), projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UnstableApi.GetChangeRequestsForProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetChangeRequestsForProject`: []ChangeRequestSchema
    fmt.Fprintf(os.Stdout, "Response from `UnstableApi.GetChangeRequestsForProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetChangeRequestsForProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ChangeRequestSchema**](ChangeRequestSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoginHistory

> LoginHistorySchema GetLoginHistory(ctx).Execute()

Get all login events.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Unleash/unleash-server-api-go/client"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UnstableApi.GetLoginHistory(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UnstableApi.GetLoginHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLoginHistory`: LoginHistorySchema
    fmt.Fprintf(os.Stdout, "Response from `UnstableApi.GetLoginHistory`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLoginHistoryRequest struct via the builder pattern


### Return type

[**LoginHistorySchema**](LoginHistorySchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNotifications

> []NotificationsSchemaInner GetNotifications(ctx).Execute()

Retrieves a list of notifications



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Unleash/unleash-server-api-go/client"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UnstableApi.GetNotifications(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UnstableApi.GetNotifications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNotifications`: []NotificationsSchemaInner
    fmt.Fprintf(os.Stdout, "Response from `UnstableApi.GetNotifications`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetNotificationsRequest struct via the builder pattern


### Return type

[**[]NotificationsSchemaInner**](NotificationsSchemaInner.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOpenChangeRequestsForUser

> []ChangeRequestSchema GetOpenChangeRequestsForUser(ctx, projectId).Execute()

Retrieves pending change requests in configured environments



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Unleash/unleash-server-api-go/client"
)

func main() {
    projectId := "projectId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UnstableApi.GetOpenChangeRequestsForUser(context.Background(), projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UnstableApi.GetOpenChangeRequestsForUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOpenChangeRequestsForUser`: []ChangeRequestSchema
    fmt.Fprintf(os.Stdout, "Response from `UnstableApi.GetOpenChangeRequestsForUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOpenChangeRequestsForUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ChangeRequestSchema**](ChangeRequestSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPendingChangeRequestsForFeature

> []ChangeRequestSchema GetPendingChangeRequestsForFeature(ctx, projectId, featureName).Execute()

Retrieves all pending change requests referencing a feature in the project



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Unleash/unleash-server-api-go/client"
)

func main() {
    projectId := "projectId_example" // string | 
    featureName := "featureName_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UnstableApi.GetPendingChangeRequestsForFeature(context.Background(), projectId, featureName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UnstableApi.GetPendingChangeRequestsForFeature``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPendingChangeRequestsForFeature`: []ChangeRequestSchema
    fmt.Fprintf(os.Stdout, "Response from `UnstableApi.GetPendingChangeRequestsForFeature`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 
**featureName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPendingChangeRequestsForFeatureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]ChangeRequestSchema**](ChangeRequestSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPendingChangeRequestsForUser

> []ChangeRequestSchema GetPendingChangeRequestsForUser(ctx, projectId).Execute()

Retrieves pending change requests in configured environments



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Unleash/unleash-server-api-go/client"
)

func main() {
    projectId := "projectId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UnstableApi.GetPendingChangeRequestsForUser(context.Background(), projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UnstableApi.GetPendingChangeRequestsForUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPendingChangeRequestsForUser`: []ChangeRequestSchema
    fmt.Fprintf(os.Stdout, "Response from `UnstableApi.GetPendingChangeRequestsForUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPendingChangeRequestsForUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ChangeRequestSchema**](ChangeRequestSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectChangeRequestConfig

> []ChangeRequestEnvironmentConfigSchema GetProjectChangeRequestConfig(ctx, projectId).Execute()

Retrieves change request configuration for a project



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Unleash/unleash-server-api-go/client"
)

func main() {
    projectId := "projectId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UnstableApi.GetProjectChangeRequestConfig(context.Background(), projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UnstableApi.GetProjectChangeRequestConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProjectChangeRequestConfig`: []ChangeRequestEnvironmentConfigSchema
    fmt.Fprintf(os.Stdout, "Response from `UnstableApi.GetProjectChangeRequestConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectChangeRequestConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ChangeRequestEnvironmentConfigSchema**](ChangeRequestEnvironmentConfigSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarkNotificationsAsRead

> MarkNotificationsAsRead(ctx).RequestBody(requestBody).Execute()

Mark notifications as read



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Unleash/unleash-server-api-go/client"
)

func main() {
    requestBody := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} | markNotificationsAsReadSchema

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UnstableApi.MarkNotificationsAsRead(context.Background()).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UnstableApi.MarkNotificationsAsRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarkNotificationsAsReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **map[string]interface{}** | markNotificationsAsReadSchema | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateChangeRequestState

> ChangeRequestStateSchema UpdateChangeRequestState(ctx, projectId, id).Execute()

This endpoint will update the state of a change request



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Unleash/unleash-server-api-go/client"
)

func main() {
    projectId := "projectId_example" // string | 
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UnstableApi.UpdateChangeRequestState(context.Background(), projectId, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UnstableApi.UpdateChangeRequestState``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateChangeRequestState`: ChangeRequestStateSchema
    fmt.Fprintf(os.Stdout, "Response from `UnstableApi.UpdateChangeRequestState`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateChangeRequestStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ChangeRequestStateSchema**](ChangeRequestStateSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateChangeRequestTitle

> UpdateChangeRequestTitle(ctx, projectId, id).Execute()

This endpoint will update the custom title of a change request



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Unleash/unleash-server-api-go/client"
)

func main() {
    projectId := "projectId_example" // string | 
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UnstableApi.UpdateChangeRequestTitle(context.Background(), projectId, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UnstableApi.UpdateChangeRequestTitle``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateChangeRequestTitleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProjectChangeRequestConfig

> UpdateProjectChangeRequestConfig(ctx, projectId, environment).UpdateChangeRequestEnvironmentConfigSchema(updateChangeRequestEnvironmentConfigSchema).Execute()

Updates change request configuration for an environment in the project



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Unleash/unleash-server-api-go/client"
)

func main() {
    projectId := "projectId_example" // string | 
    environment := "environment_example" // string | 
    updateChangeRequestEnvironmentConfigSchema := *openapiclient.NewUpdateChangeRequestEnvironmentConfigSchema(false) // UpdateChangeRequestEnvironmentConfigSchema | updateChangeRequestEnvironmentConfigSchema

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UnstableApi.UpdateProjectChangeRequestConfig(context.Background(), projectId, environment).UpdateChangeRequestEnvironmentConfigSchema(updateChangeRequestEnvironmentConfigSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UnstableApi.UpdateProjectChangeRequestConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 
**environment** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProjectChangeRequestConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateChangeRequestEnvironmentConfigSchema** | [**UpdateChangeRequestEnvironmentConfigSchema**](UpdateChangeRequestEnvironmentConfigSchema.md) | updateChangeRequestEnvironmentConfigSchema | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

