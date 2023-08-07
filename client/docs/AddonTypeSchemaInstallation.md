# AddonTypeSchemaInstallation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | A URL to where the addon configuration should redirect to install addons of this type. | 
**Title** | Pointer to **string** | The title of the installation configuration. This will be displayed to the user when installing addons of this type. | [optional] 
**HelpText** | Pointer to **string** | The help text of the installation configuration. This will be displayed to the user when installing addons of this type. | [optional] 

## Methods

### NewAddonTypeSchemaInstallation

`func NewAddonTypeSchemaInstallation(url string, ) *AddonTypeSchemaInstallation`

NewAddonTypeSchemaInstallation instantiates a new AddonTypeSchemaInstallation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddonTypeSchemaInstallationWithDefaults

`func NewAddonTypeSchemaInstallationWithDefaults() *AddonTypeSchemaInstallation`

NewAddonTypeSchemaInstallationWithDefaults instantiates a new AddonTypeSchemaInstallation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *AddonTypeSchemaInstallation) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AddonTypeSchemaInstallation) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AddonTypeSchemaInstallation) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetTitle

`func (o *AddonTypeSchemaInstallation) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AddonTypeSchemaInstallation) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AddonTypeSchemaInstallation) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AddonTypeSchemaInstallation) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetHelpText

`func (o *AddonTypeSchemaInstallation) GetHelpText() string`

GetHelpText returns the HelpText field if non-nil, zero value otherwise.

### GetHelpTextOk

`func (o *AddonTypeSchemaInstallation) GetHelpTextOk() (*string, bool)`

GetHelpTextOk returns a tuple with the HelpText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelpText

`func (o *AddonTypeSchemaInstallation) SetHelpText(v string)`

SetHelpText sets HelpText field to given value.

### HasHelpText

`func (o *AddonTypeSchemaInstallation) HasHelpText() bool`

HasHelpText returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


