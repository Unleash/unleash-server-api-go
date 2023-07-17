# {{classname}}

All URIs are relative to *https://us.app.unleash-hosted.com/ushosted*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChangeMyPassword**](UsersApi.md#ChangeMyPassword) | **Post** /api/admin/user/change-password | Change your own password
[**ChangeUserPassword**](UsersApi.md#ChangeUserPassword) | **Post** /api/admin/user-admin/{id}/change-password | Change password for a user
[**CreateGroup**](UsersApi.md#CreateGroup) | **Post** /api/admin/groups | 
[**CreateRole**](UsersApi.md#CreateRole) | **Post** /api/admin/roles | 
[**CreateUser**](UsersApi.md#CreateUser) | **Post** /api/admin/user-admin | Create a new user
[**DeleteGroup**](UsersApi.md#DeleteGroup) | **Delete** /api/admin/groups/{groupId} | 
[**DeleteRole**](UsersApi.md#DeleteRole) | **Delete** /api/admin/roles/{roleId} | 
[**DeleteUser**](UsersApi.md#DeleteUser) | **Delete** /api/admin/user-admin/{id} | Delete a user
[**GetAdminCount**](UsersApi.md#GetAdminCount) | **Get** /api/admin/user-admin/admin-count | Get total count of admin accounts
[**GetBaseUsersAndGroups**](UsersApi.md#GetBaseUsersAndGroups) | **Get** /api/admin/user-admin/access | Get basic user and group information
[**GetGroup**](UsersApi.md#GetGroup) | **Get** /api/admin/groups/{groupId} | 
[**GetGroups**](UsersApi.md#GetGroups) | **Get** /api/admin/groups | 
[**GetMe**](UsersApi.md#GetMe) | **Get** /api/admin/user | Get your own user details
[**GetProfile**](UsersApi.md#GetProfile) | **Get** /api/admin/user/profile | Get your own user profile
[**GetRoleById**](UsersApi.md#GetRoleById) | **Get** /api/admin/roles/{roleId} | 
[**GetRoles**](UsersApi.md#GetRoles) | **Get** /api/admin/roles | 
[**GetUser**](UsersApi.md#GetUser) | **Get** /api/admin/user-admin/{id} | Get user
[**GetUsers**](UsersApi.md#GetUsers) | **Get** /api/admin/user-admin | Get all users and [root roles](https://docs.getunleash.io/reference/rbac#standard-roles)
[**ResetUserPassword**](UsersApi.md#ResetUserPassword) | **Post** /api/admin/user-admin/reset-password | Reset user password
[**SearchUsers**](UsersApi.md#SearchUsers) | **Get** /api/admin/user-admin/search | Search users
[**UpdateGroup**](UsersApi.md#UpdateGroup) | **Put** /api/admin/groups/{groupId} | 
[**UpdateRole**](UsersApi.md#UpdateRole) | **Put** /api/admin/roles/{roleId} | 
[**UpdateUser**](UsersApi.md#UpdateUser) | **Put** /api/admin/user-admin/{id} | Update a user
[**ValidateRole**](UsersApi.md#ValidateRole) | **Post** /api/admin/roles/validate | 
[**ValidateUserPassword**](UsersApi.md#ValidateUserPassword) | **Post** /api/admin/user-admin/validate-password | Validate password for a user

# **ChangeMyPassword**
> ChangeMyPassword(ctx, body)
Change your own password

Requires specifying old password and confirming new password

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PasswordSchema**](PasswordSchema.md)| passwordSchema | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChangeUserPassword**
> ChangeUserPassword(ctx, body, id)
Change password for a user

Change password for a user as an admin

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PasswordSchema**](PasswordSchema.md)| passwordSchema | 
  **id** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateGroup**
> ModelMap CreateGroup(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ModelMap**](map.md)| groupSchema | 

### Return type

[**ModelMap**](map.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateRole**
> RoleWithVersionSchema CreateRole(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateRoleWithPermissionsSchema**](CreateRoleWithPermissionsSchema.md)| createRoleWithPermissionsSchema | 

### Return type

[**RoleWithVersionSchema**](roleWithVersionSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateUser**
> CreateUserResponseSchema CreateUser(ctx, body)
Create a new user

Creates a new user with the given root role.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateUserSchema**](CreateUserSchema.md)| createUserSchema | 

### Return type

[**CreateUserResponseSchema**](createUserResponseSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteGroup**
> DeleteGroup(ctx, groupId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteRole**
> DeleteRole(ctx, roleId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **roleId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteUser**
> DeleteUser(ctx, id)
Delete a user

Deletes the user with the given userId

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAdminCount**
> AdminCountSchema GetAdminCount(ctx, )
Get total count of admin accounts

Get a total count of admins with password, without password and admin service accounts

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**AdminCountSchema**](adminCountSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBaseUsersAndGroups**
> UsersGroupsBaseSchema GetBaseUsersAndGroups(ctx, )
Get basic user and group information

Get a subset of user and group information eligible even for non-admin users

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**UsersGroupsBaseSchema**](usersGroupsBaseSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGroup**
> ModelMap GetGroup(ctx, groupId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupId** | **string**|  | 

### Return type

[**ModelMap**](map.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGroups**
> GroupsSchema GetGroups(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GroupsSchema**](groupsSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMe**
> MeSchema GetMe(ctx, )
Get your own user details

Detailed information about the current user, user permissions and user feedback

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**MeSchema**](meSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProfile**
> ProfileSchema GetProfile(ctx, )
Get your own user profile

Detailed information about the current user root role and project membership

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ProfileSchema**](profileSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRoleById**
> RoleWithPermissionsSchema GetRoleById(ctx, roleId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **roleId** | **string**|  | 

### Return type

[**RoleWithPermissionsSchema**](roleWithPermissionsSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRoles**
> RolesWithVersionSchema GetRoles(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

[**RolesWithVersionSchema**](rolesWithVersionSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUser**
> UserSchema GetUser(ctx, id)
Get user

Will return a single user by id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**UserSchema**](userSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUsers**
> UsersSchema GetUsers(ctx, )
Get all users and [root roles](https://docs.getunleash.io/reference/rbac#standard-roles)

Will return all users and all available root roles for the Unleash instance.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**UsersSchema**](usersSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ResetUserPassword**
> ResetPasswordSchema ResetUserPassword(ctx, body)
Reset user password

Reset user password as an admin

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**IdSchema**](IdSchema.md)| idSchema | 

### Return type

[**ResetPasswordSchema**](resetPasswordSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchUsers**
> UsersSchema SearchUsers(ctx, optional)
Search users

 It will preform a simple search based on name and email matching the given query. Requires minimum 2 characters

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UsersApiSearchUsersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiSearchUsersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **optional.String**| The pattern to search in the username or email | 

### Return type

[**UsersSchema**](usersSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateGroup**
> ModelMap UpdateGroup(ctx, body, groupId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ModelMap**](map.md)| groupSchema | 
  **groupId** | **string**|  | 

### Return type

[**ModelMap**](map.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRole**
> RoleWithVersionSchema UpdateRole(ctx, body, roleId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateRoleWithPermissionsSchema**](CreateRoleWithPermissionsSchema.md)| createRoleWithPermissionsSchema | 
  **roleId** | **string**|  | 

### Return type

[**RoleWithVersionSchema**](roleWithVersionSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateUser**
> CreateUserResponseSchema UpdateUser(ctx, body, id)
Update a user

Only the explicitly specified fields get updated.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ModelMap**](map.md)| updateUserSchema | 
  **id** | **string**|  | 

### Return type

[**CreateUserResponseSchema**](createUserResponseSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ValidateRole**
> ValidateRole(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateRoleWithPermissionsSchema**](CreateRoleWithPermissionsSchema.md)| createRoleWithPermissionsSchema | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ValidateUserPassword**
> ValidateUserPassword(ctx, body)
Validate password for a user

Validate the password strength. Minimum 10 characters, uppercase letter, number, special character.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PasswordSchema**](PasswordSchema.md)| passwordSchema | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

