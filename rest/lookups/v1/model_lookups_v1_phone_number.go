/*
 * Twilio - Lookups
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.25.1
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// LookupsV1PhoneNumber struct for LookupsV1PhoneNumber
type LookupsV1PhoneNumber struct {
	// A JSON string with the results of the Add-ons you specified
	AddOns *map[string]interface{} `json:"add_ons,omitempty"`
	// The name of the phone number's owner
	CallerName *map[string]interface{} `json:"caller_name,omitempty"`
	// The telecom company that provides the phone number
	Carrier *map[string]interface{} `json:"carrier,omitempty"`
	// The ISO country code for the phone number
	CountryCode *string `json:"country_code,omitempty"`
	// The phone number, in national format
	NationalFormat *string `json:"national_format,omitempty"`
	// The phone number in E.164 format
	PhoneNumber *string `json:"phone_number,omitempty"`
	// The absolute URL of the resource
	Url *string `json:"url,omitempty"`
}
