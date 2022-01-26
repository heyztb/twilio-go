/*
 * Twilio - Conversations
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.25.1
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListServiceConversationParticipantResponse struct for ListServiceConversationParticipantResponse
type ListServiceConversationParticipantResponse struct {
	Meta         ListConversationResponseMeta                    `json:"meta,omitempty"`
	Participants []ConversationsV1ServiceConversationParticipant `json:"participants,omitempty"`
}
