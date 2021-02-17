/*
 * Twilio - Chat
 *
 * This is the public Twilio REST API.
 *
 * API version: 0.1.1
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// CreateRoleRequest struct for CreateRoleRequest
type CreateRoleRequest struct {
	// A descriptive string that you create to describe the new resource. It can be up to 64 characters long.
	FriendlyName string `json:"FriendlyName"`
	// A permission that you grant to the new role. Only one permission can be granted per parameter. To assign more than one permission, repeat this parameter for each permission value. The values for this parameter depend on the role's `type` and are described in the documentation.
	Permission []string `json:"Permission"`
	// The type of role. Can be: `channel` for [Channel](https://www.twilio.com/docs/chat/api/channels) roles or `deployment` for [Service](https://www.twilio.com/docs/chat/api/services) roles.
	Type string `json:"Type"`
}
