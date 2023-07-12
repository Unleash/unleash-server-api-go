/*
 * Unleash API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 5.2.0-main
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type InlineResponse400 struct {
	// The ID of the error instance
	Id string `json:"id,omitempty"`
	// The name of the error kind
	Name string `json:"name,omitempty"`
	// A description of what went wrong.
	Message string `json:"message,omitempty"`
}
