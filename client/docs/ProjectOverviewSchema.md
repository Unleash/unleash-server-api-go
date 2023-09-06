# ProjectOverviewSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Stats** | Pointer to [**ProjectStatsSchema**](ProjectStatsSchema.md) |  | [optional] 
**Version** | **int32** | The schema version used to describe the project overview | 
**Name** | **string** | The name of this project | 
**Description** | Pointer to **NullableString** | Additional information about the project | [optional] 
**DefaultStickiness** | Pointer to **string** | A default stickiness for the project affecting the default stickiness value for variants and Gradual Rollout strategy | [optional] 
**Mode** | Pointer to **string** | The project&#39;s [collaboration mode](https://docs.getunleash.io/reference/project-collaboration-mode). Determines whether non-project members can submit change requests or not. | [optional] 
**FeatureLimit** | Pointer to **NullableFloat32** | A limit on the number of features allowed in the project. Null if no limit. | [optional] 
**FeatureNaming** | Pointer to [**CreateFeatureNamingPatternSchema**](CreateFeatureNamingPatternSchema.md) |  | [optional] 
**Members** | Pointer to **float32** | The number of members this project has | [optional] 
**Health** | Pointer to **float32** | An indicator of the [project&#39;s health](https://docs.getunleash.io/reference/technical-debt#health-rating) on a scale from 0 to 100 | [optional] 
**Environments** | Pointer to [**[]ProjectEnvironmentSchema**](ProjectEnvironmentSchema.md) | The environments that are enabled for this project | [optional] 
**Features** | Pointer to [**[]FeatureSchema**](FeatureSchema.md) | The full list of features in this project (excluding archived features) | [optional] 
**UpdatedAt** | Pointer to **NullableTime** | When the project was last updated. | [optional] 
**CreatedAt** | Pointer to **NullableTime** | When the project was created. | [optional] 
**Favorite** | Pointer to **bool** | &#x60;true&#x60; if the project was favorited, otherwise &#x60;false&#x60;. | [optional] 

## Methods

### NewProjectOverviewSchema

`func NewProjectOverviewSchema(version int32, name string, ) *ProjectOverviewSchema`

NewProjectOverviewSchema instantiates a new ProjectOverviewSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectOverviewSchemaWithDefaults

`func NewProjectOverviewSchemaWithDefaults() *ProjectOverviewSchema`

NewProjectOverviewSchemaWithDefaults instantiates a new ProjectOverviewSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStats

`func (o *ProjectOverviewSchema) GetStats() ProjectStatsSchema`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *ProjectOverviewSchema) GetStatsOk() (*ProjectStatsSchema, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *ProjectOverviewSchema) SetStats(v ProjectStatsSchema)`

SetStats sets Stats field to given value.

### HasStats

`func (o *ProjectOverviewSchema) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetVersion

`func (o *ProjectOverviewSchema) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ProjectOverviewSchema) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ProjectOverviewSchema) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetName

`func (o *ProjectOverviewSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectOverviewSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectOverviewSchema) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ProjectOverviewSchema) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProjectOverviewSchema) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProjectOverviewSchema) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProjectOverviewSchema) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ProjectOverviewSchema) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ProjectOverviewSchema) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDefaultStickiness

`func (o *ProjectOverviewSchema) GetDefaultStickiness() string`

GetDefaultStickiness returns the DefaultStickiness field if non-nil, zero value otherwise.

### GetDefaultStickinessOk

`func (o *ProjectOverviewSchema) GetDefaultStickinessOk() (*string, bool)`

GetDefaultStickinessOk returns a tuple with the DefaultStickiness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultStickiness

`func (o *ProjectOverviewSchema) SetDefaultStickiness(v string)`

SetDefaultStickiness sets DefaultStickiness field to given value.

### HasDefaultStickiness

`func (o *ProjectOverviewSchema) HasDefaultStickiness() bool`

HasDefaultStickiness returns a boolean if a field has been set.

### GetMode

