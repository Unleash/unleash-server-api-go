# UpdateProjectEnterpriseSettingsSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | Pointer to **string** | A mode of the project affecting what actions are possible in this project | [optional] 
**FeatureNaming** | Pointer to [**CreateFeatureNamingPatternSchema**](CreateFeatureNamingPatternSchema.md) |  | [optional] 
**LinkTemplates** | Pointer to [**[]ProjectLinkTemplateSchema**](ProjectLinkTemplateSchema.md) | A list of link templates that can be automatically added to new feature flags. | [optional] 

## Methods

### NewUpdateProjectEnterpriseSettingsSchema

`func NewUpdateProjectEnterpriseSettingsSchema() *UpdateProjectEnterpriseSettingsSchema`

NewUpdateProjectEnterpriseSettingsSchema instantiates a new UpdateProjectEnterpriseSettingsSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateProjectEnterpriseSettingsSchemaWithDefaults

`func NewUpdateProjectEnterpriseSettingsSchemaWithDefaults() *UpdateProjectEnterpriseSettingsSchema`

NewUpdateProjectEnterpriseSettingsSchemaWithDefaults instantiates a new UpdateProjectEnterpriseSettingsSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *UpdateProjectEnterpriseSettingsSchema) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *UpdateProjectEnterpriseSettingsSchema) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *UpdateProjectEnterpriseSettingsSchema) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *UpdateProjectEnterpriseSettingsSchema) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetFeatureNaming

`func (o *UpdateProjectEnterpriseSettingsSchema) GetFeatureNaming() CreateFeatureNamingPatternSchema`

GetFeatureNaming returns the FeatureNaming field if non-nil, zero value otherwise.

### GetFeatureNamingOk

`func (o *UpdateProjectEnterpriseSettingsSchema) GetFeatureNamingOk() (*CreateFeatureNamingPatternSchema, bool)`

GetFeatureNamingOk returns a tuple with the FeatureNaming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureNaming

`func (o *UpdateProjectEnterpriseSettingsSchema) SetFeatureNaming(v CreateFeatureNamingPatternSchema)`

SetFeatureNaming sets FeatureNaming field to given value.

### HasFeatureNaming

`func (o *UpdateProjectEnterpriseSettingsSchema) HasFeatureNaming() bool`

HasFeatureNaming returns a boolean if a field has been set.

### GetLinkTemplates

`func (o *UpdateProjectEnterpriseSettingsSchema) GetLinkTemplates() []ProjectLinkTemplateSchema`

GetLinkTemplates returns the LinkTemplates field if non-nil, zero value otherwise.

### GetLinkTemplatesOk

`func (o *UpdateProjectEnterpriseSettingsSchema) GetLinkTemplatesOk() (*[]ProjectLinkTemplateSchema, bool)`

GetLinkTemplatesOk returns a tuple with the LinkTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkTemplates

`func (o *UpdateProjectEnterpriseSettingsSchema) SetLinkTemplates(v []ProjectLinkTemplateSchema)`

SetLinkTemplates sets LinkTemplates field to given value.

### HasLinkTemplates

`func (o *UpdateProjectEnterpriseSettingsSchema) HasLinkTemplates() bool`

HasLinkTemplates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


