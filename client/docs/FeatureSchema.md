# FeatureSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Unique feature name | 
**Type** | Pointer to **string** | Type of the toggle e.g. experiment, kill-switch, release, operational, permission | [optional] 
**Description** | Pointer to **NullableString** | Detailed description of the feature | [optional] 
**Archived** | Pointer to **bool** | &#x60;true&#x60; if the feature is archived | [optional] 
**Project** | Pointer to **string** | Name of the project the feature belongs to | [optional] 
**Enabled** | Pointer to **bool** | &#x60;true&#x60; if the feature is enabled, otherwise &#x60;false&#x60;. | [optional] 
**Stale** | Pointer to **bool** | &#x60;true&#x60; if the feature is stale based on the age and feature type, otherwise &#x60;false&#x60;. | [optional] 
**Favorite** | Pointer to **bool** | &#x60;true&#x60; if the feature was favorited, otherwise &#x60;false&#x60;. | [optional] 
**ImpressionData** | Pointer to **bool** | &#x60;true&#x60; if the impression data collection is enabled for the feature, otherwise &#x60;false&#x60;. | [optional] 
**CreatedAt** | Pointer to **NullableTime** | The date the feature was created | [optional] 
**ArchivedAt** | Pointer to **NullableTime** | The date the feature was archived | [optional] 
**LastSeenAt** | Pointer to **NullableTime** | The date when metrics where last collected for the feature. This field is deprecated, use the one in featureEnvironmentSchema | [optional] 
**Environments** | Pointer to [**[]FeatureEnvironmentSchema**](FeatureEnvironmentSchema.md) | The list of environments where the feature can be used | [optional] 
**Variants** | Pointer to [**[]VariantSchema**](VariantSchema.md) | The list of feature variants | [optional] 
**Strategies** | Pointer to **[]map[string]interface{}** | This is a legacy field that will be deprecated | [optional] 
**Tags** | Pointer to [**[]TagSchema**](TagSchema.md) | The list of feature tags | [optional] 

## Methods

### NewFeatureSchema

`func NewFeatureSchema(name string, ) *FeatureSchema`

NewFeatureSchema instantiates a new FeatureSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeatureSchemaWithDefaults

`func NewFeatureSchemaWithDefaults() *FeatureSchema`

NewFeatureSchemaWithDefaults instantiates a new FeatureSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *FeatureSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FeatureSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FeatureSchema) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *FeatureSchema) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FeatureSchema) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FeatureSchema) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *FeatureSchema) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDescription

`func (o *FeatureSchema) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FeatureSchema) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FeatureSchema) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FeatureSchema) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *FeatureSchema) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *FeatureSchema) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetArchived

`func (o *FeatureSchema) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *FeatureSchema) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *FeatureSchema) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *FeatureSchema) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### GetProject

`func (o *FeatureSchema) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *FeatureSchema) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *FeatureSchema) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *FeatureSchema) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetEnabled

`func (o *FeatureSchema) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *FeatureSchema) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *FeatureSchema) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *FeatureSchema) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetStale

`func (o *FeatureSchema) GetStale() bool`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *FeatureSchema) GetStaleOk() (*bool, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *FeatureSchema) SetStale(v bool)`

SetStale sets Stale field to given value.

### HasStale

`func (o *FeatureSchema) HasStale() bool`

HasStale returns a boolean if a field has been set.

### GetFavorite

`func (o *FeatureSchema) GetFavorite() bool`

GetFavorite returns the Favorite field if non-nil, zero value otherwise.

### GetFavoriteOk

`func (o *FeatureSchema) GetFavoriteOk() (*bool, bool)`

GetFavoriteOk returns a tuple with the Favorite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFavorite

`func (o *FeatureSchema) SetFavorite(v bool)`

SetFavorite sets Favorite field to given value.

### HasFavorite

`func (o *FeatureSchema) HasFavorite() bool`

HasFavorite returns a boolean if a field has been set.

### GetImpressionData

`func (o *FeatureSchema) GetImpressionData() bool`

GetImpressionData returns the ImpressionData field if non-nil, zero value otherwise.

### GetImpressionDataOk

`func (o *FeatureSchema) GetImpressionDataOk() (*bool, bool)`

GetImpressionDataOk returns a tuple with the ImpressionData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpressionData

`func (o *FeatureSchema) SetImpressionData(v bool)`

SetImpressionData sets ImpressionData field to given value.

### HasImpressionData

`func (o *FeatureSchema) HasImpressionData() bool`

HasImpressionData returns a boolean if a field has been set.

### GetCreatedAt

`func (o *FeatureSchema) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FeatureSchema) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FeatureSchema) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *FeatureSchema) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *FeatureSchema) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *FeatureSchema) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetArchivedAt

