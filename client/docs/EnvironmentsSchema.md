# EnvironmentsSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | **int32** |  | 
**Environments** | [**[]EnvironmentSchema**](EnvironmentSchema.md) |  | 

## Methods

### NewEnvironmentsSchema

`func NewEnvironmentsSchema(version int32, environments []EnvironmentSchema, ) *EnvironmentsSchema`

NewEnvironmentsSchema instantiates a new EnvironmentsSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentsSchemaWithDefaults

`func NewEnvironmentsSchemaWithDefaults() *EnvironmentsSchema`

NewEnvironmentsSchemaWithDefaults instantiates a new EnvironmentsSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *EnvironmentsSchema) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *EnvironmentsSchema) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *EnvironmentsSchema) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetEnvironments

`func (o *EnvironmentsSchema) GetEnvironments() []EnvironmentSchema`

GetEnvironments returns the Environments field if non-nil, zero value otherwise.

### GetEnvironmentsOk

`func (o *EnvironmentsSchema) GetEnvironmentsOk() (*[]EnvironmentSchema, bool)`

GetEnvironmentsOk returns a tuple with the Environments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironments

`func (o *EnvironmentsSchema) SetEnvironments(v []EnvironmentSchema)`

SetEnvironments sets Environments field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


