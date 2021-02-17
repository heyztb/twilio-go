/*
 * Twilio - Trunking
 *
 * This is the public Twilio REST API.
 *
 * API version: 0.1.1
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// CreateCredentialListRequest struct for CreateCredentialListRequest
type CreateCredentialListRequest struct {
	// The SID of the [Credential List](https://www.twilio.com/docs/voice/sip/api/sip-credentiallist-resource) that you want to associate with the trunk. Once associated, we will authenticate access to the trunk against this list.
	CredentialListSid string `json:"CredentialListSid"`
}
