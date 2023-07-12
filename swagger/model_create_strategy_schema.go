/*
 * Unleash API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 5.2.0-main
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// The data required to create a strategy type. Refer to the docs on [custom strategy types](https://docs.getunleash.io/reference/custom-activation-strategies) for more information.
type CreateStrategySchema struct {
	// The name of the strategy type. Must be unique.
	Name string `json:"name"`
	// A description of the strategy type.
	Description string `json:"description,omitempty"`
	// The parameter list lets you pass arguments to your custom activation strategy. These will be made available to your custom strategy implementation.
	Parameters []CreateStrategySchemaParameters `json:"parameters"`
}
