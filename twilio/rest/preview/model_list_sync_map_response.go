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

// ListSyncMapResponse struct for ListSyncMapResponse
type ListSyncMapResponse struct {
	Maps []PreviewSyncServiceSyncMap `json:"Maps,omitempty"`
	Meta ListDayResponseMeta         `json:"Meta,omitempty"`
}
