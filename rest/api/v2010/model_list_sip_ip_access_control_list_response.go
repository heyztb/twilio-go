/*
 * Twilio - Api
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.25.1
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListSipIpAccessControlListResponse struct for ListSipIpAccessControlListResponse
type ListSipIpAccessControlListResponse struct {
	End                  int                              `json:"end,omitempty"`
	FirstPageUri         string                           `json:"first_page_uri,omitempty"`
	IpAccessControlLists []ApiV2010SipIpAccessControlList `json:"ip_access_control_lists,omitempty"`
	NextPageUri          string                           `json:"next_page_uri,omitempty"`
	Page                 int                              `json:"page,omitempty"`
	PageSize             int                              `json:"page_size,omitempty"`
	PreviousPageUri      string                           `json:"previous_page_uri,omitempty"`
	Start                int                              `json:"start,omitempty"`
	Uri                  string                           `json:"uri,omitempty"`
}
