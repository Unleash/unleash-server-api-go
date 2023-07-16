# SplashRequestSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **int32** | The ID of the user that was shown the splash screen. | 
**SplashId** | **string** | The ID of the splash screen that was shown. | 

## Methods

### NewSplashRequestSchema

`func NewSplashRequestSchema(userId int32, splashId string, ) *SplashRequestSchema`

NewSplashRequestSchema instantiates a new SplashRequestSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSplashRequestSchemaWithDefaults

`func NewSplashRequestSchemaWithDefaults() *SplashRequestSchema`

NewSplashRequestSchemaWithDefaults instantiates a new SplashRequestSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *SplashRequestSchema) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *SplashRequestSchema) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *SplashRequestSchema) SetUserId(v int32)`

SetUserId sets UserId field to given value.


### GetSplashId

`func (o *SplashRequestSchema) GetSplashId() string`

GetSplashId returns the SplashId field if non-nil, zero value otherwise.

### GetSplashIdOk

`func (o *SplashRequestSchema) GetSplashIdOk() (*string, bool)`

GetSplashIdOk returns a tuple with the SplashId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplashId

`func (o *SplashRequestSchema) SetSplashId(v string)`

SetSplashId sets SplashId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


