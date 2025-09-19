# UpdateEnvironmentSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | Updates the type of environment (i.e. development or production). | [optional] 
**SortOrder** | Pointer to **int32** | Changes the sort order of this environment. | [optional] 
**RequiredApprovals** | Pointer to **NullableInt32** | Experimental field. The number of approvals required before a change request can be applied in this environment. | [optional] 

## Methods

### NewUpdateEnvironmentSchema

`func NewUpdateEnvironmentSchema() *UpdateEnvironmentSchema`

NewUpdateEnvironmentSchema instantiates a new UpdateEnvironmentSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateEnvironmentSchemaWithDefaults

`func NewUpdateEnvironmentSchemaWithDefaults() *UpdateEnvironmentSchema`

NewUpdateEnvironmentSchemaWithDefaults instantiates a new UpdateEnvironmentSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *UpdateEnvironmentSchema) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateEnvironmentSchema) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateEnvironmentSchema) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UpdateEnvironmentSchema) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSortOrder

`func (o *UpdateEnvironmentSchema) GetSortOrder() int32`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *UpdateEnvironmentSchema) GetSortOrderOk() (*int32, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *UpdateEnvironmentSchema) SetSortOrder(v int32)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *UpdateEnvironmentSchema) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### GetRequiredApprovals

`func (o *UpdateEnvironmentSchema) GetRequiredApprovals() int32`

GetRequiredApprovals returns the RequiredApprovals field if non-nil, zero value otherwise.

### GetRequiredApprovalsOk

`func (o *UpdateEnvironmentSchema) GetRequiredApprovalsOk() (*int32, bool)`

GetRequiredApprovalsOk returns a tuple with the RequiredApprovals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredApprovals

`func (o *UpdateEnvironmentSchema) SetRequiredApprovals(v int32)`

SetRequiredApprovals sets RequiredApprovals field to given value.

### HasRequiredApprovals

`func (o *UpdateEnvironmentSchema) HasRequiredApprovals() bool`

HasRequiredApprovals returns a boolean if a field has been set.

### SetRequiredApprovalsNil

`func (o *UpdateEnvironmentSchema) SetRequiredApprovalsNil(b bool)`

 SetRequiredApprovalsNil sets the value for RequiredApprovals to be an explicit nil

### UnsetRequiredApprovals
`func (o *UpdateEnvironmentSchema) UnsetRequiredApprovals()`

UnsetRequiredApprovals ensures that no value is present for RequiredApprovals, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


