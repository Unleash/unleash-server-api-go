/*
 * Unleash API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 5.2.0-main
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Frontend API feature
type ProxyFeatureSchema struct {
	// Unique feature name.
	Name string `json:"name"`
	// Always set to `true`.
	Enabled bool `json:"enabled"`
	// `true` if the impression data collection is enabled for the feature, otherwise `false`.
	ImpressionData bool `json:"impressionData"`
	Variant *ProxyFeatureSchemaVariant `json:"variant,omitempty"`
}
