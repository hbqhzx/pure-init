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

type CreateCouponByLinkRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 用户pin (Optional) */
    Pin *string `json:"pin"`

    /* 批次号 (Optional) */
    BatchID *string `json:"batchID"`
}

/*
 * param regionId: Region ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateCouponByLinkRequest(
    regionId string,
) *CreateCouponByLinkRequest {

	return &CreateCouponByLinkRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/coupon/link",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: Region ID (Required)
 * param pin: 用户pin (Optional)
 * param batchID: 批次号 (Optional)
 */
func NewCreateCouponByLinkRequestWithAllParams(
    regionId string,
    pin *string,
    batchID *string,
) *CreateCouponByLinkRequest {

    return &CreateCouponByLinkRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/coupon/link",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        Pin: pin,
        BatchID: batchID,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateCouponByLinkRequestWithoutParam() *CreateCouponByLinkRequest {

    return &CreateCouponByLinkRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/coupon/link",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *CreateCouponByLinkRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param pin: 用户pin(Optional) */
func (r *CreateCouponByLinkRequest) SetPin(pin string) {
    r.Pin = &pin
}

/* param batchID: 批次号(Optional) */
func (r *CreateCouponByLinkRequest) SetBatchID(batchID string) {
    r.BatchID = &batchID
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateCouponByLinkRequest) GetRegionId() string {
    return r.RegionId
}

type CreateCouponByLinkResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateCouponByLinkResult `json:"result"`
}

type CreateCouponByLinkResult struct {
}