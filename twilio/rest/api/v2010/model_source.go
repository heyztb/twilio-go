/*
 * Twilio - Api
 *
 * This is the public Twilio REST API.
 *
 * API version: 0.1.1
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// Source the model 'Source'
type Source string

// List of source
const (
	SOURCE_DIAL_VERB                      Source = "DialVerb"
	SOURCE_CONFERENCE                     Source = "Conference"
	SOURCE_OUTBOUND_API                   Source = "OutboundAPI"
	SOURCE_TRUNKING                       Source = "Trunking"
	SOURCE_RECORD_VERB                    Source = "RecordVerb"
	SOURCE_START_CALL_RECORDING_API       Source = "StartCallRecordingAPI"
	SOURCE_START_CONFERENCE_RECORDING_API Source = "StartConferenceRecordingAPI"
)
