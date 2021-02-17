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

import (
	"time"
)

// TaskrouterV1WorkspaceWorkerWorkerReservation struct for TaskrouterV1WorkspaceWorkerWorkerReservation
type TaskrouterV1WorkspaceWorkerWorkerReservation struct {
	AccountSid        string                 `json:"AccountSid,omitempty"`
	DateCreated       time.Time              `json:"DateCreated,omitempty"`
	DateUpdated       time.Time              `json:"DateUpdated,omitempty"`
	Links             map[string]interface{} `json:"Links,omitempty"`
	ReservationStatus Status                 `json:"ReservationStatus,omitempty"`
	Sid               string                 `json:"Sid,omitempty"`
	TaskSid           string                 `json:"TaskSid,omitempty"`
	Url               string                 `json:"Url,omitempty"`
	WorkerName        string                 `json:"WorkerName,omitempty"`
	WorkerSid         string                 `json:"WorkerSid,omitempty"`
	WorkspaceSid      string                 `json:"WorkspaceSid,omitempty"`
}
