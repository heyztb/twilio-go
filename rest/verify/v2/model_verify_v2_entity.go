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

// VerifyV2Entity struct for VerifyV2Entity
type VerifyV2Entity struct {
	// Account Sid.
	AccountSid *string `json:"account_sid,omitempty"`
	// The date this Entity was created
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The date this Entity was updated
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// Unique external identifier of the Entity
	Identity *string `json:"identity,omitempty"`
	// Nested resource URLs.
	Links *map[string]interface{} `json:"links,omitempty"`
	// Service Sid.
	ServiceSid *string `json:"service_sid,omitempty"`
	// A string that uniquely identifies this Entity.
	Sid *string `json:"sid,omitempty"`
	// The URL of this resource.
	Url *string `json:"url,omitempty"`
}
