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
	"time"
)

// checks if the ApiTokenSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiTokenSchema{}

// ApiTokenSchema An overview of an [Unleash API token](https://docs.getunleash.io/reference/api-tokens-and-client-keys).
type ApiTokenSchema struct {
	// The token used for authentication.
	Secret string `json:"secret"`
	// This property was deprecated in Unleash v5. Prefer the `tokenName` property instead.
	// Deprecated
	Username *string `json:"username,omitempty"`
	// A unique name for this particular token
	TokenName string `json:"tokenName"`
	// The type of API token
	Type string `json:"type"`
	// The environment the token has access to. `*` if it has access to all environments.
	Environment *string `json:"environment,omitempty"`
	// The project this token belongs to.
	Project string `json:"project"`
	// The list of projects this token has access to. If the token has access to specific projects they will be listed here. If the token has access to all projects it will be represented as `[*]`
	Projects []string `json:"projects"`
	// The token's expiration date. NULL if the token doesn't have an expiration set.
	ExpiresAt NullableTime `json:"expiresAt,omitempty"`
	// When the token was created.
	CreatedAt time.Time `json:"createdAt"`
	// When the token was last seen/used to authenticate with. NULL if the token has not yet been used for authentication.
	SeenAt NullableTime `json:"seenAt,omitempty"`
	// Alias is no longer in active use and will often be NULL. It's kept around as a way of allowing old proxy tokens created with the old metadata format to keep working.
	Alias NullableString `json:"alias,omitempty"`
}

type _ApiTokenSchema ApiTokenSchema

// NewApiTokenSchema instantiates a new ApiTokenSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiTokenSchema(secret string, tokenName string, type_ string, project string, projects []string, createdAt time.Time) *ApiTokenSchema {
	this := ApiTokenSchema{}
	this.Secret = secret
	this.TokenName = tokenName
	this.Type = type_
	this.Project = project
	this.Projects = projects
	this.CreatedAt = createdAt
	return &this
}

// NewApiTokenSchemaWithDefaults instantiates a new ApiTokenSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiTokenSchemaWithDefaults() *ApiTokenSchema {
	this := ApiTokenSchema{}
	return &this
}

// GetSecret returns the Secret field value
func (o *ApiTokenSchema) GetSecret() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Secret
}

// GetSecretOk returns a tuple with the Secret field value
// and a boolean to check if the value has been set.
func (o *ApiTokenSchema) GetSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Secret, true
}

// SetSecret sets field value
func (o *ApiTokenSchema) SetSecret(v string) {
	o.Secret = v
}

// GetUsername returns the Username field value if set, zero value otherwise.
// Deprecated
func (o *ApiTokenSchema) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *ApiTokenSchema) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *ApiTokenSchema) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
// Deprecated
func (o *ApiTokenSchema) SetUsername(v string) {
	o.Username = &v
}

// GetTokenName returns the TokenName field value
func (o *ApiTokenSchema) GetTokenName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TokenName
}

// GetTokenNameOk returns a tuple with the TokenName field value
// and a boolean to check if the value has been set.
func (o *ApiTokenSchema) GetTokenNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenName, true
}

// SetTokenName sets field value
func (o *ApiTokenSchema) SetTokenName(v string) {
	o.TokenName = v
}

// GetType returns the Type field value
func (o *ApiTokenSchema) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ApiTokenSchema) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ApiTokenSchema) SetType(v string) {
	o.Type = v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *ApiTokenSchema) GetEnvironment() string {
	if o == nil || IsNil(o.Environment) {
		var ret string
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiTokenSchema) GetEnvironmentOk() (*string, bool) {
	if o == nil || IsNil(o.Environment) {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *ApiTokenSchema) HasEnvironment() bool {
	if o != nil && !IsNil(o.Environment) {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given string and assigns it to the Environment field.
func (o *ApiTokenSchema) SetEnvironment(v string) {
	o.Environment = &v
}

// GetProject returns the Project field value
func (o *ApiTokenSchema) GetProject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Project
}

// GetProjectOk returns a tuple with the Project field value
// and a boolean to check if the value has been set.
func (o *ApiTokenSchema) GetProjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Project, true
}

// SetProject sets field value
func (o *ApiTokenSchema) SetProject(v string) {
	o.Project = v
}

// GetProjects returns the Projects field value
func (o *ApiTokenSchema) GetProjects() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Projects
}

// GetProjectsOk returns a tuple with the Projects field value
// and a boolean to check if the value has been set.
func (o *ApiTokenSchema) GetProjectsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Projects, true
}

