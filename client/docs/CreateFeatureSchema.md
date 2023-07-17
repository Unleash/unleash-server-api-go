# CreateFeatureSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Unique feature name | 
**Type** | Pointer to **string** | The feature toggle&#39;s [type](https://docs.getunleash.io/reference/feature-toggle-types). One of experiment, kill-switch, release, operational, or permission | [optional] 
**Description** | Pointer to **NullableString** | Detailed description of the feature | [optional] 
**ImpressionData** | Pointer to **bool** | &#x60;true&#x60; if the impression data collection is enabled for the feature, otherwise &#x60;false&#x60;. | [optional] 

## Methods

### NewCreateFeatureSchema

`func NewCreateFeatureSchema(name string, ) *CreateFeatureSchema`

NewCreateFeatureSchema instantiates a new CreateFeatureSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateFeatureSchemaWithDefaults

`func NewCreateFeatureSchemaWithDefaults() *CreateFeatureSchema`

NewCreateFeatureSchemaWithDefaults instantiates a new CreateFeatureSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateFeatureSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateFeatureSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateFeatureSchema) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *CreateFeatureSchema) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateFeatureSchema) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateFeatureSchema) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CreateFeatureSchema) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDescription

`func (o *CreateFeatureSchema) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateFeatureSchema) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateFeatureSchema) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateFeatureSchema) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreateFeatureSchema) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreateFeatureSchema) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetImpressionData

`func (o *CreateFeatureSchema) GetImpressionData() bool`

GetImpressionData returns the ImpressionData field if non-nil, zero value otherwise.

### GetImpressionDataOk

`func (o *CreateFeatureSchema) GetImpressionDataOk() (*bool, bool)`

GetImpressionDataOk returns a tuple with the ImpressionData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpressionData

`func (o *CreateFeatureSchema) SetImpressionData(v bool)`

SetImpressionData sets ImpressionData field to given value.

### HasImpressionData

`func (o *CreateFeatureSchema) HasImpressionData() bool`

HasImpressionData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


