/*
 * Unleash API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 5.2.0-main
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// A simplified feature toggle model intended for the Unleash playground.
type PlaygroundFeatureSchema struct {
	// The feature's name.
	Name string `json:"name"`
	// The ID of the project that contains this feature.
	ProjectId  string                             `json:"projectId"`
	Strategies *PlaygroundFeatureSchemaStrategies `json:"strategies"`
	// Whether the feature is active and would be evaluated in the provided environment in a normal SDK context.
	IsEnabledInCurrentEnvironment bool `json:"isEnabledInCurrentEnvironment"`
	// Whether this feature is enabled or not in the current environment.                           If a feature can't be fully evaluated (that is, `strategies.result` is `unknown`),                           this will be `false` to align with how client SDKs treat unresolved feature states.
	IsEnabled bool                                               `json:"isEnabled"`
	Variant   *AdvancedPlaygroundEnvironmentFeatureSchemaVariant `json:"variant"`
	// The feature variants.
	Variants []VariantSchema `json:"variants"`
}