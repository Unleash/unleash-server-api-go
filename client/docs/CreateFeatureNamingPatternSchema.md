# CreateFeatureNamingPatternSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pattern** | **NullableString** | A JavaScript regular expression pattern, without the start and end delimiters. Optional flags are not allowed. | 
**Example** | Pointer to **NullableString** | An example of a feature name that matches the pattern. Must itself match the pattern supplied. | [optional] 

## Methods

### NewCreateFeatureNamingPatternSchema

`func NewCreateFeatureNamingPatternSchema(pattern NullableString, ) *CreateFeatureNamingPatternSchema`

NewCreateFeatureNamingPatternSchema instantiates a new CreateFeatureNamingPatternSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateFeatureNamingPatternSchemaWithDefaults

`func NewCreateFeatureNamingPatternSchemaWithDefaults() *CreateFeatureNamingPatternSchema`

NewCreateFeatureNamingPatternSchemaWithDefaults instantiates a new CreateFeatureNamingPatternSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPattern

`func (o *CreateFeatureNamingPatternSchema) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *CreateFeatureNamingPatternSchema) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *CreateFeatureNamingPatternSchema) SetPattern(v string)`

SetPattern sets Pattern field to given value.


### SetPatternNil

`func (o *CreateFeatureNamingPatternSchema) SetPatternNil(b bool)`

 SetPatternNil sets the value for Pattern to be an explicit nil

### UnsetPattern
`func (o *CreateFeatureNamingPatternSchema) UnsetPattern()`

UnsetPattern ensures that no value is present for Pattern, not even an explicit nil
### GetExample

`func (o *CreateFeatureNamingPatternSchema) GetExample() string`

GetExample returns the Example field if non-nil, zero value otherwise.

### GetExampleOk

`func (o *CreateFeatureNamingPatternSchema) GetExampleOk() (*string, bool)`

GetExampleOk returns a tuple with the Example field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExample

`func (o *CreateFeatureNamingPatternSchema) SetExample(v string)`

SetExample sets Example field to given value.

### HasExample

`func (o *CreateFeatureNamingPatternSchema) HasExample() bool`

HasExample returns a boolean if a field has been set.

### SetExampleNil

`func (o *CreateFeatureNamingPatternSchema) SetExampleNil(b bool)`

 SetExampleNil sets the value for Example to be an explicit nil

### UnsetExample
`func (o *CreateFeatureNamingPatternSchema) UnsetExample()`

UnsetExample ensures that no value is present for Example, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


