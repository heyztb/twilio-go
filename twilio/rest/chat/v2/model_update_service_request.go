/*
 * Twilio - Chat
 *
 * This is the public Twilio REST API.
 *
 * API version: 0.1.1
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// UpdateServiceRequest struct for UpdateServiceRequest
type UpdateServiceRequest struct {
	// DEPRECATED. The interval in seconds between consumption reports submission batches from client endpoints.
	ConsumptionReportInterval int32 `json:"ConsumptionReportInterval,omitempty"`
	// The channel role assigned to a channel creator when they join a new channel. See the [Role resource](https://www.twilio.com/docs/chat/rest/role-resource) for more info about roles.
	DefaultChannelCreatorRoleSid string `json:"DefaultChannelCreatorRoleSid,omitempty"`
	// The channel role assigned to users when they are added to a channel. See the [Role resource](https://www.twilio.com/docs/chat/rest/role-resource) for more info about roles.
	DefaultChannelRoleSid string `json:"DefaultChannelRoleSid,omitempty"`
	// The service role assigned to users when they are added to the service. See the [Role resource](https://www.twilio.com/docs/chat/rest/role-resource) for more info about roles.
	DefaultServiceRoleSid string `json:"DefaultServiceRoleSid,omitempty"`
	// A descriptive string that you create to describe the resource.
	FriendlyName string `json:"FriendlyName,omitempty"`
	// The maximum number of Members that can be added to Channels within this Service. Can be up to 1,000.
	LimitsChannelMembers int32 `json:"LimitsChannelMembers,omitempty"`
	// The maximum number of Channels Users can be a Member of within this Service. Can be up to 1,000.
	LimitsUserChannels int32 `json:"LimitsUserChannels,omitempty"`
	// The message to send when a media message has no text. Can be used as placeholder message.
	MediaCompatibilityMessage string `json:"MediaCompatibilityMessage,omitempty"`
	// Whether to send a notification when a member is added to a channel. The default is `false`.
	NotificationsAddedToChannelEnabled bool `json:"NotificationsAddedToChannelEnabled,omitempty"`
	// The name of the sound to play when a member is added to a channel and `notifications.added_to_channel.enabled` is `true`.
	NotificationsAddedToChannelSound string `json:"NotificationsAddedToChannelSound,omitempty"`
	// The template to use to create the notification text displayed when a member is added to a channel and `notifications.added_to_channel.enabled` is `true`.
	NotificationsAddedToChannelTemplate string `json:"NotificationsAddedToChannelTemplate,omitempty"`
	// Whether to send a notification when a user is invited to a channel. The default is `false`.
	NotificationsInvitedToChannelEnabled bool `json:"NotificationsInvitedToChannelEnabled,omitempty"`
	// The name of the sound to play when a user is invited to a channel and `notifications.invited_to_channel.enabled` is `true`.
	NotificationsInvitedToChannelSound string `json:"NotificationsInvitedToChannelSound,omitempty"`
	// The template to use to create the notification text displayed when a user is invited to a channel and `notifications.invited_to_channel.enabled` is `true`.
	NotificationsInvitedToChannelTemplate string `json:"NotificationsInvitedToChannelTemplate,omitempty"`
	// Whether to log notifications. The default is `false`.
	NotificationsLogEnabled bool `json:"NotificationsLogEnabled,omitempty"`
	// Whether the new message badge is enabled. The default is `false`.
	NotificationsNewMessageBadgeCountEnabled bool `json:"NotificationsNewMessageBadgeCountEnabled,omitempty"`
	// Whether to send a notification when a new message is added to a channel. The default is `false`.
	NotificationsNewMessageEnabled bool `json:"NotificationsNewMessageEnabled,omitempty"`
	// The name of the sound to play when a new message is added to a channel and `notifications.new_message.enabled` is `true`.
	NotificationsNewMessageSound string `json:"NotificationsNewMessageSound,omitempty"`
	// The template to use to create the notification text displayed when a new message is added to a channel and `notifications.new_message.enabled` is `true`.
	NotificationsNewMessageTemplate string `json:"NotificationsNewMessageTemplate,omitempty"`
	// Whether to send a notification to a user when they are removed from a channel. The default is `false`.
	NotificationsRemovedFromChannelEnabled bool `json:"NotificationsRemovedFromChannelEnabled,omitempty"`
	// The name of the sound to play to a user when they are removed from a channel and `notifications.removed_from_channel.enabled` is `true`.
	NotificationsRemovedFromChannelSound string `json:"NotificationsRemovedFromChannelSound,omitempty"`
	// The template to use to create the notification text displayed to a user when they are removed from a channel and `notifications.removed_from_channel.enabled` is `true`.
	NotificationsRemovedFromChannelTemplate string `json:"NotificationsRemovedFromChannelTemplate,omitempty"`
	// The number of times to retry a call to the `post_webhook_url` if the request times out (after 5 seconds) or it receives a 429, 503, or 504 HTTP response. The default is 0, which means the call won't be retried.
	PostWebhookRetryCount int32 `json:"PostWebhookRetryCount,omitempty"`
	// The URL for post-event webhooks, which are called by using the `webhook_method`. See [Webhook Events](https://www.twilio.com/docs/chat/webhook-events) for more details.
	PostWebhookUrl string `json:"PostWebhookUrl,omitempty"`
	// The number of times to retry a call to the `pre_webhook_url` if the request times out (after 5 seconds) or it receives a 429, 503, or 504 HTTP response. Default retry count is 0 times, which means the call won't be retried.
	PreWebhookRetryCount int32 `json:"PreWebhookRetryCount,omitempty"`
	// The URL for pre-event webhooks, which are called by using the `webhook_method`. See [Webhook Events](https://www.twilio.com/docs/chat/webhook-events) for more details.
	PreWebhookUrl string `json:"PreWebhookUrl,omitempty"`
	// Whether to enable the [Reachability Indicator](https://www.twilio.com/docs/chat/reachability-indicator) for this Service instance. The default is `false`.
	ReachabilityEnabled bool `json:"ReachabilityEnabled,omitempty"`
	// Whether to enable the [Message Consumption Horizon](https://www.twilio.com/docs/chat/consumption-horizon) feature. The default is `true`.
	ReadStatusEnabled bool `json:"ReadStatusEnabled,omitempty"`
	// How long in seconds after a `started typing` event until clients should assume that user is no longer typing, even if no `ended typing` message was received.  The default is 5 seconds.
	TypingIndicatorTimeout int32 `json:"TypingIndicatorTimeout,omitempty"`
	// The list of webhook events that are enabled for this Service instance. See [Webhook Events](https://www.twilio.com/docs/chat/webhook-events) for more details.
	WebhookFilters []string `json:"WebhookFilters,omitempty"`
	// The HTTP method to use for calls to the `pre_webhook_url` and `post_webhook_url` webhooks.  Can be: `POST` or `GET` and the default is `POST`. See [Webhook Events](https://www.twilio.com/docs/chat/webhook-events) for more details.
	WebhookMethod string `json:"WebhookMethod,omitempty"`
}
