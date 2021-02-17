/*
 * Twilio - Api
 *
 * This is the public Twilio REST API.
 *
 * API version: 0.1.1
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// ApiV2010AccountToken struct for ApiV2010AccountToken
type ApiV2010AccountToken struct {
	AccountSid  string                   `json:"AccountSid,omitempty"`
	DateCreated string                   `json:"DateCreated,omitempty"`
	DateUpdated string                   `json:"DateUpdated,omitempty"`
	IceServers  []map[string]interface{} `json:"IceServers,omitempty"`
	Password    string                   `json:"Password,omitempty"`
	Ttl         string                   `json:"Ttl,omitempty"`
	Username    string                   `json:"Username,omitempty"`
}
