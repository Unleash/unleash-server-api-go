# CreatePatSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | The PAT&#39;s description. | 
**ExpiresAt** | **time.Time** | The PAT&#39;s expiration date. | 

## Methods

### NewCreatePatSchema

`func NewCreatePatSchema(description string, expiresAt time.Time, ) *CreatePatSchema`

NewCreatePatSchema instantiates a new CreatePatSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePatSchemaWithDefaults

`func NewCreatePatSchemaWithDefaults() *CreatePatSchema`

NewCreatePatSchemaWithDefaults instantiates a new CreatePatSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CreatePatSchema) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreatePatSchema) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreatePatSchema) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetExpiresAt

`func (o *CreatePatSchema) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *CreatePatSchema) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *CreatePatSchema) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


