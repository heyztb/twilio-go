/*
 * Twilio - Wireless
 *
 * This is the public Twilio REST API.
 *
 * API version: 0.1.1
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// ListAccountUsageRecordResponse struct for ListAccountUsageRecordResponse
type ListAccountUsageRecordResponse struct {
	Meta         ListCommandResponseMeta        `json:"Meta,omitempty"`
	UsageRecords []WirelessV1AccountUsageRecord `json:"UsageRecords,omitempty"`
}
