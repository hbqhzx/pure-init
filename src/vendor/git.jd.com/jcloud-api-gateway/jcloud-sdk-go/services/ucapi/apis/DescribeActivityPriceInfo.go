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

type DescribeActivityPriceInfoRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 用户pin  */
    Pin string `json:"pin"`

    /* 商品标识 (Optional) */
    ActivityProductCode *string `json:"activityProductCode"`

    /* 购买云主机时选的规格唯一标识 (Optional) */
    Element *string `json:"element"`
}

/*
 * param regionId: Region ID (Required)
 * param pin: 用户pin (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeActivityPriceInfoRequest(
    regionId string,
    pin string,
) *DescribeActivityPriceInfoRequest {

	return &DescribeActivityPriceInfoRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/user/{pin}:describeActivityPriceInfo",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        Pin: pin,
	}
}

/*
 * param regionId: Region ID (Required)
 * param pin: 用户pin (Required)
 * param activityProductCode: 商品标识 (Optional)
 * param element: 购买云主机时选的规格唯一标识 (Optional)
 */
func NewDescribeActivityPriceInfoRequestWithAllParams(
    regionId string,
    pin string,
    activityProductCode *string,
    element *string,
) *DescribeActivityPriceInfoRequest {

    return &DescribeActivityPriceInfoRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/user/{pin}:describeActivityPriceInfo",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        Pin: pin,
        ActivityProductCode: activityProductCode,
        Element: element,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeActivityPriceInfoRequestWithoutParam() *DescribeActivityPriceInfoRequest {

    return &DescribeActivityPriceInfoRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/user/{pin}:describeActivityPriceInfo",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *DescribeActivityPriceInfoRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param pin: 用户pin(Required) */
func (r *DescribeActivityPriceInfoRequest) SetPin(pin string) {
    r.Pin = pin
}

/* param activityProductCode: 商品标识(Optional) */
func (r *DescribeActivityPriceInfoRequest) SetActivityProductCode(activityProductCode string) {
    r.ActivityProductCode = &activityProductCode
}

/* param element: 购买云主机时选的规格唯一标识(Optional) */
func (r *DescribeActivityPriceInfoRequest) SetElement(element string) {
    r.Element = &element
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeActivityPriceInfoRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeActivityPriceInfoResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeActivityPriceInfoResult `json:"result"`
}

type DescribeActivityPriceInfoResult struct {
    PriceInfo ucapi.ActivityProductPrice `json:"priceInfo"`
}