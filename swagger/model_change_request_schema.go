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

type ChangeRequestSchema struct {
	Id           float64                            `json:"id"`
	Title        string                             `json:"title,omitempty"`
	Environment  string                             `json:"environment"`
	State        string                             `json:"state"`
	MinApprovals float64                            `json:"minApprovals"`
	Project      string                             `json:"project"`
	Features     []ChangeRequestFeatureSchema       `json:"features"`
	Approvals    []ChangeRequestApprovalSchema      `json:"approvals,omitempty"`
	Comments     []ChangeRequestCommentSchema       `json:"comments,omitempty"`
	CreatedBy    *ChangeRequestEventSchemaCreatedBy `json:"createdBy"`
	CreatedAt    time.Time                          `json:"createdAt"`
}
