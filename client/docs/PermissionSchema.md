# PermissionSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Permission** | **string** | [Project](https://docs.getunleash.io/reference/rbac#project-permissions) or [environment](https://docs.getunleash.io/reference/rbac#environment-permissions) permission name | 
**Project** | Pointer to **string** | The project this permission applies to | [optional] 
**Environment** | Pointer to **string** | The environment this permission applies to | [optional] 

## Methods

### NewPermissionSchema

`func NewPermissionSchema(permission string, ) *PermissionSchema`

NewPermissionSchema instantiates a new PermissionSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionSchemaWithDefaults

`func NewPermissionSchemaWithDefaults() *PermissionSchema`

NewPermissionSchemaWithDefaults instantiates a new PermissionSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPermission

`func (o *PermissionSchema) GetPermission() string`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *PermissionSchema) GetPermissionOk() (*string, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *PermissionSchema) SetPermission(v string)`

SetPermission sets Permission field to given value.


### GetProject

`func (o *PermissionSchema) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *PermissionSchema) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *PermissionSchema) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *PermissionSchema) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetEnvironment

`func (o *PermissionSchema) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *PermissionSchema) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *PermissionSchema) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *PermissionSchema) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


