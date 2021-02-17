/*
 * Twilio - Conversations
 *
 * This is the public Twilio REST API.
 *
 * API version: 0.1.1
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// ListServiceBindingResponse struct for ListServiceBindingResponse
type ListServiceBindingResponse struct {
	Bindings []ConversationsV1ServiceServiceBinding `json:"Bindings,omitempty"`
	Meta     ListConversationResponseMeta           `json:"Meta,omitempty"`
}
