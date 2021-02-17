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

// UpdateTaskQueueRequest struct for UpdateTaskQueueRequest
type UpdateTaskQueueRequest struct {
	// The SID of the Activity to assign Workers when a task is assigned for them.
	AssignmentActivitySid string `json:"AssignmentActivitySid,omitempty"`
	// A descriptive string that you create to describe the TaskQueue. For example `Support-Tier 1`, `Sales`, or `Escalation`.
	FriendlyName string `json:"FriendlyName,omitempty"`
	// The maximum number of Workers to create reservations for the assignment of a task while in the queue. Maximum of 50.
	MaxReservedWorkers int32 `json:"MaxReservedWorkers,omitempty"`
	// The SID of the Activity to assign Workers when a task is reserved for them.
	ReservationActivitySid string `json:"ReservationActivitySid,omitempty"`
	// A string describing the Worker selection criteria for any Tasks that enter the TaskQueue. For example '\"language\" == \"spanish\"' If no TargetWorkers parameter is provided, Tasks will wait in the queue until they are either deleted or moved to another queue. Additional examples on how to describing Worker selection criteria below.
	TargetWorkers string `json:"TargetWorkers,omitempty"`
	// How Tasks will be assigned to Workers. Can be: `FIFO` or `LIFO` and the default is `FIFO`. Use `FIFO` to assign the oldest task first and `LIFO` to assign the most recent task first. For more information, see [Queue Ordering](https://www.twilio.com/docs/taskrouter/queue-ordering-last-first-out-lifo).
	TaskOrder string `json:"TaskOrder,omitempty"`
}
