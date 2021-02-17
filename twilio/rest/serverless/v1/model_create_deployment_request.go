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

// CreateDeploymentRequest struct for CreateDeploymentRequest
type CreateDeploymentRequest struct {
	// The SID of the Build for the Deployment.
	BuildSid string `json:"BuildSid,omitempty"`
}
