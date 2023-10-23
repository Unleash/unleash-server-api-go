# VersionSchemaCurrent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Oss** | Pointer to **string** | The OSS version used when building this Unleash instance, represented as a git revision belonging to the [main Unleash git repo](https://github.com/Unleash/unleash/) | [optional] 
**Enterprise** | Pointer to **string** | The Enterpris version of Unleash used to build this instance, represented as a git revision belonging to the [Unleash Enterprise](https://github.com/ivarconr/unleash-enterprise) repository. Will be an empty string if no enterprise version was used, | [optional] 

## Methods

### NewVersionSchemaCurrent

`func NewVersionSchemaCurrent() *VersionSchemaCurrent`

NewVersionSchemaCurrent instantiates a new VersionSchemaCurrent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVersionSchemaCurrentWithDefaults

`func NewVersionSchemaCurrentWithDefaults() *VersionSchemaCurrent`

NewVersionSchemaCurrentWithDefaults instantiates a new VersionSchemaCurrent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOss

`func (o *VersionSchemaCurrent) GetOss() string`

GetOss returns the Oss field if non-nil, zero value otherwise.

### GetOssOk

`func (o *VersionSchemaCurrent) GetOssOk() (*string, bool)`

GetOssOk returns a tuple with the Oss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOss

`func (o *VersionSchemaCurrent) SetOss(v string)`

SetOss sets Oss field to given value.

### HasOss

`func (o *VersionSchemaCurrent) HasOss() bool`

HasOss returns a boolean if a field has been set.

### GetEnterprise

`func (o *VersionSchemaCurrent) GetEnterprise() string`

GetEnterprise returns the Enterprise field if non-nil, zero value otherwise.

### GetEnterpriseOk

`func (o *VersionSchemaCurrent) GetEnterpriseOk() (*string, bool)`

GetEnterpriseOk returns a tuple with the Enterprise field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnterprise

`func (o *VersionSchemaCurrent) SetEnterprise(v string)`

SetEnterprise sets Enterprise field to given value.

### HasEnterprise

`func (o *VersionSchemaCurrent) HasEnterprise() bool`

HasEnterprise returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


