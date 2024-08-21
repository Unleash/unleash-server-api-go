# ProjectSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The id of this project | 
**Name** | **string** | The name of this project | 
**Description** | Pointer to **NullableString** | Additional information about the project | [optional] 
**Health** | Pointer to **float32** | An indicator of the [project&#39;s health](https://docs.getunleash.io/reference/technical-debt#health-rating) on a scale from 0 to 100 | [optional] 
**FeatureCount** | Pointer to **float32** | The number of features this project has | [optional] 

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


