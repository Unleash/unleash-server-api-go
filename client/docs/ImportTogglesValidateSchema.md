# ImportTogglesValidateSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | [**[]ImportTogglesValidateItemSchema**](ImportTogglesValidateItemSchema.md) | A list of errors that prevent the provided data from being successfully imported. | 
**Warnings** | [**[]ImportTogglesValidateItemSchema**](ImportTogglesValidateItemSchema.md) | A list of warnings related to the provided data. | 
**Permissions** | Pointer to [**[]ImportTogglesValidateItemSchema**](ImportTogglesValidateItemSchema.md) | Any additional permissions required to import the data. If the list is empty, you require no additional permissions beyond what your user already has. | [optional] 

## Methods

### NewImportTogglesValidateSchema

`func NewImportTogglesValidateSchema(errors []ImportTogglesValidateItemSchema, warnings []ImportTogglesValidateItemSchema, ) *ImportTogglesValidateSchema`

NewImportTogglesValidateSchema instantiates a new ImportTogglesValidateSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportTogglesValidateSchemaWithDefaults

`func NewImportTogglesValidateSchemaWithDefaults() *ImportTogglesValidateSchema`

NewImportTogglesValidateSchemaWithDefaults instantiates a new ImportTogglesValidateSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *ImportTogglesValidateSchema) GetErrors() []ImportTogglesValidateItemSchema`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ImportTogglesValidateSchema) GetErrorsOk() (*[]ImportTogglesValidateItemSchema, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ImportTogglesValidateSchema) SetErrors(v []ImportTogglesValidateItemSchema)`

SetErrors sets Errors field to given value.


### GetWarnings

`func (o *ImportTogglesValidateSchema) GetWarnings() []ImportTogglesValidateItemSchema`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *ImportTogglesValidateSchema) GetWarningsOk() (*[]ImportTogglesValidateItemSchema, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *ImportTogglesValidateSchema) SetWarnings(v []ImportTogglesValidateItemSchema)`

SetWarnings sets Warnings field to given value.


### GetPermissions

`func (o *ImportTogglesValidateSchema) GetPermissions() []ImportTogglesValidateItemSchema`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *ImportTogglesValidateSchema) GetPermissionsOk() (*[]ImportTogglesValidateItemSchema, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *ImportTogglesValidateSchema) SetPermissions(v []ImportTogglesValidateItemSchema)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *ImportTogglesValidateSchema) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


