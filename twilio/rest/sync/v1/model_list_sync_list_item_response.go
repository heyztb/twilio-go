/*
 * Twilio - Sync
 *
 * This is the public Twilio REST API.
 *
 * API version: 0.1.1
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// ListSyncListItemResponse struct for ListSyncListItemResponse
type ListSyncListItemResponse struct {
	Items []SyncV1ServiceSyncListSyncListItem `json:"Items,omitempty"`
	Meta  ListServiceResponseMeta             `json:"Meta,omitempty"`
}
