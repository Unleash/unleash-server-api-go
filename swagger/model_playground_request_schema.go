/*
 * Unleash API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 5.2.0-main
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Data for the playground API to evaluate toggles
type PlaygroundRequestSchema struct {
	// The environment to evaluate toggles in.
	Environment string `json:"environment"`
	// A list of projects to check for toggles in.
	Projects *OneOfplaygroundRequestSchemaProjects `json:"projects,omitempty"`
	Context  *interface{}                          `json:"context"`
}
