/*
 * Unleash API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 5.2.0-main
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Configuration data for server-side SDKs for evaluating feature flags.
type ClientFeaturesSchema struct {
	// A version number for the format used in the response. Most Unleash instances now return version 2, which includes segments as a separate array
	Version float64 `json:"version"`
	// A list of feature toggles with their configuration
	Features []ClientFeatureSchema `json:"features"`
	// A list of [Segments](https://docs.getunleash.io/reference/segments) configured for this Unleash instance
	Segments []SegmentSchema `json:"segments,omitempty"`
	Query *ClientFeaturesQuerySchema `json:"query,omitempty"`
}
