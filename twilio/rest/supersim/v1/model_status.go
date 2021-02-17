/*
 * Twilio - Supersim
 *
 * This is the public Twilio REST API.
 *
 * API version: 0.1.1
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// Status the model 'Status'
type Status string

// List of status
const (
	STATUS_NEW       Status = "new"
	STATUS_READY     Status = "ready"
	STATUS_ACTIVE    Status = "active"
	STATUS_INACTIVE  Status = "inactive"
	STATUS_SCHEDULED Status = "scheduled"
)
