/*
 * Unleash API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 5.2.0-main
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// A wrapper object containing all strategies that use a specific context field
type ContextFieldStrategiesSchema struct {
	// List of strategies using the context field
	Strategies []ContextFieldStrategiesSchemaStrategies `json:"strategies"`
}
