# ApplicationSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppName** | **string** | Name of the application | 
**SdkVersion** | Pointer to **string** | Which SDK and version the application reporting uses. Typically represented as &#x60;&lt;identifier&gt;:&lt;version&gt;&#x60; | [optional] 
**Strategies** | Pointer to **[]string** | Which [strategies](https://docs.getunleash.io/topics/the-anatomy-of-unleash#activation-strategies) the application has loaded. Useful when trying to figure out if your [custom strategy](https://docs.getunleash.io/reference/custom-activation-strategies) has been loaded in the SDK | [optional] 
**Description** | Pointer to **string** | Extra information added about the application reporting the metrics. Only present if added via the Unleash Admin interface | [optional] 
**Url** | Pointer to **string** | A link to reference the application reporting the metrics. Could for instance be a GitHub link to the repository of the application | [optional] 
**Color** | Pointer to **string** | The CSS color that is used to color the application&#39;s entry in the application list | [optional] 
**Icon** | Pointer to **string** | An URL to an icon file to be used for the applications&#39;s entry in the application list | [optional] 

## Methods

### NewApplicationSchema

`func NewApplicationSchema(appName string, ) *ApplicationSchema`

NewApplicationSchema instantiates a new ApplicationSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationSchemaWithDefaults

`func NewApplicationSchemaWithDefaults() *ApplicationSchema`

NewApplicationSchemaWithDefaults instantiates a new ApplicationSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppName

`func (o *ApplicationSchema) GetAppName() string`

GetAppName returns the AppName field if non-nil, zero value otherwise.

### GetAppNameOk

`func (o *ApplicationSchema) GetAppNameOk() (*string, bool)`

GetAppNameOk returns a tuple with the AppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppName

`func (o *ApplicationSchema) SetAppName(v string)`

SetAppName sets AppName field to given value.


### GetSdkVersion

`func (o *ApplicationSchema) GetSdkVersion() string`

GetSdkVersion returns the SdkVersion field if non-nil, zero value otherwise.

### GetSdkVersionOk

`func (o *ApplicationSchema) GetSdkVersionOk() (*string, bool)`

GetSdkVersionOk returns a tuple with the SdkVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdkVersion

`func (o *ApplicationSchema) SetSdkVersion(v string)`

SetSdkVersion sets SdkVersion field to given value.

### HasSdkVersion

`func (o *ApplicationSchema) HasSdkVersion() bool`

HasSdkVersion returns a boolean if a field has been set.

### GetStrategies

`func (o *ApplicationSchema) GetStrategies() []string`

GetStrategies returns the Strategies field if non-nil, zero value otherwise.

### GetStrategiesOk

`func (o *ApplicationSchema) GetStrategiesOk() (*[]string, bool)`

GetStrategiesOk returns a tuple with the Strategies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategies

`func (o *ApplicationSchema) SetStrategies(v []string)`

SetStrategies sets Strategies field to given value.

### HasStrategies

`func (o *ApplicationSchema) HasStrategies() bool`

HasStrategies returns a boolean if a field has been set.

### GetDescription

`func (o *ApplicationSchema) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApplicationSchema) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApplicationSchema) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApplicationSchema) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetUrl

`func (o *ApplicationSchema) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ApplicationSchema) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ApplicationSchema) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ApplicationSchema) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetColor

`func (o *ApplicationSchema) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *ApplicationSchema) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *ApplicationSchema) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *ApplicationSchema) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetIcon

`func (o *ApplicationSchema) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *ApplicationSchema) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *ApplicationSchema) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *ApplicationSchema) HasIcon() bool`

HasIcon returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