`func (o *ProjectOverviewSchema) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *ProjectOverviewSchema) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *ProjectOverviewSchema) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *ProjectOverviewSchema) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetFeatureLimit

`func (o *ProjectOverviewSchema) GetFeatureLimit() float32`

GetFeatureLimit returns the FeatureLimit field if non-nil, zero value otherwise.

### GetFeatureLimitOk

`func (o *ProjectOverviewSchema) GetFeatureLimitOk() (*float32, bool)`

GetFeatureLimitOk returns a tuple with the FeatureLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureLimit

`func (o *ProjectOverviewSchema) SetFeatureLimit(v float32)`

SetFeatureLimit sets FeatureLimit field to given value.

### HasFeatureLimit

`func (o *ProjectOverviewSchema) HasFeatureLimit() bool`

HasFeatureLimit returns a boolean if a field has been set.

### SetFeatureLimitNil

`func (o *ProjectOverviewSchema) SetFeatureLimitNil(b bool)`

 SetFeatureLimitNil sets the value for FeatureLimit to be an explicit nil

### UnsetFeatureLimit
`func (o *ProjectOverviewSchema) UnsetFeatureLimit()`

UnsetFeatureLimit ensures that no value is present for FeatureLimit, not even an explicit nil
### GetFeatureNaming

`func (o *ProjectOverviewSchema) GetFeatureNaming() CreateFeatureNamingPatternSchema`

GetFeatureNaming returns the FeatureNaming field if non-nil, zero value otherwise.

### GetFeatureNamingOk

`func (o *ProjectOverviewSchema) GetFeatureNamingOk() (*CreateFeatureNamingPatternSchema, bool)`

GetFeatureNamingOk returns a tuple with the FeatureNaming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureNaming

`func (o *ProjectOverviewSchema) SetFeatureNaming(v CreateFeatureNamingPatternSchema)`

SetFeatureNaming sets FeatureNaming field to given value.

### HasFeatureNaming

`func (o *ProjectOverviewSchema) HasFeatureNaming() bool`

HasFeatureNaming returns a boolean if a field has been set.

### GetMembers

`func (o *ProjectOverviewSchema) GetMembers() float32`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *ProjectOverviewSchema) GetMembersOk() (*float32, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *ProjectOverviewSchema) SetMembers(v float32)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *ProjectOverviewSchema) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetHealth

`func (o *ProjectOverviewSchema) GetHealth() float32`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *ProjectOverviewSchema) GetHealthOk() (*float32, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *ProjectOverviewSchema) SetHealth(v float32)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *ProjectOverviewSchema) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetEnvironments

`func (o *ProjectOverviewSchema) GetEnvironments() []ProjectEnvironmentSchema`

GetEnvironments returns the Environments field if non-nil, zero value otherwise.

### GetEnvironmentsOk

`func (o *ProjectOverviewSchema) GetEnvironmentsOk() (*[]ProjectEnvironmentSchema, bool)`

GetEnvironmentsOk returns a tuple with the Environments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironments

`func (o *ProjectOverviewSchema) SetEnvironments(v []ProjectEnvironmentSchema)`

SetEnvironments sets Environments field to given value.

### HasEnvironments

`func (o *ProjectOverviewSchema) HasEnvironments() bool`

HasEnvironments returns a boolean if a field has been set.

### GetFeatures

`func (o *ProjectOverviewSchema) GetFeatures() []FeatureSchema`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *ProjectOverviewSchema) GetFeaturesOk() (*[]FeatureSchema, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *ProjectOverviewSchema) SetFeatures(v []FeatureSchema)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *ProjectOverviewSchema) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ProjectOverviewSchema) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ProjectOverviewSchema) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ProjectOverviewSchema) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ProjectOverviewSchema) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *ProjectOverviewSchema) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *ProjectOverviewSchema) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetCreatedAt

`func (o *ProjectOverviewSchema) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ProjectOverviewSchema) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ProjectOverviewSchema) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ProjectOverviewSchema) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *ProjectOverviewSchema) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *ProjectOverviewSchema) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetFavorite

`func (o *ProjectOverviewSchema) GetFavorite() bool`

GetFavorite returns the Favorite field if non-nil, zero value otherwise.

### GetFavoriteOk

`func (o *ProjectOverviewSchema) GetFavoriteOk() (*bool, bool)`

GetFavoriteOk returns a tuple with the Favorite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFavorite

`func (o *ProjectOverviewSchema) SetFavorite(v bool)`

SetFavorite sets Favorite field to given value.

### HasFavorite

`func (o *ProjectOverviewSchema) HasFavorite() bool`

HasFavorite returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


