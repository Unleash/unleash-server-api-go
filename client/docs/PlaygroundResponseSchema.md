# PlaygroundResponseSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Input** | [**PlaygroundRequestSchema**](PlaygroundRequestSchema.md) |  | 
**Features** | [**[]PlaygroundFeatureSchema**](PlaygroundFeatureSchema.md) | The list of features that have been evaluated. | 

## Methods

### NewPlaygroundResponseSchema

`func NewPlaygroundResponseSchema(input PlaygroundRequestSchema, features []PlaygroundFeatureSchema, ) *PlaygroundResponseSchema`

NewPlaygroundResponseSchema instantiates a new PlaygroundResponseSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlaygroundResponseSchemaWithDefaults

`func NewPlaygroundResponseSchemaWithDefaults() *PlaygroundResponseSchema`

NewPlaygroundResponseSchemaWithDefaults instantiates a new PlaygroundResponseSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInput

`func (o *PlaygroundResponseSchema) GetInput() PlaygroundRequestSchema`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *PlaygroundResponseSchema) GetInputOk() (*PlaygroundRequestSchema, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *PlaygroundResponseSchema) SetInput(v PlaygroundRequestSchema)`

SetInput sets Input field to given value.


### GetFeatures

`func (o *PlaygroundResponseSchema) GetFeatures() []PlaygroundFeatureSchema`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *PlaygroundResponseSchema) GetFeaturesOk() (*[]PlaygroundFeatureSchema, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *PlaygroundResponseSchema) SetFeatures(v []PlaygroundFeatureSchema)`

SetFeatures sets Features field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


