/*
 * Twilio - Supersim
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

	"github.com/twilio/twilio-go/client"
)

// SupersimV1UsageRecord struct for SupersimV1UsageRecord
type SupersimV1UsageRecord struct {
	// The SID of the Account that incurred the usage.
	AccountSid *string `json:"account_sid,omitempty"`
	// The currency in which the billed amounts are measured, specified in the 3 letter ISO 4127 format (e.g. `USD`, `EUR`, `JPY`).
	BilledUnit *string `json:"billed_unit,omitempty"`
	// Total data downloaded in bytes, aggregated by the query parameters.
	DataDownload *int `json:"data_download,omitempty"`
	// Total of data_upload and data_download.
	DataTotal *int `json:"data_total,omitempty"`
	// Total amount in the `billed_unit` that was charged for the data uploaded or downloaded.
	DataTotalBilled *float32 `json:"data_total_billed,omitempty"`
	// Total data uploaded in bytes, aggregated by the query parameters.
	DataUpload *int `json:"data_upload,omitempty"`
	// SID of the Fleet resource on which the usage occurred.
	FleetSid *string `json:"fleet_sid,omitempty"`
	// Alpha-2 ISO Country Code of the country the usage occurred in.
	IsoCountry *string `json:"iso_country,omitempty"`
	// SID of the Network resource on which the usage occurred.
	NetworkSid *string `json:"network_sid,omitempty"`
	// The time period for which the usage is reported.
	Period *interface{} `json:"period,omitempty"`
	// SID of a Sim resource to which the UsageRecord belongs.
	SimSid *string `json:"sim_sid,omitempty"`
}

func (response *SupersimV1UsageRecord) UnmarshalJSON(bytes []byte) (err error) {
	raw := struct {
		AccountSid      *string      `json:"account_sid"`
		BilledUnit      *string      `json:"billed_unit"`
		DataDownload    *int         `json:"data_download"`
		DataTotal       *int         `json:"data_total"`
		DataTotalBilled *interface{} `json:"data_total_billed"`
		DataUpload      *int         `json:"data_upload"`
		FleetSid        *string      `json:"fleet_sid"`
		IsoCountry      *string      `json:"iso_country"`
		NetworkSid      *string      `json:"network_sid"`
		Period          *interface{} `json:"period"`
		SimSid          *string      `json:"sim_sid"`
	}{}

	if err = json.Unmarshal(bytes, &raw); err != nil {
		return err
	}

	*response = SupersimV1UsageRecord{
		AccountSid:   raw.AccountSid,
		BilledUnit:   raw.BilledUnit,
		DataDownload: raw.DataDownload,
		DataTotal:    raw.DataTotal,
		DataUpload:   raw.DataUpload,
		FleetSid:     raw.FleetSid,
		IsoCountry:   raw.IsoCountry,
		NetworkSid:   raw.NetworkSid,
		Period:       raw.Period,
		SimSid:       raw.SimSid,
	}

	responseDataTotalBilled, err := client.UnmarshalFloat32(raw.DataTotalBilled)
	if err != nil {
		return err
	}
	response.DataTotalBilled = responseDataTotalBilled

	return
}
