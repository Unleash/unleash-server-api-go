# LoginEventSchema

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The event&#x27;s ID. Event IDs are incrementing integers. In other words, a more recent event will always have a higher ID than an older event. | [default to null]
**Username** | **string** | The username of the user that attempted to log in. Will return \&quot;Incorrectly configured provider\&quot; when attempting to log in using a misconfigured provider. | [optional] [default to null]
**AuthType** | **string** | The authentication type used to log in. | [optional] [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | The date and time of when the login was attempted. | [optional] [default to null]
**Successful** | **bool** | Whether the login was successful or not. | [optional] [default to null]
**Ip** | **string** | The IP address of the client that attempted to log in. | [optional] [default to null]
**FailureReason** | **string** | The reason for the login failure. This property is only present if the login was unsuccessful. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

