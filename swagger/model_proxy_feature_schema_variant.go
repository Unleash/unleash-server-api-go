/*
 * Unleash API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 5.2.0-main
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Variant details
type ProxyFeatureSchemaVariant struct {
	// The variants name. Is unique for this feature toggle
	Name string `json:"name"`
	// Whether the variant is enabled or not.
	Enabled bool `json:"enabled"`
	Payload *ProxyFeatureSchemaVariantPayload `json:"payload,omitempty"`
}
