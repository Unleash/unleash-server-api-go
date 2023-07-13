# UiConfigSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Slogan** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Version** | **string** |  | 
**Environment** | Pointer to **string** |  | [optional] 
**UnleashUrl** | **string** |  | 
**BaseUriPath** | **string** |  | 
**DisablePasswordAuth** | Pointer to **bool** |  | [optional] 
**EmailEnabled** | Pointer to **bool** |  | [optional] 
**MaintenanceMode** | Pointer to **bool** |  | [optional] 
**SegmentValuesLimit** | Pointer to **float32** |  | [optional] 
**StrategySegmentsLimit** | Pointer to **float32** |  | [optional] 
**NetworkViewEnabled** | Pointer to **bool** |  | [optional] 
**FrontendApiOrigins** | Pointer to **[]string** |  | [optional] 
**Flags** | Pointer to [**map[string]UiConfigSchemaFlagsValue**](UiConfigSchemaFlagsValue.md) |  | [optional] 
**Links** | Pointer to **[]map[string]interface{}** |  | [optional] 
**AuthenticationType** | Pointer to **string** |  | [optional] 
**VersionInfo** | [**VersionSchema**](VersionSchema.md) |  | 

## Methods

### NewUiConfigSchema

`func NewUiConfigSchema(version string, unleashUrl string, baseUriPath string, versionInfo VersionSchema, ) *UiConfigSchema`

NewUiConfigSchema instantiates a new UiConfigSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUiConfigSchemaWithDefaults

`func NewUiConfigSchemaWithDefaults() *UiConfigSchema`

NewUiConfigSchemaWithDefaults instantiates a new UiConfigSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSlogan

`func (o *UiConfigSchema) GetSlogan() string`

GetSlogan returns the Slogan field if non-nil, zero value otherwise.

### GetSloganOk

`func (o *UiConfigSchema) GetSloganOk() (*string, bool)`

GetSloganOk returns a tuple with the Slogan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlogan

`func (o *UiConfigSchema) SetSlogan(v string)`

SetSlogan sets Slogan field to given value.

### HasSlogan

`func (o *UiConfigSchema) HasSlogan() bool`

HasSlogan returns a boolean if a field has been set.

### GetName

`func (o *UiConfigSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UiConfigSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UiConfigSchema) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UiConfigSchema) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVersion

`func (o *UiConfigSchema) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *UiConfigSchema) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *UiConfigSchema) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetEnvironment

