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

// CreateUserRequest struct for CreateUserRequest
type CreateUserRequest struct {
	// A valid JSON string that contains application-specific data.
	Attributes string `json:"Attributes,omitempty"`
	// A descriptive string that you create to describe the new resource. This value is often used for display purposes.
	FriendlyName string `json:"FriendlyName,omitempty"`
	// The `identity` value that uniquely identifies the new resource's [User](https://www.twilio.com/docs/api/chat/rest/v1/user) within the [Service](https://www.twilio.com/docs/api/chat/rest/v1/service). This value is often a username or email address. See the Identity documentation for more details.
	Identity string `json:"Identity"`
	// The SID of the [Role](https://www.twilio.com/docs/api/chat/rest/roles) assigned to the new User.
	RoleSid string `json:"RoleSid,omitempty"`
}
