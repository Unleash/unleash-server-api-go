/*
 * Unleash API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 5.2.0-main
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// The payload must contain at least one of the name and email properties, though which one is up to you. For the user to be able to log in to the system, the user must have an email.
type CreateUserSchema struct {
	// The user's username. Must be provided if email is not provided.
	Username string `json:"username,omitempty"`
	// The user's email address. Must be provided if username is not provided.
	Email string `json:"email,omitempty"`
	// The user's name (not the user's username).
	Name string `json:"name,omitempty"`
	// Password for the user
	Password string `json:"password,omitempty"`
	// The role to assign to the user. Can be either the role's ID or its unique name.
	RootRole *OneOfcreateUserSchemaRootRole `json:"rootRole"`
	// Whether to send a welcome email with a login link to the user or not. Defaults to `true`.
	SendEmail bool `json:"sendEmail,omitempty"`
}