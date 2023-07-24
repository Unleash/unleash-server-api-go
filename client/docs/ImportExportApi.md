# \ImportExportApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CallImport**](ImportExportApi.md#CallImport) | **Post** /api/admin/state/import | Import state (deprecated)
[**Export**](ImportExportApi.md#Export) | **Get** /api/admin/state/export | Export state (deprecated)



## CallImport

> CallImport(ctx).RequestBody(requestBody).Execute()

Import state (deprecated)



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
    requestBody := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} | stateSchema

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ImportExportApi.CallImport(context.Background()).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportExportApi.CallImport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCallImportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **map[string]interface{}** | stateSchema | 

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


## Export

> StateSchema Export(ctx).Format(format).Download(download).Strategies(strategies).FeatureToggles(featureToggles).Projects(projects).Tags(tags).Environments(environments).Execute()

Export state (deprecated)



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
    format := "format_example" // string | Desired export format. Must be either `json` or `yaml`. (optional) (default to "json")
    download := *openapiclient.NewExportDownloadParameter() // ExportDownloadParameter | Whether exported data should be downloaded as a file. (optional) (default to false)
    strategies := *openapiclient.NewExportStrategiesParameter() // ExportStrategiesParameter | Whether strategies should be included in the exported data. (optional) (default to true)
    featureToggles := *openapiclient.NewExportStrategiesParameter() // ExportStrategiesParameter | Whether feature toggles should be included in the exported data. (optional) (default to true)
    projects := *openapiclient.NewExportStrategiesParameter() // ExportStrategiesParameter | Whether projects should be included in the exported data. (optional) (default to true)
    tags := *openapiclient.NewExportStrategiesParameter() // ExportStrategiesParameter | Whether tag types, tags, and feature_tags should be included in the exported data. (optional) (default to true)
    environments := *openapiclient.NewExportStrategiesParameter() // ExportStrategiesParameter | Whether environments should be included in the exported data. (optional) (default to true)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImportExportApi.Export(context.Background()).Format(format).Download(download).Strategies(strategies).FeatureToggles(featureToggles).Projects(projects).Tags(tags).Environments(environments).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportExportApi.Export``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Export`: StateSchema
    fmt.Fprintf(os.Stdout, "Response from `ImportExportApi.Export`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | **string** | Desired export format. Must be either &#x60;json&#x60; or &#x60;yaml&#x60;. | [default to &quot;json&quot;]
 **download** | [**ExportDownloadParameter**](ExportDownloadParameter.md) | Whether exported data should be downloaded as a file. | [default to false]
 **strategies** | [**ExportStrategiesParameter**](ExportStrategiesParameter.md) | Whether strategies should be included in the exported data. | [default to true]
 **featureToggles** | [**ExportStrategiesParameter**](ExportStrategiesParameter.md) | Whether feature toggles should be included in the exported data. | [default to true]
 **projects** | [**ExportStrategiesParameter**](ExportStrategiesParameter.md) | Whether projects should be included in the exported data. | [default to true]
 **tags** | [**ExportStrategiesParameter**](ExportStrategiesParameter.md) | Whether tag types, tags, and feature_tags should be included in the exported data. | [default to true]
 **environments** | [**ExportStrategiesParameter**](ExportStrategiesParameter.md) | Whether environments should be included in the exported data. | [default to true]

### Return type

[**StateSchema**](StateSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

