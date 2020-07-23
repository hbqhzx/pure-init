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

type CheckIpBlackWhiteListRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 实例 ID  */
    InstanceId int `json:"instanceId"`

    /* 检测类型, 1: 白名单, 2: 黑名单  */
    Type int `json:"type"`

    /* ip 列表  */
    IpList []string `json:"ipList"`
}

/*
 * param regionId: Region ID (Required)
 * param instanceId: 实例 ID (Required)
 * param type_: 检测类型, 1: 白名单, 2: 黑名单 (Required)
 * param ipList: ip 列表 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCheckIpBlackWhiteListRequest(
    regionId string,
    instanceId int,
    type_ int,
    ipList []string,
) *CheckIpBlackWhiteListRequest {

	return &CheckIpBlackWhiteListRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}:checkIpBlackWhiteList",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceId: instanceId,
        Type: type_,
        IpList: ipList,
	}
}

/*
 * param regionId: Region ID (Required)
 * param instanceId: 实例 ID (Required)
 * param type_: 检测类型, 1: 白名单, 2: 黑名单 (Required)
 * param ipList: ip 列表 (Required)
 */
func NewCheckIpBlackWhiteListRequestWithAllParams(
    regionId string,
    instanceId int,
    type_ int,
    ipList []string,
) *CheckIpBlackWhiteListRequest {

    return &CheckIpBlackWhiteListRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}:checkIpBlackWhiteList",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InstanceId: instanceId,
        Type: type_,
        IpList: ipList,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCheckIpBlackWhiteListRequestWithoutParam() *CheckIpBlackWhiteListRequest {

    return &CheckIpBlackWhiteListRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}:checkIpBlackWhiteList",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *CheckIpBlackWhiteListRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceId: 实例 ID(Required) */
func (r *CheckIpBlackWhiteListRequest) SetInstanceId(instanceId int) {
    r.InstanceId = instanceId
}

/* param type_: 检测类型, 1: 白名单, 2: 黑名单(Required) */
func (r *CheckIpBlackWhiteListRequest) SetType(type_ int) {
    r.Type = type_
}

/* param ipList: ip 列表(Required) */
func (r *CheckIpBlackWhiteListRequest) SetIpList(ipList []string) {
    r.IpList = ipList
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CheckIpBlackWhiteListRequest) GetRegionId() string {
    return r.RegionId
}

type CheckIpBlackWhiteListResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CheckIpBlackWhiteListResult `json:"result"`
}

type CheckIpBlackWhiteListResult struct {
    Code int `json:"code"`
    Message string `json:"message"`
}