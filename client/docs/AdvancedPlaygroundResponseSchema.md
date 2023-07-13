# AdvancedPlaygroundResponseSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Input** | [**AdvancedPlaygroundRequestSchema**](AdvancedPlaygroundRequestSchema.md) |  | 
**Features** | [**[]AdvancedPlaygroundFeatureSchema**](AdvancedPlaygroundFeatureSchema.md) | The list of features that have been evaluated. | 

## Methods

### NewAdvancedPlaygroundResponseSchema

`func NewAdvancedPlaygroundResponseSchema(input AdvancedPlaygroundRequestSchema, features []AdvancedPlaygroundFeatureSchema, ) *AdvancedPlaygroundResponseSchema`

NewAdvancedPlaygroundResponseSchema instantiates a new AdvancedPlaygroundResponseSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvancedPlaygroundResponseSchemaWithDefaults

`func NewAdvancedPlaygroundResponseSchemaWithDefaults() *AdvancedPlaygroundResponseSchema`

NewAdvancedPlaygroundResponseSchemaWithDefaults instantiates a new AdvancedPlaygroundResponseSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInput

`func (o *AdvancedPlaygroundResponseSchema) GetInput() AdvancedPlaygroundRequestSchema`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *AdvancedPlaygroundResponseSchema) GetInputOk() (*AdvancedPlaygroundRequestSchema, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *AdvancedPlaygroundResponseSchema) SetInput(v AdvancedPlaygroundRequestSchema)`

SetInput sets Input field to given value.


### GetFeatures

`func (o *AdvancedPlaygroundResponseSchema) GetFeatures() []AdvancedPlaygroundFeatureSchema`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *AdvancedPlaygroundResponseSchema) GetFeaturesOk() (*[]AdvancedPlaygroundFeatureSchema, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *AdvancedPlaygroundResponseSchema) SetFeatures(v []AdvancedPlaygroundFeatureSchema)`

SetFeatures sets Features field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