`func (o *UiConfigSchema) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *UiConfigSchema) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *UiConfigSchema) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *UiConfigSchema) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetUnleashUrl

`func (o *UiConfigSchema) GetUnleashUrl() string`

GetUnleashUrl returns the UnleashUrl field if non-nil, zero value otherwise.

### GetUnleashUrlOk

`func (o *UiConfigSchema) GetUnleashUrlOk() (*string, bool)`

GetUnleashUrlOk returns a tuple with the UnleashUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnleashUrl

`func (o *UiConfigSchema) SetUnleashUrl(v string)`

SetUnleashUrl sets UnleashUrl field to given value.


### GetBaseUriPath

`func (o *UiConfigSchema) GetBaseUriPath() string`

GetBaseUriPath returns the BaseUriPath field if non-nil, zero value otherwise.

### GetBaseUriPathOk

`func (o *UiConfigSchema) GetBaseUriPathOk() (*string, bool)`

GetBaseUriPathOk returns a tuple with the BaseUriPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUriPath

`func (o *UiConfigSchema) SetBaseUriPath(v string)`

SetBaseUriPath sets BaseUriPath field to given value.


### GetDisablePasswordAuth

`func (o *UiConfigSchema) GetDisablePasswordAuth() bool`

GetDisablePasswordAuth returns the DisablePasswordAuth field if non-nil, zero value otherwise.

### GetDisablePasswordAuthOk

`func (o *UiConfigSchema) GetDisablePasswordAuthOk() (*bool, bool)`

GetDisablePasswordAuthOk returns a tuple with the DisablePasswordAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisablePasswordAuth

`func (o *UiConfigSchema) SetDisablePasswordAuth(v bool)`

SetDisablePasswordAuth sets DisablePasswordAuth field to given value.

### HasDisablePasswordAuth

`func (o *UiConfigSchema) HasDisablePasswordAuth() bool`

HasDisablePasswordAuth returns a boolean if a field has been set.

### GetEmailEnabled

`func (o *UiConfigSchema) GetEmailEnabled() bool`

GetEmailEnabled returns the EmailEnabled field if non-nil, zero value otherwise.

### GetEmailEnabledOk

`func (o *UiConfigSchema) GetEmailEnabledOk() (*bool, bool)`

GetEmailEnabledOk returns a tuple with the EmailEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailEnabled

`func (o *UiConfigSchema) SetEmailEnabled(v bool)`

SetEmailEnabled sets EmailEnabled field to given value.

### HasEmailEnabled

`func (o *UiConfigSchema) HasEmailEnabled() bool`

HasEmailEnabled returns a boolean if a field has been set.

### GetMaintenanceMode

`func (o *UiConfigSchema) GetMaintenanceMode() bool`

GetMaintenanceMode returns the MaintenanceMode field if non-nil, zero value otherwise.

### GetMaintenanceModeOk

`func (o *UiConfigSchema) GetMaintenanceModeOk() (*bool, bool)`

GetMaintenanceModeOk returns a tuple with the MaintenanceMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceMode

`func (o *UiConfigSchema) SetMaintenanceMode(v bool)`

SetMaintenanceMode sets MaintenanceMode field to given value.

### HasMaintenanceMode

`func (o *UiConfigSchema) HasMaintenanceMode() bool`

HasMaintenanceMode returns a boolean if a field has been set.

### GetSegmentValuesLimit

`func (o *UiConfigSchema) GetSegmentValuesLimit() float32`

GetSegmentValuesLimit returns the SegmentValuesLimit field if non-nil, zero value otherwise.

### GetSegmentValuesLimitOk

`func (o *UiConfigSchema) GetSegmentValuesLimitOk() (*float32, bool)`

GetSegmentValuesLimitOk returns a tuple with the SegmentValuesLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentValuesLimit

`func (o *UiConfigSchema) SetSegmentValuesLimit(v float32)`

SetSegmentValuesLimit sets SegmentValuesLimit field to given value.

### HasSegmentValuesLimit

`func (o *UiConfigSchema) HasSegmentValuesLimit() bool`

HasSegmentValuesLimit returns a boolean if a field has been set.

### GetStrategySegmentsLimit

`func (o *UiConfigSchema) GetStrategySegmentsLimit() float32`

GetStrategySegmentsLimit returns the StrategySegmentsLimit field if non-nil, zero value otherwise.

### GetStrategySegmentsLimitOk

`func (o *UiConfigSchema) GetStrategySegmentsLimitOk() (*float32, bool)`

GetStrategySegmentsLimitOk returns a tuple with the StrategySegmentsLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategySegmentsLimit

`func (o *UiConfigSchema) SetStrategySegmentsLimit(v float32)`

SetStrategySegmentsLimit sets StrategySegmentsLimit field to given value.

### HasStrategySegmentsLimit

`func (o *UiConfigSchema) HasStrategySegmentsLimit() bool`

HasStrategySegmentsLimit returns a boolean if a field has been set.

### GetNetworkViewEnabled

`func (o *UiConfigSchema) GetNetworkViewEnabled() bool`

GetNetworkViewEnabled returns the NetworkViewEnabled field if non-nil, zero value otherwise.

### GetNetworkViewEnabledOk

`func (o *UiConfigSchema) GetNetworkViewEnabledOk() (*bool, bool)`

GetNetworkViewEnabledOk returns a tuple with the NetworkViewEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkViewEnabled

`func (o *UiConfigSchema) SetNetworkViewEnabled(v bool)`

SetNetworkViewEnabled sets NetworkViewEnabled field to given value.

### HasNetworkViewEnabled

`func (o *UiConfigSchema) HasNetworkViewEnabled() bool`

HasNetworkViewEnabled returns a boolean if a field has been set.

### GetFrontendApiOrigins

`func (o *UiConfigSchema) GetFrontendApiOrigins() []string`

GetFrontendApiOrigins returns the FrontendApiOrigins field if non-nil, zero value otherwise.

### GetFrontendApiOriginsOk

`func (o *UiConfigSchema) GetFrontendApiOriginsOk() (*[]string, bool)`

GetFrontendApiOriginsOk returns a tuple with the FrontendApiOrigins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrontendApiOrigins

`func (o *UiConfigSchema) SetFrontendApiOrigins(v []string)`

SetFrontendApiOrigins sets FrontendApiOrigins field to given value.

### HasFrontendApiOrigins

`func (o *UiConfigSchema) HasFrontendApiOrigins() bool`

HasFrontendApiOrigins returns a boolean if a field has been set.

### GetFlags

`func (o *UiConfigSchema) GetFlags() map[string]UiConfigSchemaFlagsValue`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *UiConfigSchema) GetFlagsOk() (*map[string]UiConfigSchemaFlagsValue, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *UiConfigSchema) SetFlags(v map[string]UiConfigSchemaFlagsValue)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *UiConfigSchema) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### GetLinks

`func (o *UiConfigSchema) GetLinks() []map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *UiConfigSchema) GetLinksOk() (*[]map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *UiConfigSchema) SetLinks(v []map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *UiConfigSchema) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetAuthenticationType

`func (o *UiConfigSchema) GetAuthenticationType() string`

GetAuthenticationType returns the AuthenticationType field if non-nil, zero value otherwise.

### GetAuthenticationTypeOk

`func (o *UiConfigSchema) GetAuthenticationTypeOk() (*string, bool)`

GetAuthenticationTypeOk returns a tuple with the AuthenticationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationType

`func (o *UiConfigSchema) SetAuthenticationType(v string)`

SetAuthenticationType sets AuthenticationType field to given value.

### HasAuthenticationType

`func (o *UiConfigSchema) HasAuthenticationType() bool`

HasAuthenticationType returns a boolean if a field has been set.

### GetVersionInfo

`func (o *UiConfigSchema) GetVersionInfo() VersionSchema`

GetVersionInfo returns the VersionInfo field if non-nil, zero value otherwise.

### GetVersionInfoOk

`func (o *UiConfigSchema) GetVersionInfoOk() (*VersionSchema, bool)`

GetVersionInfoOk returns a tuple with the VersionInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionInfo

`func (o *UiConfigSchema) SetVersionInfo(v VersionSchema)`

SetVersionInfo sets VersionInfo field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


