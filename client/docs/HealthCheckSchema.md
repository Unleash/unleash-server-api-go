# HealthCheckSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Health** | **string** | The state this Unleash instance is in. GOOD if everything is ok, BAD if the instance should be restarted | 

## Methods

### NewHealthCheckSchema

`func NewHealthCheckSchema(health string, ) *HealthCheckSchema`

NewHealthCheckSchema instantiates a new HealthCheckSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHealthCheckSchemaWithDefaults

`func NewHealthCheckSchemaWithDefaults() *HealthCheckSchema`

NewHealthCheckSchemaWithDefaults instantiates a new HealthCheckSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHealth

`func (o *HealthCheckSchema) GetHealth() string`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *HealthCheckSchema) GetHealthOk() (*string, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *HealthCheckSchema) SetHealth(v string)`

SetHealth sets Health field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


