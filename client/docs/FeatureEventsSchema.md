# FeatureEventsSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to **float32** |  | [optional] 
**ToggleName** | Pointer to **string** |  | [optional] 
**Events** | [**[]EventSchema**](EventSchema.md) |  | 
**TotalEvents** | Pointer to **int32** |  | [optional] 

## Methods

### NewFeatureEventsSchema

`func NewFeatureEventsSchema(events []EventSchema, ) *FeatureEventsSchema`

NewFeatureEventsSchema instantiates a new FeatureEventsSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeatureEventsSchemaWithDefaults

`func NewFeatureEventsSchemaWithDefaults() *FeatureEventsSchema`

NewFeatureEventsSchemaWithDefaults instantiates a new FeatureEventsSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *FeatureEventsSchema) GetVersion() float32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *FeatureEventsSchema) GetVersionOk() (*float32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *FeatureEventsSchema) SetVersion(v float32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *FeatureEventsSchema) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetToggleName

`func (o *FeatureEventsSchema) GetToggleName() string`

GetToggleName returns the ToggleName field if non-nil, zero value otherwise.

### GetToggleNameOk

`func (o *FeatureEventsSchema) GetToggleNameOk() (*string, bool)`

GetToggleNameOk returns a tuple with the ToggleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToggleName

`func (o *FeatureEventsSchema) SetToggleName(v string)`

SetToggleName sets ToggleName field to given value.

### HasToggleName

`func (o *FeatureEventsSchema) HasToggleName() bool`

HasToggleName returns a boolean if a field has been set.

### GetEvents

`func (o *FeatureEventsSchema) GetEvents() []EventSchema`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *FeatureEventsSchema) GetEventsOk() (*[]EventSchema, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *FeatureEventsSchema) SetEvents(v []EventSchema)`

SetEvents sets Events field to given value.


### GetTotalEvents

`func (o *FeatureEventsSchema) GetTotalEvents() int32`

GetTotalEvents returns the TotalEvents field if non-nil, zero value otherwise.

### GetTotalEventsOk

`func (o *FeatureEventsSchema) GetTotalEventsOk() (*int32, bool)`

GetTotalEventsOk returns a tuple with the TotalEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalEvents

`func (o *FeatureEventsSchema) SetTotalEvents(v int32)`

SetTotalEvents sets TotalEvents field to given value.

### HasTotalEvents

`func (o *FeatureEventsSchema) HasTotalEvents() bool`

HasTotalEvents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


