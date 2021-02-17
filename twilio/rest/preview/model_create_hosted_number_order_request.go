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

// CreateHostedNumberOrderRequest struct for CreateHostedNumberOrderRequest
type CreateHostedNumberOrderRequest struct {
	// This defaults to the AccountSid of the authorization the user is using. This can be provided to specify a subaccount to add the HostedNumberOrder to.
	AccountSid string `json:"AccountSid,omitempty"`
	// Optional. A 34 character string that uniquely identifies the Address resource that represents the address of the owner of this phone number.
	AddressSid string `json:"AddressSid,omitempty"`
	// Optional. A list of emails that the LOA document for this HostedNumberOrder will be carbon copied to.
	CcEmails []string `json:"CcEmails,omitempty"`
	// Optional. Email of the owner of this phone number that is being hosted.
	Email string `json:"Email,omitempty"`
	// A 64 character string that is a human readable text that describes this resource.
	FriendlyName string `json:"FriendlyName,omitempty"`
	// The number to host in [+E.164](https://en.wikipedia.org/wiki/E.164) format
	PhoneNumber string `json:"PhoneNumber"`
	// Optional. The 34 character sid of the application Twilio should use to handle SMS messages sent to this number. If a `SmsApplicationSid` is present, Twilio will ignore all of the SMS urls above and use those set on the application.
	SmsApplicationSid string `json:"SmsApplicationSid,omitempty"`
	// Used to specify that the SMS capability will be hosted on Twilio's platform.
	SmsCapability bool `json:"SmsCapability"`
	// The HTTP method that should be used to request the SmsFallbackUrl. Must be either `GET` or `POST`. This will be copied onto the IncomingPhoneNumber resource.
	SmsFallbackMethod string `json:"SmsFallbackMethod,omitempty"`
	// A URL that Twilio will request if an error occurs requesting or executing the TwiML defined by SmsUrl. This will be copied onto the IncomingPhoneNumber resource.
	SmsFallbackUrl string `json:"SmsFallbackUrl,omitempty"`
	// The HTTP method that should be used to request the SmsUrl. Must be either `GET` or `POST`.  This will be copied onto the IncomingPhoneNumber resource.
	SmsMethod string `json:"SmsMethod,omitempty"`
	// The URL that Twilio should request when somebody sends an SMS to the phone number. This will be copied onto the IncomingPhoneNumber resource.
	SmsUrl string `json:"SmsUrl,omitempty"`
	// Optional. The Status Callback Method attached to the IncomingPhoneNumber resource.
	StatusCallbackMethod string `json:"StatusCallbackMethod,omitempty"`
	// Optional. The Status Callback URL attached to the IncomingPhoneNumber resource.
	StatusCallbackUrl string `json:"StatusCallbackUrl,omitempty"`
	// Optional. Provides a unique and addressable name to be assigned to this HostedNumberOrder, assigned by the developer, to be optionally used in addition to SID.
	UniqueName string `json:"UniqueName,omitempty"`
	// Optional. The unique sid identifier of the Identity Document that represents the document for verifying ownership of the number to be hosted. Required when VerificationType is phone-bill.
	VerificationDocumentSid string `json:"VerificationDocumentSid,omitempty"`
	// Optional. The method used for verifying ownership of the number to be hosted. One of phone-call (default) or phone-bill.
	VerificationType string `json:"VerificationType,omitempty"`
}
