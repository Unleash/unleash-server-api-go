# EnvironmentProjectSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the environment | 
**Type** | **string** | The [type of environment](https://docs.getunleash.io/reference/environments#environment-types). | 
**Enabled** | **bool** | &#x60;true&#x60; if the environment is enabled for the project, otherwise &#x60;false&#x60; | 
**Protected** | **bool** | &#x60;true&#x60; if the environment is protected, otherwise &#x60;false&#x60;. A *protected* environment can not be deleted. | 
**SortOrder** | **int32** | Priority of the environment in a list of environments, the lower the value, the higher up in the list the environment will appear | 
**ProjectApiTokenCount** | Pointer to **int32** | The number of client and front-end API tokens that have access to this project | [optional] 
**ProjectEnabledToggleCount** | Pointer to **int32** | The number of features enabled in this environment for this project | [optional] 
**DefaultStrategy** | Pointer to [**CreateFeatureStrategySchema**](CreateFeatureStrategySchema.md) |  | [optional] 
**Visible** | Pointer to **bool** | Indicates whether the environment can be enabled for feature flags in the project | [optional] 

## Methods

### NewEnvironmentProjectSchema

`func NewEnvironmentProjectSchema(name string, type_ string, enabled bool, protected bool, sortOrder int32, ) *EnvironmentProjectSchema`

NewEnvironmentProjectSchema instantiates a new EnvironmentProjectSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentProjectSchemaWithDefaults

`func NewEnvironmentProjectSchemaWithDefaults() *EnvironmentProjectSchema`

NewEnvironmentProjectSchemaWithDefaults instantiates a new EnvironmentProjectSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EnvironmentProjectSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EnvironmentProjectSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EnvironmentProjectSchema) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *EnvironmentProjectSchema) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EnvironmentProjectSchema) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EnvironmentProjectSchema) SetType(v string)`

SetType sets Type field to given value.


### GetEnabled

`func (o *EnvironmentProjectSchema) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *EnvironmentProjectSchema) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *EnvironmentProjectSchema) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetProtected

`func (o *EnvironmentProjectSchema) GetProtected() bool`

GetProtected returns the Protected field if non-nil, zero value otherwise.

### GetProtectedOk

`func (o *EnvironmentProjectSchema) GetProtectedOk() (*bool, bool)`

GetProtectedOk returns a tuple with the Protected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtected

`func (o *EnvironmentProjectSchema) SetProtected(v bool)`

SetProtected sets Protected field to given value.


### GetSortOrder

`func (o *EnvironmentProjectSchema) GetSortOrder() int32`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *EnvironmentProjectSchema) GetSortOrderOk() (*int32, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *EnvironmentProjectSchema) SetSortOrder(v int32)`

SetSortOrder sets SortOrder field to given value.


### GetProjectApiTokenCount

`func (o *EnvironmentProjectSchema) GetProjectApiTokenCount() int32`

GetProjectApiTokenCount returns the ProjectApiTokenCount field if non-nil, zero value otherwise.

### GetProjectApiTokenCountOk

`func (o *EnvironmentProjectSchema) GetProjectApiTokenCountOk() (*int32, bool)`

GetProjectApiTokenCountOk returns a tuple with the ProjectApiTokenCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectApiTokenCount

`func (o *EnvironmentProjectSchema) SetProjectApiTokenCount(v int32)`

SetProjectApiTokenCount sets ProjectApiTokenCount field to given value.

### HasProjectApiTokenCount

`func (o *EnvironmentProjectSchema) HasProjectApiTokenCount() bool`

HasProjectApiTokenCount returns a boolean if a field has been set.

### GetProjectEnabledToggleCount

`func (o *EnvironmentProjectSchema) GetProjectEnabledToggleCount() int32`

GetProjectEnabledToggleCount returns the ProjectEnabledToggleCount field if non-nil, zero value otherwise.

### GetProjectEnabledToggleCountOk

`func (o *EnvironmentProjectSchema) GetProjectEnabledToggleCountOk() (*int32, bool)`

GetProjectEnabledToggleCountOk returns a tuple with the ProjectEnabledToggleCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectEnabledToggleCount

`func (o *EnvironmentProjectSchema) SetProjectEnabledToggleCount(v int32)`

SetProjectEnabledToggleCount sets ProjectEnabledToggleCount field to given value.

### HasProjectEnabledToggleCount

`func (o *EnvironmentProjectSchema) HasProjectEnabledToggleCount() bool`

HasProjectEnabledToggleCount returns a boolean if a field has been set.

### GetDefaultStrategy

`func (o *EnvironmentProjectSchema) GetDefaultStrategy() CreateFeatureStrategySchema`

GetDefaultStrategy returns the DefaultStrategy field if non-nil, zero value otherwise.

### GetDefaultStrategyOk

`func (o *EnvironmentProjectSchema) GetDefaultStrategyOk() (*CreateFeatureStrategySchema, bool)`

GetDefaultStrategyOk returns a tuple with the DefaultStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultStrategy

`func (o *EnvironmentProjectSchema) SetDefaultStrategy(v CreateFeatureStrategySchema)`

SetDefaultStrategy sets DefaultStrategy field to given value.

### HasDefaultStrategy

`func (o *EnvironmentProjectSchema) HasDefaultStrategy() bool`

HasDefaultStrategy returns a boolean if a field has been set.

### GetVisible

`func (o *EnvironmentProjectSchema) GetVisible() bool`

GetVisible returns the Visible field if non-nil, zero value otherwise.

### GetVisibleOk

`func (o *EnvironmentProjectSchema) GetVisibleOk() (*bool, bool)`

GetVisibleOk returns a tuple with the Visible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisible

`func (o *EnvironmentProjectSchema) SetVisible(v bool)`

SetVisible sets Visible field to given value.

### HasVisible

`func (o *EnvironmentProjectSchema) HasVisible() bool`

HasVisible returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


