/*
 * Twilio - Voice
 *
 * This is the public Twilio REST API.
 *
 * API version: 0.1.1
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// VoiceV1DialingPermissionsDialingPermissionsCountry struct for VoiceV1DialingPermissionsDialingPermissionsCountry
type VoiceV1DialingPermissionsDialingPermissionsCountry struct {
	Continent                       string                 `json:"Continent,omitempty"`
	CountryCodes                    []string               `json:"CountryCodes,omitempty"`
	HighRiskSpecialNumbersEnabled   bool                   `json:"HighRiskSpecialNumbersEnabled,omitempty"`
	HighRiskTollfraudNumbersEnabled bool                   `json:"HighRiskTollfraudNumbersEnabled,omitempty"`
	IsoCode                         string                 `json:"IsoCode,omitempty"`
	Links                           map[string]interface{} `json:"Links,omitempty"`
	LowRiskNumbersEnabled           bool                   `json:"LowRiskNumbersEnabled,omitempty"`
	Name                            string                 `json:"Name,omitempty"`
	Url                             string                 `json:"Url,omitempty"`
}
