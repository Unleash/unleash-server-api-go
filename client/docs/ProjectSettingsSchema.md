# ProjectSettingsSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultStickiness** | **NullableString** | Default stickiness for project | 
**Mode** | **NullableString** |  | 

## Methods

### NewProjectSettingsSchema

`func NewProjectSettingsSchema(defaultStickiness NullableString, mode NullableString, ) *ProjectSettingsSchema`

NewProjectSettingsSchema instantiates a new ProjectSettingsSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectSettingsSchemaWithDefaults

`func NewProjectSettingsSchemaWithDefaults() *ProjectSettingsSchema`

NewProjectSettingsSchemaWithDefaults instantiates a new ProjectSettingsSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultStickiness

`func (o *ProjectSettingsSchema) GetDefaultStickiness() string`

GetDefaultStickiness returns the DefaultStickiness field if non-nil, zero value otherwise.

### GetDefaultStickinessOk

`func (o *ProjectSettingsSchema) GetDefaultStickinessOk() (*string, bool)`

GetDefaultStickinessOk returns a tuple with the DefaultStickiness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultStickiness

`func (o *ProjectSettingsSchema) SetDefaultStickiness(v string)`

SetDefaultStickiness sets DefaultStickiness field to given value.


### SetDefaultStickinessNil

`func (o *ProjectSettingsSchema) SetDefaultStickinessNil(b bool)`

 SetDefaultStickinessNil sets the value for DefaultStickiness to be an explicit nil

### UnsetDefaultStickiness
`func (o *ProjectSettingsSchema) UnsetDefaultStickiness()`

UnsetDefaultStickiness ensures that no value is present for DefaultStickiness, not even an explicit nil
### GetMode

`func (o *ProjectSettingsSchema) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *ProjectSettingsSchema) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *ProjectSettingsSchema) SetMode(v string)`

SetMode sets Mode field to given value.


### SetModeNil

`func (o *ProjectSettingsSchema) SetModeNil(b bool)`

 SetModeNil sets the value for Mode to be an explicit nil

### UnsetMode
`func (o *ProjectSettingsSchema) UnsetMode()`

UnsetMode ensures that no value is present for Mode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


