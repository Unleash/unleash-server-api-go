# ChangeRequestCommentSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **float32** |  | [optional] 
**Text** | **string** |  | 
**CreatedBy** | [**ChangeRequestCommentSchemaCreatedBy**](ChangeRequestCommentSchemaCreatedBy.md) |  | 
**CreatedAt** | **time.Time** |  | 

## Methods

### NewChangeRequestCommentSchema

`func NewChangeRequestCommentSchema(text string, createdBy ChangeRequestCommentSchemaCreatedBy, createdAt time.Time, ) *ChangeRequestCommentSchema`

NewChangeRequestCommentSchema instantiates a new ChangeRequestCommentSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChangeRequestCommentSchemaWithDefaults

`func NewChangeRequestCommentSchemaWithDefaults() *ChangeRequestCommentSchema`

NewChangeRequestCommentSchemaWithDefaults instantiates a new ChangeRequestCommentSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ChangeRequestCommentSchema) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ChangeRequestCommentSchema) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ChangeRequestCommentSchema) SetId(v float32)`

SetId sets Id field to given value.

### HasId

`func (o *ChangeRequestCommentSchema) HasId() bool`

HasId returns a boolean if a field has been set.

### GetText

`func (o *ChangeRequestCommentSchema) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *ChangeRequestCommentSchema) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *ChangeRequestCommentSchema) SetText(v string)`

SetText sets Text field to given value.


### GetCreatedBy

`func (o *ChangeRequestCommentSchema) GetCreatedBy() ChangeRequestCommentSchemaCreatedBy`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ChangeRequestCommentSchema) GetCreatedByOk() (*ChangeRequestCommentSchemaCreatedBy, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ChangeRequestCommentSchema) SetCreatedBy(v ChangeRequestCommentSchemaCreatedBy)`

SetCreatedBy sets CreatedBy field to given value.


### GetCreatedAt

`func (o *ChangeRequestCommentSchema) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ChangeRequestCommentSchema) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ChangeRequestCommentSchema) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


