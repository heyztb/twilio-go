/*
 * Twilio - Serverless
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

// ServerlessV1ServiceEnvironmentDeployment struct for ServerlessV1ServiceEnvironmentDeployment
type ServerlessV1ServiceEnvironmentDeployment struct {
	AccountSid     string    `json:"AccountSid,omitempty"`
	BuildSid       string    `json:"BuildSid,omitempty"`
	DateCreated    time.Time `json:"DateCreated,omitempty"`
	DateUpdated    time.Time `json:"DateUpdated,omitempty"`
	EnvironmentSid string    `json:"EnvironmentSid,omitempty"`
	ServiceSid     string    `json:"ServiceSid,omitempty"`
	Sid            string    `json:"Sid,omitempty"`
	Url            string    `json:"Url,omitempty"`
}
