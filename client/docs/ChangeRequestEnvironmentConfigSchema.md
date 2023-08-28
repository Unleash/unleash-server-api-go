# ChangeRequestEnvironmentConfigSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Environment** | **string** | The environment that this configuration applies to. | 
**Type** | **string** | The [type of the environment](https://docs.getunleash.io/reference/environments#environment-types) listed in &#x60;environment&#x60;. | 
**ChangeRequestEnabled** | **bool** | &#x60;true&#x60; if this environment has change requests enabled, otherwise &#x60;false&#x60;. | 
**RequiredApprovals** | **NullableFloat32** | The number of approvals that are required for a change request to be fully approved and ready to be applied in this environment. | 

## Methods

### NewChangeRequestEnvironmentConfigSchema

`func NewChangeRequestEnvironmentConfigSchema(environment string, type_ string, changeRequestEnabled bool, requiredApprovals NullableFloat32, ) *ChangeRequestEnvironmentConfigSchema`

NewChangeRequestEnvironmentConfigSchema instantiates a new ChangeRequestEnvironmentConfigSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChangeRequestEnvironmentConfigSchemaWithDefaults

`func NewChangeRequestEnvironmentConfigSchemaWithDefaults() *ChangeRequestEnvironmentConfigSchema`

NewChangeRequestEnvironmentConfigSchemaWithDefaults instantiates a new ChangeRequestEnvironmentConfigSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironment

`func (o *ChangeRequestEnvironmentConfigSchema) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *ChangeRequestEnvironmentConfigSchema) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *ChangeRequestEnvironmentConfigSchema) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.


### GetType

`func (o *ChangeRequestEnvironmentConfigSchema) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ChangeRequestEnvironmentConfigSchema) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ChangeRequestEnvironmentConfigSchema) SetType(v string)`

SetType sets Type field to given value.


### GetChangeRequestEnabled

`func (o *ChangeRequestEnvironmentConfigSchema) GetChangeRequestEnabled() bool`

GetChangeRequestEnabled returns the ChangeRequestEnabled field if non-nil, zero value otherwise.

### GetChangeRequestEnabledOk

`func (o *ChangeRequestEnvironmentConfigSchema) GetChangeRequestEnabledOk() (*bool, bool)`

GetChangeRequestEnabledOk returns a tuple with the ChangeRequestEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeRequestEnabled

`func (o *ChangeRequestEnvironmentConfigSchema) SetChangeRequestEnabled(v bool)`

SetChangeRequestEnabled sets ChangeRequestEnabled field to given value.


### GetRequiredApprovals

`func (o *ChangeRequestEnvironmentConfigSchema) GetRequiredApprovals() float32`

GetRequiredApprovals returns the RequiredApprovals field if non-nil, zero value otherwise.

### GetRequiredApprovalsOk

`func (o *ChangeRequestEnvironmentConfigSchema) GetRequiredApprovalsOk() (*float32, bool)`

GetRequiredApprovalsOk returns a tuple with the RequiredApprovals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredApprovals

`func (o *ChangeRequestEnvironmentConfigSchema) SetRequiredApprovals(v float32)`

SetRequiredApprovals sets RequiredApprovals field to given value.


### SetRequiredApprovalsNil

`func (o *ChangeRequestEnvironmentConfigSchema) SetRequiredApprovalsNil(b bool)`

 SetRequiredApprovalsNil sets the value for RequiredApprovals to be an explicit nil

### UnsetRequiredApprovals
`func (o *ChangeRequestEnvironmentConfigSchema) UnsetRequiredApprovals()`

UnsetRequiredApprovals ensures that no value is present for RequiredApprovals, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


