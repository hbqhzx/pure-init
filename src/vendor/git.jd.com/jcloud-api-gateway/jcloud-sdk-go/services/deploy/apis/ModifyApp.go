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

type ModifyAppRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* App Id  */
    AppId string `json:"appId"`

    /* 描述  */
    Desc string `json:"desc"`
}

/*
 * param regionId: 地域ID (Required)
 * param appId: App Id (Required)
 * param desc: 描述 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewModifyAppRequest(
    regionId string,
    appId string,
    desc string,
) *ModifyAppRequest {

	return &ModifyAppRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/app/{appId}",
			Method:  "PUT",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        AppId: appId,
        Desc: desc,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param appId: App Id (Required)
 * param desc: 描述 (Required)
 */
func NewModifyAppRequestWithAllParams(
    regionId string,
    appId string,
    desc string,
) *ModifyAppRequest {

    return &ModifyAppRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/app/{appId}",
            Method:  "PUT",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        AppId: appId,
        Desc: desc,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewModifyAppRequestWithoutParam() *ModifyAppRequest {

    return &ModifyAppRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/app/{appId}",
            Method:  "PUT",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *ModifyAppRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param appId: App Id(Required) */
func (r *ModifyAppRequest) SetAppId(appId string) {
    r.AppId = appId
}

/* param desc: 描述(Required) */
func (r *ModifyAppRequest) SetDesc(desc string) {
    r.Desc = desc
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ModifyAppRequest) GetRegionId() string {
    return r.RegionId
}

type ModifyAppResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ModifyAppResult `json:"result"`
}

type ModifyAppResult struct {
    Result bool `json:"result"`
}