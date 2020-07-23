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

type DescribeDeployAppsRequest struct {

    core.JDCloudRequest

    /* 可用区id  */
    RegionId string `json:"regionId"`

    /* 应用名称 支持模糊查询 (Optional) */
    AppName *string `json:"appName"`

    /* 分页大小 (Optional) */
    PageSize *string `json:"pageSize"`

    /* 当前请求页码 (Optional) */
    PageNumber *string `json:"pageNumber"`
}

/*
 * param regionId: 可用区id (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeDeployAppsRequest(
    regionId string,
) *DescribeDeployAppsRequest {

	return &DescribeDeployAppsRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/deployapps",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: 可用区id (Required)
 * param appName: 应用名称 支持模糊查询 (Optional)
 * param pageSize: 分页大小 (Optional)
 * param pageNumber: 当前请求页码 (Optional)
 */
func NewDescribeDeployAppsRequestWithAllParams(
    regionId string,
    appName *string,
    pageSize *string,
    pageNumber *string,
) *DescribeDeployAppsRequest {

    return &DescribeDeployAppsRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/deployapps",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        AppName: appName,
        PageSize: pageSize,
        PageNumber: pageNumber,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeDeployAppsRequestWithoutParam() *DescribeDeployAppsRequest {

    return &DescribeDeployAppsRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/deployapps",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 可用区id(Required) */
func (r *DescribeDeployAppsRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param appName: 应用名称 支持模糊查询(Optional) */
func (r *DescribeDeployAppsRequest) SetAppName(appName string) {
    r.AppName = &appName
}

/* param pageSize: 分页大小(Optional) */
func (r *DescribeDeployAppsRequest) SetPageSize(pageSize string) {
    r.PageSize = &pageSize
}

/* param pageNumber: 当前请求页码(Optional) */
func (r *DescribeDeployAppsRequest) SetPageNumber(pageNumber string) {
    r.PageNumber = &pageNumber
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeDeployAppsRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeDeployAppsResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeDeployAppsResult `json:"result"`
}

type DescribeDeployAppsResult struct {
    DeployApps []jdsf.DeployApp `json:"deployApps"`
    TotalCount int64 `json:"totalCount"`
}