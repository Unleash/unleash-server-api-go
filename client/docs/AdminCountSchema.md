# AdminCountSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | **float32** | Total number of admins that have a password set. | 
**NoPassword** | **float32** | Total number of admins that do not have a password set. May be SSO, but may also be users that did not set a password yet. | 
**Service** | **float32** | Total number of service accounts that have the admin root role. | 

## Methods

### NewAdminCountSchema

`func NewAdminCountSchema(password float32, noPassword float32, service float32, ) *AdminCountSchema`

NewAdminCountSchema instantiates a new AdminCountSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminCountSchemaWithDefaults

`func NewAdminCountSchemaWithDefaults() *AdminCountSchema`

NewAdminCountSchemaWithDefaults instantiates a new AdminCountSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *AdminCountSchema) GetPassword() float32`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *AdminCountSchema) GetPasswordOk() (*float32, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *AdminCountSchema) SetPassword(v float32)`

SetPassword sets Password field to given value.


### GetNoPassword

`func (o *AdminCountSchema) GetNoPassword() float32`

GetNoPassword returns the NoPassword field if non-nil, zero value otherwise.

### GetNoPasswordOk

`func (o *AdminCountSchema) GetNoPasswordOk() (*float32, bool)`

GetNoPasswordOk returns a tuple with the NoPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoPassword

`func (o *AdminCountSchema) SetNoPassword(v float32)`

SetNoPassword sets NoPassword field to given value.


### GetService

`func (o *AdminCountSchema) GetService() float32`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *AdminCountSchema) GetServiceOk() (*float32, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *AdminCountSchema) SetService(v float32)`

SetService sets Service field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


