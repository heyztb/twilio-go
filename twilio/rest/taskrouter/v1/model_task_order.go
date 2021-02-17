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

// TaskOrder the model 'TaskOrder'
type TaskOrder string

// List of task_order
const (
	TASKORDER_FIFO TaskOrder = "FIFO"
	TASKORDER_LIFO TaskOrder = "LIFO"
)
