/*
 * Twilio - Voice
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.29.1
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListByocTrunkResponse struct for ListByocTrunkResponse
type ListByocTrunkResponse struct {
	ByocTrunks []VoiceV1ByocTrunk        `json:"byoc_trunks,omitempty"`
	Meta       ListByocTrunkResponseMeta `json:"meta,omitempty"`
}
