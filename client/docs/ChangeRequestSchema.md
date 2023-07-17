# ChangeRequestSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** |  | 
**Title** | Pointer to **string** |  | [optional] 
**Environment** | **string** |  | 
**State** | **string** |  | 
**MinApprovals** | **float32** |  | 
**Project** | **string** |  | 
**Features** | [**[]ChangeRequestFeatureSchema**](ChangeRequestFeatureSchema.md) |  | 
**Approvals** | Pointer to [**[]ChangeRequestApprovalSchema**](ChangeRequestApprovalSchema.md) |  | [optional] 
**Comments** | Pointer to [**[]ChangeRequestCommentSchema**](ChangeRequestCommentSchema.md) |  | [optional] 
**CreatedBy** | [**ChangeRequestEventSchemaCreatedBy**](ChangeRequestEventSchemaCreatedBy.md) |  | 
**CreatedAt** | **time.Time** |  | 

## Methods

### NewChangeRequestSchema

`func NewChangeRequestSchema(id float32, environment string, state string, minApprovals float32, project string, features []ChangeRequestFeatureSchema, createdBy ChangeRequestEventSchemaCreatedBy, createdAt time.Time, ) *ChangeRequestSchema`

NewChangeRequestSchema instantiates a new ChangeRequestSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChangeRequestSchemaWithDefaults

`func NewChangeRequestSchemaWithDefaults() *ChangeRequestSchema`

NewChangeRequestSchemaWithDefaults instantiates a new ChangeRequestSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ChangeRequestSchema) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ChangeRequestSchema) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ChangeRequestSchema) SetId(v float32)`

SetId sets Id field to given value.


### GetTitle

`func (o *ChangeRequestSchema) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ChangeRequestSchema) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ChangeRequestSchema) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ChangeRequestSchema) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetEnvironment

`func (o *ChangeRequestSchema) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *ChangeRequestSchema) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *ChangeRequestSchema) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.


### GetState

`func (o *ChangeRequestSchema) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ChangeRequestSchema) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ChangeRequestSchema) SetState(v string)`

SetState sets State field to given value.


### GetMinApprovals

`func (o *ChangeRequestSchema) GetMinApprovals() float32`

GetMinApprovals returns the MinApprovals field if non-nil, zero value otherwise.

### GetMinApprovalsOk

`func (o *ChangeRequestSchema) GetMinApprovalsOk() (*float32, bool)`

GetMinApprovalsOk returns a tuple with the MinApprovals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinApprovals

`func (o *ChangeRequestSchema) SetMinApprovals(v float32)`

SetMinApprovals sets MinApprovals field to given value.


### GetProject

`func (o *ChangeRequestSchema) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ChangeRequestSchema) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ChangeRequestSchema) SetProject(v string)`

SetProject sets Project field to given value.


### GetFeatures

`func (o *ChangeRequestSchema) GetFeatures() []ChangeRequestFeatureSchema`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *ChangeRequestSchema) GetFeaturesOk() (*[]ChangeRequestFeatureSchema, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *ChangeRequestSchema) SetFeatures(v []ChangeRequestFeatureSchema)`

SetFeatures sets Features field to given value.


### GetApprovals

`func (o *ChangeRequestSchema) GetApprovals() []ChangeRequestApprovalSchema`

GetApprovals returns the Approvals field if non-nil, zero value otherwise.

### GetApprovalsOk

`func (o *ChangeRequestSchema) GetApprovalsOk() (*[]ChangeRequestApprovalSchema, bool)`

GetApprovalsOk returns a tuple with the Approvals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovals

`func (o *ChangeRequestSchema) SetApprovals(v []ChangeRequestApprovalSchema)`

SetApprovals sets Approvals field to given value.

### HasApprovals

`func (o *ChangeRequestSchema) HasApprovals() bool`

HasApprovals returns a boolean if a field has been set.

### GetComments

`func (o *ChangeRequestSchema) GetComments() []ChangeRequestCommentSchema`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *ChangeRequestSchema) GetCommentsOk() (*[]ChangeRequestCommentSchema, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *ChangeRequestSchema) SetComments(v []ChangeRequestCommentSchema)`

SetComments sets Comments field to given value.

### HasComments

`func (o *ChangeRequestSchema) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetCreatedBy

`func (o *ChangeRequestSchema) GetCreatedBy() ChangeRequestEventSchemaCreatedBy`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ChangeRequestSchema) GetCreatedByOk() (*ChangeRequestEventSchemaCreatedBy, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ChangeRequestSchema) SetCreatedBy(v ChangeRequestEventSchemaCreatedBy)`

SetCreatedBy sets CreatedBy field to given value.


### GetCreatedAt

`func (o *ChangeRequestSchema) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ChangeRequestSchema) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ChangeRequestSchema) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


