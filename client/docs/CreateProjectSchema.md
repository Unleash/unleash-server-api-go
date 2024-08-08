# CreateProjectSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The project&#39;s identifier. If this property is not present or is an empty string, Unleash will generate the project id automatically. This property is deprecated. | [optional] 
**Name** | **string** | The project&#39;s name. The name must contain at least one non-whitespace character. | 
**Description** | Pointer to **NullableString** | The project&#39;s description. | [optional] 
**Mode** | Pointer to **string** | A mode of the project affecting what actions are possible in this project | [optional] [default to "open"]
**DefaultStickiness** | Pointer to **string** | A default stickiness for the project affecting the default stickiness value for variants and Gradual Rollout strategy | [optional] [default to "default"]
**Environments** | Pointer to **[]string** | A list of environments that should be enabled for this project. When provided, the list must contain at least one environment. If this property is missing, Unleash will default to enabling all non-deprecated environments for the project. | [optional] 
**ChangeRequestEnvironments** | Pointer to [**[]CreateProjectSchemaChangeRequestEnvironmentsInner**](CreateProjectSchemaChangeRequestEnvironmentsInner.md) | A list of environments that should have change requests enabled. If the list includes environments not in the &#x60;environments&#x60; list, they will still have change requests enabled. | [optional] 

## Methods

### NewCreateProjectSchema

`func NewCreateProjectSchema(name string, ) *CreateProjectSchema`

NewCreateProjectSchema instantiates a new CreateProjectSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateProjectSchemaWithDefaults

`func NewCreateProjectSchemaWithDefaults() *CreateProjectSchema`

NewCreateProjectSchemaWithDefaults instantiates a new CreateProjectSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateProjectSchema) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateProjectSchema) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateProjectSchema) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CreateProjectSchema) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CreateProjectSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateProjectSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateProjectSchema) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateProjectSchema) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateProjectSchema) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateProjectSchema) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateProjectSchema) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreateProjectSchema) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreateProjectSchema) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetMode

`func (o *CreateProjectSchema) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *CreateProjectSchema) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *CreateProjectSchema) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *CreateProjectSchema) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetDefaultStickiness

`func (o *CreateProjectSchema) GetDefaultStickiness() string`

GetDefaultStickiness returns the DefaultStickiness field if non-nil, zero value otherwise.

### GetDefaultStickinessOk

`func (o *CreateProjectSchema) GetDefaultStickinessOk() (*string, bool)`

GetDefaultStickinessOk returns a tuple with the DefaultStickiness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultStickiness

`func (o *CreateProjectSchema) SetDefaultStickiness(v string)`

SetDefaultStickiness sets DefaultStickiness field to given value.

### HasDefaultStickiness

`func (o *CreateProjectSchema) HasDefaultStickiness() bool`

HasDefaultStickiness returns a boolean if a field has been set.

### GetEnvironments

`func (o *CreateProjectSchema) GetEnvironments() []string`

GetEnvironments returns the Environments field if non-nil, zero value otherwise.

### GetEnvironmentsOk

`func (o *CreateProjectSchema) GetEnvironmentsOk() (*[]string, bool)`

GetEnvironmentsOk returns a tuple with the Environments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironments

`func (o *CreateProjectSchema) SetEnvironments(v []string)`

SetEnvironments sets Environments field to given value.

### HasEnvironments

`func (o *CreateProjectSchema) HasEnvironments() bool`

HasEnvironments returns a boolean if a field has been set.

### GetChangeRequestEnvironments

`func (o *CreateProjectSchema) GetChangeRequestEnvironments() []CreateProjectSchemaChangeRequestEnvironmentsInner`

GetChangeRequestEnvironments returns the ChangeRequestEnvironments field if non-nil, zero value otherwise.

### GetChangeRequestEnvironmentsOk

`func (o *CreateProjectSchema) GetChangeRequestEnvironmentsOk() (*[]CreateProjectSchemaChangeRequestEnvironmentsInner, bool)`

GetChangeRequestEnvironmentsOk returns a tuple with the ChangeRequestEnvironments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeRequestEnvironments

`func (o *CreateProjectSchema) SetChangeRequestEnvironments(v []CreateProjectSchemaChangeRequestEnvironmentsInner)`

SetChangeRequestEnvironments sets ChangeRequestEnvironments field to given value.

### HasChangeRequestEnvironments

`func (o *CreateProjectSchema) HasChangeRequestEnvironments() bool`

HasChangeRequestEnvironments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


