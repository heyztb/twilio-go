/*
 * Twilio - Autopilot
 *
 * This is the public Twilio REST API.
 *
 * API version: 0.1.1
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// AutopilotV1AssistantDialogue struct for AutopilotV1AssistantDialogue
type AutopilotV1AssistantDialogue struct {
	AccountSid   string                 `json:"AccountSid,omitempty"`
	AssistantSid string                 `json:"AssistantSid,omitempty"`
	Data         map[string]interface{} `json:"Data,omitempty"`
	Sid          string                 `json:"Sid,omitempty"`
	Url          string                 `json:"Url,omitempty"`
}
