/*
Unleash API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 6.1.10+main
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
	"time"
)

// checks if the UserWithProjectRoleSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserWithProjectRoleSchema{}

// UserWithProjectRoleSchema Data about a user including their project role
type UserWithProjectRoleSchema struct {
	// Whether this user is authenticated through Unleash tokens or logged in with a session
	// Deprecated
	IsAPI *bool `json:"isAPI,omitempty"`
	// The name of the user
	Name *string `json:"name,omitempty"`
	// The user's email address
	Email NullableString `json:"email,omitempty"`
	// The user's ID in the Unleash system
	Id int32 `json:"id"`
	// A URL pointing to the user's image.
	ImageUrl NullableString `json:"imageUrl,omitempty"`
	// When this user was added to the project
	AddedAt *time.Time `json:"addedAt,omitempty"`
	// The ID of the role this user has in the given project
	RoleId *int32 `json:"roleId,omitempty"`
	// A list of roles this user has in the given project
	Roles                []int32 `json:"roles,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserWithProjectRoleSchema UserWithProjectRoleSchema

// NewUserWithProjectRoleSchema instantiates a new UserWithProjectRoleSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserWithProjectRoleSchema(id int32) *UserWithProjectRoleSchema {
	this := UserWithProjectRoleSchema{}
	this.Id = id
	return &this
}

// NewUserWithProjectRoleSchemaWithDefaults instantiates a new UserWithProjectRoleSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserWithProjectRoleSchemaWithDefaults() *UserWithProjectRoleSchema {
	this := UserWithProjectRoleSchema{}
	return &this
}

// GetIsAPI returns the IsAPI field value if set, zero value otherwise.
// Deprecated
func (o *UserWithProjectRoleSchema) GetIsAPI() bool {
	if o == nil || IsNil(o.IsAPI) {
		var ret bool
		return ret
	}
	return *o.IsAPI
}

// GetIsAPIOk returns a tuple with the IsAPI field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *UserWithProjectRoleSchema) GetIsAPIOk() (*bool, bool) {
	if o == nil || IsNil(o.IsAPI) {
		return nil, false
	}
	return o.IsAPI, true
}

// HasIsAPI returns a boolean if a field has been set.
func (o *UserWithProjectRoleSchema) HasIsAPI() bool {
	if o != nil && !IsNil(o.IsAPI) {
		return true
	}

	return false
}

// SetIsAPI gets a reference to the given bool and assigns it to the IsAPI field.
// Deprecated
func (o *UserWithProjectRoleSchema) SetIsAPI(v bool) {
	o.IsAPI = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UserWithProjectRoleSchema) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserWithProjectRoleSchema) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UserWithProjectRoleSchema) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UserWithProjectRoleSchema) SetName(v string) {
	o.Name = &v
}

// GetEmail returns the Email field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserWithProjectRoleSchema) GetEmail() string {
	if o == nil || IsNil(o.Email.Get()) {
		var ret string
		return ret
	}
	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserWithProjectRoleSchema) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// HasEmail returns a boolean if a field has been set.
func (o *UserWithProjectRoleSchema) HasEmail() bool {
	if o != nil && o.Email.IsSet() {
		return true
	}

	return false
}

// SetEmail gets a reference to the given NullableString and assigns it to the Email field.
func (o *UserWithProjectRoleSchema) SetEmail(v string) {
	o.Email.Set(&v)
}

// SetEmailNil sets the value for Email to be an explicit nil
func (o *UserWithProjectRoleSchema) SetEmailNil() {
	o.Email.Set(nil)
}

// UnsetEmail ensures that no value is present for Email, not even an explicit nil
func (o *UserWithProjectRoleSchema) UnsetEmail() {
	o.Email.Unset()
}

// GetId returns the Id field value
func (o *UserWithProjectRoleSchema) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UserWithProjectRoleSchema) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UserWithProjectRoleSchema) SetId(v int32) {
	o.Id = v
}

// GetImageUrl returns the ImageUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserWithProjectRoleSchema) GetImageUrl() string {
	if o == nil || IsNil(o.ImageUrl.Get()) {
		var ret string
		return ret
	}
	return *o.ImageUrl.Get()
}

// GetImageUrlOk returns a tuple with the ImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserWithProjectRoleSchema) GetImageUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ImageUrl.Get(), o.ImageUrl.IsSet()
}

// HasImageUrl returns a boolean if a field has been set.
func (o *UserWithProjectRoleSchema) HasImageUrl() bool {
	if o != nil && o.ImageUrl.IsSet() {
		return true
	}

	return false
}

// SetImageUrl gets a reference to the given NullableString and assigns it to the ImageUrl field.
func (o *UserWithProjectRoleSchema) SetImageUrl(v string) {
	o.ImageUrl.Set(&v)
}

// SetImageUrlNil sets the value for ImageUrl to be an explicit nil
func (o *UserWithProjectRoleSchema) SetImageUrlNil() {
	o.ImageUrl.Set(nil)
}

// UnsetImageUrl ensures that no value is present for ImageUrl, not even an explicit nil
func (o *UserWithProjectRoleSchema) UnsetImageUrl() {
	o.ImageUrl.Unset()
}

// GetAddedAt returns the AddedAt field value if set, zero value otherwise.
func (o *UserWithProjectRoleSchema) GetAddedAt() time.Time {
	if o == nil || IsNil(o.AddedAt) {
		var ret time.Time
		return ret
	}
	return *o.AddedAt
}

// GetAddedAtOk returns a tuple with the AddedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserWithProjectRoleSchema) GetAddedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.AddedAt) {
		return nil, false
	}
	return o.AddedAt, true
}

// HasAddedAt returns a boolean if a field has been set.
func (o *UserWithProjectRoleSchema) HasAddedAt() bool {
	if o != nil && !IsNil(o.AddedAt) {
		return true
	}

	return false
}

// SetAddedAt gets a reference to the given time.Time and assigns it to the AddedAt field.
func (o *UserWithProjectRoleSchema) SetAddedAt(v time.Time) {
	o.AddedAt = &v
}

// GetRoleId returns the RoleId field value if set, zero value otherwise.
func (o *UserWithProjectRoleSchema) GetRoleId() int32 {
	if o == nil || IsNil(o.RoleId) {
		var ret int32
		return ret
	}
	return *o.RoleId
}

// GetRoleIdOk returns a tuple with the RoleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserWithProjectRoleSchema) GetRoleIdOk() (*int32, bool) {
	if o == nil || IsNil(o.RoleId) {
		return nil, false
	}
	return o.RoleId, true
}

// HasRoleId returns a boolean if a field has been set.
func (o *UserWithProjectRoleSchema) HasRoleId() bool {
	if o != nil && !IsNil(o.RoleId) {
		return true
	}

	return false
}

// SetRoleId gets a reference to the given int32 and assigns it to the RoleId field.
func (o *UserWithProjectRoleSchema) SetRoleId(v int32) {
	o.RoleId = &v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *UserWithProjectRoleSchema) GetRoles() []int32 {
	if o == nil || IsNil(o.Roles) {
		var ret []int32
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserWithProjectRoleSchema) GetRolesOk() ([]int32, bool) {
	if o == nil || IsNil(o.Roles) {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *UserWithProjectRoleSchema) HasRoles() bool {
	if o != nil && !IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []int32 and assigns it to the Roles field.
func (o *UserWithProjectRoleSchema) SetRoles(v []int32) {
	o.Roles = v
}

func (o UserWithProjectRoleSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserWithProjectRoleSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IsAPI) {
		toSerialize["isAPI"] = o.IsAPI
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if o.Email.IsSet() {
		toSerialize["email"] = o.Email.Get()
	}
	toSerialize["id"] = o.Id
	if o.ImageUrl.IsSet() {
		toSerialize["imageUrl"] = o.ImageUrl.Get()
	}
	if !IsNil(o.AddedAt) {
		toSerialize["addedAt"] = o.AddedAt
	}
	if !IsNil(o.RoleId) {
		toSerialize["roleId"] = o.RoleId
	}
	if !IsNil(o.Roles) {
		toSerialize["roles"] = o.Roles
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UserWithProjectRoleSchema) UnmarshalJSON(data []byte) (err error) {
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

	varUserWithProjectRoleSchema := _UserWithProjectRoleSchema{}

	err = json.Unmarshal(data, &varUserWithProjectRoleSchema)

	if err != nil {
		return err
	}

	*o = UserWithProjectRoleSchema(varUserWithProjectRoleSchema)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "isAPI")
		delete(additionalProperties, "name")
		delete(additionalProperties, "email")
		delete(additionalProperties, "id")
		delete(additionalProperties, "imageUrl")
		delete(additionalProperties, "addedAt")
		delete(additionalProperties, "roleId")
		delete(additionalProperties, "roles")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserWithProjectRoleSchema struct {
	value *UserWithProjectRoleSchema
	isSet bool
}

func (v NullableUserWithProjectRoleSchema) Get() *UserWithProjectRoleSchema {
	return v.value
}

func (v *NullableUserWithProjectRoleSchema) Set(val *UserWithProjectRoleSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableUserWithProjectRoleSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableUserWithProjectRoleSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserWithProjectRoleSchema(val *UserWithProjectRoleSchema) *NullableUserWithProjectRoleSchema {
	return &NullableUserWithProjectRoleSchema{value: val, isSet: true}
}

func (v NullableUserWithProjectRoleSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserWithProjectRoleSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
