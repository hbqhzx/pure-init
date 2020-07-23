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
    vpcwaf "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/vpcwaf/models"
)

type ModifyCCEngineConfigRequest struct {

    core.JDCloudRequest

    /* 地域 Id  */
    RegionId string `json:"regionId"`

    /* waf 实例id  */
    InstanceId string `json:"instanceId"`

    /* 新的ccEngine配置信息  */
    CcEngineConfig *vpcwaf.CcEngineConfig `json:"ccEngineConfig"`
}

/*
 * param regionId: 地域 Id (Required)
 * param instanceId: waf 实例id (Required)
 * param ccEngineConfig: 新的ccEngine配置信息 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewModifyCCEngineConfigRequest(
    regionId string,
    instanceId string,
    ccEngineConfig *vpcwaf.CcEngineConfig,
) *ModifyCCEngineConfigRequest {

	return &ModifyCCEngineConfigRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}:modifyCCEngineConfig",
			Method:  "PATCH",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceId: instanceId,
        CcEngineConfig: ccEngineConfig,
	}
}

/*
 * param regionId: 地域 Id (Required)
 * param instanceId: waf 实例id (Required)
 * param ccEngineConfig: 新的ccEngine配置信息 (Required)
 */
func NewModifyCCEngineConfigRequestWithAllParams(
    regionId string,
    instanceId string,
    ccEngineConfig *vpcwaf.CcEngineConfig,
) *ModifyCCEngineConfigRequest {

    return &ModifyCCEngineConfigRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}:modifyCCEngineConfig",
            Method:  "PATCH",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InstanceId: instanceId,
        CcEngineConfig: ccEngineConfig,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewModifyCCEngineConfigRequestWithoutParam() *ModifyCCEngineConfigRequest {

    return &ModifyCCEngineConfigRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}:modifyCCEngineConfig",
            Method:  "PATCH",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域 Id(Required) */
func (r *ModifyCCEngineConfigRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceId: waf 实例id(Required) */
func (r *ModifyCCEngineConfigRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

/* param ccEngineConfig: 新的ccEngine配置信息(Required) */
func (r *ModifyCCEngineConfigRequest) SetCcEngineConfig(ccEngineConfig *vpcwaf.CcEngineConfig) {
    r.CcEngineConfig = ccEngineConfig
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ModifyCCEngineConfigRequest) GetRegionId() string {
    return r.RegionId
}

type ModifyCCEngineConfigResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ModifyCCEngineConfigResult `json:"result"`
}

type ModifyCCEngineConfigResult struct {
}