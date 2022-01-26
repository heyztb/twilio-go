/*
 * Twilio - Trunking
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

// Optional parameters for the method 'CreateIpAccessControlList'
type CreateIpAccessControlListParams struct {
	// The SID of the [IP Access Control List](https://www.twilio.com/docs/voice/sip/api/sip-ipaccesscontrollist-resource) that you want to associate with the trunk.
	IpAccessControlListSid *string `json:"IpAccessControlListSid,omitempty"`
}

func (params *CreateIpAccessControlListParams) SetIpAccessControlListSid(IpAccessControlListSid string) *CreateIpAccessControlListParams {
	params.IpAccessControlListSid = &IpAccessControlListSid
	return params
}

// Associate an IP Access Control List with a Trunk
func (c *ApiService) CreateIpAccessControlList(TrunkSid string, params *CreateIpAccessControlListParams) (*TrunkingV1IpAccessControlList, error) {
	path := "/v1/Trunks/{TrunkSid}/IpAccessControlLists"
	path = strings.Replace(path, "{"+"TrunkSid"+"}", TrunkSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.IpAccessControlListSid != nil {
		data.Set("IpAccessControlListSid", *params.IpAccessControlListSid)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TrunkingV1IpAccessControlList{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Remove an associated IP Access Control List from a Trunk
func (c *ApiService) DeleteIpAccessControlList(TrunkSid string, Sid string) error {
	path := "/v1/Trunks/{TrunkSid}/IpAccessControlLists/{Sid}"
	path = strings.Replace(path, "{"+"TrunkSid"+"}", TrunkSid, -1)
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

func (c *ApiService) FetchIpAccessControlList(TrunkSid string, Sid string) (*TrunkingV1IpAccessControlList, error) {
	path := "/v1/Trunks/{TrunkSid}/IpAccessControlLists/{Sid}"
	path = strings.Replace(path, "{"+"TrunkSid"+"}", TrunkSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TrunkingV1IpAccessControlList{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListIpAccessControlList'
type ListIpAccessControlListParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListIpAccessControlListParams) SetPageSize(PageSize int) *ListIpAccessControlListParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListIpAccessControlListParams) SetLimit(Limit int) *ListIpAccessControlListParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of IpAccessControlList records from the API. Request is executed immediately.
func (c *ApiService) PageIpAccessControlList(TrunkSid string, params *ListIpAccessControlListParams, pageToken, pageNumber string) (*ListIpAccessControlListResponse, error) {
	path := "/v1/Trunks/{TrunkSid}/IpAccessControlLists"

	path = strings.Replace(path, "{"+"TrunkSid"+"}", TrunkSid, -1)

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

	ps := &ListIpAccessControlListResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists IpAccessControlList records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListIpAccessControlList(TrunkSid string, params *ListIpAccessControlListParams) ([]TrunkingV1IpAccessControlList, error) {
	if params == nil {
		params = &ListIpAccessControlListParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageIpAccessControlList(TrunkSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []TrunkingV1IpAccessControlList

	for response != nil {
		records = append(records, response.IpAccessControlLists...)

		var record interface{}
		if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListIpAccessControlListResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListIpAccessControlListResponse)
	}

	return records, err
}

// Streams IpAccessControlList records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamIpAccessControlList(TrunkSid string, params *ListIpAccessControlListParams) (chan TrunkingV1IpAccessControlList, error) {
	if params == nil {
		params = &ListIpAccessControlListParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageIpAccessControlList(TrunkSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan TrunkingV1IpAccessControlList, 1)

	go func() {
		for response != nil {
			for item := range response.IpAccessControlLists {
				channel <- response.IpAccessControlLists[item]
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListIpAccessControlListResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListIpAccessControlListResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListIpAccessControlListResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListIpAccessControlListResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
