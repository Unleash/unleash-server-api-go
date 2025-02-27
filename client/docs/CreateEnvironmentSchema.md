# CreateEnvironmentSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the environment. Must be a URL-friendly string according to [RFC 3968, section 2.3](https://www.rfc-editor.org/rfc/rfc3986#section-2.3) | 
**Type** | **string** | The [type of environment](https://docs.getunleash.io/reference/environments#environment-types) you would like to create. Unleash officially recognizes the following values: - &#x60;development&#x60; - &#x60;test&#x60; - &#x60;preproduction&#x60; - &#x60;production&#x60;  If you pass a string that is not one of the recognized values, Unleash will accept it, but it will carry no special semantics. | 
**Enabled** | Pointer to **bool** | Newly created environments are enabled by default. Set this property to &#x60;false&#x60; to create the environment in a disabled state. | [optional] 
**SortOrder** | Pointer to **int32** | Defines where in the list of environments to place this environment. The list uses an ascending sort, so lower numbers are shown first. You can change this value later. | [optional] 

## Methods

### NewCreateEnvironmentSchema

`func NewCreateEnvironmentSchema(name string, type_ string, ) *CreateEnvironmentSchema`

NewCreateEnvironmentSchema instantiates a new CreateEnvironmentSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateEnvironmentSchemaWithDefaults

`func NewCreateEnvironmentSchemaWithDefaults() *CreateEnvironmentSchema`

NewCreateEnvironmentSchemaWithDefaults instantiates a new CreateEnvironmentSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateEnvironmentSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateEnvironmentSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateEnvironmentSchema) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *CreateEnvironmentSchema) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateEnvironmentSchema) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateEnvironmentSchema) SetType(v string)`

SetType sets Type field to given value.


### GetEnabled

`func (o *CreateEnvironmentSchema) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CreateEnvironmentSchema) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CreateEnvironmentSchema) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CreateEnvironmentSchema) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetSortOrder

`func (o *CreateEnvironmentSchema) GetSortOrder() int32`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *CreateEnvironmentSchema) GetSortOrderOk() (*int32, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *CreateEnvironmentSchema) SetSortOrder(v int32)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *CreateEnvironmentSchema) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


