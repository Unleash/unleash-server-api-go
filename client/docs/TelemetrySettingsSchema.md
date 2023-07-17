# TelemetrySettingsSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VersionInfoCollectionEnabled** | **bool** | Whether collection of version info is enabled/active. | 
**FeatureInfoCollectionEnabled** | **bool** | Whether collection of feature usage metrics is enabled/active. | 

## Methods

### NewTelemetrySettingsSchema

`func NewTelemetrySettingsSchema(versionInfoCollectionEnabled bool, featureInfoCollectionEnabled bool, ) *TelemetrySettingsSchema`

NewTelemetrySettingsSchema instantiates a new TelemetrySettingsSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelemetrySettingsSchemaWithDefaults

`func NewTelemetrySettingsSchemaWithDefaults() *TelemetrySettingsSchema`

NewTelemetrySettingsSchemaWithDefaults instantiates a new TelemetrySettingsSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersionInfoCollectionEnabled

`func (o *TelemetrySettingsSchema) GetVersionInfoCollectionEnabled() bool`

GetVersionInfoCollectionEnabled returns the VersionInfoCollectionEnabled field if non-nil, zero value otherwise.

### GetVersionInfoCollectionEnabledOk

`func (o *TelemetrySettingsSchema) GetVersionInfoCollectionEnabledOk() (*bool, bool)`

GetVersionInfoCollectionEnabledOk returns a tuple with the VersionInfoCollectionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionInfoCollectionEnabled

`func (o *TelemetrySettingsSchema) SetVersionInfoCollectionEnabled(v bool)`

SetVersionInfoCollectionEnabled sets VersionInfoCollectionEnabled field to given value.


### GetFeatureInfoCollectionEnabled

`func (o *TelemetrySettingsSchema) GetFeatureInfoCollectionEnabled() bool`

GetFeatureInfoCollectionEnabled returns the FeatureInfoCollectionEnabled field if non-nil, zero value otherwise.

### GetFeatureInfoCollectionEnabledOk

`func (o *TelemetrySettingsSchema) GetFeatureInfoCollectionEnabledOk() (*bool, bool)`

GetFeatureInfoCollectionEnabledOk returns a tuple with the FeatureInfoCollectionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureInfoCollectionEnabled

`func (o *TelemetrySettingsSchema) SetFeatureInfoCollectionEnabled(v bool)`

SetFeatureInfoCollectionEnabled sets FeatureInfoCollectionEnabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


