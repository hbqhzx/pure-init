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
    order "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/order/models"
)

type DescribeOrdersRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* Order Number  */
    OrderNumber string `json:"orderNumber"`
}

/*
 * param regionId: Region ID (Required)
 * param orderNumber: Order Number (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeOrdersRequest(
    regionId string,
    orderNumber string,
) *DescribeOrdersRequest {

	return &DescribeOrdersRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/orders/{orderNumber}",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        OrderNumber: orderNumber,
	}
}

/*
 * param regionId: Region ID (Required)
 * param orderNumber: Order Number (Required)
 */
func NewDescribeOrdersRequestWithAllParams(
    regionId string,
    orderNumber string,
) *DescribeOrdersRequest {

    return &DescribeOrdersRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/orders/{orderNumber}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        OrderNumber: orderNumber,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeOrdersRequestWithoutParam() *DescribeOrdersRequest {

    return &DescribeOrdersRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/orders/{orderNumber}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *DescribeOrdersRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param orderNumber: Order Number(Required) */
func (r *DescribeOrdersRequest) SetOrderNumber(orderNumber string) {
    r.OrderNumber = orderNumber
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeOrdersRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeOrdersResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeOrdersResult `json:"result"`
}

type DescribeOrdersResult struct {
    OrderDetail order.OrderDetail `json:"orderDetail"`
}