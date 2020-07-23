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
    invoice "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/invoice/models"
)

type DetailRequest struct {

    core.JDCloudRequest

    /* 地域编码，参考OpenAPI公共说明  */
    RegionId string `json:"regionId"`

    /* 发票id  */
    InvoiceId int `json:"invoiceId"`

    /* 发票类型  */
    Type int `json:"type"`
}

/*
 * param regionId: 地域编码，参考OpenAPI公共说明 (Required)
 * param invoiceId: 发票id (Required)
 * param type_: 发票类型 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDetailRequest(
    regionId string,
    invoiceId int,
    type_ int,
) *DetailRequest {

	return &DetailRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/invoice/{invoiceId}",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InvoiceId: invoiceId,
        Type: type_,
	}
}

/*
 * param regionId: 地域编码，参考OpenAPI公共说明 (Required)
 * param invoiceId: 发票id (Required)
 * param type_: 发票类型 (Required)
 */
func NewDetailRequestWithAllParams(
    regionId string,
    invoiceId int,
    type_ int,
) *DetailRequest {

    return &DetailRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/invoice/{invoiceId}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InvoiceId: invoiceId,
        Type: type_,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDetailRequestWithoutParam() *DetailRequest {

    return &DetailRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/invoice/{invoiceId}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域编码，参考OpenAPI公共说明(Required) */
func (r *DetailRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param invoiceId: 发票id(Required) */
func (r *DetailRequest) SetInvoiceId(invoiceId int) {
    r.InvoiceId = invoiceId
}

/* param type_: 发票类型(Required) */
func (r *DetailRequest) SetType(type_ int) {
    r.Type = type_
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DetailRequest) GetRegionId() string {
    return r.RegionId
}

type DetailResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DetailResult `json:"result"`
}

type DetailResult struct {
    Id int `json:"id"`
    Pin string `json:"pin"`
    InvoiceType int `json:"invoiceType"`
    InvoiceTitle string `json:"invoiceTitle"`
    TaxNo string `json:"taxNo"`
    TotalPrice int `json:"totalPrice"`
    Remark string `json:"remark"`
    InvoiceContent string `json:"invoiceContent"`
    UserType int `json:"userType"`
    AddressId int `json:"addressId"`
    TransportId int `json:"transportId"`
    CreateTime string `json:"createTime"`
    UpdateTime string `json:"updateTime"`
    Status int `json:"status"`
    BillingTime string `json:"billingTime"`
    Reason string `json:"reason"`
    InvoiceOrg string `json:"invoiceOrg"`
    RegisterAddress string `json:"registerAddress"`
    AccountBank string `json:"accountBank"`
    Account string `json:"account"`
    Type int `json:"type"`
    Logistics invoice.LogisticsVo `json:"logistics"`
    PostAddress invoice.PostAddressVo `json:"postAddress"`
    InvoiceOrders []invoice.InvoiceOrderVo `json:"invoiceOrders"`
}