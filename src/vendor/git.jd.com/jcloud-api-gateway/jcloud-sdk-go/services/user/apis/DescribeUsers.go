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
    user "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/user/models"
)

type DescribeUsersRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 用户信息  */
    UserReqVo *user.UserReqVo `json:"userReqVo"`
}

/*
 * param regionId: Region ID (Required)
 * param userReqVo: 用户信息 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeUsersRequest(
    regionId string,
    userReqVo *user.UserReqVo,
) *DescribeUsersRequest {

	return &DescribeUsersRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/user:describeUsers",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        UserReqVo: userReqVo,
	}
}

/*
 * param regionId: Region ID (Required)
 * param userReqVo: 用户信息 (Required)
 */
func NewDescribeUsersRequestWithAllParams(
    regionId string,
    userReqVo *user.UserReqVo,
) *DescribeUsersRequest {

    return &DescribeUsersRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/user:describeUsers",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        UserReqVo: userReqVo,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeUsersRequestWithoutParam() *DescribeUsersRequest {

    return &DescribeUsersRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/user:describeUsers",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *DescribeUsersRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param userReqVo: 用户信息(Required) */
func (r *DescribeUsersRequest) SetUserReqVo(userReqVo *user.UserReqVo) {
    r.UserReqVo = userReqVo
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeUsersRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeUsersResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeUsersResult `json:"result"`
}

type DescribeUsersResult struct {
    Data []user.UserResVo `json:"data"`
}