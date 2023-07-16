# ChangeRequestEventSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** |  | 
**Action** | **string** |  | 
**Conflict** | Pointer to **string** |  | [optional] 
**Payload** | [**ChangeRequestEventSchemaPayload**](ChangeRequestEventSchemaPayload.md) |  | 
**UpdatedBy** | Pointer to **string** |  | [optional] 
**CreatedBy** | Pointer to [**ChangeRequestEventSchemaCreatedBy**](ChangeRequestEventSchemaCreatedBy.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewChangeRequestEventSchema

`func NewChangeRequestEventSchema(id float32, action string, payload ChangeRequestEventSchemaPayload, ) *ChangeRequestEventSchema`

NewChangeRequestEventSchema instantiates a new ChangeRequestEventSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChangeRequestEventSchemaWithDefaults

`func NewChangeRequestEventSchemaWithDefaults() *ChangeRequestEventSchema`

NewChangeRequestEventSchemaWithDefaults instantiates a new ChangeRequestEventSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ChangeRequestEventSchema) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ChangeRequestEventSchema) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ChangeRequestEventSchema) SetId(v float32)`

SetId sets Id field to given value.


### GetAction

`func (o *ChangeRequestEventSchema) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ChangeRequestEventSchema) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ChangeRequestEventSchema) SetAction(v string)`

SetAction sets Action field to given value.


### GetConflict

`func (o *ChangeRequestEventSchema) GetConflict() string`

GetConflict returns the Conflict field if non-nil, zero value otherwise.

### GetConflictOk

`func (o *ChangeRequestEventSchema) GetConflictOk() (*string, bool)`

GetConflictOk returns a tuple with the Conflict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConflict

`func (o *ChangeRequestEventSchema) SetConflict(v string)`

SetConflict sets Conflict field to given value.

### HasConflict

`func (o *ChangeRequestEventSchema) HasConflict() bool`

HasConflict returns a boolean if a field has been set.

### GetPayload

`func (o *ChangeRequestEventSchema) GetPayload() ChangeRequestEventSchemaPayload`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *ChangeRequestEventSchema) GetPayloadOk() (*ChangeRequestEventSchemaPayload, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *ChangeRequestEventSchema) SetPayload(v ChangeRequestEventSchemaPayload)`

SetPayload sets Payload field to given value.


### GetUpdatedBy

`func (o *ChangeRequestEventSchema) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *ChangeRequestEventSchema) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *ChangeRequestEventSchema) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *ChangeRequestEventSchema) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetCreatedBy

`func (o *ChangeRequestEventSchema) GetCreatedBy() ChangeRequestEventSchemaCreatedBy`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ChangeRequestEventSchema) GetCreatedByOk() (*ChangeRequestEventSchemaCreatedBy, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ChangeRequestEventSchema) SetCreatedBy(v ChangeRequestEventSchemaCreatedBy)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ChangeRequestEventSchema) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ChangeRequestEventSchema) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ChangeRequestEventSchema) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ChangeRequestEventSchema) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ChangeRequestEventSchema) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


