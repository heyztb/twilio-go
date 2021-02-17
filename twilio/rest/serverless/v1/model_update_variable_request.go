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

// UpdateVariableRequest struct for UpdateVariableRequest
type UpdateVariableRequest struct {
	// A string by which the Variable resource can be referenced. It can be a maximum of 128 characters.
	Key string `json:"Key,omitempty"`
	// A string that contains the actual value of the Variable. It can be a maximum of 450 bytes in size.
	Value string `json:"Value,omitempty"`
}
