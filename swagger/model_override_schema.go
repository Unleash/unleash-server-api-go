/*
 * Unleash API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 5.2.0-main
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// An override for deciding which variant should be assigned to a user based on the context name
type OverrideSchema struct {
	// The name of the context field used to determine overrides
	ContextName string `json:"contextName"`
	// Which values that should be overriden
	Values []string `json:"values"`
}
