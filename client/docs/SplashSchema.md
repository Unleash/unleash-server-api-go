# SplashSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **float32** |  | 
**SplashId** | **string** |  | 
**Seen** | **bool** |  | 

## Methods

### NewSplashSchema

`func NewSplashSchema(userId float32, splashId string, seen bool, ) *SplashSchema`

NewSplashSchema instantiates a new SplashSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSplashSchemaWithDefaults

`func NewSplashSchemaWithDefaults() *SplashSchema`

NewSplashSchemaWithDefaults instantiates a new SplashSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *SplashSchema) GetUserId() float32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *SplashSchema) GetUserIdOk() (*float32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *SplashSchema) SetUserId(v float32)`

SetUserId sets UserId field to given value.


### GetSplashId

`func (o *SplashSchema) GetSplashId() string`

GetSplashId returns the SplashId field if non-nil, zero value otherwise.

### GetSplashIdOk

`func (o *SplashSchema) GetSplashIdOk() (*string, bool)`

GetSplashIdOk returns a tuple with the SplashId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplashId

`func (o *SplashSchema) SetSplashId(v string)`

SetSplashId sets SplashId field to given value.


### GetSeen

`func (o *SplashSchema) GetSeen() bool`

GetSeen returns the Seen field if non-nil, zero value otherwise.

### GetSeenOk

`func (o *SplashSchema) GetSeenOk() (*bool, bool)`

GetSeenOk returns a tuple with the Seen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeen

`func (o *SplashSchema) SetSeen(v bool)`

SetSeen sets Seen field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


