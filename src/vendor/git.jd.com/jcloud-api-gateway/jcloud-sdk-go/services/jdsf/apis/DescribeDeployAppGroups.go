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
    jdsf "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/jdsf/models"
)

type DescribeDeployAppGroupsRequest struct {

    core.JDCloudRequest

    /* 可用区id  */
    RegionId string `json:"regionId"`

    /* 云部署应用 id (Optional) */
    AppId *string `json:"appId"`
}

/*
 * param regionId: 可用区id (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeDeployAppGroupsRequest(
    regionId string,
) *DescribeDeployAppGroupsRequest {

	return &DescribeDeployAppGroupsRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/deployapps/{appId}/groups",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: 可用区id (Required)
 * param appId: 云部署应用 id (Optional)
 */
func NewDescribeDeployAppGroupsRequestWithAllParams(
    regionId string,
    appId *string,
) *DescribeDeployAppGroupsRequest {

    return &DescribeDeployAppGroupsRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/deployapps/{appId}/groups",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        AppId: appId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeDeployAppGroupsRequestWithoutParam() *DescribeDeployAppGroupsRequest {

    return &DescribeDeployAppGroupsRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/deployapps/{appId}/groups",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 可用区id(Required) */
func (r *DescribeDeployAppGroupsRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param appId: 云部署应用 id(Optional) */
func (r *DescribeDeployAppGroupsRequest) SetAppId(appId string) {
    r.AppId = &appId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeDeployAppGroupsRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeDeployAppGroupsResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeDeployAppGroupsResult `json:"result"`
}

type DescribeDeployAppGroupsResult struct {
    Groups []jdsf.Group `json:"groups"`
    TotalCount int64 `json:"totalCount"`
}