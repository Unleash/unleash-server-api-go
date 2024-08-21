# ProjectSchemaOwnersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OwnerType** | **string** |  | 
**Name** | **string** |  | 
**ImageUrl** | Pointer to **NullableString** |  | [optional] 
**Email** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewProjectSchemaOwnersInner

`func NewProjectSchemaOwnersInner(ownerType string, name string, ) *ProjectSchemaOwnersInner`

NewProjectSchemaOwnersInner instantiates a new ProjectSchemaOwnersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectSchemaOwnersInnerWithDefaults

`func NewProjectSchemaOwnersInnerWithDefaults() *ProjectSchemaOwnersInner`

NewProjectSchemaOwnersInnerWithDefaults instantiates a new ProjectSchemaOwnersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOwnerType

`func (o *ProjectSchemaOwnersInner) GetOwnerType() string`

GetOwnerType returns the OwnerType field if non-nil, zero value otherwise.

### GetOwnerTypeOk

`func (o *ProjectSchemaOwnersInner) GetOwnerTypeOk() (*string, bool)`

GetOwnerTypeOk returns a tuple with the OwnerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerType

`func (o *ProjectSchemaOwnersInner) SetOwnerType(v string)`

SetOwnerType sets OwnerType field to given value.


### GetName

`func (o *ProjectSchemaOwnersInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectSchemaOwnersInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectSchemaOwnersInner) SetName(v string)`

SetName sets Name field to given value.


### GetImageUrl

`func (o *ProjectSchemaOwnersInner) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *ProjectSchemaOwnersInner) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *ProjectSchemaOwnersInner) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *ProjectSchemaOwnersInner) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### SetImageUrlNil

`func (o *ProjectSchemaOwnersInner) SetImageUrlNil(b bool)`

 SetImageUrlNil sets the value for ImageUrl to be an explicit nil

### UnsetImageUrl
`func (o *ProjectSchemaOwnersInner) UnsetImageUrl()`

UnsetImageUrl ensures that no value is present for ImageUrl, not even an explicit nil
### GetEmail

`func (o *ProjectSchemaOwnersInner) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ProjectSchemaOwnersInner) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ProjectSchemaOwnersInner) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ProjectSchemaOwnersInner) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *ProjectSchemaOwnersInner) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *ProjectSchemaOwnersInner) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


