/*
Unleash API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 6.1.10+main
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the ProjectAccessConfigurationSchemaRolesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectAccessConfigurationSchemaRolesInner{}

// ProjectAccessConfigurationSchemaRolesInner struct for ProjectAccessConfigurationSchemaRolesInner
type ProjectAccessConfigurationSchemaRolesInner struct {
	// The id of the role.
	Id *int32 `json:"id,omitempty"`
	// A list of group ids that will be assigned this role
	Groups []int32 `json:"groups,omitempty"`
	// A list of user ids that will be assigned this role
	Users                []int32 `json:"users,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProjectAccessConfigurationSchemaRolesInner ProjectAccessConfigurationSchemaRolesInner

// NewProjectAccessConfigurationSchemaRolesInner instantiates a new ProjectAccessConfigurationSchemaRolesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectAccessConfigurationSchemaRolesInner() *ProjectAccessConfigurationSchemaRolesInner {
	this := ProjectAccessConfigurationSchemaRolesInner{}
	return &this
}

// NewProjectAccessConfigurationSchemaRolesInnerWithDefaults instantiates a new ProjectAccessConfigurationSchemaRolesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectAccessConfigurationSchemaRolesInnerWithDefaults() *ProjectAccessConfigurationSchemaRolesInner {
	this := ProjectAccessConfigurationSchemaRolesInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ProjectAccessConfigurationSchemaRolesInner) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectAccessConfigurationSchemaRolesInner) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ProjectAccessConfigurationSchemaRolesInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ProjectAccessConfigurationSchemaRolesInner) SetId(v int32) {
	o.Id = &v
}

// GetGroups returns the Groups field value if set, zero value otherwise.
func (o *ProjectAccessConfigurationSchemaRolesInner) GetGroups() []int32 {
	if o == nil || IsNil(o.Groups) {
		var ret []int32
		return ret
	}
	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectAccessConfigurationSchemaRolesInner) GetGroupsOk() ([]int32, bool) {
	if o == nil || IsNil(o.Groups) {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *ProjectAccessConfigurationSchemaRolesInner) HasGroups() bool {
	if o != nil && !IsNil(o.Groups) {
		return true
	}

	return false
}

// SetGroups gets a reference to the given []int32 and assigns it to the Groups field.
func (o *ProjectAccessConfigurationSchemaRolesInner) SetGroups(v []int32) {
	o.Groups = v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *ProjectAccessConfigurationSchemaRolesInner) GetUsers() []int32 {
	if o == nil || IsNil(o.Users) {
		var ret []int32
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectAccessConfigurationSchemaRolesInner) GetUsersOk() ([]int32, bool) {
	if o == nil || IsNil(o.Users) {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *ProjectAccessConfigurationSchemaRolesInner) HasUsers() bool {
	if o != nil && !IsNil(o.Users) {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []int32 and assigns it to the Users field.
func (o *ProjectAccessConfigurationSchemaRolesInner) SetUsers(v []int32) {
	o.Users = v
}

func (o ProjectAccessConfigurationSchemaRolesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectAccessConfigurationSchemaRolesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Groups) {
		toSerialize["groups"] = o.Groups
	}
	if !IsNil(o.Users) {
		toSerialize["users"] = o.Users
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProjectAccessConfigurationSchemaRolesInner) UnmarshalJSON(data []byte) (err error) {
	varProjectAccessConfigurationSchemaRolesInner := _ProjectAccessConfigurationSchemaRolesInner{}

	err = json.Unmarshal(data, &varProjectAccessConfigurationSchemaRolesInner)

	if err != nil {
		return err
	}

	*o = ProjectAccessConfigurationSchemaRolesInner(varProjectAccessConfigurationSchemaRolesInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "groups")
		delete(additionalProperties, "users")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProjectAccessConfigurationSchemaRolesInner struct {
	value *ProjectAccessConfigurationSchemaRolesInner
	isSet bool
}

func (v NullableProjectAccessConfigurationSchemaRolesInner) Get() *ProjectAccessConfigurationSchemaRolesInner {
	return v.value
}

func (v *NullableProjectAccessConfigurationSchemaRolesInner) Set(val *ProjectAccessConfigurationSchemaRolesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectAccessConfigurationSchemaRolesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectAccessConfigurationSchemaRolesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectAccessConfigurationSchemaRolesInner(val *ProjectAccessConfigurationSchemaRolesInner) *NullableProjectAccessConfigurationSchemaRolesInner {
	return &NullableProjectAccessConfigurationSchemaRolesInner{value: val, isSet: true}
}

func (v NullableProjectAccessConfigurationSchemaRolesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectAccessConfigurationSchemaRolesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
