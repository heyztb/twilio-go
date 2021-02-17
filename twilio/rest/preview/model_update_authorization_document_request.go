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

// UpdateAuthorizationDocumentRequest struct for UpdateAuthorizationDocumentRequest
type UpdateAuthorizationDocumentRequest struct {
	// A 34 character string that uniquely identifies the Address resource that is associated with this AuthorizationDocument.
	AddressSid string `json:"AddressSid,omitempty"`
	// Email recipients who will be informed when an Authorization Document has been sent and signed
	CcEmails []string `json:"CcEmails,omitempty"`
	// The contact phone number of the person authorized to sign the Authorization Document.
	ContactPhoneNumber string `json:"ContactPhoneNumber,omitempty"`
	// The title of the person authorized to sign the Authorization Document for this phone number.
	ContactTitle string `json:"ContactTitle,omitempty"`
	// Email that this AuthorizationDocument will be sent to for signing.
	Email string `json:"Email,omitempty"`
	// A list of HostedNumberOrder sids that this AuthorizationDocument will authorize for hosting phone number capabilities on Twilio's platform.
	HostedNumberOrderSids []string `json:"HostedNumberOrderSids,omitempty"`
	// Status of an instance resource. It can hold one of the values: 1. opened 2. signing, 3. signed LOA, 4. canceled, 5. failed. See the section entitled [Status Values](https://www.twilio.com/docs/api/phone-numbers/hosted-number-authorization-documents#status-values) for more information on each of these statuses.
	Status string `json:"Status,omitempty"`
}
