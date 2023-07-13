# HealthReportSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | **int32** | The project overview version. | 
**Name** | **string** | The project&#39;s name | 
**Description** | Pointer to **NullableString** | The project&#39;s description | [optional] 
**DefaultStickiness** | **string** | A default stickiness for the project affecting the default stickiness value for variants and Gradual Rollout strategy | 
**Mode** | **string** | The project&#39;s [collaboration mode](https://docs.getunleash.io/reference/project-collaboration-mode). Determines whether non-project members can submit change requests or not. | 
**Members** | **int32** | The number of users/members in the project. | 
**Health** | **int32** | The overall [health rating](https://docs.getunleash.io/reference/technical-debt#health-rating) of the project. | 
**Environments** | [**[]ProjectEnvironmentSchema**](ProjectEnvironmentSchema.md) | An array containing the names of all the environments configured for the project. | 
**Features** | [**[]FeatureSchema**](FeatureSchema.md) | An array containing an overview of all the features of the project and their individual status | 
**UpdatedAt** | Pointer to **NullableTime** | When the project was last updated. | [optional] 
**CreatedAt** | Pointer to **NullableTime** | When the project was last updated. | [optional] 
**Favorite** | Pointer to **bool** | Indicates if the project has been marked as a favorite by the current user requesting the project health overview. | [optional] 
**Stats** | Pointer to [**ProjectStatsSchema**](ProjectStatsSchema.md) |  | [optional] 
**PotentiallyStaleCount** | **float32** | The number of potentially stale feature toggles. | 
**ActiveCount** | **float32** | The number of active feature toggles. | 
**StaleCount** | **float32** | The number of stale feature toggles. | 

## Methods

### NewHealthReportSchema

`func NewHealthReportSchema(version int32, name string, defaultStickiness string, mode string, members int32, health int32, environments []ProjectEnvironmentSchema, features []FeatureSchema, potentiallyStaleCount float32, activeCount float32, staleCount float32, ) *HealthReportSchema`

NewHealthReportSchema instantiates a new HealthReportSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHealthReportSchemaWithDefaults

`func NewHealthReportSchemaWithDefaults() *HealthReportSchema`

NewHealthReportSchemaWithDefaults instantiates a new HealthReportSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *HealthReportSchema) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *HealthReportSchema) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *HealthReportSchema) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetName

`func (o *HealthReportSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HealthReportSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HealthReportSchema) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *HealthReportSchema) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *HealthReportSchema) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *HealthReportSchema) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *HealthReportSchema) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *HealthReportSchema) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *HealthReportSchema) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDefaultStickiness

`func (o *HealthReportSchema) GetDefaultStickiness() string`

GetDefaultStickiness returns the DefaultStickiness field if non-nil, zero value otherwise.

### GetDefaultStickinessOk

`func (o *HealthReportSchema) GetDefaultStickinessOk() (*string, bool)`

GetDefaultStickinessOk returns a tuple with the DefaultStickiness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultStickiness

`func (o *HealthReportSchema) SetDefaultStickiness(v string)`

SetDefaultStickiness sets DefaultStickiness field to given value.


### GetMode

`func (o *HealthReportSchema) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *HealthReportSchema) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *HealthReportSchema) SetMode(v string)`

SetMode sets Mode field to given value.


### GetMembers

`func (o *HealthReportSchema) GetMembers() int32`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *HealthReportSchema) GetMembersOk() (*int32, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *HealthReportSchema) SetMembers(v int32)`

SetMembers sets Members field to given value.


### GetHealth

`func (o *HealthReportSchema) GetHealth() int32`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *HealthReportSchema) GetHealthOk() (*int32, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *HealthReportSchema) SetHealth(v int32)`

SetHealth sets Health field to given value.


### GetEnvironments

`func (o *HealthReportSchema) GetEnvironments() []ProjectEnvironmentSchema`

GetEnvironments returns the Environments field if non-nil, zero value otherwise.

### GetEnvironmentsOk

`func (o *HealthReportSchema) GetEnvironmentsOk() (*[]ProjectEnvironmentSchema, bool)`

GetEnvironmentsOk returns a tuple with the Environments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironments

`func (o *HealthReportSchema) SetEnvironments(v []ProjectEnvironmentSchema)`

SetEnvironments sets Environments field to given value.


### GetFeatures

`func (o *HealthReportSchema) GetFeatures() []FeatureSchema`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *HealthReportSchema) GetFeaturesOk() (*[]FeatureSchema, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *HealthReportSchema) SetFeatures(v []FeatureSchema)`

SetFeatures sets Features field to given value.


### GetUpdatedAt

`func (o *HealthReportSchema) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *HealthReportSchema) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *HealthReportSchema) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *HealthReportSchema) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *HealthReportSchema) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *HealthReportSchema) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetCreatedAt

`func (o *HealthReportSchema) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *HealthReportSchema) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *HealthReportSchema) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *HealthReportSchema) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *HealthReportSchema) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *HealthReportSchema) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetFavorite

`func (o *HealthReportSchema) GetFavorite() bool`

GetFavorite returns the Favorite field if non-nil, zero value otherwise.

### GetFavoriteOk

`func (o *HealthReportSchema) GetFavoriteOk() (*bool, bool)`

GetFavoriteOk returns a tuple with the Favorite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFavorite

`func (o *HealthReportSchema) SetFavorite(v bool)`

SetFavorite sets Favorite field to given value.

### HasFavorite

`func (o *HealthReportSchema) HasFavorite() bool`

HasFavorite returns a boolean if a field has been set.

### GetStats

`func (o *HealthReportSchema) GetStats() ProjectStatsSchema`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *HealthReportSchema) GetStatsOk() (*ProjectStatsSchema, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *HealthReportSchema) SetStats(v ProjectStatsSchema)`

SetStats sets Stats field to given value.

### HasStats

`func (o *HealthReportSchema) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetPotentiallyStaleCount

`func (o *HealthReportSchema) GetPotentiallyStaleCount() float32`

GetPotentiallyStaleCount returns the PotentiallyStaleCount field if non-nil, zero value otherwise.

### GetPotentiallyStaleCountOk

`func (o *HealthReportSchema) GetPotentiallyStaleCountOk() (*float32, bool)`

GetPotentiallyStaleCountOk returns a tuple with the PotentiallyStaleCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPotentiallyStaleCount

`func (o *HealthReportSchema) SetPotentiallyStaleCount(v float32)`

SetPotentiallyStaleCount sets PotentiallyStaleCount field to given value.


### GetActiveCount

`func (o *HealthReportSchema) GetActiveCount() float32`

GetActiveCount returns the ActiveCount field if non-nil, zero value otherwise.

### GetActiveCountOk

`func (o *HealthReportSchema) GetActiveCountOk() (*float32, bool)`

GetActiveCountOk returns a tuple with the ActiveCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveCount

`func (o *HealthReportSchema) SetActiveCount(v float32)`

SetActiveCount sets ActiveCount field to given value.


### GetStaleCount

`func (o *HealthReportSchema) GetStaleCount() float32`

GetStaleCount returns the StaleCount field if non-nil, zero value otherwise.

### GetStaleCountOk

`func (o *HealthReportSchema) GetStaleCountOk() (*float32, bool)`

GetStaleCountOk returns a tuple with the StaleCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaleCount

`func (o *HealthReportSchema) SetStaleCount(v float32)`

SetStaleCount sets StaleCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


