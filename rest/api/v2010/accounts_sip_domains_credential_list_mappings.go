/*
 * Twilio - Api
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.29.1
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

// Optional parameters for the method 'CreateSipCredentialListMapping'
type CreateSipCredentialListMappingParams struct {
	// The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// A 34 character string that uniquely identifies the CredentialList resource to map to the SIP domain.
	CredentialListSid *string `json:"CredentialListSid,omitempty"`
}

func (params *CreateSipCredentialListMappingParams) SetPathAccountSid(PathAccountSid string) *CreateSipCredentialListMappingParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *CreateSipCredentialListMappingParams) SetCredentialListSid(CredentialListSid string) *CreateSipCredentialListMappingParams {
	params.CredentialListSid = &CredentialListSid
	return params
}

// Create a CredentialListMapping resource for an account.
func (c *ApiService) CreateSipCredentialListMapping(DomainSid string, params *CreateSipCredentialListMappingParams) (*ApiV2010SipCredentialListMapping, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/SIP/Domains/{DomainSid}/CredentialListMappings.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"DomainSid"+"}", DomainSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.CredentialListSid != nil {
		data.Set("CredentialListSid", *params.CredentialListSid)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ApiV2010SipCredentialListMapping{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'DeleteSipCredentialListMapping'
type DeleteSipCredentialListMappingParams struct {
	// The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
}

func (params *DeleteSipCredentialListMappingParams) SetPathAccountSid(PathAccountSid string) *DeleteSipCredentialListMappingParams {
	params.PathAccountSid = &PathAccountSid
	return params
}

// Delete a CredentialListMapping resource from an account.
func (c *ApiService) DeleteSipCredentialListMapping(DomainSid string, Sid string, params *DeleteSipCredentialListMappingParams) error {
	path := "/2010-04-01/Accounts/{AccountSid}/SIP/Domains/{DomainSid}/CredentialListMappings/{Sid}.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"DomainSid"+"}", DomainSid, -1)
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

// Optional parameters for the method 'FetchSipCredentialListMapping'
type FetchSipCredentialListMappingParams struct {
	// The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
}

func (params *FetchSipCredentialListMappingParams) SetPathAccountSid(PathAccountSid string) *FetchSipCredentialListMappingParams {
	params.PathAccountSid = &PathAccountSid
	return params
}

// Fetch a single CredentialListMapping resource from an account.
func (c *ApiService) FetchSipCredentialListMapping(DomainSid string, Sid string, params *FetchSipCredentialListMappingParams) (*ApiV2010SipCredentialListMapping, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/SIP/Domains/{DomainSid}/CredentialListMappings/{Sid}.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"DomainSid"+"}", DomainSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ApiV2010SipCredentialListMapping{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListSipCredentialListMapping'
type ListSipCredentialListMappingParams struct {
	// The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListSipCredentialListMappingParams) SetPathAccountSid(PathAccountSid string) *ListSipCredentialListMappingParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *ListSipCredentialListMappingParams) SetPageSize(PageSize int) *ListSipCredentialListMappingParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListSipCredentialListMappingParams) SetLimit(Limit int) *ListSipCredentialListMappingParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of SipCredentialListMapping records from the API. Request is executed immediately.
func (c *ApiService) PageSipCredentialListMapping(DomainSid string, params *ListSipCredentialListMappingParams, pageToken, pageNumber string) (*ListSipCredentialListMappingResponse, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/SIP/Domains/{DomainSid}/CredentialListMappings.json"

	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"DomainSid"+"}", DomainSid, -1)

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

	ps := &ListSipCredentialListMappingResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists SipCredentialListMapping records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListSipCredentialListMapping(DomainSid string, params *ListSipCredentialListMappingParams) ([]ApiV2010SipCredentialListMapping, error) {
	response, errors := c.StreamSipCredentialListMapping(DomainSid, params)

	records := make([]ApiV2010SipCredentialListMapping, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams SipCredentialListMapping records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamSipCredentialListMapping(DomainSid string, params *ListSipCredentialListMappingParams) (chan ApiV2010SipCredentialListMapping, chan error) {
	if params == nil {
		params = &ListSipCredentialListMappingParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan ApiV2010SipCredentialListMapping, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageSipCredentialListMapping(DomainSid, params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamSipCredentialListMapping(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamSipCredentialListMapping(response *ListSipCredentialListMappingResponse, params *ListSipCredentialListMappingParams, recordChannel chan ApiV2010SipCredentialListMapping, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.CredentialListMappings
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListSipCredentialListMappingResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListSipCredentialListMappingResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListSipCredentialListMappingResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListSipCredentialListMappingResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
