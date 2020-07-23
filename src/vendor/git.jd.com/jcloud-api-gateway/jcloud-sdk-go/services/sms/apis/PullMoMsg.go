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

type PullMoMsgRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 拉取回复短信请求参数  */
    PullMoMsgSpec *sms.PullMoMsgSpec `json:"pullMoMsgSpec"`
}

/*
 * param regionId: Region ID (Required)
 * param pullMoMsgSpec: 拉取回复短信请求参数 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewPullMoMsgRequest(
    regionId string,
    pullMoMsgSpec *sms.PullMoMsgSpec,
) *PullMoMsgRequest {

	return &PullMoMsgRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/pullMoMsg",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        PullMoMsgSpec: pullMoMsgSpec,
	}
}

/*
 * param regionId: Region ID (Required)
 * param pullMoMsgSpec: 拉取回复短信请求参数 (Required)
 */
func NewPullMoMsgRequestWithAllParams(
    regionId string,
    pullMoMsgSpec *sms.PullMoMsgSpec,
) *PullMoMsgRequest {

    return &PullMoMsgRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/pullMoMsg",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        PullMoMsgSpec: pullMoMsgSpec,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewPullMoMsgRequestWithoutParam() *PullMoMsgRequest {

    return &PullMoMsgRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/pullMoMsg",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *PullMoMsgRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param pullMoMsgSpec: 拉取回复短信请求参数(Required) */
func (r *PullMoMsgRequest) SetPullMoMsgSpec(pullMoMsgSpec *sms.PullMoMsgSpec) {
    r.PullMoMsgSpec = pullMoMsgSpec
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r PullMoMsgRequest) GetRegionId() string {
    return r.RegionId
}

type PullMoMsgResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result PullMoMsgResult `json:"result"`
}

type PullMoMsgResult struct {
    Data []sms.PullMoMsg `json:"data"`
}