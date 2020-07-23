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
    es "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/es/models"
    common "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/common/models"
)

type DescribeSnapshotsRequest struct {

    core.JDCloudRequest

    /* regionId  */
    RegionId string `json:"regionId"`

    /* 实例ID  */
    InstanceId string `json:"instanceId"`

    /* 页码 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 分页大小 (Optional) */
    PageSize *int `json:"pageSize"`

    /* snapshotId - 快照编号，模糊匹配
status - 快照状态，精确匹配，支持多个(Available：可用、Unavailable：不可用Creating：创建中)
type - 快照类型，精确匹配，支持单个(Auto:自动、Manual:人工）
 (Optional) */
    Filters []common.Filter `json:"filters"`
}

/*
 * param regionId: regionId (Required)
 * param instanceId: 实例ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeSnapshotsRequest(
    regionId string,
    instanceId string,
) *DescribeSnapshotsRequest {

	return &DescribeSnapshotsRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}/snapshots",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceId: instanceId,
	}
}

/*
 * param regionId: regionId (Required)
 * param instanceId: 实例ID (Required)
 * param pageNumber: 页码 (Optional)
 * param pageSize: 分页大小 (Optional)
 * param filters: snapshotId - 快照编号，模糊匹配
status - 快照状态，精确匹配，支持多个(Available：可用、Unavailable：不可用Creating：创建中)
type - 快照类型，精确匹配，支持单个(Auto:自动、Manual:人工）
 (Optional)
 */
func NewDescribeSnapshotsRequestWithAllParams(
    regionId string,
    instanceId string,
    pageNumber *int,
    pageSize *int,
    filters []common.Filter,
) *DescribeSnapshotsRequest {

    return &DescribeSnapshotsRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/snapshots",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InstanceId: instanceId,
        PageNumber: pageNumber,
        PageSize: pageSize,
        Filters: filters,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeSnapshotsRequestWithoutParam() *DescribeSnapshotsRequest {

    return &DescribeSnapshotsRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/snapshots",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: regionId(Required) */
func (r *DescribeSnapshotsRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceId: 实例ID(Required) */
func (r *DescribeSnapshotsRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

/* param pageNumber: 页码(Optional) */
func (r *DescribeSnapshotsRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param pageSize: 分页大小(Optional) */
func (r *DescribeSnapshotsRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

/* param filters: snapshotId - 快照编号，模糊匹配
status - 快照状态，精确匹配，支持多个(Available：可用、Unavailable：不可用Creating：创建中)
type - 快照类型，精确匹配，支持单个(Auto:自动、Manual:人工）
(Optional) */
func (r *DescribeSnapshotsRequest) SetFilters(filters []common.Filter) {
    r.Filters = filters
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeSnapshotsRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeSnapshotsResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeSnapshotsResult `json:"result"`
}

type DescribeSnapshotsResult struct {
    Snapshots []es.Snapshot `json:"snapshots"`
    TotalCount int `json:"totalCount"`
}