/*
 * Twilio - Fax
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.15.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
	"time"

	twilio "github.com/twilio/twilio-go/client"
)

type DefaultApiService struct {
	baseURL string
	client  twilio.BaseClient
}

func NewDefaultApiService(client twilio.BaseClient) *DefaultApiService {
	return &DefaultApiService{
		client:  client,
		baseURL: "https://fax.twilio.com",
	}
}

// CreateFaxParams Optional parameters for the method 'CreateFax'
type CreateFaxParams struct {
	From            *string `json:"From,omitempty"`
	MediaUrl        *string `json:"MediaUrl,omitempty"`
	Quality         *string `json:"Quality,omitempty"`
	SipAuthPassword *string `json:"SipAuthPassword,omitempty"`
	SipAuthUsername *string `json:"SipAuthUsername,omitempty"`
	StatusCallback  *string `json:"StatusCallback,omitempty"`
	StoreMedia      *bool   `json:"StoreMedia,omitempty"`
	To              *string `json:"To,omitempty"`
	Ttl             *int32  `json:"Ttl,omitempty"`
}

// CreateFax Method for CreateFax
//
// Create a new fax to send to a phone number or SIP endpoint.
//
// param: optional nil or *CreateFaxParams - Optional Parameters:
//
// param: "From" (string) - The number the fax was sent from. Can be the phone number in [E.164](https://www.twilio.com/docs/glossary/what-e164) format or the SIP `from` value. The caller ID displayed to the recipient uses this value. If this is a phone number, it must be a Twilio number or a verified outgoing caller id from your account. If `to` is a SIP address, this can be any alphanumeric string (and also the characters `+`, `_`, `.`, and `-`), which will be used in the `from` header of the SIP request.
//
// param: "MediaUrl" (string) - The URL of the PDF that contains the fax. See our [security](https://www.twilio.com/docs/usage/security) page for information on how to ensure the request for your media comes from Twilio.
//
// param: "Quality" (string) - The [Fax Quality value](https://www.twilio.com/docs/fax/api/fax-resource#fax-quality-values) that describes the fax quality. Can be: `standard`, `fine`, or `superfine` and defaults to `fine`.
//
// param: "SipAuthPassword" (string) - The password to use with `sip_auth_username` to authenticate faxes sent to a SIP address.
//
// param: "SipAuthUsername" (string) - The username to use with the `sip_auth_password` to authenticate faxes sent to a SIP address.
//
// param: "StatusCallback" (string) - The URL we should call using the `POST` method to send [status information](https://www.twilio.com/docs/fax/api/fax-resource#fax-status-callback) to your application when the status of the fax changes.
//
// param: "StoreMedia" (bool) - Whether to store a copy of the sent media on our servers for later retrieval. Can be: `true` or `false` and the default is `true`.
//
// param: "To" (string) - The phone number to receive the fax in [E.164](https://www.twilio.com/docs/glossary/what-e164) format or the recipient's SIP URI.
//
// param: "Ttl" (int32) - How long in minutes from when the fax is initiated that we should try to send the fax.
//
// return: FaxV1Fax
func (c *DefaultApiService) CreateFax(params *CreateFaxParams) (*FaxV1Fax, error) {
	path := "/v1/Faxes"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.From != nil {
		data.Set("From", *params.From)
	}
	if params != nil && params.MediaUrl != nil {
		data.Set("MediaUrl", *params.MediaUrl)
	}
	if params != nil && params.Quality != nil {
		data.Set("Quality", *params.Quality)
	}
	if params != nil && params.SipAuthPassword != nil {
		data.Set("SipAuthPassword", *params.SipAuthPassword)
	}
	if params != nil && params.SipAuthUsername != nil {
		data.Set("SipAuthUsername", *params.SipAuthUsername)
	}
	if params != nil && params.StatusCallback != nil {
		data.Set("StatusCallback", *params.StatusCallback)
	}
	if params != nil && params.StoreMedia != nil {
		data.Set("StoreMedia", fmt.Sprint(*params.StoreMedia))
	}
	if params != nil && params.To != nil {
		data.Set("To", *params.To)
	}
	if params != nil && params.Ttl != nil {
		data.Set("Ttl", fmt.Sprint(*params.Ttl))
	}

	resp, err := c.client.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &FaxV1Fax{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// DeleteFax Method for DeleteFax
//
// Delete a specific fax and its associated media.
//
// param: Sid The Twilio-provided string that uniquely identifies the Fax resource to delete.
//
func (c *DefaultApiService) DeleteFax(Sid string) error {
	path := "/v1/Faxes/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.client.Delete(c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

// DeleteFaxMedia Method for DeleteFaxMedia
//
// Delete a specific fax media instance.
//
// param: FaxSid The SID of the fax with the FaxMedia resource to delete.
//
// param: Sid The Twilio-provided string that uniquely identifies the FaxMedia resource to delete.
//
func (c *DefaultApiService) DeleteFaxMedia(FaxSid string, Sid string) error {
	path := "/v1/Faxes/{FaxSid}/Media/{Sid}"
	path = strings.Replace(path, "{"+"FaxSid"+"}", FaxSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.client.Delete(c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

// FetchFax Method for FetchFax
//
// Fetch a specific fax.
//
// param: Sid The Twilio-provided string that uniquely identifies the Fax resource to fetch.
//
// return: FaxV1Fax
func (c *DefaultApiService) FetchFax(Sid string) (*FaxV1Fax, error) {
	path := "/v1/Faxes/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.client.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &FaxV1Fax{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// FetchFaxMedia Method for FetchFaxMedia
//
// Fetch a specific fax media instance.
//
// param: FaxSid The SID of the fax with the FaxMedia resource to fetch.
//
// param: Sid The Twilio-provided string that uniquely identifies the FaxMedia resource to fetch.
//
// return: FaxV1FaxFaxMedia
func (c *DefaultApiService) FetchFaxMedia(FaxSid string, Sid string) (*FaxV1FaxFaxMedia, error) {
	path := "/v1/Faxes/{FaxSid}/Media/{Sid}"
	path = strings.Replace(path, "{"+"FaxSid"+"}", FaxSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.client.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &FaxV1FaxFaxMedia{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// ListFaxParams Optional parameters for the method 'ListFax'
type ListFaxParams struct {
	From                  *string    `json:"From,omitempty"`
	To                    *string    `json:"To,omitempty"`
	DateCreatedOnOrBefore *time.Time `json:"DateCreatedOnOrBefore,omitempty"`
	DateCreatedAfter      *time.Time `json:"DateCreatedAfter,omitempty"`
	PageSize              *int32     `json:"PageSize,omitempty"`
}

// ListFax Method for ListFax
//
// Retrieve a list of all faxes.
//
// param: optional nil or *ListFaxParams - Optional Parameters:
//
// param: "From" (string) - Retrieve only those faxes sent from this phone number, specified in [E.164](https://www.twilio.com/docs/glossary/what-e164) format.
//
// param: "To" (string) - Retrieve only those faxes sent to this phone number, specified in [E.164](https://www.twilio.com/docs/glossary/what-e164) format.
//
// param: "DateCreatedOnOrBefore" (time.Time) - Retrieve only those faxes with a `date_created` that is before or equal to this value, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
//
// param: "DateCreatedAfter" (time.Time) - Retrieve only those faxes with a `date_created` that is later than this value, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
//
// param: "PageSize" (int32) - How many resources to return in each list page. The default is 50, and the maximum is 1000.
//
// return: ListFaxResponse
func (c *DefaultApiService) ListFax(params *ListFaxParams) (*ListFaxResponse, error) {
	path := "/v1/Faxes"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.From != nil {
		data.Set("From", *params.From)
	}
	if params != nil && params.To != nil {
		data.Set("To", *params.To)
	}
	if params != nil && params.DateCreatedOnOrBefore != nil {
		data.Set("DateCreatedOnOrBefore", fmt.Sprint((*params.DateCreatedOnOrBefore).Format(time.RFC3339)))
	}
	if params != nil && params.DateCreatedAfter != nil {
		data.Set("DateCreatedAfter", fmt.Sprint((*params.DateCreatedAfter).Format(time.RFC3339)))
	}
	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	resp, err := c.client.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListFaxResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// ListFaxMediaParams Optional parameters for the method 'ListFaxMedia'
type ListFaxMediaParams struct {
	PageSize *int32 `json:"PageSize,omitempty"`
}

// ListFaxMedia Method for ListFaxMedia
//
// Retrieve a list of all fax media instances for the specified fax.
//
// param: FaxSid The SID of the fax with the FaxMedia resources to read.
//
// param: optional nil or *ListFaxMediaParams - Optional Parameters:
//
// param: "PageSize" (int32) - How many resources to return in each list page. The default is 50, and the maximum is 1000.
//
// return: ListFaxMediaResponse
func (c *DefaultApiService) ListFaxMedia(FaxSid string, params *ListFaxMediaParams) (*ListFaxMediaResponse, error) {
	path := "/v1/Faxes/{FaxSid}/Media"
	path = strings.Replace(path, "{"+"FaxSid"+"}", FaxSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	resp, err := c.client.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListFaxMediaResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// UpdateFaxParams Optional parameters for the method 'UpdateFax'
type UpdateFaxParams struct {
	Status *string `json:"Status,omitempty"`
}

// UpdateFax Method for UpdateFax
//
// Update a specific fax.
//
// param: Sid The Twilio-provided string that uniquely identifies the Fax resource to update.
//
// param: optional nil or *UpdateFaxParams - Optional Parameters:
//
// param: "Status" (string) - The new [status](https://www.twilio.com/docs/fax/api/fax-resource#fax-status-values) of the resource. Can be only `canceled`. This may fail if transmission has already started.
//
// return: FaxV1Fax
func (c *DefaultApiService) UpdateFax(Sid string, params *UpdateFaxParams) (*FaxV1Fax, error) {
	path := "/v1/Faxes/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Status != nil {
		data.Set("Status", *params.Status)
	}

	resp, err := c.client.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &FaxV1Fax{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
