/*
 * Twilio - Accounts
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

// AccountsV1SecondaryAuthToken struct for AccountsV1SecondaryAuthToken
type AccountsV1SecondaryAuthToken struct {
	AccountSid         string    `json:"AccountSid,omitempty"`
	DateCreated        time.Time `json:"DateCreated,omitempty"`
	DateUpdated        time.Time `json:"DateUpdated,omitempty"`
	SecondaryAuthToken string    `json:"SecondaryAuthToken,omitempty"`
	Url                string    `json:"Url,omitempty"`
}
