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

type AddNetworkAclRulesRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* networkAclId ID  */
    NetworkAclId string `json:"networkAclId"`

    /* networkAcl规则列表  */
    NetworkAclRuleSpecs []vpc.AddNetworkAclRuleSpec `json:"networkAclRuleSpecs"`
}

/*
 * param regionId: Region ID (Required)
 * param networkAclId: networkAclId ID (Required)
 * param networkAclRuleSpecs: networkAcl规则列表 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewAddNetworkAclRulesRequest(
    regionId string,
    networkAclId string,
    networkAclRuleSpecs []vpc.AddNetworkAclRuleSpec,
) *AddNetworkAclRulesRequest {

	return &AddNetworkAclRulesRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/networkAcls/{networkAclId}:addNetworkAclRules",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        NetworkAclId: networkAclId,
        NetworkAclRuleSpecs: networkAclRuleSpecs,
	}
}

/*
 * param regionId: Region ID (Required)
 * param networkAclId: networkAclId ID (Required)
 * param networkAclRuleSpecs: networkAcl规则列表 (Required)
 */
func NewAddNetworkAclRulesRequestWithAllParams(
    regionId string,
    networkAclId string,
    networkAclRuleSpecs []vpc.AddNetworkAclRuleSpec,
) *AddNetworkAclRulesRequest {

    return &AddNetworkAclRulesRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/networkAcls/{networkAclId}:addNetworkAclRules",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        NetworkAclId: networkAclId,
        NetworkAclRuleSpecs: networkAclRuleSpecs,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewAddNetworkAclRulesRequestWithoutParam() *AddNetworkAclRulesRequest {

    return &AddNetworkAclRulesRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/networkAcls/{networkAclId}:addNetworkAclRules",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *AddNetworkAclRulesRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param networkAclId: networkAclId ID(Required) */
func (r *AddNetworkAclRulesRequest) SetNetworkAclId(networkAclId string) {
    r.NetworkAclId = networkAclId
}

/* param networkAclRuleSpecs: networkAcl规则列表(Required) */
func (r *AddNetworkAclRulesRequest) SetNetworkAclRuleSpecs(networkAclRuleSpecs []vpc.AddNetworkAclRuleSpec) {
    r.NetworkAclRuleSpecs = networkAclRuleSpecs
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r AddNetworkAclRulesRequest) GetRegionId() string {
    return r.RegionId
}

type AddNetworkAclRulesResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result AddNetworkAclRulesResult `json:"result"`
}

type AddNetworkAclRulesResult struct {
}