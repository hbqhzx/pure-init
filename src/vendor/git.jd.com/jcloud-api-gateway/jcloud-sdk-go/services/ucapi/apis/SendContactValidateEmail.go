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

type SendContactValidateEmailRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 验证码 (Optional) */
    Code *string `json:"code"`
}

/*
 * param regionId: Region ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewSendContactValidateEmailRequest(
    regionId string,
) *SendContactValidateEmailRequest {

	return &SendContactValidateEmailRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/contact/validate/email",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: Region ID (Required)
 * param code: 验证码 (Optional)
 */
func NewSendContactValidateEmailRequestWithAllParams(
    regionId string,
    code *string,
) *SendContactValidateEmailRequest {

    return &SendContactValidateEmailRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/contact/validate/email",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        Code: code,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewSendContactValidateEmailRequestWithoutParam() *SendContactValidateEmailRequest {

    return &SendContactValidateEmailRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/contact/validate/email",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *SendContactValidateEmailRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param code: 验证码(Optional) */
func (r *SendContactValidateEmailRequest) SetCode(code string) {
    r.Code = &code
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r SendContactValidateEmailRequest) GetRegionId() string {
    return r.RegionId
}

type SendContactValidateEmailResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result SendContactValidateEmailResult `json:"result"`
}

type SendContactValidateEmailResult struct {
}