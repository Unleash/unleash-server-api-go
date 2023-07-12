/*
 * Unleash API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 5.2.0-main
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// The feature variant you receive based on the provided context or the _disabled                           variant_. If a feature is disabled or doesn't have any                           variants, you would get the _disabled variant_.                           Otherwise, you'll get one of thefeature's defined variants.
type AdvancedPlaygroundEnvironmentFeatureSchemaVariant struct {
	// The variant's name. If there is no variant or if the toggle is disabled, this will be `disabled`
	Name string `json:"name"`
	// Whether the variant is enabled or not. If the feature is disabled or if it doesn't have variants, this property will be `false`
	Enabled bool                                                      `json:"enabled"`
	Payload *AdvancedPlaygroundEnvironmentFeatureSchemaVariantPayload `json:"payload,omitempty"`
}
