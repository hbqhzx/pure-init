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
)

type DeleteInvoiceMsgTemplateRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 用户pin (Optional) */
    Pin *string `json:"pin"`

    /* 专票资质id (Optional) */
    VatId *int `json:"vatId"`
}

/*
 * param regionId: Region ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDeleteInvoiceMsgTemplateRequest(
    regionId string,
) *DeleteInvoiceMsgTemplateRequest {

	return &DeleteInvoiceMsgTemplateRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/invoices/invoiceMsgTemplate",
			Method:  "DELETE",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: Region ID (Required)
 * param pin: 用户pin (Optional)
 * param vatId: 专票资质id (Optional)
 */
func NewDeleteInvoiceMsgTemplateRequestWithAllParams(
    regionId string,
    pin *string,
    vatId *int,
) *DeleteInvoiceMsgTemplateRequest {

    return &DeleteInvoiceMsgTemplateRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/invoices/invoiceMsgTemplate",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        Pin: pin,
        VatId: vatId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDeleteInvoiceMsgTemplateRequestWithoutParam() *DeleteInvoiceMsgTemplateRequest {

    return &DeleteInvoiceMsgTemplateRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/invoices/invoiceMsgTemplate",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *DeleteInvoiceMsgTemplateRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param pin: 用户pin(Optional) */
func (r *DeleteInvoiceMsgTemplateRequest) SetPin(pin string) {
    r.Pin = &pin
}

/* param vatId: 专票资质id(Optional) */
func (r *DeleteInvoiceMsgTemplateRequest) SetVatId(vatId int) {
    r.VatId = &vatId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DeleteInvoiceMsgTemplateRequest) GetRegionId() string {
    return r.RegionId
}

type DeleteInvoiceMsgTemplateResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DeleteInvoiceMsgTemplateResult `json:"result"`
}

type DeleteInvoiceMsgTemplateResult struct {
}