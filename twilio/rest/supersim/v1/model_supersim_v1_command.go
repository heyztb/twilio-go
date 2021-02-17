/*
 * Twilio - Supersim
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

// SupersimV1Command struct for SupersimV1Command
type SupersimV1Command struct {
	AccountSid  string    `json:"AccountSid,omitempty"`
	Command     string    `json:"Command,omitempty"`
	DateCreated time.Time `json:"DateCreated,omitempty"`
	DateUpdated time.Time `json:"DateUpdated,omitempty"`
	Direction   Direction `json:"Direction,omitempty"`
	Sid         string    `json:"Sid,omitempty"`
	SimSid      string    `json:"SimSid,omitempty"`
	Status      Status    `json:"Status,omitempty"`
	Url         string    `json:"Url,omitempty"`
}
