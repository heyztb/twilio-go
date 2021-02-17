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

// CreateInviteRequest struct for CreateInviteRequest
type CreateInviteRequest struct {
	// The `identity` value that uniquely identifies the new resource's [User](https://www.twilio.com/docs/api/chat/rest/v1/user) within the [Service](https://www.twilio.com/docs/api/chat/rest/v1/service). See [access tokens](https://www.twilio.com/docs/api/chat/guides/create-tokens) for more info.
	Identity string `json:"Identity"`
	// The SID of the [Role](https://www.twilio.com/docs/api/chat/rest/roles) assigned to the new member.
	RoleSid string `json:"RoleSid,omitempty"`
}
