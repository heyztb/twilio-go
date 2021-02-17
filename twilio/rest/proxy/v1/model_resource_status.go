/*
 * Twilio - Proxy
 *
 * This is the public Twilio REST API.
 *
 * API version: 0.1.1
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// ResourceStatus the model 'ResourceStatus'
type ResourceStatus string

// List of resource_status
const (
	RESOURCESTATUS_ACCEPTED         ResourceStatus = "accepted"
	RESOURCESTATUS_ANSWERED         ResourceStatus = "answered"
	RESOURCESTATUS_BUSY             ResourceStatus = "busy"
	RESOURCESTATUS_CANCELED         ResourceStatus = "canceled"
	RESOURCESTATUS_COMPLETED        ResourceStatus = "completed"
	RESOURCESTATUS_DELETED          ResourceStatus = "deleted"
	RESOURCESTATUS_DELIVERED        ResourceStatus = "delivered"
	RESOURCESTATUS_DELIVERY_UNKNOWN ResourceStatus = "delivery-unknown"
	RESOURCESTATUS_FAILED           ResourceStatus = "failed"
	RESOURCESTATUS_IN_PROGRESS      ResourceStatus = "in-progress"
	RESOURCESTATUS_INITIATED        ResourceStatus = "initiated"
	RESOURCESTATUS_NO_ANSWER        ResourceStatus = "no-answer"
	RESOURCESTATUS_QUEUED           ResourceStatus = "queued"
	RESOURCESTATUS_RECEIVED         ResourceStatus = "received"
	RESOURCESTATUS_RECEIVING        ResourceStatus = "receiving"
	RESOURCESTATUS_RINGING          ResourceStatus = "ringing"
	RESOURCESTATUS_SCHEDULED        ResourceStatus = "scheduled"
	RESOURCESTATUS_SENDING          ResourceStatus = "sending"
	RESOURCESTATUS_SENT             ResourceStatus = "sent"
	RESOURCESTATUS_UNDELIVERED      ResourceStatus = "undelivered"
	RESOURCESTATUS_UNKNOWN          ResourceStatus = "unknown"
)
