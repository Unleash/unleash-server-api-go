# AddonCreateUpdateSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Provider** | **string** | The addon provider, such as \&quot;webhook\&quot; or \&quot;slack\&quot;. This string is **case sensitive** and maps to the provider&#39;s &#x60;name&#x60; property.  The list of all supported providers and their parameters for a specific Unleash instance can be found by making a GET request to the &#x60;api/admin/addons&#x60; endpoint: the &#x60;providers&#x60; property of that response will contain all available providers.  The default set of providers can be found in the [addons reference documentation](https://docs.getunleash.io/reference/addons). The default supported options are: - &#x60;datadog&#x60; for [Datadog](https://docs.getunleash.io/reference/addons/datadog) - &#x60;slack&#x60; for [Slack](https://docs.getunleash.io/reference/addons/slack) - &#x60;teams&#x60; for [Microsoft Teams](https://docs.getunleash.io/reference/addons/teams) - &#x60;webhook&#x60; for [webhooks](https://docs.getunleash.io/reference/addons/webhook)  The provider you choose for your addon dictates what properties the &#x60;parameters&#x60; object needs. Refer to the documentation for each provider for more information.  | 
**Description** | Pointer to **string** | A description of the addon. | [optional] 
**Enabled** | **bool** | Whether the addon should be enabled or not. | 
**Parameters** | **map[string]interface{}** | Parameters for the addon provider. This object has different required and optional properties depending on the provider you choose. Consult the documentation for details. | 
**Events** | **[]string** | The event types that will trigger this specific addon. | 
**Projects** | Pointer to **[]string** | The projects that this addon will listen to events from. An empty list means it will listen to events from **all** projects. | [optional] 
**Environments** | Pointer to **[]string** | The list of environments that this addon will listen to events from. An empty list means it will listen to events from **all** environments. | [optional] 

## Methods

### NewAddonCreateUpdateSchema

`func NewAddonCreateUpdateSchema(provider string, enabled bool, parameters map[string]interface{}, events []string, ) *AddonCreateUpdateSchema`

NewAddonCreateUpdateSchema instantiates a new AddonCreateUpdateSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddonCreateUpdateSchemaWithDefaults

`func NewAddonCreateUpdateSchemaWithDefaults() *AddonCreateUpdateSchema`

NewAddonCreateUpdateSchemaWithDefaults instantiates a new AddonCreateUpdateSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProvider

`func (o *AddonCreateUpdateSchema) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *AddonCreateUpdateSchema) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *AddonCreateUpdateSchema) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetDescription

`func (o *AddonCreateUpdateSchema) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddonCreateUpdateSchema) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddonCreateUpdateSchema) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddonCreateUpdateSchema) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddonCreateUpdateSchema) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddonCreateUpdateSchema) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddonCreateUpdateSchema) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetParameters

`func (o *AddonCreateUpdateSchema) GetParameters() map[string]interface{}`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *AddonCreateUpdateSchema) GetParametersOk() (*map[string]interface{}, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *AddonCreateUpdateSchema) SetParameters(v map[string]interface{})`

SetParameters sets Parameters field to given value.


### GetEvents

`func (o *AddonCreateUpdateSchema) GetEvents() []string`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *AddonCreateUpdateSchema) GetEventsOk() (*[]string, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *AddonCreateUpdateSchema) SetEvents(v []string)`

SetEvents sets Events field to given value.


### GetProjects

`func (o *AddonCreateUpdateSchema) GetProjects() []string`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *AddonCreateUpdateSchema) GetProjectsOk() (*[]string, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *AddonCreateUpdateSchema) SetProjects(v []string)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *AddonCreateUpdateSchema) HasProjects() bool`

HasProjects returns a boolean if a field has been set.

### GetEnvironments

`func (o *AddonCreateUpdateSchema) GetEnvironments() []string`

GetEnvironments returns the Environments field if non-nil, zero value otherwise.

### GetEnvironmentsOk

`func (o *AddonCreateUpdateSchema) GetEnvironmentsOk() (*[]string, bool)`

GetEnvironmentsOk returns a tuple with the Environments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironments

`func (o *AddonCreateUpdateSchema) SetEnvironments(v []string)`

SetEnvironments sets Environments field to given value.

### HasEnvironments

`func (o *AddonCreateUpdateSchema) HasEnvironments() bool`

HasEnvironments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


