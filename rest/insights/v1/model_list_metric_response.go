/*
 * Twilio - Insights
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.25.1
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListMetricResponse struct for ListMetricResponse
type ListMetricResponse struct {
	Meta    ListConferenceResponseMeta `json:"meta,omitempty"`
	Metrics []InsightsV1Metric         `json:"metrics,omitempty"`
}
