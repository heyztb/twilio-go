/*
 * Twilio - Preview
 *
 * This is the public Twilio REST API.
 *
 * API version: 0.1.1
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// UpdateDeviceRequest struct for UpdateDeviceRequest
type UpdateDeviceRequest struct {
	// Specifies the unique string identifier of the Deployment group that this Device is going to be associated with.
	DeploymentSid string `json:"DeploymentSid,omitempty"`
	Enabled       bool   `json:"Enabled,omitempty"`
	// Provides a human readable descriptive text to be assigned to this Device, up to 256 characters long.
	FriendlyName string `json:"FriendlyName,omitempty"`
	// Provides an arbitrary string identifier representing a human user to be associated with this Device, up to 256 characters long.
	Identity string `json:"Identity,omitempty"`
}
