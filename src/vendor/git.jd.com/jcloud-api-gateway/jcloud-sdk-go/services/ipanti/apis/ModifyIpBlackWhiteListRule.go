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

type ModifyIpBlackWhiteListRuleRequest struct {

    core.JDCloudRequest

    /* 区域 Id  */
    RegionId string `json:"regionId"`

    /* 高防实例 Id  */
    InstanceId int `json:"instanceId"`

    /* IP 黑白名单规则 Id  */
    IpRuleId string `json:"ipRuleId"`

    /* 修改实例的 IP 黑/白名单规则请求参数  */
    IpBlackWhiteListRuleSpec *ipanti.IpBlackWhiteListRuleSpec `json:"ipBlackWhiteListRuleSpec"`
}

/*
 * param regionId: 区域 Id (Required)
 * param instanceId: 高防实例 Id (Required)
 * param ipRuleId: IP 黑白名单规则 Id (Required)
 * param ipBlackWhiteListRuleSpec: 修改实例的 IP 黑/白名单规则请求参数 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewModifyIpBlackWhiteListRuleRequest(
    regionId string,
    instanceId int,
    ipRuleId string,
    ipBlackWhiteListRuleSpec *ipanti.IpBlackWhiteListRuleSpec,
) *ModifyIpBlackWhiteListRuleRequest {

	return &ModifyIpBlackWhiteListRuleRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}/IpBlackWhiteListRules/{ipRuleId}",
			Method:  "PATCH",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceId: instanceId,
        IpRuleId: ipRuleId,
        IpBlackWhiteListRuleSpec: ipBlackWhiteListRuleSpec,
	}
}

/*
 * param regionId: 区域 Id (Required)
 * param instanceId: 高防实例 Id (Required)
 * param ipRuleId: IP 黑白名单规则 Id (Required)
 * param ipBlackWhiteListRuleSpec: 修改实例的 IP 黑/白名单规则请求参数 (Required)
 */
func NewModifyIpBlackWhiteListRuleRequestWithAllParams(
    regionId string,
    instanceId int,
    ipRuleId string,
    ipBlackWhiteListRuleSpec *ipanti.IpBlackWhiteListRuleSpec,
) *ModifyIpBlackWhiteListRuleRequest {

    return &ModifyIpBlackWhiteListRuleRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/IpBlackWhiteListRules/{ipRuleId}",
            Method:  "PATCH",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InstanceId: instanceId,
        IpRuleId: ipRuleId,
        IpBlackWhiteListRuleSpec: ipBlackWhiteListRuleSpec,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewModifyIpBlackWhiteListRuleRequestWithoutParam() *ModifyIpBlackWhiteListRuleRequest {

    return &ModifyIpBlackWhiteListRuleRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/IpBlackWhiteListRules/{ipRuleId}",
            Method:  "PATCH",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 区域 Id(Required) */
func (r *ModifyIpBlackWhiteListRuleRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceId: 高防实例 Id(Required) */
func (r *ModifyIpBlackWhiteListRuleRequest) SetInstanceId(instanceId int) {
    r.InstanceId = instanceId
}

/* param ipRuleId: IP 黑白名单规则 Id(Required) */
func (r *ModifyIpBlackWhiteListRuleRequest) SetIpRuleId(ipRuleId string) {
    r.IpRuleId = ipRuleId
}

/* param ipBlackWhiteListRuleSpec: 修改实例的 IP 黑/白名单规则请求参数(Required) */
func (r *ModifyIpBlackWhiteListRuleRequest) SetIpBlackWhiteListRuleSpec(ipBlackWhiteListRuleSpec *ipanti.IpBlackWhiteListRuleSpec) {
    r.IpBlackWhiteListRuleSpec = ipBlackWhiteListRuleSpec
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ModifyIpBlackWhiteListRuleRequest) GetRegionId() string {
    return r.RegionId
}

type ModifyIpBlackWhiteListRuleResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ModifyIpBlackWhiteListRuleResult `json:"result"`
}

type ModifyIpBlackWhiteListRuleResult struct {
    Code int `json:"code"`
    Message string `json:"message"`
}