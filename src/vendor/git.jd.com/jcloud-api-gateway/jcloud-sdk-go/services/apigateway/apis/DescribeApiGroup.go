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
    apigateway "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/apigateway/models"
)

type DescribeApiGroupRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 分组ID  */
    ApiGroupId string `json:"apiGroupId"`
}

/*
 * param regionId: 地域ID (Required)
 * param apiGroupId: 分组ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeApiGroupRequest(
    regionId string,
    apiGroupId string,
) *DescribeApiGroupRequest {

	return &DescribeApiGroupRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/apiGroups/{apiGroupId}",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        ApiGroupId: apiGroupId,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param apiGroupId: 分组ID (Required)
 */
func NewDescribeApiGroupRequestWithAllParams(
    regionId string,
    apiGroupId string,
) *DescribeApiGroupRequest {

    return &DescribeApiGroupRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/apiGroups/{apiGroupId}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        ApiGroupId: apiGroupId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeApiGroupRequestWithoutParam() *DescribeApiGroupRequest {

    return &DescribeApiGroupRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/apiGroups/{apiGroupId}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *DescribeApiGroupRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param apiGroupId: 分组ID(Required) */
func (r *DescribeApiGroupRequest) SetApiGroupId(apiGroupId string) {
    r.ApiGroupId = apiGroupId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeApiGroupRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeApiGroupResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeApiGroupResult `json:"result"`
}

type DescribeApiGroupResult struct {
    ApiGroup apigateway.ApiGroup `json:"apiGroup"`
}