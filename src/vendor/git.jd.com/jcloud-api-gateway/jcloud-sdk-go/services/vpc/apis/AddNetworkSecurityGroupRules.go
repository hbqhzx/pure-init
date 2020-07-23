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
    vpc "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/vpc/models"
)

type AddNetworkSecurityGroupRulesRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* NetworkSecurityGroup ID  */
    NetworkSecurityGroupId string `json:"networkSecurityGroupId"`

    /* 安全组规则信息  */
    NetworkSecurityGroupRuleSpecs []vpc.AddSecurityGroupRules `json:"networkSecurityGroupRuleSpecs"`
}

/*
 * param regionId: Region ID (Required)
 * param networkSecurityGroupId: NetworkSecurityGroup ID (Required)
 * param networkSecurityGroupRuleSpecs: 安全组规则信息 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewAddNetworkSecurityGroupRulesRequest(
    regionId string,
    networkSecurityGroupId string,
    networkSecurityGroupRuleSpecs []vpc.AddSecurityGroupRules,
) *AddNetworkSecurityGroupRulesRequest {

	return &AddNetworkSecurityGroupRulesRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/networkSecurityGroups/{networkSecurityGroupId}:addNetworkSecurityGroupRules",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        NetworkSecurityGroupId: networkSecurityGroupId,
        NetworkSecurityGroupRuleSpecs: networkSecurityGroupRuleSpecs,
	}
}

/*
 * param regionId: Region ID (Required)
 * param networkSecurityGroupId: NetworkSecurityGroup ID (Required)
 * param networkSecurityGroupRuleSpecs: 安全组规则信息 (Required)
 */
func NewAddNetworkSecurityGroupRulesRequestWithAllParams(
    regionId string,
    networkSecurityGroupId string,
    networkSecurityGroupRuleSpecs []vpc.AddSecurityGroupRules,
) *AddNetworkSecurityGroupRulesRequest {

    return &AddNetworkSecurityGroupRulesRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/networkSecurityGroups/{networkSecurityGroupId}:addNetworkSecurityGroupRules",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        NetworkSecurityGroupId: networkSecurityGroupId,
        NetworkSecurityGroupRuleSpecs: networkSecurityGroupRuleSpecs,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewAddNetworkSecurityGroupRulesRequestWithoutParam() *AddNetworkSecurityGroupRulesRequest {

    return &AddNetworkSecurityGroupRulesRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/networkSecurityGroups/{networkSecurityGroupId}:addNetworkSecurityGroupRules",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *AddNetworkSecurityGroupRulesRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param networkSecurityGroupId: NetworkSecurityGroup ID(Required) */
func (r *AddNetworkSecurityGroupRulesRequest) SetNetworkSecurityGroupId(networkSecurityGroupId string) {
    r.NetworkSecurityGroupId = networkSecurityGroupId
}

/* param networkSecurityGroupRuleSpecs: 安全组规则信息(Required) */
func (r *AddNetworkSecurityGroupRulesRequest) SetNetworkSecurityGroupRuleSpecs(networkSecurityGroupRuleSpecs []vpc.AddSecurityGroupRules) {
    r.NetworkSecurityGroupRuleSpecs = networkSecurityGroupRuleSpecs
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r AddNetworkSecurityGroupRulesRequest) GetRegionId() string {
    return r.RegionId
}

type AddNetworkSecurityGroupRulesResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result AddNetworkSecurityGroupRulesResult `json:"result"`
}

type AddNetworkSecurityGroupRulesResult struct {
}