`func (o *FeatureSchema) GetArchivedAt() time.Time`

GetArchivedAt returns the ArchivedAt field if non-nil, zero value otherwise.

### GetArchivedAtOk

`func (o *FeatureSchema) GetArchivedAtOk() (*time.Time, bool)`

GetArchivedAtOk returns a tuple with the ArchivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedAt

`func (o *FeatureSchema) SetArchivedAt(v time.Time)`

SetArchivedAt sets ArchivedAt field to given value.

### HasArchivedAt

`func (o *FeatureSchema) HasArchivedAt() bool`

HasArchivedAt returns a boolean if a field has been set.

### SetArchivedAtNil

`func (o *FeatureSchema) SetArchivedAtNil(b bool)`

 SetArchivedAtNil sets the value for ArchivedAt to be an explicit nil

### UnsetArchivedAt
`func (o *FeatureSchema) UnsetArchivedAt()`

UnsetArchivedAt ensures that no value is present for ArchivedAt, not even an explicit nil
### GetLastSeenAt

`func (o *FeatureSchema) GetLastSeenAt() time.Time`

GetLastSeenAt returns the LastSeenAt field if non-nil, zero value otherwise.

### GetLastSeenAtOk

`func (o *FeatureSchema) GetLastSeenAtOk() (*time.Time, bool)`

GetLastSeenAtOk returns a tuple with the LastSeenAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeenAt

`func (o *FeatureSchema) SetLastSeenAt(v time.Time)`

SetLastSeenAt sets LastSeenAt field to given value.

### HasLastSeenAt

`func (o *FeatureSchema) HasLastSeenAt() bool`

HasLastSeenAt returns a boolean if a field has been set.

### SetLastSeenAtNil

`func (o *FeatureSchema) SetLastSeenAtNil(b bool)`

 SetLastSeenAtNil sets the value for LastSeenAt to be an explicit nil

### UnsetLastSeenAt
`func (o *FeatureSchema) UnsetLastSeenAt()`

UnsetLastSeenAt ensures that no value is present for LastSeenAt, not even an explicit nil
### GetEnvironments

`func (o *FeatureSchema) GetEnvironments() []FeatureEnvironmentSchema`

GetEnvironments returns the Environments field if non-nil, zero value otherwise.

### GetEnvironmentsOk

`func (o *FeatureSchema) GetEnvironmentsOk() (*[]FeatureEnvironmentSchema, bool)`

GetEnvironmentsOk returns a tuple with the Environments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironments

`func (o *FeatureSchema) SetEnvironments(v []FeatureEnvironmentSchema)`

SetEnvironments sets Environments field to given value.

### HasEnvironments

`func (o *FeatureSchema) HasEnvironments() bool`

HasEnvironments returns a boolean if a field has been set.

### GetVariants

`func (o *FeatureSchema) GetVariants() []VariantSchema`

GetVariants returns the Variants field if non-nil, zero value otherwise.

### GetVariantsOk

`func (o *FeatureSchema) GetVariantsOk() (*[]VariantSchema, bool)`

GetVariantsOk returns a tuple with the Variants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariants

`func (o *FeatureSchema) SetVariants(v []VariantSchema)`

SetVariants sets Variants field to given value.

### HasVariants

`func (o *FeatureSchema) HasVariants() bool`

HasVariants returns a boolean if a field has been set.

### GetStrategies

`func (o *FeatureSchema) GetStrategies() []map[string]interface{}`

GetStrategies returns the Strategies field if non-nil, zero value otherwise.

### GetStrategiesOk

`func (o *FeatureSchema) GetStrategiesOk() (*[]map[string]interface{}, bool)`

GetStrategiesOk returns a tuple with the Strategies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategies

`func (o *FeatureSchema) SetStrategies(v []map[string]interface{})`

SetStrategies sets Strategies field to given value.

### HasStrategies

`func (o *FeatureSchema) HasStrategies() bool`

HasStrategies returns a boolean if a field has been set.

### GetTags

`func (o *FeatureSchema) GetTags() []TagSchema`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *FeatureSchema) GetTagsOk() (*[]TagSchema, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *FeatureSchema) SetTags(v []TagSchema)`

SetTags sets Tags field to given value.

### HasTags

`func (o *FeatureSchema) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *FeatureSchema) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *FeatureSchema) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


