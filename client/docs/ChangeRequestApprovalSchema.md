# ChangeRequestApprovalSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedBy** | [**ChangeRequestApprovalSchemaCreatedBy**](ChangeRequestApprovalSchemaCreatedBy.md) |  | 
**CreatedAt** | **time.Time** | When the approval was given. | 

## Methods

### NewChangeRequestApprovalSchema

`func NewChangeRequestApprovalSchema(createdBy ChangeRequestApprovalSchemaCreatedBy, createdAt time.Time, ) *ChangeRequestApprovalSchema`

NewChangeRequestApprovalSchema instantiates a new ChangeRequestApprovalSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChangeRequestApprovalSchemaWithDefaults

`func NewChangeRequestApprovalSchemaWithDefaults() *ChangeRequestApprovalSchema`

NewChangeRequestApprovalSchemaWithDefaults instantiates a new ChangeRequestApprovalSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedBy

`func (o *ChangeRequestApprovalSchema) GetCreatedBy() ChangeRequestApprovalSchemaCreatedBy`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ChangeRequestApprovalSchema) GetCreatedByOk() (*ChangeRequestApprovalSchemaCreatedBy, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ChangeRequestApprovalSchema) SetCreatedBy(v ChangeRequestApprovalSchemaCreatedBy)`

SetCreatedBy sets CreatedBy field to given value.


### GetCreatedAt

`func (o *ChangeRequestApprovalSchema) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ChangeRequestApprovalSchema) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ChangeRequestApprovalSchema) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


