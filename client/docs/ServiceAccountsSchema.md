# ServiceAccountsSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceAccounts** | [**[]ServiceAccountSchema**](ServiceAccountSchema.md) | A list of service accounts | 
**RootRoles** | Pointer to [**[]RoleSchema**](RoleSchema.md) | A list of root roles that are referenced from service account objects in the &#x60;serviceAccounts&#x60; list | [optional] 

## Methods

### NewServiceAccountsSchema

`func NewServiceAccountsSchema(serviceAccounts []ServiceAccountSchema, ) *ServiceAccountsSchema`

NewServiceAccountsSchema instantiates a new ServiceAccountsSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceAccountsSchemaWithDefaults

`func NewServiceAccountsSchemaWithDefaults() *ServiceAccountsSchema`

NewServiceAccountsSchemaWithDefaults instantiates a new ServiceAccountsSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceAccounts

`func (o *ServiceAccountsSchema) GetServiceAccounts() []ServiceAccountSchema`

GetServiceAccounts returns the ServiceAccounts field if non-nil, zero value otherwise.

### GetServiceAccountsOk

`func (o *ServiceAccountsSchema) GetServiceAccountsOk() (*[]ServiceAccountSchema, bool)`

GetServiceAccountsOk returns a tuple with the ServiceAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccounts

`func (o *ServiceAccountsSchema) SetServiceAccounts(v []ServiceAccountSchema)`

SetServiceAccounts sets ServiceAccounts field to given value.


### GetRootRoles

`func (o *ServiceAccountsSchema) GetRootRoles() []RoleSchema`

GetRootRoles returns the RootRoles field if non-nil, zero value otherwise.

### GetRootRolesOk

`func (o *ServiceAccountsSchema) GetRootRolesOk() (*[]RoleSchema, bool)`

GetRootRolesOk returns a tuple with the RootRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootRoles

`func (o *ServiceAccountsSchema) SetRootRoles(v []RoleSchema)`

SetRootRoles sets RootRoles field to given value.

### HasRootRoles

`func (o *ServiceAccountsSchema) HasRootRoles() bool`

HasRootRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


