# AddonTypeSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the addon type. When creating new addons, this goes in the payload&#39;s &#x60;type&#x60; field. | 
**DisplayName** | **string** | The addon type&#39;s name as it should be displayed in the admin UI. | 
**DocumentationUrl** | **string** | A URL to where you can find more information about using this addon type. | 
**Description** | **string** | A description of the addon type. | 
**TagTypes** | Pointer to [**[]TagTypeSchema**](TagTypeSchema.md) | A list of [Unleash tag types](https://docs.getunleash.io/reference/tags#tag-types) that this addon uses. These tags will be added to the Unleash instance when an addon of this type is created. | [optional] 
**Parameters** | Pointer to [**[]AddonParameterSchema**](AddonParameterSchema.md) | The addon provider&#39;s parameters. Use these to configure an addon of this provider type. Items with &#x60;required: true&#x60; must be provided. | [optional] 
**Events** | Pointer to **[]string** | All the [event types](https://docs.getunleash.io/reference/api/legacy/unleash/admin/events#feature-toggle-events) that are available for this addon provider. | [optional] 
**Installation** | Pointer to [**AddonTypeSchemaInstallation**](AddonTypeSchemaInstallation.md) |  | [optional] 
**Alerts** | Pointer to [**[]AddonTypeSchemaAlertsInner**](AddonTypeSchemaAlertsInner.md) | A list of alerts to display to the user when installing addons of this type. | [optional] 
**Deprecated** | Pointer to **string** | This should be used to inform the user that this addon type is deprecated and should not be used. Deprecated addons will show a badge with this information on the UI. | [optional] 

## Methods

### NewAddonTypeSchema

`func NewAddonTypeSchema(name string, displayName string, documentationUrl string, description string, ) *AddonTypeSchema`

NewAddonTypeSchema instantiates a new AddonTypeSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddonTypeSchemaWithDefaults

`func NewAddonTypeSchemaWithDefaults() *AddonTypeSchema`

NewAddonTypeSchemaWithDefaults instantiates a new AddonTypeSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AddonTypeSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddonTypeSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddonTypeSchema) SetName(v string)`

SetName sets Name field to given value.


### GetDisplayName

`func (o *AddonTypeSchema) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AddonTypeSchema) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AddonTypeSchema) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetDocumentationUrl

`func (o *AddonTypeSchema) GetDocumentationUrl() string`

GetDocumentationUrl returns the DocumentationUrl field if non-nil, zero value otherwise.

### GetDocumentationUrlOk

`func (o *AddonTypeSchema) GetDocumentationUrlOk() (*string, bool)`

GetDocumentationUrlOk returns a tuple with the DocumentationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentationUrl

`func (o *AddonTypeSchema) SetDocumentationUrl(v string)`

SetDocumentationUrl sets DocumentationUrl field to given value.


### GetDescription

`func (o *AddonTypeSchema) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddonTypeSchema) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddonTypeSchema) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetTagTypes

`func (o *AddonTypeSchema) GetTagTypes() []TagTypeSchema`

GetTagTypes returns the TagTypes field if non-nil, zero value otherwise.

### GetTagTypesOk

`func (o *AddonTypeSchema) GetTagTypesOk() (*[]TagTypeSchema, bool)`

GetTagTypesOk returns a tuple with the TagTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagTypes

`func (o *AddonTypeSchema) SetTagTypes(v []TagTypeSchema)`

SetTagTypes sets TagTypes field to given value.

### HasTagTypes

`func (o *AddonTypeSchema) HasTagTypes() bool`

HasTagTypes returns a boolean if a field has been set.

### GetParameters

`func (o *AddonTypeSchema) GetParameters() []AddonParameterSchema`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *AddonTypeSchema) GetParametersOk() (*[]AddonParameterSchema, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *AddonTypeSchema) SetParameters(v []AddonParameterSchema)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *AddonTypeSchema) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetEvents

`func (o *AddonTypeSchema) GetEvents() []string`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *AddonTypeSchema) GetEventsOk() (*[]string, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *AddonTypeSchema) SetEvents(v []string)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *AddonTypeSchema) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetInstallation

`func (o *AddonTypeSchema) GetInstallation() AddonTypeSchemaInstallation`

GetInstallation returns the Installation field if non-nil, zero value otherwise.

### GetInstallationOk

`func (o *AddonTypeSchema) GetInstallationOk() (*AddonTypeSchemaInstallation, bool)`

GetInstallationOk returns a tuple with the Installation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallation

`func (o *AddonTypeSchema) SetInstallation(v AddonTypeSchemaInstallation)`

SetInstallation sets Installation field to given value.

### HasInstallation

`func (o *AddonTypeSchema) HasInstallation() bool`

HasInstallation returns a boolean if a field has been set.

### GetAlerts

`func (o *AddonTypeSchema) GetAlerts() []AddonTypeSchemaAlertsInner`

GetAlerts returns the Alerts field if non-nil, zero value otherwise.

### GetAlertsOk

`func (o *AddonTypeSchema) GetAlertsOk() (*[]AddonTypeSchemaAlertsInner, bool)`

GetAlertsOk returns a tuple with the Alerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlerts

`func (o *AddonTypeSchema) SetAlerts(v []AddonTypeSchemaAlertsInner)`

SetAlerts sets Alerts field to given value.

### HasAlerts

`func (o *AddonTypeSchema) HasAlerts() bool`

HasAlerts returns a boolean if a field has been set.

### GetDeprecated

`func (o *AddonTypeSchema) GetDeprecated() string`

GetDeprecated returns the Deprecated field if non-nil, zero value otherwise.

### GetDeprecatedOk

`func (o *AddonTypeSchema) GetDeprecatedOk() (*string, bool)`

GetDeprecatedOk returns a tuple with the Deprecated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecated

`func (o *AddonTypeSchema) SetDeprecated(v string)`

SetDeprecated sets Deprecated field to given value.

### HasDeprecated

`func (o *AddonTypeSchema) HasDeprecated() bool`

HasDeprecated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


