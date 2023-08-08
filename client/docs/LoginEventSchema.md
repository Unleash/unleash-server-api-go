# LoginEventSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The event&#39;s ID. Event IDs are incrementing integers. In other words, a more recent event will always have a higher ID than an older event. | 
**Username** | Pointer to **string** | The username of the user that attempted to log in. Will return \&quot;Incorrectly configured provider\&quot; when attempting to log in using a misconfigured provider. | [optional] 
**AuthType** | Pointer to **string** | The authentication type used to log in. | [optional] 
**CreatedAt** | Pointer to **time.Time** | The date and time of when the login was attempted. | [optional] 
**Successful** | Pointer to **bool** | Whether the login was successful or not. | [optional] 
**Ip** | Pointer to **NullableString** | The IP address of the client that attempted to log in. | [optional] 
**FailureReason** | Pointer to **NullableString** | The reason for the login failure. This property is only present if the login was unsuccessful. | [optional] 

## Methods

### NewLoginEventSchema

`func NewLoginEventSchema(id int32, ) *LoginEventSchema`

NewLoginEventSchema instantiates a new LoginEventSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoginEventSchemaWithDefaults

`func NewLoginEventSchemaWithDefaults() *LoginEventSchema`

NewLoginEventSchemaWithDefaults instantiates a new LoginEventSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LoginEventSchema) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LoginEventSchema) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LoginEventSchema) SetId(v int32)`

SetId sets Id field to given value.


### GetUsername

`func (o *LoginEventSchema) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *LoginEventSchema) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *LoginEventSchema) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *LoginEventSchema) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetAuthType

`func (o *LoginEventSchema) GetAuthType() string`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *LoginEventSchema) GetAuthTypeOk() (*string, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *LoginEventSchema) SetAuthType(v string)`

SetAuthType sets AuthType field to given value.

### HasAuthType

`func (o *LoginEventSchema) HasAuthType() bool`

HasAuthType returns a boolean if a field has been set.

### GetCreatedAt

`func (o *LoginEventSchema) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *LoginEventSchema) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *LoginEventSchema) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *LoginEventSchema) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetSuccessful

`func (o *LoginEventSchema) GetSuccessful() bool`

GetSuccessful returns the Successful field if non-nil, zero value otherwise.

### GetSuccessfulOk

`func (o *LoginEventSchema) GetSuccessfulOk() (*bool, bool)`

GetSuccessfulOk returns a tuple with the Successful field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessful

`func (o *LoginEventSchema) SetSuccessful(v bool)`

SetSuccessful sets Successful field to given value.

### HasSuccessful

`func (o *LoginEventSchema) HasSuccessful() bool`

HasSuccessful returns a boolean if a field has been set.

### GetIp

`func (o *LoginEventSchema) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *LoginEventSchema) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *LoginEventSchema) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *LoginEventSchema) HasIp() bool`

HasIp returns a boolean if a field has been set.

### SetIpNil

`func (o *LoginEventSchema) SetIpNil(b bool)`

 SetIpNil sets the value for Ip to be an explicit nil

### UnsetIp
`func (o *LoginEventSchema) UnsetIp()`

UnsetIp ensures that no value is present for Ip, not even an explicit nil
### GetFailureReason

`func (o *LoginEventSchema) GetFailureReason() string`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *LoginEventSchema) GetFailureReasonOk() (*string, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *LoginEventSchema) SetFailureReason(v string)`

SetFailureReason sets FailureReason field to given value.

### HasFailureReason

`func (o *LoginEventSchema) HasFailureReason() bool`

HasFailureReason returns a boolean if a field has been set.

### SetFailureReasonNil

`func (o *LoginEventSchema) SetFailureReasonNil(b bool)`

 SetFailureReasonNil sets the value for FailureReason to be an explicit nil

### UnsetFailureReason
`func (o *LoginEventSchema) UnsetFailureReason()`

UnsetFailureReason ensures that no value is present for FailureReason, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


