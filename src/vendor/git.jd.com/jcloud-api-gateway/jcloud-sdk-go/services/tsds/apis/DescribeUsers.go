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

type DescribeUsersRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* Instance ID  */
    InstanceId string `json:"instanceId"`

    /* The page number of instance user list，the value range is [1,1000]，when the page number execceds the total number of pages，the last page is displayed. (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* The number of instance user per page, the value range is [10/20/30/50/100] (Optional) */
    PageSize *int `json:"pageSize"`
}

/*
 * param regionId: Region ID (Required)
 * param instanceId: Instance ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeUsersRequest(
    regionId string,
    instanceId string,
) *DescribeUsersRequest {

	return &DescribeUsersRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}/users",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceId: instanceId,
	}
}

/*
 * param regionId: Region ID (Required)
 * param instanceId: Instance ID (Required)
 * param pageNumber: The page number of instance user list，the value range is [1,1000]，when the page number execceds the total number of pages，the last page is displayed. (Optional)
 * param pageSize: The number of instance user per page, the value range is [10/20/30/50/100] (Optional)
 */
func NewDescribeUsersRequestWithAllParams(
    regionId string,
    instanceId string,
    pageNumber *int,
    pageSize *int,
) *DescribeUsersRequest {

    return &DescribeUsersRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/users",
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
func NewDescribeUsersRequestWithoutParam() *DescribeUsersRequest {

    return &DescribeUsersRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/users",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *DescribeUsersRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceId: Instance ID(Required) */
func (r *DescribeUsersRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

/* param pageNumber: The page number of instance user list，the value range is [1,1000]，when the page number execceds the total number of pages，the last page is displayed.(Optional) */
func (r *DescribeUsersRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param pageSize: The number of instance user per page, the value range is [10/20/30/50/100](Optional) */
func (r *DescribeUsersRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeUsersRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeUsersResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeUsersResult `json:"result"`
}

type DescribeUsersResult struct {
    Users []tsds.InstanceUser `json:"users"`
    TotalCount int `json:"totalCount"`
}