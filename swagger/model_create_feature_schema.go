/*
 * Unleash API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 5.2.0-main
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type CreateFeatureSchema struct {
	Name           string `json:"name"`
	Type_          string `json:"type,omitempty"`
	Description    string `json:"description,omitempty"`
	ImpressionData bool   `json:"impressionData,omitempty"`
}