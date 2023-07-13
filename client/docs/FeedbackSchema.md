# FeedbackSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | Pointer to **int32** | Identifier of the current user giving feedback | [optional] 
**FeedbackId** | Pointer to **string** | The name of the feedback session | [optional] 
**NeverShow** | Pointer to **bool** | &#x60;true&#x60; when user opts-in to never show the feedback again. | [optional] 
**Given** | Pointer to **NullableTime** | When this feedback was given | [optional] 

## Methods

### NewFeedbackSchema

`func NewFeedbackSchema() *FeedbackSchema`

NewFeedbackSchema instantiates a new FeedbackSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeedbackSchemaWithDefaults

`func NewFeedbackSchemaWithDefaults() *FeedbackSchema`

NewFeedbackSchemaWithDefaults instantiates a new FeedbackSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *FeedbackSchema) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *FeedbackSchema) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *FeedbackSchema) SetUserId(v int32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *FeedbackSchema) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetFeedbackId

`func (o *FeedbackSchema) GetFeedbackId() string`

GetFeedbackId returns the FeedbackId field if non-nil, zero value otherwise.

### GetFeedbackIdOk

`func (o *FeedbackSchema) GetFeedbackIdOk() (*string, bool)`

GetFeedbackIdOk returns a tuple with the FeedbackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedbackId

`func (o *FeedbackSchema) SetFeedbackId(v string)`

SetFeedbackId sets FeedbackId field to given value.

### HasFeedbackId

`func (o *FeedbackSchema) HasFeedbackId() bool`

HasFeedbackId returns a boolean if a field has been set.

### GetNeverShow

`func (o *FeedbackSchema) GetNeverShow() bool`

GetNeverShow returns the NeverShow field if non-nil, zero value otherwise.

### GetNeverShowOk

`func (o *FeedbackSchema) GetNeverShowOk() (*bool, bool)`

GetNeverShowOk returns a tuple with the NeverShow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeverShow

`func (o *FeedbackSchema) SetNeverShow(v bool)`

SetNeverShow sets NeverShow field to given value.

### HasNeverShow

`func (o *FeedbackSchema) HasNeverShow() bool`

HasNeverShow returns a boolean if a field has been set.

### GetGiven

`func (o *FeedbackSchema) GetGiven() time.Time`

GetGiven returns the Given field if non-nil, zero value otherwise.

### GetGivenOk

`func (o *FeedbackSchema) GetGivenOk() (*time.Time, bool)`

GetGivenOk returns a tuple with the Given field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiven

`func (o *FeedbackSchema) SetGiven(v time.Time)`

SetGiven sets Given field to given value.

### HasGiven

`func (o *FeedbackSchema) HasGiven() bool`

HasGiven returns a boolean if a field has been set.

### SetGivenNil

`func (o *FeedbackSchema) SetGivenNil(b bool)`

 SetGivenNil sets the value for Given to be an explicit nil

### UnsetGiven
`func (o *FeedbackSchema) UnsetGiven()`

UnsetGiven ensures that no value is present for Given, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


