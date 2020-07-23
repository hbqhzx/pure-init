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

type TryAntiProRequest struct {

    core.JDCloudRequest

    /* 所属地域 Id, 防护包目前支持 华北-北京, 华东-上海<br>- cn-north-1: 华北-北京<br>- cn-east-1: 华东-宿迁<br>- cn-east-2: 华东-上海<br>- cn-south-1: 华南-广州  */
    RegionId string `json:"regionId"`

    /* 防护包实例名称, 长度限制为 1-80 个字符, 只允许包含中文, 字母, 数字, -, ., /, _  */
    Name string `json:"name"`

    /* 被防护 IP  */
    Ip string `json:"ip"`
}

/*
 * param regionId: 所属地域 Id, 防护包目前支持 华北-北京, 华东-上海<br>- cn-north-1: 华北-北京<br>- cn-east-1: 华东-宿迁<br>- cn-east-2: 华东-上海<br>- cn-south-1: 华南-广州 (Required)
 * param name: 防护包实例名称, 长度限制为 1-80 个字符, 只允许包含中文, 字母, 数字, -, ., /, _ (Required)
 * param ip: 被防护 IP (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewTryAntiProRequest(
    regionId string,
    name string,
    ip string,
) *TryAntiProRequest {

	return &TryAntiProRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/tryAntiPro",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        Name: name,
        Ip: ip,
	}
}

/*
 * param regionId: 所属地域 Id, 防护包目前支持 华北-北京, 华东-上海<br>- cn-north-1: 华北-北京<br>- cn-east-1: 华东-宿迁<br>- cn-east-2: 华东-上海<br>- cn-south-1: 华南-广州 (Required)
 * param name: 防护包实例名称, 长度限制为 1-80 个字符, 只允许包含中文, 字母, 数字, -, ., /, _ (Required)
 * param ip: 被防护 IP (Required)
 */
func NewTryAntiProRequestWithAllParams(
    regionId string,
    name string,
    ip string,
) *TryAntiProRequest {

    return &TryAntiProRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/tryAntiPro",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        Name: name,
        Ip: ip,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewTryAntiProRequestWithoutParam() *TryAntiProRequest {

    return &TryAntiProRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/tryAntiPro",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 所属地域 Id, 防护包目前支持 华北-北京, 华东-上海<br>- cn-north-1: 华北-北京<br>- cn-east-1: 华东-宿迁<br>- cn-east-2: 华东-上海<br>- cn-south-1: 华南-广州(Required) */
func (r *TryAntiProRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param name: 防护包实例名称, 长度限制为 1-80 个字符, 只允许包含中文, 字母, 数字, -, ., /, _(Required) */
func (r *TryAntiProRequest) SetName(name string) {
    r.Name = name
}

/* param ip: 被防护 IP(Required) */
func (r *TryAntiProRequest) SetIp(ip string) {
    r.Ip = ip
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r TryAntiProRequest) GetRegionId() string {
    return r.RegionId
}

type TryAntiProResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result TryAntiProResult `json:"result"`
}

type TryAntiProResult struct {
    Code int `json:"code"`
    Message string `json:"message"`
}