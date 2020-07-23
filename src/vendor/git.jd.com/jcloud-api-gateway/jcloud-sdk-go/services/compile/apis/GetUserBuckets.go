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

type GetUserBucketsRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 查询用户指定区域的oss存储空间  */
    UserRegionId string `json:"userRegionId"`
}

/*
 * param regionId: Region ID (Required)
 * param userRegionId: 查询用户指定区域的oss存储空间 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewGetUserBucketsRequest(
    regionId string,
    userRegionId string,
) *GetUserBucketsRequest {

	return &GetUserBucketsRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/jobs/user/buckets",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        UserRegionId: userRegionId,
	}
}

/*
 * param regionId: Region ID (Required)
 * param userRegionId: 查询用户指定区域的oss存储空间 (Required)
 */
func NewGetUserBucketsRequestWithAllParams(
    regionId string,
    userRegionId string,
) *GetUserBucketsRequest {

    return &GetUserBucketsRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/jobs/user/buckets",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        UserRegionId: userRegionId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewGetUserBucketsRequestWithoutParam() *GetUserBucketsRequest {

    return &GetUserBucketsRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/jobs/user/buckets",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *GetUserBucketsRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param userRegionId: 查询用户指定区域的oss存储空间(Required) */
func (r *GetUserBucketsRequest) SetUserRegionId(userRegionId string) {
    r.UserRegionId = userRegionId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r GetUserBucketsRequest) GetRegionId() string {
    return r.RegionId
}

type GetUserBucketsResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result GetUserBucketsResult `json:"result"`
}

type GetUserBucketsResult struct {
    Buckets []string `json:"buckets"`
}