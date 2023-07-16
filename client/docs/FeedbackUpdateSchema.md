# FeedbackUpdateSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | Pointer to **int32** | The ID of the user that gave the feedback. | [optional] 
**NeverShow** | Pointer to **bool** | &#x60;true&#x60; if the user has asked never to see this feedback questionnaire again. | [optional] 
**Given** | Pointer to **NullableTime** | When this feedback was given | [optional] 

## Methods

### NewFeedbackUpdateSchema

`func NewFeedbackUpdateSchema() *FeedbackUpdateSchema`

NewFeedbackUpdateSchema instantiates a new FeedbackUpdateSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeedbackUpdateSchemaWithDefaults

`func NewFeedbackUpdateSchemaWithDefaults() *FeedbackUpdateSchema`

NewFeedbackUpdateSchemaWithDefaults instantiates a new FeedbackUpdateSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *FeedbackUpdateSchema) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *FeedbackUpdateSchema) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *FeedbackUpdateSchema) SetUserId(v int32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *FeedbackUpdateSchema) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetNeverShow

`func (o *FeedbackUpdateSchema) GetNeverShow() bool`

GetNeverShow returns the NeverShow field if non-nil, zero value otherwise.

### GetNeverShowOk

`func (o *FeedbackUpdateSchema) GetNeverShowOk() (*bool, bool)`

GetNeverShowOk returns a tuple with the NeverShow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeverShow

`func (o *FeedbackUpdateSchema) SetNeverShow(v bool)`

SetNeverShow sets NeverShow field to given value.

### HasNeverShow

`func (o *FeedbackUpdateSchema) HasNeverShow() bool`

HasNeverShow returns a boolean if a field has been set.

### GetGiven

`func (o *FeedbackUpdateSchema) GetGiven() time.Time`

GetGiven returns the Given field if non-nil, zero value otherwise.

### GetGivenOk

`func (o *FeedbackUpdateSchema) GetGivenOk() (*time.Time, bool)`

GetGivenOk returns a tuple with the Given field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiven

`func (o *FeedbackUpdateSchema) SetGiven(v time.Time)`

SetGiven sets Given field to given value.

### HasGiven

`func (o *FeedbackUpdateSchema) HasGiven() bool`

HasGiven returns a boolean if a field has been set.

### SetGivenNil

`func (o *FeedbackUpdateSchema) SetGivenNil(b bool)`

 SetGivenNil sets the value for Given to be an explicit nil

### UnsetGiven
`func (o *FeedbackUpdateSchema) UnsetGiven()`

UnsetGiven ensures that no value is present for Given, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


