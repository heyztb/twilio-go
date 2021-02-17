/*
 * Twilio - Preview
 *
 * This is the public Twilio REST API.
 *
 * API version: 0.1.1
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// UpdateSyncListPermissionRequest struct for UpdateSyncListPermissionRequest
type UpdateSyncListPermissionRequest struct {
	// Boolean flag specifying whether the identity can delete the Sync List.
	Manage bool `json:"Manage"`
	// Boolean flag specifying whether the identity can read the Sync List.
	Read bool `json:"Read"`
	// Boolean flag specifying whether the identity can create, update and delete Items of the Sync List.
	Write bool `json:"Write"`
}
