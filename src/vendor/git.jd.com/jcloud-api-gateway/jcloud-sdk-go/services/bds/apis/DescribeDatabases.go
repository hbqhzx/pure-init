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
    bds "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/bds/models"
)

type DescribeDatabasesRequest struct {

    core.JDCloudRequest

    /* 区域代码  */
    RegionId string `json:"regionId"`

    /* 实例ID  */
    InstanceId string `json:"instanceId"`

    /* 显示数据的页码，取值范围：[1,1000)，页码超过总页数时，显示最后一页，用于查询列表的接口 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 每页显示的数据条数，取值范围：10/20/30/50/100 (Optional) */
    PageSize *int `json:"pageSize"`
}

/*
 * param regionId: 区域代码 (Required)
 * param instanceId: 实例ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeDatabasesRequest(
    regionId string,
    instanceId string,
) *DescribeDatabasesRequest {

	return &DescribeDatabasesRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}/databases",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceId: instanceId,
	}
}

/*
 * param regionId: 区域代码 (Required)
 * param instanceId: 实例ID (Required)
 * param pageNumber: 显示数据的页码，取值范围：[1,1000)，页码超过总页数时，显示最后一页，用于查询列表的接口 (Optional)
 * param pageSize: 每页显示的数据条数，取值范围：10/20/30/50/100 (Optional)
 */
func NewDescribeDatabasesRequestWithAllParams(
    regionId string,
    instanceId string,
    pageNumber *int,
    pageSize *int,
) *DescribeDatabasesRequest {

    return &DescribeDatabasesRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/databases",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InstanceId: instanceId,
        PageNumber: pageNumber,
        PageSize: pageSize,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeDatabasesRequestWithoutParam() *DescribeDatabasesRequest {

    return &DescribeDatabasesRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/databases",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 区域代码(Required) */
func (r *DescribeDatabasesRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceId: 实例ID(Required) */
func (r *DescribeDatabasesRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

/* param pageNumber: 显示数据的页码，取值范围：[1,1000)，页码超过总页数时，显示最后一页，用于查询列表的接口(Optional) */
func (r *DescribeDatabasesRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param pageSize: 每页显示的数据条数，取值范围：10/20/30/50/100(Optional) */
func (r *DescribeDatabasesRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeDatabasesRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeDatabasesResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeDatabasesResult `json:"result"`
}

type DescribeDatabasesResult struct {
    Databases []bds.Database `json:"databases"`
    TotalCount int `json:"totalCount"`
}