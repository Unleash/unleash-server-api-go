# SearchEventsSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | Find events by event type (case-sensitive). | [optional] 
**Project** | Pointer to **string** | Find events by project ID (case-sensitive). | [optional] 
**Feature** | Pointer to **string** | Find events by feature toggle name (case-sensitive). | [optional] 
**Query** | Pointer to **string** |                  Find events by a free-text search query.                 The query will be matched against the event type,                 the username or email that created the event (if any),                 and the event data payload (if any).              | [optional] 
**Limit** | Pointer to **int32** | The maximum amount of events to return in the search result | [optional] [default to 100]
**Offset** | Pointer to **int32** | Which event id to start listing from | [optional] [default to 0]

## Methods

### NewSearchEventsSchema

`func NewSearchEventsSchema() *SearchEventsSchema`

NewSearchEventsSchema instantiates a new SearchEventsSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchEventsSchemaWithDefaults

`func NewSearchEventsSchemaWithDefaults() *SearchEventsSchema`

NewSearchEventsSchemaWithDefaults instantiates a new SearchEventsSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SearchEventsSchema) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SearchEventsSchema) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SearchEventsSchema) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SearchEventsSchema) HasType() bool`

HasType returns a boolean if a field has been set.

### GetProject

`func (o *SearchEventsSchema) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *SearchEventsSchema) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *SearchEventsSchema) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *SearchEventsSchema) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetFeature

`func (o *SearchEventsSchema) GetFeature() string`

GetFeature returns the Feature field if non-nil, zero value otherwise.

### GetFeatureOk

`func (o *SearchEventsSchema) GetFeatureOk() (*string, bool)`

GetFeatureOk returns a tuple with the Feature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeature

`func (o *SearchEventsSchema) SetFeature(v string)`

SetFeature sets Feature field to given value.

### HasFeature

`func (o *SearchEventsSchema) HasFeature() bool`

HasFeature returns a boolean if a field has been set.

### GetQuery

`func (o *SearchEventsSchema) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *SearchEventsSchema) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *SearchEventsSchema) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *SearchEventsSchema) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetLimit

`func (o *SearchEventsSchema) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *SearchEventsSchema) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *SearchEventsSchema) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *SearchEventsSchema) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *SearchEventsSchema) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *SearchEventsSchema) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *SearchEventsSchema) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *SearchEventsSchema) HasOffset() bool`

HasOffset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


