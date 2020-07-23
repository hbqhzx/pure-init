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
    ucapi "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/ucapi/models"
)

type CreateSensitiveOpSettingRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 设置信息 (Optional) */
    Setting *ucapi.SensitiveOpSettingParam `json:"setting"`
}

/*
 * param regionId: Region ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateSensitiveOpSettingRequest(
    regionId string,
) *CreateSensitiveOpSettingRequest {

	return &CreateSensitiveOpSettingRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/sensitiveOpSetting",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: Region ID (Required)
 * param setting: 设置信息 (Optional)
 */
func NewCreateSensitiveOpSettingRequestWithAllParams(
    regionId string,
    setting *ucapi.SensitiveOpSettingParam,
) *CreateSensitiveOpSettingRequest {

    return &CreateSensitiveOpSettingRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/sensitiveOpSetting",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        Setting: setting,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateSensitiveOpSettingRequestWithoutParam() *CreateSensitiveOpSettingRequest {

    return &CreateSensitiveOpSettingRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/sensitiveOpSetting",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *CreateSensitiveOpSettingRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param setting: 设置信息(Optional) */
func (r *CreateSensitiveOpSettingRequest) SetSetting(setting *ucapi.SensitiveOpSettingParam) {
    r.Setting = setting
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateSensitiveOpSettingRequest) GetRegionId() string {
    return r.RegionId
}

type CreateSensitiveOpSettingResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateSensitiveOpSettingResult `json:"result"`
}

type CreateSensitiveOpSettingResult struct {
}