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

type BindEmailRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 用户pin  */
    Pin string `json:"pin"`

    /* 绑定邮箱信息 (Optional) */
    AccountEmail *ucapi.AccountEmail `json:"accountEmail"`
}

/*
 * param regionId: Region ID (Required)
 * param pin: 用户pin (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewBindEmailRequest(
    regionId string,
    pin string,
) *BindEmailRequest {

	return &BindEmailRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/user/{pin}:bindEmail",
			Method:  "POST",
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
 * param accountEmail: 绑定邮箱信息 (Optional)
 */
func NewBindEmailRequestWithAllParams(
    regionId string,
    pin string,
    accountEmail *ucapi.AccountEmail,
) *BindEmailRequest {

    return &BindEmailRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/user/{pin}:bindEmail",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        Pin: pin,
        AccountEmail: accountEmail,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewBindEmailRequestWithoutParam() *BindEmailRequest {

    return &BindEmailRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/user/{pin}:bindEmail",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *BindEmailRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param pin: 用户pin(Required) */
func (r *BindEmailRequest) SetPin(pin string) {
    r.Pin = pin
}

/* param accountEmail: 绑定邮箱信息(Optional) */
func (r *BindEmailRequest) SetAccountEmail(accountEmail *ucapi.AccountEmail) {
    r.AccountEmail = accountEmail
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r BindEmailRequest) GetRegionId() string {
    return r.RegionId
}

type BindEmailResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result BindEmailResult `json:"result"`
}

type BindEmailResult struct {
}