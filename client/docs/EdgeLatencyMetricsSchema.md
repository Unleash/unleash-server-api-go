# EdgeLatencyMetricsSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Avg** | **float32** | Average time per request in milliseconds. | 
**Count** | **float32** | Total number of requests made. | 
**P99** | **float32** | 99% of requests finished within this amount of milliseconds. | 

## Methods

### NewEdgeLatencyMetricsSchema

`func NewEdgeLatencyMetricsSchema(avg float32, count float32, p99 float32, ) *EdgeLatencyMetricsSchema`

NewEdgeLatencyMetricsSchema instantiates a new EdgeLatencyMetricsSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEdgeLatencyMetricsSchemaWithDefaults

`func NewEdgeLatencyMetricsSchemaWithDefaults() *EdgeLatencyMetricsSchema`

NewEdgeLatencyMetricsSchemaWithDefaults instantiates a new EdgeLatencyMetricsSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvg

`func (o *EdgeLatencyMetricsSchema) GetAvg() float32`

GetAvg returns the Avg field if non-nil, zero value otherwise.

### GetAvgOk

`func (o *EdgeLatencyMetricsSchema) GetAvgOk() (*float32, bool)`

GetAvgOk returns a tuple with the Avg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvg

`func (o *EdgeLatencyMetricsSchema) SetAvg(v float32)`

SetAvg sets Avg field to given value.


### GetCount

`func (o *EdgeLatencyMetricsSchema) GetCount() float32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *EdgeLatencyMetricsSchema) GetCountOk() (*float32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *EdgeLatencyMetricsSchema) SetCount(v float32)`

SetCount sets Count field to given value.


### GetP99

`func (o *EdgeLatencyMetricsSchema) GetP99() float32`

GetP99 returns the P99 field if non-nil, zero value otherwise.

### GetP99Ok

`func (o *EdgeLatencyMetricsSchema) GetP99Ok() (*float32, bool)`

GetP99Ok returns a tuple with the P99 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP99

`func (o *EdgeLatencyMetricsSchema) SetP99(v float32)`

SetP99 sets P99 field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


