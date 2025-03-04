# \ProjectsAPI

All URIs are relative to *https://us.app.unleash-hosted.com/ushosted*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateProject**](ProjectsAPI.md#CreateProject) | **Post** /api/admin/projects | Create project
[**DeleteProject**](ProjectsAPI.md#DeleteProject) | **Delete** /api/admin/projects/{projectId} | Delete project
[**GetProjectAccess**](ProjectsAPI.md#GetProjectAccess) | **Get** /api/admin/projects/{projectId}/access | Get users and groups in project
[**GetProjectOverview**](ProjectsAPI.md#GetProjectOverview) | **Get** /api/admin/projects/{projectId}/overview | Get an overview of a project.
[**GetProjects**](ProjectsAPI.md#GetProjects) | **Get** /api/admin/projects | Get a list of all projects.
[**SetProjectAccess**](ProjectsAPI.md#SetProjectAccess) | **Put** /api/admin/projects/{projectId}/access | Set users and groups to roles in the current project
[**UpdateProject**](ProjectsAPI.md#UpdateProject) | **Put** /api/admin/projects/{projectId} | Update project
[**UpdateProjectEnterpriseSettings**](ProjectsAPI.md#UpdateProjectEnterpriseSettings) | **Put** /api/admin/projects/{projectId}/settings | Update project enterprise settings



## CreateProject

> ProjectCreatedSchema CreateProject(ctx).CreateProjectSchema(createProjectSchema).Execute()

Create project



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
	createProjectSchema := *openapiclient.NewCreateProjectSchema("Pet shop") // CreateProjectSchema | createProjectSchema

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.CreateProject(context.Background()).CreateProjectSchema(createProjectSchema).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.CreateProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateProject`: ProjectCreatedSchema
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.CreateProject`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createProjectSchema** | [**CreateProjectSchema**](CreateProjectSchema.md) | createProjectSchema | 

### Return type

[**ProjectCreatedSchema**](ProjectCreatedSchema.md)

### Authorization

[bearerToken](../README.md#bearerToken), [apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteProject

> DeleteProject(ctx, projectId).Execute()

Delete project



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
	projectId := "projectId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectsAPI.DeleteProject(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.DeleteProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProjectRequest struct via the builder pattern


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


## GetProjectAccess

> ProjectAccessSchema GetProjectAccess(ctx, projectId).Execute()

Get users and groups in project



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
	projectId := "projectId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.GetProjectAccess(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.GetProjectAccess``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectAccess`: ProjectAccessSchema
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.GetProjectAccess`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectAccessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProjectAccessSchema**](ProjectAccessSchema.md)

### Authorization

[bearerToken](../README.md#bearerToken), [apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectOverview

> ProjectOverviewSchema GetProjectOverview(ctx, projectId).Execute()

Get an overview of a project.



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
	projectId := "projectId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.GetProjectOverview(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.GetProjectOverview``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectOverview`: ProjectOverviewSchema
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.GetProjectOverview`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectOverviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProjectOverviewSchema**](ProjectOverviewSchema.md)

### Authorization

[bearerToken](../README.md#bearerToken), [apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjects

> ProjectsSchema GetProjects(ctx).Archived(archived).Execute()

Get a list of all projects.



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
	archived := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.GetProjects(context.Background()).Archived(archived).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.GetProjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjects`: ProjectsSchema
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.GetProjects`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **archived** | **bool** |  | 

### Return type

[**ProjectsSchema**](ProjectsSchema.md)

### Authorization

[bearerToken](../README.md#bearerToken), [apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetProjectAccess

> SetProjectAccess(ctx, projectId).ProjectAccessConfigurationSchema(projectAccessConfigurationSchema).Execute()

Set users and groups to roles in the current project



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
	projectId := "projectId_example" // string | 
	projectAccessConfigurationSchema := *openapiclient.NewProjectAccessConfigurationSchema([]openapiclient.ProjectAccessConfigurationSchemaRolesInner{*openapiclient.NewProjectAccessConfigurationSchemaRolesInner()}) // ProjectAccessConfigurationSchema | projectAccessConfigurationSchema

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectsAPI.SetProjectAccess(context.Background(), projectId).ProjectAccessConfigurationSchema(projectAccessConfigurationSchema).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.SetProjectAccess``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetProjectAccessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **projectAccessConfigurationSchema** | [**ProjectAccessConfigurationSchema**](ProjectAccessConfigurationSchema.md) | projectAccessConfigurationSchema | 

### Return type

 (empty response body)

### Authorization

[bearerToken](../README.md#bearerToken), [apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProject

> UpdateProject(ctx, projectId).UpdateProjectSchema(updateProjectSchema).Execute()

Update project



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
	projectId := "projectId_example" // string | 
	updateProjectSchema := *openapiclient.NewUpdateProjectSchema("my-renamed-project") // UpdateProjectSchema | updateProjectSchema

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectsAPI.UpdateProject(context.Background(), projectId).UpdateProjectSchema(updateProjectSchema).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.UpdateProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateProjectSchema** | [**UpdateProjectSchema**](UpdateProjectSchema.md) | updateProjectSchema | 

### Return type

 (empty response body)

### Authorization

[bearerToken](../README.md#bearerToken), [apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProjectEnterpriseSettings

> UpdateProjectEnterpriseSettings(ctx, projectId).UpdateProjectEnterpriseSettingsSchema(updateProjectEnterpriseSettingsSchema).Execute()

Update project enterprise settings



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
	projectId := "projectId_example" // string | 
	updateProjectEnterpriseSettingsSchema := *openapiclient.NewUpdateProjectEnterpriseSettingsSchema() // UpdateProjectEnterpriseSettingsSchema | updateProjectEnterpriseSettingsSchema

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectsAPI.UpdateProjectEnterpriseSettings(context.Background(), projectId).UpdateProjectEnterpriseSettingsSchema(updateProjectEnterpriseSettingsSchema).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.UpdateProjectEnterpriseSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProjectEnterpriseSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateProjectEnterpriseSettingsSchema** | [**UpdateProjectEnterpriseSettingsSchema**](UpdateProjectEnterpriseSettingsSchema.md) | updateProjectEnterpriseSettingsSchema | 

### Return type

 (empty response body)

### Authorization

[bearerToken](../README.md#bearerToken), [apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

