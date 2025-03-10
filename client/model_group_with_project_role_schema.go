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
	"time"
)

// checks if the GroupWithProjectRoleSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupWithProjectRoleSchema{}

// GroupWithProjectRoleSchema Data about a group including their project role
type GroupWithProjectRoleSchema struct {
	// The name of the group
	Name *string `json:"name,omitempty"`
	// The group's ID in the Unleash system
	Id int32 `json:"id"`
	// When this group was added to the project
	AddedAt *time.Time `json:"addedAt,omitempty"`
	// The ID of the role this group has in the given project
	RoleId *int32 `json:"roleId,omitempty"`
	// A list of roles this user has in the given project
	Roles []int32 `json:"roles,omitempty"`
	// A custom description of the group
	Description NullableString `json:"description,omitempty"`
	// A list of SSO groups that should map to this Unleash group
	MappingsSSO []string `json:"mappingsSSO,omitempty"`
	// A role id that is used as the root role for all users in this group. This can be either the id of the Viewer, Editor or Admin role.
	RootRole NullableFloat32 `json:"rootRole,omitempty"`
	// A user who created this group
	CreatedBy NullableString `json:"createdBy,omitempty"`
	// When was this group created
	CreatedAt NullableTime `json:"createdAt,omitempty"`
	// A list of users belonging to this group
	Users []GroupUserModelSchema `json:"users,omitempty"`
	// The SCIM ID of the group, only present if managed by SCIM
	ScimId               NullableString `json:"scimId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GroupWithProjectRoleSchema GroupWithProjectRoleSchema

// NewGroupWithProjectRoleSchema instantiates a new GroupWithProjectRoleSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupWithProjectRoleSchema(id int32) *GroupWithProjectRoleSchema {
	this := GroupWithProjectRoleSchema{}
	this.Id = id
	return &this
}

// NewGroupWithProjectRoleSchemaWithDefaults instantiates a new GroupWithProjectRoleSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupWithProjectRoleSchemaWithDefaults() *GroupWithProjectRoleSchema {
	this := GroupWithProjectRoleSchema{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GroupWithProjectRoleSchema) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupWithProjectRoleSchema) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GroupWithProjectRoleSchema) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GroupWithProjectRoleSchema) SetName(v string) {
	o.Name = &v
}

// GetId returns the Id field value
func (o *GroupWithProjectRoleSchema) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GroupWithProjectRoleSchema) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GroupWithProjectRoleSchema) SetId(v int32) {
	o.Id = v
}

// GetAddedAt returns the AddedAt field value if set, zero value otherwise.
func (o *GroupWithProjectRoleSchema) GetAddedAt() time.Time {
	if o == nil || IsNil(o.AddedAt) {
		var ret time.Time
		return ret
	}
	return *o.AddedAt
}

// GetAddedAtOk returns a tuple with the AddedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupWithProjectRoleSchema) GetAddedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.AddedAt) {
		return nil, false
	}
	return o.AddedAt, true
}

// HasAddedAt returns a boolean if a field has been set.
func (o *GroupWithProjectRoleSchema) HasAddedAt() bool {
	if o != nil && !IsNil(o.AddedAt) {
		return true
	}

	return false
}

// SetAddedAt gets a reference to the given time.Time and assigns it to the AddedAt field.
func (o *GroupWithProjectRoleSchema) SetAddedAt(v time.Time) {
	o.AddedAt = &v
}

// GetRoleId returns the RoleId field value if set, zero value otherwise.
func (o *GroupWithProjectRoleSchema) GetRoleId() int32 {
	if o == nil || IsNil(o.RoleId) {
		var ret int32
		return ret
	}
	return *o.RoleId
}

// GetRoleIdOk returns a tuple with the RoleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupWithProjectRoleSchema) GetRoleIdOk() (*int32, bool) {
	if o == nil || IsNil(o.RoleId) {
		return nil, false
	}
	return o.RoleId, true
}

// HasRoleId returns a boolean if a field has been set.
func (o *GroupWithProjectRoleSchema) HasRoleId() bool {
	if o != nil && !IsNil(o.RoleId) {
		return true
	}

	return false
}

// SetRoleId gets a reference to the given int32 and assigns it to the RoleId field.
func (o *GroupWithProjectRoleSchema) SetRoleId(v int32) {
	o.RoleId = &v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *GroupWithProjectRoleSchema) GetRoles() []int32 {
	if o == nil || IsNil(o.Roles) {
		var ret []int32
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupWithProjectRoleSchema) GetRolesOk() ([]int32, bool) {
	if o == nil || IsNil(o.Roles) {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *GroupWithProjectRoleSchema) HasRoles() bool {
	if o != nil && !IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []int32 and assigns it to the Roles field.
func (o *GroupWithProjectRoleSchema) SetRoles(v []int32) {
	o.Roles = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupWithProjectRoleSchema) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupWithProjectRoleSchema) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *GroupWithProjectRoleSchema) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *GroupWithProjectRoleSchema) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *GroupWithProjectRoleSchema) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *GroupWithProjectRoleSchema) UnsetDescription() {
	o.Description.Unset()
}

// GetMappingsSSO returns the MappingsSSO field value if set, zero value otherwise.
func (o *GroupWithProjectRoleSchema) GetMappingsSSO() []string {
	if o == nil || IsNil(o.MappingsSSO) {
		var ret []string
		return ret
	}
	return o.MappingsSSO
}

// GetMappingsSSOOk returns a tuple with the MappingsSSO field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupWithProjectRoleSchema) GetMappingsSSOOk() ([]string, bool) {
	if o == nil || IsNil(o.MappingsSSO) {
		return nil, false
	}
	return o.MappingsSSO, true
}

// HasMappingsSSO returns a boolean if a field has been set.
func (o *GroupWithProjectRoleSchema) HasMappingsSSO() bool {
	if o != nil && !IsNil(o.MappingsSSO) {
		return true
	}

	return false
}

// SetMappingsSSO gets a reference to the given []string and assigns it to the MappingsSSO field.
func (o *GroupWithProjectRoleSchema) SetMappingsSSO(v []string) {
	o.MappingsSSO = v
}

// GetRootRole returns the RootRole field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupWithProjectRoleSchema) GetRootRole() float32 {
	if o == nil || IsNil(o.RootRole.Get()) {
		var ret float32
		return ret
	}
	return *o.RootRole.Get()
}

// GetRootRoleOk returns a tuple with the RootRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupWithProjectRoleSchema) GetRootRoleOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.RootRole.Get(), o.RootRole.IsSet()
}

// HasRootRole returns a boolean if a field has been set.
func (o *GroupWithProjectRoleSchema) HasRootRole() bool {
	if o != nil && o.RootRole.IsSet() {
		return true
	}

	return false
}

// SetRootRole gets a reference to the given NullableFloat32 and assigns it to the RootRole field.
func (o *GroupWithProjectRoleSchema) SetRootRole(v float32) {
	o.RootRole.Set(&v)
}

// SetRootRoleNil sets the value for RootRole to be an explicit nil
func (o *GroupWithProjectRoleSchema) SetRootRoleNil() {
	o.RootRole.Set(nil)
}

// UnsetRootRole ensures that no value is present for RootRole, not even an explicit nil
func (o *GroupWithProjectRoleSchema) UnsetRootRole() {
	o.RootRole.Unset()
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupWithProjectRoleSchema) GetCreatedBy() string {
	if o == nil || IsNil(o.CreatedBy.Get()) {
		var ret string
		return ret
	}
	return *o.CreatedBy.Get()
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupWithProjectRoleSchema) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedBy.Get(), o.CreatedBy.IsSet()
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *GroupWithProjectRoleSchema) HasCreatedBy() bool {
	if o != nil && o.CreatedBy.IsSet() {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given NullableString and assigns it to the CreatedBy field.
func (o *GroupWithProjectRoleSchema) SetCreatedBy(v string) {
	o.CreatedBy.Set(&v)
}

// SetCreatedByNil sets the value for CreatedBy to be an explicit nil
func (o *GroupWithProjectRoleSchema) SetCreatedByNil() {
	o.CreatedBy.Set(nil)
}

// UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
func (o *GroupWithProjectRoleSchema) UnsetCreatedBy() {
	o.CreatedBy.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupWithProjectRoleSchema) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupWithProjectRoleSchema) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GroupWithProjectRoleSchema) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableTime and assigns it to the CreatedAt field.
func (o *GroupWithProjectRoleSchema) SetCreatedAt(v time.Time) {
	o.CreatedAt.Set(&v)
}

// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *GroupWithProjectRoleSchema) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *GroupWithProjectRoleSchema) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *GroupWithProjectRoleSchema) GetUsers() []GroupUserModelSchema {
	if o == nil || IsNil(o.Users) {
		var ret []GroupUserModelSchema
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupWithProjectRoleSchema) GetUsersOk() ([]GroupUserModelSchema, bool) {
	if o == nil || IsNil(o.Users) {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *GroupWithProjectRoleSchema) HasUsers() bool {
	if o != nil && !IsNil(o.Users) {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []GroupUserModelSchema and assigns it to the Users field.
func (o *GroupWithProjectRoleSchema) SetUsers(v []GroupUserModelSchema) {
	o.Users = v
}

// GetScimId returns the ScimId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupWithProjectRoleSchema) GetScimId() string {
	if o == nil || IsNil(o.ScimId.Get()) {
		var ret string
		return ret
	}
	return *o.ScimId.Get()
}

// GetScimIdOk returns a tuple with the ScimId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupWithProjectRoleSchema) GetScimIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ScimId.Get(), o.ScimId.IsSet()
}

// HasScimId returns a boolean if a field has been set.
func (o *GroupWithProjectRoleSchema) HasScimId() bool {
	if o != nil && o.ScimId.IsSet() {
		return true
	}

	return false
}

// SetScimId gets a reference to the given NullableString and assigns it to the ScimId field.
func (o *GroupWithProjectRoleSchema) SetScimId(v string) {
	o.ScimId.Set(&v)
}

// SetScimIdNil sets the value for ScimId to be an explicit nil
func (o *GroupWithProjectRoleSchema) SetScimIdNil() {
	o.ScimId.Set(nil)
}

// UnsetScimId ensures that no value is present for ScimId, not even an explicit nil
func (o *GroupWithProjectRoleSchema) UnsetScimId() {
	o.ScimId.Unset()
}

func (o GroupWithProjectRoleSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupWithProjectRoleSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	toSerialize["id"] = o.Id
	if !IsNil(o.AddedAt) {
		toSerialize["addedAt"] = o.AddedAt
	}
	if !IsNil(o.RoleId) {
		toSerialize["roleId"] = o.RoleId
	}
	if !IsNil(o.Roles) {
		toSerialize["roles"] = o.Roles
	}
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
	if o.ScimId.IsSet() {
		toSerialize["scimId"] = o.ScimId.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GroupWithProjectRoleSchema) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
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

	varGroupWithProjectRoleSchema := _GroupWithProjectRoleSchema{}

	err = json.Unmarshal(data, &varGroupWithProjectRoleSchema)

	if err != nil {
		return err
	}

	*o = GroupWithProjectRoleSchema(varGroupWithProjectRoleSchema)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "id")
		delete(additionalProperties, "addedAt")
		delete(additionalProperties, "roleId")
		delete(additionalProperties, "roles")
		delete(additionalProperties, "description")
		delete(additionalProperties, "mappingsSSO")
		delete(additionalProperties, "rootRole")
		delete(additionalProperties, "createdBy")
		delete(additionalProperties, "createdAt")
		delete(additionalProperties, "users")
		delete(additionalProperties, "scimId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGroupWithProjectRoleSchema struct {
	value *GroupWithProjectRoleSchema
	isSet bool
}

func (v NullableGroupWithProjectRoleSchema) Get() *GroupWithProjectRoleSchema {
	return v.value
}

func (v *NullableGroupWithProjectRoleSchema) Set(val *GroupWithProjectRoleSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupWithProjectRoleSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupWithProjectRoleSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupWithProjectRoleSchema(val *GroupWithProjectRoleSchema) *NullableGroupWithProjectRoleSchema {
	return &NullableGroupWithProjectRoleSchema{value: val, isSet: true}
}

func (v NullableGroupWithProjectRoleSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupWithProjectRoleSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
