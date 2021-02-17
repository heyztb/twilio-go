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

// PreviewMarketplaceAvailableAddOn struct for PreviewMarketplaceAvailableAddOn
type PreviewMarketplaceAvailableAddOn struct {
	ConfigurationSchema map[string]interface{} `json:"ConfigurationSchema,omitempty"`
	Description         string                 `json:"Description,omitempty"`
	FriendlyName        string                 `json:"FriendlyName,omitempty"`
	Links               map[string]interface{} `json:"Links,omitempty"`
	PricingType         string                 `json:"PricingType,omitempty"`
	Sid                 string                 `json:"Sid,omitempty"`
	Url                 string                 `json:"Url,omitempty"`
}
