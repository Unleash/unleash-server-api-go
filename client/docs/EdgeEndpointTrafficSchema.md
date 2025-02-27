# EdgeEndpointTrafficSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Requests200** | Pointer to **float32** | Number of 20x requests | [optional] 
**Requests304** | Pointer to **float32** | Number of 30x requests | [optional] 

## Methods

### NewEdgeEndpointTrafficSchema

`func NewEdgeEndpointTrafficSchema() *EdgeEndpointTrafficSchema`

NewEdgeEndpointTrafficSchema instantiates a new EdgeEndpointTrafficSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEdgeEndpointTrafficSchemaWithDefaults

`func NewEdgeEndpointTrafficSchemaWithDefaults() *EdgeEndpointTrafficSchema`

NewEdgeEndpointTrafficSchemaWithDefaults instantiates a new EdgeEndpointTrafficSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequests200

`func (o *EdgeEndpointTrafficSchema) GetRequests200() float32`

GetRequests200 returns the Requests200 field if non-nil, zero value otherwise.

### GetRequests200Ok

`func (o *EdgeEndpointTrafficSchema) GetRequests200Ok() (*float32, bool)`

GetRequests200Ok returns a tuple with the Requests200 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests200

`func (o *EdgeEndpointTrafficSchema) SetRequests200(v float32)`

SetRequests200 sets Requests200 field to given value.

### HasRequests200

`func (o *EdgeEndpointTrafficSchema) HasRequests200() bool`

HasRequests200 returns a boolean if a field has been set.

### GetRequests304

`func (o *EdgeEndpointTrafficSchema) GetRequests304() float32`

GetRequests304 returns the Requests304 field if non-nil, zero value otherwise.

### GetRequests304Ok

`func (o *EdgeEndpointTrafficSchema) GetRequests304Ok() (*float32, bool)`

GetRequests304Ok returns a tuple with the Requests304 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests304

`func (o *EdgeEndpointTrafficSchema) SetRequests304(v float32)`

SetRequests304 sets Requests304 field to given value.

### HasRequests304

`func (o *EdgeEndpointTrafficSchema) HasRequests304() bool`

HasRequests304 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


