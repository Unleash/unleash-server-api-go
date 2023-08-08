# UpdateChangeRequestEnvironmentConfigSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChangeRequestsEnabled** | **bool** |  | 
**RequiredApprovals** | Pointer to **float32** |  | [optional] 

## Methods

### NewUpdateChangeRequestEnvironmentConfigSchema

`func NewUpdateChangeRequestEnvironmentConfigSchema(changeRequestsEnabled bool, ) *UpdateChangeRequestEnvironmentConfigSchema`

NewUpdateChangeRequestEnvironmentConfigSchema instantiates a new UpdateChangeRequestEnvironmentConfigSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateChangeRequestEnvironmentConfigSchemaWithDefaults

`func NewUpdateChangeRequestEnvironmentConfigSchemaWithDefaults() *UpdateChangeRequestEnvironmentConfigSchema`

NewUpdateChangeRequestEnvironmentConfigSchemaWithDefaults instantiates a new UpdateChangeRequestEnvironmentConfigSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChangeRequestsEnabled

`func (o *UpdateChangeRequestEnvironmentConfigSchema) GetChangeRequestsEnabled() bool`

GetChangeRequestsEnabled returns the ChangeRequestsEnabled field if non-nil, zero value otherwise.

### GetChangeRequestsEnabledOk

`func (o *UpdateChangeRequestEnvironmentConfigSchema) GetChangeRequestsEnabledOk() (*bool, bool)`

GetChangeRequestsEnabledOk returns a tuple with the ChangeRequestsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeRequestsEnabled

`func (o *UpdateChangeRequestEnvironmentConfigSchema) SetChangeRequestsEnabled(v bool)`

SetChangeRequestsEnabled sets ChangeRequestsEnabled field to given value.


### GetRequiredApprovals

`func (o *UpdateChangeRequestEnvironmentConfigSchema) GetRequiredApprovals() float32`

GetRequiredApprovals returns the RequiredApprovals field if non-nil, zero value otherwise.

### GetRequiredApprovalsOk

`func (o *UpdateChangeRequestEnvironmentConfigSchema) GetRequiredApprovalsOk() (*float32, bool)`

GetRequiredApprovalsOk returns a tuple with the RequiredApprovals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredApprovals

`func (o *UpdateChangeRequestEnvironmentConfigSchema) SetRequiredApprovals(v float32)`

SetRequiredApprovals sets RequiredApprovals field to given value.

### HasRequiredApprovals

`func (o *UpdateChangeRequestEnvironmentConfigSchema) HasRequiredApprovals() bool`

HasRequiredApprovals returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


