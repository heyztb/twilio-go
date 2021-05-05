/*
 * Twilio - Conversations
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.15.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"time"
)

// ConversationsV1ServiceServiceUser struct for ConversationsV1ServiceServiceUser
type ConversationsV1ServiceServiceUser struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// The JSON Object string that stores application-specific data
	Attributes *string `json:"attributes,omitempty"`
	// The SID of the Conversation Service that the resource is associated with
	ChatServiceSid *string `json:"chat_service_sid,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The string that you assigned to describe the resource
	FriendlyName *string `json:"friendly_name,omitempty"`
	// The string that identifies the resource's User
	Identity *string `json:"identity,omitempty"`
	// Whether the User has a potentially valid Push Notification registration for this Conversations Service
	IsNotifiable *bool `json:"is_notifiable,omitempty"`
	// Whether the User is actively connected to this Conversations Service and online
	IsOnline *bool `json:"is_online,omitempty"`
	// The SID of a service-level Role assigned to the user
	RoleSid *string `json:"role_sid,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"sid,omitempty"`
	// An absolute URL for this user.
	Url *string `json:"url,omitempty"`
}
