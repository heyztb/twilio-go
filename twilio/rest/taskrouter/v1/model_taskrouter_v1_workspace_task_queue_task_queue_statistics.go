/*
 * Twilio - Taskrouter
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.15.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// TaskrouterV1WorkspaceTaskQueueTaskQueueStatistics struct for TaskrouterV1WorkspaceTaskQueueTaskQueueStatistics
type TaskrouterV1WorkspaceTaskQueueTaskQueueStatistics struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// An object that contains the cumulative statistics for the TaskQueue
	Cumulative *map[string]interface{} `json:"cumulative,omitempty"`
	// An object that contains the real-time statistics for the TaskQueue
	Realtime *map[string]interface{} `json:"realtime,omitempty"`
	// The SID of the TaskQueue from which these statistics were calculated
	TaskQueueSid *string `json:"task_queue_sid,omitempty"`
	// The absolute URL of the TaskQueue statistics resource
	Url *string `json:"url,omitempty"`
	// The SID of the Workspace that contains the TaskQueue
	WorkspaceSid *string `json:"workspace_sid,omitempty"`
}
