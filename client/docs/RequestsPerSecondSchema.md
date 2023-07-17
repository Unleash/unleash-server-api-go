# RequestsPerSecondSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | Whether the query against prometheus succeeded or failed | [optional] 
**Data** | Pointer to [**RequestsPerSecondSchemaData**](RequestsPerSecondSchemaData.md) |  | [optional] 

## Methods

### NewRequestsPerSecondSchema

`func NewRequestsPerSecondSchema() *RequestsPerSecondSchema`

NewRequestsPerSecondSchema instantiates a new RequestsPerSecondSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestsPerSecondSchemaWithDefaults

`func NewRequestsPerSecondSchemaWithDefaults() *RequestsPerSecondSchema`

NewRequestsPerSecondSchemaWithDefaults instantiates a new RequestsPerSecondSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *RequestsPerSecondSchema) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RequestsPerSecondSchema) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RequestsPerSecondSchema) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RequestsPerSecondSchema) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetData

`func (o *RequestsPerSecondSchema) GetData() RequestsPerSecondSchemaData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *RequestsPerSecondSchema) GetDataOk() (*RequestsPerSecondSchemaData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *RequestsPerSecondSchema) SetData(v RequestsPerSecondSchemaData)`

SetData sets Data field to given value.

### HasData

`func (o *RequestsPerSecondSchema) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


