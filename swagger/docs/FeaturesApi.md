# {{classname}}

All URIs are relative to *https://us.app.unleash-hosted.com/ushosted*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddFavoriteFeature**](FeaturesApi.md#AddFavoriteFeature) | **Post** /api/admin/projects/{projectId}/features/{featureName}/favorites | Add feature to favorites
[**AddFavoriteProject**](FeaturesApi.md#AddFavoriteProject) | **Post** /api/admin/projects/{projectId}/favorites | Add project to favorites
[**AddFeatureStrategy**](FeaturesApi.md#AddFeatureStrategy) | **Post** /api/admin/projects/{projectId}/features/{featureName}/environments/{environment}/strategies | Add a strategy to a feature toggle.
[**AddTag**](FeaturesApi.md#AddTag) | **Post** /api/admin/features/{featureName}/tags | Adds a tag to a feature.
[**ArchiveFeature**](FeaturesApi.md#ArchiveFeature) | **Delete** /api/admin/projects/{projectId}/features/{featureName} | Archive a feature.
[**ArchiveFeatures**](FeaturesApi.md#ArchiveFeatures) | **Post** /api/admin/projects/{projectId}/archive | Archives a list of features
[**BulkToggleFeaturesEnvironmentOff**](FeaturesApi.md#BulkToggleFeaturesEnvironmentOff) | **Post** /api/admin/projects/{projectId}/bulk_features/environments/{environment}/off | Bulk disabled a list of features.
[**BulkToggleFeaturesEnvironmentOn**](FeaturesApi.md#BulkToggleFeaturesEnvironmentOn) | **Post** /api/admin/projects/{projectId}/bulk_features/environments/{environment}/on | Bulk enable a list of features.
[**ChangeProject**](FeaturesApi.md#ChangeProject) | **Post** /api/admin/projects/{projectId}/features/{featureName}/changeProject | Move feature to project
[**CloneFeature**](FeaturesApi.md#CloneFeature) | **Post** /api/admin/projects/{projectId}/features/{featureName}/clone | 
[**CreateFeature**](FeaturesApi.md#CreateFeature) | **Post** /api/admin/projects/{projectId}/features | 
[**DeleteFeatureStrategy**](FeaturesApi.md#DeleteFeatureStrategy) | **Delete** /api/admin/projects/{projectId}/features/{featureName}/environments/{environment}/strategies/{strategyId} | Delete a strategy from a feature toggle.
[**GetAllFeatureTypes**](FeaturesApi.md#GetAllFeatureTypes) | **Get** /api/admin/feature-types | Get all feature types
[**GetAllToggles**](FeaturesApi.md#GetAllToggles) | **Get** /api/admin/features | Get all features (deprecated)
[**GetEnvironmentFeatureVariants**](FeaturesApi.md#GetEnvironmentFeatureVariants) | **Get** /api/admin/projects/{projectId}/features/{featureName}/environments/{environment}/variants | Get variants for a feature in an environment
[**GetFeature**](FeaturesApi.md#GetFeature) | **Get** /api/admin/projects/{projectId}/features/{featureName} | Get a feature.
[**GetFeatureEnvironment**](FeaturesApi.md#GetFeatureEnvironment) | **Get** /api/admin/projects/{projectId}/features/{featureName}/environments/{environment} | Get a feature environment.
[**GetFeatureStrategies**](FeaturesApi.md#GetFeatureStrategies) | **Get** /api/admin/projects/{projectId}/features/{featureName}/environments/{environment}/strategies | Get feature toggle strategies.
[**GetFeatureStrategy**](FeaturesApi.md#GetFeatureStrategy) | **Get** /api/admin/projects/{projectId}/features/{featureName}/environments/{environment}/strategies/{strategyId} | Get a strategy configuration.
[**GetFeatureVariants**](FeaturesApi.md#GetFeatureVariants) | **Get** /api/admin/projects/{projectId}/features/{featureName}/variants | Retrieve variants for a feature (deprecated) 
[**GetFeatures**](FeaturesApi.md#GetFeatures) | **Get** /api/admin/projects/{projectId}/features | 
[**ListTags**](FeaturesApi.md#ListTags) | **Get** /api/admin/features/{featureName}/tags | Get all tags for a feature.
[**OverwriteEnvironmentFeatureVariants**](FeaturesApi.md#OverwriteEnvironmentFeatureVariants) | **Put** /api/admin/projects/{projectId}/features/{featureName}/environments/{environment}/variants | Create (overwrite) variants for a feature in an environment
[**OverwriteFeatureVariants**](FeaturesApi.md#OverwriteFeatureVariants) | **Put** /api/admin/projects/{projectId}/features/{featureName}/variants | Create (overwrite) variants for a feature toggle in all environments
[**OverwriteFeatureVariantsOnEnvironments**](FeaturesApi.md#OverwriteFeatureVariantsOnEnvironments) | **Put** /api/admin/projects/{projectId}/features/{featureName}/variants-batch | 
[**PatchEnvironmentsFeatureVariants**](FeaturesApi.md#PatchEnvironmentsFeatureVariants) | **Patch** /api/admin/projects/{projectId}/features/{featureName}/environments/{environment}/variants | Patch a feature&#x27;s variants in an environment
[**PatchFeature**](FeaturesApi.md#PatchFeature) | **Patch** /api/admin/projects/{projectId}/features/{featureName} | 
[**PatchFeatureStrategy**](FeaturesApi.md#PatchFeatureStrategy) | **Patch** /api/admin/projects/{projectId}/features/{featureName}/environments/{environment}/strategies/{strategyId} | Change specific properties of a strategy.
[**PatchFeatureVariants**](FeaturesApi.md#PatchFeatureVariants) | **Patch** /api/admin/projects/{projectId}/features/{featureName}/variants | Apply a patch to a feature&#x27;s variants (in all environments).
[**RemoveFavoriteFeature**](FeaturesApi.md#RemoveFavoriteFeature) | **Delete** /api/admin/projects/{projectId}/features/{featureName}/favorites | Remove feature from favorites
[**RemoveFavoriteProject**](FeaturesApi.md#RemoveFavoriteProject) | **Delete** /api/admin/projects/{projectId}/favorites | Remove project from favorites
[**RemoveTag**](FeaturesApi.md#RemoveTag) | **Delete** /api/admin/features/{featureName}/tags/{type}/{value} | Removes a tag from a feature.
[**SetStrategySortOrder**](FeaturesApi.md#SetStrategySortOrder) | **Post** /api/admin/projects/{projectId}/features/{featureName}/environments/{environment}/strategies/set-sort-order | Set the order of strategies on the list.
[**StaleFeatures**](FeaturesApi.md#StaleFeatures) | **Post** /api/admin/projects/{projectId}/stale | Mark features as stale / not stale
[**ToggleFeatureEnvironmentOff**](FeaturesApi.md#ToggleFeatureEnvironmentOff) | **Post** /api/admin/projects/{projectId}/features/{featureName}/environments/{environment}/off | Disable a feature toggle.
[**ToggleFeatureEnvironmentOn**](FeaturesApi.md#ToggleFeatureEnvironmentOn) | **Post** /api/admin/projects/{projectId}/features/{featureName}/environments/{environment}/on | Enable a feature toggle.
[**UpdateFeature**](FeaturesApi.md#UpdateFeature) | **Put** /api/admin/projects/{projectId}/features/{featureName} | 
[**UpdateFeatureStrategy**](FeaturesApi.md#UpdateFeatureStrategy) | **Put** /api/admin/projects/{projectId}/features/{featureName}/environments/{environment}/strategies/{strategyId} | Update a strategy.
[**UpdateTags**](FeaturesApi.md#UpdateTags) | **Put** /api/admin/features/{featureName}/tags | Updates multiple tags for a feature.
[**ValidateConstraint**](FeaturesApi.md#ValidateConstraint) | **Post** /api/admin/constraints/validate | Validate constraint
[**ValidateFeature**](FeaturesApi.md#ValidateFeature) | **Post** /api/admin/features/validate | Validate feature name

# **AddFavoriteFeature**
> AddFavoriteFeature(ctx, projectId, featureName)
Add feature to favorites

This endpoint marks the feature in the url as favorite

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**|  | 
  **featureName** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddFavoriteProject**
> AddFavoriteProject(ctx, projectId)
Add project to favorites

This endpoint marks the project in the url as favorite

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
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddFeatureStrategy**
> FeatureStrategySchema AddFeatureStrategy(ctx, body, projectId, featureName, environment)
Add a strategy to a feature toggle.

Add a strategy to a feature toggle in the specified environment.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateFeatureStrategySchema**](CreateFeatureStrategySchema.md)| createFeatureStrategySchema | 
  **projectId** | **string**|  | 
  **featureName** | **string**|  | 
  **environment** | **string**|  | 

### Return type

[**FeatureStrategySchema**](featureStrategySchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddTag**
> TagSchema AddTag(ctx, body, featureName)
Adds a tag to a feature.

Adds a tag to a feature if the feature and tag type exist in the system. The operation is idempotent, so adding an existing tag will result in a successful response.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TagSchema**](TagSchema.md)| tagSchema | 
  **featureName** | **string**|  | 

### Return type

[**TagSchema**](tagSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ArchiveFeature**
> ArchiveFeature(ctx, projectId, featureName)
Archive a feature.

This endpoint archives the specified feature if the feature belongs to the specified project.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**|  | 
  **featureName** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ArchiveFeatures**
> ArchiveFeatures(ctx, body, projectId)
Archives a list of features

This endpoint archives the specified features. Any features that are already archived or that don't exist are ignored. All existing features (whether already archived or not) that are provided must belong to the specified project.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BatchFeaturesSchema**](BatchFeaturesSchema.md)| batchFeaturesSchema | 
  **projectId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BulkToggleFeaturesEnvironmentOff**
> BulkToggleFeaturesEnvironmentOff(ctx, body, projectId, environment)
Bulk disabled a list of features.

This endpoint disables multiple feature toggles.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BulkToggleFeaturesSchema**](BulkToggleFeaturesSchema.md)| bulkToggleFeaturesSchema | 
  **projectId** | **string**|  | 
  **environment** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BulkToggleFeaturesEnvironmentOn**
> BulkToggleFeaturesEnvironmentOn(ctx, body, projectId, environment)
Bulk enable a list of features.

This endpoint enables multiple feature toggles.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BulkToggleFeaturesSchema**](BulkToggleFeaturesSchema.md)| bulkToggleFeaturesSchema | 
  **projectId** | **string**|  | 
  **environment** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChangeProject**
> ChangeProject(ctx, body, projectId, featureName)
Move feature to project

Moves the specified feature to the new project in the request schema. Requires you to have permissions to move the feature toggle in both projects. Features that are included in any active change requests can not be moved.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ChangeProjectSchema**](ChangeProjectSchema.md)| changeProjectSchema | 
  **projectId** | **string**|  | 
  **featureName** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CloneFeature**
> FeatureSchema CloneFeature(ctx, body, projectId, featureName)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CloneFeatureSchema**](CloneFeatureSchema.md)| cloneFeatureSchema | 
  **projectId** | **string**|  | 
  **featureName** | **string**|  | 

### Return type

[**FeatureSchema**](featureSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateFeature**
> FeatureSchema CreateFeature(ctx, body, projectId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateFeatureSchema**](CreateFeatureSchema.md)| createFeatureSchema | 
  **projectId** | **string**|  | 

### Return type

[**FeatureSchema**](featureSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteFeatureStrategy**
> DeleteFeatureStrategy(ctx, projectId, featureName, environment, strategyId)
Delete a strategy from a feature toggle.

Delete a strategy configuration from a feature toggle in the specified environment.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**|  | 
  **featureName** | **string**|  | 
  **environment** | **string**|  | 
  **strategyId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllFeatureTypes**
> FeatureTypesSchema GetAllFeatureTypes(ctx, )
Get all feature types

Retrieves all feature types that exist in this Unleash instance, along with their descriptions and lifetimes.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**FeatureTypesSchema**](featureTypesSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllToggles**
> FeaturesSchema GetAllToggles(ctx, )
Get all features (deprecated)

Gets all feature toggles with their full configuration. This endpoint is **deprecated**. You should  use the project-based endpoint instead (`/api/admin/projects/<project-id>/features`).

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**FeaturesSchema**](featuresSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEnvironmentFeatureVariants**
> FeatureVariantsSchema GetEnvironmentFeatureVariants(ctx, projectId, featureName, environment)
Get variants for a feature in an environment

Returns the variants for a feature in a specific environment. If the feature has no variants it will return an empty array of variants

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**|  | 
  **featureName** | **string**|  | 
  **environment** | **string**|  | 

### Return type

[**FeatureVariantsSchema**](featureVariantsSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFeature**
> FeatureSchema GetFeature(ctx, projectId, featureName)
Get a feature.

This endpoint returns the information about the requested feature if the feature belongs to the specified project.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**|  | 
  **featureName** | **string**|  | 

### Return type

[**FeatureSchema**](featureSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFeatureEnvironment**
> FeatureEnvironmentSchema GetFeatureEnvironment(ctx, projectId, featureName, environment)
Get a feature environment.

Information about the enablement status and strategies for a feature toggle in specified environment.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**|  | 
  **featureName** | **string**|  | 
  **environment** | **string**|  | 

### Return type

[**FeatureEnvironmentSchema**](featureEnvironmentSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFeatureStrategies**
> FeatureStrategySchema GetFeatureStrategies(ctx, projectId, featureName, environment)
Get feature toggle strategies.

Get strategies defined for a feature toggle in the specified environment.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**|  | 
  **featureName** | **string**|  | 
  **environment** | **string**|  | 

### Return type

[**FeatureStrategySchema**](featureStrategySchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFeatureStrategy**
> FeatureStrategySchema GetFeatureStrategy(ctx, projectId, featureName, environment, strategyId)
Get a strategy configuration.

Get a strategy configuration for an environment in a feature toggle.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**|  | 
  **featureName** | **string**|  | 
  **environment** | **string**|  | 
  **strategyId** | **string**|  | 

### Return type

[**FeatureStrategySchema**](featureStrategySchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFeatureVariants**
> FeatureVariantsSchema GetFeatureVariants(ctx, projectId, featureName)
Retrieve variants for a feature (deprecated) 

(deprecated from 4.21) Retrieve the variants for the specified feature. From Unleash 4.21 onwards, this endpoint will attempt to choose a [production-type environment](https://docs.getunleash.io/reference/environments) as the source of truth. If more than one production environment is found, the first one will be used.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**|  | 
  **featureName** | **string**|  | 

### Return type

[**FeatureVariantsSchema**](featureVariantsSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFeatures**
> FeaturesSchema GetFeatures(ctx, projectId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**|  | 

### Return type

[**FeaturesSchema**](featuresSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListTags**
> TagsSchema ListTags(ctx, featureName)
Get all tags for a feature.

Retrieves all the tags for a feature name. If the feature does not exist it returns an empty list.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **featureName** | **string**|  | 

### Return type

[**TagsSchema**](tagsSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OverwriteEnvironmentFeatureVariants**
> FeatureVariantsSchema OverwriteEnvironmentFeatureVariants(ctx, body, projectId, featureName, environment)
Create (overwrite) variants for a feature in an environment

This overwrites the current variants for the feature toggle in the :featureName parameter for the :environment parameter.                                                  The backend will validate the input for the following invariants:                                              * If there are variants, there needs to be at least one variant with `weightType: variable`                     * The sum of the weights of variants with `weightType: fix` must be strictly less than 1000 (< 1000)                      The backend will also distribute remaining weight up to 1000 after adding the variants with `weightType: fix` together amongst the variants of `weightType: variable`

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]VariantSchema**](variantSchema.md)| variantsSchema | 
  **projectId** | **string**|  | 
  **featureName** | **string**|  | 
  **environment** | **string**|  | 

### Return type

[**FeatureVariantsSchema**](featureVariantsSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OverwriteFeatureVariants**
> FeatureVariantsSchema OverwriteFeatureVariants(ctx, body, projectId, featureName)
Create (overwrite) variants for a feature toggle in all environments

This overwrites the current variants for the feature specified in the :featureName parameter in all environments.                      The backend will validate the input for the following invariants                      * If there are variants, there needs to be at least one variant with `weightType: variable`                     * The sum of the weights of variants with `weightType: fix` must be strictly less than 1000 (< 1000)                      The backend will also distribute remaining weight up to 1000 after adding the variants with `weightType: fix` together amongst the variants of `weightType: variable`

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]VariantSchema**](variantSchema.md)| variantsSchema | 
  **projectId** | **string**|  | 
  **featureName** | **string**|  | 

### Return type

[**FeatureVariantsSchema**](featureVariantsSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OverwriteFeatureVariantsOnEnvironments**
> FeatureVariantsSchema OverwriteFeatureVariantsOnEnvironments(ctx, body, projectId, featureName)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PushVariantsSchema**](PushVariantsSchema.md)| pushVariantsSchema | 
  **projectId** | **string**|  | 
  **featureName** | **string**|  | 

### Return type

[**FeatureVariantsSchema**](featureVariantsSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchEnvironmentsFeatureVariants**
> FeatureVariantsSchema PatchEnvironmentsFeatureVariants(ctx, body, projectId, featureName, environment)
Patch a feature's variants in an environment

Apply a list of patches to the features environments in the specified environment. The patch objects should conform to the [JSON-patch format (RFC 6902)](https://www.rfc-editor.org/rfc/rfc6902).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]PatchSchema**](patchSchema.md)| patchesSchema | 
  **projectId** | **string**|  | 
  **featureName** | **string**|  | 
  **environment** | **string**|  | 

### Return type

[**FeatureVariantsSchema**](featureVariantsSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchFeature**
> FeatureSchema PatchFeature(ctx, body, projectId, featureName)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]PatchSchema**](patchSchema.md)| patchesSchema | 
  **projectId** | **string**|  | 
  **featureName** | **string**|  | 

### Return type

[**FeatureSchema**](featureSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchFeatureStrategy**
> FeatureStrategySchema PatchFeatureStrategy(ctx, body, projectId, featureName, environment, strategyId)
Change specific properties of a strategy.

Change specific properties of a strategy configuration in a feature toggle.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]PatchSchema**](patchSchema.md)| patchesSchema | 
  **projectId** | **string**|  | 
  **featureName** | **string**|  | 
  **environment** | **string**|  | 
  **strategyId** | **string**|  | 

### Return type

[**FeatureStrategySchema**](featureStrategySchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchFeatureVariants**
> FeatureVariantsSchema PatchFeatureVariants(ctx, body, projectId, featureName)
Apply a patch to a feature's variants (in all environments).

Apply a list of patches patch to the specified feature's variants. The patch objects should conform to the [JSON-patch format (RFC 6902)](https://www.rfc-editor.org/rfc/rfc6902).                                                  ⚠️ **Warning**: This method is not atomic. If something fails in the middle of applying the patch, you can be left with a half-applied patch. We recommend that you instead [patch variants on a per-environment basis](/docs/reference/api/unleash/patch-environments-feature-variants.api.mdx), which **is** an atomic operation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]PatchSchema**](patchSchema.md)| patchesSchema | 
  **projectId** | **string**|  | 
  **featureName** | **string**|  | 

### Return type

[**FeatureVariantsSchema**](featureVariantsSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveFavoriteFeature**
> RemoveFavoriteFeature(ctx, projectId, featureName)
Remove feature from favorites

This endpoint removes the feature in the url from favorites

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**|  | 
  **featureName** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveFavoriteProject**
> RemoveFavoriteProject(ctx, projectId)
Remove project from favorites

This endpoint removes the project in the url from favorites

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
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveTag**
> RemoveTag(ctx, featureName, type_, value)
Removes a tag from a feature.

Removes a tag from a feature. If the feature exists but the tag does not, it returns a successful response.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **featureName** | **string**|  | 
  **type_** | **string**|  | 
  **value** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetStrategySortOrder**
> SetStrategySortOrder(ctx, body, projectId, featureName, environment)
Set the order of strategies on the list.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]SetStrategySortOrderSchemaInner**](setStrategySortOrderSchema_inner.md)| setStrategySortOrderSchema | 
  **projectId** | **string**|  | 
  **featureName** | **string**|  | 
  **environment** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StaleFeatures**
> StaleFeatures(ctx, body, projectId)
Mark features as stale / not stale

This endpoint marks the provided list of features as either [stale](https://docs.getunleash.io/reference/technical-debt#stale-and-potentially-stale-toggles) or not stale depending on the request body you send. Any provided features that don't exist are ignored.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BatchStaleSchema**](BatchStaleSchema.md)| batchStaleSchema | 
  **projectId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ToggleFeatureEnvironmentOff**
> FeatureSchema ToggleFeatureEnvironmentOff(ctx, projectId, featureName, environment)
Disable a feature toggle.

Disable a feature toggle in the specified environment.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**|  | 
  **featureName** | **string**|  | 
  **environment** | **string**|  | 

### Return type

[**FeatureSchema**](featureSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ToggleFeatureEnvironmentOn**
> FeatureSchema ToggleFeatureEnvironmentOn(ctx, projectId, featureName, environment)
Enable a feature toggle.

Enable a feature toggle in the specified environment.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**|  | 
  **featureName** | **string**|  | 
  **environment** | **string**|  | 

### Return type

[**FeatureSchema**](featureSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateFeature**
> FeatureSchema UpdateFeature(ctx, body, projectId, featureName)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateFeatureSchema**](UpdateFeatureSchema.md)| updateFeatureSchema | 
  **projectId** | **string**|  | 
  **featureName** | **string**|  | 

### Return type

[**FeatureSchema**](featureSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateFeatureStrategy**
> FeatureStrategySchema UpdateFeatureStrategy(ctx, body, projectId, featureName, environment, strategyId)
Update a strategy.

Replace strategy configuration for a feature toggle in the specified environment.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateFeatureStrategySchema**](UpdateFeatureStrategySchema.md)| updateFeatureStrategySchema | 
  **projectId** | **string**|  | 
  **featureName** | **string**|  | 
  **environment** | **string**|  | 
  **strategyId** | **string**|  | 

### Return type

[**FeatureStrategySchema**](featureStrategySchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateTags**
> TagsSchema UpdateTags(ctx, body, featureName)
Updates multiple tags for a feature.

Receives a list of tags to add and a list of tags to remove that are mandatory but can be empty. All tags under addedTags are first added to the feature and then all tags under removedTags are removed from the feature.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateTagsSchema**](UpdateTagsSchema.md)| updateTagsSchema | 
  **featureName** | **string**|  | 

### Return type

[**TagsSchema**](tagsSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ValidateConstraint**
> ValidateConstraint(ctx, body)
Validate constraint

Validates a constraint definition. Checks whether the context field exists and whether the applied configuration is valid. Additional properties are not allowed on data objects that you send to this endpoint.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ConstraintSchema**](ConstraintSchema.md)| constraintSchema | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ValidateFeature**
> ValidateFeature(ctx, body)
Validate feature name

Validates a feature toggle name: checks whether the name is URL-friendly and whether a feature with the given name already exists. Returns 200 if the feature name is compliant and unused.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ValidateFeatureSchema**](ValidateFeatureSchema.md)| validateFeatureSchema | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

