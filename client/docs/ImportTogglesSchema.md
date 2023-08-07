# ImportTogglesSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Project** | **string** | The exported [project](https://docs.getunleash.io/reference/projects) | 
**Environment** | **string** | The exported [environment](https://docs.getunleash.io/reference/environments) | 
**Data** | [**ExportResultSchema**](ExportResultSchema.md) |  | 

## Methods

### NewImportTogglesSchema

`func NewImportTogglesSchema(project string, environment string, data ExportResultSchema, ) *ImportTogglesSchema`

NewImportTogglesSchema instantiates a new ImportTogglesSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportTogglesSchemaWithDefaults

`func NewImportTogglesSchemaWithDefaults() *ImportTogglesSchema`

NewImportTogglesSchemaWithDefaults instantiates a new ImportTogglesSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProject

`func (o *ImportTogglesSchema) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ImportTogglesSchema) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ImportTogglesSchema) SetProject(v string)`

SetProject sets Project field to given value.


### GetEnvironment

`func (o *ImportTogglesSchema) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *ImportTogglesSchema) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *ImportTogglesSchema) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.


### GetData

`func (o *ImportTogglesSchema) GetData() ExportResultSchema`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ImportTogglesSchema) GetDataOk() (*ExportResultSchema, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ImportTogglesSchema) SetData(v ExportResultSchema)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


