// Copyright 2018 JDCLOUD.COM
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// NOTE: This class is auto generated by the jdcloud code generator program.

package apis

import (
    "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/core"
    ucapi "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/ucapi/models"
)

type DescribeInvoiceRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 发票id  */
    Id string `json:"id"`

    /* 用户pin  */
    Pin string `json:"pin"`
}

/*
 * param regionId: Region ID (Required)
 * param id: 发票id (Required)
 * param pin: 用户pin (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeInvoiceRequest(
    regionId string,
    id string,
    pin string,
) *DescribeInvoiceRequest {

	return &DescribeInvoiceRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/invoices/{id}/{pin}",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        Id: id,
        Pin: pin,
	}
}

/*
 * param regionId: Region ID (Required)
 * param id: 发票id (Required)
 * param pin: 用户pin (Required)
 */
func NewDescribeInvoiceRequestWithAllParams(
    regionId string,
    id string,
    pin string,
) *DescribeInvoiceRequest {

    return &DescribeInvoiceRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/invoices/{id}/{pin}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        Id: id,
        Pin: pin,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeInvoiceRequestWithoutParam() *DescribeInvoiceRequest {

    return &DescribeInvoiceRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/invoices/{id}/{pin}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *DescribeInvoiceRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param id: 发票id(Required) */
func (r *DescribeInvoiceRequest) SetId(id string) {
    r.Id = id
}

/* param pin: 用户pin(Required) */
func (r *DescribeInvoiceRequest) SetPin(pin string) {
    r.Pin = pin
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeInvoiceRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeInvoiceResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeInvoiceResult `json:"result"`
}

type DescribeInvoiceResult struct {
    Detail ucapi.Invoice `json:"detail"`
}