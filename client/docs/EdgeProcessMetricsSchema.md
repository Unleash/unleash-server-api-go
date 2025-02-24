# EdgeProcessMetricsSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CpuUsage** | **float32** | CPU usage, in seconds, since start of process. | 
**MemoryUsage** | **float32** | Current process_resident_memory (in bytes) usage. | 

## Methods

### NewEdgeProcessMetricsSchema

`func NewEdgeProcessMetricsSchema(cpuUsage float32, memoryUsage float32, ) *EdgeProcessMetricsSchema`

NewEdgeProcessMetricsSchema instantiates a new EdgeProcessMetricsSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEdgeProcessMetricsSchemaWithDefaults

`func NewEdgeProcessMetricsSchemaWithDefaults() *EdgeProcessMetricsSchema`

NewEdgeProcessMetricsSchemaWithDefaults instantiates a new EdgeProcessMetricsSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpuUsage

`func (o *EdgeProcessMetricsSchema) GetCpuUsage() float32`

GetCpuUsage returns the CpuUsage field if non-nil, zero value otherwise.

### GetCpuUsageOk

`func (o *EdgeProcessMetricsSchema) GetCpuUsageOk() (*float32, bool)`

GetCpuUsageOk returns a tuple with the CpuUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuUsage

`func (o *EdgeProcessMetricsSchema) SetCpuUsage(v float32)`

SetCpuUsage sets CpuUsage field to given value.


### GetMemoryUsage

`func (o *EdgeProcessMetricsSchema) GetMemoryUsage() float32`

GetMemoryUsage returns the MemoryUsage field if non-nil, zero value otherwise.

### GetMemoryUsageOk

`func (o *EdgeProcessMetricsSchema) GetMemoryUsageOk() (*float32, bool)`

GetMemoryUsageOk returns a tuple with the MemoryUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryUsage

`func (o *EdgeProcessMetricsSchema) SetMemoryUsage(v float32)`

SetMemoryUsage sets MemoryUsage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


