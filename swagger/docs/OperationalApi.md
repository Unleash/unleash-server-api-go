# {{classname}}

All URIs are relative to *https://us.app.unleash-hosted.com/ushosted*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetHealth**](OperationalApi.md#GetHealth) | **Get** /health | Get instance operational status

# **GetHealth**
> HealthCheckSchema GetHealth(ctx, )
Get instance operational status

This operation returns information about whether this Unleash instance is healthy and ready to serve requests or not. Typically used by your deployment orchestrator (e.g. Kubernetes, Docker Swarm, Mesos, et al.).

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**HealthCheckSchema**](healthCheckSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

