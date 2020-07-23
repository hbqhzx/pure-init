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
    ipanti "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/ipanti/models"
)

type DescribeTotalPriceRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 查询套餐价格请求参数  */
    PriceSpec *ipanti.PriceSpec `json:"priceSpec"`
}

/*
 * param regionId: Region ID (Required)
 * param priceSpec: 查询套餐价格请求参数 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeTotalPriceRequest(
    regionId string,
    priceSpec *ipanti.PriceSpec,
) *DescribeTotalPriceRequest {

	return &DescribeTotalPriceRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/describeTotalPrice",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        PriceSpec: priceSpec,
	}
}

/*
 * param regionId: Region ID (Required)
 * param priceSpec: 查询套餐价格请求参数 (Required)
 */
func NewDescribeTotalPriceRequestWithAllParams(
    regionId string,
    priceSpec *ipanti.PriceSpec,
) *DescribeTotalPriceRequest {

    return &DescribeTotalPriceRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/describeTotalPrice",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        PriceSpec: priceSpec,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeTotalPriceRequestWithoutParam() *DescribeTotalPriceRequest {

    return &DescribeTotalPriceRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/describeTotalPrice",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *DescribeTotalPriceRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param priceSpec: 查询套餐价格请求参数(Required) */
func (r *DescribeTotalPriceRequest) SetPriceSpec(priceSpec *ipanti.PriceSpec) {
    r.PriceSpec = priceSpec
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeTotalPriceRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeTotalPriceResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeTotalPriceResult `json:"result"`
}

type DescribeTotalPriceResult struct {
    Price float64 `json:"price"`
    DiscountPrice float64 `json:"discountPrice"`
}