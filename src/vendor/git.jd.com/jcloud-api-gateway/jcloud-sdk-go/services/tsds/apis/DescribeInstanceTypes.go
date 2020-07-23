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
    tsds "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/tsds/models"
)

type DescribeInstanceTypesRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* The page number of instance list，the value range is [1,1000]，when the page number execceds the total number of pages，the last page is displayed. (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* The number of instance per page, the value range is [10/20/30/50/100] (Optional) */
    PageSize *int `json:"pageSize"`
}

/*
 * param regionId: Region ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeInstanceTypesRequest(
    regionId string,
) *DescribeInstanceTypesRequest {

	return &DescribeInstanceTypesRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instance_types",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: Region ID (Required)
 * param pageNumber: The page number of instance list，the value range is [1,1000]，when the page number execceds the total number of pages，the last page is displayed. (Optional)
 * param pageSize: The number of instance per page, the value range is [10/20/30/50/100] (Optional)
 */
func NewDescribeInstanceTypesRequestWithAllParams(
    regionId string,
    pageNumber *int,
    pageSize *int,
) *DescribeInstanceTypesRequest {

    return &DescribeInstanceTypesRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instance_types",
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
func NewDescribeInstanceTypesRequestWithoutParam() *DescribeInstanceTypesRequest {

    return &DescribeInstanceTypesRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instance_types",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *DescribeInstanceTypesRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param pageNumber: The page number of instance list，the value range is [1,1000]，when the page number execceds the total number of pages，the last page is displayed.(Optional) */
func (r *DescribeInstanceTypesRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param pageSize: The number of instance per page, the value range is [10/20/30/50/100](Optional) */
func (r *DescribeInstanceTypesRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeInstanceTypesRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeInstanceTypesResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeInstanceTypesResult `json:"result"`
}

type DescribeInstanceTypesResult struct {
    Instances []tsds.InstanceType `json:"instances"`
    TotalCount int `json:"totalCount"`
}