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

type QueryUnPayableCouponListRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* resourceId ID  */
    OrderNumber string `json:"orderNumber"`
}

/*
 * param regionId: Region ID (Required)
 * param orderNumber: resourceId ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewQueryUnPayableCouponListRequest(
    regionId string,
    orderNumber string,
) *QueryUnPayableCouponListRequest {

	return &QueryUnPayableCouponListRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/order/{orderNumber}:unPayableCouponList",
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
 * param orderNumber: resourceId ID (Required)
 */
func NewQueryUnPayableCouponListRequestWithAllParams(
    regionId string,
    orderNumber string,
) *QueryUnPayableCouponListRequest {

    return &QueryUnPayableCouponListRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/order/{orderNumber}:unPayableCouponList",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        OrderNumber: orderNumber,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewQueryUnPayableCouponListRequestWithoutParam() *QueryUnPayableCouponListRequest {

    return &QueryUnPayableCouponListRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/order/{orderNumber}:unPayableCouponList",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *QueryUnPayableCouponListRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param orderNumber: resourceId ID(Required) */
func (r *QueryUnPayableCouponListRequest) SetOrderNumber(orderNumber string) {
    r.OrderNumber = orderNumber
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r QueryUnPayableCouponListRequest) GetRegionId() string {
    return r.RegionId
}

type QueryUnPayableCouponListResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result QueryUnPayableCouponListResult `json:"result"`
}

type QueryUnPayableCouponListResult struct {
    UnCoupons string `json:"unCoupons"`
    UnCouponCount int `json:"unCouponCount"`
}