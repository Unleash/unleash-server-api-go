/*
Unleash API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 5.2.0-main
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// checks if the GroupSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupSchema{}

// GroupSchema struct for GroupSchema
type GroupSchema struct {
	Id          *float32       `json:"id,omitempty"`
	Name        string         `json:"name"`
	Description NullableString `json:"description,omitempty"`
	MappingsSSO []string       `json:"mappingsSSO,omitempty"`
	// A role id that is used as the root role for all users in this group. This can be either the id of the Editor or Admin role.
	RootRole             NullableFloat32        `json:"rootRole,omitempty"`
	CreatedBy            NullableString         `json:"createdBy,omitempty"`
	CreatedAt            NullableTime           `json:"createdAt,omitempty"`
	Users                []GroupUserModelSchema `json:"users,omitempty"`
	Projects             []string               `json:"projects,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GroupSchema GroupSchema

// NewGroupSchema instantiates a new GroupSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupSchema(name string) *GroupSchema {
	this := GroupSchema{}
	this.Name = name
	return &this
}

// NewGroupSchemaWithDefaults instantiates a new GroupSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupSchemaWithDefaults() *GroupSchema {
	this := GroupSchema{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GroupSchema) GetId() float32 {
	if o == nil || IsNil(o.Id) {
		var ret float32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupSchema) GetIdOk() (*float32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GroupSchema) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given float32 and assigns it to the Id field.
func (o *GroupSchema) SetId(v float32) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *GroupSchema) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GroupSchema) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GroupSchema) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupSchema) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupSchema) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *GroupSchema) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *GroupSchema) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *GroupSchema) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *GroupSchema) UnsetDescription() {
	o.Description.Unset()
}

// GetMappingsSSO returns the MappingsSSO field value if set, zero value otherwise.
func (o *GroupSchema) GetMappingsSSO() []string {
	if o == nil || IsNil(o.MappingsSSO) {
		var ret []string
		return ret
	}
	return o.MappingsSSO
}

// GetMappingsSSOOk returns a tuple with the MappingsSSO field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupSchema) GetMappingsSSOOk() ([]string, bool) {
	if o == nil || IsNil(o.MappingsSSO) {
		return nil, false
	}
	return o.MappingsSSO, true
}

// HasMappingsSSO returns a boolean if a field has been set.
func (o *GroupSchema) HasMappingsSSO() bool {
	if o != nil && !IsNil(o.MappingsSSO) {
		return true
	}

	return false
}

// SetMappingsSSO gets a reference to the given []string and assigns it to the MappingsSSO field.
func (o *GroupSchema) SetMappingsSSO(v []string) {
	o.MappingsSSO = v
}

// GetRootRole returns the RootRole field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupSchema) GetRootRole() float32 {
	if o == nil || IsNil(o.RootRole.Get()) {
		var ret float32
		return ret
	}
	return *o.RootRole.Get()
}

// GetRootRoleOk returns a tuple with the RootRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupSchema) GetRootRoleOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.RootRole.Get(), o.RootRole.IsSet()
}

// HasRootRole returns a boolean if a field has been set.
func (o *GroupSchema) HasRootRole() bool {
	if o != nil && o.RootRole.IsSet() {
		return true
	}

	return false
}

// SetRootRole gets a reference to the given NullableFloat32 and assigns it to the RootRole field.
func (o *GroupSchema) SetRootRole(v float32) {
	o.RootRole.Set(&v)
}

// SetRootRoleNil sets the value for RootRole to be an explicit nil
func (o *GroupSchema) SetRootRoleNil() {
	o.RootRole.Set(nil)
}

// UnsetRootRole ensures that no value is present for RootRole, not even an explicit nil
func (o *GroupSchema) UnsetRootRole() {
	o.RootRole.Unset()
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupSchema) GetCreatedBy() string {
	if o == nil || IsNil(o.CreatedBy.Get()) {
		var ret string
		return ret
	}
	return *o.CreatedBy.Get()
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupSchema) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedBy.Get(), o.CreatedBy.IsSet()
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *GroupSchema) HasCreatedBy() bool {
	if o != nil && o.CreatedBy.IsSet() {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given NullableString and assigns it to the CreatedBy field.
func (o *GroupSchema) SetCreatedBy(v string) {
	o.CreatedBy.Set(&v)
}

// SetCreatedByNil sets the value for CreatedBy to be an explicit nil
func (o *GroupSchema) SetCreatedByNil() {
	o.CreatedBy.Set(nil)
}

// UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
func (o *GroupSchema) UnsetCreatedBy() {
	o.CreatedBy.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupSchema) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupSchema) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GroupSchema) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableTime and assigns it to the CreatedAt field.
func (o *GroupSchema) SetCreatedAt(v time.Time) {
	o.CreatedAt.Set(&v)
}

// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *GroupSchema) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *GroupSchema) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *GroupSchema) GetUsers() []GroupUserModelSchema {
	if o == nil || IsNil(o.Users) {
		var ret []GroupUserModelSchema
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupSchema) GetUsersOk() ([]GroupUserModelSchema, bool) {
	if o == nil || IsNil(o.Users) {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *GroupSchema) HasUsers() bool {
	if o != nil && !IsNil(o.Users) {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []GroupUserModelSchema and assigns it to the Users field.
func (o *GroupSchema) SetUsers(v []GroupUserModelSchema) {
	o.Users = v
}

// GetProjects returns the Projects field value if set, zero value otherwise.
func (o *GroupSchema) GetProjects() []string {
	if o == nil || IsNil(o.Projects) {
		var ret []string
		return ret
	}
	return o.Projects
}

// GetProjectsOk returns a tuple with the Projects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupSchema) GetProjectsOk() ([]string, bool) {
	if o == nil || IsNil(o.Projects) {
		return nil, false
	}
	return o.Projects, true
}

// HasProjects returns a boolean if a field has been set.
func (o *GroupSchema) HasProjects() bool {
	if o != nil && !IsNil(o.Projects) {
		return true
	}

	return false
}

// SetProjects gets a reference to the given []string and assigns it to the Projects field.
func (o *GroupSchema) SetProjects(v []string) {
	o.Projects = v
}

func (o GroupSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["name"] = o.Name
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if !IsNil(o.MappingsSSO) {
		toSerialize["mappingsSSO"] = o.MappingsSSO
	}
	if o.RootRole.IsSet() {
		toSerialize["rootRole"] = o.RootRole.Get()
	}
	if o.CreatedBy.IsSet() {
		toSerialize["createdBy"] = o.CreatedBy.Get()
	}
	if o.CreatedAt.IsSet() {
		toSerialize["createdAt"] = o.CreatedAt.Get()
	}
	if !IsNil(o.Users) {
		toSerialize["users"] = o.Users
	}
	if !IsNil(o.Projects) {
		toSerialize["projects"] = o.Projects
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GroupSchema) UnmarshalJSON(bytes []byte) (err error) {
	varGroupSchema := _GroupSchema{}

	if err = json.Unmarshal(bytes, &varGroupSchema); err == nil {
		*o = GroupSchema(varGroupSchema)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "mappingsSSO")
		delete(additionalProperties, "rootRole")
		delete(additionalProperties, "createdBy")
		delete(additionalProperties, "createdAt")
		delete(additionalProperties, "users")
		delete(additionalProperties, "projects")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGroupSchema struct {
	value *GroupSchema
	isSet bool
}

func (v NullableGroupSchema) Get() *GroupSchema {
	return v.value
}

func (v *NullableGroupSchema) Set(val *GroupSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupSchema(val *GroupSchema) *NullableGroupSchema {
	return &NullableGroupSchema{value: val, isSet: true}
}

func (v NullableGroupSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
