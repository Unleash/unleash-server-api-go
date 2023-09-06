# ChangeRequestStateSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | **string** | The new desired state for the change request | 
**Comment** | Pointer to **string** | Any comments accompanying the state changed. Used when sending a draft to review. | [optional] 

## Methods

### NewChangeRequestStateSchema

`func NewChangeRequestStateSchema(state string, ) *ChangeRequestStateSchema`

NewChangeRequestStateSchema instantiates a new ChangeRequestStateSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChangeRequestStateSchemaWithDefaults

`func NewChangeRequestStateSchemaWithDefaults() *ChangeRequestStateSchema`

NewChangeRequestStateSchemaWithDefaults instantiates a new ChangeRequestStateSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *ChangeRequestStateSchema) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ChangeRequestStateSchema) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ChangeRequestStateSchema) SetState(v string)`

SetState sets State field to given value.


### GetComment

`func (o *ChangeRequestStateSchema) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ChangeRequestStateSchema) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ChangeRequestStateSchema) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ChangeRequestStateSchema) HasComment() bool`

HasComment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


