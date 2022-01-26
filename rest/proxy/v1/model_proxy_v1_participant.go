/*
 * Twilio - Proxy
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.25.1
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"time"
)

// ProxyV1Participant struct for ProxyV1Participant
type ProxyV1Participant struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The ISO 8601 date the Participant was removed
	DateDeleted *time.Time `json:"date_deleted,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The string that you assigned to describe the participant
	FriendlyName *string `json:"friendly_name,omitempty"`
	// The phone number or channel identifier of the Participant
	Identifier *string `json:"identifier,omitempty"`
	// The URLs to resources related the participant
	Links *map[string]interface{} `json:"links,omitempty"`
	// The phone number or short code of the participant's partner
	ProxyIdentifier *string `json:"proxy_identifier,omitempty"`
	// The SID of the Proxy Identifier assigned to the Participant
	ProxyIdentifierSid *string `json:"proxy_identifier_sid,omitempty"`
	// The SID of the resource's parent Service
	ServiceSid *string `json:"service_sid,omitempty"`
	// The SID of the resource's parent Session
	SessionSid *string `json:"session_sid,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"sid,omitempty"`
	// The absolute URL of the Participant resource
	Url *string `json:"url,omitempty"`
}
