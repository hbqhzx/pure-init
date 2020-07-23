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

type DisableApplicationServiceUrlRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 资源id  */
    DirectoryId string `json:"directoryId"`

    /* 应用名称  */
    AppName string `json:"appName"`
}

/*
 * param regionId: 地域ID (Required)
 * param directoryId: 资源id (Required)
 * param appName: 应用名称 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDisableApplicationServiceUrlRequest(
    regionId string,
    directoryId string,
    appName string,
) *DisableApplicationServiceUrlRequest {

	return &DisableApplicationServiceUrlRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/directory/{directoryId}/applicationService:disableApplicationServiceUrl",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        DirectoryId: directoryId,
        AppName: appName,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param directoryId: 资源id (Required)
 * param appName: 应用名称 (Required)
 */
func NewDisableApplicationServiceUrlRequestWithAllParams(
    regionId string,
    directoryId string,
    appName string,
) *DisableApplicationServiceUrlRequest {

    return &DisableApplicationServiceUrlRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/directory/{directoryId}/applicationService:disableApplicationServiceUrl",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        DirectoryId: directoryId,
        AppName: appName,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDisableApplicationServiceUrlRequestWithoutParam() *DisableApplicationServiceUrlRequest {

    return &DisableApplicationServiceUrlRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/directory/{directoryId}/applicationService:disableApplicationServiceUrl",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *DisableApplicationServiceUrlRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param directoryId: 资源id(Required) */
func (r *DisableApplicationServiceUrlRequest) SetDirectoryId(directoryId string) {
    r.DirectoryId = directoryId
}

/* param appName: 应用名称(Required) */
func (r *DisableApplicationServiceUrlRequest) SetAppName(appName string) {
    r.AppName = appName
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DisableApplicationServiceUrlRequest) GetRegionId() string {
    return r.RegionId
}

type DisableApplicationServiceUrlResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DisableApplicationServiceUrlResult `json:"result"`
}

type DisableApplicationServiceUrlResult struct {
}