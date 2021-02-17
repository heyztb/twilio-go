/*
 * Twilio - Studio
 *
 * This is the public Twilio REST API.
 *
 * API version: 0.1.1
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// ListFlowRevisionResponse struct for ListFlowRevisionResponse
type ListFlowRevisionResponse struct {
	Meta      ListFlowResponseMeta       `json:"Meta,omitempty"`
	Revisions []StudioV2FlowFlowRevision `json:"Revisions,omitempty"`
}
