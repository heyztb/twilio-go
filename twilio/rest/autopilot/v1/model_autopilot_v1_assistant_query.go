/*
 * Twilio - Autopilot
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

// AutopilotV1AssistantQuery struct for AutopilotV1AssistantQuery
type AutopilotV1AssistantQuery struct {
	AccountSid    string                 `json:"AccountSid,omitempty"`
	AssistantSid  string                 `json:"AssistantSid,omitempty"`
	DateCreated   time.Time              `json:"DateCreated,omitempty"`
	DateUpdated   time.Time              `json:"DateUpdated,omitempty"`
	DialogueSid   string                 `json:"DialogueSid,omitempty"`
	Language      string                 `json:"Language,omitempty"`
	ModelBuildSid string                 `json:"ModelBuildSid,omitempty"`
	Query         string                 `json:"Query,omitempty"`
	Results       map[string]interface{} `json:"Results,omitempty"`
	SampleSid     string                 `json:"SampleSid,omitempty"`
	Sid           string                 `json:"Sid,omitempty"`
	SourceChannel string                 `json:"SourceChannel,omitempty"`
	Status        string                 `json:"Status,omitempty"`
	Url           string                 `json:"Url,omitempty"`
}
