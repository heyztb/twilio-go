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

// ListServiceResponse struct for ListServiceResponse
type ListServiceResponse struct {
	Meta     ListServiceResponseMeta `json:"Meta,omitempty"`
	Services []ServerlessV1Service   `json:"Services,omitempty"`
}
