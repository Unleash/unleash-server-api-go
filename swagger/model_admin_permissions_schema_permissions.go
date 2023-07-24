/*
 * Unleash API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 5.2.0-main
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Returns permissions available at all three levels (root|project|environment)
type AdminPermissionsSchemaPermissions struct {
	// Permissions available at the root level, i.e. not connected to any specific project or environment
	Root []AdminPermissionSchema `json:"root,omitempty"`
	// Permissions available at the project level
	Project []AdminPermissionSchema `json:"project"`
	// A list of environments with available permissions per environment
	Environments []AdminPermissionsSchemaPermissionsEnvironments `json:"environments"`
}