# NotificationsSchemaInner

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float64** | The id of this notification | [default to null]
**Message** | **string** | The actual notification message | [default to null]
**Link** | **string** | The link to change request or feature toggle the notification refers to | [default to null]
**NotificationType** | **string** | The type of the notification used e.g. for the graphical hints | [default to null]
**CreatedBy** | [***interface{}**](interface{}.md) |  | [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | The date and time when the notification was created | [default to null]
**ReadAt** | [**time.Time**](time.Time.md) | The date and time when the notification was read or marked as read, otherwise &#x60;null&#x60; | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

