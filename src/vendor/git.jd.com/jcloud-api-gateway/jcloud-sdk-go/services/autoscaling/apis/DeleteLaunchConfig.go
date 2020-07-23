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

type DeleteLaunchConfigRequest struct {

    core.JDCloudRequest

    /* 地域 Id  */
    RegionId string `json:"regionId"`

    /* 启动配置uuid  */
    LaunchConfigUuid string `json:"launchConfigUuid"`
}

/*
 * param regionId: 地域 Id (Required)
 * param launchConfigUuid: 启动配置uuid (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDeleteLaunchConfigRequest(
    regionId string,
    launchConfigUuid string,
) *DeleteLaunchConfigRequest {

	return &DeleteLaunchConfigRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/launchConfigs/{launchConfigUuid}",
			Method:  "DELETE",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        LaunchConfigUuid: launchConfigUuid,
	}
}

/*
 * param regionId: 地域 Id (Required)
 * param launchConfigUuid: 启动配置uuid (Required)
 */
func NewDeleteLaunchConfigRequestWithAllParams(
    regionId string,
    launchConfigUuid string,
) *DeleteLaunchConfigRequest {

    return &DeleteLaunchConfigRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/launchConfigs/{launchConfigUuid}",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        LaunchConfigUuid: launchConfigUuid,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDeleteLaunchConfigRequestWithoutParam() *DeleteLaunchConfigRequest {

    return &DeleteLaunchConfigRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/launchConfigs/{launchConfigUuid}",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域 Id(Required) */
func (r *DeleteLaunchConfigRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param launchConfigUuid: 启动配置uuid(Required) */
func (r *DeleteLaunchConfigRequest) SetLaunchConfigUuid(launchConfigUuid string) {
    r.LaunchConfigUuid = launchConfigUuid
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DeleteLaunchConfigRequest) GetRegionId() string {
    return r.RegionId
}

type DeleteLaunchConfigResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DeleteLaunchConfigResult `json:"result"`
}

type DeleteLaunchConfigResult struct {
}