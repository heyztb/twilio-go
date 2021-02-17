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
	// The channel role assigned to a channel creator when they join a new channel. See the [Roles endpoint](https://www.twilio.com/docs/chat/api/roles) for more details.
	DefaultChannelCreatorRoleSid string `json:"DefaultChannelCreatorRoleSid,omitempty"`
	// The channel role assigned to users when they are added to a channel. See the [Roles endpoint](https://www.twilio.com/docs/chat/api/roles) for more details.
	DefaultChannelRoleSid string `json:"DefaultChannelRoleSid,omitempty"`
	// The service role assigned to users when they are added to the service. See the [Roles endpoint](https://www.twilio.com/docs/chat/api/roles) for more details.
	DefaultServiceRoleSid string `json:"DefaultServiceRoleSid,omitempty"`
	// A descriptive string that you create to describe the resource. It can be up to 64 characters long.
	FriendlyName string `json:"FriendlyName,omitempty"`
	// The maximum number of Members that can be added to Channels within this Service. Can be up to 1,000.
	LimitsChannelMembers int32 `json:"LimitsChannelMembers,omitempty"`
	// The maximum number of Channels Users can be a Member of within this Service. Can be up to 1,000.
	LimitsUserChannels int32 `json:"LimitsUserChannels,omitempty"`
	// Whether to send a notification when a member is added to a channel. Can be: `true` or `false` and the default is `false`.
	NotificationsAddedToChannelEnabled bool `json:"NotificationsAddedToChannelEnabled,omitempty"`
	// The template to use to create the notification text displayed when a member is added to a channel and `notifications.added_to_channel.enabled` is `true`.
	NotificationsAddedToChannelTemplate string `json:"NotificationsAddedToChannelTemplate,omitempty"`
	// Whether to send a notification when a user is invited to a channel. Can be: `true` or `false` and the default is `false`.
	NotificationsInvitedToChannelEnabled bool `json:"NotificationsInvitedToChannelEnabled,omitempty"`
	// The template to use to create the notification text displayed when a user is invited to a channel and `notifications.invited_to_channel.enabled` is `true`.
	NotificationsInvitedToChannelTemplate string `json:"NotificationsInvitedToChannelTemplate,omitempty"`
	// Whether to send a notification when a new message is added to a channel. Can be: `true` or `false` and the default is `false`.
	NotificationsNewMessageEnabled bool `json:"NotificationsNewMessageEnabled,omitempty"`
	// The template to use to create the notification text displayed when a new message is added to a channel and `notifications.new_message.enabled` is `true`.
	NotificationsNewMessageTemplate string `json:"NotificationsNewMessageTemplate,omitempty"`
	// Whether to send a notification to a user when they are removed from a channel. Can be: `true` or `false` and the default is `false`.
	NotificationsRemovedFromChannelEnabled bool `json:"NotificationsRemovedFromChannelEnabled,omitempty"`
	// The template to use to create the notification text displayed to a user when they are removed from a channel and `notifications.removed_from_channel.enabled` is `true`.
	NotificationsRemovedFromChannelTemplate string `json:"NotificationsRemovedFromChannelTemplate,omitempty"`
	// The URL for post-event webhooks, which are called by using the `webhook_method`. See [Webhook Events](https://www.twilio.com/docs/api/chat/webhooks) for more details.
	PostWebhookUrl string `json:"PostWebhookUrl,omitempty"`
	// The URL for pre-event webhooks, which are called by using the `webhook_method`. See [Webhook Events](https://www.twilio.com/docs/api/chat/webhooks) for more details.
	PreWebhookUrl string `json:"PreWebhookUrl,omitempty"`
	// Whether to enable the [Reachability Indicator](https://www.twilio.com/docs/chat/reachability-indicator) for this Service instance. The default is `false`.
	ReachabilityEnabled bool `json:"ReachabilityEnabled,omitempty"`
	// Whether to enable the [Message Consumption Horizon](https://www.twilio.com/docs/chat/consumption-horizon) feature. The default is `true`.
	ReadStatusEnabled bool `json:"ReadStatusEnabled,omitempty"`
	// How long in seconds after a `started typing` event until clients should assume that user is no longer typing, even if no `ended typing` message was received.  The default is 5 seconds.
	TypingIndicatorTimeout int32 `json:"TypingIndicatorTimeout,omitempty"`
	// The list of WebHook events that are enabled for this Service instance. See [Webhook Events](https://www.twilio.com/docs/chat/webhook-events) for more details.
	WebhookFilters []string `json:"WebhookFilters,omitempty"`
	// The HTTP method to use for calls to the `pre_webhook_url` and `post_webhook_url` webhooks.  Can be: `POST` or `GET` and the default is `POST`. See [Webhook Events](https://www.twilio.com/docs/chat/webhook-events) for more details.
	WebhookMethod string `json:"WebhookMethod,omitempty"`
	// The HTTP method to use when calling the `webhooks.on_channel_add.url`.
	WebhooksOnChannelAddMethod string `json:"WebhooksOnChannelAddMethod,omitempty"`
	// The URL of the webhook to call in response to the `on_channel_add` event using the `webhooks.on_channel_add.method` HTTP method.
	WebhooksOnChannelAddUrl string `json:"WebhooksOnChannelAddUrl,omitempty"`
	// The URL of the webhook to call in response to the `on_channel_added` event`.
	WebhooksOnChannelAddedMethod string `json:"WebhooksOnChannelAddedMethod,omitempty"`
	// The URL of the webhook to call in response to the `on_channel_added` event using the `webhooks.on_channel_added.method` HTTP method.
	WebhooksOnChannelAddedUrl string `json:"WebhooksOnChannelAddedUrl,omitempty"`
	// The HTTP method to use when calling the `webhooks.on_channel_destroy.url`.
	WebhooksOnChannelDestroyMethod string `json:"WebhooksOnChannelDestroyMethod,omitempty"`
	// The URL of the webhook to call in response to the `on_channel_destroy` event using the `webhooks.on_channel_destroy.method` HTTP method.
	WebhooksOnChannelDestroyUrl string `json:"WebhooksOnChannelDestroyUrl,omitempty"`
	// The HTTP method to use when calling the `webhooks.on_channel_destroyed.url`.
	WebhooksOnChannelDestroyedMethod string `json:"WebhooksOnChannelDestroyedMethod,omitempty"`
	// The URL of the webhook to call in response to the `on_channel_added` event using the `webhooks.on_channel_destroyed.method` HTTP method.
	WebhooksOnChannelDestroyedUrl string `json:"WebhooksOnChannelDestroyedUrl,omitempty"`
	// The HTTP method to use when calling the `webhooks.on_channel_update.url`.
	WebhooksOnChannelUpdateMethod string `json:"WebhooksOnChannelUpdateMethod,omitempty"`
	// The URL of the webhook to call in response to the `on_channel_update` event using the `webhooks.on_channel_update.method` HTTP method.
	WebhooksOnChannelUpdateUrl string `json:"WebhooksOnChannelUpdateUrl,omitempty"`
	// The HTTP method to use when calling the `webhooks.on_channel_updated.url`.
	WebhooksOnChannelUpdatedMethod string `json:"WebhooksOnChannelUpdatedMethod,omitempty"`
	// The URL of the webhook to call in response to the `on_channel_updated` event using the `webhooks.on_channel_updated.method` HTTP method.
	WebhooksOnChannelUpdatedUrl string `json:"WebhooksOnChannelUpdatedUrl,omitempty"`
	// The HTTP method to use when calling the `webhooks.on_member_add.url`.
	WebhooksOnMemberAddMethod string `json:"WebhooksOnMemberAddMethod,omitempty"`
	// The URL of the webhook to call in response to the `on_member_add` event using the `webhooks.on_member_add.method` HTTP method.
	WebhooksOnMemberAddUrl string `json:"WebhooksOnMemberAddUrl,omitempty"`
	// The HTTP method to use when calling the `webhooks.on_channel_updated.url`.
	WebhooksOnMemberAddedMethod string `json:"WebhooksOnMemberAddedMethod,omitempty"`
	// The URL of the webhook to call in response to the `on_channel_updated` event using the `webhooks.on_channel_updated.method` HTTP method.
	WebhooksOnMemberAddedUrl string `json:"WebhooksOnMemberAddedUrl,omitempty"`
	// The HTTP method to use when calling the `webhooks.on_member_remove.url`.
	WebhooksOnMemberRemoveMethod string `json:"WebhooksOnMemberRemoveMethod,omitempty"`
	// The URL of the webhook to call in response to the `on_member_remove` event using the `webhooks.on_member_remove.method` HTTP method.
	WebhooksOnMemberRemoveUrl string `json:"WebhooksOnMemberRemoveUrl,omitempty"`
	// The HTTP method to use when calling the `webhooks.on_member_removed.url`.
	WebhooksOnMemberRemovedMethod string `json:"WebhooksOnMemberRemovedMethod,omitempty"`
	// The URL of the webhook to call in response to the `on_member_removed` event using the `webhooks.on_member_removed.method` HTTP method.
	WebhooksOnMemberRemovedUrl string `json:"WebhooksOnMemberRemovedUrl,omitempty"`
	// The HTTP method to use when calling the `webhooks.on_message_remove.url`.
	WebhooksOnMessageRemoveMethod string `json:"WebhooksOnMessageRemoveMethod,omitempty"`
	// The URL of the webhook to call in response to the `on_message_remove` event using the `webhooks.on_message_remove.method` HTTP method.
	WebhooksOnMessageRemoveUrl string `json:"WebhooksOnMessageRemoveUrl,omitempty"`
	// The HTTP method to use when calling the `webhooks.on_message_removed.url`.
	WebhooksOnMessageRemovedMethod string `json:"WebhooksOnMessageRemovedMethod,omitempty"`
	// The URL of the webhook to call in response to the `on_message_removed` event using the `webhooks.on_message_removed.method` HTTP method.
	WebhooksOnMessageRemovedUrl string `json:"WebhooksOnMessageRemovedUrl,omitempty"`
	// The HTTP method to use when calling the `webhooks.on_message_send.url`.
	WebhooksOnMessageSendMethod string `json:"WebhooksOnMessageSendMethod,omitempty"`
	// The URL of the webhook to call in response to the `on_message_send` event using the `webhooks.on_message_send.method` HTTP method.
	WebhooksOnMessageSendUrl string `json:"WebhooksOnMessageSendUrl,omitempty"`
	// The URL of the webhook to call in response to the `on_message_sent` event`.
	WebhooksOnMessageSentMethod string `json:"WebhooksOnMessageSentMethod,omitempty"`
	// The URL of the webhook to call in response to the `on_message_sent` event using the `webhooks.on_message_sent.method` HTTP method.
	WebhooksOnMessageSentUrl string `json:"WebhooksOnMessageSentUrl,omitempty"`
	// The HTTP method to use when calling the `webhooks.on_message_update.url`.
	WebhooksOnMessageUpdateMethod string `json:"WebhooksOnMessageUpdateMethod,omitempty"`
	// The URL of the webhook to call in response to the `on_message_update` event using the `webhooks.on_message_update.method` HTTP method.
	WebhooksOnMessageUpdateUrl string `json:"WebhooksOnMessageUpdateUrl,omitempty"`
	// The HTTP method to use when calling the `webhooks.on_message_updated.url`.
	WebhooksOnMessageUpdatedMethod string `json:"WebhooksOnMessageUpdatedMethod,omitempty"`
	// The URL of the webhook to call in response to the `on_message_updated` event using the `webhooks.on_message_updated.method` HTTP method.
	WebhooksOnMessageUpdatedUrl string `json:"WebhooksOnMessageUpdatedUrl,omitempty"`
}
