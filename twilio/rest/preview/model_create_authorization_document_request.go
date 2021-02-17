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

// CreateAuthorizationDocumentRequest struct for CreateAuthorizationDocumentRequest
type CreateAuthorizationDocumentRequest struct {
	// A 34 character string that uniquely identifies the Address resource that is associated with this AuthorizationDocument.
	AddressSid string `json:"AddressSid"`
	// Email recipients who will be informed when an Authorization Document has been sent and signed.
	CcEmails []string `json:"CcEmails,omitempty"`
	// The contact phone number of the person authorized to sign the Authorization Document.
	ContactPhoneNumber string `json:"ContactPhoneNumber"`
	// The title of the person authorized to sign the Authorization Document for this phone number.
	ContactTitle string `json:"ContactTitle"`
	// Email that this AuthorizationDocument will be sent to for signing.
	Email string `json:"Email"`
	// A list of HostedNumberOrder sids that this AuthorizationDocument will authorize for hosting phone number capabilities on Twilio's platform.
	HostedNumberOrderSids []string `json:"HostedNumberOrderSids"`
}
