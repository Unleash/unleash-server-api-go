# PlaygroundStrategySchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The strategy&#39;s name. | 
**Title** | Pointer to **string** | Description of the feature&#39;s purpose. | [optional] 
**Id** | **string** | The strategy&#39;s id. | 
**Result** | [**PlaygroundStrategySchemaResult**](PlaygroundStrategySchemaResult.md) |  | 
**Disabled** | **NullableBool** | The strategy&#39;s status. Disabled strategies are not evaluated | 
**Segments** | [**[]PlaygroundSegmentSchema**](PlaygroundSegmentSchema.md) | The strategy&#39;s segments and their evaluation results. | 
**Constraints** | [**[]PlaygroundConstraintSchema**](PlaygroundConstraintSchema.md) | The strategy&#39;s constraints and their evaluation results. | 
**Parameters** | **map[string]string** | A list of parameters for a strategy | 
**Links** | [**PlaygroundStrategySchemaLinks**](PlaygroundStrategySchemaLinks.md) |  | 

## Methods

### NewPlaygroundStrategySchema

`func NewPlaygroundStrategySchema(name string, id string, result PlaygroundStrategySchemaResult, disabled NullableBool, segments []PlaygroundSegmentSchema, constraints []PlaygroundConstraintSchema, parameters map[string]string, links PlaygroundStrategySchemaLinks, ) *PlaygroundStrategySchema`

NewPlaygroundStrategySchema instantiates a new PlaygroundStrategySchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlaygroundStrategySchemaWithDefaults

`func NewPlaygroundStrategySchemaWithDefaults() *PlaygroundStrategySchema`

NewPlaygroundStrategySchemaWithDefaults instantiates a new PlaygroundStrategySchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PlaygroundStrategySchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PlaygroundStrategySchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PlaygroundStrategySchema) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *PlaygroundStrategySchema) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PlaygroundStrategySchema) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PlaygroundStrategySchema) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *PlaygroundStrategySchema) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetId

`func (o *PlaygroundStrategySchema) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PlaygroundStrategySchema) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PlaygroundStrategySchema) SetId(v string)`

SetId sets Id field to given value.


### GetResult

`func (o *PlaygroundStrategySchema) GetResult() PlaygroundStrategySchemaResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *PlaygroundStrategySchema) GetResultOk() (*PlaygroundStrategySchemaResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *PlaygroundStrategySchema) SetResult(v PlaygroundStrategySchemaResult)`

SetResult sets Result field to given value.


### GetDisabled

`func (o *PlaygroundStrategySchema) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *PlaygroundStrategySchema) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *PlaygroundStrategySchema) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.


### SetDisabledNil

`func (o *PlaygroundStrategySchema) SetDisabledNil(b bool)`

 SetDisabledNil sets the value for Disabled to be an explicit nil

### UnsetDisabled
`func (o *PlaygroundStrategySchema) UnsetDisabled()`

UnsetDisabled ensures that no value is present for Disabled, not even an explicit nil
### GetSegments

`func (o *PlaygroundStrategySchema) GetSegments() []PlaygroundSegmentSchema`

GetSegments returns the Segments field if non-nil, zero value otherwise.

### GetSegmentsOk

`func (o *PlaygroundStrategySchema) GetSegmentsOk() (*[]PlaygroundSegmentSchema, bool)`

GetSegmentsOk returns a tuple with the Segments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegments

`func (o *PlaygroundStrategySchema) SetSegments(v []PlaygroundSegmentSchema)`

SetSegments sets Segments field to given value.


### GetConstraints

`func (o *PlaygroundStrategySchema) GetConstraints() []PlaygroundConstraintSchema`

GetConstraints returns the Constraints field if non-nil, zero value otherwise.

### GetConstraintsOk

`func (o *PlaygroundStrategySchema) GetConstraintsOk() (*[]PlaygroundConstraintSchema, bool)`

GetConstraintsOk returns a tuple with the Constraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraints

`func (o *PlaygroundStrategySchema) SetConstraints(v []PlaygroundConstraintSchema)`

SetConstraints sets Constraints field to given value.


### GetParameters

`func (o *PlaygroundStrategySchema) GetParameters() map[string]string`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *PlaygroundStrategySchema) GetParametersOk() (*map[string]string, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *PlaygroundStrategySchema) SetParameters(v map[string]string)`

SetParameters sets Parameters field to given value.


### GetLinks

`func (o *PlaygroundStrategySchema) GetLinks() PlaygroundStrategySchemaLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PlaygroundStrategySchema) GetLinksOk() (*PlaygroundStrategySchemaLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PlaygroundStrategySchema) SetLinks(v PlaygroundStrategySchemaLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


