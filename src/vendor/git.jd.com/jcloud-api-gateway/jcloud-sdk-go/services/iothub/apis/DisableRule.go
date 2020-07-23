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

type DisableRuleRequest struct {

    core.JDCloudRequest

    /* 规则Id (Optional) */
    RuleId *string `json:"ruleId"`

    /* 实例Id (Optional) */
    InstanceId *string `json:"instanceId"`
}

/*
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDisableRuleRequest(
) *DisableRuleRequest {

	return &DisableRuleRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/rule:disable",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
	}
}

/*
 * param ruleId: 规则Id (Optional)
 * param instanceId: 实例Id (Optional)
 */
func NewDisableRuleRequestWithAllParams(
    ruleId *string,
    instanceId *string,
) *DisableRuleRequest {

    return &DisableRuleRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/rule:disable",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RuleId: ruleId,
        InstanceId: instanceId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDisableRuleRequestWithoutParam() *DisableRuleRequest {

    return &DisableRuleRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/rule:disable",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param ruleId: 规则Id(Optional) */
func (r *DisableRuleRequest) SetRuleId(ruleId string) {
    r.RuleId = &ruleId
}

/* param instanceId: 实例Id(Optional) */
func (r *DisableRuleRequest) SetInstanceId(instanceId string) {
    r.InstanceId = &instanceId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DisableRuleRequest) GetRegionId() string {
    return ""
}

type DisableRuleResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DisableRuleResult `json:"result"`
}

type DisableRuleResult struct {
}