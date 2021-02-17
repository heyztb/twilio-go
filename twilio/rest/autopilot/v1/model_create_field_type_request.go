/*
 * Twilio - Autopilot
 *
 * This is the public Twilio REST API.
 *
 * API version: 0.1.1
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// CreateFieldTypeRequest struct for CreateFieldTypeRequest
type CreateFieldTypeRequest struct {
	// A descriptive string that you create to describe the new resource. It is not unique and can be up to 255 characters long.
	FriendlyName string `json:"FriendlyName,omitempty"`
	// An application-defined string that uniquely identifies the new resource. It can be used as an alternative to the `sid` in the URL path to address the resource. The first 64 characters must be unique.
	UniqueName string `json:"UniqueName"`
}
