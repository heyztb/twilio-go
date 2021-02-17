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

// CreateOriginationUrlRequest struct for CreateOriginationUrlRequest
type CreateOriginationUrlRequest struct {
	// Whether the URL is enabled. The default is `true`.
	Enabled bool `json:"Enabled"`
	// A descriptive string that you create to describe the resource. It can be up to 64 characters long.
	FriendlyName string `json:"FriendlyName"`
	// The relative importance of the URI. Can be an integer from 0 to 65535, inclusive, and the default is 10. The lowest number represents the most important URI.
	Priority int32 `json:"Priority"`
	// The SIP address you want Twilio to route your Origination calls to. This must be a `sip:` schema.
	SipUrl string `json:"SipUrl"`
	// The value that determines the relative share of the load the URI should receive compared to other URIs with the same priority. Can be an integer from 1 to 65535, inclusive, and the default is 10. URLs with higher values receive more load than those with lower ones with the same priority.
	Weight int32 `json:"Weight"`
}
