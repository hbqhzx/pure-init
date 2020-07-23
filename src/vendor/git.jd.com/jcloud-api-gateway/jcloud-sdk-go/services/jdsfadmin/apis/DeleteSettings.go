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

type DeleteSettingsRequest struct {

    core.JDCloudRequest

    /* 可用区id  */
    RegionId string `json:"regionId"`

    /* 设置的key  */
    SettingKey string `json:"settingKey"`
}

/*
 * param regionId: 可用区id (Required)
 * param settingKey: 设置的key (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDeleteSettingsRequest(
    regionId string,
    settingKey string,
) *DeleteSettingsRequest {

	return &DeleteSettingsRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/settings/{settingKey}",
			Method:  "DELETE",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        SettingKey: settingKey,
	}
}

/*
 * param regionId: 可用区id (Required)
 * param settingKey: 设置的key (Required)
 */
func NewDeleteSettingsRequestWithAllParams(
    regionId string,
    settingKey string,
) *DeleteSettingsRequest {

    return &DeleteSettingsRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/settings/{settingKey}",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        SettingKey: settingKey,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDeleteSettingsRequestWithoutParam() *DeleteSettingsRequest {

    return &DeleteSettingsRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/settings/{settingKey}",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 可用区id(Required) */
func (r *DeleteSettingsRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param settingKey: 设置的key(Required) */
func (r *DeleteSettingsRequest) SetSettingKey(settingKey string) {
    r.SettingKey = settingKey
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DeleteSettingsRequest) GetRegionId() string {
    return r.RegionId
}

type DeleteSettingsResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DeleteSettingsResult `json:"result"`
}

type DeleteSettingsResult struct {
    DeleteSettingResult string `json:"deleteSettingResult"`
}