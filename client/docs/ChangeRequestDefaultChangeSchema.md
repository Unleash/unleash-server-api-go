# ChangeRequestDefaultChangeSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** | The kind of action this is. | 
**Payload** | **map[string]interface{}** | The necessary data to perform this change. | 

## Methods

### NewChangeRequestDefaultChangeSchema

`func NewChangeRequestDefaultChangeSchema(action string, payload map[string]interface{}, ) *ChangeRequestDefaultChangeSchema`

NewChangeRequestDefaultChangeSchema instantiates a new ChangeRequestDefaultChangeSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChangeRequestDefaultChangeSchemaWithDefaults

`func NewChangeRequestDefaultChangeSchemaWithDefaults() *ChangeRequestDefaultChangeSchema`

NewChangeRequestDefaultChangeSchemaWithDefaults instantiates a new ChangeRequestDefaultChangeSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *ChangeRequestDefaultChangeSchema) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ChangeRequestDefaultChangeSchema) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ChangeRequestDefaultChangeSchema) SetAction(v string)`

SetAction sets Action field to given value.


### GetPayload

`func (o *ChangeRequestDefaultChangeSchema) GetPayload() map[string]interface{}`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *ChangeRequestDefaultChangeSchema) GetPayloadOk() (*map[string]interface{}, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *ChangeRequestDefaultChangeSchema) SetPayload(v map[string]interface{})`

SetPayload sets Payload field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


