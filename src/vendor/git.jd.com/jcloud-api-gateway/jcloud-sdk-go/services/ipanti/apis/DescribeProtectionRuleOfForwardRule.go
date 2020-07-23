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
    ipanti "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/ipanti/models"
)

type DescribeProtectionRuleOfForwardRuleRequest struct {

    core.JDCloudRequest

    /* 区域 Id  */
    RegionId string `json:"regionId"`

    /* 高防实例 Id  */
    InstanceId int `json:"instanceId"`

    /* 转发规则 Id  */
    ForwardRuleId int `json:"forwardRuleId"`
}

/*
 * param regionId: 区域 Id (Required)
 * param instanceId: 高防实例 Id (Required)
 * param forwardRuleId: 转发规则 Id (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeProtectionRuleOfForwardRuleRequest(
    regionId string,
    instanceId int,
    forwardRuleId int,
) *DescribeProtectionRuleOfForwardRuleRequest {

	return &DescribeProtectionRuleOfForwardRuleRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}/forwardRules/{forwardRuleId}:describeProtectionRule",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceId: instanceId,
        ForwardRuleId: forwardRuleId,
	}
}

/*
 * param regionId: 区域 Id (Required)
 * param instanceId: 高防实例 Id (Required)
 * param forwardRuleId: 转发规则 Id (Required)
 */
func NewDescribeProtectionRuleOfForwardRuleRequestWithAllParams(
    regionId string,
    instanceId int,
    forwardRuleId int,
) *DescribeProtectionRuleOfForwardRuleRequest {

    return &DescribeProtectionRuleOfForwardRuleRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/forwardRules/{forwardRuleId}:describeProtectionRule",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InstanceId: instanceId,
        ForwardRuleId: forwardRuleId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeProtectionRuleOfForwardRuleRequestWithoutParam() *DescribeProtectionRuleOfForwardRuleRequest {

    return &DescribeProtectionRuleOfForwardRuleRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/forwardRules/{forwardRuleId}:describeProtectionRule",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 区域 Id(Required) */
func (r *DescribeProtectionRuleOfForwardRuleRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceId: 高防实例 Id(Required) */
func (r *DescribeProtectionRuleOfForwardRuleRequest) SetInstanceId(instanceId int) {
    r.InstanceId = instanceId
}

/* param forwardRuleId: 转发规则 Id(Required) */
func (r *DescribeProtectionRuleOfForwardRuleRequest) SetForwardRuleId(forwardRuleId int) {
    r.ForwardRuleId = forwardRuleId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeProtectionRuleOfForwardRuleRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeProtectionRuleOfForwardRuleResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeProtectionRuleOfForwardRuleResult `json:"result"`
}

type DescribeProtectionRuleOfForwardRuleResult struct {
    ProtectionRule ipanti.ForwardProtectionRule `json:"protectionRule"`
}