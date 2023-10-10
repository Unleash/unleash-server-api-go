# UpdateProjectSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The new name of the project | 
**Description** | Pointer to **string** | A new description for the project | [optional] 
**Mode** | Pointer to **string** | A mode of the project affecting what actions are possible in this project | [optional] 
**DefaultStickiness** | Pointer to **string** | A default stickiness for the project affecting the default stickiness value for variants and Gradual Rollout strategy | [optional] 

## Methods

### NewUpdateProjectSchema

`func NewUpdateProjectSchema(name string, ) *UpdateProjectSchema`

NewUpdateProjectSchema instantiates a new UpdateProjectSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateProjectSchemaWithDefaults

`func NewUpdateProjectSchemaWithDefaults() *UpdateProjectSchema`

NewUpdateProjectSchemaWithDefaults instantiates a new UpdateProjectSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateProjectSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateProjectSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateProjectSchema) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *UpdateProjectSchema) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateProjectSchema) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateProjectSchema) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateProjectSchema) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMode

`func (o *UpdateProjectSchema) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *UpdateProjectSchema) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *UpdateProjectSchema) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *UpdateProjectSchema) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetDefaultStickiness

`func (o *UpdateProjectSchema) GetDefaultStickiness() string`

GetDefaultStickiness returns the DefaultStickiness field if non-nil, zero value otherwise.

### GetDefaultStickinessOk

`func (o *UpdateProjectSchema) GetDefaultStickinessOk() (*string, bool)`

GetDefaultStickinessOk returns a tuple with the DefaultStickiness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultStickiness

`func (o *UpdateProjectSchema) SetDefaultStickiness(v string)`

SetDefaultStickiness sets DefaultStickiness field to given value.

### HasDefaultStickiness

`func (o *UpdateProjectSchema) HasDefaultStickiness() bool`

HasDefaultStickiness returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


