# OverrideSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContextName** | **string** | The name of the context field used to determine overrides | 
**Values** | **[]string** | Which values that should be overriden | 

## Methods

### NewOverrideSchema

`func NewOverrideSchema(contextName string, values []string, ) *OverrideSchema`

NewOverrideSchema instantiates a new OverrideSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOverrideSchemaWithDefaults

`func NewOverrideSchemaWithDefaults() *OverrideSchema`

NewOverrideSchemaWithDefaults instantiates a new OverrideSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContextName

`func (o *OverrideSchema) GetContextName() string`

GetContextName returns the ContextName field if non-nil, zero value otherwise.

### GetContextNameOk

`func (o *OverrideSchema) GetContextNameOk() (*string, bool)`

GetContextNameOk returns a tuple with the ContextName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextName

`func (o *OverrideSchema) SetContextName(v string)`

SetContextName sets ContextName field to given value.


### GetValues

`func (o *OverrideSchema) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *OverrideSchema) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *OverrideSchema) SetValues(v []string)`

SetValues sets Values field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


