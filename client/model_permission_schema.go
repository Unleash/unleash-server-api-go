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

// checks if the PermissionSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PermissionSchema{}

// PermissionSchema struct for PermissionSchema
type PermissionSchema struct {
	Permission  string  `json:"permission"`
	Project     *string `json:"project,omitempty"`
	Environment *string `json:"environment,omitempty"`
}

// NewPermissionSchema instantiates a new PermissionSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPermissionSchema(permission string) *PermissionSchema {
	this := PermissionSchema{}
	this.Permission = permission
	return &this
}

// NewPermissionSchemaWithDefaults instantiates a new PermissionSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPermissionSchemaWithDefaults() *PermissionSchema {
	this := PermissionSchema{}
	return &this
}

// GetPermission returns the Permission field value
func (o *PermissionSchema) GetPermission() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Permission
}

// GetPermissionOk returns a tuple with the Permission field value
// and a boolean to check if the value has been set.
func (o *PermissionSchema) GetPermissionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Permission, true
}

// SetPermission sets field value
func (o *PermissionSchema) SetPermission(v string) {
	o.Permission = v
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *PermissionSchema) GetProject() string {
	if o == nil || IsNil(o.Project) {
		var ret string
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionSchema) GetProjectOk() (*string, bool) {
	if o == nil || IsNil(o.Project) {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *PermissionSchema) HasProject() bool {
	if o != nil && !IsNil(o.Project) {
		return true
	}

	return false
}

// SetProject gets a reference to the given string and assigns it to the Project field.
func (o *PermissionSchema) SetProject(v string) {
	o.Project = &v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *PermissionSchema) GetEnvironment() string {
	if o == nil || IsNil(o.Environment) {
		var ret string
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionSchema) GetEnvironmentOk() (*string, bool) {
	if o == nil || IsNil(o.Environment) {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *PermissionSchema) HasEnvironment() bool {
	if o != nil && !IsNil(o.Environment) {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given string and assigns it to the Environment field.
func (o *PermissionSchema) SetEnvironment(v string) {
	o.Environment = &v
}

func (o PermissionSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PermissionSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["permission"] = o.Permission
	if !IsNil(o.Project) {
		toSerialize["project"] = o.Project
	}
	if !IsNil(o.Environment) {
		toSerialize["environment"] = o.Environment
	}
	return toSerialize, nil
}

type NullablePermissionSchema struct {
	value *PermissionSchema
	isSet bool
}

func (v NullablePermissionSchema) Get() *PermissionSchema {
	return v.value
}

func (v *NullablePermissionSchema) Set(val *PermissionSchema) {
	v.value = val
	v.isSet = true
}

func (v NullablePermissionSchema) IsSet() bool {
	return v.isSet
}

func (v *NullablePermissionSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePermissionSchema(val *PermissionSchema) *NullablePermissionSchema {
	return &NullablePermissionSchema{value: val, isSet: true}
}

func (v NullablePermissionSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePermissionSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
