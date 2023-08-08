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

// checks if the AdminPermissionSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdminPermissionSchema{}

// AdminPermissionSchema Describes a single permission
type AdminPermissionSchema struct {
	// The identifier for this permission
	Id int32 `json:"id"`
	// The name of this permission
	Name string `json:"name"`
	// The name to display in listings of permissions
	DisplayName string `json:"displayName"`
	// What level this permission applies to. Either root, project or the name of the environment it applies to
	Type string `json:"type"`
	// Which environment this permission applies to
	Environment *string `json:"environment,omitempty"`
}

// NewAdminPermissionSchema instantiates a new AdminPermissionSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdminPermissionSchema(id int32, name string, displayName string, type_ string) *AdminPermissionSchema {
	this := AdminPermissionSchema{}
	this.Id = id
	this.Name = name
	this.DisplayName = displayName
	this.Type = type_
	return &this
}

// NewAdminPermissionSchemaWithDefaults instantiates a new AdminPermissionSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdminPermissionSchemaWithDefaults() *AdminPermissionSchema {
	this := AdminPermissionSchema{}
	return &this
}

// GetId returns the Id field value
func (o *AdminPermissionSchema) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AdminPermissionSchema) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AdminPermissionSchema) SetId(v int32) {
	o.Id = v
}

// GetName returns the Name field value
func (o *AdminPermissionSchema) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AdminPermissionSchema) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AdminPermissionSchema) SetName(v string) {
	o.Name = v
}

// GetDisplayName returns the DisplayName field value
func (o *AdminPermissionSchema) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *AdminPermissionSchema) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *AdminPermissionSchema) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetType returns the Type field value
func (o *AdminPermissionSchema) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AdminPermissionSchema) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AdminPermissionSchema) SetType(v string) {
	o.Type = v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *AdminPermissionSchema) GetEnvironment() string {
	if o == nil || IsNil(o.Environment) {
		var ret string
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminPermissionSchema) GetEnvironmentOk() (*string, bool) {
	if o == nil || IsNil(o.Environment) {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *AdminPermissionSchema) HasEnvironment() bool {
	if o != nil && !IsNil(o.Environment) {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given string and assigns it to the Environment field.
func (o *AdminPermissionSchema) SetEnvironment(v string) {
	o.Environment = &v
}

func (o AdminPermissionSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdminPermissionSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["displayName"] = o.DisplayName
	toSerialize["type"] = o.Type
	if !IsNil(o.Environment) {
		toSerialize["environment"] = o.Environment
	}
	return toSerialize, nil
}

type NullableAdminPermissionSchema struct {
	value *AdminPermissionSchema
	isSet bool
}

func (v NullableAdminPermissionSchema) Get() *AdminPermissionSchema {
	return v.value
}

func (v *NullableAdminPermissionSchema) Set(val *AdminPermissionSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableAdminPermissionSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableAdminPermissionSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdminPermissionSchema(val *AdminPermissionSchema) *NullableAdminPermissionSchema {
	return &NullableAdminPermissionSchema{value: val, isSet: true}
}

func (v NullableAdminPermissionSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdminPermissionSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
