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
    iam "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/iam/models"
)

type SendOutSiteNoticeRequest struct {

    core.JDCloudRequest

    /* 站外信接口  */
    SendOutSiteNoticeReqVo *iam.SendOutSiteNotice `json:"sendOutSiteNoticeReqVo"`
}

/*
 * param sendOutSiteNoticeReqVo: 站外信接口 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewSendOutSiteNoticeRequest(
    sendOutSiteNoticeReqVo *iam.SendOutSiteNotice,
) *SendOutSiteNoticeRequest {

	return &SendOutSiteNoticeRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/sms:sendOutSiteNotice",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        SendOutSiteNoticeReqVo: sendOutSiteNoticeReqVo,
	}
}

/*
 * param sendOutSiteNoticeReqVo: 站外信接口 (Required)
 */
func NewSendOutSiteNoticeRequestWithAllParams(
    sendOutSiteNoticeReqVo *iam.SendOutSiteNotice,
) *SendOutSiteNoticeRequest {

    return &SendOutSiteNoticeRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/sms:sendOutSiteNotice",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        SendOutSiteNoticeReqVo: sendOutSiteNoticeReqVo,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewSendOutSiteNoticeRequestWithoutParam() *SendOutSiteNoticeRequest {

    return &SendOutSiteNoticeRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/sms:sendOutSiteNotice",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param sendOutSiteNoticeReqVo: 站外信接口(Required) */
func (r *SendOutSiteNoticeRequest) SetSendOutSiteNoticeReqVo(sendOutSiteNoticeReqVo *iam.SendOutSiteNotice) {
    r.SendOutSiteNoticeReqVo = sendOutSiteNoticeReqVo
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r SendOutSiteNoticeRequest) GetRegionId() string {
    return ""
}

type SendOutSiteNoticeResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result SendOutSiteNoticeResult `json:"result"`
}

type SendOutSiteNoticeResult struct {
    Data bool `json:"data"`
}