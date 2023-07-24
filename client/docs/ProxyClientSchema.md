# ProxyClientSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppName** | **string** | Name of the application using Unleash | 
**InstanceId** | Pointer to **string** | Instance id for this application (typically hostname, podId or similar) | [optional] 
**SdkVersion** | Pointer to **string** | Optional field that describes the sdk version (name:version) | [optional] 
**Environment** | Pointer to **string** |  | [optional] 
**Interval** | **float32** | At which interval, in milliseconds, will this client be expected to send metrics | 
**Started** | [**ProxyClientSchemaStarted**](ProxyClientSchemaStarted.md) |  | 
**Strategies** | **[]string** | List of strategies implemented by this application | 

## Methods

### NewProxyClientSchema

`func NewProxyClientSchema(appName string, interval float32, started ProxyClientSchemaStarted, strategies []string, ) *ProxyClientSchema`

NewProxyClientSchema instantiates a new ProxyClientSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProxyClientSchemaWithDefaults

`func NewProxyClientSchemaWithDefaults() *ProxyClientSchema`

NewProxyClientSchemaWithDefaults instantiates a new ProxyClientSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppName

`func (o *ProxyClientSchema) GetAppName() string`

GetAppName returns the AppName field if non-nil, zero value otherwise.

### GetAppNameOk

`func (o *ProxyClientSchema) GetAppNameOk() (*string, bool)`

GetAppNameOk returns a tuple with the AppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppName

`func (o *ProxyClientSchema) SetAppName(v string)`

SetAppName sets AppName field to given value.


### GetInstanceId

`func (o *ProxyClientSchema) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *ProxyClientSchema) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *ProxyClientSchema) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *ProxyClientSchema) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### GetSdkVersion

`func (o *ProxyClientSchema) GetSdkVersion() string`

GetSdkVersion returns the SdkVersion field if non-nil, zero value otherwise.

### GetSdkVersionOk

`func (o *ProxyClientSchema) GetSdkVersionOk() (*string, bool)`

GetSdkVersionOk returns a tuple with the SdkVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdkVersion

`func (o *ProxyClientSchema) SetSdkVersion(v string)`

SetSdkVersion sets SdkVersion field to given value.

### HasSdkVersion

`func (o *ProxyClientSchema) HasSdkVersion() bool`

HasSdkVersion returns a boolean if a field has been set.

### GetEnvironment

`func (o *ProxyClientSchema) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *ProxyClientSchema) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *ProxyClientSchema) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *ProxyClientSchema) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetInterval

`func (o *ProxyClientSchema) GetInterval() float32`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *ProxyClientSchema) GetIntervalOk() (*float32, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *ProxyClientSchema) SetInterval(v float32)`

SetInterval sets Interval field to given value.


### GetStarted

`func (o *ProxyClientSchema) GetStarted() ProxyClientSchemaStarted`

GetStarted returns the Started field if non-nil, zero value otherwise.

### GetStartedOk

`func (o *ProxyClientSchema) GetStartedOk() (*ProxyClientSchemaStarted, bool)`

GetStartedOk returns a tuple with the Started field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarted

`func (o *ProxyClientSchema) SetStarted(v ProxyClientSchemaStarted)`

SetStarted sets Started field to given value.


### GetStrategies

`func (o *ProxyClientSchema) GetStrategies() []string`

GetStrategies returns the Strategies field if non-nil, zero value otherwise.

### GetStrategiesOk

`func (o *ProxyClientSchema) GetStrategiesOk() (*[]string, bool)`

GetStrategiesOk returns a tuple with the Strategies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategies

`func (o *ProxyClientSchema) SetStrategies(v []string)`

SetStrategies sets Strategies field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


