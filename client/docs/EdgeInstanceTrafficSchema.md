# EdgeInstanceTrafficSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Get** | [**map[string]EdgeLatencyMetricsSchema**](EdgeLatencyMetricsSchema.md) | A map containing GET requests. | 
**Post** | [**map[string]EdgeLatencyMetricsSchema**](EdgeLatencyMetricsSchema.md) | A map containing POST requests. | 
**AccessDenied** | [**map[string]EdgeLatencyMetricsSchema**](EdgeLatencyMetricsSchema.md) | A map containing requests that were denied. | 
**CachedResponses** | [**map[string]EdgeLatencyMetricsSchema**](EdgeLatencyMetricsSchema.md) | A map containing requests that had cached responses. | 

## Methods

### NewEdgeInstanceTrafficSchema

`func NewEdgeInstanceTrafficSchema(get map[string]EdgeLatencyMetricsSchema, post map[string]EdgeLatencyMetricsSchema, accessDenied map[string]EdgeLatencyMetricsSchema, cachedResponses map[string]EdgeLatencyMetricsSchema, ) *EdgeInstanceTrafficSchema`

NewEdgeInstanceTrafficSchema instantiates a new EdgeInstanceTrafficSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEdgeInstanceTrafficSchemaWithDefaults

`func NewEdgeInstanceTrafficSchemaWithDefaults() *EdgeInstanceTrafficSchema`

NewEdgeInstanceTrafficSchemaWithDefaults instantiates a new EdgeInstanceTrafficSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGet

`func (o *EdgeInstanceTrafficSchema) GetGet() map[string]EdgeLatencyMetricsSchema`

GetGet returns the Get field if non-nil, zero value otherwise.

### GetGetOk

`func (o *EdgeInstanceTrafficSchema) GetGetOk() (*map[string]EdgeLatencyMetricsSchema, bool)`

GetGetOk returns a tuple with the Get field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGet

`func (o *EdgeInstanceTrafficSchema) SetGet(v map[string]EdgeLatencyMetricsSchema)`

SetGet sets Get field to given value.


### GetPost

`func (o *EdgeInstanceTrafficSchema) GetPost() map[string]EdgeLatencyMetricsSchema`

GetPost returns the Post field if non-nil, zero value otherwise.

### GetPostOk

`func (o *EdgeInstanceTrafficSchema) GetPostOk() (*map[string]EdgeLatencyMetricsSchema, bool)`

GetPostOk returns a tuple with the Post field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPost

`func (o *EdgeInstanceTrafficSchema) SetPost(v map[string]EdgeLatencyMetricsSchema)`

SetPost sets Post field to given value.


### GetAccessDenied

`func (o *EdgeInstanceTrafficSchema) GetAccessDenied() map[string]EdgeLatencyMetricsSchema`

GetAccessDenied returns the AccessDenied field if non-nil, zero value otherwise.

### GetAccessDeniedOk

`func (o *EdgeInstanceTrafficSchema) GetAccessDeniedOk() (*map[string]EdgeLatencyMetricsSchema, bool)`

GetAccessDeniedOk returns a tuple with the AccessDenied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessDenied

`func (o *EdgeInstanceTrafficSchema) SetAccessDenied(v map[string]EdgeLatencyMetricsSchema)`

SetAccessDenied sets AccessDenied field to given value.


### GetCachedResponses

`func (o *EdgeInstanceTrafficSchema) GetCachedResponses() map[string]EdgeLatencyMetricsSchema`

GetCachedResponses returns the CachedResponses field if non-nil, zero value otherwise.

### GetCachedResponsesOk

`func (o *EdgeInstanceTrafficSchema) GetCachedResponsesOk() (*map[string]EdgeLatencyMetricsSchema, bool)`

GetCachedResponsesOk returns a tuple with the CachedResponses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCachedResponses

`func (o *EdgeInstanceTrafficSchema) SetCachedResponses(v map[string]EdgeLatencyMetricsSchema)`

SetCachedResponses sets CachedResponses field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


