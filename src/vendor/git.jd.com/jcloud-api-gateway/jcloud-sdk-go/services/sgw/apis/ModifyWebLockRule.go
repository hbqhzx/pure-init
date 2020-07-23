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
    sgw "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/sgw/models"
)

type ModifyWebLockRuleRequest struct {

    core.JDCloudRequest

    /* 地域 Id  */
    RegionId string `json:"regionId"`

    /* waf 实例id  */
    InstanceId string `json:"instanceId"`

    /* 网页防篡改规则id  */
    RuleId string `json:"ruleId"`

    /* 创建防护配置规则所需基本信息。  */
    WeblockRule *sgw.WebLockRule `json:"weblockRule"`
}

/*
 * param regionId: 地域 Id (Required)
 * param instanceId: waf 实例id (Required)
 * param ruleId: 网页防篡改规则id (Required)
 * param weblockRule: 创建防护配置规则所需基本信息。 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewModifyWebLockRuleRequest(
    regionId string,
    instanceId string,
    ruleId string,
    weblockRule *sgw.WebLockRule,
) *ModifyWebLockRuleRequest {

	return &ModifyWebLockRuleRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}/weblockrules/{ruleId}",
			Method:  "PATCH",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceId: instanceId,
        RuleId: ruleId,
        WeblockRule: weblockRule,
	}
}

/*
 * param regionId: 地域 Id (Required)
 * param instanceId: waf 实例id (Required)
 * param ruleId: 网页防篡改规则id (Required)
 * param weblockRule: 创建防护配置规则所需基本信息。 (Required)
 */
func NewModifyWebLockRuleRequestWithAllParams(
    regionId string,
    instanceId string,
    ruleId string,
    weblockRule *sgw.WebLockRule,
) *ModifyWebLockRuleRequest {

    return &ModifyWebLockRuleRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/weblockrules/{ruleId}",
            Method:  "PATCH",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InstanceId: instanceId,
        RuleId: ruleId,
        WeblockRule: weblockRule,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewModifyWebLockRuleRequestWithoutParam() *ModifyWebLockRuleRequest {

    return &ModifyWebLockRuleRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/weblockrules/{ruleId}",
            Method:  "PATCH",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域 Id(Required) */
func (r *ModifyWebLockRuleRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceId: waf 实例id(Required) */
func (r *ModifyWebLockRuleRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

/* param ruleId: 网页防篡改规则id(Required) */
func (r *ModifyWebLockRuleRequest) SetRuleId(ruleId string) {
    r.RuleId = ruleId
}

/* param weblockRule: 创建防护配置规则所需基本信息。(Required) */
func (r *ModifyWebLockRuleRequest) SetWeblockRule(weblockRule *sgw.WebLockRule) {
    r.WeblockRule = weblockRule
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ModifyWebLockRuleRequest) GetRegionId() string {
    return r.RegionId
}

type ModifyWebLockRuleResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ModifyWebLockRuleResult `json:"result"`
}

type ModifyWebLockRuleResult struct {
}