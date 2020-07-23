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

type DeleteAppConfigRequest struct {

    core.JDCloudRequest

    /* 可用区id  */
    RegionId string `json:"regionId"`

    /* 要删除的配置 id  */
    AppConfigId int `json:"appConfigId"`
}

/*
 * param regionId: 可用区id (Required)
 * param appConfigId: 要删除的配置 id (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDeleteAppConfigRequest(
    regionId string,
    appConfigId int,
) *DeleteAppConfigRequest {

	return &DeleteAppConfigRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/appconfig/{appConfigId}",
			Method:  "DELETE",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        AppConfigId: appConfigId,
	}
}

/*
 * param regionId: 可用区id (Required)
 * param appConfigId: 要删除的配置 id (Required)
 */
func NewDeleteAppConfigRequestWithAllParams(
    regionId string,
    appConfigId int,
) *DeleteAppConfigRequest {

    return &DeleteAppConfigRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/appconfig/{appConfigId}",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        AppConfigId: appConfigId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDeleteAppConfigRequestWithoutParam() *DeleteAppConfigRequest {

    return &DeleteAppConfigRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/appconfig/{appConfigId}",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 可用区id(Required) */
func (r *DeleteAppConfigRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param appConfigId: 要删除的配置 id(Required) */
func (r *DeleteAppConfigRequest) SetAppConfigId(appConfigId int) {
    r.AppConfigId = appConfigId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DeleteAppConfigRequest) GetRegionId() string {
    return r.RegionId
}

type DeleteAppConfigResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DeleteAppConfigResult `json:"result"`
}

type DeleteAppConfigResult struct {
    DeleteResult string `json:"deleteResult"`
}