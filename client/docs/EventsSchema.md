# EventsSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | **int32** |  | 
**Events** | [**[]EventSchema**](EventSchema.md) |  | 
**TotalEvents** | Pointer to **int32** |  | [optional] 

## Methods

### NewEventsSchema

`func NewEventsSchema(version int32, events []EventSchema, ) *EventsSchema`

NewEventsSchema instantiates a new EventsSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventsSchemaWithDefaults

`func NewEventsSchemaWithDefaults() *EventsSchema`

NewEventsSchemaWithDefaults instantiates a new EventsSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *EventsSchema) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *EventsSchema) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *EventsSchema) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetEvents

`func (o *EventsSchema) GetEvents() []EventSchema`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *EventsSchema) GetEventsOk() (*[]EventSchema, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *EventsSchema) SetEvents(v []EventSchema)`

SetEvents sets Events field to given value.


### GetTotalEvents

`func (o *EventsSchema) GetTotalEvents() int32`

GetTotalEvents returns the TotalEvents field if non-nil, zero value otherwise.

### GetTotalEventsOk

`func (o *EventsSchema) GetTotalEventsOk() (*int32, bool)`

GetTotalEventsOk returns a tuple with the TotalEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalEvents

`func (o *EventsSchema) SetTotalEvents(v int32)`

SetTotalEvents sets TotalEvents field to given value.

### HasTotalEvents

`func (o *EventsSchema) HasTotalEvents() bool`

HasTotalEvents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


