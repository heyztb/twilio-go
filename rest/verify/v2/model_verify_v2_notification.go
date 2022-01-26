/*
 * Twilio - Verify
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

// VerifyV2Notification struct for VerifyV2Notification
type VerifyV2Notification struct {
	// Account Sid.
	AccountSid *string `json:"account_sid,omitempty"`
	// Challenge Sid.
	ChallengeSid *string `json:"challenge_sid,omitempty"`
	// The date this Notification was created
	DateCreated *time.Time `json:"date_created,omitempty"`
	// Entity Sid.
	EntitySid *string `json:"entity_sid,omitempty"`
	// Unique external identifier of the Entity
	Identity *string `json:"identity,omitempty"`
	// The priority of the notification.
	Priority *string `json:"priority,omitempty"`
	// Service Sid.
	ServiceSid *string `json:"service_sid,omitempty"`
	// A string that uniquely identifies this Notification.
	Sid *string `json:"sid,omitempty"`
	// How long, in seconds, the notification is valid.
	Ttl *int `json:"ttl,omitempty"`
}
