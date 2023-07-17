# SearchEventsSchema

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type_** | **string** | Find events by event type (case-sensitive). | [optional] [default to null]
**Project** | **string** | Find events by project ID (case-sensitive). | [optional] [default to null]
**Feature** | **string** | Find events by feature toggle name (case-sensitive). | [optional] [default to null]
**Query** | **string** |                  Find events by a free-text search query.                 The query will be matched against the event type,                 the username or email that created the event (if any),                 and the event data payload (if any).              | [optional] [default to null]
**Limit** | **int32** | The maximum amount of events to return in the search result | [optional] [default to 100]
**Offset** | **int32** | Which event id to start listing from | [optional] [default to 0]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

