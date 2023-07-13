# BulkMetricsSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Applications** | [**[]BulkRegistrationSchema**](BulkRegistrationSchema.md) | A list of applications registered by an Unleash SDK | 
**Metrics** | [**[]ClientMetricsEnvSchema**](ClientMetricsEnvSchema.md) | a list of client usage metrics registered by downstream providers. (Typically Unleash Edge) | 

## Methods

### NewBulkMetricsSchema

`func NewBulkMetricsSchema(applications []BulkRegistrationSchema, metrics []ClientMetricsEnvSchema, ) *BulkMetricsSchema`

NewBulkMetricsSchema instantiates a new BulkMetricsSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkMetricsSchemaWithDefaults

`func NewBulkMetricsSchemaWithDefaults() *BulkMetricsSchema`

NewBulkMetricsSchemaWithDefaults instantiates a new BulkMetricsSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplications

`func (o *BulkMetricsSchema) GetApplications() []BulkRegistrationSchema`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *BulkMetricsSchema) GetApplicationsOk() (*[]BulkRegistrationSchema, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *BulkMetricsSchema) SetApplications(v []BulkRegistrationSchema)`

SetApplications sets Applications field to given value.


### GetMetrics

`func (o *BulkMetricsSchema) GetMetrics() []ClientMetricsEnvSchema`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *BulkMetricsSchema) GetMetricsOk() (*[]ClientMetricsEnvSchema, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *BulkMetricsSchema) SetMetrics(v []ClientMetricsEnvSchema)`

SetMetrics sets Metrics field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


