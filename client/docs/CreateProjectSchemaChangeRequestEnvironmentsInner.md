# CreateProjectSchemaChangeRequestEnvironmentsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the environment to configure change requests for. | 
**RequiredApprovals** | Pointer to **int32** | The number of approvals required for a change request to be fully approved and ready to applied in this environment. If no value is provided, it will be set to the default number, which is 1. Values will be clamped to between 1 and 10 inclusive. | [optional] [default to 1]

## Methods

### NewCreateProjectSchemaChangeRequestEnvironmentsInner

`func NewCreateProjectSchemaChangeRequestEnvironmentsInner(name string, ) *CreateProjectSchemaChangeRequestEnvironmentsInner`

NewCreateProjectSchemaChangeRequestEnvironmentsInner instantiates a new CreateProjectSchemaChangeRequestEnvironmentsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateProjectSchemaChangeRequestEnvironmentsInnerWithDefaults

`func NewCreateProjectSchemaChangeRequestEnvironmentsInnerWithDefaults() *CreateProjectSchemaChangeRequestEnvironmentsInner`

NewCreateProjectSchemaChangeRequestEnvironmentsInnerWithDefaults instantiates a new CreateProjectSchemaChangeRequestEnvironmentsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateProjectSchemaChangeRequestEnvironmentsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateProjectSchemaChangeRequestEnvironmentsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateProjectSchemaChangeRequestEnvironmentsInner) SetName(v string)`

SetName sets Name field to given value.


### GetRequiredApprovals

`func (o *CreateProjectSchemaChangeRequestEnvironmentsInner) GetRequiredApprovals() int32`

GetRequiredApprovals returns the RequiredApprovals field if non-nil, zero value otherwise.

### GetRequiredApprovalsOk

`func (o *CreateProjectSchemaChangeRequestEnvironmentsInner) GetRequiredApprovalsOk() (*int32, bool)`

GetRequiredApprovalsOk returns a tuple with the RequiredApprovals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredApprovals

`func (o *CreateProjectSchemaChangeRequestEnvironmentsInner) SetRequiredApprovals(v int32)`

SetRequiredApprovals sets RequiredApprovals field to given value.

### HasRequiredApprovals

`func (o *CreateProjectSchemaChangeRequestEnvironmentsInner) HasRequiredApprovals() bool`

HasRequiredApprovals returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


