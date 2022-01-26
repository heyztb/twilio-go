/*
 * Twilio - Proxy
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.25.1
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
	"net/url"

	"strings"

	"github.com/twilio/twilio-go/client"
)

// Delete a specific Interaction.
func (c *ApiService) DeleteInteraction(ServiceSid string, SessionSid string, Sid string) error {
	path := "/v1/Services/{ServiceSid}/Sessions/{SessionSid}/Interactions/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"SessionSid"+"}", SessionSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Delete(c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

// Retrieve a list of Interactions for a given [Session](https://www.twilio.com/docs/proxy/api/session).
func (c *ApiService) FetchInteraction(ServiceSid string, SessionSid string, Sid string) (*ProxyV1Interaction, error) {
	path := "/v1/Services/{ServiceSid}/Sessions/{SessionSid}/Interactions/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"SessionSid"+"}", SessionSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ProxyV1Interaction{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListInteraction'
type ListInteractionParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListInteractionParams) SetPageSize(PageSize int) *ListInteractionParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListInteractionParams) SetLimit(Limit int) *ListInteractionParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of Interaction records from the API. Request is executed immediately.
func (c *ApiService) PageInteraction(ServiceSid string, SessionSid string, params *ListInteractionParams, pageToken, pageNumber string) (*ListInteractionResponse, error) {
	path := "/v1/Services/{ServiceSid}/Sessions/{SessionSid}/Interactions"

	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"SessionSid"+"}", SessionSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	if pageToken != "" {
		data.Set("PageToken", pageToken)
	}
	if pageNumber != "" {
		data.Set("Page", pageNumber)
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListInteractionResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists Interaction records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListInteraction(ServiceSid string, SessionSid string, params *ListInteractionParams) ([]ProxyV1Interaction, error) {
	if params == nil {
		params = &ListInteractionParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageInteraction(ServiceSid, SessionSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []ProxyV1Interaction

	for response != nil {
		records = append(records, response.Interactions...)

		var record interface{}
		if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListInteractionResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListInteractionResponse)
	}

	return records, err
}

// Streams Interaction records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamInteraction(ServiceSid string, SessionSid string, params *ListInteractionParams) (chan ProxyV1Interaction, error) {
	if params == nil {
		params = &ListInteractionParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageInteraction(ServiceSid, SessionSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan ProxyV1Interaction, 1)

	go func() {
		for response != nil {
			for item := range response.Interactions {
				channel <- response.Interactions[item]
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListInteractionResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListInteractionResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListInteractionResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListInteractionResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
