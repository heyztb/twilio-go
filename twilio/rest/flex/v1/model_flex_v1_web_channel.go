/*
 * Twilio - Flex
 *
 * This is the public Twilio REST API.
 *
 * API version: 0.1.1
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"time"
)

// FlexV1WebChannel struct for FlexV1WebChannel
type FlexV1WebChannel struct {
	AccountSid  string    `json:"AccountSid,omitempty"`
	DateCreated time.Time `json:"DateCreated,omitempty"`
	DateUpdated time.Time `json:"DateUpdated,omitempty"`
	FlexFlowSid string    `json:"FlexFlowSid,omitempty"`
	Sid         string    `json:"Sid,omitempty"`
	Url         string    `json:"Url,omitempty"`
}
