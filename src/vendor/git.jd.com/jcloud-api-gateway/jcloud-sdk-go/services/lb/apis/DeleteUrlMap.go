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

type DeleteUrlMapRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 转发规则组id  */
    UrlMapId string `json:"urlMapId"`
}

/*
 * param regionId: Region ID (Required)
 * param urlMapId: 转发规则组id (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDeleteUrlMapRequest(
    regionId string,
    urlMapId string,
) *DeleteUrlMapRequest {

	return &DeleteUrlMapRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/urlMaps/{urlMapId}",
			Method:  "DELETE",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        UrlMapId: urlMapId,
	}
}

/*
 * param regionId: Region ID (Required)
 * param urlMapId: 转发规则组id (Required)
 */
func NewDeleteUrlMapRequestWithAllParams(
    regionId string,
    urlMapId string,
) *DeleteUrlMapRequest {

    return &DeleteUrlMapRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/urlMaps/{urlMapId}",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        UrlMapId: urlMapId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDeleteUrlMapRequestWithoutParam() *DeleteUrlMapRequest {

    return &DeleteUrlMapRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/urlMaps/{urlMapId}",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *DeleteUrlMapRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param urlMapId: 转发规则组id(Required) */
func (r *DeleteUrlMapRequest) SetUrlMapId(urlMapId string) {
    r.UrlMapId = urlMapId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DeleteUrlMapRequest) GetRegionId() string {
    return r.RegionId
}

type DeleteUrlMapResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DeleteUrlMapResult `json:"result"`
}

type DeleteUrlMapResult struct {
}