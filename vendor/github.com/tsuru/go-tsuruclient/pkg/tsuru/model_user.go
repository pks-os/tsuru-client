/*
 * Tsuru
 *
 * Open source, extensible and Docker-based Platform as a Service (PaaS)
 *
 * API version: 1.6
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package tsuru

// User object.
type User struct {
	Email       string           `json:"email,omitempty"`
	Roles       []RoleUser       `json:"roles,omitempty"`
	Permissions []PermissionUser `json:"permissions,omitempty"`
}