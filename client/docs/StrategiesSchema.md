# StrategiesSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | **int32** |  | 
**Strategies** | [**[]StrategySchema**](StrategySchema.md) |  | 

## Methods

### NewStrategiesSchema

`func NewStrategiesSchema(version int32, strategies []StrategySchema, ) *StrategiesSchema`

NewStrategiesSchema instantiates a new StrategiesSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStrategiesSchemaWithDefaults

`func NewStrategiesSchemaWithDefaults() *StrategiesSchema`

NewStrategiesSchemaWithDefaults instantiates a new StrategiesSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *StrategiesSchema) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *StrategiesSchema) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *StrategiesSchema) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetStrategies

`func (o *StrategiesSchema) GetStrategies() []StrategySchema`

GetStrategies returns the Strategies field if non-nil, zero value otherwise.

### GetStrategiesOk

`func (o *StrategiesSchema) GetStrategiesOk() (*[]StrategySchema, bool)`

GetStrategiesOk returns a tuple with the Strategies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategies

`func (o *StrategiesSchema) SetStrategies(v []StrategySchema)`

SetStrategies sets Strategies field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


