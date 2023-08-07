# TagTypeSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the tag type. | 
**Description** | Pointer to **string** | The description of the tag type. | [optional] 
**Icon** | Pointer to **NullableString** | The icon of the tag type. | [optional] 

## Methods

### NewTagTypeSchema

`func NewTagTypeSchema(name string, ) *TagTypeSchema`

NewTagTypeSchema instantiates a new TagTypeSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagTypeSchemaWithDefaults

`func NewTagTypeSchemaWithDefaults() *TagTypeSchema`

NewTagTypeSchemaWithDefaults instantiates a new TagTypeSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TagTypeSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TagTypeSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TagTypeSchema) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *TagTypeSchema) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TagTypeSchema) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TagTypeSchema) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TagTypeSchema) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIcon

`func (o *TagTypeSchema) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *TagTypeSchema) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *TagTypeSchema) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *TagTypeSchema) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### SetIconNil

`func (o *TagTypeSchema) SetIconNil(b bool)`

 SetIconNil sets the value for Icon to be an explicit nil

### UnsetIcon
`func (o *TagTypeSchema) UnsetIcon()`

UnsetIcon ensures that no value is present for Icon, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


