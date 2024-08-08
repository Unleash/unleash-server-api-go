# ProjectSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The id of this project | 
**Name** | **string** | The name of this project | 
**Description** | Pointer to **NullableString** | Additional information about the project | [optional] 
**Health** | Pointer to **float32** | An indicator of the [project&#39;s health](https://docs.getunleash.io/reference/technical-debt#health-rating) on a scale from 0 to 100 | [optional] 
**FeatureCount** | Pointer to **float32** | The number of features this project has | [optional] 
**StaleFeatureCount** | Pointer to **float32** | The number of stale features this project has | [optional] 
**PotentiallyStaleFeatureCount** | Pointer to **float32** | The number of potentially stale features this project has | [optional] 
**MemberCount** | Pointer to **float32** | The number of members this project has | [optional] 
**CreatedAt** | Pointer to **time.Time** | When this project was created. | [optional] 
**UpdatedAt** | Pointer to **NullableTime** | When this project was last updated. | [optional] 
**Favorite** | Pointer to **bool** | &#x60;true&#x60; if the project was favorited, otherwise &#x60;false&#x60;. | [optional] 
**Mode** | Pointer to **string** | The project&#39;s [collaboration mode](https://docs.getunleash.io/reference/project-collaboration-mode). Determines whether non-project members can submit change requests or not. | [optional] 
**DefaultStickiness** | Pointer to **string** | A default stickiness for the project affecting the default stickiness value for variants and Gradual Rollout strategy | [optional] 
**AvgTimeToProduction** | Pointer to **float32** | The average time from when a feature was created to when it was enabled in the \&quot;production\&quot; environment during the current window | [optional] 
**Owners** | Pointer to [**ProjectSchemaOwners**](ProjectSchemaOwners.md) |  | [optional] 

## Methods

### NewProjectSchema

`func NewProjectSchema(id string, name string, ) *ProjectSchema`

NewProjectSchema instantiates a new ProjectSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectSchemaWithDefaults

`func NewProjectSchemaWithDefaults() *ProjectSchema`

NewProjectSchemaWithDefaults instantiates a new ProjectSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProjectSchema) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectSchema) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectSchema) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ProjectSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectSchema) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ProjectSchema) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProjectSchema) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProjectSchema) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProjectSchema) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ProjectSchema) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ProjectSchema) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetHealth

`func (o *ProjectSchema) GetHealth() float32`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *ProjectSchema) GetHealthOk() (*float32, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *ProjectSchema) SetHealth(v float32)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *ProjectSchema) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetFeatureCount

`func (o *ProjectSchema) GetFeatureCount() float32`

GetFeatureCount returns the FeatureCount field if non-nil, zero value otherwise.

### GetFeatureCountOk

`func (o *ProjectSchema) GetFeatureCountOk() (*float32, bool)`

GetFeatureCountOk returns a tuple with the FeatureCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureCount

`func (o *ProjectSchema) SetFeatureCount(v float32)`

SetFeatureCount sets FeatureCount field to given value.

### HasFeatureCount

`func (o *ProjectSchema) HasFeatureCount() bool`

HasFeatureCount returns a boolean if a field has been set.

### GetStaleFeatureCount

`func (o *ProjectSchema) GetStaleFeatureCount() float32`

GetStaleFeatureCount returns the StaleFeatureCount field if non-nil, zero value otherwise.

### GetStaleFeatureCountOk

`func (o *ProjectSchema) GetStaleFeatureCountOk() (*float32, bool)`

GetStaleFeatureCountOk returns a tuple with the StaleFeatureCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaleFeatureCount

`func (o *ProjectSchema) SetStaleFeatureCount(v float32)`

SetStaleFeatureCount sets StaleFeatureCount field to given value.

### HasStaleFeatureCount

`func (o *ProjectSchema) HasStaleFeatureCount() bool`

HasStaleFeatureCount returns a boolean if a field has been set.

### GetPotentiallyStaleFeatureCount

`func (o *ProjectSchema) GetPotentiallyStaleFeatureCount() float32`

