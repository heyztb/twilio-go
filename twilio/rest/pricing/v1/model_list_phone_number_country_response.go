/*
 * Twilio - Pricing
 *
 * This is the public Twilio REST API.
 *
 * API version: 0.1.1
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// ListPhoneNumberCountryResponse struct for ListPhoneNumberCountryResponse
type ListPhoneNumberCountryResponse struct {
	Countries []PricingV1PhoneNumberPhoneNumberCountry `json:"Countries,omitempty"`
	Meta      ListMessagingCountryResponseMeta         `json:"Meta,omitempty"`
}
