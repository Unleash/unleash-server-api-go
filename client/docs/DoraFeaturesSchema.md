# DoraFeaturesSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of a feature toggle | 
**TimeToProduction** | **float32** | The average number of days it takes a feature toggle to get into production | 

## Methods

### NewDoraFeaturesSchema

`func NewDoraFeaturesSchema(name string, timeToProduction float32, ) *DoraFeaturesSchema`

NewDoraFeaturesSchema instantiates a new DoraFeaturesSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDoraFeaturesSchemaWithDefaults

`func NewDoraFeaturesSchemaWithDefaults() *DoraFeaturesSchema`

NewDoraFeaturesSchemaWithDefaults instantiates a new DoraFeaturesSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DoraFeaturesSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DoraFeaturesSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DoraFeaturesSchema) SetName(v string)`

SetName sets Name field to given value.


### GetTimeToProduction

`func (o *DoraFeaturesSchema) GetTimeToProduction() float32`

GetTimeToProduction returns the TimeToProduction field if non-nil, zero value otherwise.

### GetTimeToProductionOk

`func (o *DoraFeaturesSchema) GetTimeToProductionOk() (*float32, bool)`

GetTimeToProductionOk returns a tuple with the TimeToProduction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeToProduction

`func (o *DoraFeaturesSchema) SetTimeToProduction(v float32)`

SetTimeToProduction sets TimeToProduction field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


