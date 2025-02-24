/*
Unleash API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 6.7.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
	"time"
)

// checks if the CreateUserResponseSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateUserResponseSchema{}

// CreateUserResponseSchema An Unleash user after creation
type CreateUserResponseSchema struct {
	// The user id
	Id int32 `json:"id"`
	// Deprecated in v5. Used internally to know which operations the user should be allowed to perform
	// Deprecated
	IsAPI *bool `json:"isAPI,omitempty"`
	// Name of the user
	Name NullableString `json:"name,omitempty"`
	// Email of the user
	Email *string `json:"email,omitempty"`
	// A unique username for the user
	Username NullableString `json:"username,omitempty"`
	// URL used for the user profile image
	ImageUrl *string `json:"imageUrl,omitempty"`
	// If the user is actively inviting other users, this is the link that can be shared with other users
	InviteLink *string `json:"inviteLink,omitempty"`
	// How many unsuccessful attempts at logging in has the user made
	LoginAttempts *int32 `json:"loginAttempts,omitempty"`
	// Is the welcome email sent to the user or not
	EmailSent *bool                             `json:"emailSent,omitempty"`
	RootRole  *CreateUserResponseSchemaRootRole `json:"rootRole,omitempty"`
	// The last time this user logged in
	SeenAt NullableTime `json:"seenAt,omitempty"`
	// The user was created at this time
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// A user is either an actual User or a Service Account
	AccountType *string `json:"accountType,omitempty"`
	// Deprecated
	Permissions []string `json:"permissions,omitempty"`
	// The SCIM ID of the user, only present if managed by SCIM
	ScimId NullableString `json:"scimId,omitempty"`
	// Count of active browser sessions for this user
	ActiveSessions NullableInt32 `json:"activeSessions,omitempty"`
	// Experimental. The number of deleted browser sessions after last login
	DeletedSessions      *float32 `json:"deletedSessions,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateUserResponseSchema CreateUserResponseSchema

// NewCreateUserResponseSchema instantiates a new CreateUserResponseSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateUserResponseSchema(id int32) *CreateUserResponseSchema {
	this := CreateUserResponseSchema{}
	this.Id = id
	return &this
}

// NewCreateUserResponseSchemaWithDefaults instantiates a new CreateUserResponseSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateUserResponseSchemaWithDefaults() *CreateUserResponseSchema {
	this := CreateUserResponseSchema{}
	return &this
}

// GetId returns the Id field value
func (o *CreateUserResponseSchema) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CreateUserResponseSchema) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CreateUserResponseSchema) SetId(v int32) {
	o.Id = v
}

// GetIsAPI returns the IsAPI field value if set, zero value otherwise.
// Deprecated
func (o *CreateUserResponseSchema) GetIsAPI() bool {
	if o == nil || IsNil(o.IsAPI) {
		var ret bool
		return ret
	}
	return *o.IsAPI
}

// GetIsAPIOk returns a tuple with the IsAPI field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *CreateUserResponseSchema) GetIsAPIOk() (*bool, bool) {
	if o == nil || IsNil(o.IsAPI) {
		return nil, false
	}
	return o.IsAPI, true
}

// HasIsAPI returns a boolean if a field has been set.
func (o *CreateUserResponseSchema) HasIsAPI() bool {
	if o != nil && !IsNil(o.IsAPI) {
		return true
	}

	return false
}

// SetIsAPI gets a reference to the given bool and assigns it to the IsAPI field.
// Deprecated
func (o *CreateUserResponseSchema) SetIsAPI(v bool) {
	o.IsAPI = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateUserResponseSchema) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateUserResponseSchema) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *CreateUserResponseSchema) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *CreateUserResponseSchema) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *CreateUserResponseSchema) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *CreateUserResponseSchema) UnsetName() {
	o.Name.Unset()
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *CreateUserResponseSchema) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUserResponseSchema) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *CreateUserResponseSchema) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *CreateUserResponseSchema) SetEmail(v string) {
	o.Email = &v
}

// GetUsername returns the Username field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateUserResponseSchema) GetUsername() string {
	if o == nil || IsNil(o.Username.Get()) {
		var ret string
		return ret
	}
	return *o.Username.Get()
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateUserResponseSchema) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Username.Get(), o.Username.IsSet()
}

// HasUsername returns a boolean if a field has been set.
func (o *CreateUserResponseSchema) HasUsername() bool {
	if o != nil && o.Username.IsSet() {
		return true
	}

	return false
}

// SetUsername gets a reference to the given NullableString and assigns it to the Username field.
func (o *CreateUserResponseSchema) SetUsername(v string) {
	o.Username.Set(&v)
}

// SetUsernameNil sets the value for Username to be an explicit nil
func (o *CreateUserResponseSchema) SetUsernameNil() {
	o.Username.Set(nil)
}

// UnsetUsername ensures that no value is present for Username, not even an explicit nil
func (o *CreateUserResponseSchema) UnsetUsername() {
	o.Username.Unset()
}

// GetImageUrl returns the ImageUrl field value if set, zero value otherwise.
func (o *CreateUserResponseSchema) GetImageUrl() string {
	if o == nil || IsNil(o.ImageUrl) {
		var ret string
		return ret
	}
	return *o.ImageUrl
}

// GetImageUrlOk returns a tuple with the ImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUserResponseSchema) GetImageUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ImageUrl) {
		return nil, false
	}
	return o.ImageUrl, true
}

// HasImageUrl returns a boolean if a field has been set.
func (o *CreateUserResponseSchema) HasImageUrl() bool {
	if o != nil && !IsNil(o.ImageUrl) {
		return true
	}

	return false
}

// SetImageUrl gets a reference to the given string and assigns it to the ImageUrl field.
func (o *CreateUserResponseSchema) SetImageUrl(v string) {
	o.ImageUrl = &v
}

// GetInviteLink returns the InviteLink field value if set, zero value otherwise.
func (o *CreateUserResponseSchema) GetInviteLink() string {
	if o == nil || IsNil(o.InviteLink) {
		var ret string
		return ret
	}
	return *o.InviteLink
}

// GetInviteLinkOk returns a tuple with the InviteLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUserResponseSchema) GetInviteLinkOk() (*string, bool) {
	if o == nil || IsNil(o.InviteLink) {
		return nil, false
	}
	return o.InviteLink, true
}

// HasInviteLink returns a boolean if a field has been set.
func (o *CreateUserResponseSchema) HasInviteLink() bool {
	if o != nil && !IsNil(o.InviteLink) {
		return true
	}

	return false
}

// SetInviteLink gets a reference to the given string and assigns it to the InviteLink field.
func (o *CreateUserResponseSchema) SetInviteLink(v string) {
	o.InviteLink = &v
}

// GetLoginAttempts returns the LoginAttempts field value if set, zero value otherwise.
func (o *CreateUserResponseSchema) GetLoginAttempts() int32 {
	if o == nil || IsNil(o.LoginAttempts) {
		var ret int32
		return ret
	}
	return *o.LoginAttempts
}

// GetLoginAttemptsOk returns a tuple with the LoginAttempts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUserResponseSchema) GetLoginAttemptsOk() (*int32, bool) {
	if o == nil || IsNil(o.LoginAttempts) {
		return nil, false
	}
	return o.LoginAttempts, true
}

// HasLoginAttempts returns a boolean if a field has been set.
func (o *CreateUserResponseSchema) HasLoginAttempts() bool {
	if o != nil && !IsNil(o.LoginAttempts) {
		return true
	}

	return false
}

// SetLoginAttempts gets a reference to the given int32 and assigns it to the LoginAttempts field.
func (o *CreateUserResponseSchema) SetLoginAttempts(v int32) {
	o.LoginAttempts = &v
}

// GetEmailSent returns the EmailSent field value if set, zero value otherwise.
func (o *CreateUserResponseSchema) GetEmailSent() bool {
	if o == nil || IsNil(o.EmailSent) {
		var ret bool
		return ret
	}
	return *o.EmailSent
}

// GetEmailSentOk returns a tuple with the EmailSent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUserResponseSchema) GetEmailSentOk() (*bool, bool) {
	if o == nil || IsNil(o.EmailSent) {
		return nil, false
	}
	return o.EmailSent, true
}

// HasEmailSent returns a boolean if a field has been set.
func (o *CreateUserResponseSchema) HasEmailSent() bool {
	if o != nil && !IsNil(o.EmailSent) {
		return true
	}

	return false
}

// SetEmailSent gets a reference to the given bool and assigns it to the EmailSent field.
func (o *CreateUserResponseSchema) SetEmailSent(v bool) {
	o.EmailSent = &v
}

// GetRootRole returns the RootRole field value if set, zero value otherwise.
func (o *CreateUserResponseSchema) GetRootRole() CreateUserResponseSchemaRootRole {
	if o == nil || IsNil(o.RootRole) {
		var ret CreateUserResponseSchemaRootRole
		return ret
	}
	return *o.RootRole
}

// GetRootRoleOk returns a tuple with the RootRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUserResponseSchema) GetRootRoleOk() (*CreateUserResponseSchemaRootRole, bool) {
	if o == nil || IsNil(o.RootRole) {
		return nil, false
	}
	return o.RootRole, true
}

// HasRootRole returns a boolean if a field has been set.
func (o *CreateUserResponseSchema) HasRootRole() bool {
	if o != nil && !IsNil(o.RootRole) {
		return true
	}

	return false
}

// SetRootRole gets a reference to the given CreateUserResponseSchemaRootRole and assigns it to the RootRole field.
func (o *CreateUserResponseSchema) SetRootRole(v CreateUserResponseSchemaRootRole) {
	o.RootRole = &v
}

// GetSeenAt returns the SeenAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateUserResponseSchema) GetSeenAt() time.Time {
	if o == nil || IsNil(o.SeenAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.SeenAt.Get()
}

// GetSeenAtOk returns a tuple with the SeenAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateUserResponseSchema) GetSeenAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.SeenAt.Get(), o.SeenAt.IsSet()
}

// HasSeenAt returns a boolean if a field has been set.
func (o *CreateUserResponseSchema) HasSeenAt() bool {
	if o != nil && o.SeenAt.IsSet() {
		return true
	}

	return false
}

// SetSeenAt gets a reference to the given NullableTime and assigns it to the SeenAt field.
func (o *CreateUserResponseSchema) SetSeenAt(v time.Time) {
	o.SeenAt.Set(&v)
}

// SetSeenAtNil sets the value for SeenAt to be an explicit nil
func (o *CreateUserResponseSchema) SetSeenAtNil() {
	o.SeenAt.Set(nil)
}

// UnsetSeenAt ensures that no value is present for SeenAt, not even an explicit nil
func (o *CreateUserResponseSchema) UnsetSeenAt() {
	o.SeenAt.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *CreateUserResponseSchema) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUserResponseSchema) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *CreateUserResponseSchema) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *CreateUserResponseSchema) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetAccountType returns the AccountType field value if set, zero value otherwise.
func (o *CreateUserResponseSchema) GetAccountType() string {
	if o == nil || IsNil(o.AccountType) {
		var ret string
		return ret
	}
	return *o.AccountType
}

// GetAccountTypeOk returns a tuple with the AccountType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUserResponseSchema) GetAccountTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AccountType) {
		return nil, false
	}
	return o.AccountType, true
}

// HasAccountType returns a boolean if a field has been set.
func (o *CreateUserResponseSchema) HasAccountType() bool {
	if o != nil && !IsNil(o.AccountType) {
		return true
	}

	return false
}

// SetAccountType gets a reference to the given string and assigns it to the AccountType field.
func (o *CreateUserResponseSchema) SetAccountType(v string) {
	o.AccountType = &v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *CreateUserResponseSchema) GetPermissions() []string {
	if o == nil || IsNil(o.Permissions) {
		var ret []string
		return ret
	}
	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUserResponseSchema) GetPermissionsOk() ([]string, bool) {
	if o == nil || IsNil(o.Permissions) {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *CreateUserResponseSchema) HasPermissions() bool {
	if o != nil && !IsNil(o.Permissions) {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given []string and assigns it to the Permissions field.
func (o *CreateUserResponseSchema) SetPermissions(v []string) {
	o.Permissions = v
}

// GetScimId returns the ScimId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateUserResponseSchema) GetScimId() string {
	if o == nil || IsNil(o.ScimId.Get()) {
		var ret string
		return ret
	}
	return *o.ScimId.Get()
}

// GetScimIdOk returns a tuple with the ScimId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateUserResponseSchema) GetScimIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ScimId.Get(), o.ScimId.IsSet()
}

// HasScimId returns a boolean if a field has been set.
func (o *CreateUserResponseSchema) HasScimId() bool {
	if o != nil && o.ScimId.IsSet() {
		return true
	}

	return false
}

// SetScimId gets a reference to the given NullableString and assigns it to the ScimId field.
func (o *CreateUserResponseSchema) SetScimId(v string) {
	o.ScimId.Set(&v)
}

// SetScimIdNil sets the value for ScimId to be an explicit nil
func (o *CreateUserResponseSchema) SetScimIdNil() {
	o.ScimId.Set(nil)
}

// UnsetScimId ensures that no value is present for ScimId, not even an explicit nil
func (o *CreateUserResponseSchema) UnsetScimId() {
	o.ScimId.Unset()
}

// GetActiveSessions returns the ActiveSessions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateUserResponseSchema) GetActiveSessions() int32 {
	if o == nil || IsNil(o.ActiveSessions.Get()) {
		var ret int32
		return ret
	}
	return *o.ActiveSessions.Get()
}

// GetActiveSessionsOk returns a tuple with the ActiveSessions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateUserResponseSchema) GetActiveSessionsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ActiveSessions.Get(), o.ActiveSessions.IsSet()
}

// HasActiveSessions returns a boolean if a field has been set.
func (o *CreateUserResponseSchema) HasActiveSessions() bool {
	if o != nil && o.ActiveSessions.IsSet() {
		return true
	}

	return false
}

// SetActiveSessions gets a reference to the given NullableInt32 and assigns it to the ActiveSessions field.
func (o *CreateUserResponseSchema) SetActiveSessions(v int32) {
	o.ActiveSessions.Set(&v)
}

// SetActiveSessionsNil sets the value for ActiveSessions to be an explicit nil
func (o *CreateUserResponseSchema) SetActiveSessionsNil() {
	o.ActiveSessions.Set(nil)
}

// UnsetActiveSessions ensures that no value is present for ActiveSessions, not even an explicit nil
func (o *CreateUserResponseSchema) UnsetActiveSessions() {
	o.ActiveSessions.Unset()
}

// GetDeletedSessions returns the DeletedSessions field value if set, zero value otherwise.
func (o *CreateUserResponseSchema) GetDeletedSessions() float32 {
	if o == nil || IsNil(o.DeletedSessions) {
		var ret float32
		return ret
	}
	return *o.DeletedSessions
}

// GetDeletedSessionsOk returns a tuple with the DeletedSessions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUserResponseSchema) GetDeletedSessionsOk() (*float32, bool) {
	if o == nil || IsNil(o.DeletedSessions) {
		return nil, false
	}
	return o.DeletedSessions, true
}

// HasDeletedSessions returns a boolean if a field has been set.
func (o *CreateUserResponseSchema) HasDeletedSessions() bool {
	if o != nil && !IsNil(o.DeletedSessions) {
		return true
	}

	return false
}

// SetDeletedSessions gets a reference to the given float32 and assigns it to the DeletedSessions field.
func (o *CreateUserResponseSchema) SetDeletedSessions(v float32) {
	o.DeletedSessions = &v
}

func (o CreateUserResponseSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateUserResponseSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.IsAPI) {
		toSerialize["isAPI"] = o.IsAPI
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if o.Username.IsSet() {
		toSerialize["username"] = o.Username.Get()
	}
	if !IsNil(o.ImageUrl) {
		toSerialize["imageUrl"] = o.ImageUrl
	}
	if !IsNil(o.InviteLink) {
		toSerialize["inviteLink"] = o.InviteLink
	}
	if !IsNil(o.LoginAttempts) {
		toSerialize["loginAttempts"] = o.LoginAttempts
	}
	if !IsNil(o.EmailSent) {
		toSerialize["emailSent"] = o.EmailSent
	}
	if !IsNil(o.RootRole) {
		toSerialize["rootRole"] = o.RootRole
	}
	if o.SeenAt.IsSet() {
		toSerialize["seenAt"] = o.SeenAt.Get()
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.AccountType) {
		toSerialize["accountType"] = o.AccountType
	}
	if !IsNil(o.Permissions) {
		toSerialize["permissions"] = o.Permissions
	}
	if o.ScimId.IsSet() {
		toSerialize["scimId"] = o.ScimId.Get()
	}
	if o.ActiveSessions.IsSet() {
		toSerialize["activeSessions"] = o.ActiveSessions.Get()
	}
	if !IsNil(o.DeletedSessions) {
		toSerialize["deletedSessions"] = o.DeletedSessions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateUserResponseSchema) UnmarshalJSON(data []byte) (err error) {
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

	varCreateUserResponseSchema := _CreateUserResponseSchema{}

	err = json.Unmarshal(data, &varCreateUserResponseSchema)

	if err != nil {
		return err
	}

	*o = CreateUserResponseSchema(varCreateUserResponseSchema)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "isAPI")
		delete(additionalProperties, "name")
		delete(additionalProperties, "email")
		delete(additionalProperties, "username")
		delete(additionalProperties, "imageUrl")
		delete(additionalProperties, "inviteLink")
		delete(additionalProperties, "loginAttempts")
		delete(additionalProperties, "emailSent")
		delete(additionalProperties, "rootRole")
		delete(additionalProperties, "seenAt")
		delete(additionalProperties, "createdAt")
		delete(additionalProperties, "accountType")
		delete(additionalProperties, "permissions")
		delete(additionalProperties, "scimId")
		delete(additionalProperties, "activeSessions")
		delete(additionalProperties, "deletedSessions")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateUserResponseSchema struct {
	value *CreateUserResponseSchema
	isSet bool
}

func (v NullableCreateUserResponseSchema) Get() *CreateUserResponseSchema {
	return v.value
}

func (v *NullableCreateUserResponseSchema) Set(val *CreateUserResponseSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateUserResponseSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateUserResponseSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateUserResponseSchema(val *CreateUserResponseSchema) *NullableCreateUserResponseSchema {
	return &NullableCreateUserResponseSchema{value: val, isSet: true}
}

func (v NullableCreateUserResponseSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateUserResponseSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
