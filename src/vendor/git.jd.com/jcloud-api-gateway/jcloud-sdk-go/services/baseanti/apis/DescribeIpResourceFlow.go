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

type DescribeIpResourceFlowRequest struct {

    core.JDCloudRequest

    /* 所属地域 Id, cn-north-1: 华北-北京, cn-east-1: 华东-宿迁, cn-east-2: 华东-上海, cn-south-1: 华南-广州  */
    RegionId string `json:"regionId"`

    /* 公网 IP 地址  */
    Ip string `json:"ip"`

    /* 查询的结束时间, UTC时间, 格式: yyyy-MM-dd'T'HH:mm:ssZ, 为空时查询当前时间之前 15 分钟内监控流量 (Optional) */
    EndTime *string `json:"endTime"`
}

/*
 * param regionId: 所属地域 Id, cn-north-1: 华北-北京, cn-east-1: 华东-宿迁, cn-east-2: 华东-上海, cn-south-1: 华南-广州 (Required)
 * param ip: 公网 IP 地址 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeIpResourceFlowRequest(
    regionId string,
    ip string,
) *DescribeIpResourceFlowRequest {

	return &DescribeIpResourceFlowRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/ipResources/{ip}/monitorFlow",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        Ip: ip,
	}
}

/*
 * param regionId: 所属地域 Id, cn-north-1: 华北-北京, cn-east-1: 华东-宿迁, cn-east-2: 华东-上海, cn-south-1: 华南-广州 (Required)
 * param ip: 公网 IP 地址 (Required)
 * param endTime: 查询的结束时间, UTC时间, 格式: yyyy-MM-dd'T'HH:mm:ssZ, 为空时查询当前时间之前 15 分钟内监控流量 (Optional)
 */
func NewDescribeIpResourceFlowRequestWithAllParams(
    regionId string,
    ip string,
    endTime *string,
) *DescribeIpResourceFlowRequest {

    return &DescribeIpResourceFlowRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/ipResources/{ip}/monitorFlow",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        Ip: ip,
        EndTime: endTime,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeIpResourceFlowRequestWithoutParam() *DescribeIpResourceFlowRequest {

    return &DescribeIpResourceFlowRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/ipResources/{ip}/monitorFlow",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 所属地域 Id, cn-north-1: 华北-北京, cn-east-1: 华东-宿迁, cn-east-2: 华东-上海, cn-south-1: 华南-广州(Required) */
func (r *DescribeIpResourceFlowRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param ip: 公网 IP 地址(Required) */
func (r *DescribeIpResourceFlowRequest) SetIp(ip string) {
    r.Ip = ip
}

/* param endTime: 查询的结束时间, UTC时间, 格式: yyyy-MM-dd'T'HH:mm:ssZ, 为空时查询当前时间之前 15 分钟内监控流量(Optional) */
func (r *DescribeIpResourceFlowRequest) SetEndTime(endTime string) {
    r.EndTime = &endTime
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeIpResourceFlowRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeIpResourceFlowResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeIpResourceFlowResult `json:"result"`
}

type DescribeIpResourceFlowResult struct {
    Bps baseanti.IpResourceFlow `json:"bps"`
    Pps baseanti.IpResourceFlow `json:"pps"`
}