/*
 * Unleash API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 5.2.0-main
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger
import (
	"time"
)

// Data relating to the previous state of the event's subject.
type EventSchemaPreData struct {
	// Name of the feature toggle/strategy/environment that this event relates to
	Name string `json:"name,omitempty"`
	// The description of the object this event relates to
	Description string `json:"description,omitempty"`
	// If this event relates to a feature toggle, the type of feature toggle.
	Type_ string `json:"type,omitempty"`
	// The project this event relates to
	Project string `json:"project,omitempty"`
	// Is the feature toggle this event relates to stale
	Stale bool `json:"stale,omitempty"`
	// Variants configured for this toggle
	Variants []VariantSchema `json:"variants,omitempty"`
	// The time the event happened as a RFC 3339-conformant timestamp.
	CreatedAt time.Time `json:"createdAt,omitempty"`
	// The time the feature was last seen
	LastSeenAt time.Time `json:"lastSeenAt,omitempty"`
	// Should [impression events](https://docs.getunleash.io/reference/impression-data) activate for this feature toggle
	ImpressionData bool `json:"impressionData,omitempty"`
}
