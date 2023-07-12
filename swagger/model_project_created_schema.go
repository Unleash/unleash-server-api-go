/*
 * Unleash API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 5.2.0-main
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Details about the newly created project.
type ProjectCreatedSchema struct {
	// The project's identifier.
	Id string `json:"id"`
	// The project's name.
	Name string `json:"name"`
	// The project's description.
	Description string `json:"description,omitempty"`
	// A mode of the project affecting what actions are possible in this project
	Mode string `json:"mode,omitempty"`
	// A default stickiness for the project affecting the default stickiness value for variants and Gradual Rollout strategy
	DefaultStickiness string `json:"defaultStickiness,omitempty"`
}
