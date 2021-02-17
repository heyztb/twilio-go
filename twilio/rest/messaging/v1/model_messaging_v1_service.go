/*
 * Twilio - Messaging
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

// MessagingV1Service struct for MessagingV1Service
type MessagingV1Service struct {
	AccountSid            string                 `json:"AccountSid,omitempty"`
	AreaCodeGeomatch      bool                   `json:"AreaCodeGeomatch,omitempty"`
	DateCreated           time.Time              `json:"DateCreated,omitempty"`
	DateUpdated           time.Time              `json:"DateUpdated,omitempty"`
	FallbackMethod        HttpMethod             `json:"FallbackMethod,omitempty"`
	FallbackToLongCode    bool                   `json:"FallbackToLongCode,omitempty"`
	FallbackUrl           string                 `json:"FallbackUrl,omitempty"`
	FriendlyName          string                 `json:"FriendlyName,omitempty"`
	InboundMethod         HttpMethod             `json:"InboundMethod,omitempty"`
	InboundRequestUrl     string                 `json:"InboundRequestUrl,omitempty"`
	Links                 map[string]interface{} `json:"Links,omitempty"`
	MmsConverter          bool                   `json:"MmsConverter,omitempty"`
	ScanMessageContent    ScanMessageContent     `json:"ScanMessageContent,omitempty"`
	Sid                   string                 `json:"Sid,omitempty"`
	SmartEncoding         bool                   `json:"SmartEncoding,omitempty"`
	StatusCallback        string                 `json:"StatusCallback,omitempty"`
	StickySender          bool                   `json:"StickySender,omitempty"`
	SynchronousValidation bool                   `json:"SynchronousValidation,omitempty"`
	Url                   string                 `json:"Url,omitempty"`
	ValidityPeriod        int32                  `json:"ValidityPeriod,omitempty"`
}
