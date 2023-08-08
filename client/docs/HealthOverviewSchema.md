# HealthOverviewSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | **int32** | The project overview version. | 
**Name** | **string** | The project&#39;s name | 
**Description** | Pointer to **NullableString** | The project&#39;s description | [optional] 
**DefaultStickiness** | **string** | A default stickiness for the project affecting the default stickiness value for variants and Gradual Rollout strategy | 
**Mode** | **string** | The project&#39;s [collaboration mode](https://docs.getunleash.io/reference/project-collaboration-mode). Determines whether non-project members can submit change requests or not. | 
**FeatureLimit** | Pointer to **NullableFloat32** | A limit on the number of features allowed in the project. Null if no limit. | [optional] 
**Members** | **int32** | The number of users/members in the project. | 
**Health** | **int32** | The overall [health rating](https://docs.getunleash.io/reference/technical-debt#health-rating) of the project. | 
**Environments** | [**[]ProjectEnvironmentSchema**](ProjectEnvironmentSchema.md) | An array containing the names of all the environments configured for the project. | 
**Features** | [**[]FeatureSchema**](FeatureSchema.md) | An array containing an overview of all the features of the project and their individual status | 
**UpdatedAt** | Pointer to **NullableTime** | When the project was last updated. | [optional] 
**CreatedAt** | Pointer to **NullableTime** | When the project was last updated. | [optional] 
**Favorite** | Pointer to **bool** | Indicates if the project has been marked as a favorite by the current user requesting the project health overview. | [optional] 
**Stats** | Pointer to [**ProjectStatsSchema**](ProjectStatsSchema.md) |  | [optional] 

## Methods

### NewHealthOverviewSchema

`func NewHealthOverviewSchema(version int32, name string, defaultStickiness string, mode string, members int32, health int32, environments []ProjectEnvironmentSchema, features []FeatureSchema, ) *HealthOverviewSchema`

NewHealthOverviewSchema instantiates a new HealthOverviewSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHealthOverviewSchemaWithDefaults

`func NewHealthOverviewSchemaWithDefaults() *HealthOverviewSchema`

NewHealthOverviewSchemaWithDefaults instantiates a new HealthOverviewSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *HealthOverviewSchema) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *HealthOverviewSchema) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *HealthOverviewSchema) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetName

`func (o *HealthOverviewSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HealthOverviewSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HealthOverviewSchema) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *HealthOverviewSchema) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *HealthOverviewSchema) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *HealthOverviewSchema) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *HealthOverviewSchema) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *HealthOverviewSchema) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *HealthOverviewSchema) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDefaultStickiness

`func (o *HealthOverviewSchema) GetDefaultStickiness() string`

GetDefaultStickiness returns the DefaultStickiness field if non-nil, zero value otherwise.

### GetDefaultStickinessOk

`func (o *HealthOverviewSchema) GetDefaultStickinessOk() (*string, bool)`

GetDefaultStickinessOk returns a tuple with the DefaultStickiness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultStickiness

`func (o *HealthOverviewSchema) SetDefaultStickiness(v string)`

SetDefaultStickiness sets DefaultStickiness field to given value.


### GetMode

`func (o *HealthOverviewSchema) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *HealthOverviewSchema) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *HealthOverviewSchema) SetMode(v string)`

SetMode sets Mode field to given value.


### GetFeatureLimit

`func (o *HealthOverviewSchema) GetFeatureLimit() float32`

GetFeatureLimit returns the FeatureLimit field if non-nil, zero value otherwise.

### GetFeatureLimitOk

`func (o *HealthOverviewSchema) GetFeatureLimitOk() (*float32, bool)`

GetFeatureLimitOk returns a tuple with the FeatureLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureLimit

`func (o *HealthOverviewSchema) SetFeatureLimit(v float32)`

SetFeatureLimit sets FeatureLimit field to given value.

### HasFeatureLimit

`func (o *HealthOverviewSchema) HasFeatureLimit() bool`

HasFeatureLimit returns a boolean if a field has been set.

### SetFeatureLimitNil

`func (o *HealthOverviewSchema) SetFeatureLimitNil(b bool)`

 SetFeatureLimitNil sets the value for FeatureLimit to be an explicit nil

### UnsetFeatureLimit
`func (o *HealthOverviewSchema) UnsetFeatureLimit()`

UnsetFeatureLimit ensures that no value is present for FeatureLimit, not even an explicit nil
### GetMembers

`func (o *HealthOverviewSchema) GetMembers() int32`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *HealthOverviewSchema) GetMembersOk() (*int32, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *HealthOverviewSchema) SetMembers(v int32)`

SetMembers sets Members field to given value.


### GetHealth

`func (o *HealthOverviewSchema) GetHealth() int32`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *HealthOverviewSchema) GetHealthOk() (*int32, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *HealthOverviewSchema) SetHealth(v int32)`

SetHealth sets Health field to given value.


### GetEnvironments

`func (o *HealthOverviewSchema) GetEnvironments() []ProjectEnvironmentSchema`

GetEnvironments returns the Environments field if non-nil, zero value otherwise.

### GetEnvironmentsOk

`func (o *HealthOverviewSchema) GetEnvironmentsOk() (*[]ProjectEnvironmentSchema, bool)`

GetEnvironmentsOk returns a tuple with the Environments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironments

`func (o *HealthOverviewSchema) SetEnvironments(v []ProjectEnvironmentSchema)`

SetEnvironments sets Environments field to given value.


### GetFeatures

`func (o *HealthOverviewSchema) GetFeatures() []FeatureSchema`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *HealthOverviewSchema) GetFeaturesOk() (*[]FeatureSchema, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *HealthOverviewSchema) SetFeatures(v []FeatureSchema)`

SetFeatures sets Features field to given value.


### GetUpdatedAt

`func (o *HealthOverviewSchema) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *HealthOverviewSchema) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *HealthOverviewSchema) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *HealthOverviewSchema) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *HealthOverviewSchema) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *HealthOverviewSchema) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetCreatedAt

`func (o *HealthOverviewSchema) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *HealthOverviewSchema) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *HealthOverviewSchema) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *HealthOverviewSchema) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *HealthOverviewSchema) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *HealthOverviewSchema) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetFavorite

`func (o *HealthOverviewSchema) GetFavorite() bool`

GetFavorite returns the Favorite field if non-nil, zero value otherwise.

### GetFavoriteOk

`func (o *HealthOverviewSchema) GetFavoriteOk() (*bool, bool)`

GetFavoriteOk returns a tuple with the Favorite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFavorite

`func (o *HealthOverviewSchema) SetFavorite(v bool)`

SetFavorite sets Favorite field to given value.

### HasFavorite

`func (o *HealthOverviewSchema) HasFavorite() bool`

HasFavorite returns a boolean if a field has been set.

### GetStats

`func (o *HealthOverviewSchema) GetStats() ProjectStatsSchema`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *HealthOverviewSchema) GetStatsOk() (*ProjectStatsSchema, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *HealthOverviewSchema) SetStats(v ProjectStatsSchema)`

SetStats sets Stats field to given value.

### HasStats

`func (o *HealthOverviewSchema) HasStats() bool`

HasStats returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


