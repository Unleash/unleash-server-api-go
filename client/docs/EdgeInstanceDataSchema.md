# EdgeInstanceDataSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | **string** | The ID of the Edge process, typically a ULID. Newly generated for each restart of the instance. | 
**AppName** | **string** | The name of the application, configured by the user, typically persistent across restarts of Edge. | 
**Region** | Pointer to **NullableString** | Which region the Edge instance is running in. Set to AWS_REGION by default (if present). | [optional] 
**EdgeVersion** | **string** | Which version (semver) of Edge is the Edge instance running. | 
**ProcessMetrics** | Pointer to [**EdgeProcessMetricsSchema**](EdgeProcessMetricsSchema.md) |  | [optional] 
**Started** | **time.Time** | RFC3339 timestamp for when the Edge instance was started. | 
**Traffic** | [**EdgeInstanceTrafficSchema**](EdgeInstanceTrafficSchema.md) |  | 
**LatencyUpstream** | [**EdgeUpstreamLatencySchema**](EdgeUpstreamLatencySchema.md) |  | 
**ConnectedStreamingClients** | **float32** | How many streaming clients are connected to the Edge instance. | 
**ConnectedEdges** | [**[]EdgeInstanceDataSchema**](EdgeInstanceDataSchema.md) | A list of Edge instances connected to the Edge instance. | 

## Methods

### NewEdgeInstanceDataSchema

`func NewEdgeInstanceDataSchema(identifier string, appName string, edgeVersion string, started time.Time, traffic EdgeInstanceTrafficSchema, latencyUpstream EdgeUpstreamLatencySchema, connectedStreamingClients float32, connectedEdges []EdgeInstanceDataSchema, ) *EdgeInstanceDataSchema`

NewEdgeInstanceDataSchema instantiates a new EdgeInstanceDataSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEdgeInstanceDataSchemaWithDefaults

`func NewEdgeInstanceDataSchemaWithDefaults() *EdgeInstanceDataSchema`

NewEdgeInstanceDataSchemaWithDefaults instantiates a new EdgeInstanceDataSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifier

`func (o *EdgeInstanceDataSchema) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *EdgeInstanceDataSchema) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *EdgeInstanceDataSchema) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetAppName

`func (o *EdgeInstanceDataSchema) GetAppName() string`

GetAppName returns the AppName field if non-nil, zero value otherwise.

### GetAppNameOk

`func (o *EdgeInstanceDataSchema) GetAppNameOk() (*string, bool)`

GetAppNameOk returns a tuple with the AppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppName

`func (o *EdgeInstanceDataSchema) SetAppName(v string)`

SetAppName sets AppName field to given value.


### GetRegion

`func (o *EdgeInstanceDataSchema) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *EdgeInstanceDataSchema) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *EdgeInstanceDataSchema) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *EdgeInstanceDataSchema) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### SetRegionNil

`func (o *EdgeInstanceDataSchema) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *EdgeInstanceDataSchema) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetEdgeVersion

`func (o *EdgeInstanceDataSchema) GetEdgeVersion() string`

GetEdgeVersion returns the EdgeVersion field if non-nil, zero value otherwise.

### GetEdgeVersionOk

`func (o *EdgeInstanceDataSchema) GetEdgeVersionOk() (*string, bool)`

GetEdgeVersionOk returns a tuple with the EdgeVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeVersion

`func (o *EdgeInstanceDataSchema) SetEdgeVersion(v string)`

SetEdgeVersion sets EdgeVersion field to given value.


### GetProcessMetrics

`func (o *EdgeInstanceDataSchema) GetProcessMetrics() EdgeProcessMetricsSchema`

GetProcessMetrics returns the ProcessMetrics field if non-nil, zero value otherwise.

### GetProcessMetricsOk

`func (o *EdgeInstanceDataSchema) GetProcessMetricsOk() (*EdgeProcessMetricsSchema, bool)`

GetProcessMetricsOk returns a tuple with the ProcessMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessMetrics

`func (o *EdgeInstanceDataSchema) SetProcessMetrics(v EdgeProcessMetricsSchema)`

SetProcessMetrics sets ProcessMetrics field to given value.

### HasProcessMetrics

`func (o *EdgeInstanceDataSchema) HasProcessMetrics() bool`

HasProcessMetrics returns a boolean if a field has been set.

### GetStarted

`func (o *EdgeInstanceDataSchema) GetStarted() time.Time`

GetStarted returns the Started field if non-nil, zero value otherwise.

### GetStartedOk

`func (o *EdgeInstanceDataSchema) GetStartedOk() (*time.Time, bool)`

GetStartedOk returns a tuple with the Started field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarted

`func (o *EdgeInstanceDataSchema) SetStarted(v time.Time)`

SetStarted sets Started field to given value.


### GetTraffic

`func (o *EdgeInstanceDataSchema) GetTraffic() EdgeInstanceTrafficSchema`

GetTraffic returns the Traffic field if non-nil, zero value otherwise.

### GetTrafficOk

`func (o *EdgeInstanceDataSchema) GetTrafficOk() (*EdgeInstanceTrafficSchema, bool)`

GetTrafficOk returns a tuple with the Traffic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraffic

`func (o *EdgeInstanceDataSchema) SetTraffic(v EdgeInstanceTrafficSchema)`

SetTraffic sets Traffic field to given value.


### GetLatencyUpstream

`func (o *EdgeInstanceDataSchema) GetLatencyUpstream() EdgeUpstreamLatencySchema`

GetLatencyUpstream returns the LatencyUpstream field if non-nil, zero value otherwise.

### GetLatencyUpstreamOk

`func (o *EdgeInstanceDataSchema) GetLatencyUpstreamOk() (*EdgeUpstreamLatencySchema, bool)`

GetLatencyUpstreamOk returns a tuple with the LatencyUpstream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatencyUpstream

`func (o *EdgeInstanceDataSchema) SetLatencyUpstream(v EdgeUpstreamLatencySchema)`

SetLatencyUpstream sets LatencyUpstream field to given value.


### GetConnectedStreamingClients

`func (o *EdgeInstanceDataSchema) GetConnectedStreamingClients() float32`

GetConnectedStreamingClients returns the ConnectedStreamingClients field if non-nil, zero value otherwise.

### GetConnectedStreamingClientsOk

`func (o *EdgeInstanceDataSchema) GetConnectedStreamingClientsOk() (*float32, bool)`

GetConnectedStreamingClientsOk returns a tuple with the ConnectedStreamingClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedStreamingClients

`func (o *EdgeInstanceDataSchema) SetConnectedStreamingClients(v float32)`

SetConnectedStreamingClients sets ConnectedStreamingClients field to given value.


### GetConnectedEdges

`func (o *EdgeInstanceDataSchema) GetConnectedEdges() []EdgeInstanceDataSchema`

GetConnectedEdges returns the ConnectedEdges field if non-nil, zero value otherwise.

### GetConnectedEdgesOk

`func (o *EdgeInstanceDataSchema) GetConnectedEdgesOk() (*[]EdgeInstanceDataSchema, bool)`

GetConnectedEdgesOk returns a tuple with the ConnectedEdges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedEdges

`func (o *EdgeInstanceDataSchema) SetConnectedEdges(v []EdgeInstanceDataSchema)`

SetConnectedEdges sets ConnectedEdges field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


