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

type GetExternalLinkRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 构建任务uuid  */
    Id string `json:"id"`

    /* 过期时间，单位秒， 默认1800秒 (Optional) */
    Expires *int `json:"expires"`
}

/*
 * param regionId: Region ID (Required)
 * param id: 构建任务uuid (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewGetExternalLinkRequest(
    regionId string,
    id string,
) *GetExternalLinkRequest {

	return &GetExternalLinkRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/builds/{id}/externalLink",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        Id: id,
	}
}

/*
 * param regionId: Region ID (Required)
 * param id: 构建任务uuid (Required)
 * param expires: 过期时间，单位秒， 默认1800秒 (Optional)
 */
func NewGetExternalLinkRequestWithAllParams(
    regionId string,
    id string,
    expires *int,
) *GetExternalLinkRequest {

    return &GetExternalLinkRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/builds/{id}/externalLink",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        Id: id,
        Expires: expires,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewGetExternalLinkRequestWithoutParam() *GetExternalLinkRequest {

    return &GetExternalLinkRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/builds/{id}/externalLink",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *GetExternalLinkRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param id: 构建任务uuid(Required) */
func (r *GetExternalLinkRequest) SetId(id string) {
    r.Id = id
}

/* param expires: 过期时间，单位秒， 默认1800秒(Optional) */
func (r *GetExternalLinkRequest) SetExpires(expires int) {
    r.Expires = &expires
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r GetExternalLinkRequest) GetRegionId() string {
    return r.RegionId
}

type GetExternalLinkResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result GetExternalLinkResult `json:"result"`
}

type GetExternalLinkResult struct {
    Url string `json:"url"`
    Token string `json:"token"`
}