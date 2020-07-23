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

type DisableWhiteListRuleOfForwardRuleRequest struct {

    core.JDCloudRequest

    /* 区域 Id  */
    RegionId string `json:"regionId"`

    /* 高防实例 Id  */
    InstanceId int `json:"instanceId"`

    /* 转发规则 Id  */
    ForwardRuleId int `json:"forwardRuleId"`

    /* 待引用的 IP 黑白名单规则 Id  */
    IpRuleId string `json:"ipRuleId"`
}

/*
 * param regionId: 区域 Id (Required)
 * param instanceId: 高防实例 Id (Required)
 * param forwardRuleId: 转发规则 Id (Required)
 * param ipRuleId: 待引用的 IP 黑白名单规则 Id (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDisableWhiteListRuleOfForwardRuleRequest(
    regionId string,
    instanceId int,
    forwardRuleId int,
    ipRuleId string,
) *DisableWhiteListRuleOfForwardRuleRequest {

	return &DisableWhiteListRuleOfForwardRuleRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}/forwardRules/{forwardRuleId}:disableWhiteListRuleOfForwardRule",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceId: instanceId,
        ForwardRuleId: forwardRuleId,
        IpRuleId: ipRuleId,
	}
}

/*
 * param regionId: 区域 Id (Required)
 * param instanceId: 高防实例 Id (Required)
 * param forwardRuleId: 转发规则 Id (Required)
 * param ipRuleId: 待引用的 IP 黑白名单规则 Id (Required)
 */
func NewDisableWhiteListRuleOfForwardRuleRequestWithAllParams(
    regionId string,
    instanceId int,
    forwardRuleId int,
    ipRuleId string,
) *DisableWhiteListRuleOfForwardRuleRequest {

    return &DisableWhiteListRuleOfForwardRuleRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/forwardRules/{forwardRuleId}:disableWhiteListRuleOfForwardRule",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InstanceId: instanceId,
        ForwardRuleId: forwardRuleId,
        IpRuleId: ipRuleId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDisableWhiteListRuleOfForwardRuleRequestWithoutParam() *DisableWhiteListRuleOfForwardRuleRequest {

    return &DisableWhiteListRuleOfForwardRuleRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/forwardRules/{forwardRuleId}:disableWhiteListRuleOfForwardRule",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 区域 Id(Required) */
func (r *DisableWhiteListRuleOfForwardRuleRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceId: 高防实例 Id(Required) */
func (r *DisableWhiteListRuleOfForwardRuleRequest) SetInstanceId(instanceId int) {
    r.InstanceId = instanceId
}

/* param forwardRuleId: 转发规则 Id(Required) */
func (r *DisableWhiteListRuleOfForwardRuleRequest) SetForwardRuleId(forwardRuleId int) {
    r.ForwardRuleId = forwardRuleId
}

/* param ipRuleId: 待引用的 IP 黑白名单规则 Id(Required) */
func (r *DisableWhiteListRuleOfForwardRuleRequest) SetIpRuleId(ipRuleId string) {
    r.IpRuleId = ipRuleId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DisableWhiteListRuleOfForwardRuleRequest) GetRegionId() string {
    return r.RegionId
}

type DisableWhiteListRuleOfForwardRuleResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DisableWhiteListRuleOfForwardRuleResult `json:"result"`
}

type DisableWhiteListRuleOfForwardRuleResult struct {
    Code int `json:"code"`
    Message string `json:"message"`
}