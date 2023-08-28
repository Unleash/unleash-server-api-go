# SdkContextSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppName** | **string** | The name of the application. | 
**CurrentTime** | Pointer to **time.Time** | A DateTime (or similar) data class instance or a string in an RFC3339-compatible format. Defaults to the current time if not set by the user. | [optional] 
**Environment** | Pointer to **string** | The environment the app is running in. | [optional] 
**Properties** | Pointer to **map[string]string** | Additional Unleash context properties | [optional] 
**RemoteAddress** | Pointer to **string** | The app&#39;s IP address | [optional] 
**SessionId** | Pointer to **string** | An identifier for the current session | [optional] 
**UserId** | Pointer to **string** | An identifier for the current user | [optional] 

## Methods

### NewSdkContextSchema

`func NewSdkContextSchema(appName string, ) *SdkContextSchema`

NewSdkContextSchema instantiates a new SdkContextSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSdkContextSchemaWithDefaults

`func NewSdkContextSchemaWithDefaults() *SdkContextSchema`

NewSdkContextSchemaWithDefaults instantiates a new SdkContextSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppName

`func (o *SdkContextSchema) GetAppName() string`

GetAppName returns the AppName field if non-nil, zero value otherwise.

### GetAppNameOk

`func (o *SdkContextSchema) GetAppNameOk() (*string, bool)`

GetAppNameOk returns a tuple with the AppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppName

`func (o *SdkContextSchema) SetAppName(v string)`

SetAppName sets AppName field to given value.


### GetCurrentTime

`func (o *SdkContextSchema) GetCurrentTime() time.Time`

GetCurrentTime returns the CurrentTime field if non-nil, zero value otherwise.

### GetCurrentTimeOk

`func (o *SdkContextSchema) GetCurrentTimeOk() (*time.Time, bool)`

GetCurrentTimeOk returns a tuple with the CurrentTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentTime

`func (o *SdkContextSchema) SetCurrentTime(v time.Time)`

SetCurrentTime sets CurrentTime field to given value.

### HasCurrentTime

`func (o *SdkContextSchema) HasCurrentTime() bool`

HasCurrentTime returns a boolean if a field has been set.

### GetEnvironment

`func (o *SdkContextSchema) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *SdkContextSchema) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *SdkContextSchema) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *SdkContextSchema) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetProperties

`func (o *SdkContextSchema) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *SdkContextSchema) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *SdkContextSchema) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *SdkContextSchema) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetRemoteAddress

`func (o *SdkContextSchema) GetRemoteAddress() string`

GetRemoteAddress returns the RemoteAddress field if non-nil, zero value otherwise.

### GetRemoteAddressOk

`func (o *SdkContextSchema) GetRemoteAddressOk() (*string, bool)`

GetRemoteAddressOk returns a tuple with the RemoteAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAddress

`func (o *SdkContextSchema) SetRemoteAddress(v string)`

SetRemoteAddress sets RemoteAddress field to given value.

### HasRemoteAddress

`func (o *SdkContextSchema) HasRemoteAddress() bool`

HasRemoteAddress returns a boolean if a field has been set.

### GetSessionId

`func (o *SdkContextSchema) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *SdkContextSchema) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *SdkContextSchema) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *SdkContextSchema) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetUserId

`func (o *SdkContextSchema) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *SdkContextSchema) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *SdkContextSchema) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *SdkContextSchema) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


