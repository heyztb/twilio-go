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

// TaskrouterV1WorkspaceWorkerWorkersCumulativeStatistics struct for TaskrouterV1WorkspaceWorkerWorkersCumulativeStatistics
type TaskrouterV1WorkspaceWorkerWorkersCumulativeStatistics struct {
	AccountSid            string                   `json:"AccountSid,omitempty"`
	ActivityDurations     []map[string]interface{} `json:"ActivityDurations,omitempty"`
	EndTime               time.Time                `json:"EndTime,omitempty"`
	ReservationsAccepted  int32                    `json:"ReservationsAccepted,omitempty"`
	ReservationsCanceled  int32                    `json:"ReservationsCanceled,omitempty"`
	ReservationsCreated   int32                    `json:"ReservationsCreated,omitempty"`
	ReservationsRejected  int32                    `json:"ReservationsRejected,omitempty"`
	ReservationsRescinded int32                    `json:"ReservationsRescinded,omitempty"`
	ReservationsTimedOut  int32                    `json:"ReservationsTimedOut,omitempty"`
	StartTime             time.Time                `json:"StartTime,omitempty"`
	Url                   string                   `json:"Url,omitempty"`
	WorkspaceSid          string                   `json:"WorkspaceSid,omitempty"`
}
