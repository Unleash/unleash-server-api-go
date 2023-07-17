# AdminFeaturesQuerySchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tag** | Pointer to **[]string** | Used to filter by tags. For each entry, a TAGTYPE:TAGVALUE is expected | [optional] 
**NamePrefix** | Pointer to **string** | A case-insensitive prefix filter for the names of feature toggles | [optional] 

## Methods

### NewAdminFeaturesQuerySchema

`func NewAdminFeaturesQuerySchema() *AdminFeaturesQuerySchema`

NewAdminFeaturesQuerySchema instantiates a new AdminFeaturesQuerySchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminFeaturesQuerySchemaWithDefaults

`func NewAdminFeaturesQuerySchemaWithDefaults() *AdminFeaturesQuerySchema`

NewAdminFeaturesQuerySchemaWithDefaults instantiates a new AdminFeaturesQuerySchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTag

`func (o *AdminFeaturesQuerySchema) GetTag() []string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *AdminFeaturesQuerySchema) GetTagOk() (*[]string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *AdminFeaturesQuerySchema) SetTag(v []string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *AdminFeaturesQuerySchema) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetNamePrefix

`func (o *AdminFeaturesQuerySchema) GetNamePrefix() string`

GetNamePrefix returns the NamePrefix field if non-nil, zero value otherwise.

### GetNamePrefixOk

`func (o *AdminFeaturesQuerySchema) GetNamePrefixOk() (*string, bool)`

GetNamePrefixOk returns a tuple with the NamePrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamePrefix

`func (o *AdminFeaturesQuerySchema) SetNamePrefix(v string)`

SetNamePrefix sets NamePrefix field to given value.

### HasNamePrefix

`func (o *AdminFeaturesQuerySchema) HasNamePrefix() bool`

HasNamePrefix returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


