/*
 * Twilio - Events
 *
 * This is the public Twilio REST API.
 *
 * API version: 0.1.1
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// ListSinkResponse struct for ListSinkResponse
type ListSinkResponse struct {
	Meta  ListVersionResponseMeta `json:"Meta,omitempty"`
	Sinks []EventsV1Sink          `json:"Sinks,omitempty"`
}
