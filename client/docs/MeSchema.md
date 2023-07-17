# MeSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | [**UserSchema**](UserSchema.md) |  | 
**Permissions** | [**[]PermissionSchema**](PermissionSchema.md) | User permissions for projects and environments | 
**Feedback** | [**[]FeedbackResponseSchema**](FeedbackResponseSchema.md) | User feedback information | 
**Splash** | **map[string]bool** | Splash screen configuration | 

## Methods

### NewMeSchema

`func NewMeSchema(user UserSchema, permissions []PermissionSchema, feedback []FeedbackResponseSchema, splash map[string]bool, ) *MeSchema`

NewMeSchema instantiates a new MeSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMeSchemaWithDefaults

`func NewMeSchemaWithDefaults() *MeSchema`

NewMeSchemaWithDefaults instantiates a new MeSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *MeSchema) GetUser() UserSchema`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *MeSchema) GetUserOk() (*UserSchema, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *MeSchema) SetUser(v UserSchema)`

SetUser sets User field to given value.


### GetPermissions

`func (o *MeSchema) GetPermissions() []PermissionSchema`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *MeSchema) GetPermissionsOk() (*[]PermissionSchema, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *MeSchema) SetPermissions(v []PermissionSchema)`

SetPermissions sets Permissions field to given value.


### GetFeedback

`func (o *MeSchema) GetFeedback() []FeedbackResponseSchema`

GetFeedback returns the Feedback field if non-nil, zero value otherwise.

### GetFeedbackOk

`func (o *MeSchema) GetFeedbackOk() (*[]FeedbackResponseSchema, bool)`

GetFeedbackOk returns a tuple with the Feedback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedback

`func (o *MeSchema) SetFeedback(v []FeedbackResponseSchema)`

SetFeedback sets Feedback field to given value.


### GetSplash

`func (o *MeSchema) GetSplash() map[string]bool`

GetSplash returns the Splash field if non-nil, zero value otherwise.

### GetSplashOk

`func (o *MeSchema) GetSplashOk() (*map[string]bool, bool)`

GetSplashOk returns a tuple with the Splash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplash

`func (o *MeSchema) SetSplash(v map[string]bool)`

SetSplash sets Splash field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


