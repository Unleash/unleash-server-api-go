# ClientMetricsSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppName** | **string** | The name of the application that is evaluating toggles | 
**InstanceId** | Pointer to **string** | A [(somewhat) unique identifier](https://docs.getunleash.io/reference/sdks/node#advanced-usage) for the application | [optional] 
**Environment** | Pointer to **string** | Which environment the application is running in | [optional] 
**Bucket** | [**ClientMetricsSchemaBucket**](ClientMetricsSchemaBucket.md) |  | 

## Methods

### NewClientMetricsSchema

`func NewClientMetricsSchema(appName string, bucket ClientMetricsSchemaBucket, ) *ClientMetricsSchema`

NewClientMetricsSchema instantiates a new ClientMetricsSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientMetricsSchemaWithDefaults

`func NewClientMetricsSchemaWithDefaults() *ClientMetricsSchema`

NewClientMetricsSchemaWithDefaults instantiates a new ClientMetricsSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppName

`func (o *ClientMetricsSchema) GetAppName() string`

GetAppName returns the AppName field if non-nil, zero value otherwise.

### GetAppNameOk

`func (o *ClientMetricsSchema) GetAppNameOk() (*string, bool)`

GetAppNameOk returns a tuple with the AppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppName

`func (o *ClientMetricsSchema) SetAppName(v string)`

SetAppName sets AppName field to given value.


### GetInstanceId

`func (o *ClientMetricsSchema) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *ClientMetricsSchema) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *ClientMetricsSchema) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *ClientMetricsSchema) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### GetEnvironment

`func (o *ClientMetricsSchema) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *ClientMetricsSchema) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *ClientMetricsSchema) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *ClientMetricsSchema) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetBucket

`func (o *ClientMetricsSchema) GetBucket() ClientMetricsSchemaBucket`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *ClientMetricsSchema) GetBucketOk() (*ClientMetricsSchemaBucket, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *ClientMetricsSchema) SetBucket(v ClientMetricsSchemaBucket)`

SetBucket sets Bucket field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


