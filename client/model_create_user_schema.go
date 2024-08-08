/*
Unleash API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 6.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the CreateUserSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateUserSchema{}

// CreateUserSchema The payload must contain at least one of the name and email properties, though which one is up to you. For the user to be able to log in to the system, the user must have an email.
type CreateUserSchema struct {
	// The user's username. Must be provided if email is not provided.
	Username *string `json:"username,omitempty"`
	// The user's email address. Must be provided if username is not provided.
	Email *string `json:"email,omitempty"`
	// The user's name (not the user's username).
	Name *string `json:"name,omitempty"`
	// Password for the user
	Password *string                  `json:"password,omitempty"`
	RootRole CreateUserSchemaRootRole `json:"rootRole"`
	// Whether to send a welcome email with a login link to the user or not. Defaults to `true`.
	SendEmail *bool `json:"sendEmail,omitempty"`
}

type _CreateUserSchema CreateUserSchema

// NewCreateUserSchema instantiates a new CreateUserSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateUserSchema(rootRole CreateUserSchemaRootRole) *CreateUserSchema {
	this := CreateUserSchema{}
	this.RootRole = rootRole
	return &this
}

// NewCreateUserSchemaWithDefaults instantiates a new CreateUserSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateUserSchemaWithDefaults() *CreateUserSchema {
	this := CreateUserSchema{}
	return &this
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *CreateUserSchema) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUserSchema) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *CreateUserSchema) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *CreateUserSchema) SetUsername(v string) {
	o.Username = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *CreateUserSchema) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUserSchema) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *CreateUserSchema) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *CreateUserSchema) SetEmail(v string) {
	o.Email = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateUserSchema) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUserSchema) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateUserSchema) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreateUserSchema) SetName(v string) {
	o.Name = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *CreateUserSchema) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUserSchema) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *CreateUserSchema) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *CreateUserSchema) SetPassword(v string) {
	o.Password = &v
}

// GetRootRole returns the RootRole field value
func (o *CreateUserSchema) GetRootRole() CreateUserSchemaRootRole {
	if o == nil {
		var ret CreateUserSchemaRootRole
		return ret
	}

	return o.RootRole
}

// GetRootRoleOk returns a tuple with the RootRole field value
// and a boolean to check if the value has been set.
func (o *CreateUserSchema) GetRootRoleOk() (*CreateUserSchemaRootRole, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RootRole, true
}

// SetRootRole sets field value
func (o *CreateUserSchema) SetRootRole(v CreateUserSchemaRootRole) {
	o.RootRole = v
}

// GetSendEmail returns the SendEmail field value if set, zero value otherwise.
func (o *CreateUserSchema) GetSendEmail() bool {
	if o == nil || IsNil(o.SendEmail) {
		var ret bool
		return ret
	}
	return *o.SendEmail
}

// GetSendEmailOk returns a tuple with the SendEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUserSchema) GetSendEmailOk() (*bool, bool) {
	if o == nil || IsNil(o.SendEmail) {
		return nil, false
	}
	return o.SendEmail, true
}

// HasSendEmail returns a boolean if a field has been set.
func (o *CreateUserSchema) HasSendEmail() bool {
	if o != nil && !IsNil(o.SendEmail) {
		return true
	}

	return false
}

// SetSendEmail gets a reference to the given bool and assigns it to the SendEmail field.
func (o *CreateUserSchema) SetSendEmail(v bool) {
	o.SendEmail = &v
}

func (o CreateUserSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateUserSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	toSerialize["rootRole"] = o.RootRole
	if !IsNil(o.SendEmail) {
		toSerialize["sendEmail"] = o.SendEmail
	}
	return toSerialize, nil
}

func (o *CreateUserSchema) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"rootRole",
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

	varCreateUserSchema := _CreateUserSchema{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateUserSchema)

	if err != nil {
		return err
	}

	*o = CreateUserSchema(varCreateUserSchema)

	return err
}

type NullableCreateUserSchema struct {
	value *CreateUserSchema
	isSet bool
}

func (v NullableCreateUserSchema) Get() *CreateUserSchema {
	return v.value
}

func (v *NullableCreateUserSchema) Set(val *CreateUserSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateUserSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateUserSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateUserSchema(val *CreateUserSchema) *NullableCreateUserSchema {
	return &NullableCreateUserSchema{value: val, isSet: true}
}

func (v NullableCreateUserSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateUserSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
