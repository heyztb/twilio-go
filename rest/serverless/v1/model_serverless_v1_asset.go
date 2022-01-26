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

// ServerlessV1Asset struct for ServerlessV1Asset
type ServerlessV1Asset struct {
	// The SID of the Account that created the Asset resource
	AccountSid *string `json:"account_sid,omitempty"`
	// The ISO 8601 date and time in GMT when the Asset resource was created
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The ISO 8601 date and time in GMT when the Asset resource was last updated
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The string that you assigned to describe the Asset resource
	FriendlyName *string `json:"friendly_name,omitempty"`
	// The URLs of the Asset resource's nested resources
	Links *map[string]interface{} `json:"links,omitempty"`
	// The SID of the Service that the Asset resource is associated with
	ServiceSid *string `json:"service_sid,omitempty"`
	// The unique string that identifies the Asset resource
	Sid *string `json:"sid,omitempty"`
	// The absolute URL of the Asset resource
	Url *string `json:"url,omitempty"`
}
