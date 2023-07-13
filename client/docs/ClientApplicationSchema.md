# ClientApplicationSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppName** | **string** | An identifier for the app that uses the sdk, should be static across SDK restarts | 
**InstanceId** | Pointer to **string** | A unique identifier identifying the instance of the application running the SDK. Often changes based on execution environment. For instance: two pods in Kubernetes will have two different instanceIds | [optional] 
**SdkVersion** | Pointer to **string** | An SDK version identifier. Usually formatted as \&quot;unleash-client-&lt;language&gt;:&lt;version&gt;\&quot; | [optional] 
**Environment** | Pointer to **string** | The SDK&#39;s configured &#39;environment&#39; property. Deprecated. This property  does **not** control which Unleash environment the SDK gets toggles for. To control Unleash environments, use the SDKs API key. | [optional] 
**Interval** | **float32** | How often (in seconds) does the client refresh its toggles | 
**Started** | [**ClientApplicationSchemaStarted**](ClientApplicationSchemaStarted.md) |  | 
**Strategies** | **[]string** | Which strategies the SDKs runtime knows about | 

## Methods

### NewClientApplicationSchema

`func NewClientApplicationSchema(appName string, interval float32, started ClientApplicationSchemaStarted, strategies []string, ) *ClientApplicationSchema`

NewClientApplicationSchema instantiates a new ClientApplicationSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientApplicationSchemaWithDefaults

`func NewClientApplicationSchemaWithDefaults() *ClientApplicationSchema`

NewClientApplicationSchemaWithDefaults instantiates a new ClientApplicationSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppName

`func (o *ClientApplicationSchema) GetAppName() string`

GetAppName returns the AppName field if non-nil, zero value otherwise.

### GetAppNameOk

`func (o *ClientApplicationSchema) GetAppNameOk() (*string, bool)`

GetAppNameOk returns a tuple with the AppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppName

`func (o *ClientApplicationSchema) SetAppName(v string)`

SetAppName sets AppName field to given value.


### GetInstanceId

`func (o *ClientApplicationSchema) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *ClientApplicationSchema) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *ClientApplicationSchema) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *ClientApplicationSchema) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### GetSdkVersion

`func (o *ClientApplicationSchema) GetSdkVersion() string`

GetSdkVersion returns the SdkVersion field if non-nil, zero value otherwise.

### GetSdkVersionOk

`func (o *ClientApplicationSchema) GetSdkVersionOk() (*string, bool)`

GetSdkVersionOk returns a tuple with the SdkVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdkVersion

`func (o *ClientApplicationSchema) SetSdkVersion(v string)`

SetSdkVersion sets SdkVersion field to given value.

### HasSdkVersion

`func (o *ClientApplicationSchema) HasSdkVersion() bool`

HasSdkVersion returns a boolean if a field has been set.

### GetEnvironment

`func (o *ClientApplicationSchema) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *ClientApplicationSchema) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *ClientApplicationSchema) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *ClientApplicationSchema) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetInterval

`func (o *ClientApplicationSchema) GetInterval() float32`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *ClientApplicationSchema) GetIntervalOk() (*float32, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *ClientApplicationSchema) SetInterval(v float32)`

SetInterval sets Interval field to given value.


### GetStarted

`func (o *ClientApplicationSchema) GetStarted() ClientApplicationSchemaStarted`

GetStarted returns the Started field if non-nil, zero value otherwise.

### GetStartedOk

`func (o *ClientApplicationSchema) GetStartedOk() (*ClientApplicationSchemaStarted, bool)`

GetStartedOk returns a tuple with the Started field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarted

`func (o *ClientApplicationSchema) SetStarted(v ClientApplicationSchemaStarted)`

SetStarted sets Started field to given value.


### GetStrategies

`func (o *ClientApplicationSchema) GetStrategies() []string`

GetStrategies returns the Strategies field if non-nil, zero value otherwise.

### GetStrategiesOk

`func (o *ClientApplicationSchema) GetStrategiesOk() (*[]string, bool)`

GetStrategiesOk returns a tuple with the Strategies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategies

`func (o *ClientApplicationSchema) SetStrategies(v []string)`

SetStrategies sets Strategies field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


