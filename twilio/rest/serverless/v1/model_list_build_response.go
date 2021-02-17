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

// ListBuildResponse struct for ListBuildResponse
type ListBuildResponse struct {
	Builds []ServerlessV1ServiceBuild `json:"Builds,omitempty"`
	Meta   ListServiceResponseMeta    `json:"Meta,omitempty"`
}
