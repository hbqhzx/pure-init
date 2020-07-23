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

type SetUserQuotaOpsRequest struct {

    core.JDCloudRequest

    /* 所在区域的Region ID  */
    RegionId string `json:"regionId"`

    /* 运维人员erp账号  */
    Erp string `json:"erp"`

    /* 用户pin  */
    UserPin string `json:"userPin"`

    /* 要修改的quota值，大于0的整数  */
    Quota int `json:"quota"`
}

/*
 * param regionId: 所在区域的Region ID (Required)
 * param erp: 运维人员erp账号 (Required)
 * param userPin: 用户pin (Required)
 * param quota: 要修改的quota值，大于0的整数 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewSetUserQuotaOpsRequest(
    regionId string,
    erp string,
    userPin string,
    quota int,
) *SetUserQuotaOpsRequest {

	return &SetUserQuotaOpsRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/ops/{erp}/quota",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        Erp: erp,
        UserPin: userPin,
        Quota: quota,
	}
}

/*
 * param regionId: 所在区域的Region ID (Required)
 * param erp: 运维人员erp账号 (Required)
 * param userPin: 用户pin (Required)
 * param quota: 要修改的quota值，大于0的整数 (Required)
 */
func NewSetUserQuotaOpsRequestWithAllParams(
    regionId string,
    erp string,
    userPin string,
    quota int,
) *SetUserQuotaOpsRequest {

    return &SetUserQuotaOpsRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/ops/{erp}/quota",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        Erp: erp,
        UserPin: userPin,
        Quota: quota,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewSetUserQuotaOpsRequestWithoutParam() *SetUserQuotaOpsRequest {

    return &SetUserQuotaOpsRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/ops/{erp}/quota",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 所在区域的Region ID(Required) */
func (r *SetUserQuotaOpsRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param erp: 运维人员erp账号(Required) */
func (r *SetUserQuotaOpsRequest) SetErp(erp string) {
    r.Erp = erp
}

/* param userPin: 用户pin(Required) */
func (r *SetUserQuotaOpsRequest) SetUserPin(userPin string) {
    r.UserPin = userPin
}

/* param quota: 要修改的quota值，大于0的整数(Required) */
func (r *SetUserQuotaOpsRequest) SetQuota(quota int) {
    r.Quota = quota
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r SetUserQuotaOpsRequest) GetRegionId() string {
    return r.RegionId
}

type SetUserQuotaOpsResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result SetUserQuotaOpsResult `json:"result"`
}

type SetUserQuotaOpsResult struct {
}