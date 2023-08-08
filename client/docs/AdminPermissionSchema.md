# AdminPermissionSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The identifier for this permission | 
**Name** | **string** | The name of this permission | 
**DisplayName** | **string** | The name to display in listings of permissions | 
**Type** | **string** | What level this permission applies to. Either root, project or the name of the environment it applies to | 
**Environment** | Pointer to **string** | Which environment this permission applies to | [optional] 

## Methods

### NewAdminPermissionSchema

`func NewAdminPermissionSchema(id int32, name string, displayName string, type_ string, ) *AdminPermissionSchema`

NewAdminPermissionSchema instantiates a new AdminPermissionSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminPermissionSchemaWithDefaults

`func NewAdminPermissionSchemaWithDefaults() *AdminPermissionSchema`

NewAdminPermissionSchemaWithDefaults instantiates a new AdminPermissionSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AdminPermissionSchema) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AdminPermissionSchema) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AdminPermissionSchema) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *AdminPermissionSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AdminPermissionSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AdminPermissionSchema) SetName(v string)`

SetName sets Name field to given value.


### GetDisplayName

`func (o *AdminPermissionSchema) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AdminPermissionSchema) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AdminPermissionSchema) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetType

`func (o *AdminPermissionSchema) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AdminPermissionSchema) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AdminPermissionSchema) SetType(v string)`

SetType sets Type field to given value.


### GetEnvironment

`func (o *AdminPermissionSchema) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *AdminPermissionSchema) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *AdminPermissionSchema) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *AdminPermissionSchema) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


