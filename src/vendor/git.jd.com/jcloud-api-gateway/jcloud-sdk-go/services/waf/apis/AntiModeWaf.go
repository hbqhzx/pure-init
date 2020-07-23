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
    waf "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/waf/models"
)

type AntiModeWafRequest struct {

    core.JDCloudRequest

    /* 实例所属的地域ID  */
    RegionId string `json:"regionId"`

    /* 请求  */
    Req *waf.AntiModeWafReq `json:"req"`
}

/*
 * param regionId: 实例所属的地域ID (Required)
 * param req: 请求 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewAntiModeWafRequest(
    regionId string,
    req *waf.AntiModeWafReq,
) *AntiModeWafRequest {

	return &AntiModeWafRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/waf:antiMode",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        Req: req,
	}
}

/*
 * param regionId: 实例所属的地域ID (Required)
 * param req: 请求 (Required)
 */
func NewAntiModeWafRequestWithAllParams(
    regionId string,
    req *waf.AntiModeWafReq,
) *AntiModeWafRequest {

    return &AntiModeWafRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/waf:antiMode",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        Req: req,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewAntiModeWafRequestWithoutParam() *AntiModeWafRequest {

    return &AntiModeWafRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/waf:antiMode",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 实例所属的地域ID(Required) */
func (r *AntiModeWafRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param req: 请求(Required) */
func (r *AntiModeWafRequest) SetReq(req *waf.AntiModeWafReq) {
    r.Req = req
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r AntiModeWafRequest) GetRegionId() string {
    return r.RegionId
}

type AntiModeWafResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result AntiModeWafResult `json:"result"`
}

type AntiModeWafResult struct {
}