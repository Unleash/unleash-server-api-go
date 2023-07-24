# EventSchemaData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the feature toggle/strategy/environment that this event relates to | [optional] 
**Description** | Pointer to **string** | The description of the object this event relates to | [optional] 
**Type** | Pointer to **string** | If this event relates to a feature toggle, the type of feature toggle. | [optional] 
**Project** | Pointer to **string** | The project this event relates to | [optional] 
**Stale** | Pointer to **bool** | Is the feature toggle this event relates to stale | [optional] 
**Variants** | Pointer to [**[]VariantSchema**](VariantSchema.md) | Variants configured for this toggle | [optional] 
**CreatedAt** | Pointer to **time.Time** | The time the event happened as a RFC 3339-conformant timestamp. | [optional] 
**LastSeenAt** | Pointer to **NullableTime** | The time the feature was last seen | [optional] 
**ImpressionData** | Pointer to **bool** | Should [impression events](https://docs.getunleash.io/reference/impression-data) activate for this feature toggle | [optional] 

## Methods

### NewEventSchemaData

`func NewEventSchemaData() *EventSchemaData`

NewEventSchemaData instantiates a new EventSchemaData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventSchemaDataWithDefaults

`func NewEventSchemaDataWithDefaults() *EventSchemaData`

NewEventSchemaDataWithDefaults instantiates a new EventSchemaData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EventSchemaData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EventSchemaData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EventSchemaData) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EventSchemaData) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *EventSchemaData) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EventSchemaData) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EventSchemaData) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EventSchemaData) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *EventSchemaData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EventSchemaData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EventSchemaData) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EventSchemaData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetProject

`func (o *EventSchemaData) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *EventSchemaData) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *EventSchemaData) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *EventSchemaData) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetStale

`func (o *EventSchemaData) GetStale() bool`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *EventSchemaData) GetStaleOk() (*bool, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *EventSchemaData) SetStale(v bool)`

SetStale sets Stale field to given value.

### HasStale

`func (o *EventSchemaData) HasStale() bool`

HasStale returns a boolean if a field has been set.

### GetVariants

`func (o *EventSchemaData) GetVariants() []VariantSchema`

GetVariants returns the Variants field if non-nil, zero value otherwise.

### GetVariantsOk

`func (o *EventSchemaData) GetVariantsOk() (*[]VariantSchema, bool)`

GetVariantsOk returns a tuple with the Variants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariants

`func (o *EventSchemaData) SetVariants(v []VariantSchema)`

SetVariants sets Variants field to given value.

### HasVariants

`func (o *EventSchemaData) HasVariants() bool`

HasVariants returns a boolean if a field has been set.

### GetCreatedAt

`func (o *EventSchemaData) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EventSchemaData) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EventSchemaData) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *EventSchemaData) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLastSeenAt

`func (o *EventSchemaData) GetLastSeenAt() time.Time`

GetLastSeenAt returns the LastSeenAt field if non-nil, zero value otherwise.

### GetLastSeenAtOk

`func (o *EventSchemaData) GetLastSeenAtOk() (*time.Time, bool)`

GetLastSeenAtOk returns a tuple with the LastSeenAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeenAt

`func (o *EventSchemaData) SetLastSeenAt(v time.Time)`

SetLastSeenAt sets LastSeenAt field to given value.

### HasLastSeenAt

`func (o *EventSchemaData) HasLastSeenAt() bool`

HasLastSeenAt returns a boolean if a field has been set.

### SetLastSeenAtNil

`func (o *EventSchemaData) SetLastSeenAtNil(b bool)`

 SetLastSeenAtNil sets the value for LastSeenAt to be an explicit nil

### UnsetLastSeenAt
`func (o *EventSchemaData) UnsetLastSeenAt()`

UnsetLastSeenAt ensures that no value is present for LastSeenAt, not even an explicit nil
### GetImpressionData

`func (o *EventSchemaData) GetImpressionData() bool`

GetImpressionData returns the ImpressionData field if non-nil, zero value otherwise.

### GetImpressionDataOk

`func (o *EventSchemaData) GetImpressionDataOk() (*bool, bool)`

GetImpressionDataOk returns a tuple with the ImpressionData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpressionData

`func (o *EventSchemaData) SetImpressionData(v bool)`

SetImpressionData sets ImpressionData field to given value.

### HasImpressionData

`func (o *EventSchemaData) HasImpressionData() bool`

HasImpressionData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


