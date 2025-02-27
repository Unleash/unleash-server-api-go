# EnvironmentSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the environment | 
**Type** | **string** | The [type of environment](https://docs.getunleash.io/reference/environments#environment-types). | 
**Enabled** | **bool** | &#x60;true&#x60; if the environment is enabled for the project, otherwise &#x60;false&#x60;. | 
**Protected** | **bool** | &#x60;true&#x60; if the environment is protected, otherwise &#x60;false&#x60;. A *protected* environment can not be deleted. | 
**SortOrder** | **int32** | Priority of the environment in a list of environments, the lower the value, the higher up in the list the environment will appear. Needs to be an integer | 
**ProjectCount** | Pointer to **NullableInt32** | The number of projects with this environment | [optional] 
**ApiTokenCount** | Pointer to **NullableInt32** | The number of API tokens for the project environment | [optional] 
**EnabledToggleCount** | Pointer to **NullableInt32** | The number of enabled toggles for the project environment | [optional] 

## Methods

### NewEnvironmentSchema

`func NewEnvironmentSchema(name string, type_ string, enabled bool, protected bool, sortOrder int32, ) *EnvironmentSchema`

NewEnvironmentSchema instantiates a new EnvironmentSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentSchemaWithDefaults

`func NewEnvironmentSchemaWithDefaults() *EnvironmentSchema`

NewEnvironmentSchemaWithDefaults instantiates a new EnvironmentSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EnvironmentSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EnvironmentSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EnvironmentSchema) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *EnvironmentSchema) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EnvironmentSchema) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EnvironmentSchema) SetType(v string)`

SetType sets Type field to given value.


### GetEnabled

`func (o *EnvironmentSchema) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *EnvironmentSchema) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *EnvironmentSchema) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetProtected

`func (o *EnvironmentSchema) GetProtected() bool`

GetProtected returns the Protected field if non-nil, zero value otherwise.

### GetProtectedOk

`func (o *EnvironmentSchema) GetProtectedOk() (*bool, bool)`

GetProtectedOk returns a tuple with the Protected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtected

`func (o *EnvironmentSchema) SetProtected(v bool)`

SetProtected sets Protected field to given value.


### GetSortOrder

`func (o *EnvironmentSchema) GetSortOrder() int32`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *EnvironmentSchema) GetSortOrderOk() (*int32, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *EnvironmentSchema) SetSortOrder(v int32)`

SetSortOrder sets SortOrder field to given value.


### GetProjectCount

`func (o *EnvironmentSchema) GetProjectCount() int32`

GetProjectCount returns the ProjectCount field if non-nil, zero value otherwise.

### GetProjectCountOk

`func (o *EnvironmentSchema) GetProjectCountOk() (*int32, bool)`

GetProjectCountOk returns a tuple with the ProjectCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectCount

`func (o *EnvironmentSchema) SetProjectCount(v int32)`

SetProjectCount sets ProjectCount field to given value.

### HasProjectCount

`func (o *EnvironmentSchema) HasProjectCount() bool`

HasProjectCount returns a boolean if a field has been set.

### SetProjectCountNil

`func (o *EnvironmentSchema) SetProjectCountNil(b bool)`

 SetProjectCountNil sets the value for ProjectCount to be an explicit nil

### UnsetProjectCount
`func (o *EnvironmentSchema) UnsetProjectCount()`

UnsetProjectCount ensures that no value is present for ProjectCount, not even an explicit nil
### GetApiTokenCount

`func (o *EnvironmentSchema) GetApiTokenCount() int32`

GetApiTokenCount returns the ApiTokenCount field if non-nil, zero value otherwise.

### GetApiTokenCountOk

`func (o *EnvironmentSchema) GetApiTokenCountOk() (*int32, bool)`

GetApiTokenCountOk returns a tuple with the ApiTokenCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiTokenCount

`func (o *EnvironmentSchema) SetApiTokenCount(v int32)`

SetApiTokenCount sets ApiTokenCount field to given value.

### HasApiTokenCount

`func (o *EnvironmentSchema) HasApiTokenCount() bool`

HasApiTokenCount returns a boolean if a field has been set.

### SetApiTokenCountNil

`func (o *EnvironmentSchema) SetApiTokenCountNil(b bool)`

 SetApiTokenCountNil sets the value for ApiTokenCount to be an explicit nil

### UnsetApiTokenCount
`func (o *EnvironmentSchema) UnsetApiTokenCount()`

UnsetApiTokenCount ensures that no value is present for ApiTokenCount, not even an explicit nil
### GetEnabledToggleCount

`func (o *EnvironmentSchema) GetEnabledToggleCount() int32`

GetEnabledToggleCount returns the EnabledToggleCount field if non-nil, zero value otherwise.

### GetEnabledToggleCountOk

`func (o *EnvironmentSchema) GetEnabledToggleCountOk() (*int32, bool)`

GetEnabledToggleCountOk returns a tuple with the EnabledToggleCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledToggleCount

`func (o *EnvironmentSchema) SetEnabledToggleCount(v int32)`

SetEnabledToggleCount sets EnabledToggleCount field to given value.

### HasEnabledToggleCount

`func (o *EnvironmentSchema) HasEnabledToggleCount() bool`

HasEnabledToggleCount returns a boolean if a field has been set.

### SetEnabledToggleCountNil

`func (o *EnvironmentSchema) SetEnabledToggleCountNil(b bool)`

 SetEnabledToggleCountNil sets the value for EnabledToggleCount to be an explicit nil

### UnsetEnabledToggleCount
`func (o *EnvironmentSchema) UnsetEnabledToggleCount()`

UnsetEnabledToggleCount ensures that no value is present for EnabledToggleCount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


