# FeedbackResponseSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | Pointer to **int32** | The ID of the user that gave the feedback. | [optional] 
**NeverShow** | Pointer to **bool** | &#x60;true&#x60; if the user has asked never to see this feedback questionnaire again. | [optional] 
**Given** | Pointer to **NullableTime** | When this feedback was given | [optional] 
**FeedbackId** | Pointer to **string** | The name of the feedback session | [optional] 

## Methods

### NewFeedbackResponseSchema

`func NewFeedbackResponseSchema() *FeedbackResponseSchema`

NewFeedbackResponseSchema instantiates a new FeedbackResponseSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeedbackResponseSchemaWithDefaults

`func NewFeedbackResponseSchemaWithDefaults() *FeedbackResponseSchema`

NewFeedbackResponseSchemaWithDefaults instantiates a new FeedbackResponseSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *FeedbackResponseSchema) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *FeedbackResponseSchema) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *FeedbackResponseSchema) SetUserId(v int32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *FeedbackResponseSchema) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetNeverShow

`func (o *FeedbackResponseSchema) GetNeverShow() bool`

GetNeverShow returns the NeverShow field if non-nil, zero value otherwise.

### GetNeverShowOk

`func (o *FeedbackResponseSchema) GetNeverShowOk() (*bool, bool)`

GetNeverShowOk returns a tuple with the NeverShow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeverShow

`func (o *FeedbackResponseSchema) SetNeverShow(v bool)`

SetNeverShow sets NeverShow field to given value.

### HasNeverShow

`func (o *FeedbackResponseSchema) HasNeverShow() bool`

HasNeverShow returns a boolean if a field has been set.

### GetGiven

`func (o *FeedbackResponseSchema) GetGiven() time.Time`

GetGiven returns the Given field if non-nil, zero value otherwise.

### GetGivenOk

`func (o *FeedbackResponseSchema) GetGivenOk() (*time.Time, bool)`

GetGivenOk returns a tuple with the Given field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiven

`func (o *FeedbackResponseSchema) SetGiven(v time.Time)`

SetGiven sets Given field to given value.

### HasGiven

`func (o *FeedbackResponseSchema) HasGiven() bool`

HasGiven returns a boolean if a field has been set.

### SetGivenNil

`func (o *FeedbackResponseSchema) SetGivenNil(b bool)`

 SetGivenNil sets the value for Given to be an explicit nil

### UnsetGiven
`func (o *FeedbackResponseSchema) UnsetGiven()`

UnsetGiven ensures that no value is present for Given, not even an explicit nil
### GetFeedbackId

`func (o *FeedbackResponseSchema) GetFeedbackId() string`

GetFeedbackId returns the FeedbackId field if non-nil, zero value otherwise.

### GetFeedbackIdOk

`func (o *FeedbackResponseSchema) GetFeedbackIdOk() (*string, bool)`

GetFeedbackIdOk returns a tuple with the FeedbackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedbackId

`func (o *FeedbackResponseSchema) SetFeedbackId(v string)`

SetFeedbackId sets FeedbackId field to given value.

### HasFeedbackId

`func (o *FeedbackResponseSchema) HasFeedbackId() bool`

HasFeedbackId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


