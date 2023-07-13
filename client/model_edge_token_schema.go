/*
Unleash API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 5.2.0-main
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the EdgeTokenSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EdgeTokenSchema{}

// EdgeTokenSchema A representation of a client token, limiting access to [CLIENT](https://docs.getunleash.io/reference/api-tokens-and-client-keys#client-tokens) (used by serverside SDKs) or [FRONTEND](https://docs.getunleash.io/reference/api-tokens-and-client-keys#front-end-tokens) (used by proxy SDKs)
type EdgeTokenSchema struct {
	// The list of projects this token has access to. If the token has access to specific projects they will be listed here. If the token has access to all projects it will be represented as [`*`]
	Projects []string `json:"projects"`
	// The [API token](https://docs.getunleash.io/reference/api-tokens-and-client-keys#api-tokens)'s **type**. Unleash supports three different types of API tokens ([ADMIN](https://docs.getunleash.io/reference/api-tokens-and-client-keys#admin-tokens), [CLIENT](https://docs.getunleash.io/reference/api-tokens-and-client-keys#client-tokens), [FRONTEND](https://docs.getunleash.io/reference/api-tokens-and-client-keys#front-end-tokens)). They all have varying access, so when validating a token it's important to know what kind you're dealing with
	Type string `json:"type"`
	// The actual token value. [Unleash API tokens](https://docs.getunleash.io/reference/api-tokens-and-client-keys) are comprised of three parts. <project(s)>:<environment>.randomcharacters
	Token string `json:"token"`
}

// NewEdgeTokenSchema instantiates a new EdgeTokenSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEdgeTokenSchema(projects []string, type_ string, token string) *EdgeTokenSchema {
	this := EdgeTokenSchema{}
	this.Projects = projects
	this.Type = type_
	this.Token = token
	return &this
}

// NewEdgeTokenSchemaWithDefaults instantiates a new EdgeTokenSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEdgeTokenSchemaWithDefaults() *EdgeTokenSchema {
	this := EdgeTokenSchema{}
	return &this
}

// GetProjects returns the Projects field value
func (o *EdgeTokenSchema) GetProjects() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Projects
}

// GetProjectsOk returns a tuple with the Projects field value
// and a boolean to check if the value has been set.
func (o *EdgeTokenSchema) GetProjectsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Projects, true
}

// SetProjects sets field value
func (o *EdgeTokenSchema) SetProjects(v []string) {
	o.Projects = v
}

// GetType returns the Type field value
func (o *EdgeTokenSchema) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *EdgeTokenSchema) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *EdgeTokenSchema) SetType(v string) {
	o.Type = v
}

// GetToken returns the Token field value
func (o *EdgeTokenSchema) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *EdgeTokenSchema) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *EdgeTokenSchema) SetToken(v string) {
	o.Token = v
}

func (o EdgeTokenSchema) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EdgeTokenSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["projects"] = o.Projects
	toSerialize["type"] = o.Type
	toSerialize["token"] = o.Token
	return toSerialize, nil
}

type NullableEdgeTokenSchema struct {
	value *EdgeTokenSchema
	isSet bool
}

func (v NullableEdgeTokenSchema) Get() *EdgeTokenSchema {
	return v.value
}

func (v *NullableEdgeTokenSchema) Set(val *EdgeTokenSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableEdgeTokenSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableEdgeTokenSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEdgeTokenSchema(val *EdgeTokenSchema) *NullableEdgeTokenSchema {
	return &NullableEdgeTokenSchema{value: val, isSet: true}
}

func (v NullableEdgeTokenSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEdgeTokenSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


