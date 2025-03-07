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

// checks if the ServiceAccountSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceAccountSchema{}

// ServiceAccountSchema Represents a [service account](https://docs.getunleash.io/reference/service-accounts). Service accounts are used to let systems interact with the Unleash API.
type ServiceAccountSchema struct {
	// The service account id
	Id float32 `json:"id"`
	// Deprecated: for internal use only, should not be exposed through the API
	// Deprecated
	IsAPI *bool `json:"isAPI,omitempty"`
	// The name of the service account
	Name *string `json:"name,omitempty"`
	// Deprecated: service accounts don't have emails associated with them
	// Deprecated
	Email *string `json:"email,omitempty"`
	// The service account username
	Username *string `json:"username,omitempty"`
	// The service account image url
	ImageUrl *string `json:"imageUrl,omitempty"`
	// Deprecated: service accounts cannot be invited via an invitation link
	// Deprecated
	InviteLink *string `json:"inviteLink,omitempty"`
	// Deprecated: service accounts cannot log in to Unleash
	// Deprecated
	LoginAttempts *float32 `json:"loginAttempts,omitempty"`
	// Deprecated: internal use only
	// Deprecated
	EmailSent *bool `json:"emailSent,omitempty"`
	// The root role id associated with the service account
	RootRole *int32 `json:"rootRole,omitempty"`
	// Deprecated. This property is always `null`. To find out when a service account was last seen, check its `tokens` list and refer to each token's `lastSeen` property instead.
	// Deprecated
	SeenAt NullableTime `json:"seenAt,omitempty"`
	// The service account creation date
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// The list of tokens associated with the service account
	Tokens               []PatSchema `json:"tokens,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ServiceAccountSchema ServiceAccountSchema

// NewServiceAccountSchema instantiates a new ServiceAccountSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceAccountSchema(id float32) *ServiceAccountSchema {
	this := ServiceAccountSchema{}
	this.Id = id
	return &this
}

// NewServiceAccountSchemaWithDefaults instantiates a new ServiceAccountSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceAccountSchemaWithDefaults() *ServiceAccountSchema {
	this := ServiceAccountSchema{}
	return &this
}

// GetId returns the Id field value
func (o *ServiceAccountSchema) GetId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ServiceAccountSchema) GetIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ServiceAccountSchema) SetId(v float32) {
	o.Id = v
}

// GetIsAPI returns the IsAPI field value if set, zero value otherwise.
// Deprecated
func (o *ServiceAccountSchema) GetIsAPI() bool {
	if o == nil || IsNil(o.IsAPI) {
		var ret bool
		return ret
	}
	return *o.IsAPI
}

// GetIsAPIOk returns a tuple with the IsAPI field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *ServiceAccountSchema) GetIsAPIOk() (*bool, bool) {
	if o == nil || IsNil(o.IsAPI) {
		return nil, false
	}
	return o.IsAPI, true
}

// HasIsAPI returns a boolean if a field has been set.
func (o *ServiceAccountSchema) HasIsAPI() bool {
	if o != nil && !IsNil(o.IsAPI) {
		return true
	}

	return false
}

// SetIsAPI gets a reference to the given bool and assigns it to the IsAPI field.
// Deprecated
func (o *ServiceAccountSchema) SetIsAPI(v bool) {
	o.IsAPI = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ServiceAccountSchema) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccountSchema) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ServiceAccountSchema) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ServiceAccountSchema) SetName(v string) {
	o.Name = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
// Deprecated
func (o *ServiceAccountSchema) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *ServiceAccountSchema) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *ServiceAccountSchema) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
// Deprecated
func (o *ServiceAccountSchema) SetEmail(v string) {
	o.Email = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *ServiceAccountSchema) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccountSchema) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *ServiceAccountSchema) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *ServiceAccountSchema) SetUsername(v string) {
	o.Username = &v
}

// GetImageUrl returns the ImageUrl field value if set, zero value otherwise.
func (o *ServiceAccountSchema) GetImageUrl() string {
	if o == nil || IsNil(o.ImageUrl) {
		var ret string
		return ret
	}
	return *o.ImageUrl
}

// GetImageUrlOk returns a tuple with the ImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccountSchema) GetImageUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ImageUrl) {
		return nil, false
	}
	return o.ImageUrl, true
}

// HasImageUrl returns a boolean if a field has been set.
func (o *ServiceAccountSchema) HasImageUrl() bool {
	if o != nil && !IsNil(o.ImageUrl) {
		return true
	}

	return false
}

// SetImageUrl gets a reference to the given string and assigns it to the ImageUrl field.
func (o *ServiceAccountSchema) SetImageUrl(v string) {
	o.ImageUrl = &v
}

// GetInviteLink returns the InviteLink field value if set, zero value otherwise.
// Deprecated
func (o *ServiceAccountSchema) GetInviteLink() string {
	if o == nil || IsNil(o.InviteLink) {
		var ret string
		return ret
	}
	return *o.InviteLink
}

// GetInviteLinkOk returns a tuple with the InviteLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *ServiceAccountSchema) GetInviteLinkOk() (*string, bool) {
	if o == nil || IsNil(o.InviteLink) {
		return nil, false
	}
	return o.InviteLink, true
}

// HasInviteLink returns a boolean if a field has been set.
func (o *ServiceAccountSchema) HasInviteLink() bool {
	if o != nil && !IsNil(o.InviteLink) {
		return true
	}

	return false
}

// SetInviteLink gets a reference to the given string and assigns it to the InviteLink field.
// Deprecated
func (o *ServiceAccountSchema) SetInviteLink(v string) {
	o.InviteLink = &v
}

// GetLoginAttempts returns the LoginAttempts field value if set, zero value otherwise.
// Deprecated
func (o *ServiceAccountSchema) GetLoginAttempts() float32 {
	if o == nil || IsNil(o.LoginAttempts) {
		var ret float32
		return ret
	}
	return *o.LoginAttempts
}

// GetLoginAttemptsOk returns a tuple with the LoginAttempts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *ServiceAccountSchema) GetLoginAttemptsOk() (*float32, bool) {
	if o == nil || IsNil(o.LoginAttempts) {
		return nil, false
	}
	return o.LoginAttempts, true
}

// HasLoginAttempts returns a boolean if a field has been set.
func (o *ServiceAccountSchema) HasLoginAttempts() bool {
	if o != nil && !IsNil(o.LoginAttempts) {
		return true
	}

	return false
}

// SetLoginAttempts gets a reference to the given float32 and assigns it to the LoginAttempts field.
// Deprecated
func (o *ServiceAccountSchema) SetLoginAttempts(v float32) {
	o.LoginAttempts = &v
}

// GetEmailSent returns the EmailSent field value if set, zero value otherwise.
// Deprecated
func (o *ServiceAccountSchema) GetEmailSent() bool {
	if o == nil || IsNil(o.EmailSent) {
		var ret bool
		return ret
	}
	return *o.EmailSent
}

// GetEmailSentOk returns a tuple with the EmailSent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *ServiceAccountSchema) GetEmailSentOk() (*bool, bool) {
	if o == nil || IsNil(o.EmailSent) {
		return nil, false
	}
	return o.EmailSent, true
}

// HasEmailSent returns a boolean if a field has been set.
func (o *ServiceAccountSchema) HasEmailSent() bool {
	if o != nil && !IsNil(o.EmailSent) {
		return true
	}

	return false
}

// SetEmailSent gets a reference to the given bool and assigns it to the EmailSent field.
// Deprecated
func (o *ServiceAccountSchema) SetEmailSent(v bool) {
	o.EmailSent = &v
}

// GetRootRole returns the RootRole field value if set, zero value otherwise.
func (o *ServiceAccountSchema) GetRootRole() int32 {
	if o == nil || IsNil(o.RootRole) {
		var ret int32
		return ret
	}
	return *o.RootRole
}

// GetRootRoleOk returns a tuple with the RootRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccountSchema) GetRootRoleOk() (*int32, bool) {
	if o == nil || IsNil(o.RootRole) {
		return nil, false
	}
	return o.RootRole, true
}

// HasRootRole returns a boolean if a field has been set.
func (o *ServiceAccountSchema) HasRootRole() bool {
	if o != nil && !IsNil(o.RootRole) {
		return true
	}

	return false
}

// SetRootRole gets a reference to the given int32 and assigns it to the RootRole field.
func (o *ServiceAccountSchema) SetRootRole(v int32) {
	o.RootRole = &v
}

// GetSeenAt returns the SeenAt field value if set, zero value otherwise (both if not set or set to explicit null).
// Deprecated
func (o *ServiceAccountSchema) GetSeenAt() time.Time {
	if o == nil || IsNil(o.SeenAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.SeenAt.Get()
}

// GetSeenAtOk returns a tuple with the SeenAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
// Deprecated
func (o *ServiceAccountSchema) GetSeenAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.SeenAt.Get(), o.SeenAt.IsSet()
}

// HasSeenAt returns a boolean if a field has been set.
func (o *ServiceAccountSchema) HasSeenAt() bool {
	if o != nil && o.SeenAt.IsSet() {
		return true
	}

	return false
}

// SetSeenAt gets a reference to the given NullableTime and assigns it to the SeenAt field.
// Deprecated
func (o *ServiceAccountSchema) SetSeenAt(v time.Time) {
	o.SeenAt.Set(&v)
}

// SetSeenAtNil sets the value for SeenAt to be an explicit nil
func (o *ServiceAccountSchema) SetSeenAtNil() {
	o.SeenAt.Set(nil)
}

// UnsetSeenAt ensures that no value is present for SeenAt, not even an explicit nil
func (o *ServiceAccountSchema) UnsetSeenAt() {
	o.SeenAt.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ServiceAccountSchema) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccountSchema) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ServiceAccountSchema) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *ServiceAccountSchema) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetTokens returns the Tokens field value if set, zero value otherwise.
func (o *ServiceAccountSchema) GetTokens() []PatSchema {
	if o == nil || IsNil(o.Tokens) {
		var ret []PatSchema
		return ret
	}
	return o.Tokens
}

// GetTokensOk returns a tuple with the Tokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccountSchema) GetTokensOk() ([]PatSchema, bool) {
	if o == nil || IsNil(o.Tokens) {
		return nil, false
	}
	return o.Tokens, true
}

// HasTokens returns a boolean if a field has been set.
func (o *ServiceAccountSchema) HasTokens() bool {
	if o != nil && !IsNil(o.Tokens) {
		return true
	}

	return false
}

// SetTokens gets a reference to the given []PatSchema and assigns it to the Tokens field.
func (o *ServiceAccountSchema) SetTokens(v []PatSchema) {
	o.Tokens = v
}

func (o ServiceAccountSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceAccountSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.IsAPI) {
		toSerialize["isAPI"] = o.IsAPI
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
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
	if !IsNil(o.Tokens) {
		toSerialize["tokens"] = o.Tokens
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ServiceAccountSchema) UnmarshalJSON(data []byte) (err error) {
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

	varServiceAccountSchema := _ServiceAccountSchema{}

	err = json.Unmarshal(data, &varServiceAccountSchema)

	if err != nil {
		return err
	}

	*o = ServiceAccountSchema(varServiceAccountSchema)

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
		delete(additionalProperties, "tokens")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableServiceAccountSchema struct {
	value *ServiceAccountSchema
	isSet bool
}

func (v NullableServiceAccountSchema) Get() *ServiceAccountSchema {
	return v.value
}

func (v *NullableServiceAccountSchema) Set(val *ServiceAccountSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceAccountSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceAccountSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceAccountSchema(val *ServiceAccountSchema) *NullableServiceAccountSchema {
	return &NullableServiceAccountSchema{value: val, isSet: true}
}

func (v NullableServiceAccountSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceAccountSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
