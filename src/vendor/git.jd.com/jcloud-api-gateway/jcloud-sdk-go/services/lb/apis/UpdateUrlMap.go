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

type UpdateUrlMapRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 转发规则组id  */
    UrlMapId string `json:"urlMapId"`

    /* 转发规则组描述,允许输入UTF-8编码下的全部字符，不超过256字符 (Optional) */
    Description *string `json:"description"`

    /* 转发规则组名称，同一个负载均衡下，名称不能重复,只允许输入中文、数字、大小写字母、英文下划线“_”及中划线“-”，不允许为空且不超过32字符 (Optional) */
    UrlMapName *string `json:"urlMapName"`
}

/*
 * param regionId: Region ID (Required)
 * param urlMapId: 转发规则组id (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewUpdateUrlMapRequest(
    regionId string,
    urlMapId string,
) *UpdateUrlMapRequest {

	return &UpdateUrlMapRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/urlMaps/{urlMapId}",
			Method:  "PATCH",
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
 * param description: 转发规则组描述,允许输入UTF-8编码下的全部字符，不超过256字符 (Optional)
 * param urlMapName: 转发规则组名称，同一个负载均衡下，名称不能重复,只允许输入中文、数字、大小写字母、英文下划线“_”及中划线“-”，不允许为空且不超过32字符 (Optional)
 */
func NewUpdateUrlMapRequestWithAllParams(
    regionId string,
    urlMapId string,
    description *string,
    urlMapName *string,
) *UpdateUrlMapRequest {

    return &UpdateUrlMapRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/urlMaps/{urlMapId}",
            Method:  "PATCH",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        UrlMapId: urlMapId,
        Description: description,
        UrlMapName: urlMapName,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewUpdateUrlMapRequestWithoutParam() *UpdateUrlMapRequest {

    return &UpdateUrlMapRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/urlMaps/{urlMapId}",
            Method:  "PATCH",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *UpdateUrlMapRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param urlMapId: 转发规则组id(Required) */
func (r *UpdateUrlMapRequest) SetUrlMapId(urlMapId string) {
    r.UrlMapId = urlMapId
}

/* param description: 转发规则组描述,允许输入UTF-8编码下的全部字符，不超过256字符(Optional) */
func (r *UpdateUrlMapRequest) SetDescription(description string) {
    r.Description = &description
}

/* param urlMapName: 转发规则组名称，同一个负载均衡下，名称不能重复,只允许输入中文、数字、大小写字母、英文下划线“_”及中划线“-”，不允许为空且不超过32字符(Optional) */
func (r *UpdateUrlMapRequest) SetUrlMapName(urlMapName string) {
    r.UrlMapName = &urlMapName
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r UpdateUrlMapRequest) GetRegionId() string {
    return r.RegionId
}

type UpdateUrlMapResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result UpdateUrlMapResult `json:"result"`
}

type UpdateUrlMapResult struct {
}