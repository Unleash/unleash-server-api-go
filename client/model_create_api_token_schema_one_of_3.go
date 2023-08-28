/*
Unleash API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 5.4.0-main
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the CreateApiTokenSchemaOneOf3 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateApiTokenSchemaOneOf3{}

// CreateApiTokenSchemaOneOf3 struct for CreateApiTokenSchemaOneOf3
type CreateApiTokenSchemaOneOf3 struct {
	// A client or frontend token. Must be one of the strings \"client\" or \"frontend\" (not case sensitive).
	Type string `json:"type"`
	// The environment that the token should be valid for. Defaults to \"default\"
	Environment *string `json:"environment,omitempty"`
	// The project that the token should be valid for. Defaults to \"*\" meaning every project. This property is mutually incompatible with the `projects` property. If you specify one, you cannot specify the other.
	Project *string `json:"project,omitempty"`
	// A list of projects that the token should be valid for. This property is mutually incompatible with the `project` property. If you specify one, you cannot specify the other.
	Projects []string `json:"projects,omitempty"`
	// The name of the token. This property is deprecated. Use `tokenName` instead.
	// Deprecated
	Username string `json:"username"`
}

// NewCreateApiTokenSchemaOneOf3 instantiates a new CreateApiTokenSchemaOneOf3 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateApiTokenSchemaOneOf3(type_ string, username string) *CreateApiTokenSchemaOneOf3 {
	this := CreateApiTokenSchemaOneOf3{}
	this.Type = type_
	this.Username = username
	return &this
}

// NewCreateApiTokenSchemaOneOf3WithDefaults instantiates a new CreateApiTokenSchemaOneOf3 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateApiTokenSchemaOneOf3WithDefaults() *CreateApiTokenSchemaOneOf3 {
	this := CreateApiTokenSchemaOneOf3{}
	return &this
}

// GetType returns the Type field value
func (o *CreateApiTokenSchemaOneOf3) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CreateApiTokenSchemaOneOf3) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CreateApiTokenSchemaOneOf3) SetType(v string) {
	o.Type = v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *CreateApiTokenSchemaOneOf3) GetEnvironment() string {
	if o == nil || IsNil(o.Environment) {
		var ret string
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateApiTokenSchemaOneOf3) GetEnvironmentOk() (*string, bool) {
	if o == nil || IsNil(o.Environment) {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *CreateApiTokenSchemaOneOf3) HasEnvironment() bool {
	if o != nil && !IsNil(o.Environment) {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given string and assigns it to the Environment field.
func (o *CreateApiTokenSchemaOneOf3) SetEnvironment(v string) {
	o.Environment = &v
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *CreateApiTokenSchemaOneOf3) GetProject() string {
	if o == nil || IsNil(o.Project) {
		var ret string
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateApiTokenSchemaOneOf3) GetProjectOk() (*string, bool) {
	if o == nil || IsNil(o.Project) {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *CreateApiTokenSchemaOneOf3) HasProject() bool {
	if o != nil && !IsNil(o.Project) {
		return true
	}

	return false
}

// SetProject gets a reference to the given string and assigns it to the Project field.
func (o *CreateApiTokenSchemaOneOf3) SetProject(v string) {
	o.Project = &v
}

// GetProjects returns the Projects field value if set, zero value otherwise.
func (o *CreateApiTokenSchemaOneOf3) GetProjects() []string {
	if o == nil || IsNil(o.Projects) {
		var ret []string
		return ret
	}
	return o.Projects
}

// GetProjectsOk returns a tuple with the Projects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateApiTokenSchemaOneOf3) GetProjectsOk() ([]string, bool) {
	if o == nil || IsNil(o.Projects) {
		return nil, false
	}
	return o.Projects, true
}

// HasProjects returns a boolean if a field has been set.
func (o *CreateApiTokenSchemaOneOf3) HasProjects() bool {
	if o != nil && !IsNil(o.Projects) {
		return true
	}

	return false
}

// SetProjects gets a reference to the given []string and assigns it to the Projects field.
func (o *CreateApiTokenSchemaOneOf3) SetProjects(v []string) {
	o.Projects = v
}

// GetUsername returns the Username field value
// Deprecated
func (o *CreateApiTokenSchemaOneOf3) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
// Deprecated
func (o *CreateApiTokenSchemaOneOf3) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
// Deprecated
func (o *CreateApiTokenSchemaOneOf3) SetUsername(v string) {
	o.Username = v
}

func (o CreateApiTokenSchemaOneOf3) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateApiTokenSchemaOneOf3) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.Environment) {
		toSerialize["environment"] = o.Environment
	}
	if !IsNil(o.Project) {
		toSerialize["project"] = o.Project
	}
	if !IsNil(o.Projects) {
		toSerialize["projects"] = o.Projects
	}
	toSerialize["username"] = o.Username
	return toSerialize, nil
}

type NullableCreateApiTokenSchemaOneOf3 struct {
	value *CreateApiTokenSchemaOneOf3
	isSet bool
}

func (v NullableCreateApiTokenSchemaOneOf3) Get() *CreateApiTokenSchemaOneOf3 {
	return v.value
}

func (v *NullableCreateApiTokenSchemaOneOf3) Set(val *CreateApiTokenSchemaOneOf3) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateApiTokenSchemaOneOf3) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateApiTokenSchemaOneOf3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateApiTokenSchemaOneOf3(val *CreateApiTokenSchemaOneOf3) *NullableCreateApiTokenSchemaOneOf3 {
	return &NullableCreateApiTokenSchemaOneOf3{value: val, isSet: true}
}

func (v NullableCreateApiTokenSchemaOneOf3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateApiTokenSchemaOneOf3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
