# AdminPermissionsSchemaPermissions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Root** | Pointer to [**[]AdminPermissionSchema**](AdminPermissionSchema.md) | Permissions available at the root level, i.e. not connected to any specific project or environment | [optional] 
**Project** | [**[]AdminPermissionSchema**](AdminPermissionSchema.md) | Permissions available at the project level | 
**Environments** | [**[]AdminPermissionsSchemaPermissionsEnvironmentsInner**](AdminPermissionsSchemaPermissionsEnvironmentsInner.md) | A list of environments with available permissions per environment | 

## Methods

### NewAdminPermissionsSchemaPermissions

`func NewAdminPermissionsSchemaPermissions(project []AdminPermissionSchema, environments []AdminPermissionsSchemaPermissionsEnvironmentsInner, ) *AdminPermissionsSchemaPermissions`

NewAdminPermissionsSchemaPermissions instantiates a new AdminPermissionsSchemaPermissions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminPermissionsSchemaPermissionsWithDefaults

`func NewAdminPermissionsSchemaPermissionsWithDefaults() *AdminPermissionsSchemaPermissions`

NewAdminPermissionsSchemaPermissionsWithDefaults instantiates a new AdminPermissionsSchemaPermissions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoot

`func (o *AdminPermissionsSchemaPermissions) GetRoot() []AdminPermissionSchema`

GetRoot returns the Root field if non-nil, zero value otherwise.

### GetRootOk

`func (o *AdminPermissionsSchemaPermissions) GetRootOk() (*[]AdminPermissionSchema, bool)`

GetRootOk returns a tuple with the Root field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoot

`func (o *AdminPermissionsSchemaPermissions) SetRoot(v []AdminPermissionSchema)`

SetRoot sets Root field to given value.

### HasRoot

`func (o *AdminPermissionsSchemaPermissions) HasRoot() bool`

HasRoot returns a boolean if a field has been set.

### GetProject

`func (o *AdminPermissionsSchemaPermissions) GetProject() []AdminPermissionSchema`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *AdminPermissionsSchemaPermissions) GetProjectOk() (*[]AdminPermissionSchema, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *AdminPermissionsSchemaPermissions) SetProject(v []AdminPermissionSchema)`

SetProject sets Project field to given value.


### GetEnvironments

`func (o *AdminPermissionsSchemaPermissions) GetEnvironments() []AdminPermissionsSchemaPermissionsEnvironmentsInner`

GetEnvironments returns the Environments field if non-nil, zero value otherwise.

### GetEnvironmentsOk

`func (o *AdminPermissionsSchemaPermissions) GetEnvironmentsOk() (*[]AdminPermissionsSchemaPermissionsEnvironmentsInner, bool)`

GetEnvironmentsOk returns a tuple with the Environments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironments

`func (o *AdminPermissionsSchemaPermissions) SetEnvironments(v []AdminPermissionsSchemaPermissionsEnvironmentsInner)`

SetEnvironments sets Environments field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


