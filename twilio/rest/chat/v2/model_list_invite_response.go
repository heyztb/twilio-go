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

// ListInviteResponse struct for ListInviteResponse
type ListInviteResponse struct {
	Invites []ChatV2ServiceChannelInvite `json:"Invites,omitempty"`
	Meta    ListCredentialResponseMeta   `json:"Meta,omitempty"`
}
