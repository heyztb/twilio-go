/*
 * Twilio - Ip_messaging
 *
 * This is the public Twilio REST API.
 *
 * API version: 0.1.1
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// CreateCredentialRequest struct for CreateCredentialRequest
type CreateCredentialRequest struct {
	ApiKey       string `json:"ApiKey,omitempty"`
	Certificate  string `json:"Certificate,omitempty"`
	FriendlyName string `json:"FriendlyName,omitempty"`
	PrivateKey   string `json:"PrivateKey,omitempty"`
	Sandbox      bool   `json:"Sandbox,omitempty"`
	Secret       string `json:"Secret,omitempty"`
	Type         string `json:"Type"`
}
