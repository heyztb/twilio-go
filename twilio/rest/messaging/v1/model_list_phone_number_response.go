/*
 * Twilio - Messaging
 *
 * This is the public Twilio REST API.
 *
 * API version: 0.1.1
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// ListPhoneNumberResponse struct for ListPhoneNumberResponse
type ListPhoneNumberResponse struct {
	Meta         ListServiceResponseMeta         `json:"Meta,omitempty"`
	PhoneNumbers []MessagingV1ServicePhoneNumber `json:"PhoneNumbers,omitempty"`
}
