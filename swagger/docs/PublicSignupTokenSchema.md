# PublicSignupTokenSchema

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Secret** | **string** | The actual value of the token. This is the part that is used by Unleash to create an invite link | [default to null]
**Url** | **string** | The public signup link for the token. Users who follow this link will be taken to a signup page where they can create an Unleash user. | [default to null]
**Name** | **string** | The token&#x27;s name. Only for displaying in the UI | [default to null]
**Enabled** | **bool** | Whether the token is active. This property will always be &#x60;false&#x60; for a token that has expired. | [default to null]
**ExpiresAt** | [**time.Time**](time.Time.md) | The time when the token will expire. | [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | When the token was created. | [default to null]
**CreatedBy** | **string** | The creator&#x27;s email or username | [default to null]
**Users** | [**[]UserSchema**](userSchema.md) | Array of users that have signed up using the token. | [optional] [default to null]
**Role** | [***RoleSchema**](roleSchema.md) |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

