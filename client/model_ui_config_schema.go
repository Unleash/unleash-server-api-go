/*
Unleash API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 5.3.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the UiConfigSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UiConfigSchema{}

// UiConfigSchema A collection of properties used to configure the Unleash Admin UI.
type UiConfigSchema struct {
	// The slogan to display in the UI footer.
	Slogan *string `json:"slogan,omitempty"`
	// The name of this Unleash instance. Used to build the text in the footer.
	Name *string `json:"name,omitempty"`
	// The current version of Unleash
	Version string `json:"version"`
	// What kind of Unleash instance it is: Enterprise, Pro, or Open source
	Environment *string `json:"environment,omitempty"`
	// The URL of the Unleash instance.
	UnleashUrl string `json:"unleashUrl"`
	// The base URI path at which this Unleash instance is listening.
	BaseUriPath string `json:"baseUriPath"`
	// Whether password authentication should be disabled or not.
	DisablePasswordAuth *bool `json:"disablePasswordAuth,omitempty"`
	// Whether this instance can send out emails or not.
	EmailEnabled *bool `json:"emailEnabled,omitempty"`
	// Whether maintenance mode is currently active or not.
	MaintenanceMode *bool `json:"maintenanceMode,omitempty"`
	// The maximum number of values that can be used in a single segment.
	SegmentValuesLimit *float32 `json:"segmentValuesLimit,omitempty"`
	// The maximum number of segments that can be applied to a single strategy.
	StrategySegmentsLimit *float32 `json:"strategySegmentsLimit,omitempty"`
	// Whether to enable the Unleash network view or not.
	NetworkViewEnabled *bool `json:"networkViewEnabled,omitempty"`
	// The list of origins that the front-end API should accept requests from.
	FrontendApiOrigins []string `json:"frontendApiOrigins,omitempty"`
	// Additional (largely experimental) features that are enabled in this Unleash instance.
	Flags *map[string]UiConfigSchemaFlagsValue `json:"flags,omitempty"`
	// Relevant links to use in the UI.
	Links []map[string]interface{} `json:"links,omitempty"`
	// The type of authentication enabled for this Unleash instance
	AuthenticationType *string       `json:"authenticationType,omitempty"`
	VersionInfo        VersionSchema `json:"versionInfo"`
}

// NewUiConfigSchema instantiates a new UiConfigSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUiConfigSchema(version string, unleashUrl string, baseUriPath string, versionInfo VersionSchema) *UiConfigSchema {
	this := UiConfigSchema{}
	this.Version = version
	this.UnleashUrl = unleashUrl
	this.BaseUriPath = baseUriPath
	this.VersionInfo = versionInfo
	return &this
}

// NewUiConfigSchemaWithDefaults instantiates a new UiConfigSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUiConfigSchemaWithDefaults() *UiConfigSchema {
	this := UiConfigSchema{}
	return &this
}

// GetSlogan returns the Slogan field value if set, zero value otherwise.
func (o *UiConfigSchema) GetSlogan() string {
	if o == nil || IsNil(o.Slogan) {
		var ret string
		return ret
	}
	return *o.Slogan
}

// GetSloganOk returns a tuple with the Slogan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UiConfigSchema) GetSloganOk() (*string, bool) {
	if o == nil || IsNil(o.Slogan) {
		return nil, false
	}
	return o.Slogan, true
}

// HasSlogan returns a boolean if a field has been set.
func (o *UiConfigSchema) HasSlogan() bool {
	if o != nil && !IsNil(o.Slogan) {
		return true
	}

	return false
}

// SetSlogan gets a reference to the given string and assigns it to the Slogan field.
func (o *UiConfigSchema) SetSlogan(v string) {
	o.Slogan = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UiConfigSchema) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UiConfigSchema) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UiConfigSchema) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UiConfigSchema) SetName(v string) {
	o.Name = &v
}

// GetVersion returns the Version field value
func (o *UiConfigSchema) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *UiConfigSchema) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *UiConfigSchema) SetVersion(v string) {
	o.Version = v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *UiConfigSchema) GetEnvironment() string {
	if o == nil || IsNil(o.Environment) {
		var ret string
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UiConfigSchema) GetEnvironmentOk() (*string, bool) {
	if o == nil || IsNil(o.Environment) {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *UiConfigSchema) HasEnvironment() bool {
	if o != nil && !IsNil(o.Environment) {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given string and assigns it to the Environment field.
func (o *UiConfigSchema) SetEnvironment(v string) {
	o.Environment = &v
}

// GetUnleashUrl returns the UnleashUrl field value
func (o *UiConfigSchema) GetUnleashUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UnleashUrl
}

// GetUnleashUrlOk returns a tuple with the UnleashUrl field value
// and a boolean to check if the value has been set.
func (o *UiConfigSchema) GetUnleashUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnleashUrl, true
}

// SetUnleashUrl sets field value
func (o *UiConfigSchema) SetUnleashUrl(v string) {
	o.UnleashUrl = v
}

// GetBaseUriPath returns the BaseUriPath field value
func (o *UiConfigSchema) GetBaseUriPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BaseUriPath
}

// GetBaseUriPathOk returns a tuple with the BaseUriPath field value
// and a boolean to check if the value has been set.
func (o *UiConfigSchema) GetBaseUriPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BaseUriPath, true
}

// SetBaseUriPath sets field value
func (o *UiConfigSchema) SetBaseUriPath(v string) {
	o.BaseUriPath = v
}

// GetDisablePasswordAuth returns the DisablePasswordAuth field value if set, zero value otherwise.
func (o *UiConfigSchema) GetDisablePasswordAuth() bool {
	if o == nil || IsNil(o.DisablePasswordAuth) {
		var ret bool
		return ret
	}
	return *o.DisablePasswordAuth
}

// GetDisablePasswordAuthOk returns a tuple with the DisablePasswordAuth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UiConfigSchema) GetDisablePasswordAuthOk() (*bool, bool) {
	if o == nil || IsNil(o.DisablePasswordAuth) {
		return nil, false
	}
	return o.DisablePasswordAuth, true
}

// HasDisablePasswordAuth returns a boolean if a field has been set.
func (o *UiConfigSchema) HasDisablePasswordAuth() bool {
	if o != nil && !IsNil(o.DisablePasswordAuth) {
		return true
	}

	return false
}

// SetDisablePasswordAuth gets a reference to the given bool and assigns it to the DisablePasswordAuth field.
func (o *UiConfigSchema) SetDisablePasswordAuth(v bool) {
	o.DisablePasswordAuth = &v
}

// GetEmailEnabled returns the EmailEnabled field value if set, zero value otherwise.
func (o *UiConfigSchema) GetEmailEnabled() bool {
	if o == nil || IsNil(o.EmailEnabled) {
		var ret bool
		return ret
	}
	return *o.EmailEnabled
}

// GetEmailEnabledOk returns a tuple with the EmailEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UiConfigSchema) GetEmailEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.EmailEnabled) {
		return nil, false
	}
	return o.EmailEnabled, true
}

// HasEmailEnabled returns a boolean if a field has been set.
func (o *UiConfigSchema) HasEmailEnabled() bool {
	if o != nil && !IsNil(o.EmailEnabled) {
		return true
	}

	return false
}

// SetEmailEnabled gets a reference to the given bool and assigns it to the EmailEnabled field.
func (o *UiConfigSchema) SetEmailEnabled(v bool) {
	o.EmailEnabled = &v
}

// GetMaintenanceMode returns the MaintenanceMode field value if set, zero value otherwise.
func (o *UiConfigSchema) GetMaintenanceMode() bool {
	if o == nil || IsNil(o.MaintenanceMode) {
		var ret bool
		return ret
	}
	return *o.MaintenanceMode
}

// GetMaintenanceModeOk returns a tuple with the MaintenanceMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UiConfigSchema) GetMaintenanceModeOk() (*bool, bool) {
	if o == nil || IsNil(o.MaintenanceMode) {
		return nil, false
	}
	return o.MaintenanceMode, true
}

// HasMaintenanceMode returns a boolean if a field has been set.
func (o *UiConfigSchema) HasMaintenanceMode() bool {
	if o != nil && !IsNil(o.MaintenanceMode) {
		return true
	}

	return false
}

// SetMaintenanceMode gets a reference to the given bool and assigns it to the MaintenanceMode field.
func (o *UiConfigSchema) SetMaintenanceMode(v bool) {
	o.MaintenanceMode = &v
}

// GetSegmentValuesLimit returns the SegmentValuesLimit field value if set, zero value otherwise.
func (o *UiConfigSchema) GetSegmentValuesLimit() float32 {
	if o == nil || IsNil(o.SegmentValuesLimit) {
		var ret float32
		return ret
	}
	return *o.SegmentValuesLimit
}

// GetSegmentValuesLimitOk returns a tuple with the SegmentValuesLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UiConfigSchema) GetSegmentValuesLimitOk() (*float32, bool) {
	if o == nil || IsNil(o.SegmentValuesLimit) {
		return nil, false
	}
	return o.SegmentValuesLimit, true
}

// HasSegmentValuesLimit returns a boolean if a field has been set.
func (o *UiConfigSchema) HasSegmentValuesLimit() bool {
	if o != nil && !IsNil(o.SegmentValuesLimit) {
		return true
	}

	return false
}

// SetSegmentValuesLimit gets a reference to the given float32 and assigns it to the SegmentValuesLimit field.
func (o *UiConfigSchema) SetSegmentValuesLimit(v float32) {
	o.SegmentValuesLimit = &v
}

// GetStrategySegmentsLimit returns the StrategySegmentsLimit field value if set, zero value otherwise.
func (o *UiConfigSchema) GetStrategySegmentsLimit() float32 {
	if o == nil || IsNil(o.StrategySegmentsLimit) {
		var ret float32
		return ret
	}
	return *o.StrategySegmentsLimit
}

// GetStrategySegmentsLimitOk returns a tuple with the StrategySegmentsLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UiConfigSchema) GetStrategySegmentsLimitOk() (*float32, bool) {
	if o == nil || IsNil(o.StrategySegmentsLimit) {
		return nil, false
	}
	return o.StrategySegmentsLimit, true
}

// HasStrategySegmentsLimit returns a boolean if a field has been set.
func (o *UiConfigSchema) HasStrategySegmentsLimit() bool {
	if o != nil && !IsNil(o.StrategySegmentsLimit) {
		return true
	}

	return false
}

// SetStrategySegmentsLimit gets a reference to the given float32 and assigns it to the StrategySegmentsLimit field.
func (o *UiConfigSchema) SetStrategySegmentsLimit(v float32) {
	o.StrategySegmentsLimit = &v
}

// GetNetworkViewEnabled returns the NetworkViewEnabled field value if set, zero value otherwise.
func (o *UiConfigSchema) GetNetworkViewEnabled() bool {
	if o == nil || IsNil(o.NetworkViewEnabled) {
		var ret bool
		return ret
	}
	return *o.NetworkViewEnabled
}

// GetNetworkViewEnabledOk returns a tuple with the NetworkViewEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UiConfigSchema) GetNetworkViewEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.NetworkViewEnabled) {
		return nil, false
	}
	return o.NetworkViewEnabled, true
}

// HasNetworkViewEnabled returns a boolean if a field has been set.
func (o *UiConfigSchema) HasNetworkViewEnabled() bool {
	if o != nil && !IsNil(o.NetworkViewEnabled) {
		return true
	}

	return false
}

// SetNetworkViewEnabled gets a reference to the given bool and assigns it to the NetworkViewEnabled field.
func (o *UiConfigSchema) SetNetworkViewEnabled(v bool) {
	o.NetworkViewEnabled = &v
}

// GetFrontendApiOrigins returns the FrontendApiOrigins field value if set, zero value otherwise.
func (o *UiConfigSchema) GetFrontendApiOrigins() []string {
	if o == nil || IsNil(o.FrontendApiOrigins) {
		var ret []string
		return ret
	}
	return o.FrontendApiOrigins
}

// GetFrontendApiOriginsOk returns a tuple with the FrontendApiOrigins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UiConfigSchema) GetFrontendApiOriginsOk() ([]string, bool) {
	if o == nil || IsNil(o.FrontendApiOrigins) {
		return nil, false
	}
	return o.FrontendApiOrigins, true
}

// HasFrontendApiOrigins returns a boolean if a field has been set.
func (o *UiConfigSchema) HasFrontendApiOrigins() bool {
	if o != nil && !IsNil(o.FrontendApiOrigins) {
		return true
	}

	return false
}

// SetFrontendApiOrigins gets a reference to the given []string and assigns it to the FrontendApiOrigins field.
func (o *UiConfigSchema) SetFrontendApiOrigins(v []string) {
	o.FrontendApiOrigins = v
}

// GetFlags returns the Flags field value if set, zero value otherwise.
func (o *UiConfigSchema) GetFlags() map[string]UiConfigSchemaFlagsValue {
	if o == nil || IsNil(o.Flags) {
		var ret map[string]UiConfigSchemaFlagsValue
		return ret
	}
	return *o.Flags
}

// GetFlagsOk returns a tuple with the Flags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UiConfigSchema) GetFlagsOk() (*map[string]UiConfigSchemaFlagsValue, bool) {
	if o == nil || IsNil(o.Flags) {
		return nil, false
	}
	return o.Flags, true
}

// HasFlags returns a boolean if a field has been set.
func (o *UiConfigSchema) HasFlags() bool {
	if o != nil && !IsNil(o.Flags) {
		return true
	}

	return false
}

// SetFlags gets a reference to the given map[string]UiConfigSchemaFlagsValue and assigns it to the Flags field.
func (o *UiConfigSchema) SetFlags(v map[string]UiConfigSchemaFlagsValue) {
	o.Flags = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *UiConfigSchema) GetLinks() []map[string]interface{} {
	if o == nil || IsNil(o.Links) {
		var ret []map[string]interface{}
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UiConfigSchema) GetLinksOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *UiConfigSchema) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []map[string]interface{} and assigns it to the Links field.
func (o *UiConfigSchema) SetLinks(v []map[string]interface{}) {
	o.Links = v
}

// GetAuthenticationType returns the AuthenticationType field value if set, zero value otherwise.
func (o *UiConfigSchema) GetAuthenticationType() string {
	if o == nil || IsNil(o.AuthenticationType) {
		var ret string
		return ret
	}
	return *o.AuthenticationType
}

// GetAuthenticationTypeOk returns a tuple with the AuthenticationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UiConfigSchema) GetAuthenticationTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AuthenticationType) {
		return nil, false
	}
	return o.AuthenticationType, true
}

// HasAuthenticationType returns a boolean if a field has been set.
func (o *UiConfigSchema) HasAuthenticationType() bool {
	if o != nil && !IsNil(o.AuthenticationType) {
		return true
	}

	return false
}

// SetAuthenticationType gets a reference to the given string and assigns it to the AuthenticationType field.
func (o *UiConfigSchema) SetAuthenticationType(v string) {
	o.AuthenticationType = &v
}

// GetVersionInfo returns the VersionInfo field value
func (o *UiConfigSchema) GetVersionInfo() VersionSchema {
	if o == nil {
		var ret VersionSchema
		return ret
	}

	return o.VersionInfo
}

// GetVersionInfoOk returns a tuple with the VersionInfo field value
// and a boolean to check if the value has been set.
func (o *UiConfigSchema) GetVersionInfoOk() (*VersionSchema, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VersionInfo, true
}

// SetVersionInfo sets field value
func (o *UiConfigSchema) SetVersionInfo(v VersionSchema) {
	o.VersionInfo = v
}

func (o UiConfigSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UiConfigSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Slogan) {
		toSerialize["slogan"] = o.Slogan
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	toSerialize["version"] = o.Version
	if !IsNil(o.Environment) {
		toSerialize["environment"] = o.Environment
	}
	toSerialize["unleashUrl"] = o.UnleashUrl
	toSerialize["baseUriPath"] = o.BaseUriPath
	if !IsNil(o.DisablePasswordAuth) {
		toSerialize["disablePasswordAuth"] = o.DisablePasswordAuth
	}
	if !IsNil(o.EmailEnabled) {
		toSerialize["emailEnabled"] = o.EmailEnabled
	}
	if !IsNil(o.MaintenanceMode) {
		toSerialize["maintenanceMode"] = o.MaintenanceMode
	}
	if !IsNil(o.SegmentValuesLimit) {
		toSerialize["segmentValuesLimit"] = o.SegmentValuesLimit
	}
	if !IsNil(o.StrategySegmentsLimit) {
		toSerialize["strategySegmentsLimit"] = o.StrategySegmentsLimit
	}
	if !IsNil(o.NetworkViewEnabled) {
		toSerialize["networkViewEnabled"] = o.NetworkViewEnabled
	}
	if !IsNil(o.FrontendApiOrigins) {
		toSerialize["frontendApiOrigins"] = o.FrontendApiOrigins
	}
	if !IsNil(o.Flags) {
		toSerialize["flags"] = o.Flags
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.AuthenticationType) {
		toSerialize["authenticationType"] = o.AuthenticationType
	}
	toSerialize["versionInfo"] = o.VersionInfo
	return toSerialize, nil
}

type NullableUiConfigSchema struct {
	value *UiConfigSchema
	isSet bool
}

func (v NullableUiConfigSchema) Get() *UiConfigSchema {
	return v.value
}

func (v *NullableUiConfigSchema) Set(val *UiConfigSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableUiConfigSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableUiConfigSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUiConfigSchema(val *UiConfigSchema) *NullableUiConfigSchema {
	return &NullableUiConfigSchema{value: val, isSet: true}
}

func (v NullableUiConfigSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUiConfigSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
