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
    jdsfadmin "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/jdsfadmin/models"
)

type DescribeTasksRequest struct {

    core.JDCloudRequest

    /* 可用区id  */
    RegionId string `json:"regionId"`

    /* 分页容量,默认100 (Optional) */
    PageSize *int `json:"pageSize"`

    /* 当前要请求的页数 default 1 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 创建人ERP (Optional) */
    UserPin *string `json:"userPin"`

    /* 执行结果（1：成功，0：失败） (Optional) */
    LatestStatus *int `json:"latestStatus"`
}

/*
 * param regionId: 可用区id (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeTasksRequest(
    regionId string,
) *DescribeTasksRequest {

	return &DescribeTasksRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/scheduled_tasks",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: 可用区id (Required)
 * param pageSize: 分页容量,默认100 (Optional)
 * param pageNumber: 当前要请求的页数 default 1 (Optional)
 * param userPin: 创建人ERP (Optional)
 * param latestStatus: 执行结果（1：成功，0：失败） (Optional)
 */
func NewDescribeTasksRequestWithAllParams(
    regionId string,
    pageSize *int,
    pageNumber *int,
    userPin *string,
    latestStatus *int,
) *DescribeTasksRequest {

    return &DescribeTasksRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/scheduled_tasks",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        PageSize: pageSize,
        PageNumber: pageNumber,
        UserPin: userPin,
        LatestStatus: latestStatus,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeTasksRequestWithoutParam() *DescribeTasksRequest {

    return &DescribeTasksRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/scheduled_tasks",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 可用区id(Required) */
func (r *DescribeTasksRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param pageSize: 分页容量,默认100(Optional) */
func (r *DescribeTasksRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

/* param pageNumber: 当前要请求的页数 default 1(Optional) */
func (r *DescribeTasksRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param userPin: 创建人ERP(Optional) */
func (r *DescribeTasksRequest) SetUserPin(userPin string) {
    r.UserPin = &userPin
}

/* param latestStatus: 执行结果（1：成功，0：失败）(Optional) */
func (r *DescribeTasksRequest) SetLatestStatus(latestStatus int) {
    r.LatestStatus = &latestStatus
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeTasksRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeTasksResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeTasksResult `json:"result"`
}

type DescribeTasksResult struct {
    TaskList []jdsfadmin.ScheduledTask `json:"taskList"`
    TotalCount int `json:"totalCount"`
}