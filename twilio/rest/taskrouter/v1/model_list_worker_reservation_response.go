/*
 * Twilio - Taskrouter
 *
 * This is the public Twilio REST API.
 *
 * API version: 0.1.1
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// ListWorkerReservationResponse struct for ListWorkerReservationResponse
type ListWorkerReservationResponse struct {
	Meta         ListWorkspaceResponseMeta                      `json:"Meta,omitempty"`
	Reservations []TaskrouterV1WorkspaceWorkerWorkerReservation `json:"Reservations,omitempty"`
}
