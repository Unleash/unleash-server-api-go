# RequestsPerSecondSchemaData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResultType** | Pointer to **string** |  | [optional] 
**Result** | Pointer to [**[]RequestsPerSecondSchemaDataResultInner**](RequestsPerSecondSchemaDataResultInner.md) | An array of values per metric. Each one represents a line in the graph labeled by its metric name | [optional] 

## Methods

### NewRequestsPerSecondSchemaData

`func NewRequestsPerSecondSchemaData() *RequestsPerSecondSchemaData`

NewRequestsPerSecondSchemaData instantiates a new RequestsPerSecondSchemaData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestsPerSecondSchemaDataWithDefaults

`func NewRequestsPerSecondSchemaDataWithDefaults() *RequestsPerSecondSchemaData`

NewRequestsPerSecondSchemaDataWithDefaults instantiates a new RequestsPerSecondSchemaData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResultType

`func (o *RequestsPerSecondSchemaData) GetResultType() string`

GetResultType returns the ResultType field if non-nil, zero value otherwise.

### GetResultTypeOk

`func (o *RequestsPerSecondSchemaData) GetResultTypeOk() (*string, bool)`

GetResultTypeOk returns a tuple with the ResultType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultType

`func (o *RequestsPerSecondSchemaData) SetResultType(v string)`

SetResultType sets ResultType field to given value.

### HasResultType

`func (o *RequestsPerSecondSchemaData) HasResultType() bool`

HasResultType returns a boolean if a field has been set.

### GetResult

`func (o *RequestsPerSecondSchemaData) GetResult() []RequestsPerSecondSchemaDataResultInner`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *RequestsPerSecondSchemaData) GetResultOk() (*[]RequestsPerSecondSchemaDataResultInner, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *RequestsPerSecondSchemaData) SetResult(v []RequestsPerSecondSchemaDataResultInner)`

SetResult sets Result field to given value.

### HasResult

`func (o *RequestsPerSecondSchemaData) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


