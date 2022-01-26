/*
 * Twilio - Insights
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

// Optional parameters for the method 'FetchConferenceParticipant'
type FetchConferenceParticipantParams struct {
	//
	Events *string `json:"Events,omitempty"`
	//
	Metrics *string `json:"Metrics,omitempty"`
}

func (params *FetchConferenceParticipantParams) SetEvents(Events string) *FetchConferenceParticipantParams {
	params.Events = &Events
	return params
}
func (params *FetchConferenceParticipantParams) SetMetrics(Metrics string) *FetchConferenceParticipantParams {
	params.Metrics = &Metrics
	return params
}

func (c *ApiService) FetchConferenceParticipant(ConferenceSid string, ParticipantSid string, params *FetchConferenceParticipantParams) (*InsightsV1ConferenceParticipant, error) {
	path := "/v1/Conferences/{ConferenceSid}/Participants/{ParticipantSid}"
	path = strings.Replace(path, "{"+"ConferenceSid"+"}", ConferenceSid, -1)
	path = strings.Replace(path, "{"+"ParticipantSid"+"}", ParticipantSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Events != nil {
		data.Set("Events", *params.Events)
	}
	if params != nil && params.Metrics != nil {
		data.Set("Metrics", *params.Metrics)
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &InsightsV1ConferenceParticipant{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListConferenceParticipant'
type ListConferenceParticipantParams struct {
	//
	ParticipantSid *string `json:"ParticipantSid,omitempty"`
	//
	Label *string `json:"Label,omitempty"`
	//
	Events *string `json:"Events,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListConferenceParticipantParams) SetParticipantSid(ParticipantSid string) *ListConferenceParticipantParams {
	params.ParticipantSid = &ParticipantSid
	return params
}
func (params *ListConferenceParticipantParams) SetLabel(Label string) *ListConferenceParticipantParams {
	params.Label = &Label
	return params
}
func (params *ListConferenceParticipantParams) SetEvents(Events string) *ListConferenceParticipantParams {
	params.Events = &Events
	return params
}
func (params *ListConferenceParticipantParams) SetPageSize(PageSize int) *ListConferenceParticipantParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListConferenceParticipantParams) SetLimit(Limit int) *ListConferenceParticipantParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of ConferenceParticipant records from the API. Request is executed immediately.
func (c *ApiService) PageConferenceParticipant(ConferenceSid string, params *ListConferenceParticipantParams, pageToken, pageNumber string) (*ListConferenceParticipantResponse, error) {
	path := "/v1/Conferences/{ConferenceSid}/Participants"

	path = strings.Replace(path, "{"+"ConferenceSid"+"}", ConferenceSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.ParticipantSid != nil {
		data.Set("ParticipantSid", *params.ParticipantSid)
	}
	if params != nil && params.Label != nil {
		data.Set("Label", *params.Label)
	}
	if params != nil && params.Events != nil {
		data.Set("Events", *params.Events)
	}
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

	ps := &ListConferenceParticipantResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists ConferenceParticipant records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListConferenceParticipant(ConferenceSid string, params *ListConferenceParticipantParams) ([]InsightsV1ConferenceParticipant, error) {
	if params == nil {
		params = &ListConferenceParticipantParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageConferenceParticipant(ConferenceSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []InsightsV1ConferenceParticipant

	for response != nil {
		records = append(records, response.Participants...)

		var record interface{}
		if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListConferenceParticipantResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListConferenceParticipantResponse)
	}

	return records, err
}

// Streams ConferenceParticipant records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamConferenceParticipant(ConferenceSid string, params *ListConferenceParticipantParams) (chan InsightsV1ConferenceParticipant, error) {
	if params == nil {
		params = &ListConferenceParticipantParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageConferenceParticipant(ConferenceSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan InsightsV1ConferenceParticipant, 1)

	go func() {
		for response != nil {
			for item := range response.Participants {
				channel <- response.Participants[item]
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListConferenceParticipantResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListConferenceParticipantResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListConferenceParticipantResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListConferenceParticipantResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
