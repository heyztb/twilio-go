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

// ListServiceConversationResponse struct for ListServiceConversationResponse
type ListServiceConversationResponse struct {
	Conversations []ConversationsV1ServiceServiceConversation `json:"Conversations,omitempty"`
	Meta          ListConversationResponseMeta                `json:"Meta,omitempty"`
}
