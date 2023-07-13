# AddonParameterSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the parameter as it is used in code. References to this parameter should use this value. | 
**DisplayName** | **string** | The name of the parameter as it is shown to the end user in the Admin UI. | 
**Type** | **string** | The type of the parameter. Corresponds roughly to [HTML &#x60;input&#x60; field types](https://developer.mozilla.org/docs/Web/HTML/Element/Input#input_types). Multi-line inut fields are indicated as &#x60;textfield&#x60; (equivalent to the HTML &#x60;textarea&#x60; tag). | 
**Description** | Pointer to **string** | A description of the parameter. This should explain to the end user what the parameter is used for. | [optional] 
**Placeholder** | Pointer to **string** | The default value for this parameter. This value is used if no other value is provided. | [optional] 
**Required** | **bool** | Whether this parameter is required or not. If a parameter is required, you must give it a value when you create the addon. If it is not required it can be left out. It may receive a default value in those cases. | 
**Sensitive** | **bool** | Indicates whether this parameter is **sensitive** or not. Unleash will not return sensitive parameters to API requests. It will instead use a number of asterisks to indicate that a value is set, e.g. \&quot;******\&quot;. The number of asterisks does not correlate to the parameter&#39;s value. | 

## Methods

### NewAddonParameterSchema

`func NewAddonParameterSchema(name string, displayName string, type_ string, required bool, sensitive bool, ) *AddonParameterSchema`

NewAddonParameterSchema instantiates a new AddonParameterSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddonParameterSchemaWithDefaults

`func NewAddonParameterSchemaWithDefaults() *AddonParameterSchema`

NewAddonParameterSchemaWithDefaults instantiates a new AddonParameterSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AddonParameterSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddonParameterSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddonParameterSchema) SetName(v string)`

SetName sets Name field to given value.


### GetDisplayName

`func (o *AddonParameterSchema) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AddonParameterSchema) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AddonParameterSchema) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetType

`func (o *AddonParameterSchema) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AddonParameterSchema) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AddonParameterSchema) SetType(v string)`

SetType sets Type field to given value.


### GetDescription

`func (o *AddonParameterSchema) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddonParameterSchema) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddonParameterSchema) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddonParameterSchema) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPlaceholder

`func (o *AddonParameterSchema) GetPlaceholder() string`

GetPlaceholder returns the Placeholder field if non-nil, zero value otherwise.

### GetPlaceholderOk

`func (o *AddonParameterSchema) GetPlaceholderOk() (*string, bool)`

GetPlaceholderOk returns a tuple with the Placeholder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceholder

`func (o *AddonParameterSchema) SetPlaceholder(v string)`

SetPlaceholder sets Placeholder field to given value.

### HasPlaceholder

`func (o *AddonParameterSchema) HasPlaceholder() bool`

HasPlaceholder returns a boolean if a field has been set.

### GetRequired

`func (o *AddonParameterSchema) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *AddonParameterSchema) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *AddonParameterSchema) SetRequired(v bool)`

SetRequired sets Required field to given value.


### GetSensitive

`func (o *AddonParameterSchema) GetSensitive() bool`

GetSensitive returns the Sensitive field if non-nil, zero value otherwise.

### GetSensitiveOk

`func (o *AddonParameterSchema) GetSensitiveOk() (*bool, bool)`

GetSensitiveOk returns a tuple with the Sensitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensitive

`func (o *AddonParameterSchema) SetSensitive(v bool)`

SetSensitive sets Sensitive field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


