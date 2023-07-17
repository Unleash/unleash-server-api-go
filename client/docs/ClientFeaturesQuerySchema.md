# ClientFeaturesQuerySchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tag** | Pointer to **[][]string** | Features tagged with one of these tags are included | [optional] 
**Project** | Pointer to **[]string** | Features that are part of these projects are included in this response. (DEPRECATED) - Handled by API tokens | [optional] 
**NamePrefix** | Pointer to **string** | Features are filtered to only include features whose name starts with this prefix | [optional] 
**Environment** | Pointer to **string** | Strategies for the feature toggle configured for this environment are included. (DEPRECATED) - Handled by API tokens | [optional] 
**InlineSegmentConstraints** | Pointer to **bool** | Set to true if requesting client does not support Unleash-Client-Specification 4.2.2 or newer. Modern SDKs will have this set to false, since they will be able to merge constraints and segments themselves | [optional] 

## Methods

### NewClientFeaturesQuerySchema

`func NewClientFeaturesQuerySchema() *ClientFeaturesQuerySchema`

NewClientFeaturesQuerySchema instantiates a new ClientFeaturesQuerySchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientFeaturesQuerySchemaWithDefaults

`func NewClientFeaturesQuerySchemaWithDefaults() *ClientFeaturesQuerySchema`

NewClientFeaturesQuerySchemaWithDefaults instantiates a new ClientFeaturesQuerySchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTag

`func (o *ClientFeaturesQuerySchema) GetTag() [][]string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *ClientFeaturesQuerySchema) GetTagOk() (*[][]string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *ClientFeaturesQuerySchema) SetTag(v [][]string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *ClientFeaturesQuerySchema) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetProject

`func (o *ClientFeaturesQuerySchema) GetProject() []string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ClientFeaturesQuerySchema) GetProjectOk() (*[]string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ClientFeaturesQuerySchema) SetProject(v []string)`

SetProject sets Project field to given value.

### HasProject

`func (o *ClientFeaturesQuerySchema) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetNamePrefix

`func (o *ClientFeaturesQuerySchema) GetNamePrefix() string`

GetNamePrefix returns the NamePrefix field if non-nil, zero value otherwise.

### GetNamePrefixOk

`func (o *ClientFeaturesQuerySchema) GetNamePrefixOk() (*string, bool)`

GetNamePrefixOk returns a tuple with the NamePrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamePrefix

`func (o *ClientFeaturesQuerySchema) SetNamePrefix(v string)`

SetNamePrefix sets NamePrefix field to given value.

### HasNamePrefix

`func (o *ClientFeaturesQuerySchema) HasNamePrefix() bool`

HasNamePrefix returns a boolean if a field has been set.

### GetEnvironment

`func (o *ClientFeaturesQuerySchema) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *ClientFeaturesQuerySchema) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *ClientFeaturesQuerySchema) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *ClientFeaturesQuerySchema) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetInlineSegmentConstraints

`func (o *ClientFeaturesQuerySchema) GetInlineSegmentConstraints() bool`

GetInlineSegmentConstraints returns the InlineSegmentConstraints field if non-nil, zero value otherwise.

### GetInlineSegmentConstraintsOk

`func (o *ClientFeaturesQuerySchema) GetInlineSegmentConstraintsOk() (*bool, bool)`

GetInlineSegmentConstraintsOk returns a tuple with the InlineSegmentConstraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInlineSegmentConstraints

`func (o *ClientFeaturesQuerySchema) SetInlineSegmentConstraints(v bool)`

SetInlineSegmentConstraints sets InlineSegmentConstraints field to given value.

### HasInlineSegmentConstraints

`func (o *ClientFeaturesQuerySchema) HasInlineSegmentConstraints() bool`

HasInlineSegmentConstraints returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


