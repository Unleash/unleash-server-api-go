/*
Unleash API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 5.2.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the VariantSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VariantSchema{}

// VariantSchema A variant allows for further separation of users into segments. See [our excellent documentation](https://docs.getunleash.io/reference/feature-toggle-variants#what-are-variants) for a more detailed description
type VariantSchema struct {
	// The variants name. Is unique for this feature toggle
	Name string `json:"name"`
	// The weight is the likelihood of any one user getting this variant. It is a number between 0 and 1000. See the section on [variant weights](https://docs.getunleash.io/reference/feature-toggle-variants#variant-weight) for more information
	Weight float32 `json:"weight"`
	// Set to fix if this variant must have exactly the weight allocated to it. If the type is variable, the weight will adjust so that the total weight of all variants adds up to 1000
	WeightType *string `json:"weightType,omitempty"`
	// [Stickiness](https://docs.getunleash.io/reference/feature-toggle-variants#variant-stickiness) is how Unleash guarantees that the same user gets the same variant every time
	Stickiness *string               `json:"stickiness,omitempty"`
	Payload    *VariantSchemaPayload `json:"payload,omitempty"`
	// Overrides assigning specific variants to specific users. The weighting system automatically assigns users to specific groups for you, but any overrides in this list will take precedence.
	Overrides []OverrideSchema `json:"overrides,omitempty"`
}

// NewVariantSchema instantiates a new VariantSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVariantSchema(name string, weight float32) *VariantSchema {
	this := VariantSchema{}
	this.Name = name
	this.Weight = weight
	return &this
}

// NewVariantSchemaWithDefaults instantiates a new VariantSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVariantSchemaWithDefaults() *VariantSchema {
	this := VariantSchema{}
	return &this
}

// GetName returns the Name field value
func (o *VariantSchema) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *VariantSchema) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *VariantSchema) SetName(v string) {
	o.Name = v
}

// GetWeight returns the Weight field value
func (o *VariantSchema) GetWeight() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Weight
}

// GetWeightOk returns a tuple with the Weight field value
// and a boolean to check if the value has been set.
func (o *VariantSchema) GetWeightOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Weight, true
}

// SetWeight sets field value
func (o *VariantSchema) SetWeight(v float32) {
	o.Weight = v
}

// GetWeightType returns the WeightType field value if set, zero value otherwise.
func (o *VariantSchema) GetWeightType() string {
	if o == nil || IsNil(o.WeightType) {
		var ret string
		return ret
	}
	return *o.WeightType
}

// GetWeightTypeOk returns a tuple with the WeightType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariantSchema) GetWeightTypeOk() (*string, bool) {
	if o == nil || IsNil(o.WeightType) {
		return nil, false
	}
	return o.WeightType, true
}

// HasWeightType returns a boolean if a field has been set.
func (o *VariantSchema) HasWeightType() bool {
	if o != nil && !IsNil(o.WeightType) {
		return true
	}

	return false
}

// SetWeightType gets a reference to the given string and assigns it to the WeightType field.
func (o *VariantSchema) SetWeightType(v string) {
	o.WeightType = &v
}

// GetStickiness returns the Stickiness field value if set, zero value otherwise.
func (o *VariantSchema) GetStickiness() string {
	if o == nil || IsNil(o.Stickiness) {
		var ret string
		return ret
	}
	return *o.Stickiness
}

// GetStickinessOk returns a tuple with the Stickiness field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariantSchema) GetStickinessOk() (*string, bool) {
	if o == nil || IsNil(o.Stickiness) {
		return nil, false
	}
	return o.Stickiness, true
}

// HasStickiness returns a boolean if a field has been set.
func (o *VariantSchema) HasStickiness() bool {
	if o != nil && !IsNil(o.Stickiness) {
		return true
	}

	return false
}

// SetStickiness gets a reference to the given string and assigns it to the Stickiness field.
func (o *VariantSchema) SetStickiness(v string) {
	o.Stickiness = &v
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *VariantSchema) GetPayload() VariantSchemaPayload {
	if o == nil || IsNil(o.Payload) {
		var ret VariantSchemaPayload
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariantSchema) GetPayloadOk() (*VariantSchemaPayload, bool) {
	if o == nil || IsNil(o.Payload) {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *VariantSchema) HasPayload() bool {
	if o != nil && !IsNil(o.Payload) {
		return true
	}

	return false
}

// SetPayload gets a reference to the given VariantSchemaPayload and assigns it to the Payload field.
func (o *VariantSchema) SetPayload(v VariantSchemaPayload) {
	o.Payload = &v
}

// GetOverrides returns the Overrides field value if set, zero value otherwise.
func (o *VariantSchema) GetOverrides() []OverrideSchema {
	if o == nil || IsNil(o.Overrides) {
		var ret []OverrideSchema
		return ret
	}
	return o.Overrides
}

// GetOverridesOk returns a tuple with the Overrides field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariantSchema) GetOverridesOk() ([]OverrideSchema, bool) {
	if o == nil || IsNil(o.Overrides) {
		return nil, false
	}
	return o.Overrides, true
}

// HasOverrides returns a boolean if a field has been set.
func (o *VariantSchema) HasOverrides() bool {
	if o != nil && !IsNil(o.Overrides) {
		return true
	}

	return false
}

// SetOverrides gets a reference to the given []OverrideSchema and assigns it to the Overrides field.
func (o *VariantSchema) SetOverrides(v []OverrideSchema) {
	o.Overrides = v
}

func (o VariantSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VariantSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["weight"] = o.Weight
	if !IsNil(o.WeightType) {
		toSerialize["weightType"] = o.WeightType
	}
	if !IsNil(o.Stickiness) {
		toSerialize["stickiness"] = o.Stickiness
	}
	if !IsNil(o.Payload) {
		toSerialize["payload"] = o.Payload
	}
	if !IsNil(o.Overrides) {
		toSerialize["overrides"] = o.Overrides
	}
	return toSerialize, nil
}

type NullableVariantSchema struct {
	value *VariantSchema
	isSet bool
}

func (v NullableVariantSchema) Get() *VariantSchema {
	return v.value
}

func (v *NullableVariantSchema) Set(val *VariantSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableVariantSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableVariantSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVariantSchema(val *VariantSchema) *NullableVariantSchema {
	return &NullableVariantSchema{value: val, isSet: true}
}

func (v NullableVariantSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVariantSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
