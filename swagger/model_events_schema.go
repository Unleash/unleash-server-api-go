/*
 * Unleash API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 5.2.0-main
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// A list of events that has happened in the system
type EventsSchema struct {
	// The api version of this response. A natural increasing number. Only increases if format changes
	Version int32 `json:"version"`
	// The list of events
	Events []EventSchema `json:"events"`
	// The total count of events
	TotalEvents int32 `json:"totalEvents,omitempty"`
}
