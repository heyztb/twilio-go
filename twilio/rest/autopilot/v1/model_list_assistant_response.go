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

// ListAssistantResponse struct for ListAssistantResponse
type ListAssistantResponse struct {
	Assistants []AutopilotV1Assistant    `json:"Assistants,omitempty"`
	Meta       ListAssistantResponseMeta `json:"Meta,omitempty"`
}
