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
    sms "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/sms/models"
)

type PullMtMsgByMobileRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 拉取单个手机短信状态请求参数  */
    PullMtMsgByMobileSpec *sms.PullMtMsgByMobileSpec `json:"pullMtMsgByMobileSpec"`
}

/*
 * param regionId: Region ID (Required)
 * param pullMtMsgByMobileSpec: 拉取单个手机短信状态请求参数 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewPullMtMsgByMobileRequest(
    regionId string,
    pullMtMsgByMobileSpec *sms.PullMtMsgByMobileSpec,
) *PullMtMsgByMobileRequest {

	return &PullMtMsgByMobileRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/pullMtMsgByMobile",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        PullMtMsgByMobileSpec: pullMtMsgByMobileSpec,
	}
}

/*
 * param regionId: Region ID (Required)
 * param pullMtMsgByMobileSpec: 拉取单个手机短信状态请求参数 (Required)
 */
func NewPullMtMsgByMobileRequestWithAllParams(
    regionId string,
    pullMtMsgByMobileSpec *sms.PullMtMsgByMobileSpec,
) *PullMtMsgByMobileRequest {

    return &PullMtMsgByMobileRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/pullMtMsgByMobile",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        PullMtMsgByMobileSpec: pullMtMsgByMobileSpec,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewPullMtMsgByMobileRequestWithoutParam() *PullMtMsgByMobileRequest {

    return &PullMtMsgByMobileRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/pullMtMsgByMobile",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *PullMtMsgByMobileRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param pullMtMsgByMobileSpec: 拉取单个手机短信状态请求参数(Required) */
func (r *PullMtMsgByMobileRequest) SetPullMtMsgByMobileSpec(pullMtMsgByMobileSpec *sms.PullMtMsgByMobileSpec) {
    r.PullMtMsgByMobileSpec = pullMtMsgByMobileSpec
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r PullMtMsgByMobileRequest) GetRegionId() string {
    return r.RegionId
}

type PullMtMsgByMobileResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result PullMtMsgByMobileResult `json:"result"`
}

type PullMtMsgByMobileResult struct {
    Data []sms.PullMtMsgByMobile `json:"data"`
}