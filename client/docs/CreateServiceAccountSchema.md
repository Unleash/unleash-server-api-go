# CreateServiceAccountSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** | The username of the service account | 
**Name** | **string** | The name of the service account | 
**RootRole** | **int32** | The id of the root role for the service account | 

## Methods

### NewCreateServiceAccountSchema

`func NewCreateServiceAccountSchema(username string, name string, rootRole int32, ) *CreateServiceAccountSchema`

NewCreateServiceAccountSchema instantiates a new CreateServiceAccountSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateServiceAccountSchemaWithDefaults

`func NewCreateServiceAccountSchemaWithDefaults() *CreateServiceAccountSchema`

NewCreateServiceAccountSchemaWithDefaults instantiates a new CreateServiceAccountSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *CreateServiceAccountSchema) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CreateServiceAccountSchema) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CreateServiceAccountSchema) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetName

`func (o *CreateServiceAccountSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateServiceAccountSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateServiceAccountSchema) SetName(v string)`

SetName sets Name field to given value.


### GetRootRole

`func (o *CreateServiceAccountSchema) GetRootRole() int32`

GetRootRole returns the RootRole field if non-nil, zero value otherwise.

### GetRootRoleOk

`func (o *CreateServiceAccountSchema) GetRootRoleOk() (*int32, bool)`

GetRootRoleOk returns a tuple with the RootRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootRole

`func (o *CreateServiceAccountSchema) SetRootRole(v int32)`

SetRootRole sets RootRole field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


