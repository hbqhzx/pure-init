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

type CreateAppConfigVersionRequest struct {

    core.JDCloudRequest

    /* 可用区id  */
    RegionId string `json:"regionId"`

    /* 配置id  */
    AppConfigId string `json:"appConfigId"`

    /* 配置内容  */
    AppConfigContent string `json:"appConfigContent"`

    /* 配置版本号  */
    AppConfigVersion int `json:"appConfigVersion"`

    /* 配置父版本号 (Optional) */
    AppConfigParentVersionId *int64 `json:"appConfigParentVersionId"`

    /* 版本备注信息 (Optional) */
    AppConfigRemark *string `json:"appConfigRemark"`
}

/*
 * param regionId: 可用区id (Required)
 * param appConfigId: 配置id (Required)
 * param appConfigContent: 配置内容 (Required)
 * param appConfigVersion: 配置版本号 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateAppConfigVersionRequest(
    regionId string,
    appConfigId string,
    appConfigContent string,
    appConfigVersion int,
) *CreateAppConfigVersionRequest {

	return &CreateAppConfigVersionRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/appconfig/{appConfigId}/versions",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        AppConfigId: appConfigId,
        AppConfigContent: appConfigContent,
        AppConfigVersion: appConfigVersion,
	}
}

/*
 * param regionId: 可用区id (Required)
 * param appConfigId: 配置id (Required)
 * param appConfigContent: 配置内容 (Required)
 * param appConfigVersion: 配置版本号 (Required)
 * param appConfigParentVersionId: 配置父版本号 (Optional)
 * param appConfigRemark: 版本备注信息 (Optional)
 */
func NewCreateAppConfigVersionRequestWithAllParams(
    regionId string,
    appConfigId string,
    appConfigContent string,
    appConfigVersion int,
    appConfigParentVersionId *int64,
    appConfigRemark *string,
) *CreateAppConfigVersionRequest {

    return &CreateAppConfigVersionRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/appconfig/{appConfigId}/versions",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        AppConfigId: appConfigId,
        AppConfigContent: appConfigContent,
        AppConfigVersion: appConfigVersion,
        AppConfigParentVersionId: appConfigParentVersionId,
        AppConfigRemark: appConfigRemark,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateAppConfigVersionRequestWithoutParam() *CreateAppConfigVersionRequest {

    return &CreateAppConfigVersionRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/appconfig/{appConfigId}/versions",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 可用区id(Required) */
func (r *CreateAppConfigVersionRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param appConfigId: 配置id(Required) */
func (r *CreateAppConfigVersionRequest) SetAppConfigId(appConfigId string) {
    r.AppConfigId = appConfigId
}

/* param appConfigContent: 配置内容(Required) */
func (r *CreateAppConfigVersionRequest) SetAppConfigContent(appConfigContent string) {
    r.AppConfigContent = appConfigContent
}

/* param appConfigVersion: 配置版本号(Required) */
func (r *CreateAppConfigVersionRequest) SetAppConfigVersion(appConfigVersion int) {
    r.AppConfigVersion = appConfigVersion
}

/* param appConfigParentVersionId: 配置父版本号(Optional) */
func (r *CreateAppConfigVersionRequest) SetAppConfigParentVersionId(appConfigParentVersionId int64) {
    r.AppConfigParentVersionId = &appConfigParentVersionId
}

/* param appConfigRemark: 版本备注信息(Optional) */
func (r *CreateAppConfigVersionRequest) SetAppConfigRemark(appConfigRemark string) {
    r.AppConfigRemark = &appConfigRemark
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateAppConfigVersionRequest) GetRegionId() string {
    return r.RegionId
}

type CreateAppConfigVersionResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateAppConfigVersionResult `json:"result"`
}

type CreateAppConfigVersionResult struct {
    CreateResult string `json:"createResult"`
    AppConfigVersionId int64 `json:"appConfigVersionId"`
}