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
    sop "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/sop/models"
)

type SendSopVerificationCodeRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 发送验证码参数  */
    VerificationCodeInfo *sop.VerificationCodeInfo `json:"verificationCodeInfo"`
}

/*
 * param regionId: Region ID (Required)
 * param verificationCodeInfo: 发送验证码参数 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewSendSopVerificationCodeRequest(
    regionId string,
    verificationCodeInfo *sop.VerificationCodeInfo,
) *SendSopVerificationCodeRequest {

	return &SendSopVerificationCodeRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/securityToken:sendSopVerificationCode",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        VerificationCodeInfo: verificationCodeInfo,
	}
}

/*
 * param regionId: Region ID (Required)
 * param verificationCodeInfo: 发送验证码参数 (Required)
 */
func NewSendSopVerificationCodeRequestWithAllParams(
    regionId string,
    verificationCodeInfo *sop.VerificationCodeInfo,
) *SendSopVerificationCodeRequest {

    return &SendSopVerificationCodeRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/securityToken:sendSopVerificationCode",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        VerificationCodeInfo: verificationCodeInfo,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewSendSopVerificationCodeRequestWithoutParam() *SendSopVerificationCodeRequest {

    return &SendSopVerificationCodeRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/securityToken:sendSopVerificationCode",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *SendSopVerificationCodeRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param verificationCodeInfo: 发送验证码参数(Required) */
func (r *SendSopVerificationCodeRequest) SetVerificationCodeInfo(verificationCodeInfo *sop.VerificationCodeInfo) {
    r.VerificationCodeInfo = verificationCodeInfo
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r SendSopVerificationCodeRequest) GetRegionId() string {
    return r.RegionId
}

type SendSopVerificationCodeResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result SendSopVerificationCodeResult `json:"result"`
}

type SendSopVerificationCodeResult struct {
}