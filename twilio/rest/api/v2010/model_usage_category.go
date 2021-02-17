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

// UsageCategory the model 'UsageCategory'
type UsageCategory string

// List of usage_category
const (
	USAGECATEGORY_AGENT_CONFERENCE                                       UsageCategory = "agent-conference"
	USAGECATEGORY_ANSWERING_MACHINE_DETECTION                            UsageCategory = "answering-machine-detection"
	USAGECATEGORY_AUTHY_AUTHENTICATIONS                                  UsageCategory = "authy-authentications"
	USAGECATEGORY_AUTHY_CALLS_OUTBOUND                                   UsageCategory = "authy-calls-outbound"
	USAGECATEGORY_AUTHY_MONTHLY_FEES                                     UsageCategory = "authy-monthly-fees"
	USAGECATEGORY_AUTHY_PHONE_INTELLIGENCE                               UsageCategory = "authy-phone-intelligence"
	USAGECATEGORY_AUTHY_PHONE_VERIFICATIONS                              UsageCategory = "authy-phone-verifications"
	USAGECATEGORY_AUTHY_SMS_OUTBOUND                                     UsageCategory = "authy-sms-outbound"
	USAGECATEGORY_CALL_PROGESS_EVENTS                                    UsageCategory = "call-progess-events"
	USAGECATEGORY_CALLERIDLOOKUPS                                        UsageCategory = "calleridlookups"
	USAGECATEGORY_CALLS                                                  UsageCategory = "calls"
	USAGECATEGORY_CALLS_CLIENT                                           UsageCategory = "calls-client"
	USAGECATEGORY_CALLS_GLOBALCONFERENCE                                 UsageCategory = "calls-globalconference"
	USAGECATEGORY_CALLS_INBOUND                                          UsageCategory = "calls-inbound"
	USAGECATEGORY_CALLS_INBOUND_LOCAL                                    UsageCategory = "calls-inbound-local"
	USAGECATEGORY_CALLS_INBOUND_MOBILE                                   UsageCategory = "calls-inbound-mobile"
	USAGECATEGORY_CALLS_INBOUND_TOLLFREE                                 UsageCategory = "calls-inbound-tollfree"
	USAGECATEGORY_CALLS_OUTBOUND                                         UsageCategory = "calls-outbound"
	USAGECATEGORY_CALLS_PAY_VERB_TRANSACTIONS                            UsageCategory = "calls-pay-verb-transactions"
	USAGECATEGORY_CALLS_RECORDINGS                                       UsageCategory = "calls-recordings"
	USAGECATEGORY_CALLS_SIP                                              UsageCategory = "calls-sip"
	USAGECATEGORY_CALLS_SIP_INBOUND                                      UsageCategory = "calls-sip-inbound"
	USAGECATEGORY_CALLS_SIP_OUTBOUND                                     UsageCategory = "calls-sip-outbound"
	USAGECATEGORY_CARRIER_LOOKUPS                                        UsageCategory = "carrier-lookups"
	USAGECATEGORY_CONVERSATIONS                                          UsageCategory = "conversations"
	USAGECATEGORY_CONVERSATIONS_API_REQUESTS                             UsageCategory = "conversations-api-requests"
	USAGECATEGORY_CONVERSATIONS_CONVERSATION_EVENTS                      UsageCategory = "conversations-conversation-events"
	USAGECATEGORY_CONVERSATIONS_ENDPOINT_CONNECTIVITY                    UsageCategory = "conversations-endpoint-connectivity"
	USAGECATEGORY_CONVERSATIONS_EVENTS                                   UsageCategory = "conversations-events"
	USAGECATEGORY_CONVERSATIONS_PARTICIPANT_EVENTS                       UsageCategory = "conversations-participant-events"
	USAGECATEGORY_CONVERSATIONS_PARTICIPANTS                             UsageCategory = "conversations-participants"
	USAGECATEGORY_CPS                                                    UsageCategory = "cps"
	USAGECATEGORY_FRAUD_LOOKUPS                                          UsageCategory = "fraud-lookups"
	USAGECATEGORY_GROUP_ROOMS                                            UsageCategory = "group-rooms"
	USAGECATEGORY_GROUP_ROOMS_DATA_TRACK                                 UsageCategory = "group-rooms-data-track"
	USAGECATEGORY_GROUP_ROOMS_ENCRYPTED_MEDIA_RECORDED                   UsageCategory = "group-rooms-encrypted-media-recorded"
	USAGECATEGORY_GROUP_ROOMS_MEDIA_DOWNLOADED                           UsageCategory = "group-rooms-media-downloaded"
	USAGECATEGORY_GROUP_ROOMS_MEDIA_RECORDED                             UsageCategory = "group-rooms-media-recorded"
	USAGECATEGORY_GROUP_ROOMS_MEDIA_ROUTED                               UsageCategory = "group-rooms-media-routed"
	USAGECATEGORY_GROUP_ROOMS_MEDIA_STORED                               UsageCategory = "group-rooms-media-stored"
	USAGECATEGORY_GROUP_ROOMS_PARTICIPANT_MINUTES                        UsageCategory = "group-rooms-participant-minutes"
	USAGECATEGORY_GROUP_ROOMS_RECORDED_MINUTES                           UsageCategory = "group-rooms-recorded-minutes"
	USAGECATEGORY_IMP_V1_USAGE                                           UsageCategory = "imp-v1-usage"
	USAGECATEGORY_LOOKUPS                                                UsageCategory = "lookups"
	USAGECATEGORY_MARKETPLACE                                            UsageCategory = "marketplace"
	USAGECATEGORY_MARKETPLACE_ALGORITHMIA_NAMED_ENTITY_RECOGNITION       UsageCategory = "marketplace-algorithmia-named-entity-recognition"
	USAGECATEGORY_MARKETPLACE_CADENCE_TRANSCRIPTION                      UsageCategory = "marketplace-cadence-transcription"
	USAGECATEGORY_MARKETPLACE_CADENCE_TRANSLATION                        UsageCategory = "marketplace-cadence-translation"
	USAGECATEGORY_MARKETPLACE_CAPIO_SPEECH_TO_TEXT                       UsageCategory = "marketplace-capio-speech-to-text"
	USAGECATEGORY_MARKETPLACE_CONVRIZA_ABABA                             UsageCategory = "marketplace-convriza-ababa"
	USAGECATEGORY_MARKETPLACE_DEEPGRAM_PHRASE_DETECTOR                   UsageCategory = "marketplace-deepgram-phrase-detector"
	USAGECATEGORY_MARKETPLACE_DIGITAL_SEGMENT_BUSINESS_INFO              UsageCategory = "marketplace-digital-segment-business-info"
	USAGECATEGORY_MARKETPLACE_FACEBOOK_OFFLINE_CONVERSIONS               UsageCategory = "marketplace-facebook-offline-conversions"
	USAGECATEGORY_MARKETPLACE_GOOGLE_SPEECH_TO_TEXT                      UsageCategory = "marketplace-google-speech-to-text"
	USAGECATEGORY_MARKETPLACE_IBM_WATSON_MESSAGE_INSIGHTS                UsageCategory = "marketplace-ibm-watson-message-insights"
	USAGECATEGORY_MARKETPLACE_IBM_WATSON_MESSAGE_SENTIMENT               UsageCategory = "marketplace-ibm-watson-message-sentiment"
	USAGECATEGORY_MARKETPLACE_IBM_WATSON_RECORDING_ANALYSIS              UsageCategory = "marketplace-ibm-watson-recording-analysis"
	USAGECATEGORY_MARKETPLACE_IBM_WATSON_TONE_ANALYZER                   UsageCategory = "marketplace-ibm-watson-tone-analyzer"
	USAGECATEGORY_MARKETPLACE_ICEHOOK_SYSTEMS_SCOUT                      UsageCategory = "marketplace-icehook-systems-scout"
	USAGECATEGORY_MARKETPLACE_INFOGROUP_DATAAXLE_BIZINFO                 UsageCategory = "marketplace-infogroup-dataaxle-bizinfo"
	USAGECATEGORY_MARKETPLACE_KEEN_IO_CONTACT_CENTER_ANALYTICS           UsageCategory = "marketplace-keen-io-contact-center-analytics"
	USAGECATEGORY_MARKETPLACE_MARCHEX_CLEANCALL                          UsageCategory = "marketplace-marchex-cleancall"
	USAGECATEGORY_MARKETPLACE_MARCHEX_SENTIMENT_ANALYSIS_FOR_SMS         UsageCategory = "marketplace-marchex-sentiment-analysis-for-sms"
	USAGECATEGORY_MARKETPLACE_MARKETPLACE_NEXTCALLER_SOCIAL_ID           UsageCategory = "marketplace-marketplace-nextcaller-social-id"
	USAGECATEGORY_MARKETPLACE_MOBILE_COMMONS_OPT_OUT_CLASSIFIER          UsageCategory = "marketplace-mobile-commons-opt-out-classifier"
	USAGECATEGORY_MARKETPLACE_NEXIWAVE_VOICEMAIL_TO_TEXT                 UsageCategory = "marketplace-nexiwave-voicemail-to-text"
	USAGECATEGORY_MARKETPLACE_NEXTCALLER_ADVANCED_CALLER_IDENTIFICATION  UsageCategory = "marketplace-nextcaller-advanced-caller-identification"
	USAGECATEGORY_MARKETPLACE_NOMOROBO_SPAM_SCORE                        UsageCategory = "marketplace-nomorobo-spam-score"
	USAGECATEGORY_MARKETPLACE_PAYFONE_TCPA_COMPLIANCE                    UsageCategory = "marketplace-payfone-tcpa-compliance"
	USAGECATEGORY_MARKETPLACE_REMEETING_AUTOMATIC_SPEECH_RECOGNITION     UsageCategory = "marketplace-remeeting-automatic-speech-recognition"
	USAGECATEGORY_MARKETPLACE_TCPA_DEFENSE_SOLUTIONS_BLACKLIST_FEED      UsageCategory = "marketplace-tcpa-defense-solutions-blacklist-feed"
	USAGECATEGORY_MARKETPLACE_TELO_OPENCNAM                              UsageCategory = "marketplace-telo-opencnam"
	USAGECATEGORY_MARKETPLACE_TRUECNAM_TRUE_SPAM                         UsageCategory = "marketplace-truecnam-true-spam"
	USAGECATEGORY_MARKETPLACE_TWILIO_CALLER_NAME_LOOKUP_US               UsageCategory = "marketplace-twilio-caller-name-lookup-us"
	USAGECATEGORY_MARKETPLACE_TWILIO_CARRIER_INFORMATION_LOOKUP          UsageCategory = "marketplace-twilio-carrier-information-lookup"
	USAGECATEGORY_MARKETPLACE_VOICEBASE_PCI                              UsageCategory = "marketplace-voicebase-pci"
	USAGECATEGORY_MARKETPLACE_VOICEBASE_TRANSCRIPTION                    UsageCategory = "marketplace-voicebase-transcription"
	USAGECATEGORY_MARKETPLACE_VOICEBASE_TRANSCRIPTION_CUSTOM_VOCABULARY  UsageCategory = "marketplace-voicebase-transcription-custom-vocabulary"
	USAGECATEGORY_MARKETPLACE_WHITEPAGES_PRO_CALLER_IDENTIFICATION       UsageCategory = "marketplace-whitepages-pro-caller-identification"
	USAGECATEGORY_MARKETPLACE_WHITEPAGES_PRO_PHONE_INTELLIGENCE          UsageCategory = "marketplace-whitepages-pro-phone-intelligence"
	USAGECATEGORY_MARKETPLACE_WHITEPAGES_PRO_PHONE_REPUTATION            UsageCategory = "marketplace-whitepages-pro-phone-reputation"
	USAGECATEGORY_MARKETPLACE_WOLFARM_SPOKEN_RESULTS                     UsageCategory = "marketplace-wolfarm-spoken-results"
	USAGECATEGORY_MARKETPLACE_WOLFRAM_SHORT_ANSWER                       UsageCategory = "marketplace-wolfram-short-answer"
	USAGECATEGORY_MARKETPLACE_YTICA_CONTACT_CENTER_REPORTING_ANALYTICS   UsageCategory = "marketplace-ytica-contact-center-reporting-analytics"
	USAGECATEGORY_MEDIASTORAGE                                           UsageCategory = "mediastorage"
	USAGECATEGORY_MMS                                                    UsageCategory = "mms"
	USAGECATEGORY_MMS_INBOUND                                            UsageCategory = "mms-inbound"
	USAGECATEGORY_MMS_INBOUND_LONGCODE                                   UsageCategory = "mms-inbound-longcode"
	USAGECATEGORY_MMS_INBOUND_SHORTCODE                                  UsageCategory = "mms-inbound-shortcode"
	USAGECATEGORY_MMS_MESSAGES_CARRIERFEES                               UsageCategory = "mms-messages-carrierfees"
	USAGECATEGORY_MMS_OUTBOUND                                           UsageCategory = "mms-outbound"
	USAGECATEGORY_MMS_OUTBOUND_LONGCODE                                  UsageCategory = "mms-outbound-longcode"
	USAGECATEGORY_MMS_OUTBOUND_SHORTCODE                                 UsageCategory = "mms-outbound-shortcode"
	USAGECATEGORY_MONITOR_READS                                          UsageCategory = "monitor-reads"
	USAGECATEGORY_MONITOR_STORAGE                                        UsageCategory = "monitor-storage"
	USAGECATEGORY_MONITOR_WRITES                                         UsageCategory = "monitor-writes"
	USAGECATEGORY_NOTIFY                                                 UsageCategory = "notify"
	USAGECATEGORY_NOTIFY_ACTIONS_ATTEMPTS                                UsageCategory = "notify-actions-attempts"
	USAGECATEGORY_NOTIFY_CHANNELS                                        UsageCategory = "notify-channels"
	USAGECATEGORY_NUMBER_FORMAT_LOOKUPS                                  UsageCategory = "number-format-lookups"
	USAGECATEGORY_PCHAT                                                  UsageCategory = "pchat"
	USAGECATEGORY_PCHAT_USERS                                            UsageCategory = "pchat-users"
	USAGECATEGORY_PEER_TO_PEER_ROOMS_PARTICIPANT_MINUTES                 UsageCategory = "peer-to-peer-rooms-participant-minutes"
	USAGECATEGORY_PFAX                                                   UsageCategory = "pfax"
	USAGECATEGORY_PFAX_MINUTES                                           UsageCategory = "pfax-minutes"
	USAGECATEGORY_PFAX_MINUTES_INBOUND                                   UsageCategory = "pfax-minutes-inbound"
	USAGECATEGORY_PFAX_MINUTES_OUTBOUND                                  UsageCategory = "pfax-minutes-outbound"
	USAGECATEGORY_PFAX_PAGES                                             UsageCategory = "pfax-pages"
	USAGECATEGORY_PHONENUMBERS                                           UsageCategory = "phonenumbers"
	USAGECATEGORY_PHONENUMBERS_CPS                                       UsageCategory = "phonenumbers-cps"
	USAGECATEGORY_PHONENUMBERS_EMERGENCY                                 UsageCategory = "phonenumbers-emergency"
	USAGECATEGORY_PHONENUMBERS_LOCAL                                     UsageCategory = "phonenumbers-local"
	USAGECATEGORY_PHONENUMBERS_MOBILE                                    UsageCategory = "phonenumbers-mobile"
	USAGECATEGORY_PHONENUMBERS_SETUPS                                    UsageCategory = "phonenumbers-setups"
	USAGECATEGORY_PHONENUMBERS_TOLLFREE                                  UsageCategory = "phonenumbers-tollfree"
	USAGECATEGORY_PREMIUMSUPPORT                                         UsageCategory = "premiumsupport"
	USAGECATEGORY_PROXY                                                  UsageCategory = "proxy"
	USAGECATEGORY_PROXY_ACTIVE_SESSIONS                                  UsageCategory = "proxy-active-sessions"
	USAGECATEGORY_PSTNCONNECTIVITY                                       UsageCategory = "pstnconnectivity"
	USAGECATEGORY_PV                                                     UsageCategory = "pv"
	USAGECATEGORY_PV_COMPOSITION_MEDIA_DOWNLOADED                        UsageCategory = "pv-composition-media-downloaded"
	USAGECATEGORY_PV_COMPOSITION_MEDIA_ENCRYPTED                         UsageCategory = "pv-composition-media-encrypted"
	USAGECATEGORY_PV_COMPOSITION_MEDIA_STORED                            UsageCategory = "pv-composition-media-stored"
	USAGECATEGORY_PV_COMPOSITION_MINUTES                                 UsageCategory = "pv-composition-minutes"
	USAGECATEGORY_PV_RECORDING_COMPOSITIONS                              UsageCategory = "pv-recording-compositions"
	USAGECATEGORY_PV_ROOM_PARTICIPANTS                                   UsageCategory = "pv-room-participants"
	USAGECATEGORY_PV_ROOM_PARTICIPANTS_AU1                               UsageCategory = "pv-room-participants-au1"
	USAGECATEGORY_PV_ROOM_PARTICIPANTS_BR1                               UsageCategory = "pv-room-participants-br1"
	USAGECATEGORY_PV_ROOM_PARTICIPANTS_IE1                               UsageCategory = "pv-room-participants-ie1"
	USAGECATEGORY_PV_ROOM_PARTICIPANTS_JP1                               UsageCategory = "pv-room-participants-jp1"
	USAGECATEGORY_PV_ROOM_PARTICIPANTS_SG1                               UsageCategory = "pv-room-participants-sg1"
	USAGECATEGORY_PV_ROOM_PARTICIPANTS_US1                               UsageCategory = "pv-room-participants-us1"
	USAGECATEGORY_PV_ROOM_PARTICIPANTS_US2                               UsageCategory = "pv-room-participants-us2"
	USAGECATEGORY_PV_ROOMS                                               UsageCategory = "pv-rooms"
	USAGECATEGORY_PV_SIP_ENDPOINT_REGISTRATIONS                          UsageCategory = "pv-sip-endpoint-registrations"
	USAGECATEGORY_RECORDINGS                                             UsageCategory = "recordings"
	USAGECATEGORY_RECORDINGSTORAGE                                       UsageCategory = "recordingstorage"
	USAGECATEGORY_ROOMS_GROUP_BANDWIDTH                                  UsageCategory = "rooms-group-bandwidth"
	USAGECATEGORY_ROOMS_GROUP_MINUTES                                    UsageCategory = "rooms-group-minutes"
	USAGECATEGORY_ROOMS_PEER_TO_PEER_MINUTES                             UsageCategory = "rooms-peer-to-peer-minutes"
	USAGECATEGORY_SHORTCODES                                             UsageCategory = "shortcodes"
	USAGECATEGORY_SHORTCODES_CUSTOMEROWNED                               UsageCategory = "shortcodes-customerowned"
	USAGECATEGORY_SHORTCODES_MMS_ENABLEMENT                              UsageCategory = "shortcodes-mms-enablement"
	USAGECATEGORY_SHORTCODES_MPS                                         UsageCategory = "shortcodes-mps"
	USAGECATEGORY_SHORTCODES_RANDOM                                      UsageCategory = "shortcodes-random"
	USAGECATEGORY_SHORTCODES_UK                                          UsageCategory = "shortcodes-uk"
	USAGECATEGORY_SHORTCODES_VANITY                                      UsageCategory = "shortcodes-vanity"
	USAGECATEGORY_SMALL_GROUP_ROOMS                                      UsageCategory = "small-group-rooms"
	USAGECATEGORY_SMALL_GROUP_ROOMS_DATA_TRACK                           UsageCategory = "small-group-rooms-data-track"
	USAGECATEGORY_SMALL_GROUP_ROOMS_PARTICIPANT_MINUTES                  UsageCategory = "small-group-rooms-participant-minutes"
	USAGECATEGORY_SMS                                                    UsageCategory = "sms"
	USAGECATEGORY_SMS_INBOUND                                            UsageCategory = "sms-inbound"
	USAGECATEGORY_SMS_INBOUND_LONGCODE                                   UsageCategory = "sms-inbound-longcode"
	USAGECATEGORY_SMS_INBOUND_SHORTCODE                                  UsageCategory = "sms-inbound-shortcode"
	USAGECATEGORY_SMS_MESSAGES_CARRIERFEES                               UsageCategory = "sms-messages-carrierfees"
	USAGECATEGORY_SMS_MESSAGES_FEATURES                                  UsageCategory = "sms-messages-features"
	USAGECATEGORY_SMS_MESSAGES_FEATURES_SENDERID                         UsageCategory = "sms-messages-features-senderid"
	USAGECATEGORY_SMS_OUTBOUND                                           UsageCategory = "sms-outbound"
	USAGECATEGORY_SMS_OUTBOUND_CONTENT_INSPECTION                        UsageCategory = "sms-outbound-content-inspection"
	USAGECATEGORY_SMS_OUTBOUND_LONGCODE                                  UsageCategory = "sms-outbound-longcode"
	USAGECATEGORY_SMS_OUTBOUND_SHORTCODE                                 UsageCategory = "sms-outbound-shortcode"
	USAGECATEGORY_SPEECH_RECOGNITION                                     UsageCategory = "speech-recognition"
	USAGECATEGORY_STUDIO_ENGAGEMENTS                                     UsageCategory = "studio-engagements"
	USAGECATEGORY_SYNC                                                   UsageCategory = "sync"
	USAGECATEGORY_SYNC_ACTIONS                                           UsageCategory = "sync-actions"
	USAGECATEGORY_SYNC_ENDPOINT_HOURS                                    UsageCategory = "sync-endpoint-hours"
	USAGECATEGORY_SYNC_ENDPOINT_HOURS_ABOVE_DAILY_CAP                    UsageCategory = "sync-endpoint-hours-above-daily-cap"
	USAGECATEGORY_TASKROUTER_TASKS                                       UsageCategory = "taskrouter-tasks"
	USAGECATEGORY_TOTALPRICE                                             UsageCategory = "totalprice"
	USAGECATEGORY_TRANSCRIPTIONS                                         UsageCategory = "transcriptions"
	USAGECATEGORY_TRUNKING_CPS                                           UsageCategory = "trunking-cps"
	USAGECATEGORY_TRUNKING_EMERGENCY_CALLS                               UsageCategory = "trunking-emergency-calls"
	USAGECATEGORY_TRUNKING_ORIGINATION                                   UsageCategory = "trunking-origination"
	USAGECATEGORY_TRUNKING_ORIGINATION_LOCAL                             UsageCategory = "trunking-origination-local"
	USAGECATEGORY_TRUNKING_ORIGINATION_MOBILE                            UsageCategory = "trunking-origination-mobile"
	USAGECATEGORY_TRUNKING_ORIGINATION_TOLLFREE                          UsageCategory = "trunking-origination-tollfree"
	USAGECATEGORY_TRUNKING_RECORDINGS                                    UsageCategory = "trunking-recordings"
	USAGECATEGORY_TRUNKING_SECURE                                        UsageCategory = "trunking-secure"
	USAGECATEGORY_TRUNKING_TERMINATION                                   UsageCategory = "trunking-termination"
	USAGECATEGORY_TURNMEGABYTES                                          UsageCategory = "turnmegabytes"
	USAGECATEGORY_TURNMEGABYTES_AUSTRALIA                                UsageCategory = "turnmegabytes-australia"
	USAGECATEGORY_TURNMEGABYTES_BRASIL                                   UsageCategory = "turnmegabytes-brasil"
	USAGECATEGORY_TURNMEGABYTES_GERMANY                                  UsageCategory = "turnmegabytes-germany"
	USAGECATEGORY_TURNMEGABYTES_INDIA                                    UsageCategory = "turnmegabytes-india"
	USAGECATEGORY_TURNMEGABYTES_IRELAND                                  UsageCategory = "turnmegabytes-ireland"
	USAGECATEGORY_TURNMEGABYTES_JAPAN                                    UsageCategory = "turnmegabytes-japan"
	USAGECATEGORY_TURNMEGABYTES_SINGAPORE                                UsageCategory = "turnmegabytes-singapore"
	USAGECATEGORY_TURNMEGABYTES_USEAST                                   UsageCategory = "turnmegabytes-useast"
	USAGECATEGORY_TURNMEGABYTES_USWEST                                   UsageCategory = "turnmegabytes-uswest"
	USAGECATEGORY_TWILIO_INTERCONNECT                                    UsageCategory = "twilio-interconnect"
	USAGECATEGORY_VERIFY_PUSH                                            UsageCategory = "verify-push"
	USAGECATEGORY_VIDEO_RECORDINGS                                       UsageCategory = "video-recordings"
	USAGECATEGORY_VOICE_INSIGHTS                                         UsageCategory = "voice-insights"
	USAGECATEGORY_VOICE_INSIGHTS_CLIENT_INSIGHTS_ON_DEMAND_MINUTE        UsageCategory = "voice-insights-client-insights-on-demand-minute"
	USAGECATEGORY_VOICE_INSIGHTS_PTSN_INSIGHTS_ON_DEMAND_MINUTE          UsageCategory = "voice-insights-ptsn-insights-on-demand-minute"
	USAGECATEGORY_VOICE_INSIGHTS_SIP_INTERFACE_INSIGHTS_ON_DEMAND_MINUTE UsageCategory = "voice-insights-sip-interface-insights-on-demand-minute"
	USAGECATEGORY_VOICE_INSIGHTS_SIP_TRUNKING_INSIGHTS_ON_DEMAND_MINUTE  UsageCategory = "voice-insights-sip-trunking-insights-on-demand-minute"
	USAGECATEGORY_WIRELESS                                               UsageCategory = "wireless"
	USAGECATEGORY_WIRELESS_ORDERS                                        UsageCategory = "wireless-orders"
	USAGECATEGORY_WIRELESS_ORDERS_ARTWORK                                UsageCategory = "wireless-orders-artwork"
	USAGECATEGORY_WIRELESS_ORDERS_BULK                                   UsageCategory = "wireless-orders-bulk"
	USAGECATEGORY_WIRELESS_ORDERS_ESIM                                   UsageCategory = "wireless-orders-esim"
	USAGECATEGORY_WIRELESS_ORDERS_STARTER                                UsageCategory = "wireless-orders-starter"
	USAGECATEGORY_WIRELESS_USAGE                                         UsageCategory = "wireless-usage"
	USAGECATEGORY_WIRELESS_USAGE_COMMANDS                                UsageCategory = "wireless-usage-commands"
	USAGECATEGORY_WIRELESS_USAGE_COMMANDS_AFRICA                         UsageCategory = "wireless-usage-commands-africa"
	USAGECATEGORY_WIRELESS_USAGE_COMMANDS_ASIA                           UsageCategory = "wireless-usage-commands-asia"
	USAGECATEGORY_WIRELESS_USAGE_COMMANDS_CENTRALANDSOUTHAMERICA         UsageCategory = "wireless-usage-commands-centralandsouthamerica"
	USAGECATEGORY_WIRELESS_USAGE_COMMANDS_EUROPE                         UsageCategory = "wireless-usage-commands-europe"
	USAGECATEGORY_WIRELESS_USAGE_COMMANDS_HOME                           UsageCategory = "wireless-usage-commands-home"
	USAGECATEGORY_WIRELESS_USAGE_COMMANDS_NORTHAMERICA                   UsageCategory = "wireless-usage-commands-northamerica"
	USAGECATEGORY_WIRELESS_USAGE_COMMANDS_OCEANIA                        UsageCategory = "wireless-usage-commands-oceania"
	USAGECATEGORY_WIRELESS_USAGE_COMMANDS_ROAMING                        UsageCategory = "wireless-usage-commands-roaming"
	USAGECATEGORY_WIRELESS_USAGE_DATA                                    UsageCategory = "wireless-usage-data"
	USAGECATEGORY_WIRELESS_USAGE_DATA_AFRICA                             UsageCategory = "wireless-usage-data-africa"
	USAGECATEGORY_WIRELESS_USAGE_DATA_ASIA                               UsageCategory = "wireless-usage-data-asia"
	USAGECATEGORY_WIRELESS_USAGE_DATA_CENTRALANDSOUTHAMERICA             UsageCategory = "wireless-usage-data-centralandsouthamerica"
	USAGECATEGORY_WIRELESS_USAGE_DATA_CUSTOM_ADDITIONALMB                UsageCategory = "wireless-usage-data-custom-additionalmb"
	USAGECATEGORY_WIRELESS_USAGE_DATA_CUSTOM_FIRST5MB                    UsageCategory = "wireless-usage-data-custom-first5mb"
	USAGECATEGORY_WIRELESS_USAGE_DATA_DOMESTIC_ROAMING                   UsageCategory = "wireless-usage-data-domestic-roaming"
	USAGECATEGORY_WIRELESS_USAGE_DATA_EUROPE                             UsageCategory = "wireless-usage-data-europe"
	USAGECATEGORY_WIRELESS_USAGE_DATA_INDIVIDUAL_ADDITIONALGB            UsageCategory = "wireless-usage-data-individual-additionalgb"
	USAGECATEGORY_WIRELESS_USAGE_DATA_INDIVIDUAL_FIRSTGB                 UsageCategory = "wireless-usage-data-individual-firstgb"
	USAGECATEGORY_WIRELESS_USAGE_DATA_INTERNATIONAL_ROAMING_CANADA       UsageCategory = "wireless-usage-data-international-roaming-canada"
	USAGECATEGORY_WIRELESS_USAGE_DATA_INTERNATIONAL_ROAMING_INDIA        UsageCategory = "wireless-usage-data-international-roaming-india"
	USAGECATEGORY_WIRELESS_USAGE_DATA_INTERNATIONAL_ROAMING_MEXICO       UsageCategory = "wireless-usage-data-international-roaming-mexico"
	USAGECATEGORY_WIRELESS_USAGE_DATA_NORTHAMERICA                       UsageCategory = "wireless-usage-data-northamerica"
	USAGECATEGORY_WIRELESS_USAGE_DATA_OCEANIA                            UsageCategory = "wireless-usage-data-oceania"
	USAGECATEGORY_WIRELESS_USAGE_DATA_POOLED                             UsageCategory = "wireless-usage-data-pooled"
	USAGECATEGORY_WIRELESS_USAGE_DATA_POOLED_DOWNLINK                    UsageCategory = "wireless-usage-data-pooled-downlink"
	USAGECATEGORY_WIRELESS_USAGE_DATA_POOLED_UPLINK                      UsageCategory = "wireless-usage-data-pooled-uplink"
	USAGECATEGORY_WIRELESS_USAGE_MRC                                     UsageCategory = "wireless-usage-mrc"
	USAGECATEGORY_WIRELESS_USAGE_MRC_CUSTOM                              UsageCategory = "wireless-usage-mrc-custom"
	USAGECATEGORY_WIRELESS_USAGE_MRC_INDIVIDUAL                          UsageCategory = "wireless-usage-mrc-individual"
	USAGECATEGORY_WIRELESS_USAGE_MRC_POOLED                              UsageCategory = "wireless-usage-mrc-pooled"
	USAGECATEGORY_WIRELESS_USAGE_MRC_SUSPENDED                           UsageCategory = "wireless-usage-mrc-suspended"
	USAGECATEGORY_WIRELESS_USAGE_SMS                                     UsageCategory = "wireless-usage-sms"
	USAGECATEGORY_WIRELESS_USAGE_VOICE                                   UsageCategory = "wireless-usage-voice"
)
