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

type CreateAppRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 应用名称  */
    AppName string `json:"appName"`

    /* 部署平台  */
    Platform int `json:"platform"`

    /* 使用分布式服务框架：0不使用，1使用 (Optional) */
    JdsfEnabled *int `json:"jdsfEnabled"`

    /* 描述 (Optional) */
    Desc *string `json:"desc"`
}

/*
 * param regionId: 地域ID (Required)
 * param appName: 应用名称 (Required)
 * param platform: 部署平台 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateAppRequest(
    regionId string,
    appName string,
    platform int,
) *CreateAppRequest {

	return &CreateAppRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/app",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        AppName: appName,
        Platform: platform,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param appName: 应用名称 (Required)
 * param platform: 部署平台 (Required)
 * param jdsfEnabled: 使用分布式服务框架：0不使用，1使用 (Optional)
 * param desc: 描述 (Optional)
 */
func NewCreateAppRequestWithAllParams(
    regionId string,
    appName string,
    platform int,
    jdsfEnabled *int,
    desc *string,
) *CreateAppRequest {

    return &CreateAppRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/app",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        AppName: appName,
        Platform: platform,
        JdsfEnabled: jdsfEnabled,
        Desc: desc,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateAppRequestWithoutParam() *CreateAppRequest {

    return &CreateAppRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/app",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *CreateAppRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param appName: 应用名称(Required) */
func (r *CreateAppRequest) SetAppName(appName string) {
    r.AppName = appName
}

/* param platform: 部署平台(Required) */
func (r *CreateAppRequest) SetPlatform(platform int) {
    r.Platform = platform
}

/* param jdsfEnabled: 使用分布式服务框架：0不使用，1使用(Optional) */
func (r *CreateAppRequest) SetJdsfEnabled(jdsfEnabled int) {
    r.JdsfEnabled = &jdsfEnabled
}

/* param desc: 描述(Optional) */
func (r *CreateAppRequest) SetDesc(desc string) {
    r.Desc = &desc
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateAppRequest) GetRegionId() string {
    return r.RegionId
}

type CreateAppResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateAppResult `json:"result"`
}

type CreateAppResult struct {
    AppId string `json:"appId"`
}