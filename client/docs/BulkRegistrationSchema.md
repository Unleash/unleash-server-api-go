# BulkRegistrationSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectVia** | Pointer to [**[]BulkRegistrationSchemaConnectViaInner**](BulkRegistrationSchemaConnectViaInner.md) | A list of applications this app registration has been registered through. If connected directly to Unleash, this is an empty list.   This can be used in later visualizations to tell how many levels of proxy or Edge instances our SDKs have connected through | [optional] 
**AppName** | **string** | The name of the application that is evaluating toggles | 
**Environment** | **string** | Which environment the application is running in | 
**InstanceId** | **string** | A [(somewhat) unique identifier](https://docs.getunleash.io/reference/sdks/node#advanced-usage) for the application | 
**Interval** | Pointer to **float32** | How often (in seconds) the application refreshes its features | [optional] 
**Started** | Pointer to [**DateSchema**](DateSchema.md) |  | [optional] 
**Strategies** | Pointer to **[]string** | Enabled [strategies](https://docs.getunleash.io/reference/activation-strategies) in the application | [optional] 
**SdkVersion** | Pointer to **string** | The version the sdk is running. Typically &lt;client&gt;:&lt;version&gt; | [optional] 

## Methods

### NewBulkRegistrationSchema

`func NewBulkRegistrationSchema(appName string, environment string, instanceId string, ) *BulkRegistrationSchema`

NewBulkRegistrationSchema instantiates a new BulkRegistrationSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkRegistrationSchemaWithDefaults

`func NewBulkRegistrationSchemaWithDefaults() *BulkRegistrationSchema`

NewBulkRegistrationSchemaWithDefaults instantiates a new BulkRegistrationSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectVia

`func (o *BulkRegistrationSchema) GetConnectVia() []BulkRegistrationSchemaConnectViaInner`

GetConnectVia returns the ConnectVia field if non-nil, zero value otherwise.

### GetConnectViaOk

`func (o *BulkRegistrationSchema) GetConnectViaOk() (*[]BulkRegistrationSchemaConnectViaInner, bool)`

GetConnectViaOk returns a tuple with the ConnectVia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectVia

`func (o *BulkRegistrationSchema) SetConnectVia(v []BulkRegistrationSchemaConnectViaInner)`

SetConnectVia sets ConnectVia field to given value.

### HasConnectVia

`func (o *BulkRegistrationSchema) HasConnectVia() bool`

HasConnectVia returns a boolean if a field has been set.

### GetAppName

`func (o *BulkRegistrationSchema) GetAppName() string`

GetAppName returns the AppName field if non-nil, zero value otherwise.

### GetAppNameOk

`func (o *BulkRegistrationSchema) GetAppNameOk() (*string, bool)`

GetAppNameOk returns a tuple with the AppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppName

`func (o *BulkRegistrationSchema) SetAppName(v string)`

SetAppName sets AppName field to given value.


### GetEnvironment

`func (o *BulkRegistrationSchema) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *BulkRegistrationSchema) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *BulkRegistrationSchema) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.


### GetInstanceId

`func (o *BulkRegistrationSchema) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *BulkRegistrationSchema) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *BulkRegistrationSchema) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.


### GetInterval

`func (o *BulkRegistrationSchema) GetInterval() float32`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *BulkRegistrationSchema) GetIntervalOk() (*float32, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *BulkRegistrationSchema) SetInterval(v float32)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *BulkRegistrationSchema) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetStarted

`func (o *BulkRegistrationSchema) GetStarted() DateSchema`

GetStarted returns the Started field if non-nil, zero value otherwise.

### GetStartedOk

`func (o *BulkRegistrationSchema) GetStartedOk() (*DateSchema, bool)`

GetStartedOk returns a tuple with the Started field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarted

`func (o *BulkRegistrationSchema) SetStarted(v DateSchema)`

SetStarted sets Started field to given value.

### HasStarted

`func (o *BulkRegistrationSchema) HasStarted() bool`

HasStarted returns a boolean if a field has been set.

### GetStrategies

`func (o *BulkRegistrationSchema) GetStrategies() []string`

GetStrategies returns the Strategies field if non-nil, zero value otherwise.

### GetStrategiesOk

`func (o *BulkRegistrationSchema) GetStrategiesOk() (*[]string, bool)`

GetStrategiesOk returns a tuple with the Strategies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategies

`func (o *BulkRegistrationSchema) SetStrategies(v []string)`

SetStrategies sets Strategies field to given value.

### HasStrategies

`func (o *BulkRegistrationSchema) HasStrategies() bool`

HasStrategies returns a boolean if a field has been set.

### GetSdkVersion

`func (o *BulkRegistrationSchema) GetSdkVersion() string`

GetSdkVersion returns the SdkVersion field if non-nil, zero value otherwise.

### GetSdkVersionOk

`func (o *BulkRegistrationSchema) GetSdkVersionOk() (*string, bool)`

GetSdkVersionOk returns a tuple with the SdkVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdkVersion

`func (o *BulkRegistrationSchema) SetSdkVersion(v string)`

SetSdkVersion sets SdkVersion field to given value.

### HasSdkVersion

`func (o *BulkRegistrationSchema) HasSdkVersion() bool`

HasSdkVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


