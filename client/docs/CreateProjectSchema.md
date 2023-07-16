# CreateProjectSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The project&#39;s identifier. | 
**Name** | **string** | The project&#39;s name. | 
**Description** | Pointer to **NullableString** | The project&#39;s description. | [optional] 
**FeatureLimit** | Pointer to **NullableInt32** | A limit on the number of features allowed in the project. &#x60;null&#x60; if no limit. | [optional] 
**Mode** | Pointer to **string** | A mode of the project affecting what actions are possible in this project | [optional] 
**DefaultStickiness** | Pointer to **string** | A default stickiness for the project affecting the default stickiness value for variants and Gradual Rollout strategy | [optional] 

## Methods

### NewCreateProjectSchema

`func NewCreateProjectSchema(id string, name string, ) *CreateProjectSchema`

NewCreateProjectSchema instantiates a new CreateProjectSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateProjectSchemaWithDefaults

`func NewCreateProjectSchemaWithDefaults() *CreateProjectSchema`

NewCreateProjectSchemaWithDefaults instantiates a new CreateProjectSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateProjectSchema) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateProjectSchema) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateProjectSchema) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *CreateProjectSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateProjectSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateProjectSchema) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateProjectSchema) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateProjectSchema) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateProjectSchema) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateProjectSchema) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreateProjectSchema) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreateProjectSchema) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetFeatureLimit

`func (o *CreateProjectSchema) GetFeatureLimit() int32`

GetFeatureLimit returns the FeatureLimit field if non-nil, zero value otherwise.

### GetFeatureLimitOk

`func (o *CreateProjectSchema) GetFeatureLimitOk() (*int32, bool)`

GetFeatureLimitOk returns a tuple with the FeatureLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureLimit

`func (o *CreateProjectSchema) SetFeatureLimit(v int32)`

SetFeatureLimit sets FeatureLimit field to given value.

### HasFeatureLimit

`func (o *CreateProjectSchema) HasFeatureLimit() bool`

HasFeatureLimit returns a boolean if a field has been set.

### SetFeatureLimitNil

`func (o *CreateProjectSchema) SetFeatureLimitNil(b bool)`

 SetFeatureLimitNil sets the value for FeatureLimit to be an explicit nil

### UnsetFeatureLimit
`func (o *CreateProjectSchema) UnsetFeatureLimit()`

UnsetFeatureLimit ensures that no value is present for FeatureLimit, not even an explicit nil
### GetMode

`func (o *CreateProjectSchema) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *CreateProjectSchema) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *CreateProjectSchema) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *CreateProjectSchema) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetDefaultStickiness

`func (o *CreateProjectSchema) GetDefaultStickiness() string`

GetDefaultStickiness returns the DefaultStickiness field if non-nil, zero value otherwise.

### GetDefaultStickinessOk

`func (o *CreateProjectSchema) GetDefaultStickinessOk() (*string, bool)`

GetDefaultStickinessOk returns a tuple with the DefaultStickiness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultStickiness

`func (o *CreateProjectSchema) SetDefaultStickiness(v string)`

SetDefaultStickiness sets DefaultStickiness field to given value.

### HasDefaultStickiness

`func (o *CreateProjectSchema) HasDefaultStickiness() bool`

HasDefaultStickiness returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


