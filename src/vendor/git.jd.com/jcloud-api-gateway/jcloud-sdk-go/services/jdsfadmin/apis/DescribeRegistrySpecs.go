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

type DescribeRegistrySpecsRequest struct {

    core.JDCloudRequest

    /* 可用区id  */
    RegionId string `json:"regionId"`

    /* 分页信息要请求的页数 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 分页信息要请求每页数据的条数 (Optional) */
    PageSize *int `json:"pageSize"`
}

/*
 * param regionId: 可用区id (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeRegistrySpecsRequest(
    regionId string,
) *DescribeRegistrySpecsRequest {

	return &DescribeRegistrySpecsRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/registryspecs",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: 可用区id (Required)
 * param pageNumber: 分页信息要请求的页数 (Optional)
 * param pageSize: 分页信息要请求每页数据的条数 (Optional)
 */
func NewDescribeRegistrySpecsRequestWithAllParams(
    regionId string,
    pageNumber *int,
    pageSize *int,
) *DescribeRegistrySpecsRequest {

    return &DescribeRegistrySpecsRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/registryspecs",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        PageNumber: pageNumber,
        PageSize: pageSize,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeRegistrySpecsRequestWithoutParam() *DescribeRegistrySpecsRequest {

    return &DescribeRegistrySpecsRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/registryspecs",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 可用区id(Required) */
func (r *DescribeRegistrySpecsRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param pageNumber: 分页信息要请求的页数(Optional) */
func (r *DescribeRegistrySpecsRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param pageSize: 分页信息要请求每页数据的条数(Optional) */
func (r *DescribeRegistrySpecsRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeRegistrySpecsRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeRegistrySpecsResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeRegistrySpecsResult `json:"result"`
}

type DescribeRegistrySpecsResult struct {
    RegistrySpecInfos []jdsfadmin.RegistrySpec `json:"registrySpecInfos"`
    TotalCount int64 `json:"totalCount"`
}