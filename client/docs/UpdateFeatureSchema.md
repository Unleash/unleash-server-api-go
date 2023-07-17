# UpdateFeatureSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Detailed description of the feature | [optional] 
**Type** | Pointer to **string** | Type of the toggle e.g. experiment, kill-switch, release, operational, permission | [optional] 
**Stale** | Pointer to **bool** | &#x60;true&#x60; if the feature is archived | [optional] 
**Archived** | Pointer to **bool** | If &#x60;true&#x60; the feature toggle will be moved to the [archive](https://docs.getunleash.io/reference/archived-toggles) with a property &#x60;archivedAt&#x60; set to current time | [optional] 
**ImpressionData** | Pointer to **bool** | &#x60;true&#x60; if the impression data collection is enabled for the feature | [optional] 

## Methods

### NewUpdateFeatureSchema

`func NewUpdateFeatureSchema() *UpdateFeatureSchema`

NewUpdateFeatureSchema instantiates a new UpdateFeatureSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateFeatureSchemaWithDefaults

`func NewUpdateFeatureSchemaWithDefaults() *UpdateFeatureSchema`

NewUpdateFeatureSchemaWithDefaults instantiates a new UpdateFeatureSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *UpdateFeatureSchema) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateFeatureSchema) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateFeatureSchema) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateFeatureSchema) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *UpdateFeatureSchema) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateFeatureSchema) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateFeatureSchema) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UpdateFeatureSchema) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStale

`func (o *UpdateFeatureSchema) GetStale() bool`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *UpdateFeatureSchema) GetStaleOk() (*bool, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *UpdateFeatureSchema) SetStale(v bool)`

SetStale sets Stale field to given value.

### HasStale

`func (o *UpdateFeatureSchema) HasStale() bool`

HasStale returns a boolean if a field has been set.

### GetArchived

`func (o *UpdateFeatureSchema) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *UpdateFeatureSchema) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *UpdateFeatureSchema) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *UpdateFeatureSchema) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### GetImpressionData

`func (o *UpdateFeatureSchema) GetImpressionData() bool`

GetImpressionData returns the ImpressionData field if non-nil, zero value otherwise.

### GetImpressionDataOk

`func (o *UpdateFeatureSchema) GetImpressionDataOk() (*bool, bool)`

GetImpressionDataOk returns a tuple with the ImpressionData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpressionData

`func (o *UpdateFeatureSchema) SetImpressionData(v bool)`

SetImpressionData sets ImpressionData field to given value.

### HasImpressionData

`func (o *UpdateFeatureSchema) HasImpressionData() bool`

HasImpressionData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


