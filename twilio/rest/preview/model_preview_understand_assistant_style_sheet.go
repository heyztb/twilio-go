/*
 * Twilio - Preview
 *
 * This is the public Twilio REST API.
 *
 * API version: 0.1.1
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// PreviewUnderstandAssistantStyleSheet struct for PreviewUnderstandAssistantStyleSheet
type PreviewUnderstandAssistantStyleSheet struct {
	AccountSid   string                 `json:"AccountSid,omitempty"`
	AssistantSid string                 `json:"AssistantSid,omitempty"`
	Data         map[string]interface{} `json:"Data,omitempty"`
	Url          string                 `json:"Url,omitempty"`
}
