/*
 * Unleash API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 5.2.0-main
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// The result of the export operation, providing you with the feature toggle definitions, strategy definitions and the rest of the elements relevant to the features (tags, environments etc.)
type ExportResultSchema struct {
	// All the exported features.
	Features []FeatureSchema `json:"features"`
	// All strategy instances that are used by the exported features in the `features` list.
	FeatureStrategies []FeatureStrategySchema `json:"featureStrategies"`
	// Environment-specific configuration for all the features in the `features` list. Includes data such as whether the feature is enabled in the selected export environment, whether there are any variants assigned, etc.
	FeatureEnvironments []FeatureEnvironmentSchema `json:"featureEnvironments,omitempty"`
	// A list of all the context fields that are in use by any of the strategies in the `featureStrategies` list.
	ContextFields []ContextFieldSchema `json:"contextFields,omitempty"`
	// A list of all the tags that have been applied to any of the features in the `features` list.
	FeatureTags []FeatureTagSchema `json:"featureTags,omitempty"`
	// A list of all the segments that are used by the strategies in the `featureStrategies` list.
	Segments []ExportResultSchemaSegments `json:"segments,omitempty"`
	// A list of all of the tag types that are used in the `featureTags` list.
	TagTypes []TagTypeSchema `json:"tagTypes"`
}
