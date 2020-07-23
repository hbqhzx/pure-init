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
    zfs "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/zfs/models"
    common "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/common/models"
)

type DescribeMountTargetsRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 页码, 默认为1, 取值范围：[1,∞) (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 分页大小，默认为20，取值范围：[10,100] (Optional) */
    PageSize *int `json:"pageSize"`

    /* Tag筛选条件 (Optional) */
    Tags []zfs.TagFilter `json:"tags"`

    /* fileSystemId - 文件系统ID，精确匹配，支持多个
mountTargetId - 挂载目标ID，精确匹配，支持多个
 (Optional) */
    Filters []common.Filter `json:"filters"`
}

/*
 * param regionId: 地域ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeMountTargetsRequest(
    regionId string,
) *DescribeMountTargetsRequest {

	return &DescribeMountTargetsRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/mountTargets",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param pageNumber: 页码, 默认为1, 取值范围：[1,∞) (Optional)
 * param pageSize: 分页大小，默认为20，取值范围：[10,100] (Optional)
 * param tags: Tag筛选条件 (Optional)
 * param filters: fileSystemId - 文件系统ID，精确匹配，支持多个
mountTargetId - 挂载目标ID，精确匹配，支持多个
 (Optional)
 */
func NewDescribeMountTargetsRequestWithAllParams(
    regionId string,
    pageNumber *int,
    pageSize *int,
    tags []zfs.TagFilter,
    filters []common.Filter,
) *DescribeMountTargetsRequest {

    return &DescribeMountTargetsRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/mountTargets",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        PageNumber: pageNumber,
        PageSize: pageSize,
        Tags: tags,
        Filters: filters,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeMountTargetsRequestWithoutParam() *DescribeMountTargetsRequest {

    return &DescribeMountTargetsRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/mountTargets",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *DescribeMountTargetsRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param pageNumber: 页码, 默认为1, 取值范围：[1,∞)(Optional) */
func (r *DescribeMountTargetsRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param pageSize: 分页大小，默认为20，取值范围：[10,100](Optional) */
func (r *DescribeMountTargetsRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

/* param tags: Tag筛选条件(Optional) */
func (r *DescribeMountTargetsRequest) SetTags(tags []zfs.TagFilter) {
    r.Tags = tags
}

/* param filters: fileSystemId - 文件系统ID，精确匹配，支持多个
mountTargetId - 挂载目标ID，精确匹配，支持多个
(Optional) */
func (r *DescribeMountTargetsRequest) SetFilters(filters []common.Filter) {
    r.Filters = filters
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeMountTargetsRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeMountTargetsResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeMountTargetsResult `json:"result"`
}

type DescribeMountTargetsResult struct {
    MountTargets []zfs.MountTarget `json:"mountTargets"`
    TotalCount int `json:"totalCount"`
}