// SetProjects sets field value
func (o *ApiTokenSchema) SetProjects(v []string) {
	o.Projects = v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiTokenSchema) GetExpiresAt() time.Time {
	if o == nil || IsNil(o.ExpiresAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.ExpiresAt.Get()
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiTokenSchema) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpiresAt.Get(), o.ExpiresAt.IsSet()
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *ApiTokenSchema) HasExpiresAt() bool {
	if o != nil && o.ExpiresAt.IsSet() {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given NullableTime and assigns it to the ExpiresAt field.
func (o *ApiTokenSchema) SetExpiresAt(v time.Time) {
	o.ExpiresAt.Set(&v)
}

// SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil
func (o *ApiTokenSchema) SetExpiresAtNil() {
	o.ExpiresAt.Set(nil)
}

// UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
func (o *ApiTokenSchema) UnsetExpiresAt() {
	o.ExpiresAt.Unset()
}

// GetCreatedAt returns the CreatedAt field value
func (o *ApiTokenSchema) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ApiTokenSchema) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ApiTokenSchema) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetSeenAt returns the SeenAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiTokenSchema) GetSeenAt() time.Time {
	if o == nil || IsNil(o.SeenAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.SeenAt.Get()
}

// GetSeenAtOk returns a tuple with the SeenAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiTokenSchema) GetSeenAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.SeenAt.Get(), o.SeenAt.IsSet()
}

// HasSeenAt returns a boolean if a field has been set.
func (o *ApiTokenSchema) HasSeenAt() bool {
	if o != nil && o.SeenAt.IsSet() {
		return true
	}

	return false
}

// SetSeenAt gets a reference to the given NullableTime and assigns it to the SeenAt field.
func (o *ApiTokenSchema) SetSeenAt(v time.Time) {
	o.SeenAt.Set(&v)
}

// SetSeenAtNil sets the value for SeenAt to be an explicit nil
func (o *ApiTokenSchema) SetSeenAtNil() {
	o.SeenAt.Set(nil)
}

// UnsetSeenAt ensures that no value is present for SeenAt, not even an explicit nil
func (o *ApiTokenSchema) UnsetSeenAt() {
	o.SeenAt.Unset()
}

// GetAlias returns the Alias field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiTokenSchema) GetAlias() string {
	if o == nil || IsNil(o.Alias.Get()) {
		var ret string
		return ret
	}
	return *o.Alias.Get()
}

// GetAliasOk returns a tuple with the Alias field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiTokenSchema) GetAliasOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Alias.Get(), o.Alias.IsSet()
}

// HasAlias returns a boolean if a field has been set.
func (o *ApiTokenSchema) HasAlias() bool {
	if o != nil && o.Alias.IsSet() {
		return true
	}

	return false
}

// SetAlias gets a reference to the given NullableString and assigns it to the Alias field.
func (o *ApiTokenSchema) SetAlias(v string) {
	o.Alias.Set(&v)
}

// SetAliasNil sets the value for Alias to be an explicit nil
func (o *ApiTokenSchema) SetAliasNil() {
	o.Alias.Set(nil)
}

// UnsetAlias ensures that no value is present for Alias, not even an explicit nil
func (o *ApiTokenSchema) UnsetAlias() {
	o.Alias.Unset()
}

func (o ApiTokenSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiTokenSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["secret"] = o.Secret
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	toSerialize["tokenName"] = o.TokenName
	toSerialize["type"] = o.Type
	if !IsNil(o.Environment) {
		toSerialize["environment"] = o.Environment
	}
	toSerialize["project"] = o.Project
	toSerialize["projects"] = o.Projects
	if o.ExpiresAt.IsSet() {
		toSerialize["expiresAt"] = o.ExpiresAt.Get()
	}
	toSerialize["createdAt"] = o.CreatedAt
	if o.SeenAt.IsSet() {
		toSerialize["seenAt"] = o.SeenAt.Get()
	}
	if o.Alias.IsSet() {
		toSerialize["alias"] = o.Alias.Get()
	}
	return toSerialize, nil
}

func (o *ApiTokenSchema) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"secret",
		"tokenName",
		"type",
		"project",
		"projects",
		"createdAt",
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

	varApiTokenSchema := _ApiTokenSchema{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApiTokenSchema)

	if err != nil {
		return err
	}

	*o = ApiTokenSchema(varApiTokenSchema)

	return err
}

type NullableApiTokenSchema struct {
	value *ApiTokenSchema
	isSet bool
}

func (v NullableApiTokenSchema) Get() *ApiTokenSchema {
	return v.value
}

func (v *NullableApiTokenSchema) Set(val *ApiTokenSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableApiTokenSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableApiTokenSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiTokenSchema(val *ApiTokenSchema) *NullableApiTokenSchema {
	return &NullableApiTokenSchema{value: val, isSet: true}
}

func (v NullableApiTokenSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiTokenSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
