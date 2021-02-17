/*
 * Twilio - Ip_messaging
 *
 * This is the public Twilio REST API.
 *
 * API version: 0.1.1
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// ListChannelResponse struct for ListChannelResponse
type ListChannelResponse struct {
	Channels []IpMessagingV2ServiceChannel `json:"Channels,omitempty"`
	Meta     ListCredentialResponseMeta    `json:"Meta,omitempty"`
}
