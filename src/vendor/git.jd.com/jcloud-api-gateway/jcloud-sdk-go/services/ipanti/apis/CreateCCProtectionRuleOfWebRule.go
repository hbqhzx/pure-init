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

type CreateCCProtectionRuleOfWebRuleRequest struct {

    core.JDCloudRequest

    /* 区域 Id  */
    RegionId string `json:"regionId"`

    /* 高防实例 Id  */
    InstanceId int `json:"instanceId"`

    /* 网站规则 Id  */
    WebRuleId int `json:"webRuleId"`

    /* 添加 CC 防护规则请求参数  */
    CcProtectionRuleSpec *ipanti.CCProtectionRuleSpec `json:"ccProtectionRuleSpec"`
}

/*
 * param regionId: 区域 Id (Required)
 * param instanceId: 高防实例 Id (Required)
 * param webRuleId: 网站规则 Id (Required)
 * param ccProtectionRuleSpec: 添加 CC 防护规则请求参数 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateCCProtectionRuleOfWebRuleRequest(
    regionId string,
    instanceId int,
    webRuleId int,
    ccProtectionRuleSpec *ipanti.CCProtectionRuleSpec,
) *CreateCCProtectionRuleOfWebRuleRequest {

	return &CreateCCProtectionRuleOfWebRuleRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}/webRules/{webRuleId}/ccProtectionRules",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceId: instanceId,
        WebRuleId: webRuleId,
        CcProtectionRuleSpec: ccProtectionRuleSpec,
	}
}

/*
 * param regionId: 区域 Id (Required)
 * param instanceId: 高防实例 Id (Required)
 * param webRuleId: 网站规则 Id (Required)
 * param ccProtectionRuleSpec: 添加 CC 防护规则请求参数 (Required)
 */
func NewCreateCCProtectionRuleOfWebRuleRequestWithAllParams(
    regionId string,
    instanceId int,
    webRuleId int,
    ccProtectionRuleSpec *ipanti.CCProtectionRuleSpec,
) *CreateCCProtectionRuleOfWebRuleRequest {

    return &CreateCCProtectionRuleOfWebRuleRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/webRules/{webRuleId}/ccProtectionRules",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InstanceId: instanceId,
        WebRuleId: webRuleId,
        CcProtectionRuleSpec: ccProtectionRuleSpec,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateCCProtectionRuleOfWebRuleRequestWithoutParam() *CreateCCProtectionRuleOfWebRuleRequest {

    return &CreateCCProtectionRuleOfWebRuleRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/webRules/{webRuleId}/ccProtectionRules",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 区域 Id(Required) */
func (r *CreateCCProtectionRuleOfWebRuleRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceId: 高防实例 Id(Required) */
func (r *CreateCCProtectionRuleOfWebRuleRequest) SetInstanceId(instanceId int) {
    r.InstanceId = instanceId
}

/* param webRuleId: 网站规则 Id(Required) */
func (r *CreateCCProtectionRuleOfWebRuleRequest) SetWebRuleId(webRuleId int) {
    r.WebRuleId = webRuleId
}

/* param ccProtectionRuleSpec: 添加 CC 防护规则请求参数(Required) */
func (r *CreateCCProtectionRuleOfWebRuleRequest) SetCcProtectionRuleSpec(ccProtectionRuleSpec *ipanti.CCProtectionRuleSpec) {
    r.CcProtectionRuleSpec = ccProtectionRuleSpec
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateCCProtectionRuleOfWebRuleRequest) GetRegionId() string {
    return r.RegionId
}

type CreateCCProtectionRuleOfWebRuleResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateCCProtectionRuleOfWebRuleResult `json:"result"`
}

type CreateCCProtectionRuleOfWebRuleResult struct {
    Code int `json:"code"`
    Message string `json:"message"`
}