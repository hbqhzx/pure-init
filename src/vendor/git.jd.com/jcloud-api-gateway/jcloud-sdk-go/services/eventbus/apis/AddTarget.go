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
    eventbus "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/eventbus/models"
)

type AddTargetRequest struct {

    core.JDCloudRequest

    /* 地域 Id  */
    RegionId string `json:"regionId"`

    /* 幂等性校验参数,最长36位  */
    ClientToken string `json:"clientToken"`

    /* 路由目的地对应的规则唯一标识  */
    RuleId string `json:"ruleId"`

    /* 新增规则事件路由目的地  */
    Targets []eventbus.TargetParam `json:"targets"`
}

/*
 * param regionId: 地域 Id (Required)
 * param clientToken: 幂等性校验参数,最长36位 (Required)
 * param ruleId: 路由目的地对应的规则唯一标识 (Required)
 * param targets: 新增规则事件路由目的地 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewAddTargetRequest(
    regionId string,
    clientToken string,
    ruleId string,
    targets []eventbus.TargetParam,
) *AddTargetRequest {

	return &AddTargetRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/target",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        ClientToken: clientToken,
        RuleId: ruleId,
        Targets: targets,
	}
}

/*
 * param regionId: 地域 Id (Required)
 * param clientToken: 幂等性校验参数,最长36位 (Required)
 * param ruleId: 路由目的地对应的规则唯一标识 (Required)
 * param targets: 新增规则事件路由目的地 (Required)
 */
func NewAddTargetRequestWithAllParams(
    regionId string,
    clientToken string,
    ruleId string,
    targets []eventbus.TargetParam,
) *AddTargetRequest {

    return &AddTargetRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/target",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        ClientToken: clientToken,
        RuleId: ruleId,
        Targets: targets,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewAddTargetRequestWithoutParam() *AddTargetRequest {

    return &AddTargetRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/target",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域 Id(Required) */
func (r *AddTargetRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param clientToken: 幂等性校验参数,最长36位(Required) */
func (r *AddTargetRequest) SetClientToken(clientToken string) {
    r.ClientToken = clientToken
}

/* param ruleId: 路由目的地对应的规则唯一标识(Required) */
func (r *AddTargetRequest) SetRuleId(ruleId string) {
    r.RuleId = ruleId
}

/* param targets: 新增规则事件路由目的地(Required) */
func (r *AddTargetRequest) SetTargets(targets []eventbus.TargetParam) {
    r.Targets = targets
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r AddTargetRequest) GetRegionId() string {
    return r.RegionId
}

type AddTargetResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result AddTargetResult `json:"result"`
}

type AddTargetResult struct {
    Success bool `json:"success"`
    Uuids []string `json:"uuids"`
}