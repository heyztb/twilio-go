/*
 * Twilio - Video
 *
 * This is the public Twilio REST API.
 *
 * API version: 0.1.1
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// CreateCompositionSettingsRequest struct for CreateCompositionSettingsRequest
type CreateCompositionSettingsRequest struct {
	// The SID of the stored Credential resource.
	AwsCredentialsSid string `json:"AwsCredentialsSid,omitempty"`
	// The URL of the AWS S3 bucket where the compositions should be stored. We only support DNS-compliant URLs like `https://documentation-example-twilio-bucket/compositions`, where `compositions` is the path in which you want the compositions to be stored. This URL accepts only URI-valid characters, as described in the <a href='https://tools.ietf.org/html/rfc3986#section-2'>RFC 3986</a>.
	AwsS3Url string `json:"AwsS3Url,omitempty"`
	// Whether all compositions should be written to the `aws_s3_url`. When `false`, all compositions are stored in our cloud.
	AwsStorageEnabled bool `json:"AwsStorageEnabled,omitempty"`
	// Whether all compositions should be stored in an encrypted form. The default is `false`.
	EncryptionEnabled bool `json:"EncryptionEnabled,omitempty"`
	// The SID of the Public Key resource to use for encryption.
	EncryptionKeySid string `json:"EncryptionKeySid,omitempty"`
	// A descriptive string that you create to describe the resource and show to the user in the console
	FriendlyName string `json:"FriendlyName"`
}
