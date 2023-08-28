/*
Unleash API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 5.4.0-main
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the AddonTypeSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddonTypeSchema{}

// AddonTypeSchema An addon provider. Defines a specific addon type and what the end user must configure when creating a new addon of that type.
type AddonTypeSchema struct {
	// The name of the addon type. When creating new addons, this goes in the payload's `type` field.
	Name string `json:"name"`
	// The addon type's name as it should be displayed in the admin UI.
	DisplayName string `json:"displayName"`
	// A URL to where you can find more information about using this addon type.
	DocumentationUrl string `json:"documentationUrl"`
	// A description of the addon type.
	Description string `json:"description"`
	// A list of [Unleash tag types](https://docs.getunleash.io/reference/tags#tag-types) that this addon uses. These tags will be added to the Unleash instance when an addon of this type is created.
	TagTypes []TagTypeSchema `json:"tagTypes,omitempty"`
	// The addon provider's parameters. Use these to configure an addon of this provider type. Items with `required: true` must be provided.
	Parameters []AddonParameterSchema `json:"parameters,omitempty"`
	// All the [event types](https://docs.getunleash.io/reference/api/legacy/unleash/admin/events#feature-toggle-events) that are available for this addon provider.
	Events       []string                     `json:"events,omitempty"`
	Installation *AddonTypeSchemaInstallation `json:"installation,omitempty"`
	// A list of alerts to display to the user when installing addons of this type.
	Alerts []AddonTypeSchemaAlertsInner `json:"alerts,omitempty"`
	// This should be used to inform the user that this addon type is deprecated and should not be used. Deprecated addons will show a badge with this information on the UI.
	Deprecated *string `json:"deprecated,omitempty"`
}

// NewAddonTypeSchema instantiates a new AddonTypeSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddonTypeSchema(name string, displayName string, documentationUrl string, description string) *AddonTypeSchema {
	this := AddonTypeSchema{}
	this.Name = name
	this.DisplayName = displayName
	this.DocumentationUrl = documentationUrl
	this.Description = description
	return &this
}

// NewAddonTypeSchemaWithDefaults instantiates a new AddonTypeSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddonTypeSchemaWithDefaults() *AddonTypeSchema {
	this := AddonTypeSchema{}
	return &this
}

// GetName returns the Name field value
func (o *AddonTypeSchema) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AddonTypeSchema) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AddonTypeSchema) SetName(v string) {
	o.Name = v
}

// GetDisplayName returns the DisplayName field value
func (o *AddonTypeSchema) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *AddonTypeSchema) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *AddonTypeSchema) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetDocumentationUrl returns the DocumentationUrl field value
func (o *AddonTypeSchema) GetDocumentationUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DocumentationUrl
}

// GetDocumentationUrlOk returns a tuple with the DocumentationUrl field value
// and a boolean to check if the value has been set.
func (o *AddonTypeSchema) GetDocumentationUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DocumentationUrl, true
}

// SetDocumentationUrl sets field value
func (o *AddonTypeSchema) SetDocumentationUrl(v string) {
	o.DocumentationUrl = v
}

// GetDescription returns the Description field value
func (o *AddonTypeSchema) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *AddonTypeSchema) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *AddonTypeSchema) SetDescription(v string) {
	o.Description = v
}

// GetTagTypes returns the TagTypes field value if set, zero value otherwise.
func (o *AddonTypeSchema) GetTagTypes() []TagTypeSchema {
	if o == nil || IsNil(o.TagTypes) {
		var ret []TagTypeSchema
		return ret
	}
	return o.TagTypes
}

// GetTagTypesOk returns a tuple with the TagTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddonTypeSchema) GetTagTypesOk() ([]TagTypeSchema, bool) {
	if o == nil || IsNil(o.TagTypes) {
		return nil, false
	}
	return o.TagTypes, true
}

// HasTagTypes returns a boolean if a field has been set.
func (o *AddonTypeSchema) HasTagTypes() bool {
	if o != nil && !IsNil(o.TagTypes) {
		return true
	}

	return false
}

// SetTagTypes gets a reference to the given []TagTypeSchema and assigns it to the TagTypes field.
func (o *AddonTypeSchema) SetTagTypes(v []TagTypeSchema) {
	o.TagTypes = v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *AddonTypeSchema) GetParameters() []AddonParameterSchema {
	if o == nil || IsNil(o.Parameters) {
		var ret []AddonParameterSchema
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddonTypeSchema) GetParametersOk() ([]AddonParameterSchema, bool) {
	if o == nil || IsNil(o.Parameters) {
		return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *AddonTypeSchema) HasParameters() bool {
	if o != nil && !IsNil(o.Parameters) {
		return true
	}

	return false
}

// SetParameters gets a reference to the given []AddonParameterSchema and assigns it to the Parameters field.
func (o *AddonTypeSchema) SetParameters(v []AddonParameterSchema) {
	o.Parameters = v
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *AddonTypeSchema) GetEvents() []string {
	if o == nil || IsNil(o.Events) {
		var ret []string
		return ret
	}
	return o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddonTypeSchema) GetEventsOk() ([]string, bool) {
	if o == nil || IsNil(o.Events) {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *AddonTypeSchema) HasEvents() bool {
	if o != nil && !IsNil(o.Events) {
		return true
	}

	return false
}

// SetEvents gets a reference to the given []string and assigns it to the Events field.
func (o *AddonTypeSchema) SetEvents(v []string) {
	o.Events = v
}

// GetInstallation returns the Installation field value if set, zero value otherwise.
func (o *AddonTypeSchema) GetInstallation() AddonTypeSchemaInstallation {
	if o == nil || IsNil(o.Installation) {
		var ret AddonTypeSchemaInstallation
		return ret
	}
	return *o.Installation
}

// GetInstallationOk returns a tuple with the Installation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddonTypeSchema) GetInstallationOk() (*AddonTypeSchemaInstallation, bool) {
	if o == nil || IsNil(o.Installation) {
		return nil, false
	}
	return o.Installation, true
}

// HasInstallation returns a boolean if a field has been set.
func (o *AddonTypeSchema) HasInstallation() bool {
	if o != nil && !IsNil(o.Installation) {
		return true
	}

	return false
}

// SetInstallation gets a reference to the given AddonTypeSchemaInstallation and assigns it to the Installation field.
func (o *AddonTypeSchema) SetInstallation(v AddonTypeSchemaInstallation) {
	o.Installation = &v
}

// GetAlerts returns the Alerts field value if set, zero value otherwise.
func (o *AddonTypeSchema) GetAlerts() []AddonTypeSchemaAlertsInner {
	if o == nil || IsNil(o.Alerts) {
		var ret []AddonTypeSchemaAlertsInner
		return ret
	}
	return o.Alerts
}

// GetAlertsOk returns a tuple with the Alerts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddonTypeSchema) GetAlertsOk() ([]AddonTypeSchemaAlertsInner, bool) {
	if o == nil || IsNil(o.Alerts) {
		return nil, false
	}
	return o.Alerts, true
}

// HasAlerts returns a boolean if a field has been set.
func (o *AddonTypeSchema) HasAlerts() bool {
	if o != nil && !IsNil(o.Alerts) {
		return true
	}

	return false
}

// SetAlerts gets a reference to the given []AddonTypeSchemaAlertsInner and assigns it to the Alerts field.
func (o *AddonTypeSchema) SetAlerts(v []AddonTypeSchemaAlertsInner) {
	o.Alerts = v
}

// GetDeprecated returns the Deprecated field value if set, zero value otherwise.
func (o *AddonTypeSchema) GetDeprecated() string {
	if o == nil || IsNil(o.Deprecated) {
		var ret string
		return ret
	}
	return *o.Deprecated
}

// GetDeprecatedOk returns a tuple with the Deprecated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddonTypeSchema) GetDeprecatedOk() (*string, bool) {
	if o == nil || IsNil(o.Deprecated) {
		return nil, false
	}
	return o.Deprecated, true
}

// HasDeprecated returns a boolean if a field has been set.
func (o *AddonTypeSchema) HasDeprecated() bool {
	if o != nil && !IsNil(o.Deprecated) {
		return true
	}

	return false
}

// SetDeprecated gets a reference to the given string and assigns it to the Deprecated field.
func (o *AddonTypeSchema) SetDeprecated(v string) {
	o.Deprecated = &v
}

func (o AddonTypeSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddonTypeSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["displayName"] = o.DisplayName
	toSerialize["documentationUrl"] = o.DocumentationUrl
	toSerialize["description"] = o.Description
	if !IsNil(o.TagTypes) {
		toSerialize["tagTypes"] = o.TagTypes
	}
	if !IsNil(o.Parameters) {
		toSerialize["parameters"] = o.Parameters
	}
	if !IsNil(o.Events) {
		toSerialize["events"] = o.Events
	}
	if !IsNil(o.Installation) {
		toSerialize["installation"] = o.Installation
	}
	if !IsNil(o.Alerts) {
		toSerialize["alerts"] = o.Alerts
	}
	if !IsNil(o.Deprecated) {
		toSerialize["deprecated"] = o.Deprecated
	}
	return toSerialize, nil
}

type NullableAddonTypeSchema struct {
	value *AddonTypeSchema
	isSet bool
}

func (v NullableAddonTypeSchema) Get() *AddonTypeSchema {
	return v.value
}

func (v *NullableAddonTypeSchema) Set(val *AddonTypeSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableAddonTypeSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableAddonTypeSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddonTypeSchema(val *AddonTypeSchema) *NullableAddonTypeSchema {
	return &NullableAddonTypeSchema{value: val, isSet: true}
}

func (v NullableAddonTypeSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddonTypeSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
