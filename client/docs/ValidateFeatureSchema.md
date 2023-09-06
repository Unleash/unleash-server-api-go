# ValidateFeatureSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The feature name to validate. | 
**ProjectId** | Pointer to **NullableString** | The id of the project that the feature flag will belong to. If the target project has a feature naming pattern defined, the name will be validated against that pattern. | [optional] 

## Methods

### NewValidateFeatureSchema

`func NewValidateFeatureSchema(name string, ) *ValidateFeatureSchema`

NewValidateFeatureSchema instantiates a new ValidateFeatureSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidateFeatureSchemaWithDefaults

`func NewValidateFeatureSchemaWithDefaults() *ValidateFeatureSchema`

NewValidateFeatureSchemaWithDefaults instantiates a new ValidateFeatureSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ValidateFeatureSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ValidateFeatureSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ValidateFeatureSchema) SetName(v string)`

SetName sets Name field to given value.


### GetProjectId

`func (o *ValidateFeatureSchema) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ValidateFeatureSchema) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ValidateFeatureSchema) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *ValidateFeatureSchema) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *ValidateFeatureSchema) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *ValidateFeatureSchema) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


