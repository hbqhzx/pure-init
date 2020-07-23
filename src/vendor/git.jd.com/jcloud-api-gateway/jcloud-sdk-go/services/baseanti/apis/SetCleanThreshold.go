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
    baseanti "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/baseanti/models"
)

type SetCleanThresholdRequest struct {

    core.JDCloudRequest

    /* 所属地域 Id, cn-north-1: 华北-北京, cn-east-1: 华东-宿迁, cn-east-2: 华东-上海, cn-south-1: 华南-广州  */
    RegionId string `json:"regionId"`

    /* 公网 IP 地址  */
    Ip string `json:"ip"`

    /* cc参数  */
    CleanThresholdSpec *baseanti.CleanThresholdSpec `json:"cleanThresholdSpec"`
}

/*
 * param regionId: 所属地域 Id, cn-north-1: 华北-北京, cn-east-1: 华东-宿迁, cn-east-2: 华东-上海, cn-south-1: 华南-广州 (Required)
 * param ip: 公网 IP 地址 (Required)
 * param cleanThresholdSpec: cc参数 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewSetCleanThresholdRequest(
    regionId string,
    ip string,
    cleanThresholdSpec *baseanti.CleanThresholdSpec,
) *SetCleanThresholdRequest {

	return &SetCleanThresholdRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/ipResources/{ip}:setCleanThreshold",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        Ip: ip,
        CleanThresholdSpec: cleanThresholdSpec,
	}
}

/*
 * param regionId: 所属地域 Id, cn-north-1: 华北-北京, cn-east-1: 华东-宿迁, cn-east-2: 华东-上海, cn-south-1: 华南-广州 (Required)
 * param ip: 公网 IP 地址 (Required)
 * param cleanThresholdSpec: cc参数 (Required)
 */
func NewSetCleanThresholdRequestWithAllParams(
    regionId string,
    ip string,
    cleanThresholdSpec *baseanti.CleanThresholdSpec,
) *SetCleanThresholdRequest {

    return &SetCleanThresholdRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/ipResources/{ip}:setCleanThreshold",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        Ip: ip,
        CleanThresholdSpec: cleanThresholdSpec,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewSetCleanThresholdRequestWithoutParam() *SetCleanThresholdRequest {

    return &SetCleanThresholdRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/ipResources/{ip}:setCleanThreshold",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 所属地域 Id, cn-north-1: 华北-北京, cn-east-1: 华东-宿迁, cn-east-2: 华东-上海, cn-south-1: 华南-广州(Required) */
func (r *SetCleanThresholdRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param ip: 公网 IP 地址(Required) */
func (r *SetCleanThresholdRequest) SetIp(ip string) {
    r.Ip = ip
}

/* param cleanThresholdSpec: cc参数(Required) */
func (r *SetCleanThresholdRequest) SetCleanThresholdSpec(cleanThresholdSpec *baseanti.CleanThresholdSpec) {
    r.CleanThresholdSpec = cleanThresholdSpec
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r SetCleanThresholdRequest) GetRegionId() string {
    return r.RegionId
}

type SetCleanThresholdResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result SetCleanThresholdResult `json:"result"`
}

type SetCleanThresholdResult struct {
}