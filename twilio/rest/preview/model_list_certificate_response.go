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

// ListCertificateResponse struct for ListCertificateResponse
type ListCertificateResponse struct {
	Certificates []PreviewDeployedDevicesFleetCertificate `json:"Certificates,omitempty"`
	Meta         ListDayResponseMeta                      `json:"Meta,omitempty"`
}
