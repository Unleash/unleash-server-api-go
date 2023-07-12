/*
 * Unleash API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 5.2.0-main
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Feature toggle configuration used by SDKs to evaluate state of a toggle
type ClientFeatureSchema struct {
	// The unique name of a feature toggle. Is validated to be URL safe on creation
	Name string `json:"name"`
	// What kind of feature flag is this. Refer to the documentation on [feature toggle types](https://docs.getunleash.io/reference/feature-toggle-types) for more information
	Type_ string `json:"type,omitempty"`
	// A description of the toggle
	Description string `json:"description,omitempty"`
	// Whether the feature flag is enabled for the current API key or not. This is ANDed with the evaluation results of the strategies list, so if this is false, the evaluation result will always be false
	Enabled bool `json:"enabled"`
	// If this is true Unleash believes this feature toggle has been active longer than Unleash expects a toggle of this type to be active
	Stale bool `json:"stale,omitempty"`
	// Set to true if SDKs should trigger [impression events](https://docs.getunleash.io/reference/impression-data) when this toggle is evaluated
	ImpressionData bool `json:"impressionData,omitempty"`
	// Which project this feature toggle belongs to
	Project string `json:"project,omitempty"`
	// Evaluation strategies for this toggle. Each entry in this list will be evaluated and ORed together
	Strategies []FeatureStrategySchema `json:"strategies,omitempty"`
	// [Variants](https://docs.getunleash.io/reference/feature-toggle-variants#what-are-variants) configured for this toggle
	Variants []VariantSchema `json:"variants,omitempty"`
}
