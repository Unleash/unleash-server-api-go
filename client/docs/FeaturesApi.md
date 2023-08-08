# \FeaturesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddFavoriteFeature**](FeaturesApi.md#AddFavoriteFeature) | **Post** /api/admin/projects/{projectId}/features/{featureName}/favorites | Add feature to favorites
[**AddFavoriteProject**](FeaturesApi.md#AddFavoriteProject) | **Post** /api/admin/projects/{projectId}/favorites | Add project to favorites
[**AddFeatureStrategy**](FeaturesApi.md#AddFeatureStrategy) | **Post** /api/admin/projects/{projectId}/features/{featureName}/environments/{environment}/strategies | Add a strategy to a feature toggle
[**AddTag**](FeaturesApi.md#AddTag) | **Post** /api/admin/features/{featureName}/tags | Adds a tag to a feature.
[**ArchiveFeature**](FeaturesApi.md#ArchiveFeature) | **Delete** /api/admin/projects/{projectId}/features/{featureName} | Archive a feature toggle
[**ArchiveFeatures**](FeaturesApi.md#ArchiveFeatures) | **Post** /api/admin/projects/{projectId}/archive | Archives a list of features
[**BulkToggleFeaturesEnvironmentOff**](FeaturesApi.md#BulkToggleFeaturesEnvironmentOff) | **Post** /api/admin/projects/{projectId}/bulk_features/environments/{environment}/off | Bulk disable a list of features
[**BulkToggleFeaturesEnvironmentOn**](FeaturesApi.md#BulkToggleFeaturesEnvironmentOn) | **Post** /api/admin/projects/{projectId}/bulk_features/environments/{environment}/on | Bulk enable a list of features
[**ChangeProject**](FeaturesApi.md#ChangeProject) | **Post** /api/admin/projects/{projectId}/features/{featureName}/changeProject | Move feature to project
[**CloneFeature**](FeaturesApi.md#CloneFeature) | **Post** /api/admin/projects/{projectId}/features/{featureName}/clone | Clone a feature toggle
[**CreateFeature**](FeaturesApi.md#CreateFeature) | **Post** /api/admin/projects/{projectId}/features | Add a new feature toggle
[**DeleteFeatureStrategy**](FeaturesApi.md#DeleteFeatureStrategy) | **Delete** /api/admin/projects/{projectId}/features/{featureName}/environments/{environment}/strategies/{strategyId} | Delete a strategy from a feature toggle
[**GetAllFeatureTypes**](FeaturesApi.md#GetAllFeatureTypes) | **Get** /api/admin/feature-types | Get all feature types
[**GetAllToggles**](FeaturesApi.md#GetAllToggles) | **Get** /api/admin/features | Get all feature toggles (deprecated)
[**GetEnvironmentFeatureVariants**](FeaturesApi.md#GetEnvironmentFeatureVariants) | **Get** /api/admin/projects/{projectId}/features/{featureName}/environments/{environment}/variants | Get variants for a feature in an environment
[**GetFeature**](FeaturesApi.md#GetFeature) | **Get** /api/admin/projects/{projectId}/features/{featureName} | Get a feature
[**GetFeatureEnvironment**](FeaturesApi.md#GetFeatureEnvironment) | **Get** /api/admin/projects/{projectId}/features/{featureName}/environments/{environment} | Get a feature environment
[**GetFeatureStrategies**](FeaturesApi.md#GetFeatureStrategies) | **Get** /api/admin/projects/{projectId}/features/{featureName}/environments/{environment}/strategies | Get feature toggle strategies
[**GetFeatureStrategy**](FeaturesApi.md#GetFeatureStrategy) | **Get** /api/admin/projects/{projectId}/features/{featureName}/environments/{environment}/strategies/{strategyId} | Get a strategy configuration
[**GetFeatureVariants**](FeaturesApi.md#GetFeatureVariants) | **Get** /api/admin/projects/{projectId}/features/{featureName}/variants | Retrieve variants for a feature (deprecated) 
[**GetFeatures**](FeaturesApi.md#GetFeatures) | **Get** /api/admin/projects/{projectId}/features | Get all features in a project
[**ListTags**](FeaturesApi.md#ListTags) | **Get** /api/admin/features/{featureName}/tags | Get all tags for a feature.
[**OverwriteEnvironmentFeatureVariants**](FeaturesApi.md#OverwriteEnvironmentFeatureVariants) | **Put** /api/admin/projects/{projectId}/features/{featureName}/environments/{environment}/variants | Create (overwrite) variants for a feature in an environment
[**OverwriteFeatureVariants**](FeaturesApi.md#OverwriteFeatureVariants) | **Put** /api/admin/projects/{projectId}/features/{featureName}/variants | Create (overwrite) variants for a feature toggle in all environments
[**OverwriteFeatureVariantsOnEnvironments**](FeaturesApi.md#OverwriteFeatureVariantsOnEnvironments) | **Put** /api/admin/projects/{projectId}/features/{featureName}/variants-batch | Create (overwrite) variants for a feature toggle in multiple environments
[**PatchEnvironmentsFeatureVariants**](FeaturesApi.md#PatchEnvironmentsFeatureVariants) | **Patch** /api/admin/projects/{projectId}/features/{featureName}/environments/{environment}/variants | Patch a feature&#39;s variants in an environment
[**PatchFeature**](FeaturesApi.md#PatchFeature) | **Patch** /api/admin/projects/{projectId}/features/{featureName} | Modify a feature toggle
[**PatchFeatureStrategy**](FeaturesApi.md#PatchFeatureStrategy) | **Patch** /api/admin/projects/{projectId}/features/{featureName}/environments/{environment}/strategies/{strategyId} | Change specific properties of a strategy
[**PatchFeatureVariants**](FeaturesApi.md#PatchFeatureVariants) | **Patch** /api/admin/projects/{projectId}/features/{featureName}/variants | Apply a patch to a feature&#39;s variants (in all environments).
[**RemoveFavoriteFeature**](FeaturesApi.md#RemoveFavoriteFeature) | **Delete** /api/admin/projects/{projectId}/features/{featureName}/favorites | Remove feature from favorites
[**RemoveFavoriteProject**](FeaturesApi.md#RemoveFavoriteProject) | **Delete** /api/admin/projects/{projectId}/favorites | Remove project from favorites
[**RemoveTag**](FeaturesApi.md#RemoveTag) | **Delete** /api/admin/features/{featureName}/tags/{type}/{value} | Removes a tag from a feature.
[**SetStrategySortOrder**](FeaturesApi.md#SetStrategySortOrder) | **Post** /api/admin/projects/{projectId}/features/{featureName}/environments/{environment}/strategies/set-sort-order | Set the order of strategies on the list
[**StaleFeatures**](FeaturesApi.md#StaleFeatures) | **Post** /api/admin/projects/{projectId}/stale | Mark features as stale / not stale
[**ToggleFeatureEnvironmentOff**](FeaturesApi.md#ToggleFeatureEnvironmentOff) | **Post** /api/admin/projects/{projectId}/features/{featureName}/environments/{environment}/off | Disable a feature toggle
[**ToggleFeatureEnvironmentOn**](FeaturesApi.md#ToggleFeatureEnvironmentOn) | **Post** /api/admin/projects/{projectId}/features/{featureName}/environments/{environment}/on | Enable a feature toggle
[**UpdateFeature**](FeaturesApi.md#UpdateFeature) | **Put** /api/admin/projects/{projectId}/features/{featureName} | Update a feature toggle
[**UpdateFeatureStrategy**](FeaturesApi.md#UpdateFeatureStrategy) | **Put** /api/admin/projects/{projectId}/features/{featureName}/environments/{environment}/strategies/{strategyId} | Update a strategy
[**UpdateTags**](FeaturesApi.md#UpdateTags) | **Put** /api/admin/features/{featureName}/tags | Updates multiple tags for a feature.
[**ValidateConstraint**](FeaturesApi.md#ValidateConstraint) | **Post** /api/admin/constraints/validate | Validate constraint
[**ValidateFeature**](FeaturesApi.md#ValidateFeature) | **Post** /api/admin/features/validate | Validate a feature toggle name.



## AddFavoriteFeature

> AddFavoriteFeature(ctx, projectId, featureName).Execute()

Add feature to favorites



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
    featureName := "featureName_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.FeaturesApi.AddFavoriteFeature(context.Background(), projectId, featureName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeaturesApi.AddFavoriteFeature``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 
**featureName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddFavoriteFeatureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddFavoriteProject

> AddFavoriteProject(ctx, projectId).Execute()

Add project to favorites



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
    r, err := apiClient.FeaturesApi.AddFavoriteProject(context.Background(), projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeaturesApi.AddFavoriteProject``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAddFavoriteProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddFeatureStrategy

> FeatureStrategySchema AddFeatureStrategy(ctx, projectId, featureName, environment).CreateFeatureStrategySchema(createFeatureStrategySchema).Execute()

Add a strategy to a feature toggle



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
    featureName := "featureName_example" // string | 
    environment := "environment_example" // string | 
    createFeatureStrategySchema := *openapiclient.NewCreateFeatureStrategySchema("flexibleRollout") // CreateFeatureStrategySchema | createFeatureStrategySchema

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeaturesApi.AddFeatureStrategy(context.Background(), projectId, featureName, environment).CreateFeatureStrategySchema(createFeatureStrategySchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeaturesApi.AddFeatureStrategy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddFeatureStrategy`: FeatureStrategySchema
    fmt.Fprintf(os.Stdout, "Response from `FeaturesApi.AddFeatureStrategy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 
**featureName** | **string** |  | 
**environment** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddFeatureStrategyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **createFeatureStrategySchema** | [**CreateFeatureStrategySchema**](CreateFeatureStrategySchema.md) | createFeatureStrategySchema | 

### Return type

[**FeatureStrategySchema**](FeatureStrategySchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddTag

> TagSchema AddTag(ctx, featureName).TagSchema(tagSchema).Execute()

Adds a tag to a feature.



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
    featureName := "featureName_example" // string | 
    tagSchema := *openapiclient.NewTagSchema("a-tag-value", "simple") // TagSchema | tagSchema

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeaturesApi.AddTag(context.Background(), featureName).TagSchema(tagSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeaturesApi.AddTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddTag`: TagSchema
    fmt.Fprintf(os.Stdout, "Response from `FeaturesApi.AddTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**featureName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tagSchema** | [**TagSchema**](TagSchema.md) | tagSchema | 

### Return type

[**TagSchema**](TagSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ArchiveFeature

> ArchiveFeature(ctx, projectId, featureName).Execute()

Archive a feature toggle



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
    featureName := "featureName_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.FeaturesApi.ArchiveFeature(context.Background(), projectId, featureName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeaturesApi.ArchiveFeature``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 
**featureName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiArchiveFeatureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ArchiveFeatures

> ArchiveFeatures(ctx, projectId).BatchFeaturesSchema(batchFeaturesSchema).Execute()

Archives a list of features



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
    batchFeaturesSchema := *openapiclient.NewBatchFeaturesSchema([]string{"Features_example"}) // BatchFeaturesSchema | batchFeaturesSchema

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.FeaturesApi.ArchiveFeatures(context.Background(), projectId).BatchFeaturesSchema(batchFeaturesSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeaturesApi.ArchiveFeatures``: %v\n", err)
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

Other parameters are passed through a pointer to a apiArchiveFeaturesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **batchFeaturesSchema** | [**BatchFeaturesSchema**](BatchFeaturesSchema.md) | batchFeaturesSchema | 

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


## BulkToggleFeaturesEnvironmentOff

> BulkToggleFeaturesEnvironmentOff(ctx, projectId, environment).BulkToggleFeaturesSchema(bulkToggleFeaturesSchema).Execute()

Bulk disable a list of features



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
    environment := "environment_example" // string | 
    bulkToggleFeaturesSchema := *openapiclient.NewBulkToggleFeaturesSchema([]string{"Features_example"}) // BulkToggleFeaturesSchema | bulkToggleFeaturesSchema

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.FeaturesApi.BulkToggleFeaturesEnvironmentOff(context.Background(), projectId, environment).BulkToggleFeaturesSchema(bulkToggleFeaturesSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeaturesApi.BulkToggleFeaturesEnvironmentOff``: %v\n", err)
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

Other parameters are passed through a pointer to a apiBulkToggleFeaturesEnvironmentOffRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **bulkToggleFeaturesSchema** | [**BulkToggleFeaturesSchema**](BulkToggleFeaturesSchema.md) | bulkToggleFeaturesSchema | 

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


## BulkToggleFeaturesEnvironmentOn

> BulkToggleFeaturesEnvironmentOn(ctx, projectId, environment).BulkToggleFeaturesSchema(bulkToggleFeaturesSchema).Execute()

Bulk enable a list of features



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
    environment := "environment_example" // string | 
    bulkToggleFeaturesSchema := *openapiclient.NewBulkToggleFeaturesSchema([]string{"Features_example"}) // BulkToggleFeaturesSchema | bulkToggleFeaturesSchema

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.FeaturesApi.BulkToggleFeaturesEnvironmentOn(context.Background(), projectId, environment).BulkToggleFeaturesSchema(bulkToggleFeaturesSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeaturesApi.BulkToggleFeaturesEnvironmentOn``: %v\n", err)
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

Other parameters are passed through a pointer to a apiBulkToggleFeaturesEnvironmentOnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **bulkToggleFeaturesSchema** | [**BulkToggleFeaturesSchema**](BulkToggleFeaturesSchema.md) | bulkToggleFeaturesSchema | 

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


## ChangeProject

> ChangeProject(ctx, projectId, featureName).ChangeProjectSchema(changeProjectSchema).Execute()

Move feature to project



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
    featureName := "featureName_example" // string | 
    changeProjectSchema := *openapiclient.NewChangeProjectSchema("newProject") // ChangeProjectSchema | changeProjectSchema

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.FeaturesApi.ChangeProject(context.Background(), projectId, featureName).ChangeProjectSchema(changeProjectSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeaturesApi.ChangeProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 
**featureName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChangeProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **changeProjectSchema** | [**ChangeProjectSchema**](ChangeProjectSchema.md) | changeProjectSchema | 

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


## CloneFeature

> FeatureSchema CloneFeature(ctx, projectId, featureName).CloneFeatureSchema(cloneFeatureSchema).Execute()

Clone a feature toggle



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
    featureName := "featureName_example" // string | 
    cloneFeatureSchema := *openapiclient.NewCloneFeatureSchema("new-feature") // CloneFeatureSchema | cloneFeatureSchema

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeaturesApi.CloneFeature(context.Background(), projectId, featureName).CloneFeatureSchema(cloneFeatureSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeaturesApi.CloneFeature``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CloneFeature`: FeatureSchema
    fmt.Fprintf(os.Stdout, "Response from `FeaturesApi.CloneFeature`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 
**featureName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloneFeatureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cloneFeatureSchema** | [**CloneFeatureSchema**](CloneFeatureSchema.md) | cloneFeatureSchema | 

### Return type

[**FeatureSchema**](FeatureSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateFeature

> FeatureSchema CreateFeature(ctx, projectId).CreateFeatureSchema(createFeatureSchema).Execute()

Add a new feature toggle



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
    createFeatureSchema := *openapiclient.NewCreateFeatureSchema("disable-comments") // CreateFeatureSchema | createFeatureSchema

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeaturesApi.CreateFeature(context.Background(), projectId).CreateFeatureSchema(createFeatureSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeaturesApi.CreateFeature``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateFeature`: FeatureSchema
    fmt.Fprintf(os.Stdout, "Response from `FeaturesApi.CreateFeature`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateFeatureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createFeatureSchema** | [**CreateFeatureSchema**](CreateFeatureSchema.md) | createFeatureSchema | 

### Return type

[**FeatureSchema**](FeatureSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFeatureStrategy

> DeleteFeatureStrategy(ctx, projectId, featureName, environment, strategyId).Execute()

Delete a strategy from a feature toggle



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
    featureName := "featureName_example" // string | 
    environment := "environment_example" // string | 
    strategyId := "strategyId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.FeaturesApi.DeleteFeatureStrategy(context.Background(), projectId, featureName, environment, strategyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeaturesApi.DeleteFeatureStrategy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 
**featureName** | **string** |  | 
**environment** | **string** |  | 
**strategyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFeatureStrategyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllFeatureTypes

> FeatureTypesSchema GetAllFeatureTypes(ctx).Execute()

Get all feature types



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
    resp, r, err := apiClient.FeaturesApi.GetAllFeatureTypes(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeaturesApi.GetAllFeatureTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllFeatureTypes`: FeatureTypesSchema
    fmt.Fprintf(os.Stdout, "Response from `FeaturesApi.GetAllFeatureTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllFeatureTypesRequest struct via the builder pattern


### Return type

[**FeatureTypesSchema**](FeatureTypesSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllToggles

> FeaturesSchema GetAllToggles(ctx).Execute()

Get all feature toggles (deprecated)



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
    resp, r, err := apiClient.FeaturesApi.GetAllToggles(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeaturesApi.GetAllToggles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllToggles`: FeaturesSchema
    fmt.Fprintf(os.Stdout, "Response from `FeaturesApi.GetAllToggles`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllTogglesRequest struct via the builder pattern


### Return type

[**FeaturesSchema**](FeaturesSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEnvironmentFeatureVariants

> FeatureVariantsSchema GetEnvironmentFeatureVariants(ctx, projectId, featureName, environment).Execute()

Get variants for a feature in an environment



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
    featureName := "featureName_example" // string | 
    environment := "environment_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeaturesApi.GetEnvironmentFeatureVariants(context.Background(), projectId, featureName, environment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeaturesApi.GetEnvironmentFeatureVariants``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEnvironmentFeatureVariants`: FeatureVariantsSchema
    fmt.Fprintf(os.Stdout, "Response from `FeaturesApi.GetEnvironmentFeatureVariants`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 
**featureName** | **string** |  | 
**environment** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnvironmentFeatureVariantsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**FeatureVariantsSchema**](FeatureVariantsSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFeature

> FeatureSchema GetFeature(ctx, projectId, featureName).Execute()

Get a feature



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
    featureName := "featureName_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeaturesApi.GetFeature(context.Background(), projectId, featureName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeaturesApi.GetFeature``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFeature`: FeatureSchema
    fmt.Fprintf(os.Stdout, "Response from `FeaturesApi.GetFeature`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 
**featureName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFeatureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**FeatureSchema**](FeatureSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFeatureEnvironment

> FeatureEnvironmentSchema GetFeatureEnvironment(ctx, projectId, featureName, environment).Execute()

Get a feature environment



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
    featureName := "featureName_example" // string | 
    environment := "environment_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeaturesApi.GetFeatureEnvironment(context.Background(), projectId, featureName, environment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeaturesApi.GetFeatureEnvironment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFeatureEnvironment`: FeatureEnvironmentSchema
    fmt.Fprintf(os.Stdout, "Response from `FeaturesApi.GetFeatureEnvironment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 
**featureName** | **string** |  | 
**environment** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFeatureEnvironmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**FeatureEnvironmentSchema**](FeatureEnvironmentSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFeatureStrategies

> FeatureStrategySchema GetFeatureStrategies(ctx, projectId, featureName, environment).Execute()

Get feature toggle strategies



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
    featureName := "featureName_example" // string | 
    environment := "environment_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeaturesApi.GetFeatureStrategies(context.Background(), projectId, featureName, environment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeaturesApi.GetFeatureStrategies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFeatureStrategies`: FeatureStrategySchema
    fmt.Fprintf(os.Stdout, "Response from `FeaturesApi.GetFeatureStrategies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 
**featureName** | **string** |  | 
**environment** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFeatureStrategiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**FeatureStrategySchema**](FeatureStrategySchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFeatureStrategy

> FeatureStrategySchema GetFeatureStrategy(ctx, projectId, featureName, environment, strategyId).Execute()

Get a strategy configuration



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
    featureName := "featureName_example" // string | 
    environment := "environment_example" // string | 
    strategyId := "strategyId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeaturesApi.GetFeatureStrategy(context.Background(), projectId, featureName, environment, strategyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeaturesApi.GetFeatureStrategy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFeatureStrategy`: FeatureStrategySchema
    fmt.Fprintf(os.Stdout, "Response from `FeaturesApi.GetFeatureStrategy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 
**featureName** | **string** |  | 
**environment** | **string** |  | 
**strategyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFeatureStrategyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**FeatureStrategySchema**](FeatureStrategySchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFeatureVariants

> FeatureVariantsSchema GetFeatureVariants(ctx, projectId, featureName).Execute()

Retrieve variants for a feature (deprecated) 



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
    featureName := "featureName_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeaturesApi.GetFeatureVariants(context.Background(), projectId, featureName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeaturesApi.GetFeatureVariants``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFeatureVariants`: FeatureVariantsSchema
    fmt.Fprintf(os.Stdout, "Response from `FeaturesApi.GetFeatureVariants`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 
**featureName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFeatureVariantsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**FeatureVariantsSchema**](FeatureVariantsSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFeatures

> FeaturesSchema GetFeatures(ctx, projectId).Execute()

Get all features in a project



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
    resp, r, err := apiClient.FeaturesApi.GetFeatures(context.Background(), projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeaturesApi.GetFeatures``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFeatures`: FeaturesSchema
    fmt.Fprintf(os.Stdout, "Response from `FeaturesApi.GetFeatures`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFeaturesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FeaturesSchema**](FeaturesSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTags

> TagsSchema ListTags(ctx, featureName).Execute()

Get all tags for a feature.



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
    featureName := "featureName_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeaturesApi.ListTags(context.Background(), featureName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeaturesApi.ListTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTags`: TagsSchema
    fmt.Fprintf(os.Stdout, "Response from `FeaturesApi.ListTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**featureName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TagsSchema**](TagsSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OverwriteEnvironmentFeatureVariants

> FeatureVariantsSchema OverwriteEnvironmentFeatureVariants(ctx, projectId, featureName, environment).VariantSchema(variantSchema).Execute()

Create (overwrite) variants for a feature in an environment



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
    featureName := "featureName_example" // string | 
    environment := "environment_example" // string | 
    variantSchema := []openapiclient.VariantSchema{*openapiclient.NewVariantSchema("blue_group", float32(123))} // []VariantSchema | variantsSchema

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeaturesApi.OverwriteEnvironmentFeatureVariants(context.Background(), projectId, featureName, environment).VariantSchema(variantSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeaturesApi.OverwriteEnvironmentFeatureVariants``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OverwriteEnvironmentFeatureVariants`: FeatureVariantsSchema
    fmt.Fprintf(os.Stdout, "Response from `FeaturesApi.OverwriteEnvironmentFeatureVariants`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 
**featureName** | **string** |  | 
**environment** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOverwriteEnvironmentFeatureVariantsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **variantSchema** | [**[]VariantSchema**](VariantSchema.md) | variantsSchema | 

### Return type

[**FeatureVariantsSchema**](FeatureVariantsSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OverwriteFeatureVariants

> FeatureVariantsSchema OverwriteFeatureVariants(ctx, projectId, featureName).VariantSchema(variantSchema).Execute()

Create (overwrite) variants for a feature toggle in all environments



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
    featureName := "featureName_example" // string | 
    variantSchema := []openapiclient.VariantSchema{*openapiclient.NewVariantSchema("blue_group", float32(123))} // []VariantSchema | variantsSchema

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeaturesApi.OverwriteFeatureVariants(context.Background(), projectId, featureName).VariantSchema(variantSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeaturesApi.OverwriteFeatureVariants``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OverwriteFeatureVariants`: FeatureVariantsSchema
    fmt.Fprintf(os.Stdout, "Response from `FeaturesApi.OverwriteFeatureVariants`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 
**featureName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOverwriteFeatureVariantsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **variantSchema** | [**[]VariantSchema**](VariantSchema.md) | variantsSchema | 

### Return type

[**FeatureVariantsSchema**](FeatureVariantsSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OverwriteFeatureVariantsOnEnvironments

> FeatureVariantsSchema OverwriteFeatureVariantsOnEnvironments(ctx, projectId, featureName).PushVariantsSchema(pushVariantsSchema).Execute()

Create (overwrite) variants for a feature toggle in multiple environments



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
    featureName := "featureName_example" // string | 
    pushVariantsSchema := *openapiclient.NewPushVariantsSchema() // PushVariantsSchema | pushVariantsSchema

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeaturesApi.OverwriteFeatureVariantsOnEnvironments(context.Background(), projectId, featureName).PushVariantsSchema(pushVariantsSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeaturesApi.OverwriteFeatureVariantsOnEnvironments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OverwriteFeatureVariantsOnEnvironments`: FeatureVariantsSchema
    fmt.Fprintf(os.Stdout, "Response from `FeaturesApi.OverwriteFeatureVariantsOnEnvironments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 
**featureName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOverwriteFeatureVariantsOnEnvironmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pushVariantsSchema** | [**PushVariantsSchema**](PushVariantsSchema.md) | pushVariantsSchema | 

### Return type

[**FeatureVariantsSchema**](FeatureVariantsSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchEnvironmentsFeatureVariants

> FeatureVariantsSchema PatchEnvironmentsFeatureVariants(ctx, projectId, featureName, environment).PatchSchema(patchSchema).Execute()

Patch a feature's variants in an environment



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
    featureName := "featureName_example" // string | 
    environment := "environment_example" // string | 
    patchSchema := []openapiclient.PatchSchema{*openapiclient.NewPatchSchema("Path_example", "Op_example")} // []PatchSchema | patchesSchema

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeaturesApi.PatchEnvironmentsFeatureVariants(context.Background(), projectId, featureName, environment).PatchSchema(patchSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeaturesApi.PatchEnvironmentsFeatureVariants``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchEnvironmentsFeatureVariants`: FeatureVariantsSchema
    fmt.Fprintf(os.Stdout, "Response from `FeaturesApi.PatchEnvironmentsFeatureVariants`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 
**featureName** | **string** |  | 
**environment** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchEnvironmentsFeatureVariantsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **patchSchema** | [**[]PatchSchema**](PatchSchema.md) | patchesSchema | 

### Return type

[**FeatureVariantsSchema**](FeatureVariantsSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchFeature

> FeatureSchema PatchFeature(ctx, projectId, featureName).PatchSchema(patchSchema).Execute()

Modify a feature toggle



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
    featureName := "featureName_example" // string | 
    patchSchema := []openapiclient.PatchSchema{*openapiclient.NewPatchSchema("Path_example", "Op_example")} // []PatchSchema | patchesSchema

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeaturesApi.PatchFeature(context.Background(), projectId, featureName).PatchSchema(patchSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeaturesApi.PatchFeature``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchFeature`: FeatureSchema
    fmt.Fprintf(os.Stdout, "Response from `FeaturesApi.PatchFeature`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 
**featureName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchFeatureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchSchema** | [**[]PatchSchema**](PatchSchema.md) | patchesSchema | 

### Return type

[**FeatureSchema**](FeatureSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchFeatureStrategy

> FeatureStrategySchema PatchFeatureStrategy(ctx, projectId, featureName, environment, strategyId).PatchSchema(patchSchema).Execute()

Change specific properties of a strategy



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
    featureName := "featureName_example" // string | 
    environment := "environment_example" // string | 
    strategyId := "strategyId_example" // string | 
    patchSchema := []openapiclient.PatchSchema{*openapiclient.NewPatchSchema("Path_example", "Op_example")} // []PatchSchema | patchesSchema

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeaturesApi.PatchFeatureStrategy(context.Background(), projectId, featureName, environment, strategyId).PatchSchema(patchSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeaturesApi.PatchFeatureStrategy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchFeatureStrategy`: FeatureStrategySchema
    fmt.Fprintf(os.Stdout, "Response from `FeaturesApi.PatchFeatureStrategy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 
**featureName** | **string** |  | 
**environment** | **string** |  | 
**strategyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchFeatureStrategyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **patchSchema** | [**[]PatchSchema**](PatchSchema.md) | patchesSchema | 

### Return type

[**FeatureStrategySchema**](FeatureStrategySchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchFeatureVariants

> FeatureVariantsSchema PatchFeatureVariants(ctx, projectId, featureName).PatchSchema(patchSchema).Execute()

Apply a patch to a feature's variants (in all environments).



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
    featureName := "featureName_example" // string | 
    patchSchema := []openapiclient.PatchSchema{*openapiclient.NewPatchSchema("Path_example", "Op_example")} // []PatchSchema | patchesSchema

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeaturesApi.PatchFeatureVariants(context.Background(), projectId, featureName).PatchSchema(patchSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeaturesApi.PatchFeatureVariants``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchFeatureVariants`: FeatureVariantsSchema
    fmt.Fprintf(os.Stdout, "Response from `FeaturesApi.PatchFeatureVariants`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 
**featureName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchFeatureVariantsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchSchema** | [**[]PatchSchema**](PatchSchema.md) | patchesSchema | 

### Return type

[**FeatureVariantsSchema**](FeatureVariantsSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveFavoriteFeature

> RemoveFavoriteFeature(ctx, projectId, featureName).Execute()

Remove feature from favorites



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
    featureName := "featureName_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.FeaturesApi.RemoveFavoriteFeature(context.Background(), projectId, featureName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeaturesApi.RemoveFavoriteFeature``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 
**featureName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveFavoriteFeatureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveFavoriteProject

> RemoveFavoriteProject(ctx, projectId).Execute()

Remove project from favorites



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
    r, err := apiClient.FeaturesApi.RemoveFavoriteProject(context.Background(), projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeaturesApi.RemoveFavoriteProject``: %v\n", err)
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

Other parameters are passed through a pointer to a apiRemoveFavoriteProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveTag

> RemoveTag(ctx, featureName, type_, value).Execute()

Removes a tag from a feature.



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
    featureName := "featureName_example" // string | 
    type_ := "type__example" // string | 
    value := "value_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.FeaturesApi.RemoveTag(context.Background(), featureName, type_, value).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeaturesApi.RemoveTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**featureName** | **string** |  | 
**type_** | **string** |  | 
**value** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetStrategySortOrder

> SetStrategySortOrder(ctx, projectId, featureName, environment).SetStrategySortOrderSchemaInner(setStrategySortOrderSchemaInner).Execute()

Set the order of strategies on the list

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
    featureName := "featureName_example" // string | 
    environment := "environment_example" // string | 
    setStrategySortOrderSchemaInner := []openapiclient.SetStrategySortOrderSchemaInner{*openapiclient.NewSetStrategySortOrderSchemaInner("9c40958a-daac-400e-98fb-3bb438567008", float32(1))} // []SetStrategySortOrderSchemaInner | setStrategySortOrderSchema

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.FeaturesApi.SetStrategySortOrder(context.Background(), projectId, featureName, environment).SetStrategySortOrderSchemaInner(setStrategySortOrderSchemaInner).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeaturesApi.SetStrategySortOrder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 
**featureName** | **string** |  | 
**environment** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetStrategySortOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **setStrategySortOrderSchemaInner** | [**[]SetStrategySortOrderSchemaInner**](SetStrategySortOrderSchemaInner.md) | setStrategySortOrderSchema | 

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


## StaleFeatures

> StaleFeatures(ctx, projectId).BatchStaleSchema(batchStaleSchema).Execute()

Mark features as stale / not stale



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
    batchStaleSchema := *openapiclient.NewBatchStaleSchema([]string{"my-feature-5"}, true) // BatchStaleSchema | batchStaleSchema

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.FeaturesApi.StaleFeatures(context.Background(), projectId).BatchStaleSchema(batchStaleSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeaturesApi.StaleFeatures``: %v\n", err)
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

Other parameters are passed through a pointer to a apiStaleFeaturesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **batchStaleSchema** | [**BatchStaleSchema**](BatchStaleSchema.md) | batchStaleSchema | 

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


## ToggleFeatureEnvironmentOff

> FeatureSchema ToggleFeatureEnvironmentOff(ctx, projectId, featureName, environment).Execute()

Disable a feature toggle



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
    featureName := "featureName_example" // string | 
    environment := "environment_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeaturesApi.ToggleFeatureEnvironmentOff(context.Background(), projectId, featureName, environment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeaturesApi.ToggleFeatureEnvironmentOff``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ToggleFeatureEnvironmentOff`: FeatureSchema
    fmt.Fprintf(os.Stdout, "Response from `FeaturesApi.ToggleFeatureEnvironmentOff`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 
**featureName** | **string** |  | 
**environment** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiToggleFeatureEnvironmentOffRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**FeatureSchema**](FeatureSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ToggleFeatureEnvironmentOn

> FeatureSchema ToggleFeatureEnvironmentOn(ctx, projectId, featureName, environment).Execute()

Enable a feature toggle



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
    featureName := "featureName_example" // string | 
    environment := "environment_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeaturesApi.ToggleFeatureEnvironmentOn(context.Background(), projectId, featureName, environment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeaturesApi.ToggleFeatureEnvironmentOn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ToggleFeatureEnvironmentOn`: FeatureSchema
    fmt.Fprintf(os.Stdout, "Response from `FeaturesApi.ToggleFeatureEnvironmentOn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 
**featureName** | **string** |  | 
**environment** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiToggleFeatureEnvironmentOnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**FeatureSchema**](FeatureSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFeature

> FeatureSchema UpdateFeature(ctx, projectId, featureName).UpdateFeatureSchema(updateFeatureSchema).Execute()

Update a feature toggle



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
    featureName := "featureName_example" // string | 
    updateFeatureSchema := *openapiclient.NewUpdateFeatureSchema() // UpdateFeatureSchema | updateFeatureSchema

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeaturesApi.UpdateFeature(context.Background(), projectId, featureName).UpdateFeatureSchema(updateFeatureSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeaturesApi.UpdateFeature``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateFeature`: FeatureSchema
    fmt.Fprintf(os.Stdout, "Response from `FeaturesApi.UpdateFeature`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 
**featureName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFeatureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateFeatureSchema** | [**UpdateFeatureSchema**](UpdateFeatureSchema.md) | updateFeatureSchema | 

### Return type

[**FeatureSchema**](FeatureSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFeatureStrategy

> FeatureStrategySchema UpdateFeatureStrategy(ctx, projectId, featureName, environment, strategyId).UpdateFeatureStrategySchema(updateFeatureStrategySchema).Execute()

Update a strategy



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
    featureName := "featureName_example" // string | 
    environment := "environment_example" // string | 
    strategyId := "strategyId_example" // string | 
    updateFeatureStrategySchema := *openapiclient.NewUpdateFeatureStrategySchema() // UpdateFeatureStrategySchema | updateFeatureStrategySchema

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeaturesApi.UpdateFeatureStrategy(context.Background(), projectId, featureName, environment, strategyId).UpdateFeatureStrategySchema(updateFeatureStrategySchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeaturesApi.UpdateFeatureStrategy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateFeatureStrategy`: FeatureStrategySchema
    fmt.Fprintf(os.Stdout, "Response from `FeaturesApi.UpdateFeatureStrategy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 
**featureName** | **string** |  | 
**environment** | **string** |  | 
**strategyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFeatureStrategyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **updateFeatureStrategySchema** | [**UpdateFeatureStrategySchema**](UpdateFeatureStrategySchema.md) | updateFeatureStrategySchema | 

### Return type

[**FeatureStrategySchema**](FeatureStrategySchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTags

> TagsSchema UpdateTags(ctx, featureName).UpdateTagsSchema(updateTagsSchema).Execute()

Updates multiple tags for a feature.



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
    featureName := "featureName_example" // string | 
    updateTagsSchema := *openapiclient.NewUpdateTagsSchema([]openapiclient.TagSchema{*openapiclient.NewTagSchema("a-tag-value", "simple")}, []openapiclient.TagSchema{*openapiclient.NewTagSchema("a-tag-value", "simple")}) // UpdateTagsSchema | updateTagsSchema

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeaturesApi.UpdateTags(context.Background(), featureName).UpdateTagsSchema(updateTagsSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeaturesApi.UpdateTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTags`: TagsSchema
    fmt.Fprintf(os.Stdout, "Response from `FeaturesApi.UpdateTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**featureName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateTagsSchema** | [**UpdateTagsSchema**](UpdateTagsSchema.md) | updateTagsSchema | 

### Return type

[**TagsSchema**](TagsSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateConstraint

> ValidateConstraint(ctx).ConstraintSchema(constraintSchema).Execute()

Validate constraint



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
    constraintSchema := *openapiclient.NewConstraintSchema("appName", "IN") // ConstraintSchema | constraintSchema

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.FeaturesApi.ValidateConstraint(context.Background()).ConstraintSchema(constraintSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeaturesApi.ValidateConstraint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateConstraintRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **constraintSchema** | [**ConstraintSchema**](ConstraintSchema.md) | constraintSchema | 

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


## ValidateFeature

> ValidateFeature(ctx).ValidateFeatureSchema(validateFeatureSchema).Execute()

Validate a feature toggle name.



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
    validateFeatureSchema := *openapiclient.NewValidateFeatureSchema("my-feature-3") // ValidateFeatureSchema | validateFeatureSchema

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.FeaturesApi.ValidateFeature(context.Background()).ValidateFeatureSchema(validateFeatureSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeaturesApi.ValidateFeature``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateFeatureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **validateFeatureSchema** | [**ValidateFeatureSchema**](ValidateFeatureSchema.md) | validateFeatureSchema | 

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

