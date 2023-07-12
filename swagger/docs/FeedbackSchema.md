# FeedbackSchema

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **int32** | Identifier of the current user giving feedback | [optional] [default to null]
**FeedbackId** | **string** | The name of the feedback session | [optional] [default to null]
**NeverShow** | **bool** | &#x60;true&#x60; when user opts-in to never show the feedback again. | [optional] [default to null]
**Given** | [**time.Time**](time.Time.md) | When this feedback was given | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

