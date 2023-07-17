# {{classname}}

All URIs are relative to *https://us.app.unleash-hosted.com/ushosted*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetInstanceAdminStats**](InstanceAdminApi.md#GetInstanceAdminStats) | **Get** /api/admin/instance-admin/statistics | Instance usage statistics
[**GetInstanceAdminStatsCsv**](InstanceAdminApi.md#GetInstanceAdminStatsCsv) | **Get** /api/admin/instance-admin/statistics/csv | Instance usage statistics

# **GetInstanceAdminStats**
> InstanceAdminStatsSchema GetInstanceAdminStats(ctx, )
Instance usage statistics

Provides statistics about various features of Unleash to allow for reporting of usage for self-hosted customers. The response contains data such as the number of users, groups, features, strategies, versions, etc.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InstanceAdminStatsSchema**](instanceAdminStatsSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetInstanceAdminStatsCsv**
> string GetInstanceAdminStatsCsv(ctx, )
Instance usage statistics

Provides statistics about various features of Unleash to allow for reporting of usage for self-hosted customers. The response contains data such as the number of users, groups, features, strategies, versions, etc.

### Required Parameters
This endpoint does not need any parameter.

### Return type

**string**

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

