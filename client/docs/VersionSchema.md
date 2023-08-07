# VersionSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Current** | [**VersionSchemaCurrent**](VersionSchemaCurrent.md) |  | 
**Latest** | [**VersionSchemaLatest**](VersionSchemaLatest.md) |  | 
**IsLatest** | **bool** | Whether the Unleash server is running the latest release (&#x60;true&#x60;) or if there are updates available (&#x60;false&#x60;) | 
**InstanceId** | Pointer to **string** | The instance identifier of the Unleash instance | [optional] 

## Methods

### NewVersionSchema

`func NewVersionSchema(current VersionSchemaCurrent, latest VersionSchemaLatest, isLatest bool, ) *VersionSchema`

NewVersionSchema instantiates a new VersionSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVersionSchemaWithDefaults

`func NewVersionSchemaWithDefaults() *VersionSchema`

NewVersionSchemaWithDefaults instantiates a new VersionSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrent

`func (o *VersionSchema) GetCurrent() VersionSchemaCurrent`

GetCurrent returns the Current field if non-nil, zero value otherwise.

### GetCurrentOk

`func (o *VersionSchema) GetCurrentOk() (*VersionSchemaCurrent, bool)`

GetCurrentOk returns a tuple with the Current field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrent

`func (o *VersionSchema) SetCurrent(v VersionSchemaCurrent)`

SetCurrent sets Current field to given value.


### GetLatest

`func (o *VersionSchema) GetLatest() VersionSchemaLatest`

GetLatest returns the Latest field if non-nil, zero value otherwise.

### GetLatestOk

`func (o *VersionSchema) GetLatestOk() (*VersionSchemaLatest, bool)`

GetLatestOk returns a tuple with the Latest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatest

`func (o *VersionSchema) SetLatest(v VersionSchemaLatest)`

SetLatest sets Latest field to given value.


### GetIsLatest

`func (o *VersionSchema) GetIsLatest() bool`

GetIsLatest returns the IsLatest field if non-nil, zero value otherwise.

### GetIsLatestOk

`func (o *VersionSchema) GetIsLatestOk() (*bool, bool)`

GetIsLatestOk returns a tuple with the IsLatest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLatest

`func (o *VersionSchema) SetIsLatest(v bool)`

SetIsLatest sets IsLatest field to given value.


### GetInstanceId

`func (o *VersionSchema) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *VersionSchema) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *VersionSchema) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *VersionSchema) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


