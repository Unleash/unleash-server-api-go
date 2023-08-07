# FeedbackCreateSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NeverShow** | Pointer to **bool** | &#x60;true&#x60; if the user has asked never to see this feedback questionnaire again. Defaults to &#x60;false&#x60;. | [optional] 
**FeedbackId** | **string** | The name of the feedback session | 

## Methods

### NewFeedbackCreateSchema

`func NewFeedbackCreateSchema(feedbackId string, ) *FeedbackCreateSchema`

NewFeedbackCreateSchema instantiates a new FeedbackCreateSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeedbackCreateSchemaWithDefaults

`func NewFeedbackCreateSchemaWithDefaults() *FeedbackCreateSchema`

NewFeedbackCreateSchemaWithDefaults instantiates a new FeedbackCreateSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNeverShow

`func (o *FeedbackCreateSchema) GetNeverShow() bool`

GetNeverShow returns the NeverShow field if non-nil, zero value otherwise.

### GetNeverShowOk

`func (o *FeedbackCreateSchema) GetNeverShowOk() (*bool, bool)`

GetNeverShowOk returns a tuple with the NeverShow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeverShow

`func (o *FeedbackCreateSchema) SetNeverShow(v bool)`

SetNeverShow sets NeverShow field to given value.

### HasNeverShow

`func (o *FeedbackCreateSchema) HasNeverShow() bool`

HasNeverShow returns a boolean if a field has been set.

### GetFeedbackId

`func (o *FeedbackCreateSchema) GetFeedbackId() string`

GetFeedbackId returns the FeedbackId field if non-nil, zero value otherwise.

### GetFeedbackIdOk

`func (o *FeedbackCreateSchema) GetFeedbackIdOk() (*string, bool)`

GetFeedbackIdOk returns a tuple with the FeedbackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedbackId

`func (o *FeedbackCreateSchema) SetFeedbackId(v string)`

SetFeedbackId sets FeedbackId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


