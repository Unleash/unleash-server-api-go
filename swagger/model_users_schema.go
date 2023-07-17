/*
 * Unleash API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 5.2.0-main
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Users and root roles
type UsersSchema struct {
	// A list of users in the Unleash instance.
	Users []UserSchema `json:"users"`
	// A list of [root roles](https://docs.getunleash.io/reference/rbac#standard-roles) in the Unleash instance.
	RootRoles []RoleSchema `json:"rootRoles,omitempty"`
}
