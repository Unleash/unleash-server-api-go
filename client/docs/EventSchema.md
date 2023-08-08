# EventSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The ID of the event. An increasing natural number. | 
**CreatedAt** | **time.Time** | The time the event happened as a RFC 3339-conformant timestamp. | 
**Type** | **string** | What [type](https://docs.getunleash.io/reference/api/legacy/unleash/admin/events#event-type-description) of event this is | 
**CreatedBy** | **string** | Which user created this event | 
**Environment** | Pointer to **NullableString** | The feature toggle environment the event relates to, if applicable. | [optional] 
**Project** | Pointer to **NullableString** | The project the event relates to, if applicable. | [optional] 
**FeatureName** | Pointer to **NullableString** | The name of the feature toggle the event relates to, if applicable. | [optional] 
**Data** | Pointer to **map[string]interface{}** | Extra associated data related to the event, such as feature toggle state, segment configuration, etc., if applicable. | [optional] 
**PreData** | Pointer to **map[string]interface{}** | Data relating to the previous state of the event&#39;s subject. | [optional] 
**Tags** | Pointer to [**[]TagSchema**](TagSchema.md) | Any tags related to the event, if applicable. | [optional] 

## Methods

### NewEventSchema

`func NewEventSchema(id int32, createdAt time.Time, type_ string, createdBy string, ) *EventSchema`

NewEventSchema instantiates a new EventSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventSchemaWithDefaults

`func NewEventSchemaWithDefaults() *EventSchema`

NewEventSchemaWithDefaults instantiates a new EventSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EventSchema) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EventSchema) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EventSchema) SetId(v int32)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *EventSchema) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EventSchema) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EventSchema) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetType

`func (o *EventSchema) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EventSchema) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EventSchema) SetType(v string)`

SetType sets Type field to given value.


### GetCreatedBy

`func (o *EventSchema) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *EventSchema) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *EventSchema) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetEnvironment

`func (o *EventSchema) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *EventSchema) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *EventSchema) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *EventSchema) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### SetEnvironmentNil

`func (o *EventSchema) SetEnvironmentNil(b bool)`

 SetEnvironmentNil sets the value for Environment to be an explicit nil

### UnsetEnvironment
`func (o *EventSchema) UnsetEnvironment()`

UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
### GetProject

`func (o *EventSchema) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *EventSchema) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *EventSchema) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *EventSchema) HasProject() bool`

HasProject returns a boolean if a field has been set.

### SetProjectNil

`func (o *EventSchema) SetProjectNil(b bool)`

 SetProjectNil sets the value for Project to be an explicit nil

### UnsetProject
`func (o *EventSchema) UnsetProject()`

UnsetProject ensures that no value is present for Project, not even an explicit nil
### GetFeatureName

`func (o *EventSchema) GetFeatureName() string`

GetFeatureName returns the FeatureName field if non-nil, zero value otherwise.

### GetFeatureNameOk

`func (o *EventSchema) GetFeatureNameOk() (*string, bool)`

GetFeatureNameOk returns a tuple with the FeatureName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureName

`func (o *EventSchema) SetFeatureName(v string)`

SetFeatureName sets FeatureName field to given value.

### HasFeatureName

`func (o *EventSchema) HasFeatureName() bool`

HasFeatureName returns a boolean if a field has been set.

### SetFeatureNameNil

`func (o *EventSchema) SetFeatureNameNil(b bool)`

 SetFeatureNameNil sets the value for FeatureName to be an explicit nil

### UnsetFeatureName
`func (o *EventSchema) UnsetFeatureName()`

UnsetFeatureName ensures that no value is present for FeatureName, not even an explicit nil
### GetData

`func (o *EventSchema) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *EventSchema) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *EventSchema) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *EventSchema) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *EventSchema) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *EventSchema) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetPreData

`func (o *EventSchema) GetPreData() map[string]interface{}`

GetPreData returns the PreData field if non-nil, zero value otherwise.

### GetPreDataOk

`func (o *EventSchema) GetPreDataOk() (*map[string]interface{}, bool)`

GetPreDataOk returns a tuple with the PreData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreData

`func (o *EventSchema) SetPreData(v map[string]interface{})`

SetPreData sets PreData field to given value.

### HasPreData

`func (o *EventSchema) HasPreData() bool`

HasPreData returns a boolean if a field has been set.

### SetPreDataNil

`func (o *EventSchema) SetPreDataNil(b bool)`

 SetPreDataNil sets the value for PreData to be an explicit nil

### UnsetPreData
`func (o *EventSchema) UnsetPreData()`

UnsetPreData ensures that no value is present for PreData, not even an explicit nil
### GetTags

`func (o *EventSchema) GetTags() []TagSchema`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *EventSchema) GetTagsOk() (*[]TagSchema, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *EventSchema) SetTags(v []TagSchema)`

SetTags sets Tags field to given value.

### HasTags

`func (o *EventSchema) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *EventSchema) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *EventSchema) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


