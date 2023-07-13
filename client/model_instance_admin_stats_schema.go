/*
Unleash API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 5.2.0-main
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// checks if the InstanceAdminStatsSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InstanceAdminStatsSchema{}

// InstanceAdminStatsSchema Information about an instance and statistics about usage of various features of Unleash
type InstanceAdminStatsSchema struct {
	// A unique identifier for this instance. Generated by the database migration scripts at first run. Typically a UUID.
	InstanceId string `json:"instanceId"`
	// When these statistics were produced
	Timestamp NullableTime `json:"timestamp,omitempty"`
	// The version of Unleash OSS that is bundled in this instance
	VersionOSS *string `json:"versionOSS,omitempty"`
	// The version of Unleash Enterprise that is bundled in this instance
	VersionEnterprise *string `json:"versionEnterprise,omitempty"`
	// The number of users this instance has
	Users *float32 `json:"users,omitempty"`
	// The number of feature-toggles this instance has
	FeatureToggles *float32 `json:"featureToggles,omitempty"`
	// The number of projects defined in this instance.
	Projects *float32 `json:"projects,omitempty"`
	// The number of context fields defined in this instance.
	ContextFields *float32 `json:"contextFields,omitempty"`
	// The number of roles defined in this instance
	Roles *float32 `json:"roles,omitempty"`
	// The number of groups defined in this instance
	Groups *float32 `json:"groups,omitempty"`
	// The number of environments defined in this instance
	Environments *float32 `json:"environments,omitempty"`
	// The number of segments defined in this instance
	Segments *float32 `json:"segments,omitempty"`
	// The number of strategies defined in this instance
	Strategies *float32 `json:"strategies,omitempty"`
	// Whether or not SAML authentication is enabled for this instance
	SAMLenabled *bool `json:"SAMLenabled,omitempty"`
	// Whether or not OIDC authentication is enabled for this instance
	OIDCenabled *bool `json:"OIDCenabled,omitempty"`
	// A count of connected applications in the last week, last month and all time since last restart
	ClientApps []InstanceAdminStatsSchemaClientAppsInner `json:"clientApps,omitempty"`
	// A SHA-256 checksum of the instance statistics to be used to verify that the data in this object has not been tampered with
	Sum *string `json:"sum,omitempty"`
}

// NewInstanceAdminStatsSchema instantiates a new InstanceAdminStatsSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstanceAdminStatsSchema(instanceId string) *InstanceAdminStatsSchema {
	this := InstanceAdminStatsSchema{}
	this.InstanceId = instanceId
	return &this
}

// NewInstanceAdminStatsSchemaWithDefaults instantiates a new InstanceAdminStatsSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstanceAdminStatsSchemaWithDefaults() *InstanceAdminStatsSchema {
	this := InstanceAdminStatsSchema{}
	return &this
}

// GetInstanceId returns the InstanceId field value
func (o *InstanceAdminStatsSchema) GetInstanceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InstanceId
}

// GetInstanceIdOk returns a tuple with the InstanceId field value
// and a boolean to check if the value has been set.
func (o *InstanceAdminStatsSchema) GetInstanceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstanceId, true
}

// SetInstanceId sets field value
func (o *InstanceAdminStatsSchema) SetInstanceId(v string) {
	o.InstanceId = v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InstanceAdminStatsSchema) GetTimestamp() time.Time {
	if o == nil || IsNil(o.Timestamp.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Timestamp.Get()
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InstanceAdminStatsSchema) GetTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Timestamp.Get(), o.Timestamp.IsSet()
}

// HasTimestamp returns a boolean if a field has been set.
func (o *InstanceAdminStatsSchema) HasTimestamp() bool {
	if o != nil && o.Timestamp.IsSet() {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given NullableTime and assigns it to the Timestamp field.
func (o *InstanceAdminStatsSchema) SetTimestamp(v time.Time) {
	o.Timestamp.Set(&v)
}
// SetTimestampNil sets the value for Timestamp to be an explicit nil
func (o *InstanceAdminStatsSchema) SetTimestampNil() {
	o.Timestamp.Set(nil)
}

// UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
func (o *InstanceAdminStatsSchema) UnsetTimestamp() {
	o.Timestamp.Unset()
}

// GetVersionOSS returns the VersionOSS field value if set, zero value otherwise.
func (o *InstanceAdminStatsSchema) GetVersionOSS() string {
	if o == nil || IsNil(o.VersionOSS) {
		var ret string
		return ret
	}
	return *o.VersionOSS
}

// GetVersionOSSOk returns a tuple with the VersionOSS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceAdminStatsSchema) GetVersionOSSOk() (*string, bool) {
	if o == nil || IsNil(o.VersionOSS) {
		return nil, false
	}
	return o.VersionOSS, true
}

// HasVersionOSS returns a boolean if a field has been set.
func (o *InstanceAdminStatsSchema) HasVersionOSS() bool {
	if o != nil && !IsNil(o.VersionOSS) {
		return true
	}

	return false
}

// SetVersionOSS gets a reference to the given string and assigns it to the VersionOSS field.
func (o *InstanceAdminStatsSchema) SetVersionOSS(v string) {
	o.VersionOSS = &v
}

// GetVersionEnterprise returns the VersionEnterprise field value if set, zero value otherwise.
func (o *InstanceAdminStatsSchema) GetVersionEnterprise() string {
	if o == nil || IsNil(o.VersionEnterprise) {
		var ret string
		return ret
	}
	return *o.VersionEnterprise
}

// GetVersionEnterpriseOk returns a tuple with the VersionEnterprise field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceAdminStatsSchema) GetVersionEnterpriseOk() (*string, bool) {
	if o == nil || IsNil(o.VersionEnterprise) {
		return nil, false
	}
	return o.VersionEnterprise, true
}

// HasVersionEnterprise returns a boolean if a field has been set.
func (o *InstanceAdminStatsSchema) HasVersionEnterprise() bool {
	if o != nil && !IsNil(o.VersionEnterprise) {
		return true
	}

	return false
}

// SetVersionEnterprise gets a reference to the given string and assigns it to the VersionEnterprise field.
func (o *InstanceAdminStatsSchema) SetVersionEnterprise(v string) {
	o.VersionEnterprise = &v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *InstanceAdminStatsSchema) GetUsers() float32 {
	if o == nil || IsNil(o.Users) {
		var ret float32
		return ret
	}
	return *o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceAdminStatsSchema) GetUsersOk() (*float32, bool) {
	if o == nil || IsNil(o.Users) {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *InstanceAdminStatsSchema) HasUsers() bool {
	if o != nil && !IsNil(o.Users) {
		return true
	}

	return false
}

// SetUsers gets a reference to the given float32 and assigns it to the Users field.
func (o *InstanceAdminStatsSchema) SetUsers(v float32) {
	o.Users = &v
}

// GetFeatureToggles returns the FeatureToggles field value if set, zero value otherwise.
func (o *InstanceAdminStatsSchema) GetFeatureToggles() float32 {
	if o == nil || IsNil(o.FeatureToggles) {
		var ret float32
		return ret
	}
	return *o.FeatureToggles
}

// GetFeatureTogglesOk returns a tuple with the FeatureToggles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceAdminStatsSchema) GetFeatureTogglesOk() (*float32, bool) {
	if o == nil || IsNil(o.FeatureToggles) {
		return nil, false
	}
	return o.FeatureToggles, true
}

// HasFeatureToggles returns a boolean if a field has been set.
func (o *InstanceAdminStatsSchema) HasFeatureToggles() bool {
	if o != nil && !IsNil(o.FeatureToggles) {
		return true
	}

	return false
}

// SetFeatureToggles gets a reference to the given float32 and assigns it to the FeatureToggles field.
func (o *InstanceAdminStatsSchema) SetFeatureToggles(v float32) {
	o.FeatureToggles = &v
}

// GetProjects returns the Projects field value if set, zero value otherwise.
func (o *InstanceAdminStatsSchema) GetProjects() float32 {
	if o == nil || IsNil(o.Projects) {
		var ret float32
		return ret
	}
	return *o.Projects
}

// GetProjectsOk returns a tuple with the Projects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceAdminStatsSchema) GetProjectsOk() (*float32, bool) {
	if o == nil || IsNil(o.Projects) {
		return nil, false
	}
	return o.Projects, true
}

// HasProjects returns a boolean if a field has been set.
func (o *InstanceAdminStatsSchema) HasProjects() bool {
	if o != nil && !IsNil(o.Projects) {
		return true
	}

	return false
}

// SetProjects gets a reference to the given float32 and assigns it to the Projects field.
func (o *InstanceAdminStatsSchema) SetProjects(v float32) {
	o.Projects = &v
}

// GetContextFields returns the ContextFields field value if set, zero value otherwise.
func (o *InstanceAdminStatsSchema) GetContextFields() float32 {
	if o == nil || IsNil(o.ContextFields) {
		var ret float32
		return ret
	}
	return *o.ContextFields
}

// GetContextFieldsOk returns a tuple with the ContextFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceAdminStatsSchema) GetContextFieldsOk() (*float32, bool) {
	if o == nil || IsNil(o.ContextFields) {
		return nil, false
	}
	return o.ContextFields, true
}

// HasContextFields returns a boolean if a field has been set.
func (o *InstanceAdminStatsSchema) HasContextFields() bool {
	if o != nil && !IsNil(o.ContextFields) {
		return true
	}

	return false
}

// SetContextFields gets a reference to the given float32 and assigns it to the ContextFields field.
func (o *InstanceAdminStatsSchema) SetContextFields(v float32) {
	o.ContextFields = &v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *InstanceAdminStatsSchema) GetRoles() float32 {
	if o == nil || IsNil(o.Roles) {
		var ret float32
		return ret
	}
	return *o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceAdminStatsSchema) GetRolesOk() (*float32, bool) {
	if o == nil || IsNil(o.Roles) {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *InstanceAdminStatsSchema) HasRoles() bool {
	if o != nil && !IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given float32 and assigns it to the Roles field.
func (o *InstanceAdminStatsSchema) SetRoles(v float32) {
	o.Roles = &v
}

// GetGroups returns the Groups field value if set, zero value otherwise.
func (o *InstanceAdminStatsSchema) GetGroups() float32 {
	if o == nil || IsNil(o.Groups) {
		var ret float32
		return ret
	}
	return *o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceAdminStatsSchema) GetGroupsOk() (*float32, bool) {
	if o == nil || IsNil(o.Groups) {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *InstanceAdminStatsSchema) HasGroups() bool {
	if o != nil && !IsNil(o.Groups) {
		return true
	}

	return false
}

// SetGroups gets a reference to the given float32 and assigns it to the Groups field.
func (o *InstanceAdminStatsSchema) SetGroups(v float32) {
	o.Groups = &v
}

// GetEnvironments returns the Environments field value if set, zero value otherwise.
func (o *InstanceAdminStatsSchema) GetEnvironments() float32 {
	if o == nil || IsNil(o.Environments) {
		var ret float32
		return ret
	}
	return *o.Environments
}

// GetEnvironmentsOk returns a tuple with the Environments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceAdminStatsSchema) GetEnvironmentsOk() (*float32, bool) {
	if o == nil || IsNil(o.Environments) {
		return nil, false
	}
	return o.Environments, true
}

// HasEnvironments returns a boolean if a field has been set.
func (o *InstanceAdminStatsSchema) HasEnvironments() bool {
	if o != nil && !IsNil(o.Environments) {
		return true
	}

	return false
}

// SetEnvironments gets a reference to the given float32 and assigns it to the Environments field.
func (o *InstanceAdminStatsSchema) SetEnvironments(v float32) {
	o.Environments = &v
}

// GetSegments returns the Segments field value if set, zero value otherwise.
func (o *InstanceAdminStatsSchema) GetSegments() float32 {
	if o == nil || IsNil(o.Segments) {
		var ret float32
		return ret
	}
	return *o.Segments
}

// GetSegmentsOk returns a tuple with the Segments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceAdminStatsSchema) GetSegmentsOk() (*float32, bool) {
	if o == nil || IsNil(o.Segments) {
		return nil, false
	}
	return o.Segments, true
}

// HasSegments returns a boolean if a field has been set.
func (o *InstanceAdminStatsSchema) HasSegments() bool {
	if o != nil && !IsNil(o.Segments) {
		return true
	}

	return false
}

// SetSegments gets a reference to the given float32 and assigns it to the Segments field.
func (o *InstanceAdminStatsSchema) SetSegments(v float32) {
	o.Segments = &v
}

// GetStrategies returns the Strategies field value if set, zero value otherwise.
func (o *InstanceAdminStatsSchema) GetStrategies() float32 {
	if o == nil || IsNil(o.Strategies) {
		var ret float32
		return ret
	}
	return *o.Strategies
}

// GetStrategiesOk returns a tuple with the Strategies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceAdminStatsSchema) GetStrategiesOk() (*float32, bool) {
	if o == nil || IsNil(o.Strategies) {
		return nil, false
	}
	return o.Strategies, true
}

// HasStrategies returns a boolean if a field has been set.
func (o *InstanceAdminStatsSchema) HasStrategies() bool {
	if o != nil && !IsNil(o.Strategies) {
		return true
	}

	return false
}

// SetStrategies gets a reference to the given float32 and assigns it to the Strategies field.
func (o *InstanceAdminStatsSchema) SetStrategies(v float32) {
	o.Strategies = &v
}

// GetSAMLenabled returns the SAMLenabled field value if set, zero value otherwise.
func (o *InstanceAdminStatsSchema) GetSAMLenabled() bool {
	if o == nil || IsNil(o.SAMLenabled) {
		var ret bool
		return ret
	}
	return *o.SAMLenabled
}

// GetSAMLenabledOk returns a tuple with the SAMLenabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceAdminStatsSchema) GetSAMLenabledOk() (*bool, bool) {
	if o == nil || IsNil(o.SAMLenabled) {
		return nil, false
	}
	return o.SAMLenabled, true
}

// HasSAMLenabled returns a boolean if a field has been set.
func (o *InstanceAdminStatsSchema) HasSAMLenabled() bool {
	if o != nil && !IsNil(o.SAMLenabled) {
		return true
	}

	return false
}

// SetSAMLenabled gets a reference to the given bool and assigns it to the SAMLenabled field.
func (o *InstanceAdminStatsSchema) SetSAMLenabled(v bool) {
	o.SAMLenabled = &v
}

// GetOIDCenabled returns the OIDCenabled field value if set, zero value otherwise.
func (o *InstanceAdminStatsSchema) GetOIDCenabled() bool {
	if o == nil || IsNil(o.OIDCenabled) {
		var ret bool
		return ret
	}
	return *o.OIDCenabled
}

// GetOIDCenabledOk returns a tuple with the OIDCenabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceAdminStatsSchema) GetOIDCenabledOk() (*bool, bool) {
	if o == nil || IsNil(o.OIDCenabled) {
		return nil, false
	}
	return o.OIDCenabled, true
}

// HasOIDCenabled returns a boolean if a field has been set.
func (o *InstanceAdminStatsSchema) HasOIDCenabled() bool {
	if o != nil && !IsNil(o.OIDCenabled) {
		return true
	}

	return false
}

// SetOIDCenabled gets a reference to the given bool and assigns it to the OIDCenabled field.
func (o *InstanceAdminStatsSchema) SetOIDCenabled(v bool) {
	o.OIDCenabled = &v
}

// GetClientApps returns the ClientApps field value if set, zero value otherwise.
func (o *InstanceAdminStatsSchema) GetClientApps() []InstanceAdminStatsSchemaClientAppsInner {
	if o == nil || IsNil(o.ClientApps) {
		var ret []InstanceAdminStatsSchemaClientAppsInner
		return ret
	}
	return o.ClientApps
}

// GetClientAppsOk returns a tuple with the ClientApps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceAdminStatsSchema) GetClientAppsOk() ([]InstanceAdminStatsSchemaClientAppsInner, bool) {
	if o == nil || IsNil(o.ClientApps) {
		return nil, false
	}
	return o.ClientApps, true
}

// HasClientApps returns a boolean if a field has been set.
func (o *InstanceAdminStatsSchema) HasClientApps() bool {
	if o != nil && !IsNil(o.ClientApps) {
		return true
	}

	return false
}

// SetClientApps gets a reference to the given []InstanceAdminStatsSchemaClientAppsInner and assigns it to the ClientApps field.
func (o *InstanceAdminStatsSchema) SetClientApps(v []InstanceAdminStatsSchemaClientAppsInner) {
	o.ClientApps = v
}

// GetSum returns the Sum field value if set, zero value otherwise.
func (o *InstanceAdminStatsSchema) GetSum() string {
	if o == nil || IsNil(o.Sum) {
		var ret string
		return ret
	}
	return *o.Sum
}

// GetSumOk returns a tuple with the Sum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceAdminStatsSchema) GetSumOk() (*string, bool) {
	if o == nil || IsNil(o.Sum) {
		return nil, false
	}
	return o.Sum, true
}

// HasSum returns a boolean if a field has been set.
func (o *InstanceAdminStatsSchema) HasSum() bool {
	if o != nil && !IsNil(o.Sum) {
		return true
	}

	return false
}

// SetSum gets a reference to the given string and assigns it to the Sum field.
func (o *InstanceAdminStatsSchema) SetSum(v string) {
	o.Sum = &v
}

func (o InstanceAdminStatsSchema) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InstanceAdminStatsSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["instanceId"] = o.InstanceId
	if o.Timestamp.IsSet() {
		toSerialize["timestamp"] = o.Timestamp.Get()
	}
	if !IsNil(o.VersionOSS) {
		toSerialize["versionOSS"] = o.VersionOSS
	}
	if !IsNil(o.VersionEnterprise) {
		toSerialize["versionEnterprise"] = o.VersionEnterprise
	}
	if !IsNil(o.Users) {
		toSerialize["users"] = o.Users
	}
	if !IsNil(o.FeatureToggles) {
		toSerialize["featureToggles"] = o.FeatureToggles
	}
	if !IsNil(o.Projects) {
		toSerialize["projects"] = o.Projects
	}
	if !IsNil(o.ContextFields) {
		toSerialize["contextFields"] = o.ContextFields
	}
	if !IsNil(o.Roles) {
		toSerialize["roles"] = o.Roles
	}
	if !IsNil(o.Groups) {
		toSerialize["groups"] = o.Groups
	}
	if !IsNil(o.Environments) {
		toSerialize["environments"] = o.Environments
	}
	if !IsNil(o.Segments) {
		toSerialize["segments"] = o.Segments
	}
	if !IsNil(o.Strategies) {
		toSerialize["strategies"] = o.Strategies
	}
	if !IsNil(o.SAMLenabled) {
		toSerialize["SAMLenabled"] = o.SAMLenabled
	}
	if !IsNil(o.OIDCenabled) {
		toSerialize["OIDCenabled"] = o.OIDCenabled
	}
	if !IsNil(o.ClientApps) {
		toSerialize["clientApps"] = o.ClientApps
	}
	if !IsNil(o.Sum) {
		toSerialize["sum"] = o.Sum
	}
	return toSerialize, nil
}

type NullableInstanceAdminStatsSchema struct {
	value *InstanceAdminStatsSchema
	isSet bool
}

func (v NullableInstanceAdminStatsSchema) Get() *InstanceAdminStatsSchema {
	return v.value
}

func (v *NullableInstanceAdminStatsSchema) Set(val *InstanceAdminStatsSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableInstanceAdminStatsSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableInstanceAdminStatsSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstanceAdminStatsSchema(val *InstanceAdminStatsSchema) *NullableInstanceAdminStatsSchema {
	return &NullableInstanceAdminStatsSchema{value: val, isSet: true}
}

func (v NullableInstanceAdminStatsSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstanceAdminStatsSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


