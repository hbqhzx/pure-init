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
    yd2 "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/yd2/models"
)

type QueryVmInstancesRequest struct {

    core.JDCloudRequest

    /* Region ID （cn-north-1：华北-北京；cn-east-1：华东-宿迁；cn-east-2：华东-上海；cn-south-1：华南-广州）  */
    RegionId string `json:"regionId"`

    /* 页码；默认为1 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 分页大小；默认为20；取值范围[10, 100] (Optional) */
    PageSize *int `json:"pageSize"`
}

/*
 * param regionId: Region ID （cn-north-1：华北-北京；cn-east-1：华东-宿迁；cn-east-2：华东-上海；cn-south-1：华南-广州） (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewQueryVmInstancesRequest(
    regionId string,
) *QueryVmInstancesRequest {

	return &QueryVmInstancesRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/vmInstances",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: Region ID （cn-north-1：华北-北京；cn-east-1：华东-宿迁；cn-east-2：华东-上海；cn-south-1：华南-广州） (Required)
 * param pageNumber: 页码；默认为1 (Optional)
 * param pageSize: 分页大小；默认为20；取值范围[10, 100] (Optional)
 */
func NewQueryVmInstancesRequestWithAllParams(
    regionId string,
    pageNumber *int,
    pageSize *int,
) *QueryVmInstancesRequest {

    return &QueryVmInstancesRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/vmInstances",
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
func NewQueryVmInstancesRequestWithoutParam() *QueryVmInstancesRequest {

    return &QueryVmInstancesRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/vmInstances",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID （cn-north-1：华北-北京；cn-east-1：华东-宿迁；cn-east-2：华东-上海；cn-south-1：华南-广州）(Required) */
func (r *QueryVmInstancesRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param pageNumber: 页码；默认为1(Optional) */
func (r *QueryVmInstancesRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param pageSize: 分页大小；默认为20；取值范围[10, 100](Optional) */
func (r *QueryVmInstancesRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r QueryVmInstancesRequest) GetRegionId() string {
    return r.RegionId
}

type QueryVmInstancesResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result QueryVmInstancesResult `json:"result"`
}

type QueryVmInstancesResult struct {
    Instances []yd2.Instance `json:"instances"`
    TotalCount int `json:"totalCount"`
}