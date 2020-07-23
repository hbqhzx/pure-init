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

type CheckInstanceNameRequest struct {

    core.JDCloudRequest

    /* 待检测实例名称, 长度限制为 1-80 个字符, 只允许包含中文, 字母, 数字, -, ., /, _  */
    InstanceName string `json:"instanceName"`
}

/*
 * param instanceName: 待检测实例名称, 长度限制为 1-80 个字符, 只允许包含中文, 字母, 数字, -, ., /, _ (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCheckInstanceNameRequest(
    instanceName string,
) *CheckInstanceNameRequest {

	return &CheckInstanceNameRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/checkInstanceName",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        InstanceName: instanceName,
	}
}

/*
 * param instanceName: 待检测实例名称, 长度限制为 1-80 个字符, 只允许包含中文, 字母, 数字, -, ., /, _ (Required)
 */
func NewCheckInstanceNameRequestWithAllParams(
    instanceName string,
) *CheckInstanceNameRequest {

    return &CheckInstanceNameRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/checkInstanceName",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        InstanceName: instanceName,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCheckInstanceNameRequestWithoutParam() *CheckInstanceNameRequest {

    return &CheckInstanceNameRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/checkInstanceName",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param instanceName: 待检测实例名称, 长度限制为 1-80 个字符, 只允许包含中文, 字母, 数字, -, ., /, _(Required) */
func (r *CheckInstanceNameRequest) SetInstanceName(instanceName string) {
    r.InstanceName = instanceName
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CheckInstanceNameRequest) GetRegionId() string {
    return ""
}

type CheckInstanceNameResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CheckInstanceNameResult `json:"result"`
}

type CheckInstanceNameResult struct {
    Code int `json:"code"`
    Message string `json:"message"`
}