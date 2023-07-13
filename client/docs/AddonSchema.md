# AddonSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The addon&#39;s unique identifier. | 
**Provider** | **string** | The addon provider, such as \&quot;webhook\&quot; or \&quot;slack\&quot;. | 
**Description** | **NullableString** | A description of the addon. &#x60;null&#x60; if no description exists. | 
**Enabled** | **bool** | Whether the addon is enabled or not. | 
**Parameters** | **map[string]interface{}** | Parameters for the addon provider. This object has different required and optional properties depending on the provider you choose. | 
**Events** | **[]string** | The event types that trigger this specific addon. | 
**Projects** | Pointer to **[]string** | The projects that this addon listens to events from. An empty list means it listens to events from **all** projects. | [optional] 
**Environments** | Pointer to **[]string** | The list of environments that this addon listens to events from. An empty list means it listens to events from **all** environments. | [optional] 

## Methods

### NewAddonSchema

`func NewAddonSchema(id int32, provider string, description NullableString, enabled bool, parameters map[string]interface{}, events []string, ) *AddonSchema`

NewAddonSchema instantiates a new AddonSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddonSchemaWithDefaults

`func NewAddonSchemaWithDefaults() *AddonSchema`

NewAddonSchemaWithDefaults instantiates a new AddonSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddonSchema) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddonSchema) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddonSchema) SetId(v int32)`

SetId sets Id field to given value.


### GetProvider

`func (o *AddonSchema) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *AddonSchema) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *AddonSchema) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetDescription

`func (o *AddonSchema) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddonSchema) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddonSchema) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *AddonSchema) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AddonSchema) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEnabled

`func (o *AddonSchema) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddonSchema) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddonSchema) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetParameters

`func (o *AddonSchema) GetParameters() map[string]interface{}`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *AddonSchema) GetParametersOk() (*map[string]interface{}, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *AddonSchema) SetParameters(v map[string]interface{})`

SetParameters sets Parameters field to given value.


### GetEvents

`func (o *AddonSchema) GetEvents() []string`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *AddonSchema) GetEventsOk() (*[]string, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *AddonSchema) SetEvents(v []string)`

SetEvents sets Events field to given value.


### GetProjects

`func (o *AddonSchema) GetProjects() []string`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *AddonSchema) GetProjectsOk() (*[]string, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *AddonSchema) SetProjects(v []string)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *AddonSchema) HasProjects() bool`

HasProjects returns a boolean if a field has been set.

### GetEnvironments

`func (o *AddonSchema) GetEnvironments() []string`

GetEnvironments returns the Environments field if non-nil, zero value otherwise.

### GetEnvironmentsOk

`func (o *AddonSchema) GetEnvironmentsOk() (*[]string, bool)`

GetEnvironmentsOk returns a tuple with the Environments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironments

`func (o *AddonSchema) SetEnvironments(v []string)`

SetEnvironments sets Environments field to given value.

### HasEnvironments

`func (o *AddonSchema) HasEnvironments() bool`

HasEnvironments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


