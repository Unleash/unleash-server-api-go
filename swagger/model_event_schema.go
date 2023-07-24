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

// An event describing something happening in the system
type EventSchema struct {
	// The ID of the event. An increasing natural number.
	Id int32 `json:"id"`
	// The time the event happened as a RFC 3339-conformant timestamp.
	CreatedAt time.Time `json:"createdAt"`
	// What [type](https://docs.getunleash.io/reference/api/legacy/unleash/admin/events#event-type-description) of event this is
	Type_ string `json:"type"`
	// Which user created this event
	CreatedBy string `json:"createdBy"`
	// The feature toggle environment the event relates to, if applicable.
	Environment string `json:"environment,omitempty"`
	// The project the event relates to, if applicable.
	Project string `json:"project,omitempty"`
	// The name of the feature toggle the event relates to, if applicable.
	FeatureName string              `json:"featureName,omitempty"`
	Data        *EventSchemaData    `json:"data,omitempty"`
	PreData     *EventSchemaPreData `json:"preData,omitempty"`
	// Any tags related to the event, if applicable.
	Tags []TagSchema `json:"tags,omitempty"`
}