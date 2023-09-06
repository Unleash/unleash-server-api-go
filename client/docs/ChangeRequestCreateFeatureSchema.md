# ChangeRequestCreateFeatureSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Feature** | **string** | The name of the feature that this change applies to. | 
**Action** | **string** | The name of this action. | 
**Payload** | [**[]SetStrategySortOrderSchemaInner**](SetStrategySortOrderSchemaInner.md) | An array of strategies with their new sort order | 

## Methods

### NewChangeRequestCreateFeatureSchema

`func NewChangeRequestCreateFeatureSchema(feature string, action string, payload []SetStrategySortOrderSchemaInner, ) *ChangeRequestCreateFeatureSchema`

NewChangeRequestCreateFeatureSchema instantiates a new ChangeRequestCreateFeatureSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChangeRequestCreateFeatureSchemaWithDefaults

`func NewChangeRequestCreateFeatureSchemaWithDefaults() *ChangeRequestCreateFeatureSchema`

NewChangeRequestCreateFeatureSchemaWithDefaults instantiates a new ChangeRequestCreateFeatureSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeature

`func (o *ChangeRequestCreateFeatureSchema) GetFeature() string`

GetFeature returns the Feature field if non-nil, zero value otherwise.

### GetFeatureOk

`func (o *ChangeRequestCreateFeatureSchema) GetFeatureOk() (*string, bool)`

GetFeatureOk returns a tuple with the Feature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeature

`func (o *ChangeRequestCreateFeatureSchema) SetFeature(v string)`

SetFeature sets Feature field to given value.


### GetAction

`func (o *ChangeRequestCreateFeatureSchema) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ChangeRequestCreateFeatureSchema) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ChangeRequestCreateFeatureSchema) SetAction(v string)`

SetAction sets Action field to given value.


### GetPayload

`func (o *ChangeRequestCreateFeatureSchema) GetPayload() []SetStrategySortOrderSchemaInner`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *ChangeRequestCreateFeatureSchema) GetPayloadOk() (*[]SetStrategySortOrderSchemaInner, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *ChangeRequestCreateFeatureSchema) SetPayload(v []SetStrategySortOrderSchemaInner)`

SetPayload sets Payload field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


