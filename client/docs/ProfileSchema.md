# ProfileSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RootRole** | [**RoleSchema**](RoleSchema.md) |  | 
**Projects** | **[]string** |  | 
**Features** | [**[]FeatureSchema**](FeatureSchema.md) |  | 

## Methods

### NewProfileSchema

`func NewProfileSchema(rootRole RoleSchema, projects []string, features []FeatureSchema, ) *ProfileSchema`

NewProfileSchema instantiates a new ProfileSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProfileSchemaWithDefaults

`func NewProfileSchemaWithDefaults() *ProfileSchema`

NewProfileSchemaWithDefaults instantiates a new ProfileSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRootRole

`func (o *ProfileSchema) GetRootRole() RoleSchema`

GetRootRole returns the RootRole field if non-nil, zero value otherwise.

### GetRootRoleOk

`func (o *ProfileSchema) GetRootRoleOk() (*RoleSchema, bool)`

GetRootRoleOk returns a tuple with the RootRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootRole

`func (o *ProfileSchema) SetRootRole(v RoleSchema)`

SetRootRole sets RootRole field to given value.


### GetProjects

`func (o *ProfileSchema) GetProjects() []string`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *ProfileSchema) GetProjectsOk() (*[]string, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *ProfileSchema) SetProjects(v []string)`

SetProjects sets Projects field to given value.


### GetFeatures

`func (o *ProfileSchema) GetFeatures() []FeatureSchema`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *ProfileSchema) GetFeaturesOk() (*[]FeatureSchema, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *ProfileSchema) SetFeatures(v []FeatureSchema)`

SetFeatures sets Features field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


