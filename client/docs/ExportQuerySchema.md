# ExportQuerySchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Environment** | **string** |  | 
**DownloadFile** | Pointer to **bool** |  | [optional] 

## Methods

### NewExportQuerySchema

`func NewExportQuerySchema(environment string, ) *ExportQuerySchema`

NewExportQuerySchema instantiates a new ExportQuerySchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExportQuerySchemaWithDefaults

`func NewExportQuerySchemaWithDefaults() *ExportQuerySchema`

NewExportQuerySchemaWithDefaults instantiates a new ExportQuerySchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironment

`func (o *ExportQuerySchema) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *ExportQuerySchema) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *ExportQuerySchema) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.


### GetDownloadFile

`func (o *ExportQuerySchema) GetDownloadFile() bool`

GetDownloadFile returns the DownloadFile field if non-nil, zero value otherwise.

### GetDownloadFileOk

`func (o *ExportQuerySchema) GetDownloadFileOk() (*bool, bool)`

GetDownloadFileOk returns a tuple with the DownloadFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadFile

`func (o *ExportQuerySchema) SetDownloadFile(v bool)`

SetDownloadFile sets DownloadFile field to given value.

### HasDownloadFile

`func (o *ExportQuerySchema) HasDownloadFile() bool`

HasDownloadFile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

