/*
Unleash API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 5.3.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the LoginSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LoginSchema{}

// LoginSchema A username/password login request
type LoginSchema struct {
	// The username trying to log in
	Username string `json:"username"`
	// The password of the user trying to log in
	Password string `json:"password"`
}

// NewLoginSchema instantiates a new LoginSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoginSchema(username string, password string) *LoginSchema {
	this := LoginSchema{}
	this.Username = username
	this.Password = password
	return &this
}

// NewLoginSchemaWithDefaults instantiates a new LoginSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoginSchemaWithDefaults() *LoginSchema {
	this := LoginSchema{}
	return &this
}

// GetUsername returns the Username field value
func (o *LoginSchema) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *LoginSchema) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *LoginSchema) SetUsername(v string) {
	o.Username = v
}

// GetPassword returns the Password field value
func (o *LoginSchema) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *LoginSchema) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *LoginSchema) SetPassword(v string) {
	o.Password = v
}

func (o LoginSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LoginSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["username"] = o.Username
	toSerialize["password"] = o.Password
	return toSerialize, nil
}

type NullableLoginSchema struct {
	value *LoginSchema
	isSet bool
}

func (v NullableLoginSchema) Get() *LoginSchema {
	return v.value
}

func (v *NullableLoginSchema) Set(val *LoginSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableLoginSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableLoginSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoginSchema(val *LoginSchema) *NullableLoginSchema {
	return &NullableLoginSchema{value: val, isSet: true}
}

func (v NullableLoginSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoginSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
