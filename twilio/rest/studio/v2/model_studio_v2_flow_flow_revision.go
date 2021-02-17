/*
 * Twilio - Studio
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

// StudioV2FlowFlowRevision struct for StudioV2FlowFlowRevision
type StudioV2FlowFlowRevision struct {
	AccountSid    string                   `json:"AccountSid,omitempty"`
	CommitMessage string                   `json:"CommitMessage,omitempty"`
	DateCreated   time.Time                `json:"DateCreated,omitempty"`
	DateUpdated   time.Time                `json:"DateUpdated,omitempty"`
	Definition    map[string]interface{}   `json:"Definition,omitempty"`
	Errors        []map[string]interface{} `json:"Errors,omitempty"`
	FriendlyName  string                   `json:"FriendlyName,omitempty"`
	Revision      int32                    `json:"Revision,omitempty"`
	Sid           string                   `json:"Sid,omitempty"`
	Status        Status                   `json:"Status,omitempty"`
	Url           string                   `json:"Url,omitempty"`
	Valid         bool                     `json:"Valid,omitempty"`
}
