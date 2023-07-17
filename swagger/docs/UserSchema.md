# UserSchema

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The user id | [default to null]
**IsAPI** | **bool** | (Deprecated): Used internally to know which operations the user should be allowed to perform | [optional] [default to null]
**Name** | **string** | Name of the user | [optional] [default to null]
**Email** | **string** | Email of the user | [optional] [default to null]
**Username** | **string** | A unique username for the user | [optional] [default to null]
**ImageUrl** | **string** | URL used for the userprofile image | [optional] [default to null]
**InviteLink** | **string** | If the user is actively inviting other users, this is the link that can be shared with other users | [optional] [default to null]
**LoginAttempts** | **int32** | How many unsuccessful attempts at logging in has the user made | [optional] [default to null]
**EmailSent** | **bool** | Is the welcome email sent to the user or not | [optional] [default to null]
**RootRole** | **int32** | Which [root role](https://docs.getunleash.io/reference/rbac#standard-roles) this user is assigned | [optional] [default to null]
**SeenAt** | [**time.Time**](time.Time.md) | The last time this user logged in | [optional] [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | The user was created at this time | [optional] [default to null]
**AccountType** | **string** | A user is either an actual User or a Service Account | [optional] [default to null]
**Permissions** | **[]string** | Deprecated | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

