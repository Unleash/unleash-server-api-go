# ClientMetricsSchemaBucket

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Start** | [**DateSchema**](DateSchema.md) |  | 
**Stop** | [**DateSchema**](DateSchema.md) |  | 
**Toggles** | [**map[string]ClientMetricsSchemaBucketTogglesValue**](ClientMetricsSchemaBucketTogglesValue.md) | an object containing feature names with yes/no plus variant usage | 

## Methods

### NewClientMetricsSchemaBucket

`func NewClientMetricsSchemaBucket(start DateSchema, stop DateSchema, toggles map[string]ClientMetricsSchemaBucketTogglesValue, ) *ClientMetricsSchemaBucket`

NewClientMetricsSchemaBucket instantiates a new ClientMetricsSchemaBucket object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientMetricsSchemaBucketWithDefaults

`func NewClientMetricsSchemaBucketWithDefaults() *ClientMetricsSchemaBucket`

NewClientMetricsSchemaBucketWithDefaults instantiates a new ClientMetricsSchemaBucket object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStart

`func (o *ClientMetricsSchemaBucket) GetStart() DateSchema`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *ClientMetricsSchemaBucket) GetStartOk() (*DateSchema, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *ClientMetricsSchemaBucket) SetStart(v DateSchema)`

SetStart sets Start field to given value.


### GetStop

`func (o *ClientMetricsSchemaBucket) GetStop() DateSchema`

GetStop returns the Stop field if non-nil, zero value otherwise.

### GetStopOk

`func (o *ClientMetricsSchemaBucket) GetStopOk() (*DateSchema, bool)`

GetStopOk returns a tuple with the Stop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStop

`func (o *ClientMetricsSchemaBucket) SetStop(v DateSchema)`

SetStop sets Stop field to given value.


### GetToggles

`func (o *ClientMetricsSchemaBucket) GetToggles() map[string]ClientMetricsSchemaBucketTogglesValue`

GetToggles returns the Toggles field if non-nil, zero value otherwise.

### GetTogglesOk

`func (o *ClientMetricsSchemaBucket) GetTogglesOk() (*map[string]ClientMetricsSchemaBucketTogglesValue, bool)`

GetTogglesOk returns a tuple with the Toggles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToggles

`func (o *ClientMetricsSchemaBucket) SetToggles(v map[string]ClientMetricsSchemaBucketTogglesValue)`

SetToggles sets Toggles field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


