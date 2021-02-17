/*
 * Twilio - Accounts
 *
 * This is the public Twilio REST API.
 *
 * API version: 0.1.1
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// ListCredentialPublicKeyResponse struct for ListCredentialPublicKeyResponse
type ListCredentialPublicKeyResponse struct {
	Credentials []AccountsV1CredentialCredentialPublicKey `json:"Credentials,omitempty"`
	Meta        ListCredentialAwsResponseMeta             `json:"Meta,omitempty"`
}
