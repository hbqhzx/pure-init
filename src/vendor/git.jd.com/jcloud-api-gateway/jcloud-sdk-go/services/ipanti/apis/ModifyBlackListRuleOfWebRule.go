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

type ModifyBlackListRuleOfWebRuleRequest struct {

    core.JDCloudRequest

    /* 区域 Id  */
    RegionId string `json:"regionId"`

    /* 高防实例 Id  */
    InstanceId int `json:"instanceId"`

    /* 网站规则 Id  */
    WebRuleId int `json:"webRuleId"`

    /* 网站类规则的黑名单规则 Id  */
    WebBlackListRuleId string `json:"webBlackListRuleId"`
}

/*
 * param regionId: 区域 Id (Required)
 * param instanceId: 高防实例 Id (Required)
 * param webRuleId: 网站规则 Id (Required)
 * param webBlackListRuleId: 网站类规则的黑名单规则 Id (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewModifyBlackListRuleOfWebRuleRequest(
    regionId string,
    instanceId int,
    webRuleId int,
    webBlackListRuleId string,
) *ModifyBlackListRuleOfWebRuleRequest {

	return &ModifyBlackListRuleOfWebRuleRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}/webRules/{webRuleId}/webBlackListRules/{webBlackListRuleId}",
			Method:  "PATCH",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceId: instanceId,
        WebRuleId: webRuleId,
        WebBlackListRuleId: webBlackListRuleId,
	}
}

/*
 * param regionId: 区域 Id (Required)
 * param instanceId: 高防实例 Id (Required)
 * param webRuleId: 网站规则 Id (Required)
 * param webBlackListRuleId: 网站类规则的黑名单规则 Id (Required)
 */
func NewModifyBlackListRuleOfWebRuleRequestWithAllParams(
    regionId string,
    instanceId int,
    webRuleId int,
    webBlackListRuleId string,
) *ModifyBlackListRuleOfWebRuleRequest {

    return &ModifyBlackListRuleOfWebRuleRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/webRules/{webRuleId}/webBlackListRules/{webBlackListRuleId}",
            Method:  "PATCH",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InstanceId: instanceId,
        WebRuleId: webRuleId,
        WebBlackListRuleId: webBlackListRuleId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewModifyBlackListRuleOfWebRuleRequestWithoutParam() *ModifyBlackListRuleOfWebRuleRequest {

    return &ModifyBlackListRuleOfWebRuleRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/webRules/{webRuleId}/webBlackListRules/{webBlackListRuleId}",
            Method:  "PATCH",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 区域 Id(Required) */
func (r *ModifyBlackListRuleOfWebRuleRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceId: 高防实例 Id(Required) */
func (r *ModifyBlackListRuleOfWebRuleRequest) SetInstanceId(instanceId int) {
    r.InstanceId = instanceId
}

/* param webRuleId: 网站规则 Id(Required) */
func (r *ModifyBlackListRuleOfWebRuleRequest) SetWebRuleId(webRuleId int) {
    r.WebRuleId = webRuleId
}

/* param webBlackListRuleId: 网站类规则的黑名单规则 Id(Required) */
func (r *ModifyBlackListRuleOfWebRuleRequest) SetWebBlackListRuleId(webBlackListRuleId string) {
    r.WebBlackListRuleId = webBlackListRuleId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ModifyBlackListRuleOfWebRuleRequest) GetRegionId() string {
    return r.RegionId
}

type ModifyBlackListRuleOfWebRuleResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ModifyBlackListRuleOfWebRuleResult `json:"result"`
}

type ModifyBlackListRuleOfWebRuleResult struct {
    Code int `json:"code"`
    Message string `json:"message"`
}