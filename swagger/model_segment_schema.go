/*
 * Unleash API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 5.2.0-main
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Represents a segment of users defined by a set of constraints.
type SegmentSchema struct {
	// The segment's id.
	Id float64 `json:"id"`
	// The name of the segment.
	Name string `json:"name,omitempty"`
	// The description of the segment.
	Description string `json:"description,omitempty"`
	// List of constraints that determine which users are part of the segment
	Constraints []ConstraintSchema `json:"constraints"`
}
