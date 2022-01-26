/*
 * Twilio - Serverless
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

// ServerlessV1Function struct for ServerlessV1Function
type ServerlessV1Function struct {
	// The SID of the Account that created the Function resource
	AccountSid *string `json:"account_sid,omitempty"`
	// The ISO 8601 date and time in GMT when the Function resource was created
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The ISO 8601 date and time in GMT when the Function resource was last updated
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The string that you assigned to describe the Function resource
	FriendlyName *string `json:"friendly_name,omitempty"`
	// The URLs of nested resources of the Function resource
	Links *map[string]interface{} `json:"links,omitempty"`
	// The SID of the Service that the Function resource is associated with
	ServiceSid *string `json:"service_sid,omitempty"`
	// The unique string that identifies the Function resource
	Sid *string `json:"sid,omitempty"`
	// The absolute URL of the Function resource
	Url *string `json:"url,omitempty"`
}
