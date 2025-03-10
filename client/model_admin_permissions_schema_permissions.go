/*
Unleash API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 6.8.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
)

// checks if the AdminPermissionsSchemaPermissions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdminPermissionsSchemaPermissions{}

// AdminPermissionsSchemaPermissions Returns permissions available at all three levels (root|project|environment)
type AdminPermissionsSchemaPermissions struct {
	// Permissions available at the root level, i.e. not connected to any specific project or environment
	Root []AdminPermissionSchema `json:"root,omitempty"`
	// Permissions available at the project level
	Project []AdminPermissionSchema `json:"project"`
	// A list of environments with available permissions per environment
	Environments         []AdminPermissionsSchemaPermissionsEnvironmentsInner `json:"environments"`
	AdditionalProperties map[string]interface{}
}

type _AdminPermissionsSchemaPermissions AdminPermissionsSchemaPermissions

// NewAdminPermissionsSchemaPermissions instantiates a new AdminPermissionsSchemaPermissions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdminPermissionsSchemaPermissions(project []AdminPermissionSchema, environments []AdminPermissionsSchemaPermissionsEnvironmentsInner) *AdminPermissionsSchemaPermissions {
	this := AdminPermissionsSchemaPermissions{}
	this.Project = project
	this.Environments = environments
	return &this
}

// NewAdminPermissionsSchemaPermissionsWithDefaults instantiates a new AdminPermissionsSchemaPermissions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdminPermissionsSchemaPermissionsWithDefaults() *AdminPermissionsSchemaPermissions {
	this := AdminPermissionsSchemaPermissions{}
	return &this
}

// GetRoot returns the Root field value if set, zero value otherwise.
func (o *AdminPermissionsSchemaPermissions) GetRoot() []AdminPermissionSchema {
	if o == nil || IsNil(o.Root) {
		var ret []AdminPermissionSchema
		return ret
	}
	return o.Root
}

// GetRootOk returns a tuple with the Root field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminPermissionsSchemaPermissions) GetRootOk() ([]AdminPermissionSchema, bool) {
	if o == nil || IsNil(o.Root) {
		return nil, false
	}
	return o.Root, true
}

// HasRoot returns a boolean if a field has been set.
func (o *AdminPermissionsSchemaPermissions) HasRoot() bool {
	if o != nil && !IsNil(o.Root) {
		return true
	}

	return false
}

// SetRoot gets a reference to the given []AdminPermissionSchema and assigns it to the Root field.
func (o *AdminPermissionsSchemaPermissions) SetRoot(v []AdminPermissionSchema) {
	o.Root = v
}

// GetProject returns the Project field value
func (o *AdminPermissionsSchemaPermissions) GetProject() []AdminPermissionSchema {
	if o == nil {
		var ret []AdminPermissionSchema
		return ret
	}

	return o.Project
}

// GetProjectOk returns a tuple with the Project field value
// and a boolean to check if the value has been set.
func (o *AdminPermissionsSchemaPermissions) GetProjectOk() ([]AdminPermissionSchema, bool) {
	if o == nil {
		return nil, false
	}
	return o.Project, true
}

// SetProject sets field value
func (o *AdminPermissionsSchemaPermissions) SetProject(v []AdminPermissionSchema) {
	o.Project = v
}

// GetEnvironments returns the Environments field value
func (o *AdminPermissionsSchemaPermissions) GetEnvironments() []AdminPermissionsSchemaPermissionsEnvironmentsInner {
	if o == nil {
		var ret []AdminPermissionsSchemaPermissionsEnvironmentsInner
		return ret
	}

	return o.Environments
}

// GetEnvironmentsOk returns a tuple with the Environments field value
// and a boolean to check if the value has been set.
func (o *AdminPermissionsSchemaPermissions) GetEnvironmentsOk() ([]AdminPermissionsSchemaPermissionsEnvironmentsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Environments, true
}

// SetEnvironments sets field value
func (o *AdminPermissionsSchemaPermissions) SetEnvironments(v []AdminPermissionsSchemaPermissionsEnvironmentsInner) {
	o.Environments = v
}

func (o AdminPermissionsSchemaPermissions) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdminPermissionsSchemaPermissions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Root) {
		toSerialize["root"] = o.Root
	}
	toSerialize["project"] = o.Project
	toSerialize["environments"] = o.Environments

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AdminPermissionsSchemaPermissions) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"project",
		"environments",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varAdminPermissionsSchemaPermissions := _AdminPermissionsSchemaPermissions{}

	err = json.Unmarshal(data, &varAdminPermissionsSchemaPermissions)

	if err != nil {
		return err
	}

	*o = AdminPermissionsSchemaPermissions(varAdminPermissionsSchemaPermissions)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "root")
		delete(additionalProperties, "project")
		delete(additionalProperties, "environments")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAdminPermissionsSchemaPermissions struct {
	value *AdminPermissionsSchemaPermissions
	isSet bool
}

func (v NullableAdminPermissionsSchemaPermissions) Get() *AdminPermissionsSchemaPermissions {
	return v.value
}

func (v *NullableAdminPermissionsSchemaPermissions) Set(val *AdminPermissionsSchemaPermissions) {
	v.value = val
	v.isSet = true
}

func (v NullableAdminPermissionsSchemaPermissions) IsSet() bool {
	return v.isSet
}

func (v *NullableAdminPermissionsSchemaPermissions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdminPermissionsSchemaPermissions(val *AdminPermissionsSchemaPermissions) *NullableAdminPermissionsSchemaPermissions {
	return &NullableAdminPermissionsSchemaPermissions{value: val, isSet: true}
}

func (v NullableAdminPermissionsSchemaPermissions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdminPermissionsSchemaPermissions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
