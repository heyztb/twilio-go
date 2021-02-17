/*
 * Twilio - Insights
 *
 * This is the public Twilio REST API.
 *
 * API version: 0.1.1
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// StreamDirection the model 'StreamDirection'
type StreamDirection string

// List of stream_direction
const (
	STREAMDIRECTION_UNKNOWN  StreamDirection = "unknown"
	STREAMDIRECTION_INBOUND  StreamDirection = "inbound"
	STREAMDIRECTION_OUTBOUND StreamDirection = "outbound"
	STREAMDIRECTION_BOTH     StreamDirection = "both"
)
