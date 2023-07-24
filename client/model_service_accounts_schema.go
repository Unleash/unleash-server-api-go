/*
Unleash API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 5.2.0-main
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the ServiceAccountsSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceAccountsSchema{}

// ServiceAccountsSchema Represents a list of service accounts, and includes a list of root roles they reference
type ServiceAccountsSchema struct {
	// A list of service accounts
	ServiceAccounts []ServiceAccountSchema `json:"serviceAccounts"`
	// A list of root roles that are referenced from service account objects in the `serviceAccounts` list
	RootRoles []RoleSchema `json:"rootRoles,omitempty"`
}

// NewServiceAccountsSchema instantiates a new ServiceAccountsSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceAccountsSchema(serviceAccounts []ServiceAccountSchema) *ServiceAccountsSchema {
	this := ServiceAccountsSchema{}
	this.ServiceAccounts = serviceAccounts
	return &this
}

// NewServiceAccountsSchemaWithDefaults instantiates a new ServiceAccountsSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceAccountsSchemaWithDefaults() *ServiceAccountsSchema {
	this := ServiceAccountsSchema{}
	return &this
}

// GetServiceAccounts returns the ServiceAccounts field value
func (o *ServiceAccountsSchema) GetServiceAccounts() []ServiceAccountSchema {
	if o == nil {
		var ret []ServiceAccountSchema
		return ret
	}

	return o.ServiceAccounts
}

// GetServiceAccountsOk returns a tuple with the ServiceAccounts field value
// and a boolean to check if the value has been set.
func (o *ServiceAccountsSchema) GetServiceAccountsOk() ([]ServiceAccountSchema, bool) {
	if o == nil {
		return nil, false
	}
	return o.ServiceAccounts, true
}

// SetServiceAccounts sets field value
func (o *ServiceAccountsSchema) SetServiceAccounts(v []ServiceAccountSchema) {
	o.ServiceAccounts = v
}

// GetRootRoles returns the RootRoles field value if set, zero value otherwise.
func (o *ServiceAccountsSchema) GetRootRoles() []RoleSchema {
	if o == nil || IsNil(o.RootRoles) {
		var ret []RoleSchema
		return ret
	}
	return o.RootRoles
}

// GetRootRolesOk returns a tuple with the RootRoles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccountsSchema) GetRootRolesOk() ([]RoleSchema, bool) {
	if o == nil || IsNil(o.RootRoles) {
		return nil, false
	}
	return o.RootRoles, true
}

// HasRootRoles returns a boolean if a field has been set.
func (o *ServiceAccountsSchema) HasRootRoles() bool {
	if o != nil && !IsNil(o.RootRoles) {
		return true
	}

	return false
}

// SetRootRoles gets a reference to the given []RoleSchema and assigns it to the RootRoles field.
func (o *ServiceAccountsSchema) SetRootRoles(v []RoleSchema) {
	o.RootRoles = v
}

func (o ServiceAccountsSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceAccountsSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["serviceAccounts"] = o.ServiceAccounts
	if !IsNil(o.RootRoles) {
		toSerialize["rootRoles"] = o.RootRoles
	}
	return toSerialize, nil
}

type NullableServiceAccountsSchema struct {
	value *ServiceAccountsSchema
	isSet bool
}

func (v NullableServiceAccountsSchema) Get() *ServiceAccountsSchema {
	return v.value
}

func (v *NullableServiceAccountsSchema) Set(val *ServiceAccountsSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceAccountsSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceAccountsSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceAccountsSchema(val *ServiceAccountsSchema) *NullableServiceAccountsSchema {
	return &NullableServiceAccountsSchema{value: val, isSet: true}
}

func (v NullableServiceAccountsSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceAccountsSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
