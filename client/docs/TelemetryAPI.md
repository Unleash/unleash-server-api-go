# \TelemetryAPI

All URIs are relative to *https://us.app.unleash-hosted.com/ushosted*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTelemetrySettings**](TelemetryAPI.md#GetTelemetrySettings) | **Get** /api/admin/telemetry/settings | Get telemetry settings



## GetTelemetrySettings

> TelemetrySettingsSchema GetTelemetrySettings(ctx).Execute()

Get telemetry settings



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
    resp, r, err := apiClient.TelemetryAPI.GetTelemetrySettings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TelemetryAPI.GetTelemetrySettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTelemetrySettings`: TelemetrySettingsSchema
    fmt.Fprintf(os.Stdout, "Response from `TelemetryAPI.GetTelemetrySettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTelemetrySettingsRequest struct via the builder pattern


### Return type

[**TelemetrySettingsSchema**](TelemetrySettingsSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

