# RequestsPerSecondSchemaDataResultInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metric** | Pointer to [**RequestsPerSecondSchemaDataResultInnerMetric**](RequestsPerSecondSchemaDataResultInnerMetric.md) |  | [optional] 
**Values** | Pointer to [**[][]RequestsPerSecondSchemaDataResultInnerValuesInnerInner**]([]RequestsPerSecondSchemaDataResultInnerValuesInnerInner.md) | An array of arrays. Each element of the array is an array of size 2 consisting of the 2 axis for the graph: in position zero the x axis represented as a number and position one the y axis represented as string | [optional] 

## Methods

### NewRequestsPerSecondSchemaDataResultInner

`func NewRequestsPerSecondSchemaDataResultInner() *RequestsPerSecondSchemaDataResultInner`

NewRequestsPerSecondSchemaDataResultInner instantiates a new RequestsPerSecondSchemaDataResultInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestsPerSecondSchemaDataResultInnerWithDefaults

`func NewRequestsPerSecondSchemaDataResultInnerWithDefaults() *RequestsPerSecondSchemaDataResultInner`

NewRequestsPerSecondSchemaDataResultInnerWithDefaults instantiates a new RequestsPerSecondSchemaDataResultInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetric

`func (o *RequestsPerSecondSchemaDataResultInner) GetMetric() RequestsPerSecondSchemaDataResultInnerMetric`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *RequestsPerSecondSchemaDataResultInner) GetMetricOk() (*RequestsPerSecondSchemaDataResultInnerMetric, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *RequestsPerSecondSchemaDataResultInner) SetMetric(v RequestsPerSecondSchemaDataResultInnerMetric)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *RequestsPerSecondSchemaDataResultInner) HasMetric() bool`

HasMetric returns a boolean if a field has been set.

### GetValues

`func (o *RequestsPerSecondSchemaDataResultInner) GetValues() [][]RequestsPerSecondSchemaDataResultInnerValuesInnerInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *RequestsPerSecondSchemaDataResultInner) GetValuesOk() (*[][]RequestsPerSecondSchemaDataResultInnerValuesInnerInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *RequestsPerSecondSchemaDataResultInner) SetValues(v [][]RequestsPerSecondSchemaDataResultInnerValuesInnerInner)`

SetValues sets Values field to given value.

### HasValues

`func (o *RequestsPerSecondSchemaDataResultInner) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


