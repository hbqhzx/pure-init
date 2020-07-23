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
    bgw "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/bgw/models"
    common "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/common/models"
)

type DescribeAllConnectionsRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 页码, 默认为1, 取值范围：[1,∞), 页码超过总页数时, 显示最后一页 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 分页大小，默认为20，取值范围：[10,100] (Optional) */
    PageSize *int `json:"pageSize"`

    /* connectionIds - ConnectionID列表，支持多个
connectionNames - Connection名称列表,支持多个
partnerCode - 合作伙伴编码,通过调用describePartners接口获取,支持单个
status - Connection状态，支持单个,取值为：Ordering,Rejected,Pending,Confirming,Paying,Active,Deleting,支持单个
types - 类型，取值：jcloud_hosted:托管连接、jcloud_partner:合作伙伴连接、jcloud:自助连接, 支持多个
connectionOwner - 专线所有者的用户pin，支持单个
 (Optional) */
    Filters []common.Filter `json:"filters"`
}

/*
 * param regionId: Region ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeAllConnectionsRequest(
    regionId string,
) *DescribeAllConnectionsRequest {

	return &DescribeAllConnectionsRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/allConnections/",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: Region ID (Required)
 * param pageNumber: 页码, 默认为1, 取值范围：[1,∞), 页码超过总页数时, 显示最后一页 (Optional)
 * param pageSize: 分页大小，默认为20，取值范围：[10,100] (Optional)
 * param filters: connectionIds - ConnectionID列表，支持多个
connectionNames - Connection名称列表,支持多个
partnerCode - 合作伙伴编码,通过调用describePartners接口获取,支持单个
status - Connection状态，支持单个,取值为：Ordering,Rejected,Pending,Confirming,Paying,Active,Deleting,支持单个
types - 类型，取值：jcloud_hosted:托管连接、jcloud_partner:合作伙伴连接、jcloud:自助连接, 支持多个
connectionOwner - 专线所有者的用户pin，支持单个
 (Optional)
 */
func NewDescribeAllConnectionsRequestWithAllParams(
    regionId string,
    pageNumber *int,
    pageSize *int,
    filters []common.Filter,
) *DescribeAllConnectionsRequest {

    return &DescribeAllConnectionsRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/allConnections/",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        PageNumber: pageNumber,
        PageSize: pageSize,
        Filters: filters,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeAllConnectionsRequestWithoutParam() *DescribeAllConnectionsRequest {

    return &DescribeAllConnectionsRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/allConnections/",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *DescribeAllConnectionsRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param pageNumber: 页码, 默认为1, 取值范围：[1,∞), 页码超过总页数时, 显示最后一页(Optional) */
func (r *DescribeAllConnectionsRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param pageSize: 分页大小，默认为20，取值范围：[10,100](Optional) */
func (r *DescribeAllConnectionsRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

/* param filters: connectionIds - ConnectionID列表，支持多个
connectionNames - Connection名称列表,支持多个
partnerCode - 合作伙伴编码,通过调用describePartners接口获取,支持单个
status - Connection状态，支持单个,取值为：Ordering,Rejected,Pending,Confirming,Paying,Active,Deleting,支持单个
types - 类型，取值：jcloud_hosted:托管连接、jcloud_partner:合作伙伴连接、jcloud:自助连接, 支持多个
connectionOwner - 专线所有者的用户pin，支持单个
(Optional) */
func (r *DescribeAllConnectionsRequest) SetFilters(filters []common.Filter) {
    r.Filters = filters
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeAllConnectionsRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeAllConnectionsResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeAllConnectionsResult `json:"result"`
}

type DescribeAllConnectionsResult struct {
    Connections []bgw.OpConnection `json:"connections"`
    TotalCount int `json:"totalCount"`
}