GetPotentiallyStaleFeatureCount returns the PotentiallyStaleFeatureCount field if non-nil, zero value otherwise.

### GetPotentiallyStaleFeatureCountOk

`func (o *ProjectSchema) GetPotentiallyStaleFeatureCountOk() (*float32, bool)`

GetPotentiallyStaleFeatureCountOk returns a tuple with the PotentiallyStaleFeatureCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPotentiallyStaleFeatureCount

`func (o *ProjectSchema) SetPotentiallyStaleFeatureCount(v float32)`

SetPotentiallyStaleFeatureCount sets PotentiallyStaleFeatureCount field to given value.

### HasPotentiallyStaleFeatureCount

`func (o *ProjectSchema) HasPotentiallyStaleFeatureCount() bool`

HasPotentiallyStaleFeatureCount returns a boolean if a field has been set.

### GetMemberCount

`func (o *ProjectSchema) GetMemberCount() float32`

GetMemberCount returns the MemberCount field if non-nil, zero value otherwise.

### GetMemberCountOk

`func (o *ProjectSchema) GetMemberCountOk() (*float32, bool)`

GetMemberCountOk returns a tuple with the MemberCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberCount

`func (o *ProjectSchema) SetMemberCount(v float32)`

SetMemberCount sets MemberCount field to given value.

### HasMemberCount

`func (o *ProjectSchema) HasMemberCount() bool`

HasMemberCount returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ProjectSchema) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ProjectSchema) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ProjectSchema) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ProjectSchema) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ProjectSchema) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ProjectSchema) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ProjectSchema) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ProjectSchema) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *ProjectSchema) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *ProjectSchema) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetFavorite

`func (o *ProjectSchema) GetFavorite() bool`

GetFavorite returns the Favorite field if non-nil, zero value otherwise.

### GetFavoriteOk

`func (o *ProjectSchema) GetFavoriteOk() (*bool, bool)`

GetFavoriteOk returns a tuple with the Favorite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFavorite

`func (o *ProjectSchema) SetFavorite(v bool)`

SetFavorite sets Favorite field to given value.

### HasFavorite

`func (o *ProjectSchema) HasFavorite() bool`

HasFavorite returns a boolean if a field has been set.

### GetMode

`func (o *ProjectSchema) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *ProjectSchema) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *ProjectSchema) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *ProjectSchema) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetDefaultStickiness

`func (o *ProjectSchema) GetDefaultStickiness() string`

GetDefaultStickiness returns the DefaultStickiness field if non-nil, zero value otherwise.

### GetDefaultStickinessOk

`func (o *ProjectSchema) GetDefaultStickinessOk() (*string, bool)`

GetDefaultStickinessOk returns a tuple with the DefaultStickiness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultStickiness

`func (o *ProjectSchema) SetDefaultStickiness(v string)`

SetDefaultStickiness sets DefaultStickiness field to given value.

### HasDefaultStickiness

`func (o *ProjectSchema) HasDefaultStickiness() bool`

HasDefaultStickiness returns a boolean if a field has been set.

### GetAvgTimeToProduction

`func (o *ProjectSchema) GetAvgTimeToProduction() float32`

GetAvgTimeToProduction returns the AvgTimeToProduction field if non-nil, zero value otherwise.

### GetAvgTimeToProductionOk

`func (o *ProjectSchema) GetAvgTimeToProductionOk() (*float32, bool)`

GetAvgTimeToProductionOk returns a tuple with the AvgTimeToProduction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgTimeToProduction

`func (o *ProjectSchema) SetAvgTimeToProduction(v float32)`

SetAvgTimeToProduction sets AvgTimeToProduction field to given value.

### HasAvgTimeToProduction

`func (o *ProjectSchema) HasAvgTimeToProduction() bool`

HasAvgTimeToProduction returns a boolean if a field has been set.

### GetOwners

`func (o *ProjectSchema) GetOwners() ProjectSchemaOwners`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *ProjectSchema) GetOwnersOk() (*ProjectSchemaOwners, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *ProjectSchema) SetOwners(v ProjectSchemaOwners)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *ProjectSchema) HasOwners() bool`

HasOwners returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


