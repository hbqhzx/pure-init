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

type AddAllProtectedIpRequest struct {

    core.JDCloudRequest

    /* 所属地域 Id, 防护包目前支持 华北-北京, 华东-上海<br>- cn-north-1: 华北-北京<br>- cn-east-1: 华东-宿迁<br>- cn-east-2: 华东-上海<br>- cn-south-1: 华南-广州  */
    RegionId string `json:"regionId"`

    /* 防护包实例 Id  */
    InstanceId int `json:"instanceId"`
}

/*
 * param regionId: 所属地域 Id, 防护包目前支持 华北-北京, 华东-上海<br>- cn-north-1: 华北-北京<br>- cn-east-1: 华东-宿迁<br>- cn-east-2: 华东-上海<br>- cn-south-1: 华南-广州 (Required)
 * param instanceId: 防护包实例 Id (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewAddAllProtectedIpRequest(
    regionId string,
    instanceId int,
) *AddAllProtectedIpRequest {

	return &AddAllProtectedIpRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}:addAllProtectedIp",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceId: instanceId,
	}
}

/*
 * param regionId: 所属地域 Id, 防护包目前支持 华北-北京, 华东-上海<br>- cn-north-1: 华北-北京<br>- cn-east-1: 华东-宿迁<br>- cn-east-2: 华东-上海<br>- cn-south-1: 华南-广州 (Required)
 * param instanceId: 防护包实例 Id (Required)
 */
func NewAddAllProtectedIpRequestWithAllParams(
    regionId string,
    instanceId int,
) *AddAllProtectedIpRequest {

    return &AddAllProtectedIpRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}:addAllProtectedIp",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InstanceId: instanceId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewAddAllProtectedIpRequestWithoutParam() *AddAllProtectedIpRequest {

    return &AddAllProtectedIpRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}:addAllProtectedIp",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 所属地域 Id, 防护包目前支持 华北-北京, 华东-上海<br>- cn-north-1: 华北-北京<br>- cn-east-1: 华东-宿迁<br>- cn-east-2: 华东-上海<br>- cn-south-1: 华南-广州(Required) */
func (r *AddAllProtectedIpRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceId: 防护包实例 Id(Required) */
func (r *AddAllProtectedIpRequest) SetInstanceId(instanceId int) {
    r.InstanceId = instanceId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r AddAllProtectedIpRequest) GetRegionId() string {
    return r.RegionId
}

type AddAllProtectedIpResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result AddAllProtectedIpResult `json:"result"`
}

type AddAllProtectedIpResult struct {
    Code int `json:"code"`
    Message string `json:"message"`
}