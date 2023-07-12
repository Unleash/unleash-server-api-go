/*
 * Unleash API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 5.2.0-main
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// An addon provider. Defines a specific addon type and what the end user must configure when creating a new addon of that type.
type AddonTypeSchema struct {
	// The name of the addon type. When creating new addons, this goes in the payload's `type` field.
	Name string `json:"name"`
	// The addon type's name as it should be displayed in the admin UI.
	DisplayName string `json:"displayName"`
	// A URL to where you can find more information about using this addon type.
	DocumentationUrl string `json:"documentationUrl"`
	// A description of the addon type.
	Description string `json:"description"`
	// A list of [Unleash tag types](https://docs.getunleash.io/reference/tags#tag-types) that this addon uses. These tags will be added to the Unleash instance when an addon of this type is created.
	TagTypes []TagTypeSchema `json:"tagTypes,omitempty"`
	// The addon provider's parameters. Use these to configure an addon of this provider type. Items with `required: true` must be provided.
	Parameters []AddonParameterSchema `json:"parameters,omitempty"`
	// All the [event types](https://docs.getunleash.io/reference/api/legacy/unleash/admin/events#feature-toggle-events) that are available for this addon provider.
	Events []string `json:"events,omitempty"`
}
