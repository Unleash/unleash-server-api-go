# BatchStaleSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Features** | **[]string** | A list of features to mark as (not) stale | 
**Stale** | **bool** | Whether the list of features should be marked as stale or not stale. If &#x60;true&#x60;, the features will be marked as stale. If &#x60;false&#x60;, the features will be marked as not stale. | 

## Methods

### NewBatchStaleSchema

`func NewBatchStaleSchema(features []string, stale bool, ) *BatchStaleSchema`

NewBatchStaleSchema instantiates a new BatchStaleSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchStaleSchemaWithDefaults

`func NewBatchStaleSchemaWithDefaults() *BatchStaleSchema`

NewBatchStaleSchemaWithDefaults instantiates a new BatchStaleSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeatures

`func (o *BatchStaleSchema) GetFeatures() []string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *BatchStaleSchema) GetFeaturesOk() (*[]string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *BatchStaleSchema) SetFeatures(v []string)`

SetFeatures sets Features field to given value.


### GetStale

`func (o *BatchStaleSchema) GetStale() bool`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *BatchStaleSchema) GetStaleOk() (*bool, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *BatchStaleSchema) SetStale(v bool)`

SetStale sets Stale field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


