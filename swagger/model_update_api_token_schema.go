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

// An object with fields to updated for a given API token.
type UpdateApiTokenSchema struct {
	// The new time when this token should expire.
	ExpiresAt time.Time `json:"expiresAt"`
}
