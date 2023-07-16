# ClientFeaturesSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | **float32** | A version number for the format used in the response. Most Unleash instances now return version 2, which includes segments as a separate array | 
**Features** | [**[]ClientFeatureSchema**](ClientFeatureSchema.md) | A list of feature toggles with their configuration | 
**Segments** | Pointer to [**[]ClientSegmentSchema**](ClientSegmentSchema.md) | A list of [Segments](https://docs.getunleash.io/reference/segments) configured for this Unleash instance | [optional] 
**Query** | Pointer to [**ClientFeaturesQuerySchema**](ClientFeaturesQuerySchema.md) |  | [optional] 

## Methods

### NewClientFeaturesSchema

`func NewClientFeaturesSchema(version float32, features []ClientFeatureSchema, ) *ClientFeaturesSchema`

NewClientFeaturesSchema instantiates a new ClientFeaturesSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientFeaturesSchemaWithDefaults

`func NewClientFeaturesSchemaWithDefaults() *ClientFeaturesSchema`

NewClientFeaturesSchemaWithDefaults instantiates a new ClientFeaturesSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *ClientFeaturesSchema) GetVersion() float32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ClientFeaturesSchema) GetVersionOk() (*float32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ClientFeaturesSchema) SetVersion(v float32)`

SetVersion sets Version field to given value.


### GetFeatures

`func (o *ClientFeaturesSchema) GetFeatures() []ClientFeatureSchema`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *ClientFeaturesSchema) GetFeaturesOk() (*[]ClientFeatureSchema, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *ClientFeaturesSchema) SetFeatures(v []ClientFeatureSchema)`

SetFeatures sets Features field to given value.


### GetSegments

`func (o *ClientFeaturesSchema) GetSegments() []ClientSegmentSchema`

GetSegments returns the Segments field if non-nil, zero value otherwise.

### GetSegmentsOk

`func (o *ClientFeaturesSchema) GetSegmentsOk() (*[]ClientSegmentSchema, bool)`

GetSegmentsOk returns a tuple with the Segments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegments

`func (o *ClientFeaturesSchema) SetSegments(v []ClientSegmentSchema)`

SetSegments sets Segments field to given value.

### HasSegments

`func (o *ClientFeaturesSchema) HasSegments() bool`

HasSegments returns a boolean if a field has been set.

### GetQuery

`func (o *ClientFeaturesSchema) GetQuery() ClientFeaturesQuerySchema`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *ClientFeaturesSchema) GetQueryOk() (*ClientFeaturesQuerySchema, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *ClientFeaturesSchema) SetQuery(v ClientFeaturesQuerySchema)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *ClientFeaturesSchema) HasQuery() bool`

HasQuery returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


